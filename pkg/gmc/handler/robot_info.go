package handler

import (
	"encoding/json"
	"fmt"
	"github.com/Mrs4s/MiraiGo/client"
	"github.com/ProtobufBot/Go-Mirai-Client/pkg/bot"
	"github.com/ProtobufBot/Go-Mirai-Client/pkg/device"
	"strconv"
	"time"
)

var (
	QQINFOROOTPATH = "C:\\data"
	QQINFOPATH     = "C:\\data\\"
	QQINFOSKIN     = ".info"
)

type QQInfo struct {
	QQ       int64    `json:"qq"`
	PassWord [16]byte `json:"pass_word"`
	Token    []byte   `json:"token"`
	//对应的随机种子文件
	SeedFile string `json:"seed_file"`
}

func (q *QQInfo) StoreLoginInfo(qq int64, pw [16]byte, token []byte) bool {
	q.QQ = qq
	if q.QQ == 0 {
		return false
	}
	q.PassWord = pw
	q.Token = token
	fileName := QQINFOPATH + strconv.FormatInt(q.QQ, 10) + QQINFOSKIN
	marshal, err := json.Marshal(q)
	if err != nil {
		return false
	}
	return WriteFile(fileName, marshal)
}

func (q *QQInfo) Login() bool {
	var botClient = client.NewClientEmpty()
	deviceInfo := device.GetDevice(q.QQ)
	botClient.UseDevice(deviceInfo)
	err := botClient.TokenLogin(q.Token)
	fmt.Println("使用Token登录:", botClient.Uin)
	if err == nil {
		bot.Clients.Store(botClient.Uin, botClient)
		go AfterLogin(botClient)
		return true
	} else {
		fmt.Println("Token登录失败:", err, q.Token)
		time.Sleep(time.Second * 2)
		err = botClient.TokenLogin(q.Token)
		if err != nil {
			fmt.Println("Token第二次登录失败:", err, q.Token)
		}
	}
	success := CreateBotImplMd5(q.QQ, q.PassWord, q.QQ)
	if !success {
		success = CreateBotImplMd5(q.QQ, q.PassWord, q.QQ)
	}
	if success {
		go AfterLogin(botClient)
	}
	return success
}
