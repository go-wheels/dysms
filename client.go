package dysms

import (
	"encoding/json"

	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	dysmsapi "github.com/alibabacloud-go/dysmsapi-20170525/v4/client"
	"github.com/alibabacloud-go/tea/tea"
)

const endpoint = "dysmsapi.aliyuncs.com"

type TemplateParam map[string]string

type Client struct {
	dysmsapiClient *dysmsapi.Client
}

func NewClient(accessKeyID, accessKeySecret string) (client *Client, err error) {
	config := &openapi.Config{
		AccessKeyId:     tea.String(accessKeyID),
		AccessKeySecret: tea.String(accessKeySecret),
		Endpoint:        tea.String(endpoint),
	}
	client = new(Client)
	client.dysmsapiClient, err = dysmsapi.NewClient(config)
	return
}

// 发送短信
func (c Client) SendSMS(phoneNumbers, signName, templateCode string, templateParam interface{}) (resp *dysmsapi.SendSmsResponse, err error) {
	req := &dysmsapi.SendSmsRequest{
		PhoneNumbers: tea.String(phoneNumbers),
		SignName:     tea.String(signName),
		TemplateCode: tea.String(templateCode),
	}
	if templateParam != nil {
		var data []byte
		data, err = json.Marshal(templateParam)
		if err != nil {
			return
		}
		req.TemplateParam = tea.String(string(data))
	}
	resp, err = c.dysmsapiClient.SendSms(req)
	if err != nil {
		return
	}
	if tea.StringValue(resp.Body.Code) != "OK" {
		err = newSDKError(resp.Body.Code, resp.Body.Message)
	}
	return
}

// 创建短链
func (c Client) AddShortURL(sourceURL, shortURLName, effectiveDays string) (resp *dysmsapi.AddShortUrlResponse, err error) {
	req := &dysmsapi.AddShortUrlRequest{
		SourceUrl:     tea.String(sourceURL),
		ShortUrlName:  tea.String(shortURLName),
		EffectiveDays: tea.String(effectiveDays),
	}
	resp, err = c.dysmsapiClient.AddShortUrl(req)
	if err != nil {
		return
	}
	if tea.StringValue(resp.Body.Code) != "OK" {
		err = newSDKError(resp.Body.Code, resp.Body.Message)
	}
	return
}

func newSDKError(code, message *string) *tea.SDKError {
	return tea.NewSDKError(map[string]interface{}{
		"code":    tea.StringValue(code),
		"message": tea.StringValue(message),
	})
}
