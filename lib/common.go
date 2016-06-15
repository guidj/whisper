package lib

import (
	"net"
	"time"
)

const (
	SrvAddr         = "224.0.0.1:46789"
	MaxDatagramSize = 2048
)

type Client struct {
	Host *net.UDPAddr
	LastPing time.Time
	Payload *Payload
}

type Payload struct {
	OS string	`json:"os"`
	ARCH string `json:"arch"`
}

