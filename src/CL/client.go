package CL

import (
   // "bytes"
  //  "fmt"
    //"io"
    "log"
    "net/rpc"
   // "net/http"
    )


func Run(localDir string, myAddr string, remoteAddr string) error{
    log.Println("run")

    client, ok := rpc.DialHTTP("tcp", myAddr)
    if ok != nil {
        return ok
    }
    log.Println("connected ")

    ok = StartSrv(client, localDir)
        if ok != nil {
            return ok
        }

    return nil
}



func StartSrv(client *rpc.Client, dir string) error {
    var reply bool
    ok := client.Call("Server.Init", &dir, &reply)
    return ok
    }

