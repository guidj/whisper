package main

import (
	"log"
	"net"
	"time"
	
	"github.com/guidj/whisper/lib"
)

func main() {
	ping(lib.SrvAddr)
}

func ping(a string) {
	addr, err := net.ResolveUDPAddr("udp", a)
	if err != nil {
		log.Fatal(err)
	}
	c, err := net.DialUDP("udp", nil, addr)
	for {
		c.Write([]byte("Whiper\n"))
		log.Println("Sent ping to", addr.String())
		time.Sleep(1 * time.Second)
	}
}