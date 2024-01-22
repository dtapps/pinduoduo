package pinduoduo

import (
	"go.dtapp.net/golog"
)

func (c *Client) ConfigPid(pid string) *Client {
	c.config.pid = pid
	return c
}

// ConfigApiGormFun 接口日志配置
func (c *Client) ConfigApiGormFun(apiClientFun golog.ApiGormFun) {
	client := apiClientFun()
	if client != nil {
		c.gormLog.client = client
		c.gormLog.status = true
	}
}

// ConfigApiMongoFun 接口日志配置
func (c *Client) ConfigApiMongoFun(apiClientFun golog.ApiMongoFun) {
	client := apiClientFun()
	if client != nil {
		c.mongoLog.client = client
		c.mongoLog.status = true
	}
}
