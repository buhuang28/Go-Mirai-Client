package plugin

import (
	"encoding/json"
	"github.com/Mrs4s/MiraiGo/client"
	"github.com/Mrs4s/MiraiGo/message"
	"github.com/ProtobufBot/Go-Mirai-Client/pkg/bot"
	"github.com/ProtobufBot/Go-Mirai-Client/pkg/util"
	"github.com/gorilla/websocket"
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
	//群成员++
	cli.OnGroupMemberJoined(handleMemberJoinGroup)
	//群成员--
	cli.OnGroupMemberLeaved(handleMemberLeaveGroup)
	//cli.OnJoinGroup(handleJoinGroup)
	//cli.OnLeaveGroup(handleLeaveGroup)
	//cli.OnNewFriendRequest(handleNewFriendRequest)

	//入群请求
	cli.OnUserWantJoinGroup(handleUserJoinGroupRequest)
	//机器人被邀请入群
	cli.OnGroupInvited(handleGroupInvitedRequest)

	//system,_ := cli.GetGroupFileSystem(123)
	//system.UploadFile("filePath","name","/")
	//url := cli.GetGroupFileUrl(groupID, fileID, busID)

	//cli.OnGroupMessageRecalled(handleGroupMessageRecalled)
	//cli.OnFriendMessageRecalled(handleFriendMessageRecalled)
	//cli.OnNewFriendAdded(handleNewFriendAdded)
	//cli.OnReceivedOfflineFile(handleOfflineFile)
	//cli.OnGroupMuted(handleGroupMute)
	//cli.OnGroupMemberPermissionChanged(handleMemberPermissionChanged)
}

// 添加私聊消息插件
func AddPrivateMessagePlugin(plugin PrivateMessagePlugin) {
	PrivateMessagePluginList = append(PrivateMessagePluginList, plugin)
}

// 添加群聊消息插件
func AddGroupMessagePlugin(plugin GroupMessagePlugin) {
	GroupMessagePluginList = append(GroupMessagePluginList, plugin)
}

// 添加临时消息插件
func AddTempMessagePlugin(plugin TempMessagePlugin) {
	TempMessagePluginList = append(TempMessagePluginList, plugin)
}

// 添加群成员加入插件
func AddMemberJoinGroupPlugin(plugin MemberJoinGroupPlugin) {
	MemberJoinGroupPluginList = append(MemberJoinGroupPluginList, plugin)
}

// 添加群成员离开插件
func AddMemberLeaveGroupPlugin(plugin MemberLeaveGroupPlugin) {
	MemberLeaveGroupPluginList = append(MemberLeaveGroupPluginList, plugin)
}

// 添加机器人进群插件
func AddJoinGroupPlugin(plugin JoinGroupPlugin) {
	JoinGroupPluginList = append(JoinGroupPluginList, plugin)
}

// 添加机器人离开群插件
func AddLeaveGroupPlugin(plugin LeaveGroupPlugin) {
	LeaveGroupPluginList = append(LeaveGroupPluginList, plugin)
}

// 添加好友请求处理插件
func AddNewFriendRequestPlugin(plugin NewFriendRequestPlugin) {
	NewFriendRequestPluginList = append(NewFriendRequestPluginList, plugin)
}

// 添加加群请求处理插件
func AddUserJoinGroupRequestPlugin(plugin UserJoinGroupRequestPlugin) {
	UserJoinGroupRequestPluginList = append(UserJoinGroupRequestPluginList, plugin)
}

// 添加机器人被邀请处理插件
func AddGroupInvitedRequestPlugin(plugin GroupInvitedRequestPlugin) {
	GroupInvitedRequestPluginList = append(GroupInvitedRequestPluginList, plugin)
}

// 添加群消息撤回处理插件
func AddGroupMessageRecalledPlugin(plugin GroupMessageRecalledPlugin) {
	GroupMessageRecalledPluginList = append(GroupMessageRecalledPluginList, plugin)
}

// 添加好友消息撤回处理插件
func AddFriendMessageRecalledPlugin(plugin FriendMessageRecalledPlugin) {
	FriendMessageRecalledPluginList = append(FriendMessageRecalledPluginList, plugin)
}

// 添加好友添加处理插件
func AddNewFriendAddedPlugin(plugin NewFriendAddedPlugin) {
	NewFriendAddedPluginList = append(NewFriendAddedPluginList, plugin)
}

