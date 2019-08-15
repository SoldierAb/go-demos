package cg

import "fmt"

type Player struct {
	Name string "name"
	Level int "level"
	Exp int "exp"
	Room int "room"

	ms chan *Message //等待接收的消息
}


func NewPlayer() *Player{
	m := make(chan *Message,1024)
	player := &Player{"",1,1,1,m}

	go func(p *Player) {
		for{
			msg :=<- p.ms
			fmt.Println(p.Name," 接收到消息~: ",msg.Content)
		}
	}(player)

	return player
}


//为每个玩家创建独立的goroutine ,监听所有发送给他们的聊天消息，一旦收到就及时打印到控制台上