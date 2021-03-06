package main

import (
	"log"
	"net"
	"time"
	"net/http"
	"io"
	"encoding/json"
	
	"github.com/guidj/whisper/lib"
)

var clients map[string]*lib.Client

func main() {
	clients = make(map[string]*lib.Client)
	
	//TODO: send received message to queue for processing (with now timestamp); keep one writer, scale in receiving messages. How?	
	go serveMulticastUDP(lib.SrvAddr, msgHandler)
	http.HandleFunc("/hosts", whisperServer)
	log.Fatal(http.ListenAndServe(":46790", nil))
}


func msgHandler(src *net.UDPAddr, n int, b []byte) {
	payload := lib.Payload{}
	err := json.Unmarshal(b[:n], &payload)
	
	if err != nil {
		log.Println(err)
		return
	}
	
	c, ok := clients[src.String()]
	now := time.Now()
	
	if ok != true {
		log.Println("Received ping from", src.String(), "at", now)
		clients[src.String()] = &lib.Client{Host: src, LastPing: now, Payload: &payload}
	}else {
		log.Println("Received ping from", c.Host.String(), "at", now, ", after", time.Since(c.LastPing).Seconds(), "s")
		c.LastPing = now
	}
}

func serveMulticastUDP(a string, h func(*net.UDPAddr, int, []byte)) {
	addr, err := net.ResolveUDPAddr("udp", a)
	if err != nil {
		log.Fatal(err)
	}
	
	l, err := net.ListenMulticastUDP("udp", nil, addr)
	l.SetReadBuffer(lib.MaxDatagramSize)
	for {
		b := make([]byte, lib.MaxDatagramSize)
		n, src, err := l.ReadFromUDP(b)
		if err != nil {
			log.Fatal("ReadFromUDP failed:", err)
		}
		h(src, n, b)
	}
}

func whisperServer(w http.ResponseWriter, req *http.Request) {

	data, err := json.Marshal(clients)
	if err != nil {
		log.Println(err)
		http.Error(w, "Error reading hosts index", http.StatusInternalServerError)
		return
	}
	
	io.WriteString(w, string(data))
}