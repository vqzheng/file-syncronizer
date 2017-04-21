//Qiwei Zheng, Chen Shou, Lu Min
//CS651 Distributed System
//final project, 04/2017

//main excuting package
package main

import (
	"CL"
	"log"
	"os"
	)

const(
    id = 0
    MYADDR = "localhost:1234"
    REMOTEADDR = "localhost:1235"
)

//start the client
func main() {
    ok := CL.Run(os.Args[1], os.Args[2], id, MYADDR, REMOTEADDR)
    if ok != nil{
        log.Fatal(ok)
    }
}