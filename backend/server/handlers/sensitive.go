package handlers

import (
	"github.com/TakasakiApps/Narravo/backend/internal/assets"
	"github.com/TakasakiApps/Narravo/backend/internal/entity"
	"github.com/TakasakiApps/Narravo/backend/internal/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

var SensitiveCheck gin.HandlerFunc = func(c *gin.Context) {
	content := &entity.Content{}
	utils.ConvMapToStructure(utils.GetData(c), content)

	sensitiveWords := assets.SensitiveWordsFilter.FindAll(content.Text)

	c.JSON(http.StatusOK, map[string]any{
		"sensitiveWords": sensitiveWords,
	})
}
