package global

import "github.com/TakasakiApps/Narravo/backend/internal/entity"

var TTSResults = make(map[string]entity.TTSResult)

var (
	TTSEmotionNeutral   = entity.TTSEmotion{Name: "中性", Category: "neutral"}
	TTSEmotionSad       = entity.TTSEmotion{Name: "悲伤", Category: "sad"}
	TTSEmotionHappy     = entity.TTSEmotion{Name: "高兴", Category: "happy"}
	TTSEmotionAngry     = entity.TTSEmotion{Name: "生气", Category: "angry"}
	TTSEmotionFear      = entity.TTSEmotion{Name: "恐惧", Category: "fear"}
	TTSEmotionNews      = entity.TTSEmotion{Name: "新闻", Category: "news"}
	TTSEmotionStory     = entity.TTSEmotion{Name: "故事", Category: "story"}
	TTSEmotionRadio     = entity.TTSEmotion{Name: "广播", Category: "radio"}
	TTSEmotionPoetry    = entity.TTSEmotion{Name: "诗歌", Category: "poetry"}
	TTSEmotionCall      = entity.TTSEmotion{Name: "客服", Category: "call"}
	TTSEmotionSajiao    = entity.TTSEmotion{Name: "撒娇", Category: "sajiao"}
	TTSEmotionDisgusted = entity.TTSEmotion{Name: "厌恶", Category: "disgusted"}
	TTSEmotionAmaze     = entity.TTSEmotion{Name: "震惊", Category: "amaze"}
	TTSEmotionPeaceful  = entity.TTSEmotion{Name: "平静", Category: "peaceful"}
	TTSEmotionExciting  = entity.TTSEmotion{Name: "兴奋", Category: "exciting"}
)

var TTSModels = map[string]entity.TTSModel{
	"axg": {
		Id:                 301000,
		Name:               "成年男性1号（多情感）",
		Gender:             0,
		IsEmotionSupported: true,
		Emotions: []entity.TTSEmotion{
			TTSEmotionNeutral, TTSEmotionHappy, TTSEmotionAngry, TTSEmotionSad,
			TTSEmotionFear, TTSEmotionNews, TTSEmotionStory, TTSEmotionRadio,
			TTSEmotionPoetry, TTSEmotionCall,
		},
	},
	"axs": {
		Id:                 301031,
		Name:               "成年男性2号（多情感）",
		Gender:             0,
		IsEmotionSupported: true,
		Emotions: []entity.TTSEmotion{
			TTSEmotionNeutral, TTSEmotionSad, TTSEmotionAngry, TTSEmotionHappy,
			TTSEmotionFear, TTSEmotionAmaze, TTSEmotionSajiao, TTSEmotionDisgusted,
		},
	},
	"axx": {
		Id:                 301003,
		Name:               "成年女性1号（多情感）",
		Gender:             1,
		IsEmotionSupported: true,
		Emotions: []entity.TTSEmotion{
			TTSEmotionNeutral, TTSEmotionHappy, TTSEmotionAngry, TTSEmotionSad,
			TTSEmotionFear, TTSEmotionNews, TTSEmotionStory, TTSEmotionRadio,
			TTSEmotionPoetry, TTSEmotionCall,
		},
	},
	"axm": {
		Id:                 301035,
		Name:               "成年女性2号（多情感）",
		Gender:             1,
		IsEmotionSupported: true,
		Emotions: []entity.TTSEmotion{
			TTSEmotionNeutral, TTSEmotionSad, TTSEmotionAngry, TTSEmotionHappy,
			TTSEmotionFear, TTSEmotionAmaze, TTSEmotionSajiao, TTSEmotionDisgusted,
		},
	},
	"axt": {
		Id:                 301038,
		Name:               "成年女性3号（多情感）",
		Gender:             1,
		IsEmotionSupported: true,
		Emotions: []entity.TTSEmotion{
			TTSEmotionNeutral, TTSEmotionPeaceful, TTSEmotionExciting, TTSEmotionPoetry,
			TTSEmotionSad, TTSEmotionAngry, TTSEmotionFear, TTSEmotionSajiao,
			TTSEmotionDisgusted,
		},
	},
	"zm": {
		Id:                 101015,
		Name:               "男童声",
		Gender:             0,
		IsEmotionSupported: false,
		Emotions:           nil,
	},
	"zt": {
		Id:                 101016,
		Name:               "女童声",
		Gender:             1,
		IsEmotionSupported: false,
		Emotions:           nil,
	},
}
