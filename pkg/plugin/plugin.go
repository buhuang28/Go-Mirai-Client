package plugin

import (
	"encoding/json"
	"github.com/Mrs4s/MiraiGo/client"
	"github.com/Mrs4s/MiraiGo/message"
	"github.com/ProtobufBot/Go-Mirai-Client/pkg/bot"
	"github.com/ProtobufBot/Go-Mirai-Client/pkg/util"
	"github.com/ProtobufBot/Go-Mirai-Client/pkg/ws_data"
	log "github.com/sirupsen/logrus"
	"time"
)

type (
	PrivateMessagePlugin          = func(*client.QQClient, *message.PrivateMessage) int32
	GroupMessagePlugin            = func(*client.QQClient, *message.GroupMessage) int32
	TempMessagePlugin             = func(*client.QQClient, *client.TempMessageEvent) int32
	MemberJoinGroupPlugin         = func(*client.QQClient, *client.MemberJoinGroupEvent) int32
	MemberLeaveGroupPlugin        = func(*client.QQClient, *client.MemberLeaveGroupEvent) int32
	JoinGroupPlugin               = func(*client.QQClient, *client.GroupInfo) int32
	LeaveGroupPlugin              = func(*client.QQClient, *client.GroupLeaveEvent) int32
	NewFriendRequestPlugin        = func(*client.QQClient, *client.NewFriendRequest) int32
	UserJoinGroupRequestPlugin    = func(*client.QQClient, *client.UserJoinGroupRequest) int32
	GroupInvitedRequestPlugin     = func(*client.QQClient, *client.GroupInvitedRequest) int32
	GroupMessageRecalledPlugin    = func(*client.QQClient, *client.GroupMessageRecalledEvent) int32
	FriendMessageRecalledPlugin   = func(*client.QQClient, *client.FriendMessageRecalledEvent) int32
	NewFriendAddedPlugin          = func(*client.QQClient, *client.NewFriendEvent) int32
	OfflineFilePlugin             = func(*client.QQClient, *client.OfflineFileEvent) int32
	GroupMutePlugin               = func(*client.QQClient, *client.GroupMuteEvent) int32
	MemberPermissionChangedPlugin = func(*client.QQClient, *client.MemberPermissionChangedEvent) int32
)

const (
	MessageIgnore = 0
	MessageBlock  = 1
)

var (
	REQUEST_ACCEPT int64 = 1
	REQUEST_REJECT int64 = -1
)

var PrivateMessagePluginList = make([]PrivateMessagePlugin, 0)
var GroupMessagePluginList = make([]GroupMessagePlugin, 0)
var TempMessagePluginList = make([]TempMessagePlugin, 0)
var MemberJoinGroupPluginList = make([]MemberJoinGroupPlugin, 0)
var MemberLeaveGroupPluginList = make([]MemberLeaveGroupPlugin, 0)
var JoinGroupPluginList = make([]JoinGroupPlugin, 0)
var LeaveGroupPluginList = make([]LeaveGroupPlugin, 0)
var NewFriendRequestPluginList = make([]NewFriendRequestPlugin, 0)
var UserJoinGroupRequestPluginList = make([]UserJoinGroupRequestPlugin, 0)
var GroupInvitedRequestPluginList = make([]GroupInvitedRequestPlugin, 0)
var GroupMessageRecalledPluginList = make([]GroupMessageRecalledPlugin, 0)
var FriendMessageRecalledPluginList = make([]FriendMessageRecalledPlugin, 0)
var NewFriendAddedPluginList = make([]NewFriendAddedPlugin, 0)
var OfflineFilePluginList = make([]OfflineFilePlugin, 0)
var GroupMutePluginList = make([]GroupMutePlugin, 0)
var MemberPermissionChangedPluginList = make([]MemberPermissionChangedPlugin, 0)