// 添加离线文件处理插件
func AddOfflineFilePlugin(plugin OfflineFilePlugin) {
	OfflineFilePluginList = append(OfflineFilePluginList, plugin)
}

// 添加群成员被禁言插件
func AddGroupMutePlugin(plugin GroupMutePlugin) {
	GroupMutePluginList = append(GroupMutePluginList, plugin)
}

// 添加群成员权限变动插件
func AddMemberPermissionChangedPlugin(plugin MemberPermissionChangedPlugin) {
	MemberPermissionChangedPluginList = append(MemberPermissionChangedPluginList, plugin)
}

func handlePrivateMessage(cli *client.QQClient, event *message.PrivateMessage) {
	//fmt.Println("接收到私聊消息:",event.ToString())
	//fmt.Println("接收到私聊消息:",bot.MiraiMsgToRawMsg(cli, event.Elements))
	//util.SafeGo(func() {
	//	if ws_client.WsCon == nil || !ws_client.ConSucess {
	//		return
	//	}
	//	ws_client.WsCon.WriteMessage(websocket.TextMessage, []byte(event.ToString()))
	//})
	if bot.WsCon == nil || !bot.ConSucess {
		return
	}
	util.SafeGo(func() {
		var data bot.GMCWSData
		data.BotId = cli.Uin
		data.UserId = event.Sender.Uin
		data.MsgType = bot.GMC_PRIVATE_MESSAGE
		data.Message = bot.MiraiMsgToRawMsg(cli, event.Elements)
		marshal, _ := json.Marshal(data)
		bot.WsCon.WriteMessage(websocket.TextMessage, marshal)
	})
}

func handleGroupMessage(cli *client.QQClient, event *message.GroupMessage) {
	if bot.WsCon == nil || !bot.ConSucess {
		return
	}
	util.SafeGo(func() {
		var data bot.GMCWSData
		data.BotId = cli.Uin
		data.GroupId = event.GroupCode
		data.UserId = event.Sender.Uin
		data.MsgType = bot.GMC_GROUP_MESSAGE
		data.MessageId = int64(event.Id)
		data.InternalId = event.InternalId
		data.Message = bot.MiraiMsgToRawMsg2(cli, event.GroupCode, event.Elements)
		marshal, _ := json.Marshal(data)
		bot.WsCon.WriteMessage(websocket.TextMessage, marshal)
	})
}

func handleTempMessage(cli *client.QQClient, event *client.TempMessageEvent) {
	if bot.WsCon == nil || !bot.ConSucess {
		return
	}
	util.SafeGo(func() {
		var data bot.GMCWSData
		data.BotId = cli.Uin
		data.GroupId = event.Message.GroupCode
		data.UserId = event.Message.Sender.Uin
		data.MsgType = bot.GMC_TEMP_MESSAGE
		data.Message = bot.MiraiMsgToRawMsg(cli, event.Message.Elements)
		marshal, _ := json.Marshal(data)
		bot.WsCon.WriteMessage(websocket.TextMessage, marshal)
	})
}

func handleMemberJoinGroup(cli *client.QQClient, event *client.MemberJoinGroupEvent) {
	util.SafeGo(func() {
		for _, plugin := range MemberJoinGroupPluginList {
			if result := plugin(cli, event); result == MessageBlock {
				break
			}
		}
	})
}

func handleMemberLeaveGroup(cli *client.QQClient, event *client.MemberLeaveGroupEvent) {
	util.SafeGo(func() {
		for _, plugin := range MemberLeaveGroupPluginList {
			if result := plugin(cli, event); result == MessageBlock {
				break
			}
		}
	})
}

func handleJoinGroup(cli *client.QQClient, event *client.GroupInfo) {
	util.SafeGo(func() {
		for _, plugin := range JoinGroupPluginList {
			if result := plugin(cli, event); result == MessageBlock {
				break
			}
		}
	})
}

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

func handleUserJoinGroupRequest(cli *client.QQClient, event *client.UserJoinGroupRequest) {
	util.SafeGo(func() {
		for _, plugin := range UserJoinGroupRequestPluginList {
			if result := plugin(cli, event); result == MessageBlock {
				break
			}
		}
	})
}

func handleGroupInvitedRequest(cli *client.QQClient, event *client.GroupInvitedRequest) {
	util.SafeGo(func() {
		for _, plugin := range GroupInvitedRequestPluginList {
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
