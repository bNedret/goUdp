package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net"
	"time"
)

func Listen(receive chan [] Data) {

	serverAdd, err := net.ResolveUDPAddr("udp", ":17201")
	connection, err := net.ListenUDP("udp", serverAdd)
	if err != nil{
		fmt.Println("bye bye from listening")
	}
	defer connection.Close()
	var message Data
	msg := make([]Data, 5)
	var w int
	w = 0

	for {
		inputbytes := make([]byte, 10000)
		i, add, err := connection.ReadFromUDP(inputbytes)
		if err != nil {
			panic("no data")
		}
		fmt.Println("ip: ", net.IP.String(add.IP), " message: ", fmt.Sprintf("%s", string(inputbytes[:i])))
		message.Ip = net.IP.String(add.IP)
		buffer := bytes.NewBuffer(inputbytes[:i])
		decoder := json.NewDecoder(buffer)
		decoder.Decode(&message)
		fmt.Println(message)
		msg[w] = Data{Ip: message.Ip,
			Header: Header{
				message.MsgType,
				message.Device,
				message.FromMac,
				message.ToMac,
			},
			Data1: message.Data1}

		//fmt.Println(i)
		//fmt.Println(msg[w])
		w++
		//fmt.Println(msg)
		receive<-msg
		time.Sleep(time.Millisecond*1)

	}

}
