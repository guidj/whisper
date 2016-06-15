package main

import (
	"log"
	"net"
	"time"
	"encoding/json"
	"runtime"
	
	"github.com/guidj/whisper/lib"
)

func main() {
	ping(lib.SrvAddr)
}

func ping(address string) {
	addr, err := net.ResolveUDPAddr("udp", address)
	if err != nil {
		log.Fatal(err)
	}
	c, err := net.DialUDP("udp", nil, addr)
	
	var b []byte
	payload := lib.Payload{OS: runtime.GOOS, ARCH: runtime.GOARCH}
	b, err = json.Marshal(payload)
	if err != nil {
		log.Fatalln(err)
	}
	
	for {
		c.Write(b)
		log.Println("Sent ping to", addr.String())
		time.Sleep(1 * time.Second)
	}
}