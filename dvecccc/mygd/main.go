package main

import (
	"fmt"
	"github.com/dvecccc/mygd/config"
	"log"
	"os"
)

//git: ssh-agent bash --> ssh-add ~/.ssh/2312

func init() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	f, err := os.OpenFile("./log/rabbitmq.log", os.O_RDONLY|os.O_APPEND|os.O_CREATE, 0666)
	if err != nil {
		log.Fatal("打开日志文件失败")
	}
	log.SetOutput(f)
}

//func main() {
//	rabbitmq.Publish()
//}

//Go 运行的相对路径是相对于执行命令时的目录

func main() {
	data := config.GetConfigData()
	fmt.Println(data)
}
