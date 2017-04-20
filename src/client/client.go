package client

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

    ok = StartSrv(client, localDir)
        if ok != nil {
            return ok
        }

    return nil
}



func StartSrv(client *rpc.Client, dir string) error {
    var reply bool
    ok := client.Call("Server.init", &dir, &reply)
    return ok
    }

