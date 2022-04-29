package main

import (
	// 	"errors"
	// 	"flag"
	"context"
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
type Args struct {
	A, B int
}

type Quotient struct {
	Quo, Rem int
}

type Arith int

func (t *Arith) Femsolver(cxt context.Context, args *Args, reply *int) error {
	*reply = args.A * args.B
	binary, lookErr := exec.LookPath("mpiexec")
	if lookErr != nil {
		panic(lookErr)
	}
	args2 := []string{" -n 4", " ./femsolver", " ./temp/vo.txt"}
	pro, err := os.StartProcess(binary, args2, &os.ProcAttr{})
	if err != nil {
		fmt.Println(err)
		return nil
	}
	fmt.Println("son ", pro.Pid)
	return nil
}

// func (t *Arith) femsolver(cxt context.Context, args1 *Args, reply *int) error {
// 	binary, lookErr := exec.LookPath("mpiexec")
// 	if lookErr != nil {
// 		panic(lookErr)
// 	}
// 	args := []string{" -n 4", " ./femsolver", " ./temp/vo.txt"}
// 	pro, err := os.StartProcess(binary, args, &os.ProcAttr{})
// 	if err != nil {
// 		fmt.Println(err)
// 		return nil
// 	}

// 	fmt.Println("son ", pro.Pid)
// 	return nil
// }

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
