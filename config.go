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
