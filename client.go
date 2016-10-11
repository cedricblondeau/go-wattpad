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

// NewStories returns a list of new stories
func (c *Client) NewStories() ([]Story, error) {
	req := c.buildRequest("GET", "stories", param{name: "filter", value: "new"}, param{name: "limit", value: "20"})
	var envelope StoriesEnvelope
	err := c.exec(req, &envelope)
	if err != nil {
		return nil, err
	}
	return envelope.Stories, nil
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

// Category returns the most popular category for the given tag
func (c *Client) Category(tag string) (Category, error) {
	req := c.buildRequest("GET", "categories", param{name: "tag", value: tag})
	var envelope CategoriesEnvelope
	err := c.exec(req, &envelope)
	if err != nil {
		return Category{}, err
	}
	categories := envelope.Categories
	if len(categories) < 1 {
		return Category{}, nil
	}
	return envelope.Categories[0], nil
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
	return json.Unmarshal(blob, &envelope)
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
