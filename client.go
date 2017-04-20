package main

import (
   // "bytes"
  //  "fmt"
    //"io"
    "log"
    "net/rpc"
   // "net/http"
    )

const(
    id = 0
    LOCALDIR = "../files"
    MYADDR = ""
    REMOTEADDR = ""
)

func Run(localDir string, myAddr string, remoteAddr string) error{
    log.Println("run")
    return nil
}


func StartSrv(client *rpc.Client, dir string) error {
    var reply bool
    ok := client.Call("Server.init", &dir, &reply)
    return ok
    }


func main() {
    ok :=Run(LOCALDIR, MYADDR, REMOTEADDR)
    if ok != nil{
        log.Fatal(ok)
    }
    // if err := Run(LOCALDIR, MYADDR, REMOTEADDR); err != nil {
    //     log.Fatal(err)
    // }
}