//注册事件
func Serve(cli *client.QQClient) {
	cli.OnPrivateMessage(handlePrivateMessage)
	cli.OnGroupMessage(handleGroupMessage)
	cli.OnTempMessage(handleTempMessage)
	//群成员++  (有人入群)
	cli.OnGroupMemberJoined(handleMemberJoinGroup)
	//群成员--  (有人跑路)
	cli.OnGroupMemberLeaved(handleMemberLeaveGroup)
	//入群请求
	cli.OnUserWantJoinGroup(handleUserJoinGroupRequest)
	//机器人被邀请入群
	cli.OnGroupInvited(handleGroupInvitedRequest)

	//cli.OnJoinGroup(handleJoinGroup)  //机器人入群
	//cli.OnLeaveGroup(handleLeaveGroup) //机器人退群
	//cli.OnNewFriendRequest(handleNewFriendRequest)
	//cli.OnGroupMessageRecalled(handleGroupMessageRecalled)
	//cli.OnFriendMessageRecalled(handleFriendMessageRecalled)
	//cli.OnNewFriendAdded(handleNewFriendAdded)
	//cli.OnReceivedOfflineFile(handleOfflineFile)
	//cli.OnGroupMuted(handleGroupMute)
	//cli.OnGroupMemberPermissionChanged(handleMemberPermissionChanged)
}

//// 添加私聊消息插件
//func AddPrivateMessagePlugin(plugin PrivateMessagePlugin) {
//	PrivateMessagePluginList = append(PrivateMessagePluginList, plugin)
//}
//
//// 添加群聊消息插件
//func AddGroupMessagePlugin(plugin GroupMessagePlugin) {
//	GroupMessagePluginList = append(GroupMessagePluginList, plugin)
//}
//
//// 添加临时消息插件
//func AddTempMessagePlugin(plugin TempMessagePlugin) {
//	TempMessagePluginList = append(TempMessagePluginList, plugin)
//}
//
//// 添加群成员加入插件
//func AddMemberJoinGroupPlugin(plugin MemberJoinGroupPlugin) {
//	MemberJoinGroupPluginList = append(MemberJoinGroupPluginList, plugin)
//}
//
//// 添加群成员离开插件
//func AddMemberLeaveGroupPlugin(plugin MemberLeaveGroupPlugin) {
//	MemberLeaveGroupPluginList = append(MemberLeaveGroupPluginList, plugin)
//}
//
//// 添加机器人进群插件
//func AddJoinGroupPlugin(plugin JoinGroupPlugin) {
//	JoinGroupPluginList = append(JoinGroupPluginList, plugin)
//}
//
//// 添加机器人离开群插件
//func AddLeaveGroupPlugin(plugin LeaveGroupPlugin) {
//	LeaveGroupPluginList = append(LeaveGroupPluginList, plugin)
//}
//
//// 添加好友请求处理插件
//func AddNewFriendRequestPlugin(plugin NewFriendRequestPlugin) {
//	NewFriendRequestPluginList = append(NewFriendRequestPluginList, plugin)
//}
//
//// 添加加群请求处理插件
//func AddUserJoinGroupRequestPlugin(plugin UserJoinGroupRequestPlugin) {
//	UserJoinGroupRequestPluginList = append(UserJoinGroupRequestPluginList, plugin)
//}
//
//// 添加机器人被邀请处理插件
//func AddGroupInvitedRequestPlugin(plugin GroupInvitedRequestPlugin) {
//	GroupInvitedRequestPluginList = append(GroupInvitedRequestPluginList, plugin)
//}
//
//// 添加群消息撤回处理插件
//func AddGroupMessageRecalledPlugin(plugin GroupMessageRecalledPlugin) {
//	GroupMessageRecalledPluginList = append(GroupMessageRecalledPluginList, plugin)
//}
//
//// 添加好友消息撤回处理插件
//func AddFriendMessageRecalledPlugin(plugin FriendMessageRecalledPlugin) {
//	FriendMessageRecalledPluginList = append(FriendMessageRecalledPluginList, plugin)
//}
//
//// 添加好友添加处理插件
//func AddNewFriendAddedPlugin(plugin NewFriendAddedPlugin) {
//	NewFriendAddedPluginList = append(NewFriendAddedPluginList, plugin)
//}
//
//// 添加离线文件处理插件
//func AddOfflineFilePlugin(plugin OfflineFilePlugin) {
//	OfflineFilePluginList = append(OfflineFilePluginList, plugin)
//}
//
//// 添加群成员被禁言插件
//func AddGroupMutePlugin(plugin GroupMutePlugin) {
//	GroupMutePluginList = append(GroupMutePluginList, plugin)
//}
//
//// 添加群成员权限变动插件
//func AddMemberPermissionChangedPlugin(plugin MemberPermissionChangedPlugin) {
//	MemberPermissionChangedPluginList = append(MemberPermissionChangedPluginList, plugin)
//}

