namespace go example.shop.item

include "base.thrift"

// 商品信息
struct Item {
    1: i64 id
    2: string title
    3: string description
    4: i64 price
    5: i64 stock
}


struct GetItemRequest {
    1: required i64 id
}

struct GetItemResponse {
    1: Item item

    255: base.BaseResponse baseResponse
}


service ItemService {
    GetItemResponse GetItem(1: GetItemRequest request)
}