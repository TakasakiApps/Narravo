package handlers

import (
	"fmt"
	"github.com/TakasakiApps/Narravo/backend/config"
	"github.com/TakasakiApps/Narravo/backend/dao"
	"github.com/TakasakiApps/Narravo/backend/internal/entity"
	"github.com/TakasakiApps/Narravo/backend/internal/global"
	"github.com/TakasakiApps/Narravo/backend/internal/types"
	"github.com/TakasakiApps/Narravo/backend/internal/utils"
	"github.com/gin-gonic/gin"
	"github.com/ohanakogo/exceptiongo"
	"net/http"
)

// LegacyAuthorizationComponent Define a Gin middleware function for legacy authorization.
var LegacyAuthorizationComponent gin.HandlerFunc = func(c *gin.Context) {
	// Create an empty user struct and populate it with data from the request.
	user := &entity.User{}
	utils.ConvMapToStructure(utils.GetData(c), user)
	// Query the database for a user with the same name as the one in the request.
	queryUser := dao.GetInstance().QueryUserByName(user.Name)
	// If no matching user is found, throw an unauthorized exception with an error message.
	if queryUser == nil {
		exceptiongo.QuickThrowMsg[types.ServerUnauthorizedException](fmt.Sprintf("User %v is not registered", *user))
		// If the user exists and their password matches the provided password, call the next middleware function.
	} else if queryUser.Password == utils.MD5(user.Password) {
		c.Next()
		// If the password is incorrect, throw an unauthorized exception with an error message.
	} else {
		exceptiongo.QuickThrowMsg[types.ServerUnauthorizedException]("Incorrect password provided")
	}
}

var Register gin.HandlerFunc = func(c *gin.Context) {
	user := &entity.User{}
	utils.ConvMapToStructure(utils.GetData(c), user)

	queryUser := dao.GetInstance().QueryUserByName(user.Name)
	if queryUser != nil {
		exceptiongo.QuickThrowMsg[types.ServerUnauthorizedException](fmt.Sprintf("user%v already existed", *user))
	}

	effected := dao.GetInstance().AddUser(&entity.User{
		Name:     user.Name,
		Password: utils.MD5(user.Password),
	})

	if effected == 1 {
		c.JSON(http.StatusOK, user)
	} else {
		exceptiongo.QuickThrowMsg[types.DataBaseOperationFailedException]("register failed")
	}
}

var Login gin.HandlerFunc = func(c *gin.Context) {
	user := &entity.User{}
	utils.ConvMapToStructure(utils.GetData(c), user)

	queryUser := dao.GetInstance().QueryUserByName(user.Name)
	c.JSON(http.StatusOK, entity.UserToken{
		Token: utils.JWTSign[entity.User](*queryUser, config.GetInstance().Crypto.TokenAesKey, global.TokenExpireDelay),
	})
}

var Renew gin.HandlerFunc = func(c *gin.Context) {
	userToken := &entity.UserToken{}
	utils.ConvMapToStructure(utils.GetData(c), userToken)

	user := utils.JWTParseObj[entity.User](userToken.Token, config.GetInstance().Crypto.TokenAesKey)

	c.JSON(http.StatusOK, entity.UserToken{
		Token: utils.JWTSign[entity.User](user, config.GetInstance().Crypto.TokenAesKey, global.TokenExpireDelay),
	})
}

var ResetPassword gin.HandlerFunc = func(c *gin.Context) {
	userResetPassword := &entity.UserResetPassword{}
	utils.ConvMapToStructure(utils.GetData(c), userResetPassword)

	// Check if the user's old password matches their new password.
	if userResetPassword.Password == userResetPassword.NewPassword {
		// If the passwords are the same, throw a bad request exception with an error message.
		exceptiongo.QuickThrow[types.ServerBadRequestException](fmt.Errorf("the old password cannot be the same as the new password"))
	}

	// Get user information by name from the instance of dao.
	queryUser := dao.GetInstance().QueryUserByName(userResetPassword.Name)
	// Update the user's password to the new password, hashed using MD5.
	queryUser.Password = utils.MD5(userResetPassword.NewPassword)
	// Update the user's information in the database and check if one row is affected.
	effected := dao.GetInstance().UpdateUser(queryUser)
	// If one row is affected, return the updated user's information as a JSON response.
	if effected == 1 {
		c.JSON(http.StatusOK, &entity.User{
			Model:    queryUser.Model,
			Name:     queryUser.Name,
			Password: userResetPassword.NewPassword,
		})
		// If no rows are affected, throw a database operation failed exception with an error message.
	} else {
		exceptiongo.QuickThrowMsg[types.DataBaseOperationFailedException]("reset password failed")
	}
}
