package shopify

type ShopifyError struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

func (e *ShopifyError) Error() string {
	return e.Message
}

var (
	ErrShopifyUserNotFound = ShopifyError{Code: "ERROR_Shopify_USER_NOT_FOUND", Message: "User not found"}
	ErrShopifyServiceFail  = ShopifyError{Code: "ERROR_Shopify_SERVICE_FAIL", Message: "Failed to get Shopify service"}
)
