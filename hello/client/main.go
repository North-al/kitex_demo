package main

import (
	"context"
	"log"
	"time"

	"github.com/North-al/kitex_demo/hello/kitex_gen/api"
	"github.com/North-al/kitex_demo/hello/kitex_gen/api/hello"
	"github.com/cloudwego/kitex/client"
)

func main() {
	c, err := hello.NewClient("hello", client.WithHostPorts("0.0.0.0:8888"))

	if err != nil {
		log.Fatal("new client failed", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// req := &api.Request{Message: "hello"}
	// echo, err := c.Echo(context.Background(), req)
	// if err != nil {
	// 	log.Fatal("echo failed", err)
	// }
	//
	// log.Println("echo response:", echo.Message)
	// time.Sleep(time.Second)

	add, err := c.Add(ctx, &api.AddRequest{
		First:  2,
		Second: 3,
	})

	if err != nil {
		log.Fatal("add failed", err)
	}

	log.Println("add response:", add.Sum)
	time.Sleep(time.Second)

}
