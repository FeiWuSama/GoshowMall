package rpc

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"
	"workspace-goshow-mall/adaptor"
	"workspace-goshow-mall/adaptor/repo/dto"
	"workspace-goshow-mall/adaptor/repo/vo"
	"workspace-goshow-mall/config"
	"workspace-goshow-mall/result"
)

type LarkRpc struct {
	config *config.Config
}

func NewLarkRpc(adaptor *adaptor.Adaptor) *LarkRpc {
	return &LarkRpc{
		config: &adaptor.Config,
	}
}

func (l *LarkRpc) GetLarkUserInfo(ctx context.Context, accessToken string) (*vo.LarkUserVo, error) {
	url := "https://open.feishu.cn/open-apis/authen/v1/user_info"
	request, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, err
	}
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Authorization", "Bearer "+accessToken)
	resp, err := http.DefaultClient.Do(request)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	var larkUserVo vo.LarkUserVo
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	respData := &result.Result[vo.LarkUserVo]{}
	err = json.Unmarshal(respBody, respData)
	if respData.Code != 0 {
		return nil, result.NewBusinessErrorWithMsg(result.ServerError, respData.Msg)
	}
	larkUserVo = respData.Data
	if err != nil {
		return nil, err
	}
	return &larkUserVo, nil
}

func (l *LarkRpc) GetLarkUserAccessToken(ctx context.Context, appCode int32, code string, redirectUrl string, scope string) (*vo.LarkAccessTokenVo, error) {
	url := "https://open.feishu.cn/open-apis/authen/v2/oauth/token"
	body := map[string]interface{}{
		"grant_type":    "authorization_code",
		"client_id":     l.config.AppConfig[appCode].AppId,
		"client_secret": l.config.AppConfig[appCode].AppSecret,
		"code":          code,
		"redirect_uri":  redirectUrl,
		"scope":         scope,
	}
	bodyBytes, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	request, err := http.NewRequestWithContext(ctx, "POST", url, bytes.NewBuffer(bodyBytes))
	if err != nil {
		return nil, err
	}
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Accept", "application/json")
	response, err := http.DefaultClient.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	respBody, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	var larkAccessTokenVo vo.LarkAccessTokenVo
	err = json.Unmarshal(respBody, &larkAccessTokenVo)
	if err != nil {
		return nil, err
	}
	return &larkAccessTokenVo, nil
}

type GetTokenFunc func(ctx context.Context) (string, error)

func (l *LarkRpc) SendLarkMsg(ctx context.Context, tokenFunc GetTokenFunc, dto *dto.UserLarkMsgDto) error {
	token, err := tokenFunc(ctx)
	url := "https://open.feishu.cn/open-apis/im/v1/messages?receive_id_type=" + dto.IdType
	contentMap := map[string]interface{}{
		"text": dto.Content,
	}
	contentBytes, err := json.Marshal(contentMap)
	if err != nil {
		return err
	}
	body := map[string]interface{}{
		"receive_id": dto.OpenId,
		"msg_type":   "text",
		"content":    string(contentBytes),
	}
	bodyBytes, err := json.Marshal(body)
	request, err := http.NewRequestWithContext(ctx, "POST", url, bytes.NewBuffer(bodyBytes))
	if err != nil {
		return err
	}
	request.Header.Set("Authorization", "Bearer "+token)
	request.Header.Set("Content-Type", "application/json")
	response, err := http.DefaultClient.Do(request)
	if err != nil {
		return err
	}
	defer response.Body.Close()
	respBody, err := io.ReadAll(response.Body)
	if err != nil {
		return err
	}
	respData := &result.Result[any]{}
	err = json.Unmarshal(respBody, respData)
	if err != nil {
		return err
	}
	if respData.Code != 0 {
		return result.NewBusinessErrorWithMsg(result.ServerError, respData.Msg)
	}
	return nil
}

func (l *LarkRpc) GetLarkTenantToken(ctx context.Context, code int32) (*vo.LarkTenantTokenVo, error) {
	url := "https://open.feishu.cn/open-apis/auth/v3/tenant_access_token/internal"
	body := map[string]interface{}{
		"app_id":     l.config.AppConfig[code].AppId,
		"app_secret": l.config.AppConfig[code].AppSecret,
	}
	bodyBytes, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	request, err := http.NewRequestWithContext(ctx, "POST", url, bytes.NewBuffer(bodyBytes))
	if err != nil {
		return nil, err
	}
	request.Header.Set("Content-Type", "application/json")
	response, err := http.DefaultClient.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	respBody, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	respData := &vo.LarkTenantTokenVo{}
	err = json.Unmarshal(respBody, respData)
	if err != nil {
		return nil, err
	}
	return respData, nil
}
