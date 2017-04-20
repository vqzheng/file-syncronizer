package main

import (
    "bytes"
    "fmt"
    "io"
    "net/rpc"
    "net/http"
    )


func StartSrv(client *rpc.Client, dir string) error {
    reply bool
    ok := client.Call("Server.init", &dir, &reply)
    return ok
    }