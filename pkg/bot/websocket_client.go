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
	WsConSucess    bool = false
	WSWLock        sync.Mutex
	WSRLock        sync.Mutex
	WSCallLock     sync.Mutex
	WSClientHeader http.Header = make(map[string][]string)
	//BotClientMap               = make(map[int64]*client.QQClient)
	//BotClientLock  sync.Mutex
	SendLock sync.Mutex
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
			go func() {
				log.Info("WsConSucess:", WsConSucess)
				WSDailCall()
			}()
		}
		WSCallLock.Unlock()
	}()
	WSCallLock.Lock()
	var tempHeader http.Header = make(map[string][]string)
	tempHeader.Add("origin", WSClientOrigin)
	WSClientHeader = tempHeader
	var err error
	for {
		if WsCon != nil && WsConSucess {
			return
		}
		fmt.Println("开始连接")
		WsCon, _, err = websocket.DefaultDialer.Dial(WSServerAddr, WSClientHeader)
		if err != nil || WsCon == nil {
			log.Infof("ws连接出错:", err)
			time.Sleep(time.Second * 2)
			continue
		} else {
			WsConSucess = true
			time.Sleep(time.Second)
			Clients.Range(func(_ int64, cli *client.QQClient) bool {
				if cli.Online.Load() {
					fmt.Println(cli.Uin, "发送上线事件")
					BuhuangBotOnline(cli.Uin)
				}
				return true
			})
			return
		}
	}
}

//处理Websocket-Server的消息，一般负责调用API
func HandleWSMsg() {
	defer func() {
		e := recover()
		if e != nil {
			ws_data.PrintStackTrace(e)
		}
		go func() {
			HandleWSMsg()
		}()
	}()
	for {
		if WsCon == nil && !WsConSucess {
			time.Sleep(time.Second)
			continue
		}
		WSRLock.Lock()
		_, message, e := WsCon.ReadMessage()
		WSRLock.Unlock()
		if e != nil || WsCon == nil {
			log.Println("出错了：", e)
			time.Sleep(time.Second * 2)
			WsConSucess = false
			go func() {
				log.Println("ws-server掉线，正在重连")
				WSDailCall()
			}()
			continue
		}
		go func() {
			var data ws_data.GMCWSData
			_ = json.Unmarshal(message, &data)
			cli, ok := Clients.Load(data.BotId)
			if !ok {
				log.Info("加载QQ Cli不存在:", data.BotId)
				return
			}
			//BotClientLock.Lock()
			//cli := BotClientMap[data.BotId]
			//BotClientLock.Unlock()
			miraiMsg := RawMsgToMiraiMsg(cli, data.Message)
			switch data.MsgType {
			case ws_data.GMC_PRIVATE_MESSAGE, ws_data.GMC_TEMP_MESSAGE:
				BuHuangSendPrivateMsg(cli, miraiMsg, data.UserId, data.GroupId)
			case ws_data.GMC_GROUP_MESSAGE:
				nt := time.Now().Unix()
				retId := BuHuangSendGroupMsg(cli, miraiMsg, data.MessageId, data.GroupId)
				nt2 := time.Now().Unix()
				if nt2-nt > 3 {
					log.Info("发送消息超时:", data.Message, "\r\n返回id为:", retId)
				}
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
		}()
	}
}
