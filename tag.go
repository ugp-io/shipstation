package shipstation

import (
	"fmt"
)

const tagsBasePath = "accounts/listtags"

type TagServiceOp struct {
	client *Client
}

type TagService interface {
	List(TagListParams) (*[]Tag, error)
}

type TagListParams struct {
}
type Tag struct {
	TagID *int    `json:"tagId,omitempty"`
	Name  *string `json:"name,omitempty"`
	Color *string `json:"color,omitempty"`
}

func (s *TagServiceOp) List(params TagListParams) (*[]Tag, error) {

	// URI
	uri := fmt.Sprintf("%s/%s", apiURI, tagsBasePath)
	url := uri

	var resp []Tag

	errRequest := s.client.Request("GET", url, nil, &resp)
	if errRequest != nil {
		return nil, errRequest
	}

	return &resp, nil
}
