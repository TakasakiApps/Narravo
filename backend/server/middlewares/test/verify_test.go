package test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/TakasakiApps/Narravo/backend/internal/utils"
	"github.com/TakasakiApps/Narravo/backend/server/engine"
	"github.com/TakasakiApps/Narravo/backend/server/middlewares"
	"github.com/gin-gonic/gin"
	"github.com/gookit/goutil"
	"github.com/gookit/goutil/testutil/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func mTestVerify(t *testing.T) {
	middlewares.Register()

	engine.Gin.POST("/debug", func(c *gin.Context) {
		c.JSON(http.StatusOK, c.MustGet("data"))
	})

	// 原始数据
	marshalData, err := json.Marshal(map[string]any{
		"testNumber": 0,
	})
	if err != nil {
		t.Error(err)
	}

	// 用于加密原始数据的key, 加密后放在data
	encryptKey := "32150285b345c48aa3492f9212f61ca2"
	dataCipher, err := utils.NewAesCipher(encryptKey)
	if err != nil {
		t.Error(err)
	}
	// 加密后的data
	encryptMarshalData, err := dataCipher.Encrypt(string(marshalData))
	if err != nil {
		t.Error(err)
	}

	// 时间戳字符串(毫秒)
	milliTimestamp, err := goutil.ToString(utils.GetCurrentTimestampMilli())
	if err != nil {
		t.Error(err)
	}
	// 生成一个AES加密实例, key是当前时间戳的md5 hash
	integrityKeyCipher, err := utils.NewAesCipher(utils.MD5(milliTimestamp))
	if err != nil {
		t.Error(err)
	}
	// 把加密原始数据的key加密, 生成数据完整性密钥
	integrityKey, err := integrityKeyCipher.Encrypt(encryptKey)
	if err != nil {
		t.Error(err)
	}

	postData := map[string]any{
		"data":         encryptMarshalData,
		"integrityKey": integrityKey,
	}

	marshalPostData, err := json.Marshal(postData)
	if err != nil {
		t.Error(err)
	}

	req, err := http.NewRequest("POST", "/debug", bytes.NewBuffer(marshalPostData))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")

	res := httptest.NewRecorder()

	fmt.Printf("Sending %v\n", string(marshalPostData))
	engine.Gin.ServeHTTP(res, req)

	assert.Eq(t, http.StatusOK, res.Code)

	assert.Eq(t,
		string(marshalData),
		res.Body.String(),
	)
}
