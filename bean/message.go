package bean

//1请求登陆，2请求发送消息，3请求登出
type Message struct {
	MessId    int
	MessValue []byte
}
