package bot

type GMCWSData struct {
	MsgType    int64  `json:"msg_type"`
	BotId      int64  `json:"bot_id"`
	Message    string `json:"message"`
	GroupId    int64  `json:"group_id"`
	UserId     int64  `json:"user_id"`
	MessageId  int64  `json:"message_id"`
	InternalId int32  `json:"internal_id"`
}

const (
	GMC_PRIVATE_MESSAGE  = 1
	GMC_GROUP_MESSAGE    = 2
	GMC_TEMP_MESSAGE     = 3
	GMC_WITHDRAW_MESSAGE = 4
)
