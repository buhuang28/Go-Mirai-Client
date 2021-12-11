package bot

import (
	"encoding/json"
	"github.com/Mrs4s/MiraiGo/client"
	"github.com/gorilla/websocket"
	log "github.com/sirupsen/logrus"
)

func HandleGetAllMember(cli *client.QQClient) {
	list := BuhuangGetAllGroupListAndMemberList(cli)
	marshal, err := json.Marshal(list)
	if err != nil {
		WsCon.WriteMessage(websocket.TextMessage, marshal)
	} else {
		log.Info("%d获取全部群成员列表失败", cli.Uin)
	}
}

func HandleGroupList(cli *client.QQClient) {
	list := BuhuangGetAllGroupListAndMemberList(cli)
	marshal, err := json.Marshal(list)
	if err != nil {
		WsCon.WriteMessage(websocket.TextMessage, marshal)
	} else {
		log.Info("%d获取群列表失败", cli.Uin)
	}
}
