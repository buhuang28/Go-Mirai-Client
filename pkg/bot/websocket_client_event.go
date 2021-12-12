package bot

import (
	"encoding/json"
	"github.com/Mrs4s/MiraiGo/client"
	"github.com/gorilla/websocket"
	log "github.com/sirupsen/logrus"
)

func HandleGetAllMember(cli *client.QQClient) {
	var data GMCWSData
	data.BotId = cli.Uin
	data.MsgType = GMC_ALLGROUPMEMBER
	list := BuhuangGetAllGroupListAndMemberList(cli)
	data.AllGroupMember = list
	marshal, err := json.Marshal(data)
	if err != nil {
		log.Info(cli.Uin, "获取全部群成员列表失败:", err)
	} else {
		WsCon.WriteMessage(websocket.TextMessage, marshal)
	}
}

func HandleGroupList(cli *client.QQClient) {
	var data GMCWSData
	data.BotId = cli.Uin
	data.MsgType = GMC_GROUP_LIST
	list := BuhuangGetGroupList(cli)
	data.GroupList = list
	marshal, err := json.Marshal(data)
	if err != nil {
		log.Info("%d获取群列表失败", cli.Uin)
	} else {
		WsCon.WriteMessage(websocket.TextMessage, marshal)
	}
}
