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

type Error struct {
	StatusCode int
	URL        string
	Body       string
	RateLimit  RateLimit
}

func (e Error) Error() string {
	return fmt.Sprintf(
		"%d %s: %s (RateLimit: Remaining: %d | Limit: %d | Reset: %s )",
		e.StatusCode,
		e.URL,
		e.Body,
		e.RateLimit.Remaining,
		e.RateLimit.Limit,
		e.RateLimit.Reset,
	)
}

type RateLimit struct {
	Limit     int
	Remaining int
	Reset     time.Time
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
	AccessToken string
	BaseURL     string
	HttpClient  *http.Client
}

func New(accessToken string) *Client {
	return &Client{
		AccessToken: accessToken,
		BaseURL:     "https://api.vimeo.com",
		HttpClient:  &http.Client{},
	}
}

func (c *Client) get(uri string, jsonResponse interface{}, params ...QueryParam) (*http.Response, error) {
	u, _ := addQueryParam(c.BaseURL+"/"+strings.TrimLeft(uri, "/"), params...)

	req, err := http.NewRequest("GET", u, nil)
	req.Header.Set("Authorization", "Bearer "+c.AccessToken)
	if err != nil {
		return nil, err
	}

	res, err := c.HttpClient.Do(req)
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
		RateLimit:  parseRateLimit(res),
	}
}

func getId(uri string) int {
	s := strings.SplitN(uri, "/", -1)
	id, _ := strconv.Atoi(s[len(s)-1])
	return id
}

func parseRateLimit(r *http.Response) RateLimit {
	var l RateLimit

	if reset := r.Header.Get("X-RateLimit-Reset"); reset != "" {
		t, err := time.Parse(time.RFC3339, reset)
		if err != nil {
			l.Reset = t
		}
	}

	if remaining := r.Header.Get("X-RateLimit-Remaining"); remaining != "" {
		l.Remaining, _ = strconv.Atoi(remaining)
	}

	if limit := r.Header.Get("X-RateLimit-Limit"); limit != "" {
		l.Remaining, _ = strconv.Atoi(limit)
	}

	return l
}
