package model

type Response struct {
	Code    int         `json:"code"`
	Status  string      `json:"status"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

type PaginationResponse struct {
	Code    int         `json:"code"`
	Status  string      `json:"status"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
	Meta    MetaData    `json:"meta"`
}

type MetaData struct {
	Page      int   `json:"page"`
	PageSize  int   `json:"page_size"`
	Total     int64 `json:"total"`
	TotalPage int   `json:"total_page"`
}

func NewResponse(code int, status, message string, data interface{}) *Response {
	return &Response{
		Code:    code,
		Status:  status,
		Message: message,
		Data:    data,
	}
}

func NewPaginationResponse(code int, status, message string, data interface{}, meta MetaData) *PaginationResponse {
	return &PaginationResponse{
		Code:    code,
		Status:  status,
		Message: message,
		Data:    data,
		Meta:    meta,
	}
}
