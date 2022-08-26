package pinduoduo

func (c *Client) ConfigPid(pid string) {
	c.config.pid = pid
	return
}
