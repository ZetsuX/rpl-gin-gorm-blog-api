package utils

type Response struct {
	IsSuccess bool   `json:"success"`
	Message   string `json:"message"`
	Status    uint   `json:"status"`
	Data      any    `json:"data"`
}

type EmptyObj struct {
}

func CreateFailResponse(msg string, statusCode uint) Response {
	return Response{
		IsSuccess: false, Message: msg, Status: statusCode, Data: nil,
	}
}

func CreateSuccessResponse(msg string, statusCode uint, d any) Response {
	return Response{
		IsSuccess: true, Message: msg, Status: statusCode, Data: d,
	}
}
