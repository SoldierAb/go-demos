package ipc

import (
	"encoding/json"
	"fmt"
)

type IpcClient struct {
	conn chan string
}

func NewIpcClient(ipcserver *IpcServer) *IpcClient{
	c := ipcserver.Connect()

	return &IpcClient{c}
}


func (client *IpcClient)Call(method,params string)(res *Response,err error){
	req := &Request{method,params}

	var b []byte
		b,err = json.Marshal(req)

		if err != nil{
			fmt.Println("invalid request:  ",req)
			return
		}

		client.conn <- string(b)

		str := <- client.conn    //等待连接的返回值

		var res1 Response

		err = json.Unmarshal([]byte(str),&res1)

		res = &res1

		return

}


func (client *IpcClient) Close(){
	client.conn <- "CLOSE"
}


