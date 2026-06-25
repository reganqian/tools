package tx

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
	"encoding/binary"
	"fmt"
	"math/rand"
	"net/url"
	"sort"
	"time"

	"github.com/labstack/gommon/log"
)

type AsrConfig struct {
	AppID             string
	SecretID          string
	SecretKey         string
	EngineModelType   string
	VoiceId           string
	VoiceFormat       *int     //可选，默认值为4。1：pcm；4：speex(sp)；6：silk；8：mp3；10：opus（opus 格式音频流封装说明）；12：wav；14：m4a（每个分片须是一个完整的 m4a 音频）；16：aac 示例值：1
	Needvad           *int     //needvad0：关闭 vad，1：开启 vad，默认为0。为保证识别效果，如果语音分片长度超过60秒，会强制在60s 断一次，建议客户音频超过60s 时，开启 vad（人声检测切分功能），提升切分效果。示例值：1
	HotwordId         *string  //热词表 id。如不设置该参数，自动生效默认热词表；如果设置了该参数，那么将生效对应的热词表。示例值：da3f5f5555cf11eda6da525400aec391
	CustomizationId   *string  //自学习模型 id。如设置了该参数，将生效对应的自学习模型。
	FilterDirty       *int     //是否过滤脏词（目前支持中文普通话引擎）。默认为0。0：不过滤脏词；1：过滤脏词；2：将脏词替换为“ * ” 示例值：0
	FilterModal       *int     //是否过滤语气词（目前支持中文普通话引擎）。默认为0。0：不过滤语气词；1：部分过滤；2：严格过滤。 示例值：0
	FilterPunc        *int     //是否过滤句末的句号（目前支持中文普通话引擎）。默认为0。0：不过滤句末的句号；1：过滤句末的句号。 示例值：0
	FilterEmptyResult *int     //是否回调识别空结果，默认为1。0：回调空结果；1：不回调空结果。 示例值：1
	ConvertNumMode    *int     //是否进行阿拉伯数字智能转换（目前支持中文普通话引擎）。0：不转换，直接输出中文数字，1：根据场景智能转换为阿拉伯数字，3: 打开数学相关数字转换。默认值为1 示例值：1
	WordInfo          *int     //是否显示词级别时间戳。0：不显示；1：显示，不包含标点时间戳；2：显示，包含标点时间戳；支持引擎 8k_en、8k_zh、8k_zh_finance、16k_zh_en、16k_zh、16k_en、16k_ca、16k_zh-TW、16k_ja、8k_zh_large，默认为0 示例值：0
	VadSilenceTime    *int     //语音断句检测阈值，静音时长超过该阈值会被认为断句（多用在智能客服场景，需配合 needvad = 1 使用），取值范围：500-2000（默认1000），单位 ms，此参数建议不要随意调整，可能会影响识别效果，目前支持 8k_zh、8k_zh_finance、16k_zh、16k_zh_en、8k_zh_large、16k_en、16k_en_large 引擎 示例值：1000
	MaxSpeakTime      *int     //强制断句功能，取值范围 5000-90000（单位:毫秒），默认值60000。 在连续说话不间断情况下，该参数将实现强制断句（此时结果变成稳态，slice_type=2）。如：游戏解说场景，解说员持续不间断解说，无法断句的情况下，将此参数设置为10000，则将在每10秒收到 slice_type=2 的回调。 （目前仅支持8k_zh、16k_zh、16k_zh_en、8k_zh_large 引擎） 示例值：60000
	NoiseThreshold    *float64 //噪音参数阈值，默认为0，取值范围：[-2,2]，对于一些音频片段，取值越大，判定为噪音情况越大。取值越小，判定为人声情况越大。 示例值：0
	HotwordList       *string  //临时热词表：该参数用于提升识别准确率。
	InputSampleRate   *int     //支持 PCM 格式的 8k 音频在与引擎采样率不匹配的情况下升采样到 16k 后识别，能有效提升识别准确率。仅支持：8000。如：传入 8000 ，则 PCM 音频采样率为8k，当引擎选用16k_zh， 那么该8k 采样率的 PCM 音频可以在16k_zh 引擎下正常识别。
	ReplaceTextId     *string  //替换词汇表 ID,  适用于热词和自学习场景也无法解决的极端 case 词组，会对识别结果强制替换。具体可参考（词汇替换） ；强制替换功能可能会影响正常识别结果，请谨慎使用。
}

