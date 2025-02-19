package response

type Response struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

var (
	SuccessGetAllBook = Response{Code: "SUCCESS_GET_ALL_BOOK", Message: "Books retrieved successfully"}
	SuccessCreateBook = Response{Code: "SUCCESS_CREATE_BOOK", Message: "Book created successfully"}
	SuccessUpdateBook = Response{Code: "SUCCESS_UPDATE_BOOK", Message: "Book updated successfully"}
	SuccessDeleteBook = Response{Code: "SUCCESS_DELETE_BOOK", Message: "Book deleted successfully"}

	SuccessRegisterUser = Response{Code: "SUCCESS_REGISTER_USER", Message: "User registered successfully"}
	SuccessLoginUser    = Response{Code: "SUCCESS_LOGIN_USER", Message: "User logged in successfully"}

	SuccessGetAllAuthor = Response{Code: "SUCCESS_GET_ALL_AUTHOR", Message: "Authors retrieved successfully"}
	SuccessCreateAuthor = Response{Code: "SUCCESS_CREATE_AUTHOR", Message: "Author created successfully"}
	SuccessUpdateAuthor = Response{Code: "SUCCESS_UPDATE_AUTHOR", Message: "Author updated successfully"}
	SuccessDeleteAuthor = Response{Code: "SUCCESS_DELETE_AUTHOR", Message: "Author deleted successfully"}
)