//私聊消息
func handlePrivateMessage(cli *client.QQClient, event *message.PrivateMessage) {
	//bot.WSWLock.Lock()
	defer func() {
		e := recover()
		if e != nil {
			ws_data.PrintStackTrace(e)
		}
		//bot.WSWLock.Unlock()
	}()
	msg := bot.MiraiMsgToRawMsg(cli, event.Elements)
	log.Info("收到", event.Sender.Uin, "私聊消息:", msg)
	if bot.WsCon == nil {
		return
	}
	cli.MarkPrivateMessageReaded(event.Sender.Uin, int64(event.Time))
	var data ws_data.GMCWSData
	data.BotId = cli.Uin
	data.UserId = event.Sender.Uin
	data.MsgType = ws_data.GMC_PRIVATE_MESSAGE
	data.Message = msg
	marshal, _ := json.Marshal(data)
	//err := bot.WsCon.WriteMessage(websocket.TextMessage, marshal)
	bot.WsCon.Write(marshal)
	//if err != nil {
	//	log.Info("handlePrivateMessage出错", err)
	//}
}

//群消息
func handleGroupMessage(cli *client.QQClient, event *message.GroupMessage) {
	go func() {
		defer func() {
			e := recover()
			if e != nil {
				ws_data.PrintStackTrace(e)
			}
			//bot.WSWLock.Unlock()
		}()
		//bot.WSWLock.Lock()
		msg := bot.MiraiMsgToRawMsg2(cli, event.GroupCode, event.Elements)
		log.Info("收到群聊消息")
		if bot.WsCon == nil {
			log.Infof("WS链接爆炸")
			return
		}
		cli.MarkGroupMessageReaded(event.GroupCode, int64(event.Id))
		var data ws_data.GMCWSData
		data.BotId = cli.Uin
		data.GroupId = event.GroupCode
		data.UserId = event.Sender.Uin
		data.MsgType = ws_data.GMC_GROUP_MESSAGE
		data.MessageId = int64(event.Id)
		data.InternalId = event.InternalId
		data.Message = msg
		marshal, _ := json.Marshal(data)
		//e := bot.WsCon.WriteMessage(websocket.TextMessage, marshal)
		//if e != nil {
		//log.Info("handleGroupMessage错误:", e)
		//}
		bot.WsCon.Write(marshal)
	}()
}

//临时消息
func handleTempMessage(cli *client.QQClient, event *client.TempMessageEvent) {
	//bot.WSWLock.Lock()
	defer func() {
		e := recover()
		if e != nil {
			ws_data.PrintStackTrace(e)
		}
		//bot.WSWLock.Unlock()
	}()
	msg := bot.MiraiMsgToRawMsg(cli, event.Message.Elements)
	//log.Info("收到", event.Message.GroupCode, "群,", event.Message.Sender.Uin, "的临时私聊消息:", msg)
	if bot.WsCon == nil {
		return
	}
	var data ws_data.GMCWSData
	data.BotId = cli.Uin
	data.GroupId = event.Message.GroupCode
	data.UserId = event.Message.Sender.Uin
	data.MsgType = ws_data.GMC_TEMP_MESSAGE
	data.Message = msg
	marshal, _ := json.Marshal(data)
	//err := bot.WsCon.WriteMessage(websocket.TextMessage, marshal)
	//if err != nil {
	//log.Info("handleTempMessage出错:", err)
	//}
	bot.WsCon.Write(marshal)
}

