package dingding

import (
	"bytes"
	"github.com/parnurzeal/gorequest"
)

const GateHost = "https://oapi.dingtalk.com/robot/send"

//const AccessToken = "a8e6c50548df4d3d42478f6c807d99ea1d1fd8ababb59e150cb764013b9ab7bb"

func SendMessage(msg, accessToken string) (string, error) {
	jsonObject := map[string]interface{}{
		"msgtype": "text",
		"text": map[string]interface{}{
			"content": msg,
		},
	}
	var buf bytes.Buffer
	buf.WriteString(GateHost)
	buf.WriteString("?access_token=")
	buf.WriteString(accessToken)
	apiUrl := buf.String()
	request := gorequest.New()
	_, body, errs := request.Post(apiUrl).
		Set("Content-Type", "application/json").
		Send(jsonObject).End()
	var err error
	if len(errs) > 0 {
		err = errs[0]
	} else {
		err = nil
	}
	return body, err
}
