package utils

import (
	"encoding/json"
	"fmt"
	"github.com/TakasakiApps/Narravo/backend/internal/types"
	"github.com/ohanakogo/exceptiongo"
	"github.com/ohanakogo/ohanakoutilgo"
	"reflect"
)

func ConvMapToStructure(mp any, st any) {
	switch {
	case ohanakoutilgo.Is[map[string]any](mp):
		jsonData, _ := json.Marshal(mp)
		_ = json.Unmarshal(jsonData, st)
	default:
		typeOf := reflect.TypeOf(mp)
		exceptiongo.QuickThrowMsg[types.RuntimeException](fmt.Sprintf("can't convert <%s> as a map structure", typeOf))
	}
}
