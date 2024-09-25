package shipstation

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
)

const apiURI = "https://ssapi.shipstation.com"

type Client struct {
	APIKey    string
	APISecret string
	Order     OrderService
	Shipment  ShipmentService
	Store     StoreService
	Tag       TagService
}

type DefaultResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

func NewClient(apiKey, apiSecret string) *Client {

	c := &Client{
		APIKey:    apiKey,
		APISecret: apiSecret,
	}

	c.Order = &OrderServiceOp{client: c}
	c.Shipment = &ShipmentServiceOp{client: c}
	c.Store = &StoreServiceOp{client: c}
	c.Tag = &TagServiceOp{client: c}

	return c

}

func (c *Client) Request(method string, url string, body interface{}, v interface{}) error {

	var bodyReader io.Reader
	if body != nil {
		requestJson, errMarshal := json.Marshal(body)
		if errMarshal != nil {
			return errMarshal
		}

		bodyReader = bytes.NewBuffer(requestJson)
	}

	//HTTP
	httpReq, errNewRequest := http.NewRequest(method, url, bodyReader)
	if errNewRequest != nil {
		return errNewRequest
	}

	// Basic Auth
	httpReq.SetBasicAuth(c.APIKey, c.APISecret)

	// Content Type
	httpReq.Header.Set("Content-Type", "application/json")

	//Client
	client := &http.Client{}
	resp, errDo := client.Do(httpReq)

	if resp != nil {
		defer resp.Body.Close()
	}
	if errDo != nil {
		return errDo
	}

	if v != nil {
		decoder := json.NewDecoder(resp.Body)
		errDecode := decoder.Decode(&v)
		if errDecode != nil {
			return errDecode
		}
	}
	return nil
}
