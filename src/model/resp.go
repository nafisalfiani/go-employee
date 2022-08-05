package model

type HTTPResp struct {
	Message    HTTPMessage `json:"message"`
	Data       interface{} `json:"data,omitempty"`
	Error      string      `json:"error,omitempty"`
	Pagination *Pagination `json:"pagination,omitempty"`
}

type HTTPMessage struct {
	Title string `json:"title"`
	Body  string `json:"body"`
}

type Pagination struct {
	TotalElement int64 `json:"total_element"`
	CurentPage   int64 `json:"curent_page"`
	TotalPages   int64 `json:"total_pages"`
}
