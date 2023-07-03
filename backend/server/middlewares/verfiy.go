//go:build verify

package middlewares

import (
	"encoding/json"
	"fmt"
	"github.com/TakasakiApps/Narravo/backend/internal/entity"
	"github.com/TakasakiApps/Narravo/backend/internal/global"
	"github.com/TakasakiApps/Narravo/backend/internal/types"
	"github.com/TakasakiApps/Narravo/backend/internal/utils"
	"github.com/gin-gonic/gin"
	"github.com/ohanakogo/exceptiongo"
	"github.com/ohanakogo/exceptiongo/pkg/etype"
	"net/http"
	"strconv"
	"strings"
)

// DataExpireTime If data is sent for processing, and it takes more than 5000 milliseconds, it will be considered as expired data.
// Therefore, the data will be discarded without any further processing of the request.
const DataExpireTime = 10000

const KeyDict = "abcdefghijklmnopqrstuvwxyz0123456789"

func getVerificationComponent() gin.HandlerFunc {
	return func(c *gin.Context) {
		timeStampMilli := utils.GetCurrentTimestampMilli()
		contextUtil := utils.NewContextUtil(c)

		defer exceptiongo.NewExceptionHandler(func(exception *etype.Exception) {
			switch exception.Type() {
			default:
				contextUtil.JsonException(http.StatusBadRequest, exception)
				exception.PrintStackTrace()
				c.Abort()
			}
		}).Deploy()

		verifyBase := func(data string, integrityKey string, requestTimestamp string) {
			requestTimestampInt, _ := strconv.ParseInt(requestTimestamp, 10, 64)
			if timeStampMilli-requestTimestampInt > DataExpireTime {
				exceptiongo.QuickThrowMsg[types.RequestDataExpiredException](
					fmt.Sprintf("Request data expired %v milliseconds ago and has been discarded", DataExpireTime),
				)
			}

			requestTimestampKey := utils.MD5(requestTimestamp)
			aesCipher, _ := utils.NewAesCipher(requestTimestampKey)

			dataKey, err := aesCipher.Decrypt(integrityKey)
			if err != nil || !func() bool {
				for _, char := range dataKey {
					if !strings.Contains(KeyDict, string(char)) {
						return false
					}
				}
				return true
			}() {
				exceptiongo.QuickThrowMsg[types.RuntimeException]("failed to decrypt data, request may sent by unofficial client")
			}

			global.GetLogger().Infof("The data sent from %vms ago is valid", timeStampMilli-requestTimestampInt)

			// Initialize an empty map for dataJsonResult
			dataJsonResult := make(map[string]any)

			// Create a new AES cipher using the provided dataKey
			aesCipher, err = utils.NewAesCipher(dataKey)
			exceptiongo.QuickThrow[types.RuntimeException](err)

			// Decrypt the data using the created cipher
			decrypt, err := aesCipher.Decrypt(data)
			exceptiongo.QuickThrow[types.DecryptFailedException](err)

			// Unmarshal the decrypted data into the dataJsonResult map
			err = json.Unmarshal([]byte(decrypt), &dataJsonResult)
			exceptiongo.QuickThrow[types.JsonUnmarshalFailedException](err)

			// Set the dataJsonResult map as the value of the "data" key in the context
			utils.PutData(c, dataJsonResult)

			// Add next() trigger to make data key available
			c.Next()
		}

		verifyData := entity.VerifyData{}
		switch {
		case c.Request.Method == http.MethodGet || strings.Contains(c.FullPath(), "/assets/upload"):
			verifyData.Data = c.Query("data")
			verifyData.IntegrityKey = c.Query("integrityKey")
			verifyData.Timestamp = c.Query("timestamp")
		case c.Request.Method == http.MethodPost:
			// Bind the JSON data from the request body to the verifyData struct
			err := c.ShouldBindJSON(&verifyData)
			exceptiongo.QuickThrow[types.JsonUnmarshalFailedException](err)
		}
		verifyBase(verifyData.Data, verifyData.IntegrityKey, verifyData.Timestamp)
	}
}
