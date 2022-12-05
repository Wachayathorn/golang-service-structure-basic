package client

import (
	"github.com/wachayathorn/golang-service-structure-basic/pkg/httpinfo"
	"github.com/wachayathorn/golang-service-structure-basic/pkg/model"
)

func (c *Client) Post(req model.Post) (string, error) {
	res, err := c.Utils.DoRequest(c.Url, httpinfo.Get.HttpMethod, httpinfo.Get.ApiPath, req)
	if err != nil {
		return "", err
	}
	return string(res), nil
}
