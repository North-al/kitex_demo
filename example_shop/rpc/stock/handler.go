package main

import (
	"context"
	stock "example_shop/kitex_gen/example/shop/stock"
)

// StockServiceImpl implements the last service interface defined in the IDL.
type StockServiceImpl struct{}

// GetItemStock implements the StockServiceImpl interface.
func (s *StockServiceImpl) GetItemStock(ctx context.Context, req *stock.GetItemStockRequest) (resp *stock.GetItemStockResponse, err error) {
	// TODO: Your code here...
	resp = stock.NewGetItemStockResponse()
	resp.Stock = req.GetItemId() * 2

	return
}
