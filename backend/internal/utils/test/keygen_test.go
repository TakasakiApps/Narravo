package test

import (
	"github.com/TakasakiApps/Narravo/backend/internal/utils"
	"testing"
)

func TestKeygen(t *testing.T) {
	id := utils.GenerateStandardID()
	t.Log(id)
}
