package main

import (
	"net"
	"net/http"
	"global"
	"time"
	"sync/atomic"
)

var logFile = "./conf/log.json"
var confFile = "./conf/main.json"

var reqchan = make(chan string, 1000000)

var logidGenerator LogId

var SimpleServerListener net.Listener
var httpServer http.Server

var uri2Handler map[string]*SimpleServerHandler


var SimpleServerQuit chan int
func init() {
	logidGenerator = LogId(time.Now().Unix())

	//安全退出
	SimpleServerQuit = make(chan int)

	uri2Handler = make(map[string]*SimpleServerHandler)

	uri2Handler["/ping"] = &SimpleServerHandler{Name: "Ping", Callfunc: PingHandler}
	uri2Handler["/login"] = &SimpleServerHandler{Name: "Login", MessageType: global.KMsgTypeLogin, Callfunc: FuncHandler}
	uri2Handler["/logout"] = &SimpleServerHandler{Name: "Logout", MessageType: global.KMsgTypeLogout, Callfunc: FuncHandler}
	uri2Handler["/"] = &SimpleServerHandler{Name: "GetSimpleServerInfo", Callfunc: StaticResource}
}

type LogId int64

func (i *LogId) GetNextId() int64 {
	return atomic.AddInt64((*int64)(i), 1)
}
