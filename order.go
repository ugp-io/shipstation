package shipstation

import (
	"fmt"
	"net/url"
	"strconv"
)

const ordersBasePath = "orders"

type OrderServiceOp struct {
	client *Client
}

type OrderService interface {
	List(OrderListParams) (*OrdersResource, error)
	ListByTag(OrderListByTagParams) (*OrdersResource, error)
	Get(OrderGetParams) (*Order, error)
	CreateUpdate(Order) (*Order, error)
	CreateLabel(OrderCreateLabelParams) (*OrderCreateLabelResource, error)
	AddTag(OrderAddTagParams) (*OrderAddTagResource, error)
	RemoveTag(OrderRemoveTagParams) (*OrderRemoveTagResource, error)
	MarkOrderShipped(params MarkOrderShippedParams) (*Order, error)
}

type OrderGetParams struct {
	OrderID int
}

type OrderListParams struct {
	CustomerName       *string
	ItemKeyword        *string
	CreateDateStart    *string
	CreateDateEnd      *string
	CustomsCountryCode *string
	ModifyDateStart    *string
	ModifyDateEnd      *string
	OrderDateStart     *string
	OrderDateEnd       *string
	OrderNumber        *string
	OrderStatus        *string
	PaymentDateStart   *string
	PaymentDateEnd     *string
	StoreID            *int
	SortBy             *string
	SortDir            *string
	Page               *int
	PageSize           *int
}

type OrderListByTagParams struct {
	OrderStatus string
	TagID       int
	Page        *int
	PageSize    *int
}

type OrderCreateLabelParams struct {
	OrderID              int                   `json:"orderId,omitempty"`
	CarrierCode          string                `json:"carrierCode,omitempty"`
	ServiceCode          string                `json:"serviceCode,omitempty"`
	Confirmation         string                `json:"confirmation,omitempty"`
	ShipDate             string                `json:"shipDate,omitempty"`
	Weight               *Weight               `json:"weight,omitempty"`
	Dimensions           *Dimensions           `json:"dimensions,omitempty"`
	InsuranceOptions     *InsuranceOptions     `json:"insuranceOptions,omitempty"`
	InternationalOptions *InternationalOptions `json:"internationalOptions,omitempty"`
	AdvancedOptions      *AdvancedOptions      `json:"advancedOptions,omitempty"`
	TestLabel            bool                  `json:"testLabel,omitempty"`
}

type OrderAddTagParams struct {
	OrderID int `json:"orderId,omitempty"`
	TagID   int `json:"tagId,omitempty"`
}

type OrderAddTagResource struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

type OrderRemoveTagParams struct {
	OrderID int `json:"orderId,omitempty"`
	TagID   int `json:"tagId,omitempty"`
}

type OrderRemoveTagResource struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

type MarkOrderShippedParams struct {
	OrderId            int     `json:"orderId,omitempty"`
	CarrierCode        string  `json:"carrierCode,omitempty"`
	ShipDate           *string `json:"shipDate,omitempty"`
	TrackingNumber     *string `json:"trackingNumber,omitempty"`
	NotifyCustomer     *bool   `json:"notifyCustomer,omitempty"`
	NotifySalesChannel *bool   `json:"notifySalesChannel,omitempty"`
}

type OrderCreateLabelResource struct {
	ShipmentID     int     `json:"shipmentId,omitempty"`
	ShipmentCost   float64 `json:"shipmentCost,omitempty"`
	InsuranceCost  float64 `json:"insuranceCost,omitempty"`
	TrackingNumber string  `json:"trackingNumber,omitempty"`
	LabelData      string  `json:"labelData,omitempty"`
	FormData       string  `json:"formData,omitempty"`
}

type OrdersResource struct {
	Orders []Order `json:"orders"`
	Total  int     `json:"total"`
	Page   int     `json:"page"`
	Pages  int     `json:"pages"`
}

