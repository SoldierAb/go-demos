package cg

import (
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"sync"

	"../ipc"
)

var _ ipc.Server = &CenterServer{}

type Room struct {
	ID string "id"
}

type Message struct {
	From string "from"
	To string "to"
	Content string "content"
}

type CenterServer struct {
	servers map[string] ipc.Server

	players []*Player

	rooms []*Room

	mutex sync.RWMutex

}

//创建新的服务中心
func NewCenterServer() *CenterServer{
	servers := make(map[string] ipc.Server)
	players := make([]*Player,0)

	return &CenterServer{servers:servers,players:players}
}

//添加玩家
func (server *CenterServer)AddPlayer(params string) error{
	player := NewPlayer()
	err := json.Unmarshal([]byte(params),&player)   //json解析

	if err!=nil{
		fmt.Println("添加玩家失败")
		return err
	}

	server.mutex.Lock()
	defer server.mutex.Unlock()

	server.players =append(server.players,player)
	res,_ :=json.Marshal(server.players)
	players := string(res)
	fmt.Println("添加玩家成功： ",player,players)
	return nil
}

//删除玩家
func (server *CenterServer)RemovePlayer(name string) error{
	server.mutex.Lock()
	defer server.mutex.Unlock()

	for i,serverInstance := range server.players{
		fmt.Println(i,serverInstance,name)
		if serverInstance.Name == name{
			fmt.Println(" 找到对应用户",name)
			if len(server.players) == 1{
				fmt.Println("-------0-------")
				server.players = make([]*Player,0)
			}else if i==len(server.players)-1{
				fmt.Println("-------1------")
				server.players=server.players[:1]
			}else if i==0{
				fmt.Println("-------2------")
				server.players = server.players[1:]
			}else{
				fmt.Println("-------3------")
				server.players = append(server.players[:i-1],server.players[i:]...)
			}
			return nil
		}
	}

	return errors.New("未找到对应玩家")
}

//陈列所有玩家
func (server *CenterServer)ListPlayers()(players string,err error){
	server.mutex.RLock()
	defer server.mutex.RUnlock()
	if len(server.players)>0{
		res,_ :=json.Marshal(server.players)
		players = string(res)
	}else{
		err = errors.New("暂无玩家")
	}
	return players,nil
}

//消息广播
func (server *CenterServer)Broadcast(msg string) error{
	var message Message
	err:=json.Unmarshal([]byte(msg),&message)
	if err !=nil{
		fmt.Println("消息解析出错")
		return err
	}
	if len(server.players)>0{
		for _,player :=range server.players{
			player.ms <- &message
		}
	}else{
		err = errors.New("当前无在线玩家")
	}
	return err
}

func (server *CenterServer)Handle(method,params string) *ipc.Response{
	switch method {
	case "Broadcast":
		err := server.Broadcast(params)
		if err != nil{
			return &ipc.Response{Code:err.Error()}
		}
	case "ListPlayers":
		players,err:= server.ListPlayers()
		if err != nil {
			return &ipc.Response{Code:err.Error()}
		}
		return &ipc.Response{Code:"200",Body:players}
	case "RemovePlayer":
		err := server.RemovePlayer(params)
		if err != nil{
			return &ipc.Response{Code:err.Error()}
		}
	case "AddPlayer":
		err := server.AddPlayer(params)
		if err!=nil{
			return &ipc.Response{Code:err.Error()}
		}
	}
	return &ipc.Response{Code:"200"}
}


func (server *CenterServer) Name () string{
	return "CenterServer"
}








