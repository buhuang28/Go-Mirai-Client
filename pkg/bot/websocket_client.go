package bot

import (
	"encoding/json"
	"fmt"
	"github.com/Mrs4s/MiraiGo/client"
	"github.com/gorilla/websocket"
	"net/http"
	"sync"
	"time"
)

var (
	WsCon          *websocket.Conn
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
	var err error
	for {
		if ConSucess {
			break
		}
		WSClientHeader.Add("origin", WSClientOrigin)

		WsCon, _, err = websocket.DefaultDialer.Dial(WSServerAddr, WSClientHeader)
		if err != nil || WsCon == nil {
			fmt.Println(err)
		} else {
			Clients.Range(func(_ int64, cli *client.QQClient) bool {
				if cli.Online {
					BuhuangBotOnline(cli.Uin)
				}
				return true
			})
			ConSucess = true
			go func() {
				HandleWSMsg()
			}()
			return
		}
		time.Sleep(time.Second * 2)
	}
}

//处理Websocket-Server的消息，一般负责调用API
func HandleWSMsg() {
	for {
		_, message, e := WsCon.ReadMessage()
		fmt.Println("收到消息:", string(message))
		if e != nil {
			fmt.Println("出错了")
			time.Sleep(time.Second * 2)
			go func() {
				ConSucess = false
				fmt.Println("ws-server掉线，正在重连")
				WSDailCall()
			}()
			return
		}
		var data GMCWSData
		_ = json.Unmarshal(message, &data)
		BotClientLock.Lock()
		cli := BotClientMap[data.BotId]
		BotClientLock.Unlock()

		miraiMsg := RawMsgToMiraiMsg(cli, data.Message)
		switch data.MsgType {
		case GMC_PRIVATE_MESSAGE, GMC_TEMP_MESSAGE:
			BuHuangSendPrivateMsg(cli, miraiMsg, data.UserId, data.GroupId)
		case GMC_GROUP_MESSAGE:
			BuHuangSendGroupMsg(cli, miraiMsg, data.MessageId, data.GroupId)
		case GMC_WITHDRAW_MESSAGE:
			BuBuhuangWithDrawMsg(cli, data.GroupId, data.MessageId, data.InternalId)
		case GMC_ALLGROUPMEMBER:
			HandleGetAllMember(cli)
		case GMC_GROUP_LIST:
			HandleGroupList(cli)
		case GMC_KICK:
			BuhuangKickGroupMember(cli, data.GroupId, data.UserId)
		case GMC_BAN:
			BuhuangBanGroupMember(cli, data.GroupId, data.UserId, data.Time)
		}
		//WsCon.WriteMessage(websocket.TextMessage, []byte("message"))
	}
}
