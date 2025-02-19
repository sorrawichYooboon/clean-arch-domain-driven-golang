package shopify

type ShopifyService interface {
	GetProductByID(request *GetProductByIDRequest) (*GetProductByIDResponse, error)
}
