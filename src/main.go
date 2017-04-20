package main

import (
	"client"
	"log"
	)

const(
    id = 0
    LOCALDIR = "../files"
    MYADDR = "localhost:1234"
    REMOTEADDR = ""
)

func main() {
    ok := client.Run(LOCALDIR, MYADDR, REMOTEADDR)
    if ok != nil{
        log.Fatal(ok)
    }
}