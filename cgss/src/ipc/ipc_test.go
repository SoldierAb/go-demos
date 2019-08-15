package ipc

import "testing"

type EchoServer struct {

}


func (server *EchoServer) Handle(request string) string{
	return "ECHO: "+request
}

func (server *EchoServer) Name() string{
	return "ECHO SERVER"
}

func TestIpc (t *testing.T){
	server :=&IpcServer{&EchoServer{}}
	client1 := NewIpcClient(server)
	client2 := NewIpcClient(server)

	res1 := server.call("client 1")
	res2 := server.call("client 2")

	if res1 != "ECHO: client 1" || res2 != "ECHO: client 2"{
		t.Error("ipc client Call failed",res1," ",res2)
	}

	client1.Close()
	client2.Close()

}
