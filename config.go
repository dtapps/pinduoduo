package pinduoduo

func (c *Client) ConfigPid(pid string) *Client {
	c.config.pid = pid
	return c
}