//有人入群
func handleMemberJoinGroup(cli *client.QQClient, event *client.MemberJoinGroupEvent) {
	//bot.WSWLock.Lock()
	defer func() {
		e := recover()
		if e != nil {
			ws_data.PrintStackTrace(e)
		}
		//bot.WSWLock.Unlock()
	}()
	log.Info("收到入群信息")
	if bot.WsCon == nil {
		return
	}
	var data ws_data.GMCWSData
	data.GroupId = event.Group.Code
	data.UserId = event.Member.Uin
	data.NickName = event.Member.Nickname
	data.BotId = cli.Uin
	data.MsgType = ws_data.GMC_MEMBER_ADD
	marshal, _ := json.Marshal(data)
	//err := bot.WsCon.WriteMessage(websocket.TextMessage, marshal)
	//if err != nil {
	//log.Info("handleMemberJoinGroup出错", err)
	//}
	bot.WsCon.Write(marshal)
}

//有人离开
func handleMemberLeaveGroup(cli *client.QQClient, event *client.MemberLeaveGroupEvent) {
	//bot.WSWLock.Lock()
	defer func() {
		e := recover()
		if e != nil {
			ws_data.PrintStackTrace(e)
		}
		//bot.WSWLock.Unlock()
	}()
	log.Info("收到有人退群")
	if bot.WsCon == nil {
		return
	}
	var data ws_data.GMCWSData
	data.GroupId = event.Group.Code
	data.UserId = event.Member.Uin
	data.BotId = cli.Uin
	data.MsgType = ws_data.GMC_MEMBER_LEAVE
	marshal, _ := json.Marshal(data)
	//err := bot.WsCon.WriteMessage(websocket.TextMessage, marshal)
	//if err != nil {
	//	log.Info("handleMemberLeaveGroup出错", err)
	//}
	bot.WsCon.Write(marshal)
}

//有人申请入群
func handleUserJoinGroupRequest(cli *client.QQClient, event *client.UserJoinGroupRequest) {
	defer func() {
		e := recover()
		if e != nil {
			ws_data.PrintStackTrace(e)
		}
	}()
	log.Info("有人申请入群")
	if bot.WsCon == nil {
		return
	}
	var data ws_data.GMCWSData
	data.MsgType = ws_data.GMC_GROUP_REQUEST
	data.BotId = cli.Uin
	data.GroupId = event.GroupCode
	data.UserId = event.RequesterUin
	data.NickName = event.RequesterNick
	data.RequestId = event.RequestId
	data.Message = event.Message
	marshal, _ := json.Marshal(data)
	//bot.WSWLock.Lock()
	bot.WsCon.Write(marshal)
	//err := bot.WsCon.WriteMessage(websocket.TextMessage, marshal)
	//if err != nil {
	//	log.Info("handleUserJoinGroupRequest出错", err)
	//}
	//bot.WSWLock.Unlock()
	ch := make(chan ws_data.GMCWSData, 1)
	ws_data.ChanMapLock.Lock()
	ws_data.ChanMap[event.RequestId] = ch
	ws_data.ChanMapLock.Unlock()
	go func() {
		select {
		case r := <-ch:
			switch r.GroupRequest {
			case REQUEST_ACCEPT:
				event.Accept()
			case REQUEST_REJECT:
				event.Reject(false, "")
			}
		case <-time.After(time.Second * 30):
			delete(ws_data.ChanMap, event.RequestId)
			return
		}
	}()
}