type Order struct {
	OrderID                   *int                  `json:"orderId,omitempty"`
	OrderNumber               *string               `json:"orderNumber,omitempty"`
	OrderKey                  *string               `json:"orderKey,omitempty"`
	OrderDate                 *string               `json:"orderDate,omitempty"`
	CreateDate                *string               `json:"createDate,omitempty"`
	ModifyDate                *string               `json:"modifyDate,omitempty"`
	PaymentDate               *string               `json:"paymentDate,omitempty"`
	ShipByDate                *string               `json:"shipByDate,omitempty"`
	OrderStatus               *string               `json:"orderStatus,omitempty"`
	CustomerID                *int                  `json:"customerId,omitempty"`
	CustomerUsername          *string               `json:"customerUsername,omitempty"`
	CustomerEmail             *string               `json:"customerEmail,omitempty"`
	BillTo                    *BillTo               `json:"billTo,omitempty"`
	ShipTo                    *ShipTo               `json:"shipTo,omitempty"`
	Items                     *[]Items              `json:"items,omitempty"`
	OrderTotal                *float64              `json:"orderTotal,omitempty"`
	AmountPaid                *float64              `json:"amountPaid,omitempty"`
	TaxAmount                 *float64              `json:"taxAmount,omitempty"`
	ShippingAmount            *float64              `json:"shippingAmount,omitempty"`
	CustomerNotes             *string               `json:"customerNotes,omitempty"`
	InternalNotes             *string               `json:"internalNotes,omitempty"`
	Gift                      *bool                 `json:"gift,omitempty"`
	GiftMessage               *string               `json:"giftMessage,omitempty"`
	PaymentMethod             *string               `json:"paymentMethod,omitempty"`
	RequestedShippingService  *string               `json:"requestedShippingService,omitempty"`
	CarrierCode               *string               `json:"carrierCode,omitempty"`
	ServiceCode               *string               `json:"serviceCode,omitempty"`
	PackageCode               *string               `json:"packageCode,omitempty"`
	Confirmation              *string               `json:"confirmation,omitempty"`
	ShipDate                  *string               `json:"shipDate,omitempty"`
	HoldUntilDate             *string               `json:"holdUntilDate,omitempty"`
	Weight                    *Weight               `json:"weight,omitempty"`
	Dimensions                *Dimensions           `json:"dimensions,omitempty"`
	InsuranceOptions          *InsuranceOptions     `json:"insuranceOptions,omitempty"`
	InternationalOptions      *InternationalOptions `json:"internationalOptions,omitempty"`
	AdvancedOptions           *AdvancedOptions      `json:"advancedOptions,omitempty"`
	TagIds                    *[]int                `json:"tagIds,omitempty"`
	UserID                    *string               `json:"userId,omitempty"`
	ExternallyFulfilled       *bool                 `json:"externallyFulfilled,omitempty"`
	ExternallyFulfilledBy     *string               `json:"externallyFulfilledBy,omitempty"`
	ExternallyFulfilledByID   *string               `json:"externallyFulfilledById,omitempty"`
	ExternallyFulfilledByName *string               `json:"externallyFulfilledByName,omitempty"`
	LabelMessages             *string               `json:"labelMessages,omitempty"`
}
type BillTo struct {
	Name            *string `json:"name,omitempty"`
	Compstring      *string `json:"compstring,omitempty"`
	Street1         *string `json:"street1,omitempty"`
	Street2         *string `json:"street2,omitempty"`
	Street3         *string `json:"street3,omitempty"`
	City            *string `json:"city,omitempty"`
	State           *string `json:"state,omitempty"`
	PostalCode      *string `json:"postalCode,omitempty"`
	Country         *string `json:"country,omitempty"`
	Phone           *string `json:"phone,omitempty"`
	Residential     *string `json:"residential,omitempty"`
	AddressVerified *string `json:"addressVerified,omitempty"`
}
type ShipTo struct {
	Name            *string `json:"name,omitempty"`
	Compstring      *string `json:"compstring,omitempty"`
	Street1         *string `json:"street1,omitempty"`
	Street2         *string `json:"street2,omitempty"`
	Street3         *string `json:"street3,omitempty"`
	City            *string `json:"city,omitempty"`
	State           *string `json:"state,omitempty"`
	PostalCode      *string `json:"postalCode,omitempty"`
	Country         *string `json:"country,omitempty"`
	Phone           *string `json:"phone,omitempty"`
	Residential     *bool   `json:"residential,omitempty"`
	AddressVerified *string `json:"addressVerified,omitempty"`
}
type Items struct {
	OrderItemID       *int          `json:"orderItemId,omitempty"`
	LineItemKey       *string       `json:"lineItemKey,omitempty"`
	Sku               *string       `json:"sku,omitempty"`
	Name              *string       `json:"name,omitempty"`
	ImageURL          *string       `json:"imageUrl,omitempty"`
	Weight            *Weight       `json:"weight,omitempty"`
	Quantity          *int          `json:"quantity,omitempty"`
	UnitPrice         *float64      `json:"unitPrice,omitempty"`
	TaxAmount         *float64      `json:"taxAmount,omitempty"`
	ShippingAmount    *float64      `json:"shippingAmount,omitempty"`
	WarehouseLocation *string       `json:"warehouseLocation,omitempty"`
	Options           *[]ItemOption `json:"options,omitempty"`
	ProductID         *int          `json:"productId,omitempty"`
	FulfillmentSku    *string       `json:"fulfillmentSku,omitempty"`
	Adjustment        *bool         `json:"adjustment,omitempty"`
	Upc               *string       `json:"upc,omitempty"`
	CreateDate        *string       `json:"createDate,omitempty"`
	ModifyDate        *string       `json:"modifyDate,omitempty"`
}

