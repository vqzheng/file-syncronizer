//Qiwei Zheng, Chen Shou, Lu Min
//CS651 Distributed System
//final project, 04/2017


//server.go, run this before sync
//needs the port number as the arg
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
    "errors"
    "db"
    )

type Server int

var (
	Database *db.Database
	initialized bool
	directory = "../files"
)

func (sr *Server) Init(dir *string, reply *string) error {
	directory = *dir
	err := os.Chdir(*dir)
	if err != nil {
		log.Println("dir err")
		return err
	}
	//reply := true

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

func(sr *Server) Fetch(id *int, reply *db.Database) error {
	if !initialized {
		return errors.New("you need to initiate the server first")
	}else{
		reply = Database
		return nil
	}
}

// 
func main() {
    server := new(Server)
	rpc.Register(server)
	rpc.HandleHTTP()
	l, ok := net.Listen("tcp", ":"+os.Args[1])
	if ok != nil {
   		log.Fatal("listen error:", ok)
	}
	http.Serve(l, nil)
}

