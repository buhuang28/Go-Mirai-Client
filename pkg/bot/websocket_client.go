package bot

import (
	"encoding/json"
	"fmt"
	"github.com/Mrs4s/MiraiGo/client"
	"github.com/ProtobufBot/Go-Mirai-Client/pkg/ws_data"
	"github.com/gorilla/websocket"
	log "github.com/sirupsen/logrus"
	"net/http"
	"sync"
	"time"
)

var (
	WsCon          *websocket.Conn
	WSWLock        sync.Mutex
	WSRLock        sync.Mutex
	ConSucess      bool        = false
	WSClientHeader http.Header = make(map[string][]string)
	BotClientMap               = make(map[int64]*client.QQClient)
	BotClientLock  sync.Mutex
)

const (
	WSServerAddr   = "ws://127.0.0.1:9801/gmc_event"
	WSClientOrigin = "http://127.0.0.1:9801"
)

func WSDailCall() {
	defer func() {
		e := recover()
		if e != nil {
			ws_data.PrintStackTrace(e)
			ConSucess = false
			WsCon = nil
			go WSDailCall()
			return
		}
	}()
	var err error
	for {
		if ConSucess {
			break
		}
		WSClientHeader.Add("origin", WSClientOrigin)
		WsCon, _, err = websocket.DefaultDialer.Dial(WSServerAddr, WSClientHeader)
		if err != nil || WsCon == nil {
			log.Infof("ws连接出错:", err)
		} else {
			Clients.Range(func(_ int64, cli *client.QQClient) bool {
				if cli.Online.Load() {
					fmt.Println(cli.Uin, "发送上线事件")
					BuhuangBotOnline(cli.Uin)
				}
				return true
			})
			ConSucess = true
			return
		}
		time.Sleep(time.Second * 2)
	}
}

//处理Websocket-Server的消息，一般负责调用API
func HandleWSMsg() {
	defer func() {
		e := recover()
		if e != nil {
			ws_data.PrintStackTrace(e)
		}
	}()
	for {
		if WsCon == nil || !ConSucess {
			time.Sleep(time.Second)
			continue
		}
		WSRLock.Lock()
		_, message, e := WsCon.ReadMessage()
		WSRLock.Unlock()
		log.Println("收到消息:", string(message))
		if e != nil {
			log.Println("出错了：", e)
			time.Sleep(time.Second * 2)
			go func() {
				ConSucess = false
				log.Println("ws-server掉线，正在重连")
				WSDailCall()
			}()
			continue
		}
		var data ws_data.GMCWSData
		_ = json.Unmarshal(message, &data)
		BotClientLock.Lock()
		cli := BotClientMap[data.BotId]
		BotClientLock.Unlock()

		miraiMsg := RawMsgToMiraiMsg(cli, data.Message)
		switch data.MsgType {
		case ws_data.GMC_PRIVATE_MESSAGE, ws_data.GMC_TEMP_MESSAGE:
			BuHuangSendPrivateMsg(cli, miraiMsg, data.UserId, data.GroupId)
		case ws_data.GMC_GROUP_MESSAGE:
			BuHuangSendGroupMsg(cli, miraiMsg, data.MessageId, data.GroupId)
		case ws_data.GMC_WITHDRAW_MESSAGE:
			BuBuhuangWithDrawMsg(cli, data.GroupId, data.MessageId, data.InternalId)
		case ws_data.GMC_ALLGROUPMEMBER:
			HandleGetAllMember(cli)
		case ws_data.GMC_GROUP_LIST:
			HandleGroupList(cli)
		case ws_data.GMC_KICK:
			BuhuangKickGroupMember(cli, data.GroupId, data.UserId)
		case ws_data.GMC_BAN:
			BuhuangBanGroupMember(cli, data.GroupId, data.UserId, data.Time)
		case ws_data.GMC_GROUP_FILE:
			BuhuangUploadGroupFile(cli, data.GroupId, data.Message, data.FilePath)
		case ws_data.GMC_GROUP_REQUEST, ws_data.GMC_BOT_INVITED:
			ws_data.HandleCallBackEvent(data)
		}
	}
}
