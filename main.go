package main

import (
// 	"errors"
// 	"flag"
	"fmt"
// 	"log"
// 	"time"

// metrics "github.com/rcrowley/go-metrics"
// "github.com/smallnest/rpcx"
// "github.com/smallnest/rpcx/server"
// "github.com/smallnest/rpcx/serverplugin"
// )

// var (
// 	addr     = flag.String("addr", "localhost:8972", "server address")
// 	zkAddr   = flag.String("zkAddr", "localhost:2181", "zookeeper address")
// 	basePath = flag.String("base", "/rpcx_test", "prefix path")
)

func main() {
	// server := rpcx.NewServer()
	// fn := func(p *rpcx.AuthorizationAndServiceMethod) error {
	// 	if p.Authorization != "0b79bab50daca910b000d4f1a2b675d604257e42" || p.Tag != "Bearer" {
	// 		fmt.Printf("error: wrong Authorization: %s, %s\n", p.Authorization, p.Tag)
	// 		return errors.New("Authorization failed ")
	// 	}
	// 	fmt.Printf("Authorization success: %+v\n", p)
	// 	return nil
	// }
	// server.Auth(fn)
	// server.RegisterName("Arith", new(Arith))
	// server.Serve("tcp", "127.0.0.1:8972")
	fmt.Printf("hello, world\n")
}
