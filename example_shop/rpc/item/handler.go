package main

import (
	"context"
	item "example_shop/kitex_gen/example/shop/item"
)

// ItemServiceImpl implements the last service interface defined in the IDL.
type ItemServiceImpl struct{}

// GetItem implements the ItemServiceImpl interface.
func (s *ItemServiceImpl) GetItem(ctx context.Context, request *item.GetItemRequest) (resp *item.GetItemResponse, err error) {
	resp = item.NewGetItemResponse()
	resp.Item = &item.Item{
		Id:          request.GetId(),
		Title:       "kitex",
		Description: "kitex is a high-performance and scalable RPC framework for building distributed systems in Go.",
		Price:       100,
		Stock:       1, // TODO: get stock from rpc/stock service GetItemStock(request.GetId())
	}
	
	return
}
