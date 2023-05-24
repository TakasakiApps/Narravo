package test

import (
	"fmt"
	"github.com/TakasakiApps/Narravo/backend/internal/global"
	"github.com/TakasakiApps/Narravo/backend/server/engine"
	"github.com/TakasakiApps/Narravo/backend/server/routers"
	"github.com/gookit/goutil/testutil/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestVersionRouter(t *testing.T) {
	routers.Register()

	req, err := http.NewRequest("GET", "/api/v1/version", nil)
	if err != nil {
		t.Fatal(err)
	}

	res := httptest.NewRecorder()
	engine.Gin.ServeHTTP(res, req)

	assert.Eq(t, http.StatusOK, res.Code)

	assert.Eq(t,
		fmt.Sprintf(`{"version":"%s"}`, global.GetVersion()),
		res.Body.String(),
	)
}
