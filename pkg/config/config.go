package config

import (
	"encoding/json"

	"github.com/pkg/errors"
)

var (
	Fragment = false // 是否分片
	Port     = "9000"
	SMS      = true
	Device   = ""
	Conf     = &GmcConfig{
		//SMS:  false,
		//Port: "9000",
		ServerGroups: []*ServerGroup{
			{
				Name:         "default",
				Disabled:     false,
				Json:         false,
				Urls:         []string{"ws://localhost:8081/ws/cq/"},
				EventFilter:  []int32{},
				RegexFilter:  "",
				RegexReplace: "",
				ExtraHeader: map[string][]string{
					"User-Agent": {"GMC"},
				},
			},
		},
	}
	HttpAuth = map[string]string{}
)

type GmcConfig struct {
	//Port         string         `json:"port"`          // 管理端口
	//SMS          bool           `json:"sms"`           // 设备锁是否优先使用短信认证
	ServerGroups []*ServerGroup `json:"server_groups"` // 服务器组
}

type ServerGroup struct {
	Name         string              `json:"name"`          // 功能名称
	Disabled     bool                `json:"disabled"`      // 不填false默认启用
	Json         bool                `json:"json"`          // json上报
	Urls         []string            `json:"urls"`          // 服务器列表
	EventFilter  []int32             `json:"event_filter"`  // 事件过滤
	RegexFilter  string              `json:"regex_filter"`  // 正则过滤
	RegexReplace string              `json:"regex_replace"` // 正则替换
	ExtraHeader  map[string][]string `json:"extra_header"`  // 自定义请求头
	// TODO event filter, msg filter, regex filter, prefix filter, suffix filter
}

func (g *GmcConfig) ReadJson(d []byte) error {
	var fileConfig GmcConfig
	if err := json.Unmarshal(d, &fileConfig); err != nil {
		return errors.Wrap(err, "failed to unmarshal json GmcConfig")
	}
	*g = fileConfig
	return nil
}

func (g *GmcConfig) ToJson() []byte {
	b, _ := json.MarshalIndent(g, "", "    ")
	return b
}
