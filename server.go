package main

import(
	"net/http"
	"github.com/gorilla/rpc"
	"github.com/gorilla/rpc/json"
	"fmt"
)

func init() string{
	 s := rpc.NewServer()
    s.RegisterCodec(json.NewCodec(), "application/json")
    s.RegisterService(new(HelloService), "")
    http.Handle("/rpc", s)
    return "ok"
}

type HelloArgs struct {
    Who string
}

type HelloReply struct {
    Message string
}

type HelloService struct {}

func (h *HelloService) Say(r *http.Request, args *HelloArgs, reply *HelloReply) error {
    reply.Message = "Hello, " + args.Who + "!"
    return nil
}

func main(){
	fmt.Println("Hello")
	ok :=init()
}
