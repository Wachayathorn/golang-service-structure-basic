package utils

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"
)

func (u *Utils) DoRequest(url, httpMethod, path string, data interface{}) ([]byte, error) {
	apiUrl := url + path
	u.Logger.Infof("DoRequest : do request to %s", apiUrl)

	body, err := json.Marshal(data)
	if err != nil {
		u.Logger.Errorf("DoRequest : json marshall for api:%s got error:%s", path, err.Error())
		return nil, err
	}

	req, err := http.NewRequest(httpMethod, apiUrl, strings.NewReader(string(body)))
	if err != nil {
		u.Logger.Errorf("DoRequest : build request to %s got error:%s", apiUrl, err.Error())
		return nil, err
	}

	u.addRequestHeader(req)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		u.Logger.Errorf("DoRequest : do request to %s got error:%s", apiUrl, err.Error())
		return nil, err
	}

	defer res.Body.Close()
	resBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		u.Logger.Errorf("DoRequest : read response body got error:%s", err.Error())
		return nil, err
	}

	return resBody, nil
}

func (u *Utils) addRequestHeader(req *http.Request) *http.Request {
	req.Header = http.Header{
		"Content-Type": {"application/json"},
	}
	return req
}
