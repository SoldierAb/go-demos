package cg

import (
	"../ipc"
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
)

type CenterClient struct {
	*ipc.IpcClient				//匿名组合ipc.IpcClient
}

func (client *CenterClient) AddPlayer (player *Player) error{
	b,err := json.Marshal(*player)
	if err !=nil{
		return err
	}

	res,err := client.Call("AddPlayer",string(b))

	if err ==nil && res.Code == "200"{
		return nil
	}

	return errors.New(res.Code)
}


func (client *CenterClient)RemovePlayer (name string) error{
	res,err := client.Call("RemovePlayer",name)

	if err !=nil&&res.Code!="200"{
		fmt.Println("删除用户失败  ",name)
		return errors.New(res.Code)
	}

	return nil
}

func (client *CenterClient)ListPlayers(params string)(ps []*Player, err error) {
	res, err:= client.Call("ListPlayers", params)
	if res.Code != "200" {
		err = errors.New(res.Code)
		return
	}
	err = json.Unmarshal([]byte(res.Body), &ps)
	return
}


func (client *CenterClient)Broadcast (msg string) error{
	m := &Message{Content:msg}

	b,err := json.Marshal(m)

	if err != nil{
		fmt.Println("广播消息json解析出错")
		return err
	}

	res,err := client.Call("Broadcast",string(b))

	if err !=nil && res.Code !="200"{
		return errors.New(res.Code)
	}

	return nil
}




