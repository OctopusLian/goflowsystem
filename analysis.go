package main

import (
	"flag"
	"os"
	"time"

	"github.com/sirupsen/logrus"
)

type cmdParams struct {
	logFilePath string
	routineNum  int
}

type digData struct {
	time  string
	url   string
	refer string
	ua    string
}

type urlData struct {
	data digData
	uid  string //用户id
}

type urlNode struct {
}

type storageBlock struct {
	counterType  string
	storageModel string
	unode        urlNode
}

var log = logrus.New()

func init() {
	log.Out = os.Stdout
	log.SetLevel(logrus.DebugLevel)
}

func main() {
	//获取参数
	logFilePath := flag.String("logFilePath", "nginx/.log", "--0")
	routineNum := flag.Int("routineNum", 5, "---")
	l := flag.String("l", "tmp/log", "---")
	flag.Parse()

	params := cmdParams{
		*logFilePath,
		*routineNum,
	}

	//打日志
	logFd, err := os.OpenFile(*l, os.O_APPEND|os.O_WRONLY, 0644)
	if err == nil {
		log.Out = logFd
		defer logFd.Close()
	}
	log.Infof("Exec start.")
	log.Infof("logFilePath = %s,routneNum = %d", logFilePath, routineNum)

	//初始化一些channel，用于数据传递
	var (
		logChannel     = make(chan string, 3*params.routineNum)
		pvChannel      = make(chan urlData, params.routineNum)
		uvChannel      = make(chan urlData, params.routineNum)
		storageChannel = make(chan storageBlock, params.routineNum)
	)

	//创建日志消费者
	go readFileLinebyLine(params, logChannel)

	//创建一组日志处理
	for i := 0; i < params.routineNum; i++ {
		go logConsumer(logChannel, pvChannel, uvChannel)
	}

	//创建PV UV统计器
	go pvCounter(pvChannel, storageChannel)
	go uvCounter(uvChannel, storageChannel)
	//可以继续扩展的 xxCounter

	//创建存储器
	go dataStorage(storageChannel)

	time.Sleep(1000 * time.Second)
}

func readFileLinebyLine(params cmdParams, logChannel chan string) {

}

func logConsumer(logChannel chan string, pvChannel, uvChannel chan urlData) {

}

func pvCounter(pvChannel chan urlData, storageChannel chan storageBlock) {

}

func uvCounter(uvChannel chan urlData, storageChannel chan storageBlock) {

}

func dataStorage(storageChannel chan storageBlock) {

}
