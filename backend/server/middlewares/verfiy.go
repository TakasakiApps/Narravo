//go:build verify

package middlewares

import (
	"encoding/json"
	"fmt"
	"github.com/TakasakiApps/Narravo/backend/internal/enitity"
	"github.com/TakasakiApps/Narravo/backend/internal/types"
	"github.com/TakasakiApps/Narravo/backend/internal/utils"
	"github.com/gin-gonic/gin"
	"github.com/gookit/goutil/strutil"
	"github.com/gookit/slog"
	"github.com/ohanakogo/exceptiongo"
	"github.com/ohanakogo/exceptiongo/pkg/etype"
	"net/http"
)

// DataExpireTime If data is sent for processing, and it takes more than 5000 milliseconds, it will be considered as expired data.
// Therefore, the data will be discarded without any further processing of the request.
const DataExpireTime = 5000

func verify() gin.HandlerFunc {
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

		verifyBase := func(data string, token string) {
			// Store key to decrypt data
			var dataKey string

			// Iterate over a range of timestamps, starting from the current timestamp in milliseconds
			// and going back by DataExpireTime milliseconds
			for i := timeStampMilli; i >= timeStampMilli-DataExpireTime; i-- {
				// Generate a MD5 key based on the current timestamp
				key := utils.MD5(fmt.Sprintf("%v", i))

				// Create a new AES cipher using the key
				aesCipher, _ := utils.NewAesCipher(key)

				// Decrypt the verification data using the cipher
				decrypt, err := aesCipher.Decrypt(token)
				if err != nil {
					continue
				}

				slog.Infof("The data sent %v ago is valid", timeStampMilli-i)

				// If the data was successfully decrypted, store the decryption key
				dataKey = decrypt
			}

			// Check if the dataKey is empty
			if strutil.IsEmpty(dataKey) {
				// If it's empty, create a new RequestDataExpiredException and return error response to client
				exceptiongo.QuickThrowMsg[types.RequestDataExpiredException](
					fmt.Sprintf("Request data expired %v milliseconds ago and has been discarded", DataExpireTime),
				)
			}

			// Initialize an empty map for dataJsonResult
			dataJsonResult := make(map[string]any)

			// Create a new AES cipher using the provided dataKey
			aesCipher, err := utils.NewAesCipher(dataKey)
			exceptiongo.QuickThrow[types.RuntimeException](err)

			// Decrypt the data using the created cipher
			decrypt, err := aesCipher.Decrypt(data)
			exceptiongo.QuickThrow[types.DecryptFailedException](err)

			// Unmarshal the decrypted data into the dataJsonResult map
			err = json.Unmarshal([]byte(decrypt), &dataJsonResult)
			exceptiongo.QuickThrow[types.JsonUnmarshalFailedException](err)

			// Set the dataJsonResult map as the value of the "data" key in the context
			c.Set("data", dataJsonResult)

		}

		switch c.Request.Method {
		case http.MethodPost:
			// TODO: verify post
			verifyData := enitity.VerifyData{}
			// Bind the JSON data from the request body to the verifyData struct
			err := c.ShouldBindJSON(&verifyData)
			exceptiongo.QuickThrow[types.JsonUnmarshalFailedException](err)

			verifyBase(verifyData.Data, verifyData.Token)
		}
	}
}
