package main

import (
	"bufio"
	"flag"
	"ip-pool/pkg/global"
	"ip-pool/pkg/httpserver"
	"ip-pool/pkg/v2ray"
	"log"
	"os"
)

var (
	tag        string
	serverAddr string
	db         string
	logToFile  bool
)

func init() {
	flag.StringVar(&tag, "tag", "puppet", "傀儡节点的puppet")
	flag.StringVar(&serverAddr, "server", "195.245.242.102:10081", "服务器的端口")
	flag.StringVar(&db, "db", "./user.db", "数据库文件")
	flag.BoolVar(&logToFile, "l", false, "日志是否进入文件")
}

func main() {
	flag.Parse()
	log.SetFlags(log.Lshortfile | log.Ltime)

	if logToFile {
		fi, _ := os.OpenFile("log.log", os.O_WRONLY|os.O_APPEND, 0777)
		defer fi.Close()

		buf := bufio.NewWriter(fi)
		log.SetOutput(buf)
	}

	global.UserTag = tag
	v2ray.InitClient(serverAddr) // "195.245.242.102:10081"
	v2ray.InitDB(db)

	if err := v2ray.InitUsers(); err != nil {
		log.Println(err)
	}

	httpserver.StartServer()
}
