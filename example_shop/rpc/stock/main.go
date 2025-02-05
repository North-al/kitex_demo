package main

import (
	stock "example_shop/kitex_gen/example/shop/stock/stockservice"
	"log"
	"net"

	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	etcd "github.com/kitex-contrib/registry-etcd"
)

func main() {
	registry, err := etcd.NewEtcdRegistry([]string{"127.0.0.1:2379"})
	if err != nil {
		log.Fatal(err)
	}

	addr, _ := net.ResolveTCPAddr("tcp", "127.0.0.1:8889")
	svr := stock.NewServer(
		new(StockServiceImpl),
		server.WithServiceAddr(addr),
		server.WithRegistry(registry), // add registry
		server.WithServerBasicInfo(
			&rpcinfo.EndpointBasicInfo{
				ServiceName: "example.shop.stock",
			}),
	)

	err = svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
