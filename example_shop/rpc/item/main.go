package main

import (
	item "example_shop/kitex_gen/example/shop/item/itemservice"
	"log"

	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	etcd "github.com/kitex-contrib/registry-etcd"
)

func main() {
	registry, err := etcd.NewEtcdRegistry([]string{"127.0.0.1:2379"})

	if err != nil {
		log.Fatal(err)
	}

	itemServiceImpl := new(ItemServiceImpl)
	stockCli, err := NewStockClient("0.0.0.0:8889")
	if err != nil {
		log.Fatalln(err)
	}

	itemServiceImpl.stockCli = stockCli
	svr := item.NewServer(
		itemServiceImpl,
		server.WithRegistry(registry),
		server.WithServerBasicInfo(
			&rpcinfo.EndpointBasicInfo{
				ServiceName: "example.shop.item",
			},
		),
	)

	err = svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
