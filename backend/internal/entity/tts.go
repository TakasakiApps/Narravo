package entity

import "github.com/ohanakogo/exceptiongo/pkg/etype"

type TTSTaskInfo struct {
	Text       string         `json:"text" binding:"required"`
	Volume     float64        `json:"volume" binding:"required"`
	Speed      float64        `json:"speed" binding:"required"`
	Codec      string         `json:"codec" binding:"required"`
	ModelId    string         `json:"modelId" binding:"required"`
	TTSEmotion TTSEmotionInfo `json:"ttsEmotion" binding:"required"`
}

type TTSEmotionInfo struct {
	Enabled   bool   `json:"enabled" binding:"required"`
	Category  string `json:"category"`
	Intensity int64  `json:"intensity"`
}

type TTSModel struct {
	Id                 int64        `json:"-"`
	Name               string       `json:"name"`
	Gender             int          `json:"gender"`
	IsEmotionSupported bool         `json:"isEmotionSupported"`
	Emotions           []TTSEmotion `json:"emotions"`
}

type TTSEmotion struct {
	Name     string `json:"name"`
	Category string `json:"category"`
}

type TTSResult struct {
	Exception *etype.Exception
	AudioB64  string
	Codec     string
	Timestamp int64
}
