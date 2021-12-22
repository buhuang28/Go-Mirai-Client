package main

import (
	"github.com/ProtobufBot/Go-Mirai-Client/pkg/gmc"
	"github.com/ProtobufBot/Go-Mirai-Client/pkg/ws_data"
)

func main() {
	defer func() {
		e := recover()
		if e != nil {
			ws_data.PrintStackTrace(e)
		}
	}()
	gmc.Start()
	select {}
}
