package app

type Response struct {
	// code
	Code int `json:"code" example:"200"`
	// data set
	Data interface{} `json:"data"`
	//
	Msg string `json:"msg"`
}

type Page struct {
	List      interface{} `json:"list"`
	Count     int         `json:"count"`
	PageIndex int         `json:"pageIndex"`
	PageSize  int         `json:"pageSize"`
}

type PageResponse struct {
	// code
	Code int `json:"code" example:"200"`
	// data set
	Data Page `json:"data"`
	// n
	Msg string `json:"msg"`
}

func (res *Response) ReturnOK() *Response {
	res.Code = 200
	return res
}

func (res *Response) ReturnError(code int) *Response {
	res.Code = code
	return res
}

func (res *PageResponse) ReturnOK() *PageResponse {
	res.Code = 200
	return res
}
