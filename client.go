package wattpad

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

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

// Categories returns a list of categories
func (c *Client) Categories() ([]Category, error) {
	req := c.buildRequest("GET", "categories")
	var envelope CategoriesEnvelope
	err := c.exec(req, &envelope)
	if err != nil {
		return nil, err
	}
	return envelope.Categories, nil
}

func (c *Client) exec(req *http.Request, envelope interface{}) error {
	resp, err := c.Client.Do(req)
	if err != nil {
		return err
	}
	blob, err := read(resp)
	if err != nil {
		return err
	}
	return unmarshal(blob, &envelope)
}

func unmarshal(blob []byte, envelope interface{}) error {
	err := json.Unmarshal(blob, &envelope)
	if err != nil {
		return err
	}
	return nil
}

func read(resp *http.Response) ([]byte, error) {
	defer resp.Body.Close()
	blob, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return blob, nil
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
