package main

import (
	"syscall"
	"log"
)

func main(){
	var sysInfo *syscall.SysInfo_t
	
	err := syscall.Sysinfo(sysInfo)
	if err != nil {
		log.Fatal(err)
	}
}