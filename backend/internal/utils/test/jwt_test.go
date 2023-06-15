package test

import (
	"github.com/TakasakiApps/Narravo/backend/internal/utils"
	"github.com/gookit/goutil/testutil/assert"
	"testing"
	"time"
)

type TUser struct {
	Username string
	Password string
}

var (
	token string
	key   = "rQFZI95yFVULezez"
	user  = TUser{
		Username: "test",
		Password: "123456",
	}
)

func TestJWT(t *testing.T) {
	JWTSign(t)
	JWTParse(t)
}

func JWTSign(t *testing.T) {
	token = utils.JWTSign[TUser](user, key, 5)
	t.Log("token is:", token)
}

func JWTParse(t *testing.T) {
	utils.JWTParse[TUser](token, key, func(isValid bool, obj TUser) {
		assert.Eq(t, true, isValid)
		assert.Eq(t, user, obj)
	})
	time.Sleep(time.Second * time.Duration(5))
	utils.JWTParse[TUser](token, key, func(isValid bool, obj TUser) {
		assert.Eq(t, false, isValid)
		assert.Eq(t, user, obj)
	})
}
