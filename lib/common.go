package lib

import (
	"net"
	"time"
)

const (
	SrvAddr         = "224.0.0.1:46789"
	MaxDatagramSize = 8192
)

type Client struct {
	Host *net.UDPAddr 	`json:"host"`
	LastPing time.Time	`json:"lastPing"`
}