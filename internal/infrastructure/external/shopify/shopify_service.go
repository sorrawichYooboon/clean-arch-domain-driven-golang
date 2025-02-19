package shopify

type ShopifyServiceImpl struct {
	baseURL string
	apiKey  string
}

func NewShopifyService(baseURL string, apiKey string) ShopifyService {
	return &ShopifyServiceImpl{
		baseURL: baseURL,
		apiKey:  apiKey,
	}
}

func (s *ShopifyServiceImpl) GetProductByID(request *GetProductByIDRequest) (*GetProductByIDResponse, error) {
	// Call Shopify API to get product by ID
	return &GetProductByIDResponse{}, nil
}
