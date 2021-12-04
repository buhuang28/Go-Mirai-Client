package handler

import (
	"encoding/json"
	"fmt"
	"strings"
)

func AutoLogin() {
	dir := ReadDir(QQINFOROOTPATH)
	for _, v := range dir {
		if dir == nil {
			return
		}
		if strings.Contains(v.Name(), QQINFOSKIN) {
			var qqInfo QQInfo
			fileByte := ReadFileByte(QQINFOPATH + v.Name())
			err := json.Unmarshal(fileByte, &qqInfo)
			if err != nil {
				fmt.Println("反序列化失败:", err)
				continue
			}
			qqInfo.Login()
		}
	}
}
