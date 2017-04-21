//Qiwei Zheng, Chen Shou, Lu Min
//CS651 Distributed System
//final project, 04/2017

//client end
//needs the synchronized directories as the input. 
package CL

import (
   // "bytes"
  //  "fmt"
    //"io"
    "log"
    "net/rpc"
    "db"
   // "net/http"
    )


func Run(localDir string, remoteDir string, id int, myAddr string, remoteAddr string) error{
    log.Println("run")

    //establish server on local and remote machine
    localServer, ok := rpc.DialHTTP("tcp", myAddr)
    if ok != nil {
        return ok
    }
    remoteServer, ok := rpc.DialHTTP("tcp", remoteAddr)
    if ok != nil{
        return ok
    }
    log.Println("connected")

    ok = StartSrv(localServer, localDir)
    if ok != nil {
        return ok
    }
    ok = StartSrv(remoteServer, remoteDir)
    if ok != nil{
        return ok
    }

    //fetch the databases
    localDb := db.Database{}
    ok = localServer.Call("Server.Fetch", &id, &localDb)
    if ok != nil{
        return ok
    }   
    log.Println("got DB from local")
    remoteDb := db.Database{}
    ok = remoteServer.Call("Server.Fetch", &id, &remoteDb)
    if ok != nil{
        return ok
    }
    log.Println("got DB from remote")




    return nil
}


//start server
func StartSrv(client *rpc.Client, dir string) error {
    var reply string
    ok := client.Call("Server.Init", &dir, &reply)
    return ok
    }

