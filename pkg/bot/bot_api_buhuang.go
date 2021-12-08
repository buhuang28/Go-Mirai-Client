package bot

import (
	"fmt"
	"github.com/Mrs4s/MiraiGo/client"
	"github.com/Mrs4s/MiraiGo/message"
	log "github.com/sirupsen/logrus"
	"sync"
)

var (
	GroupMsgRecordMap  = make(map[int64][]MessageRecord)
	GroupMsgRecordLock sync.Mutex
)

func BuHuangSendGroupMsg(cli *client.QQClient, miraiMsg []message.IMessageElement, msgId, groupId int64) int64 {
	if cli.FindGroup(groupId) == nil {
		return 0
	}
	sendingMessage := &message.SendingMessage{Elements: miraiMsg}
	preProcessGroupSendingMessage(cli, groupId, sendingMessage)
	if len(sendingMessage.Elements) == 0 {
		log.Warnf("发送消息内容为空")
		return -1
	}
	//是否切片
	GroupMsgRecordLock.Lock()
	defer func() {
		e := recover()
		if e != nil {
			fmt.Println(e)
		}
		GroupMsgRecordLock.Unlock()
	}()
	var record MessageRecord
	ret := cli.SendGroupMessage(groupId, sendingMessage, false)
	record.EventId = ret.Id
	record.InternalId = ret.InternalId
	record.GroupCode = ret.GroupCode
	recordList := GroupMsgRecordMap[msgId]
	if len(recordList) == 0 {
		var newRecordList []MessageRecord
		recordList = newRecordList
	}
	recordList = append(recordList, record)
	GroupMsgRecordMap[msgId] = recordList
	return int64(ret.Id)
}

func BuHuangSendPrivateMsg(cli *client.QQClient, miraiMsg []message.IMessageElement, userId, groupId int64) int64 {
	sendingMessage := &message.SendingMessage{Elements: miraiMsg}
	if userId != 0 { // 私聊+临时
		preProcessPrivateSendingMessage(cli, userId, sendingMessage)
	}
	if groupId != 0 && userId != 0 { // 临时
		ret := cli.SendGroupTempMessage(groupId, userId, sendingMessage)
		return int64(ret.Id)
	}
	preProcessPrivateSendingMessage(cli, userId, sendingMessage)
	ret := cli.SendPrivateMessage(userId, sendingMessage)
	return int64(ret.Id)
}

func BuBuhuangWithDrawMsg(cli *client.QQClient, groupId, msgId int64, internalId int32) {
	if msgId > 0 && internalId > 0 && groupId > 0 {
		_ = cli.RecallGroupMessage(groupId, int32(msgId), internalId)
	}

	GroupMsgRecordLock.Lock()
	defer func() {
		e := recover()
		if e != nil {
			fmt.Println(e)
		}
		GroupMsgRecordLock.Unlock()
	}()
	recordList := GroupMsgRecordMap[msgId]
	for _, v := range recordList {
		_ = cli.RecallGroupMessage(v.GroupCode, v.EventId, v.InternalId)
	}
}
