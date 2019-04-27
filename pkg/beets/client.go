package beets

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

type HTTPGetter interface {
	Get(string) (*http.Response, error)
}

type Client struct {
	httpClient HTTPGetter
	BaseURL    string
}

func NewClient(baseUrl string) *Client {
	return &Client{
		httpClient: http.DefaultClient,
		BaseURL: strings.Trim(baseUrl, "/"),
	}
}

func (c *Client) GetResourceURL(resource string) string {
	return c.BaseURL + resource
}

func (c *Client) GetAlbums() ([]Album, error) {
	url := c.GetResourceURL("/album/")
	res, err := c.httpClient.Get(url)
	if err != nil {
		return nil, err
	}

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("GET %s returned code %d", url, res.StatusCode)
	}

	type albumsResponse struct {
		Albums []Album `json:"albums"`
	}
	decoder := json.NewDecoder(res.Body)
	out := albumsResponse{}
	if err := decoder.Decode(&out); err != nil {
		return nil, err
	}
	return out.Albums, nil
}
