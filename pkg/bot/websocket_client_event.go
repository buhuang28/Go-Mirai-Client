package bot

import (
	"encoding/json"
	"github.com/Mrs4s/MiraiGo/client"
	"github.com/ProtobufBot/Go-Mirai-Client/pkg/ws_data"
	log "github.com/sirupsen/logrus"
)

func HandleGetAllMember(cli *client.QQClient) {
	if WsCon == nil {
		log.Info("HandleGetAllMember-WsCon链接为null")
	}
	var data ws_data.GMCWSData
	data.BotId = cli.Uin
	data.MsgType = ws_data.GMC_ALLGROUPMEMBER
	list := BuhuangGetAllGroupListAndMemberList(cli)
	data.AllGroupMember = list
	marshal, err := json.Marshal(data)
	if err != nil {
		log.Info(cli.Uin, "获取全部群成员列表失败:", err)
	} else {
		//增加Websocket写入互斥锁
		//WSWLock.Lock()
		//defer WSWLock.Unlock()
		WsCon.Write(marshal)
	}
}

func HandleGroupList(cli *client.QQClient) {
	if WsCon == nil {
		log.Info("HandleGroupList-WsCon链接为null")
	}
	var data ws_data.GMCWSData
	data.BotId = cli.Uin
	data.MsgType = ws_data.GMC_GROUP_LIST
	list := BuhuangGetGroupList(cli)
	data.GroupList = list
	marshal, err := json.Marshal(data)
	if err != nil {
		log.Info("%d获取群列表失败", cli.Uin)
	} else {
		//增加Websocket写入互斥锁
		//WSWLock.Lock()
		//defer WSWLock.Unlock()
		WsCon.Write(marshal)
		//WsCon.WriteMessage(websocket.TextMessage, marshal)
	}
}
