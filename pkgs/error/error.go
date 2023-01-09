package error

type Error struct {
	Exception    error        // 非接口返回的异常
	HttpResponse HttpResponse // 接口返回的错误
}

type HttpResponse struct {
	Status     string   `json:"status"`
	StatusCode int      `json:"status_code"` // HTTP 状态码
	Response   Response // 接口返回的错误提示信息
}

type Response struct {
	ErrorResponse struct {
		CodeSpace string `json:"code_space"`
		Code      string `json:"code"`
		Message   string `json:"message"`
	} `json:"error"`
}