func generateSecureNonce() (int, error) {
	var b [8]byte
	if _, err := rand.Read(b[:]); err != nil {
		return 0, err
	}
	return int(binary.BigEndian.Uint64(b[:]) % 10000000000), nil
}

func GenerateAsrWssURL(cfg AsrConfig) (string, error) {
	timestamp := time.Now().Unix()
	expired := timestamp + 300 // 1小时有效期

	nonce, _ := generateSecureNonce()
	// 1. 构造参数字典
	params := map[string]interface{}{
		"secretid":          cfg.SecretID,
		"timestamp":         timestamp,
		"expired":           expired,
		"engine_model_type": cfg.EngineModelType,
		"nonce":             nonce,
		"voice_id":          cfg.VoiceId,
	}
	// 字段可配置
	// VoiceFormat       *string  //可选，默认值为4。1：pcm；4：speex(sp)；6：silk；8：mp3；10：opus（opus 格式音频流封装说明）；12：wav；14：m4a（每个分片须是一个完整的 m4a 音频）；16：aac 示例值：1
	if cfg.VoiceFormat != nil {
		params["voice_format"] = *cfg.VoiceFormat
	}
	// Needvad           *int     //needvad0：关闭 vad，1：开启 vad，默认为0。为保证识别效果，如果语音分片长度超过60秒，会强制在60s 断一次，建议客户音频超过60s 时，开启 vad（人声检测切分功能），提升切分效果。示例值：1
	if cfg.Needvad != nil {
		params["needvad"] = *cfg.Needvad
	}

	// HotwordId         *string  //热词表 id。如不设置该参数，自动生效默认热词表；如果设置了该参数，那么将生效对应的热词表。示例值：da3f5f5555cf11eda6da525400aec391
	if cfg.HotwordId != nil {
		params["hotword_id"] = *cfg.HotwordId
	}

	// CustomizationId   *string  //自学习模型 id。如设置了该参数，将生效对应的自学习模型。
	if cfg.CustomizationId != nil {
		params["customization_id"] = *cfg.CustomizationId
	}
	// FilterDirty       *int     //是否过滤脏词（目前支持中文普通话引擎）。默认为0。0：不过滤脏词；1：过滤脏词；2：将脏词替换为“ * ” 示例值：0
	if cfg.FilterDirty != nil {
		params["filter_dirty"] = *cfg.FilterDirty
	}

	// FilterModal       *int     //是否过滤语气词（目前支持中文普通话引擎）。默认为0。0：不过滤语气词；1：部分过滤；2：严格过滤。 示例值：0
	if cfg.FilterModal != nil {
		params["filter_modal"] = *cfg.FilterModal
	}

	// FilterPunc        *int     //是否过滤句末的句号（目前支持中文普通话引擎）。默认为0。0：不过滤句末的句号；1：过滤句末的句号。 示例值：0
	if cfg.FilterPunc != nil {
		params["filter_punc"] = *cfg.FilterPunc
	}
	// FilterEmptyResult *int     //是否回调识别空结果，默认为1。0：回调空结果；1：不回调空结果。 示例值：1
	if cfg.FilterEmptyResult != nil {
		params["filter_empty_result"] = *cfg.FilterEmptyResult
	}

	// ConvertNumMode    *int     //是否进行阿拉伯数字智能转换（目前支持中文普通话引擎）。0：不转换，直接输出中文数字，1：根据场景智能转换为阿拉伯数字，3: 打开数学相关数字转换。默认值为1 示例值：1
	if cfg.ConvertNumMode != nil {
		params["convert_num_mode"] = *cfg.ConvertNumMode
	}

	// WordInfo          *int     //是否显示词级别时间戳。0：不显示；1：显示，不包含标点时间戳；2：显示，包含标点时间戳；支持引擎 8k_en、8k_zh、8k_zh_finance、16k_zh_en、16k_zh、16k_en、16k_ca、16k_zh-TW、16k_ja、8k_zh_large，默认为0 示例值：0
	if cfg.WordInfo != nil {
		params["word_info"] = *cfg.WordInfo
	}
	// VadSilenceTime    *int     //语音断句检测阈值，静音时长超过该阈值会被认为断句（多用在智能客服场景，需配合 needvad = 1 使用），取值范围：500-2000（默认1000），单位 ms，此参数建议不要随意调整，可能会影响识别效果，目前支持 8k_zh、8k_zh_finance、16k_zh、16k_zh_en、8k_zh_large、16k_en、16k_en_large 引擎 示例值：1000
	if cfg.VadSilenceTime != nil {
		params["vad_silence_time"] = *cfg.VadSilenceTime
	}
	// MaxSpeakTime      *int     //强制断句功能，取值范围 5000-90000（单位:毫秒），默认值60000。 在连续说话不间断情况下，该参数将实现强制断句（此时结果变成稳态，slice_type=2）。如：游戏解说场景，解说员持续不间断解说，无法断句的情况下，将此参数设置为10000，则将在每10秒收到 slice_type=2 的回调。 （目前仅支持8k_zh、16k_zh、16k_zh_en、8k_zh_large 引擎） 示例值：60000
	if cfg.MaxSpeakTime != nil {
		params["max_speak_time"] = *cfg.MaxSpeakTime
	}
	// NoiseThreshold    *float64 //噪音参数阈值，默认为0，取值范围：[-2,2]，对于一些音频片段，取值越大，判定为噪音情况越大。取值越小，判定为人声情况越大。 示例值：0
	if cfg.NoiseThreshold != nil {
		params["noise_threshold"] = *cfg.NoiseThreshold
	}
	// HotwordList       *string  //临时热词表：该参数用于提升识别准确率。
	if cfg.HotwordList != nil {
		params["hotword_list"] = *cfg.HotwordList
	}
	// InputSampleRate   *int     //支持 PCM 格式的 8k 音频在与引擎采样率不匹配的情况下升采样到 16k 后识别，能有效提升识别准确率。仅支持：8000。如：传入 8000 ，则 PCM 音频采样率为8k，当引擎选用16k_zh， 那么该8k 采样率的 PCM 音频可以在16k_zh 引擎下正常识别。
	if cfg.InputSampleRate != nil {
		params["input_sample_rate"] = *cfg.InputSampleRate
	}

	// ReplaceTextId       *string  //替换词汇表 ID,  适用于热词和自学习场景也无法解决的极端 case 词组，会对识别结果强制替换。具体可参考（词汇替换） ；强制替换功能可能会影响正常识别结果，请谨慎使用。
	if cfg.ReplaceTextId != nil {
		params["replace_text_id"] = *cfg.ReplaceTextId
	}

	// 2. 按 key 字典序排序
	keys := make([]string, 0, len(params))
	for k := range params {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	// 3. 拼接签名原文
	var signStr string
	for i, k := range keys {
		if i > 0 {
			signStr += "&"
		}
		signStr += fmt.Sprintf("%v=%v", k, params[k])
	}

	baseURL := fmt.Sprintf(
		"asr.cloud.tencent.com/asr/v2/%s",
		cfg.AppID,
	)

	signStr = baseURL + "?" + signStr
	// 4. HMAC-SHA1
	log.Info("signStr======", signStr)
	log.Info("SecretKey======", cfg.SecretKey)
	signature := HmacSha1Base64([]byte(cfg.SecretKey), []byte(signStr))
	log.Info("signature======", signature)
	// 5. URL Encode
	signature = url.QueryEscape(signature)

	// 6. 组装 WSS URL
	log.Info("signature======", signature)
	query := signStr + fmt.Sprintf(
		"&signature=%s",
		signature,
	)

	return "wss://" + query, nil
}

func HmacSha1Base64(key, data []byte) string {
	mac := hmac.New(sha1.New, key)
	mac.Write(data)
	return base64.StdEncoding.EncodeToString(mac.Sum(nil))
}
