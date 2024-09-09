package utils

import "fmt"

type ImapConfigs struct {
	ServerUrl  string
	ServerPort int
	Username   string
	Password   string
}

func (c *ImapConfigs) Address() string {
	return fmt.Sprintf("%s:%d", c.ServerUrl, c.ServerPort)
}
