package dto

type Response struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
	Error   *string     `json:"error,omitempty"`
}

type PaginationMeta struct {
	Page       int   `json:"page"`
	PerPage    int   `json:"per_page"`
	Total      int64 `json:"total"`
	TotalPages int   `json:"total_pages"`
}

type PaginatedResponse struct {
	Success bool           `json:"success"`
	Message string         `json:"message"`
	Data    interface{}    `json:"data"`
	Meta    PaginationMeta `json:"meta"`
}

func SuccessResponse(message string, data interface{}) Response {
	return Response{
		Success: true,
		Message: message,
		Data:    data,
	}
}

func ErrorResponse(message string, err error) Response {
	errMsg := err.Error()
	return Response{
		Success: false,
		Message: message,
		Error:   &errMsg,
	}
}

func PaginatedSuccessResponse(message string, data interface{}, meta PaginationMeta) PaginatedResponse {
	return PaginatedResponse{
		Success: true,
		Message: message,
		Data:    data,
		Meta:    meta,
	}
}
