package tencentcloud

import (
	"github.com/TakasakiApps/Narravo/backend/config"
	"github.com/TakasakiApps/Narravo/backend/internal/entity"
	"github.com/TakasakiApps/Narravo/backend/internal/global"
	"github.com/TakasakiApps/Narravo/backend/internal/types"
	"github.com/TakasakiApps/Narravo/backend/internal/utils"
	"github.com/gookit/goutil/maputil"
	"github.com/gookit/goutil/strutil"
	"github.com/ohanakogo/exceptiongo"

	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
	tts "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/tts/v20190823"
)

func NewTTSTask(ttsTaskId string, ttsTaskInfo *entity.TTSTaskInfo) {
	configInstance := config.GetInstance()
	secretId := configInstance.TencentCloud.SecretId
	secretKey := configInstance.TencentCloud.SecretKey

	if strutil.HasEmpty(secretId, secretKey) {
		global.TTSResults[ttsTaskId] = entity.TTSResult{
			Exception: exceptiongo.NewException[types.TencentCloudApiKeyOrSecretInvalidException]("secretKey or secretId is empty, please contact admin to fill up api key"),
		}
		return
	}

	credential := common.NewCredential(
		secretId, secretKey,
	)

	if !maputil.HasKey(global.TTSModels, ttsTaskInfo.ModelId) {
		global.TTSResults[ttsTaskId] = entity.TTSResult{
			Exception: exceptiongo.NewExceptionF[types.TTSModelNotFoundException](`modelId "%s" not found`, ttsTaskInfo.ModelId),
		}
		return
	}

	if ttsTaskInfo.Codec != "mp3" && ttsTaskInfo.Codec != "wav" {
		global.TTSResults[ttsTaskId] = entity.TTSResult{
			Exception: exceptiongo.NewExceptionF[types.RuntimeException](`invalid codec "%s" for tts`, ttsTaskInfo.Codec),
		}
		return
	}

	cpf := profile.NewClientProfile()
	cpf.HttpProfile.Endpoint = "tts.tencentcloudapi.com"

	client, _ := tts.NewClient(credential, "ap-chengdu", cpf)

	request := tts.NewTextToVoiceRequest()

	request.Text = common.StringPtr(ttsTaskInfo.Text)
	request.SessionId = common.StringPtr(ttsTaskId)
	request.Volume = common.Float64Ptr(ttsTaskInfo.Volume)
	request.Speed = common.Float64Ptr(ttsTaskInfo.Speed)
	request.ModelType = common.Int64Ptr(1)
	request.VoiceType = common.Int64Ptr(global.TTSModels[ttsTaskInfo.ModelId].Id)
	request.PrimaryLanguage = common.Int64Ptr(1)
	request.Codec = common.StringPtr(ttsTaskInfo.Codec)
	request.EnableSubtitle = common.BoolPtr(true)

	if ttsTaskInfo.TTSEmotion.Enabled {
		request.EmotionCategory = common.StringPtr(ttsTaskInfo.TTSEmotion.Category)
		request.EmotionIntensity = common.Int64Ptr(ttsTaskInfo.TTSEmotion.Intensity)
	}

	response, err := client.TextToVoice(request)
	if _, ok := err.(*errors.TencentCloudSDKError); ok {
		global.TTSResults[ttsTaskId] = entity.TTSResult{
			Exception: exceptiongo.NewException[types.TencentCloudSDKErrorException](err.Error()),
		}
		return
	}
	if err != nil {
		global.TTSResults[ttsTaskId] = entity.TTSResult{
			Exception: exceptiongo.NewException[types.TencentCloudSDKErrorException](err.Error()),
		}
		return
	}

	if *response.Response.SessionId != ttsTaskId {
		global.TTSResults[ttsTaskId] = entity.TTSResult{
			Exception: exceptiongo.NewException[types.RuntimeException]("ttsTaskId of response is not equal to request's id"),
		}
		return
	}

	if response.Response.Audio == nil || *response.Response.Audio == "" {
		global.TTSResults[ttsTaskId] = entity.TTSResult{
			Exception: exceptiongo.NewException[types.RuntimeException]("audio in the response is empty"),
		}
		return
	}

	global.TTSResults[ttsTaskId] = entity.TTSResult{
		Exception: nil,
		AudioB64:  *response.Response.Audio,
		Codec:     ttsTaskInfo.Codec,
		Timestamp: utils.GetCurrentTimestampMilli(),
	}
}
