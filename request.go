package pinduoduo

import (
	"context"
	"go.dtapp.net/gorequest"
	"go.dtapp.net/gostring"
)

func (c *Client) request(ctx context.Context, params map[string]interface{}) (gorequest.Response, error) {

	// 签名
	c.Sign(params)

	// 创建请求
	client := c.requestClient

	// 设置参数
	client.SetParams(params)

	// 发起请求
	request, err := client.Get(ctx)
	if err != nil {
		return gorequest.Response{}, err
	}

	// 日志
	if c.log.gorm {
		go c.log.logGormClient.GormMiddlewareCustom(ctx, gostring.ToString(params["type"]), request, Version)
	}
	if c.log.mongo {
		go c.log.logMongoClient.MongoMiddlewareCustom(ctx, gostring.ToString(params["type"]), request, Version)
	}

	return request, err
}
