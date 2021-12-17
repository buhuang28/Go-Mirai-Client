package ws_data

import (
	"bytes"
	"fmt"
	log "github.com/sirupsen/logrus"
	"runtime"
	"sync"
)

var (
	ChanMap     = make(map[int64]chan GMCWSData)
	ChanMapLock sync.Mutex
)

//处理入群请求、机器人被邀请进群的回调结果
func HandleCallBackEvent(data GMCWSData) {
	ChanMapLock.Lock()
	defer func() {
		e := recover()
		if e != nil {
			fmt.Println(e)
		}
		ChanMapLock.Unlock()
	}()
	ch := ChanMap[data.RequestId]
	ch <- data
}

func PrintStackTrace(err interface{}) {
	buf := new(bytes.Buffer)
	fmt.Fprintf(buf, "%v\n", err)
	for i := 1; ; i++ {
		pc, file, line, ok := runtime.Caller(i)
		if !ok {
			break
		}
		fmt.Fprintf(buf, "%s:%d (0x%x)\n", file, line, pc)
	}
	log.Warnf(buf.String())
}