//机器人被邀请入群
func handleGroupInvitedRequest(cli *client.QQClient, event *client.GroupInvitedRequest) {
	defer func() {
		e := recover()
		if e != nil {
			ws_data.PrintStackTrace(e)
		}
	}()
	log.Info("收到机器人被邀请入群")
	if bot.WsCon == nil {
		return
	}
	var data ws_data.GMCWSData
	data.MsgType = ws_data.GMC_BOT_INVITED
	data.BotId = cli.Uin
	data.GroupId = event.GroupCode
	data.NickName = event.InvitorNick
	data.InvitorId = event.InvitorUin
	data.RequestId = event.RequestId
	marshal, _ := json.Marshal(data)
	//bot.WSWLock.Lock()
	bot.WsCon.Write(marshal)
	//err := bot.WsCon.WriteMessage(websocket.TextMessage, marshal)
	//if err != nil {
	//	log.Info("handleGroupInvitedRequest出错", err)
	//}
	//bot.WSWLock.Unlock()
	ch := make(chan ws_data.GMCWSData, 1)
	ws_data.ChanMapLock.Lock()
	ws_data.ChanMap[event.RequestId] = ch
	ws_data.ChanMapLock.Unlock()
	go func() {
		select {
		case r := <-ch:
			switch r.GroupRequest {
			case REQUEST_ACCEPT:
				cli.SolveGroupJoinRequest(event, true, false, "")
			case REQUEST_REJECT:
				cli.SolveGroupJoinRequest(event, false, false, "")
			}
			return
		case <-time.After(time.Second * 30):
			delete(ws_data.ChanMap, event.RequestId)
			return
		}
	}()
}

//机器人退群
func handleLeaveGroup(cli *client.QQClient, event *client.GroupLeaveEvent) {
	util.SafeGo(func() {
		for _, plugin := range LeaveGroupPluginList {
			if result := plugin(cli, event); result == MessageBlock {
				break
			}
		}
	})
}

func handleNewFriendRequest(cli *client.QQClient, event *client.NewFriendRequest) {
	util.SafeGo(func() {
		for _, plugin := range NewFriendRequestPluginList {
			if result := plugin(cli, event); result == MessageBlock {
				break
			}
		}
	})
}

//机器人入群
func handleJoinGroup(cli *client.QQClient, event *client.GroupInfo) {
	util.SafeGo(func() {
		for _, plugin := range JoinGroupPluginList {
			if result := plugin(cli, event); result == MessageBlock {
				break
			}
		}
	})
}

func handleGroupMessageRecalled(cli *client.QQClient, event *client.GroupMessageRecalledEvent) {
	util.SafeGo(func() {
		for _, plugin := range GroupMessageRecalledPluginList {
			if result := plugin(cli, event); result == MessageBlock {
				break
			}
		}
	})
}

func handleFriendMessageRecalled(cli *client.QQClient, event *client.FriendMessageRecalledEvent) {
	util.SafeGo(func() {
		for _, plugin := range FriendMessageRecalledPluginList {
			if result := plugin(cli, event); result == MessageBlock {
				break
			}
		}
	})
}

func handleNewFriendAdded(cli *client.QQClient, event *client.NewFriendEvent) {
	util.SafeGo(func() {
		for _, plugin := range NewFriendAddedPluginList {
			if result := plugin(cli, event); result == MessageBlock {
				break
			}
		}
	})
}

func handleOfflineFile(cli *client.QQClient, event *client.OfflineFileEvent) {
	util.SafeGo(func() {
		for _, plugin := range OfflineFilePluginList {
			if result := plugin(cli, event); result == MessageBlock {
				break
			}
		}
	})
}

func handleGroupMute(cli *client.QQClient, event *client.GroupMuteEvent) {
	util.SafeGo(func() {
		for _, plugin := range GroupMutePluginList {
			if result := plugin(cli, event); result == MessageBlock {
				break
			}
		}
	})
}

func handleMemberPermissionChanged(cli *client.QQClient, event *client.MemberPermissionChangedEvent) {
	util.SafeGo(func() {
		for _, plugin := range MemberPermissionChangedPluginList {
			if result := plugin(cli, event); result == MessageBlock {
				break
			}
		}
	})
}
