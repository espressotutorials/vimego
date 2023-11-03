package vimego

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"
	"time"
)

const apiUrl = "https://api.vimeo.com"

type Error struct {
	StatusCode int
	URL        string
	Body       string
}

func (e Error) Error() string {
	return fmt.Sprintf("%d %s: %s", e.StatusCode, e.URL, e.Body)
}

type VimeoResponse[T any] struct {
	Total   int         `json:"total,omitempty"`
	Page    int         `json:"page,omitempty"`
	PerPage int         `json:"per_page,omitempty"`
	Paging  VimeoPaging `json:"paging,omitempty"`
	Data    []T         `json:"data,omitempty"`
}

type VimeoPaging struct {
	Next     string `json:"next,omitempty"`
	Previous string `json:"previous,omitempty"`
	First    string `json:"first,omitempty"`
	Last     string `json:"last,omitempty"`
}

type VimeoPictures struct {
	Uri      string      `json:"uri,omitempty"`
	Active   bool        `json:"active,omitempty"`
	Type     string      `json:"type,omitempty"`
	BaseLink string      `json:"base_link,omitempty"`
	Sizes    []VimeoSize `json:"sizes,omitempty"`
}

type VimeoFile struct {
	Quality     string    `json:"quality,omitempty"`
	Rendition   string    `json:"rendition,omitempty"`
	Type        string    `json:"type,omitempty"`
	Width       int       `json:"width,omitempty"`
	Height      int       `json:"height,omitempty"`
	Link        string    `json:"link,omitempty"`
	CreatedTime time.Time `json:"created_time,omitempty"`
	FPS         int       `json:"fps,omitempty"`
	Size        int       `json:"size,omitempty"`
	MD5         string    `json:"md5,omitempty"`
	PublicName  string    `json:"public_name,omitempty"`
	SizeShort   string    `json:"size_short"`
}

type VimeoSize struct {
	Width  int    `json:"width,omitempty"`
	Height int    `json:"height,omitempty"`
	Link   string `json:"link,omitempty"`
}

type VimeoPrivacy struct {
	View string `json:"view,omitempty"`
}

type VimeoOptions []string

type Client struct {
	accessToken string
	httpClient  *http.Client
}

func New(accessToken string) *Client {
	return &Client{
		accessToken: accessToken,
		httpClient:  &http.Client{},
	}
}

func (c *Client) get(uri string, jsonResponse interface{}, params ...QueryParam) (*http.Response, error) {
	u, _ := addQueryParam(apiUrl+"/"+strings.TrimLeft(uri, "/"), params...)

	req, err := http.NewRequest("GET", u, nil)
	req.Header.Set("Authorization", "Bearer "+c.accessToken)
	if err != nil {
		return nil, err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	bytes, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	if res.StatusCode >= 200 && res.StatusCode < 300 {
		if err := json.Unmarshal(bytes, &jsonResponse); err != nil {
			return nil, err
		}

		return res, nil
	}

	return nil, &Error{
		StatusCode: res.StatusCode,
		URL:        res.Request.URL.String(),
		Body:       string(bytes),
	}
}

func getId(uri string) int {
	s := strings.SplitN(uri, "/", -1)
	id, _ := strconv.Atoi(s[len(s)-1])
	return id
}
