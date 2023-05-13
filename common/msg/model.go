package msg

type Response struct {
	code int
	data interface{}
	msg  string
}

func (res *Response) ReturnOK() *Response {
	res.code = 200
	return res
}

func (res *Response) ReturnError(code int) *Response {
	res.code = code
	return res
}
