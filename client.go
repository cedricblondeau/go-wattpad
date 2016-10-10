package wattpad

import "net/http"

const (
	apiURL        = "https://api.wattpad.com/v4"
	clientVersion = "0.0.1"
	userAgent     = "go-wattpad " + clientVersion
)

// NewClient returns a new HTTP client
func NewClient(Key string) Client {
	c := Client{}
	c.Key = Key
	return c
}

// Client struct defines a HTTP client
type Client struct {
	Client http.Client
	Key    string
}

type param struct {
	name  string
	value string
}

func (c *Client) buildRequest(method string, endpoint string, params ...param) *http.Request {
	req, _ := http.NewRequest(method, apiURL+"/"+endpoint, nil)
	req.Header.Add("User-agent", userAgent)
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Authorization", "Basic "+c.Key)
	q := req.URL.Query()
	for _, param := range params {
		q.Add(param.name, param.value)
	}
	req.URL.RawQuery = q.Encode()
	return req
}
