package main

/*
	http接口实现

*/

import (
	"global"
	"net/http"
	"strings"
	logger "github.com/xlog4go"
	"os"
	"io/ioutil"
	"logic"
)

func FuncHandler(w http.ResponseWriter, r *http.Request, logId int64, messageType uint64) HttpResponser {
	formData := &global.FormStruct{}
	if err := ParseForm(Input(r), formData); err != nil {
		return doErrorResponse("", global.ERR_HTTP_PARSE_FAILED, err.Error(), w)
	}
	logger.Warn("FormStruct: %v", formData)

	msg := &global.Message{
		LogId:       logId,
		Writer:      w,
		FormStruct: formData,
		MessageType: messageType,
	}
	w.Header().Set("content-type", "application/json; charset=utf-8")

	resp := logic.HandleMessage(msg)
	resp.ResponseJson(msg.Writer)
	return resp
}

var realPath string = "web/"
func StaticResource(w http.ResponseWriter, r *http.Request, logId int64, messageType uint64) HttpResponser {
	path := r.URL.Path
	logger.Error("path:%v", path)
	if path == "/" {
		path = "/index.html"
	}
	var resp *global.BaseResponse
	resp = global.NewBaseResponse()
	if !strings.Contains(path, "/") {
		w.Header().Set("content-type", "application/json; charset=utf-8")
		resp.Errmsg = "Not Found."
		resp.Errno = 404
		resp.ResponseJson(w)
		return resp
	}
	path = strings.Replace(path, "/", realPath, -1)
	var index int
	index = strings.LastIndex(path, ".")
	if index > 0 {
		request_type := path[index:]
		switch request_type {
		case ".css":
			w.Header().Set("content-type", "text/css")
		case ".js":
			w.Header().Set("content-type", "text/javascript")
		default:
		}
	}
	fin, err := os.Open(path)
	defer fin.Close()
	if err != nil {
		logger.Error("static resource:%v", err)
		w.Header().Set("content-type", "application/json; charset=utf-8")
		resp.Errmsg = "Not Found."
		resp.Errno = 404
		resp.ResponseJson(w)
		return resp
	}
	fd, _ := ioutil.ReadAll(fin)
	w.Write(fd)
	return resp
}

/*
 ************************************************************
 * PingHandler
 *************************************************************
 */
func PingHandler(w http.ResponseWriter, r *http.Request, logId int64, messageType uint64) HttpResponser {
	r.ParseForm()
	resp := &HttpResponse{
		ErrNo:  0,
		ErrMsg: "PONG",
	}
	w.Header().Set("content-type", "application/json; charset=utf-8")

	resp.ResponseJson(w)
	return resp
}