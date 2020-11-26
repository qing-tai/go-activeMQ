package active

import (
	"fmt"
	"os"
	"time"

	"github.com/go-stomp/stomp"
	"github.com/spf13/viper"
)

// 使用IP和端口连接到ActiveMQ服务器
// 返回ActiveMQ连接对象
func connActiveMq() (stompConn *stomp.Conn) { // @todo 实现断开重连
	// stompConn, err := stomp.Dial("tcp", net.JoinHostPort(host, port))

	broker := viper.GetString("active.broker")
	name := viper.GetString("active.name")
	password := viper.GetString("active.password")
	// topic := viper.GetString("active.topic")
	// broker := viper.GetString("active.broker")

	stompConn, err := stomp.Dial("tcp",
		broker,
		stomp.ConnOpt.Login(name, password),
		stomp.ConnOpt.HeartBeat(3*time.Minute, 3*time.Minute),
		stomp.ConnOpt.HeartBeatError(3*time.Second))
	if err != nil {
		fmt.Println("connect to active_mq server service, error: " + err.Error())
		os.Exit(1)
	}

	return stompConn
}

// 将消息发送到ActiveMQ中
func activeMqProducer(c chan string, queue string, conn *stomp.Conn) {
	for {
		err := conn.Send(queue, "text/plain", []byte(<-c))
		fmt.Println("send active mq..." + queue)
		if err != nil {
			fmt.Println("active mq message send erorr: " + err.Error())
		}
	}
}
