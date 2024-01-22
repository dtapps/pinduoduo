package pinduoduo

import (
	"context"
	"fmt"
	"go.dtapp.net/gorequest"
)

func (c *Client) request(ctx context.Context, param gorequest.Params) (gorequest.Response, error) {

	// 签名
	c.Sign(param)

	// 创建请求
	client := gorequest.NewHttp()

	// 设置参数
	client.SetParams(param)

	// 设置用户代理
	client.SetUserAgent(gorequest.GetRandomUserAgentSystem())

	// 发起请求
	request, err := client.Get(ctx)
	if err != nil {
		return gorequest.Response{}, err
	}

	// 日志
	if c.gormLog.status {
		go c.gormLog.client.MiddlewareCustom(ctx, fmt.Sprintf("%s", param.Get("type")), request)
	}
	if c.mongoLog.status {
		go c.mongoLog.client.MiddlewareCustom(ctx, fmt.Sprintf("%s", param.Get("type")), request)
	}

	return request, err
}
