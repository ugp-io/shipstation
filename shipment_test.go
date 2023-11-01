package shipstation

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"
)

func TestListShipments(t *testing.T) {

	client := NewClient(os.Getenv("API_KEY"), os.Getenv("API_SECRET"))
	// orderStatus := "awaiting_shipment"
	// storeID := 17797
	// orderNumber := "113-9480867-6648261"
	resp, err := client.Shipment.List(ShipmentListParams{})
	if err != nil {
		t.Errorf("Shipment.List returned error: %v", err)
	}

	bolB, _ := json.Marshal(resp)
	fmt.Println(string(bolB))

}

func TestVoidLabel(t *testing.T) {

	client := NewClient(os.Getenv("API_KEY"), os.Getenv("API_SECRET"))
	// orderStatus := "awaiting_shipment"
	// storeID := 17797
	// orderNumber := "113-9480867-6648261"
	resp, err := client.Shipment.VoidLabel(ShipmentVoidLabelParams{
		ShipmentID: 554682784,
	})
	if err != nil {
		t.Errorf("Shipment.List returned error: %v", err)
	}

	bolB, _ := json.Marshal(resp)
	fmt.Println(string(bolB))

}
