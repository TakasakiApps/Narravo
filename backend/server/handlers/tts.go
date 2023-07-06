package handlers

import (
	"github.com/TakasakiApps/Narravo/backend/internal/entity"
	"github.com/TakasakiApps/Narravo/backend/internal/global"
	"github.com/TakasakiApps/Narravo/backend/internal/tencentcloud"
	"github.com/TakasakiApps/Narravo/backend/internal/types"
	"github.com/TakasakiApps/Narravo/backend/internal/utils"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/gookit/goutil/maputil"
	"github.com/gookit/goutil/strutil"
	"github.com/ohanakogo/exceptiongo"
	"net/http"
)

var GetModelInfos gin.HandlerFunc = func(c *gin.Context) {
	c.JSON(http.StatusOK, global.TTSModels)
}

var CreateTTSTask gin.HandlerFunc = func(c *gin.Context) {
	ttsTaskInfo := &entity.TTSTaskInfo{}
	utils.ConvMapToStructure(utils.GetData(c), ttsTaskInfo)

	ttsTaskId := uuid.New().String()
	go tencentcloud.NewTTSTask(ttsTaskId, ttsTaskInfo)

	c.JSON(http.StatusOK, map[string]any{
		"ttsTaskId": ttsTaskId,
	})
}

var QueryTTSTaskResult gin.HandlerFunc = func(c *gin.Context) {
	ttsTaskId := c.Query("ttsTaskId")
	if strutil.IsEmpty(ttsTaskId) {
		exceptiongo.QuickThrowMsg[types.ServerBadRequestException](`"ttsTaskId" can not be empty`)
	}

	if !maputil.HasKey(global.TTSResults, ttsTaskId) {
		c.Status(http.StatusTooEarly)
		return
	}

	if exception := global.TTSResults[ttsTaskId].Exception; exception != nil {
		// 删除掉已经无用的结果
		delete(global.TTSResults, ttsTaskId)
		exceptiongo.Throw(exception)
	}

	ttsResult := global.TTSResults[ttsTaskId]
	// 取出值后马上删除
	delete(global.TTSResults, ttsTaskId)

	c.JSON(http.StatusOK, map[string]any{
		"audioB64": ttsResult.AudioB64,
		"codec":    ttsResult.Codec,
	})
}
