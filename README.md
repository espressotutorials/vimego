# Vimego

Vimego is a simple Vimeo Go client.

## Usage 

### Functions
```go
c := vimego.New{accessToken: "YOUR_ACCESS_TOKEN"}

json, err := c.ListMyProjects()
json, err := c.ListProjectsOfUser(1234)
json, err := c.GetMyProject(4567)
json, err := c.GetProjectOfUser(1234, 4567)
json, err := c.ListMyProjectVideos(4567)
json, err := c.ListProjectVideosOfUser(1234, 4567)

json, err := c.GetVideo(8910)
json, err := c.GetVideoTexttracks(8910)
```

### Set parameters
```go
json, err := c.ListMyProjects(Page(2), PerPage(25), Fields{"name", "uri"})
```

## Examples
### Pagination
```go
page := 1
hasNextPage := true
for hasNextPage {
	json, err := c.ListMyProjects(Page(page), PerPage(25), Fields{"uri", "name"})
	if err != nil {
	    fmt.Println(err)	
    }   
	
	fmt.Println(json)
	
	hasNextPage = len(json.Paginate.Next) != 0
	page++
}
```