type ItemOption struct {
	Name  *string `json:"name,omitempty"`
	Value *string `json:"value,omitempty"`
}
type Weight struct {
	Value       *float64 `json:"value,omitempty"`
	Units       *string  `json:"units,omitempty"`
	WeightUnits *int     `json:"WeightUnits,omitempty"`
}
type InsuranceOptions struct {
	Provider       *string  `json:"provider,omitempty"`
	InsureShipment *bool    `json:"insureShipment,omitempty"`
	InsuredValue   *float64 `json:"insuredValue,omitempty"`
}
type InternationalOptions struct {
	Contents     *string        `json:"contents,omitempty"`
	CustomsItems *[]CustomsItem `json:"customsItems,omitempty"`
	NonDelivery  *string        `json:"nonDelivery,omitempty"`
}
type CustomsItem struct {
	CustomsItemID        *int     `json:"customsItemId,omitempty"`
	Description          *string  `json:"description,omitempty"`
	Quantity             *int     `json:"quantity,omitempty"`
	Value                *float64 `json:"value,omitempty"`
	HarmonizedTariffCode *string  `json:"harmonizedTariffCode,omitempty"`
	CountryOfOrigin      *string  `json:"countryOfOrigin,omitempty"`
}

type AdvancedOptions struct {
	WarehouseID          *int    `json:"warehouseId,omitempty"`
	NonMachinable        *bool   `json:"nonMachinable,omitempty"`
	SaturdayDelivery     *bool   `json:"saturdayDelivery,omitempty"`
	ContainsAlcohol      *bool   `json:"containsAlcohol,omitempty"`
	MergedOrSplit        *bool   `json:"mergedOrSplit,omitempty"`
	MergedIds            *[]int  `json:"mergedIds,omitempty"`
	ParentID             *int    `json:"parentId,omitempty"`
	StoreID              *int    `json:"storeId,omitempty"`
	CustomField1         *string `json:"customField1,omitempty"`
	CustomField2         *string `json:"customField2,omitempty"`
	CustomField3         *string `json:"customField3,omitempty"`
	Source               *string `json:"source,omitempty"`
	BillToParty          *string `json:"billToParty,omitempty"`
	BillToAccount        *string `json:"billToAccount,omitempty"`
	BillToPostalCode     *string `json:"billToPostalCode,omitempty"`
	BillToCountryCode    *string `json:"billToCountryCode,omitempty"`
	BillToMyOtherAccount *int    `json:"billToMyOtherAccount,omitempty"`
}

type Dimensions struct {
	Height *float64 `json:"height,omitempty"`
	Width  *float64 `json:"width,omitempty"`
	Length *float64 `json:"length,omitempty"`
	Units  *string  `json:"units,omitempty"`
}

type OrderCreateUpdateParams struct {
	OrderKey    *string `json:"orderKey,omitempty"`
	CarrierCode *string `json:"carrierCode,omitempty"`
	ServiceCode *string `json:"serviceCode,omitempty"`
	Weight      *Weight `json:"weight,omitempty"`
}

