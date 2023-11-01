package shipstation

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"
)

func TestListOrders(t *testing.T) {

	client := NewClient(os.Getenv("API_KEY"), os.Getenv("API_SECRET"))
	orderStatus := "awaiting_shipment"
	storeID := 17797
	// orderNumber := "113-9480867-6648261"
	resp, err := client.Order.List(OrderListParams{
		StoreID:     &storeID,
		OrderStatus: &orderStatus,
	})
	if err != nil {
		t.Errorf("Order.List returned error: %v", err)
	}

	// bolB, _ := json.Marshal(resp)
	// fmt.Println(string(bolB))
	if resp != nil {

		//Print out all skus
		for _, order := range resp.Orders {
			// fmt.Println(order.OrderNumber)
			for _, item := range *order.Items {
				fmt.Println(*item.Sku)
			}
		}

	}

}

func TestListOrdersByTag(t *testing.T) {

	client := NewClient(os.Getenv("API_KEY"), os.Getenv("API_SECRET"))
	// orderStatus := "awaiting_shipment"
	// storeID := 17797
	_, err := client.Order.ListByTag(OrderListByTagParams{
		OrderStatus: "awaiting_shipment",
		TagID:       67150,
	})
	if err != nil {
		t.Errorf("Order.List returned error: %v", err)
	}

	// bolB, _ := json.Marshal(resp)
	// fmt.Println(string(bolB))

}

func TestGetOrder(t *testing.T) {

	client := NewClient(os.Getenv("API_KEY"), os.Getenv("API_SECRET"))

	resp, err := client.Order.Get(OrderGetParams{
		OrderID: 743647382,
	})
	if err != nil {
		t.Errorf("Order.List returned error: %v", err)
	}
	bolB, _ := json.Marshal(resp)
	fmt.Println(string(bolB))

}

func TestCreateLabelForOrder(t *testing.T) {

	client := NewClient(os.Getenv("API_KEY"), os.Getenv("API_SECRET"))

	resp, err := client.Order.CreateLabel(OrderCreateLabelParams{
		OrderID: 737986330,
	})
	if err != nil {
		t.Errorf("Order.List returned error: %v", err)
	}
	bolB, _ := json.Marshal(resp)
	fmt.Println(string(bolB))

}

func TestAddTagToOrder(t *testing.T) {

	client := NewClient(os.Getenv("API_KEY"), os.Getenv("API_SECRET"))

	resp, err := client.Order.AddTag(OrderAddTagParams{
		OrderID: 737986330,
		TagID:   6018,
	})
	if err != nil {
		t.Errorf("Order.List returned error: %v", err)
	}
	bolB, _ := json.Marshal(resp)
	fmt.Println(string(bolB))

}

func TestRemoveTagFromOrder(t *testing.T) {

	client := NewClient(os.Getenv("API_KEY"), os.Getenv("API_SECRET"))

	resp, err := client.Order.RemoveTag(OrderRemoveTagParams{
		OrderID: 737986330,
		TagID:   6018,
	})
	if err != nil {
		t.Errorf("Order.List returned error: %v", err)
	}
	bolB, _ := json.Marshal(resp)
	fmt.Println(string(bolB))

}
