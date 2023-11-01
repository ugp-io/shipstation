package shipstation

import (
	"fmt"
	"net/url"
	"strconv"
)

const shipmentsBasePath = "shipments"

type ShipmentServiceOp struct {
	client *Client
}

type ShipmentService interface {
	List(ShipmentListParams) (*ShipmentsResource, error)
	VoidLabel(ShipmentVoidLabelParams) (*ShipmentVoidLabelResource, error)
}

type ShipmentListParams struct {
	RecipientName        *string
	RecipientCountryCode *string
	OrderNumber          *string
	OrderID              *int
	CarrierCode          *string
	ServiceCode          *string
	TrackingNumber       *string
	CustomsCountryCode   *string
	CreateDateStart      *string
	CreateDateEnd        *string
	ShipDateStart        *string
	ShipDateEnd          *string
	VoidDateStart        *string
	VoidDateEnd          *string
	StoreID              *int
	IncludeShipmentItems *bool
	SortBy               *string
	SortDir              *string
	Page                 *int
	PageSize             *int
}

type ShipmentVoidLabelParams struct {
	ShipmentID int `json:"shipmentId"`
}
type ShipmentVoidLabelResource struct {
	Approved bool   `json:"approved"`
	Message  string `json:"message"`
}

type ShipmentsResource struct {
	Shipments []Shipment `json:"shipments"`
	Total     int        `json:"total"`
	Page      int        `json:"page"`
	Pages     int        `json:"pages"`
}

type Shipment struct {
	ShipmentID          *int              `json:"shipmentId,omitempty"`
	OrderID             *int              `json:"orderId,omitempty"`
	UserID              *string           `json:"userId,omitempty"`
	OrderNumber         *string           `json:"orderNumber,omitempty"`
	CreateDate          *string           `json:"createDate,omitempty"`
	ShipDate            *string           `json:"shipDate,omitempty"`
	ShipmentCost        *float64          `json:"shipmentCost,omitempty"`
	InsuranceCost       *float64          `json:"insuranceCost,omitempty"`
	TrackingNumber      *string           `json:"trackingNumber,omitempty"`
	IsReturnLabel       *bool             `json:"isReturnLabel,omitempty"`
	BatchNumber         *string           `json:"batchNumber,omitempty"`
	CarrierCode         *string           `json:"carrierCode,omitempty"`
	ServiceCode         *string           `json:"serviceCode,omitempty"`
	PackageCode         *string           `json:"packageCode,omitempty"`
	Confirmation        *string           `json:"confirmation,omitempty"`
	WarehouseID         *int              `json:"warehouseId,omitempty"`
	Voided              *bool             `json:"voided,omitempty"`
	VoidDate            *string           `json:"voidDate,omitempty"`
	MarketplaceNotified *bool             `json:"marketplaceNotified,omitempty"`
	NotifyErrorMessage  *string           `json:"notifyErrorMessage,omitempty"`
	ShipTo              *ShipTo           `json:"shipTo,omitempty"`
	Weight              *Weight           `json:"weight,omitempty"`
	Dimensions          *string           `json:"dimensions,omitempty"`
	InsuranceOptions    *InsuranceOptions `json:"insuranceOptions,omitempty"`
	AdvancedOptions     *AdvancedOptions  `json:"advancedOptions,omitempty"`
	ShipmentItems       *[]ShipmentItems  `json:"shipmentItems,omitempty"`
	LabelData           *string           `json:"labelData,omitempty"`
	FormData            *string           `json:"formData,omitempty"`
}
type ShipmentItems struct {
	OrderItemID       *int    `json:"orderItemId,omitempty"`
	LineItemKey       *string `json:"lineItemKey,omitempty"`
	Sku               *string `json:"sku,omitempty"`
	Name              *string `json:"name,omitempty"`
	ImageURL          *string `json:"imageUrl,omitempty"`
	Weight            *string `json:"weight,omitempty"`
	Quantity          *int    `json:"quantity,omitempty"`
	UnitPrice         *int    `json:"unitPrice,omitempty"`
	WarehouseLocation *string `json:"warehouseLocation,omitempty"`
	Options           *string `json:"options,omitempty"`
	ProductID         *int    `json:"productId,omitempty"`
	FulfillmentSku    *string `json:"fulfillmentSku,omitempty"`
}

func (s *ShipmentServiceOp) List(params ShipmentListParams) (*ShipmentsResource, error) {

	// URI
	uri := fmt.Sprintf("%s/%s", apiURI, shipmentsBasePath)

	values := url.Values{}
	if params.RecipientName != nil {
		values.Add("recipientName", *params.RecipientName)
	}
	if params.RecipientCountryCode != nil {
		values.Add("recipientCountryCode", *params.RecipientCountryCode)
	}
	if params.OrderNumber != nil {
		values.Add("orderNumber", *params.OrderNumber)
	}
	if params.OrderID != nil {
		values.Add("orderId", strconv.Itoa(*params.OrderID))
	}
	if params.CarrierCode != nil {
		values.Add("carrierCode", *params.CarrierCode)
	}
	if params.ServiceCode != nil {
		values.Add("serviceCode", *params.ServiceCode)
	}
	if params.TrackingNumber != nil {
		values.Add("trackingNumber", *params.TrackingNumber)
	}
	if params.CustomsCountryCode != nil {
		values.Add("customsCountryCode", *params.CustomsCountryCode)
	}
	if params.CreateDateStart != nil {
		values.Add("createDateStart", *params.CreateDateStart)
	}
	if params.CreateDateEnd != nil {
		values.Add("createDateEnd", *params.CreateDateEnd)
	}
	if params.ShipDateStart != nil {
		values.Add("shipDateStart", *params.ShipDateStart)
	}
	if params.ShipDateEnd != nil {
		values.Add("shipDateEnd", *params.ShipDateEnd)
	}
	if params.VoidDateStart != nil {
		values.Add("voidDateStart", *params.VoidDateStart)
	}
	if params.VoidDateEnd != nil {
		values.Add("voidDateEnd", *params.VoidDateEnd)
	}
	if params.StoreID != nil {
		values.Add("storeId", strconv.Itoa(*params.StoreID))
	}
	if params.IncludeShipmentItems != nil {
		values.Add("includeShipmentItems", strconv.FormatBool(*params.IncludeShipmentItems))
	}
	if params.SortBy != nil {
		values.Add("sortBy", *params.SortBy)
	}
	if params.SortDir != nil {
		values.Add("sortDir", *params.SortDir)
	}
	if params.Page != nil {
		values.Add("page", strconv.Itoa(*params.Page))
	}
	if params.PageSize != nil {
		values.Add("pageSize", strconv.Itoa(*params.PageSize))
	}

	url := uri + "?" + values.Encode()

	var resp ShipmentsResource

	errRequest := s.client.Request("GET", url, nil, &resp)
	if errRequest != nil {
		return nil, errRequest
	}

	return &resp, nil
}

func (s *ShipmentServiceOp) VoidLabel(params ShipmentVoidLabelParams) (*ShipmentVoidLabelResource, error) {

	// URI
	uri := fmt.Sprintf("%s/%s", apiURI, shipmentsBasePath)

	url := uri + "/voidlabel"

	var resp ShipmentVoidLabelResource

	errRequest := s.client.Request("POST", url, params, &resp)
	if errRequest != nil {
		return nil, errRequest
	}

	return &resp, nil
}
