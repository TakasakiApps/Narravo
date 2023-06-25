package middlewares

import (
	"fmt"
	"github.com/TakasakiApps/Narravo/backend/config"
	"github.com/TakasakiApps/Narravo/backend/dao"
	"github.com/TakasakiApps/Narravo/backend/internal/entity"
	"github.com/TakasakiApps/Narravo/backend/internal/types"
	"github.com/TakasakiApps/Narravo/backend/internal/utils"
	"github.com/gin-gonic/gin"
	"github.com/gookit/goutil/maputil"
	"github.com/gookit/goutil/strutil"
	"github.com/ohanakogo/exceptiongo"
	"github.com/ohanakogo/exceptiongo/pkg/etype"
)

func getAuthorization(skipRouterSuffix ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		defer exceptiongo.NewExceptionHandler(func(exception *etype.Exception) {
			// stop context at first
			c.Abort()
			switch exception.Type() {
			default:
				exceptiongo.Throw(exception)
			}
		}).Deploy()

		if strutil.HasOneSuffix(c.FullPath(), skipRouterSuffix) {
			c.Next()
		} else {
			data := utils.GetData(c)
			if !maputil.HasKey(data, "token") {
				exceptiongo.QuickThrow[types.ServerUnauthorizedException](fmt.Errorf("authorization failed! please login before operating"))
			}
			userToken := &entity.UserToken{}
			utils.ConvMapToStructure(utils.GetData(c), userToken)

			defer exceptiongo.TryHandle[types.JWTParseFailedException](func(e *etype.Exception) {
				exceptiongo.QuickThrowMsg[types.ServerUnauthorizedException](e.Error())
			})

			utils.JWTParse[entity.User](userToken.Token, config.GetInstance().Crypto.TokenAesKey, func(isValid bool, obj entity.User) {
				if !isValid {
					exceptiongo.QuickThrowMsg[types.ServerUnauthorizedException]("token is invalid")
				}

				user := dao.GetInstance().QueryUserByName(&obj)
				if user == nil {
					exceptiongo.QuickThrowMsg[types.ServerUnauthorizedException]("token contains invalid data")
				}

				if user.Password != obj.Password {
					exceptiongo.QuickThrowMsg[types.ServerUnauthorizedException]("user password has expired")
				}

				c.Next()
			})
		}
	}
}
