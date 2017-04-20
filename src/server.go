package main

import (
   // "bytes"
  //  "fmt"
    //"io"
    "log"
    "net/rpc"
    "net"
    "net/http"
    "os"
  // "errors"
    )

type Server int

var (
	Database int
	initialized bool
)

func (sr *Server) Init(dir *string, reply *int) error {
	err := os.Chdir(*dir)
	if err != nil {
		return err
	}

	// localDb, err = db.ParseFile()
	// if err == db.FileNotFound {
	// 	initialized = true
	// 	return nil
	// } else if err != nil {
	// 	return err
	// }

	// localDb.VersionVec[localDb.ReplicaId] += 1

	// log.Println(localDb)
	// if err := localDb.Update(); err != nil {
	// 	return err
	// }

	initialized = true
	return nil
}


func main() {
    server := new(Server)
	rpc.Register(server)
	rpc.HandleHTTP()
	l, ok := net.Listen("tcp", ":1234")
	if ok != nil {
   		log.Fatal("listen error:", ok)
	}
	http.Serve(l, nil)
}

