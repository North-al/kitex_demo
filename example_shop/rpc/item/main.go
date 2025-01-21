package main

import (
	item "example_shop/kitex_gen/example/shop/item/itemservice"
	"log"
)

func main() {
	itemServiceImpl := new(ItemServiceImpl)
	stockCli, err := NewStockClient("0.0.0.0:8889")
	if err != nil {
		log.Fatalln(err)
	}

	itemServiceImpl.stockCli = stockCli
	svr := item.NewServer(itemServiceImpl)

	err = svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
