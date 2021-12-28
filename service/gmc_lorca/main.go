package main

import (
	"fmt"
	"github.com/ProtobufBot/Go-Mirai-Client/pkg/ws_data"
	log "github.com/sirupsen/logrus"
	"os"

	"github.com/ProtobufBot/Go-Mirai-Client/pkg/config"
	"github.com/ProtobufBot/Go-Mirai-Client/pkg/gmc"
	"github.com/ProtobufBot/Go-Mirai-Client/pkg/util"

	"github.com/zserge/lorca"
)

func main() {
	defer func() {
		e := recover()
		if e != nil {
			ws_data.PrintStackTrace(e)
		}
	}()
	gmc.Start()
	ui, err := lorca.New(fmt.Sprintf("http://localhost:%s", config.Port), "", 1024, 768)
	if err != nil {
		util.FatalError(err)
		return
	}
	defer func() {
		log.Info("UI EXIT.")
		ui.Close()
	}()
	<-ui.Done()
}

func WriteFile(fileName, content string) bool {
	fd, _ := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	buf := []byte(content)
	_, err := fd.Write(buf)
	fd.Close()
	if err == nil {
		return true
	}
	return false
}
