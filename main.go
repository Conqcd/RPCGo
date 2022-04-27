package main

import (
	// 	"errors"
	// 	"flag"
	"fmt"
	"os"
	"os/exec"
	"syscall"

	// 	"log"
	// 	"time"
	// metrics "github.com/rcrowley/go-metrics"
	"github.com/smallnest/rpcx/server"
)

// var (
// 	addr     = flag.String("addr", "localhost:8972", "server address")
// 	zkAddr   = flag.String("zkAddr", "localhost:2181", "zookeeper address")
// 	basePath = flag.String("base", "/rpcx_test", "prefix path")
// )
type Arith int

func (t *Arith) femsolver() {
	binary, lookErr := exec.LookPath("mpiexec")
	if lookErr != nil {
		panic(lookErr)
	}
	args := []string{" -n 4", " ./femsolver", " ./temp/vo.txt"}
	pro, err := os.StartProcess(binary, args, &os.ProcAttr{})
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("son ", pro.Pid)
}

func main() {
	servers := server.NewServer()
	// fn := func(p *rpcx.AuthorizationAndServiceMethod) error {
	// 	if p.Authorization != "0b79bab50daca910b000d4f1a2b675d604257e42" || p.Tag != "Bearer" {
	// 		fmt.Printf("error: wrong Authorization: %s, %s\n", p.Authorization, p.Tag)
	// 		return errors.New("Authorization failed ")
	// 	}
	// 	fmt.Printf("Authorization success: %+v\n", p)
	// 	return nil
	// }
	// server.Auth(fn)
	servers.RegisterName("Arith", new(Arith), "")
	servers.Serve("tcp", "127.0.0.1:8972")
	fmt.Printf("hello, world\n")
	fmt.Println(syscall.Getpid())
	// femsolver()
}
