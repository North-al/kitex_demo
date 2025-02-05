package main

import (
	"log"

	api "github.com/North-al/kitex_demo/hello/kitex_gen/api/hello"
)

func main() {
	svr := api.NewServer(new(HelloImpl))
	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
