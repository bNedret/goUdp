package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net"
	"time"
)

func Broadcast(send chan Data){

	conn, err := net.DialUDP("udp", nil, &net.UDPAddr{IP:[]byte{255,255,255,255}, Port:17201, Zone:""})
	if err != nil{
		panic("bye bye from broadcast")
	}
	defer conn.Close()

	var buffer bytes.Buffer
	encoder := json.NewEncoder(&buffer)
	//buf :=make([]byte, 1024)
	for {
		message := <-send
		//fmt.Println(message)
		encoder.Encode(message)
		_, err = conn.Write(buffer.Bytes())
		if err != nil {
			fmt.Println("error")
		}
		fmt.Println("data sent")

		buffer.Reset()
		time.Sleep(time.Second*1)
		//fmt.Println("timer1 finsihed")
	}
}

