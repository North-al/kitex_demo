package main

import (
	"context"
	item "example_shop/kitex_gen/example/shop/item"
	"example_shop/kitex_gen/example/shop/stock"
	"example_shop/kitex_gen/example/shop/stock/stockservice"

	"github.com/cloudwego/kitex/client"
)

// ItemServiceImpl implements the last service interface defined in the IDL.
type ItemServiceImpl struct {
	stockCli stockservice.Client
}

func NewStockClient(addr string) (stockservice.Client, error) {
	return stockservice.NewClient("example.shop.stock", client.WithHostPorts(addr))
}

// GetItem implements the ItemServiceImpl interface.
func (s *ItemServiceImpl) GetItem(ctx context.Context, request *item.GetItemRequest) (resp *item.GetItemResponse, err error) {
	resp = item.NewGetItemResponse()

	// get stock
	stockRequest := &stock.GetItemStockRequest{
		ItemId: request.GetId(),
	}

	itemStock, err := s.stockCli.GetItemStock(ctx, stockRequest)
	if err != nil {
		return nil, err
	}

	resp.Item = &item.Item{
		Id:          request.GetId(),
		Title:       "kitex",
		Description: "kitex is a high-performance and scalable RPC framework for building distributed systems in Go.",
		Price:       100,
		Stock:       itemStock.GetStock(), // TODO: get stock from rpc/stock service GetItemStock(request.GetId())
	}

	return
}
