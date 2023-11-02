package vimeo

import "fmt"

type Video struct {
	URI         string        `json:"uri,omitempty"`
	Name        string        `json:"name,omitempty"`
	Description string        `json:"description,omitempty"`
	Type        string        `json:"type,omitempty"`
	Link        string        `json:"link,omitempty"`
	Duration    int           `json:"duration,omitempty"`
	Width       int           `json:"width,omitempty"`
	Height      int           `json:"height,omitempty"`
	Language    string        `json:"language,omitempty"`
	Options     VimeoOptions  `json:"options,omitempty"`
	Total       int           `json:"total,omitempty"`
	Files       []VimeoFile   `json:"files,omitempty"`
	Pictures    VimeoPictures `json:"pictures,omitempty"`
}

func (v Video) GetId() int {
	return getId(v.URI)
}

type Texttrack struct {
	Uri                string `json:"uri,omitempty"`
	Active             bool   `json:"active,omitempty"`
	Type               string `json:"type,omitempty"`
	Language           string `json:"language,omitempty"`
	DisplayLanguage    string `json:"display_language,omitempty"`
	Id                 int    `json:"id,omitempty"`
	Link               string `json:"link,omitempty"`
	LinkExpiresTime    int    `json:"link_expires_time,omitempty"`
	HLSLink            string `json:"hls_link,omitempty"`
	HLSLinkExpiresTime int    `json:"hls_link_expires_time,omitempty"`
	Name               string `json:"name,omitempty"`
}

func (c *Client) GetVideo(id int, params ...QueryParam) (*Video, error) {
	var j = &Video{}
	_, err := c.get(fmt.Sprintf("videos/%d", id), j, params...)
	if err != nil {
		return nil, err
	}

	return j, nil
}

func (c *Client) ListVideoTexttracks(id int, params ...QueryParam) (*VimeoResponse[Texttrack], error) {
	var j = &VimeoResponse[Texttrack]{}
	_, err := c.get(fmt.Sprintf("videos/%d/texttracks", id), j, params...)
	if err != nil {
		return nil, err
	}

	return j, nil
}
