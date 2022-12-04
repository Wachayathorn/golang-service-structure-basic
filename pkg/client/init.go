package client

type ClientAPI interface {
}

type Client struct {
	Url string
}

func Init() *Client {
	return &Client{
		Url: "http://localhost:8080",
	}
}

func (c *Client) ReplaceUrl(url string) {
	c.Url = url
}
