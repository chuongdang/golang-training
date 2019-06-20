package dto

type ErrorResponse struct {
	Message string `json:"message"`
	Data    string `json:"data"`
}

func (er *ErrorResponse) ShowData() {
	println(er.Data)
}
