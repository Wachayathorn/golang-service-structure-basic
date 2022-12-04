package client

import "github.com/wachayathorn/golang-service-structure-basic/pkg/httpinfo"

func (c *Client) Get() (string, error) {
	res, err := c.Utils.DoRequest(c.Url, httpinfo.Get.HttpMethod, httpinfo.Get.ApiPath, nil)
	if err != nil {
		return "", err
	}
	return string(res), nil
}