func (s *OrderServiceOp) List(params OrderListParams) (*OrdersResource, error) {

	// URI
	uri := fmt.Sprintf("%s/%s", apiURI, ordersBasePath)

	values := url.Values{}
	if params.CustomerName != nil {
		values.Add("customerName", *params.CustomerName)
	}
	if params.ItemKeyword != nil {
		values.Add("itemKeyword", *params.ItemKeyword)
	}
	if params.CreateDateStart != nil {
		values.Add("createDateStart", *params.CreateDateStart)
	}
	if params.CreateDateEnd != nil {
		values.Add("createDateEnd", *params.CreateDateEnd)
	}
	if params.CustomsCountryCode != nil {
		values.Add("customsCountryCode", *params.CustomsCountryCode)
	}
	if params.ModifyDateStart != nil {
		values.Add("modifyDateStart", *params.ModifyDateStart)
	}
	if params.ModifyDateEnd != nil {
		values.Add("modifyDateEnd", *params.ModifyDateEnd)
	}
	if params.OrderDateStart != nil {
		values.Add("orderDateStart", *params.OrderDateStart)
	}
	if params.OrderDateEnd != nil {
		values.Add("orderDateEnd", *params.OrderDateEnd)
	}
	if params.OrderNumber != nil {
		values.Add("orderNumber", *params.OrderNumber)
	}
	if params.OrderStatus != nil {
		values.Add("orderStatus", *params.OrderStatus)
	}
	if params.PaymentDateStart != nil {
		values.Add("paymentDateStart", *params.PaymentDateStart)
	}
	if params.PaymentDateEnd != nil {
		values.Add("paymentDateEnd", *params.PaymentDateEnd)
	}
	if params.StoreID != nil {
		values.Add("storeId", strconv.Itoa(*params.StoreID))
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

	var resp OrdersResource

	errRequest := s.client.Request("GET", url, nil, &resp)
	if errRequest != nil {
		return nil, errRequest
	}

	return &resp, nil
}

func (s *OrderServiceOp) ListByTag(params OrderListByTagParams) (*OrdersResource, error) {

	// URI
	uri := fmt.Sprintf("%s/%s", apiURI, ordersBasePath)

	values := url.Values{}
	values.Add("orderStatus", params.OrderStatus)
	values.Add("tagId", strconv.Itoa(params.TagID))
	if params.Page != nil {
		values.Add("page", strconv.Itoa(*params.Page))
	}
	if params.PageSize != nil {
		values.Add("pageSize", strconv.Itoa(*params.PageSize))
	}

	url := uri + "/listbytag?" + values.Encode()

	var resp OrdersResource

	errRequest := s.client.Request("GET", url, nil, &resp)
	if errRequest != nil {
		return nil, errRequest
	}

	return &resp, nil
}

func (s *OrderServiceOp) Get(params OrderGetParams) (*Order, error) {

	// URI
	uri := fmt.Sprintf("%s/%s", apiURI, ordersBasePath)

	url := uri + "/" + strconv.Itoa(params.OrderID)

	var resp Order

	errRequest := s.client.Request("GET", url, nil, &resp)
	if errRequest != nil {
		return nil, errRequest
	}

	return &resp, nil
}

func (s *OrderServiceOp) CreateLabel(params OrderCreateLabelParams) (*OrderCreateLabelResource, error) {

	// URI
	uri := fmt.Sprintf("%s/%s", apiURI, ordersBasePath)

	url := uri + "/createlabelfororder"

	var resp OrderCreateLabelResource

	errRequest := s.client.Request("POST", url, params, &resp)
	if errRequest != nil {
		return nil, errRequest
	}

	return &resp, nil
}

func (s *OrderServiceOp) CreateUpdate(params Order) (*Order, error) {

	// URI
	uri := fmt.Sprintf("%s/%s", apiURI, ordersBasePath)

	url := uri + "/createorder"

	var resp Order

	errRequest := s.client.Request("POST", url, params, &resp)
	if errRequest != nil {
		return nil, errRequest
	}

	return &resp, nil
}

func (s *OrderServiceOp) AddTag(params OrderAddTagParams) (*OrderAddTagResource, error) {

	// URI
	uri := fmt.Sprintf("%s/%s", apiURI, ordersBasePath)

	url := uri + "/addtag"

	var resp OrderAddTagResource

	errRequest := s.client.Request("POST", url, params, &resp)
	if errRequest != nil {
		return nil, errRequest
	}

	return &resp, nil
}

func (s *OrderServiceOp) RemoveTag(params OrderRemoveTagParams) (*OrderRemoveTagResource, error) {

	// URI
	uri := fmt.Sprintf("%s/%s", apiURI, ordersBasePath)

	url := uri + "/removetag"

	var resp OrderRemoveTagResource

	errRequest := s.client.Request("POST", url, params, &resp)
	if errRequest != nil {
		return nil, errRequest
	}

	return &resp, nil
}

func (s *OrderServiceOp) MarkOrderShipped(params MarkOrderShippedParams) (*Order, error) {

	// URI
	uri := fmt.Sprintf("%s/%s", apiURI, ordersBasePath)

	url := uri + "/markasshipped"

	var resp Order
	errRequest := s.client.Request("POST", url, params, &resp)
	if errRequest != nil {
		return nil, errRequest
	}

	return &resp, nil
}
