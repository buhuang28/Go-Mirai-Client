package bot

type GMCWSData struct {
	MsgType            int64                 `json:"msg_type"`
	BotId              int64                 `json:"bot_id"`
	Message            string                `json:"message"`
	GroupId            int64                 `json:"group_id"`
	UserId             int64                 `json:"user_id"`
	MessageId          int64                 `json:"message_id"`
	InternalId         int32                 `json:"internal_id"`
	MemberList         []GMCMember           `json:"member_list,omitempty"`
	GroupList          []GMCGroup            `json:"group_list,omitempty"`
	ManageGroup        []int64               `json:"manage_group,omitempty"`
	GroupAndMemberList map[int64][]GMCMember `json:"group_and_member_list,omitempty"`
}

type GMCMember struct {
	QQ         int64  `json:"qq"`
	Permission int64  `json:"permission"`
	NickName   string `json:"nick_name,omitempty"`
	Level      uint16 `json:"lv"`
}

type GMCAllGroupMember struct {
	Data map[int64][]GMCMember `json:"data"`
}

type GMCGroup struct {
	GroupId   int64  `json:"group_id"`
	GroupName string `json:"group_name,omitempty"`
}

const (
	GMC_PRIVATE_MESSAGE  = 1
	GMC_GROUP_MESSAGE    = 2
	GMC_TEMP_MESSAGE     = 3
	GMC_WITHDRAW_MESSAGE = 4
	GMC_ONLINE           = 5
	GMC_OFFLINE          = 6
	//全部群成员
	GMC_ALLGROUPMEMBER = 7
	//群列表
	GMC_GROUP_LIST = 8
)
