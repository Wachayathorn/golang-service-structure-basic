package httpinfo

type HttpInfo struct {
	HttpMethod string
	ApiPath    string
	Data       interface{}
}

func NewHttpInfo(httpMethod, apiPath string, data interface{}) HttpInfo {
	return HttpInfo{
		HttpMethod: httpMethod,
		ApiPath:    apiPath,
		Data:       data,
	}
}
