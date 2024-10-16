package domain

import "fmt"

const (
	InvalidRequestMsg   string = "Invalid Request"
	UnauthorisedRequest string = "Unauthorized Request"
)

type Code string

const (
	InvalidRequestCode Code = "BAD_REQUEST"
	ServerErrorCode    Code = "SERVER_ERROR"
)

func (code *Code) IsClientError() bool {
	validClientCodeMap := map[Code]bool{
		InvalidRequestCode: true,
		ServerErrorCode:    true,
	}
	return validClientCodeMap[*code]
}

type Entity struct {
	HttpStatusCode int    `json:"http_status_code"`
	Message        string `json:"message"`
	Code           Code   `json:"code"`
}

func (ue *Entity) Error() string {
	return fmt.Sprintf("%s: %s", ue.Code, ue.Message)
}
