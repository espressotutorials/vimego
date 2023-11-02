package vimeo

import (
	"fmt"
	"time"
)

type Project struct {
	CreatedTime             time.Time    `json:"created_time,omitempty"`
	ModifiedTime            time.Time    `json:"modified_time,omitempty"`
	LastUserActionEventDate time.Time    `json:"last_user_action_event_date,omitempty"`
	Name                    string       `json:"name,omitempty"`
	Privacy                 VimeoPrivacy `json:"privacy,omitempty"`
	ResourceKey             string       `json:"resource_key,omitempty"`
	URI                     string       `json:"uri,omitempty"`
	Link                    string       `json:"link,omitempty"`
	ManageLink              string       `json:"manage_link,omitempty"`
	PinnedOn                string       `json:"pinned_on"`
	IsPinned                bool         `json:"is_pinned,omitempty"`
	IsPrivateToUser         bool         `json:"is_private_to_user,omitempty"`
	User                    User         `json:"user,omitempty"`
	AccessGrant             string       `json:"access_grant,omitempty"`
}

func (p Project) GetId() int {
	return getId(p.URI)
}

func (c *Client) ListMyProjects(params ...QueryParam) (*VimeoResponse[Project], error) {
	var j = &VimeoResponse[Project]{}
	_, err := c.get("me/projects", j, params...)
	if err != nil {
		return nil, err
	}

	return j, nil
}

func (c *Client) ListProjectsOfUser(u int, params ...QueryParam) (*VimeoResponse[Project], error) {
	var j = &VimeoResponse[Project]{}
	_, err := c.get(fmt.Sprintf("users/%d/projects", u), j, params...)
	if err != nil {
		return nil, err
	}

	return j, nil
}

func (c *Client) GetMyProject(id int, params ...QueryParam) (*Project, error) {
	var j = &Project{}
	_, err := c.get(fmt.Sprintf("me/projects/%d", id), j, params...)
	if err != nil {
		return nil, err
	}

	return j, nil
}

func (c *Client) GetProjectOfUser(u, id int, params ...QueryParam) (*Project, error) {
	var j = &Project{}
	_, err := c.get(fmt.Sprintf("users/%d/projects/%d", u, id), j, params...)
	if err != nil {
		return nil, err
	}

	return j, nil
}

func (c *Client) ListMyProjectVideos(id int, params ...QueryParam) (*VimeoResponse[Video], error) {
	var j = &VimeoResponse[Video]{}
	_, err := c.get(fmt.Sprintf("me/projects/%d/videos", id), j, params...)
	if err != nil {
		return nil, err
	}

	return j, nil
}

func (c *Client) ListProjectVideosOfUser(u, id int, params ...QueryParam) (*VimeoResponse[Video], error) {
	var j = &VimeoResponse[Video]{}
	_, err := c.get(fmt.Sprintf("users/%d/projects/%d/videos", u, id), j, params...)
	if err != nil {
		return nil, err
	}

	return j, nil
}
