package iserver

//success
const (
	ApiCodeSuccess = 0
	ApiCodeFail = 1

	//登录相关100-200
	ApiCodeLoginInvalid = 101	//登录失效

	//非法操作200
	ApiCodeIllegalConnentServer =200 //非法连接服务器
	
)

type IApiReply struct{
	Msg string
	Code int
}
