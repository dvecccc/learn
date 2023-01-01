package main

import (
	"github.com/dvecccc/demo/math"
	"log"
	"os"
)

func init() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	f, err := os.OpenFile("./rabbitmq.log", os.O_RDONLY|os.O_APPEND|os.O_CREATE, 0666)
	if err != nil {
		log.Fatal("打开日志文件失败")
	}
	log.SetOutput(f)
}

func main() {
	log.Println("hello world")
	math.Add(1, 2)
}
