package httpinfo

import "net/http"

var (
	Get  = NewHttpInfo(http.MethodGet, "/api/v1/")
	Post = NewHttpInfo(http.MethodPost, "/api/v1/")
)
