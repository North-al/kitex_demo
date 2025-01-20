package main

import (
	"context"
	"example_shop/kitex_gen/example/shop/item"
	"example_shop/kitex_gen/example/shop/item/itemservice"
	"fmt"
	"log"

	"github.com/cloudwego/kitex/client"
)

func main() {
	newClient, err := itemservice.NewClient("example.shop.item", client.WithHostPorts("0.0.0.0:8888"))
	if err != nil {
		log.Fatalln("create client failed", err)
	}

	req := &item.GetItemRequest{Id: 1024}
	getItem, err := newClient.GetItem(context.Background(), req)
	if err != nil {
		log.Fatalln("get item failed", err)
	}

	fmt.Println("get item success", getItem.Item)

}
