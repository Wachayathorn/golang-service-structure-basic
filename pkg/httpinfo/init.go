package httpinfo

type HttpInfo struct {
	HttpMethod string
	ApiPath    string
	Data       interface{}
}

func NewHttpInfo(httpMethod, apiPath string) HttpInfo {
	return HttpInfo{
		HttpMethod: httpMethod,
		ApiPath:    apiPath,
	}
}
