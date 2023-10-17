package callback

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/OpenIMSDK/protocol/sdkws"
	"github.com/OpenIMSDK/tools/log"
	"github.com/openimsdk/open-im-server/v3/pkg/apistruct"
	"github.com/openimsdk/open-im-server/v3/pkg/common/config"
	http_client "github.com/openimsdk/open-im-server/v3/pkg/common/http"
)

type HtCommonResponse struct {
	Code string      `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"resultData"`
}

func CustomCallback(ctx context.Context, data *sdkws.MsgData) {

	content := &apistruct.CustomElem{}
	err := json.Unmarshal(data.GetContent(), content)
	if err != nil {
		log.ZError(ctx, "自定义消息不合法", err)
		return
	}
	switch content.Description {
	case "LINK_ONCE":
		if data.RecvID != config.Config.Colline.LinkOnceUser {
			// LINK_ONCE 类型， 接收人一定是 link
			return
		}
	default:

	}
	customApi := config.Config.Colline.Host + config.Config.Colline.CustomCallbackApi
	dataStr, _ := json.Marshal(data)
	log.ZDebug(ctx, "回调的请求", "request", bytes.NewBuffer(dataStr), "api", customApi)
	responseStr, err := http_client.Post(ctx, customApi, map[string]string{}, data, 5)
	if err != nil {
		log.ZError(ctx, "回调华泰服务出错", err)
		return
	}
	response := &HtCommonResponse{}
	if err = json.Unmarshal(responseStr, response); err != nil {
		log.ZError(ctx, "反序列化出错: "+string(responseStr), err)
		return
	}
	if response.Code != "0" {
		log.ZError(ctx, "回调华泰服务，业务异常", err)
	}
}
