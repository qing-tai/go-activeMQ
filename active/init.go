package active

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/go-stomp/stomp"
	"github.com/spf13/viper"
)

// Init func
func Init() {
	fmt.Println("active mq start")
	// git config file
	getMsgFunc()
}

func getMsgFunc() {
	// connActiveMq()
	stompConn := connActiveMq()
	defer stompConn.Disconnect()

	fmt.Println(stompConn.Version().String())
	fmt.Println(stompConn.Server())

	topic := "/topic/" + viper.GetString("active.topic")
	fmt.Println(topic)

	stompSub, err := stompConn.Subscribe(topic,
		stomp.AckAuto,
		// stomp.SubscribeOpt.Header("persistent", "true"),
		// stomp.SubscribeOpt.Id("Skyhawk1129")
	)
	if err != nil {
		fmt.Println("stomp subscribe err:" + err.Error())
	}
	defer stompSub.Unsubscribe()

	fmt.Println("start for get msg")

	for {
		select {
		//sub.C是一个channel，如果订阅的队列有数据就读取
		case v := <-stompSub.C:
			//读取的数据是一个*stomp.Message类型
			var msg interface{}
			json.Unmarshal(v.Body, &msg)
			fmt.Println(msg)

			//如果30秒还没有人发数据的话，就结束
		case <-time.After(time.Second * 30):
			fmt.Println("timeout")

		}

	}

}
