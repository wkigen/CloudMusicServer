package iserver

//success
const (
	ApiCodeSuccess = 0
	ApiCodeFail = 1
)

type IApiReply struct{
	Msg string
	Code int
}
