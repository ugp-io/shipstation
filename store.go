package shipstation

import (
	"fmt"
	"net/url"
	"strconv"
)

const storesBasePath = "stores"

type StoreServiceOp struct {
	client *Client
}

type StoreService interface {
	List(StoreListParams) (*[]Store, error)
	Refresh(params RefreshStoreParams) (bool, error)
}

type StoreListParams struct {
	ShowInactive  *bool
	MarketplaceID *int
}

type RefreshStoreParams struct {
	StoreId     *string `json:"storeId,omitempty"`
	RefreshDate *string `json:"refreshDate,omitempty"`
}

type Store struct {
	StoreID            *int    `json:"storeId,omitempty"`
	StoreName          *string `json:"storeName,omitempty"`
	MarketplaceID      *int    `json:"marketplaceId,omitempty"`
	MarketplaceName    *string `json:"marketplaceName,omitempty"`
	AccountName        *string `json:"accountName,omitempty"`
	Email              *string `json:"email,omitempty"`
	IntegrationURL     *string `json:"integrationUrl,omitempty"`
	Active             *bool   `json:"active,omitempty"`
	CompstringName     *string `json:"compstringName,omitempty"`
	Phone              *string `json:"phone,omitempty"`
	PublicEmail        *string `json:"publicEmail,omitempty"`
	Website            *string `json:"website,omitempty"`
	RefreshDate        *string `json:"refreshDate,omitempty"`
	LastRefreshAttempt *string `json:"lastRefreshAttempt,omitempty"`
	CreateDate         *string `json:"createDate,omitempty"`
	ModifyDate         *string `json:"modifyDate,omitempty"`
	AutoRefresh        *bool   `json:"autoRefresh,omitempty"`
}

func (s *StoreServiceOp) List(params StoreListParams) (*[]Store, error) {

	// URI
	uri := fmt.Sprintf("%s/%s", apiURI, storesBasePath)

	values := url.Values{}
	if params.ShowInactive != nil && *params.ShowInactive {
		values.Add("showInactive", "showInactive")
	}
	if params.MarketplaceID != nil {
		values.Add("marketplaceId", strconv.Itoa(*params.MarketplaceID))
	}

	url := uri + "?" + values.Encode()

	var resp []Store

	errRequest := s.client.Request("GET", url, nil, &resp)
	if errRequest != nil {
		return nil, errRequest
	}

	return &resp, nil
}

func (s *StoreServiceOp) Refresh(params RefreshStoreParams) (bool, error) {

	// URI
	uri := fmt.Sprintf("%s/%s/%s", apiURI, storesBasePath, "refreshstore")

	values := url.Values{}
	if params.StoreId != nil {
		values.Add("storeId", *params.StoreId)
	}
	if params.RefreshDate != nil {
		values.Add("refreshDate", *params.RefreshDate)
	}

	url := uri + "?" + values.Encode()

	var resp DefaultResponse
	errRequest := s.client.Request("POST", url, nil, &resp)
	if errRequest != nil {
		return false, errRequest
	}

	return resp.Success, nil
}
