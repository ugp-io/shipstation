package shipstation

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"
)

func TestListStores(t *testing.T) {

	client := NewClient(os.Getenv("API_KEY"), os.Getenv("API_SECRET"))
	resp, err := client.Store.List(StoreListParams{})
	if err != nil {
		t.Errorf("Store.List returned error: %v", err)
	}

	bolB, _ := json.Marshal(resp)
	fmt.Println(string(bolB))

}
