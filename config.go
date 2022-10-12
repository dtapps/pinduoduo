package pinduoduo

import "go.dtapp.net/golog"

func (c *Client) ConfigPid(pid string) *Client {
	c.config.pid = pid
	return c
}

// ConfigApiClientFun 日志配置
func (c *Client) ConfigApiClientFun(apiClientFun golog.ApiClientFun) {
	apiClient := apiClientFun()
	if apiClient != nil {
		c.log.client = apiClient
		c.log.status = true
	}
}
