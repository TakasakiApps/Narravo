package handlers

import (
	"fmt"
	"github.com/TakasakiApps/Narravo/backend/internal/assets"
	"github.com/TakasakiApps/Narravo/backend/internal/types"
	"github.com/gin-gonic/gin"
	"github.com/gookit/goutil/arrutil"
	"github.com/gookit/goutil/fsutil"
	"github.com/ohanakogo/exceptiongo"
	"net/http"
	"path/filepath"
)

func GetUploadCheckerComponent(allowContentType ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		_, headers, err := c.Request.FormFile("file")
		exceptiongo.QuickThrow[types.ServerBadRequestException](err)
		if headers.Size > 1024*1024*2 {
			exceptiongo.QuickThrowMsg[types.ServerBadRequestException]("file too large")
		}
		contentType := headers.Header.Get("Content-Type")
		if !arrutil.StringsHas(allowContentType, contentType) {
			exceptiongo.QuickThrowMsg[types.ServerBadRequestException](fmt.Sprintf(
				"incorrect contentType<%s>. Your contentType must be one of the following: [%s]",
				contentType,
				arrutil.JoinSlice(", ", allowContentType),
			))
		}
	}
}

var UploadImage gin.HandlerFunc = func(c *gin.Context) {
	_, headers, _ := c.Request.FormFile("file")
	imageType := c.Query("type")
	var writeFilePath string
	switch imageType {
	case "avatar":
		writeFilePath = assets.GetAvatarPath(GenerateUniqueID())
	case "cover":
		CheckAdminPermissionComponent(c)
		writeFilePath = assets.GetCoverPath(GenerateUniqueID())
	default:
		exceptiongo.QuickThrowMsg[types.ServerBadRequestException](fmt.Sprintf("type<%s> is not allowed here", imageType))
	}
	err := c.SaveUploadedFile(headers, writeFilePath)
	exceptiongo.QuickThrow[types.ServerNotModifiedException](err)
	c.JSON(http.StatusOK, map[string]any{
		"id": filepath.Base(writeFilePath),
	})
}

var FetchImage gin.HandlerFunc = func(c *gin.Context) {
	imageType := c.Query("type")
	id := c.Query("id")
	var readFilePath string
	switch imageType {
	case "avatar":
		readFilePath = assets.GetAvatarPath(id)
	case "cover":
		readFilePath = assets.GetCoverPath(id)
	default:
		exceptiongo.QuickThrowMsg[types.ServerBadRequestException](fmt.Sprintf("type<%s> is not allowed here", imageType))
	}
	if !fsutil.FileExists(readFilePath) {
		exceptiongo.QuickThrowMsg[types.ServerBadRequestException](fmt.Sprintf("cannot find image file for specified id<%s>", id))
	}
	c.File(readFilePath)
}
