//go:build !release

package engine

func setMode() {
	gin.SetMode("debug")
}
