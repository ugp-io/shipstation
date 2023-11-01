package shipstation

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"
)

func TestListTags(t *testing.T) {

	client := NewClient(os.Getenv("API_KEY"), os.Getenv("API_SECRET"))
	resp, err := client.Tag.List(TagListParams{})
	if err != nil {
		t.Errorf("Tag.List returned error: %v", err)
	}

	bolB, _ := json.Marshal(resp)
	fmt.Println(string(bolB))

}
