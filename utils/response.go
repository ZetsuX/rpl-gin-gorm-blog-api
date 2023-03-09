package utils

type Response struct {
	IsSuccess bool   `json:"success"`
	Message   string `json:"message"`
	Status    uint   `json:"status"`
	Data      any    `json:"data"`
}

type EmptyObj struct {
}

func CreateResponse(msg string, statusCode uint, d any) Response {
	if d == nil {
		return Response{
			IsSuccess: false, Message: msg, Status: statusCode, Data: nil,
		}
	} else {
		return Response{
			IsSuccess: true, Message: msg, Status: statusCode, Data: d,
		}
	}
}
