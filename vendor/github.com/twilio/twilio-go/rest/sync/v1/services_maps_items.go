/*
 * Twilio - Sync
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.27.1
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strings"

	"github.com/twilio/twilio-go/client"
)

// Optional parameters for the method 'CreateSyncMapItem'
type CreateSyncMapItemParams struct {
	// How long, [in seconds](https://www.twilio.com/docs/sync/limits#sync-payload-limits), before the Map Item's parent Sync Map expires (time-to-live) and is deleted.
	CollectionTtl *int `json:"CollectionTtl,omitempty"`
	// A JSON string that represents an arbitrary, schema-less object that the Map Item stores. Can be up to 16 KiB in length.
	Data *map[string]interface{} `json:"Data,omitempty"`
	// How long, [in seconds](https://www.twilio.com/docs/sync/limits#sync-payload-limits), before the Map Item expires (time-to-live) and is deleted.
	ItemTtl *int `json:"ItemTtl,omitempty"`
	// The unique, user-defined key for the Map Item. Can be up to 320 characters long.
	Key *string `json:"Key,omitempty"`
	// An alias for `item_ttl`. If both parameters are provided, this value is ignored.
	Ttl *int `json:"Ttl,omitempty"`
}

func (params *CreateSyncMapItemParams) SetCollectionTtl(CollectionTtl int) *CreateSyncMapItemParams {
	params.CollectionTtl = &CollectionTtl
	return params
}
func (params *CreateSyncMapItemParams) SetData(Data map[string]interface{}) *CreateSyncMapItemParams {
	params.Data = &Data
	return params
}
func (params *CreateSyncMapItemParams) SetItemTtl(ItemTtl int) *CreateSyncMapItemParams {
	params.ItemTtl = &ItemTtl
	return params
}
func (params *CreateSyncMapItemParams) SetKey(Key string) *CreateSyncMapItemParams {
	params.Key = &Key
	return params
}
func (params *CreateSyncMapItemParams) SetTtl(Ttl int) *CreateSyncMapItemParams {
	params.Ttl = &Ttl
	return params
}

func (c *ApiService) CreateSyncMapItem(ServiceSid string, MapSid string, params *CreateSyncMapItemParams) (*SyncV1SyncMapItem, error) {
	path := "/v1/Services/{ServiceSid}/Maps/{MapSid}/Items"
	path = strings.Replace(path, "{"+"ServiceSid"+"}", ServiceSid, -1)
	path = strings.Replace(path, "{"+"MapSid"+"}", MapSid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.CollectionTtl != nil {
		data.Set("CollectionTtl", fmt.Sprint(*params.CollectionTtl))
	}
	if params != nil && params.Data != nil {
		v, err := json.Marshal(params.Data)

		if err != nil {
			return nil, err
		}

		data.Set("Data", string(v))
	}
	if params != nil && params.ItemTtl != nil {
		data.Set("ItemTtl", fmt.Sprint(*params.ItemTtl))
	}
	if params != nil && params.Key != nil {
		data.Set("Key", *params.Key)
	}
	if params != nil && params.Ttl != nil {
		data.Set("Ttl", fmt.Sprint(*params.Ttl))
	}

	resp, err := c.requestHandler.Post(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &SyncV1SyncMapItem{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Optional parameters for the method 'DeleteSyncMapItem'
type DeleteSyncMapItemParams struct {
	// If provided, applies this mutation if (and only if) the “revision” field of this [map item] matches the provided value. This matches the semantics of (and is implemented with) the HTTP [If-Match header](https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/If-Match).
	IfMatch *string `json:"If-Match,omitempty"`
}

func (params *DeleteSyncMapItemParams) SetIfMatch(IfMatch string) *DeleteSyncMapItemParams {
	params.IfMatch = &IfMatch
	return params
}

func (c *ApiService) DeleteSyncMapItem(ServiceSid string, MapSid string, Key string, params *DeleteSyncMapItemParams) error {
	path := "/v1/Services/{ServiceSid}/Maps/{MapSid}/Items/{Key}"
	path = strings.Replace(path, "{"+"ServiceSid"+"}", ServiceSid, -1)
	path = strings.Replace(path, "{"+"MapSid"+"}", MapSid, -1)
	path = strings.Replace(path, "{"+"Key"+"}", Key, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.IfMatch != nil {
		headers["If-Match"] = *params.IfMatch
	}

	resp, err := c.requestHandler.Delete(c.baseURL+path, data, headers)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	return nil
}

func (c *ApiService) FetchSyncMapItem(ServiceSid string, MapSid string, Key string) (*SyncV1SyncMapItem, error) {
	path := "/v1/Services/{ServiceSid}/Maps/{MapSid}/Items/{Key}"
	path = strings.Replace(path, "{"+"ServiceSid"+"}", ServiceSid, -1)
	path = strings.Replace(path, "{"+"MapSid"+"}", MapSid, -1)
	path = strings.Replace(path, "{"+"Key"+"}", Key, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	resp, err := c.requestHandler.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &SyncV1SyncMapItem{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Optional parameters for the method 'ListSyncMapItem'
type ListSyncMapItemParams struct {
	// How to order the Map Items returned by their `key` value. Can be: `asc` (ascending) or `desc` (descending) and the default is ascending. Map Items are [ordered lexicographically](https://en.wikipedia.org/wiki/Lexicographical_order) by Item key.
	Order *string `json:"Order,omitempty"`
	// The `key` of the first Sync Map Item resource to read. See also `bounds`.
	From *string `json:"From,omitempty"`
	// Whether to include the Map Item referenced by the `from` parameter. Can be: `inclusive` to include the Map Item referenced by the `from` parameter or `exclusive` to start with the next Map Item. The default value is `inclusive`.
	Bounds *string `json:"Bounds,omitempty"`
	// How many resources to return in each list page. The default is 50, and the maximum is 1000.
	PageSize *int `json:"PageSize,omitempty"`
	// Max number of records to return.
	Limit *int `json:"limit,omitempty"`
}

func (params *ListSyncMapItemParams) SetOrder(Order string) *ListSyncMapItemParams {
	params.Order = &Order
	return params
}
func (params *ListSyncMapItemParams) SetFrom(From string) *ListSyncMapItemParams {
	params.From = &From
	return params
}
func (params *ListSyncMapItemParams) SetBounds(Bounds string) *ListSyncMapItemParams {
	params.Bounds = &Bounds
	return params
}
func (params *ListSyncMapItemParams) SetPageSize(PageSize int) *ListSyncMapItemParams {
	params.PageSize = &PageSize
	return params
}
func (params *ListSyncMapItemParams) SetLimit(Limit int) *ListSyncMapItemParams {
	params.Limit = &Limit
	return params
}

// Retrieve a single page of SyncMapItem records from the API. Request is executed immediately.
func (c *ApiService) PageSyncMapItem(ServiceSid string, MapSid string, params *ListSyncMapItemParams, pageToken, pageNumber string) (*ListSyncMapItemResponse, error) {
	path := "/v1/Services/{ServiceSid}/Maps/{MapSid}/Items"

	path = strings.Replace(path, "{"+"ServiceSid"+"}", ServiceSid, -1)
	path = strings.Replace(path, "{"+"MapSid"+"}", MapSid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.Order != nil {
		data.Set("Order", *params.Order)
	}
	if params != nil && params.From != nil {
		data.Set("From", *params.From)
	}
	if params != nil && params.Bounds != nil {
		data.Set("Bounds", *params.Bounds)
	}
	if params != nil && params.PageSize != nil {
		data.Set("PageSize", fmt.Sprint(*params.PageSize))
	}

	if pageToken != "" {
		data.Set("PageToken", pageToken)
	}
	if pageNumber != "" {
		data.Set("Page", pageNumber)
	}

	resp, err := c.requestHandler.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ListSyncMapItemResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Lists SyncMapItem records from the API as a list. Unlike stream, this operation is eager and loads 'limit' records into memory before returning.
func (c *ApiService) ListSyncMapItem(ServiceSid string, MapSid string, params *ListSyncMapItemParams) ([]SyncV1SyncMapItem, error) {
	if params == nil {
		params = &ListSyncMapItemParams{}
	}
	params.SetPageSize(client.ReadLimits(params.PageSize, params.Limit))

	response, err := c.PageSyncMapItem(ServiceSid, MapSid, params, "", "")
	if err != nil {
		return nil, err
	}

	curRecord := 0
	var records []SyncV1SyncMapItem

	for response != nil {
		records = append(records, response.Items...)

		var record interface{}
		if record, err = client.GetNext(c.baseURL, response, &curRecord, params.Limit, c.getNextListSyncMapItemResponse); record == nil || err != nil {
			return records, err
		}

		response = record.(*ListSyncMapItemResponse)
	}

	return records, err
}

// Streams SyncMapItem records from the API as a channel stream. This operation lazily loads records as efficiently as possible until the limit is reached.
func (c *ApiService) StreamSyncMapItem(ServiceSid string, MapSid string, params *ListSyncMapItemParams) (chan SyncV1SyncMapItem, error) {
	if params == nil {
		params = &ListSyncMapItemParams{}
	}
	params.SetPageSize(client.ReadLimits(params.PageSize, params.Limit))

	response, err := c.PageSyncMapItem(ServiceSid, MapSid, params, "", "")
	if err != nil {
		return nil, err
	}

	curRecord := 0
	//set buffer size of the channel to 1
	channel := make(chan SyncV1SyncMapItem, 1)

	go func() {
		for response != nil {
			for item := range response.Items {
				channel <- response.Items[item]
			}

			var record interface{}
			if record, err = client.GetNext(c.baseURL, response, &curRecord, params.Limit, c.getNextListSyncMapItemResponse); record == nil || err != nil {
				close(channel)
				return
			}

			response = record.(*ListSyncMapItemResponse)
		}
		close(channel)
	}()

	return channel, err
}

func (c *ApiService) getNextListSyncMapItemResponse(nextPageUrl string) (interface{}, error) {
	if nextPageUrl == "" {
		return nil, nil
	}
	resp, err := c.requestHandler.Get(nextPageUrl, nil, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ListSyncMapItemResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}
	return ps, nil
}

// Optional parameters for the method 'UpdateSyncMapItem'
type UpdateSyncMapItemParams struct {
	// If provided, applies this mutation if (and only if) the “revision” field of this [map item] matches the provided value. This matches the semantics of (and is implemented with) the HTTP [If-Match header](https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/If-Match).
	IfMatch *string `json:"If-Match,omitempty"`
	// How long, [in seconds](https://www.twilio.com/docs/sync/limits#sync-payload-limits), before the Map Item's parent Sync Map expires (time-to-live) and is deleted. This parameter can only be used when the Map Item's `data` or `ttl` is updated in the same request.
	CollectionTtl *int `json:"CollectionTtl,omitempty"`
	// A JSON string that represents an arbitrary, schema-less object that the Map Item stores. Can be up to 16 KiB in length.
	Data *map[string]interface{} `json:"Data,omitempty"`
	// How long, [in seconds](https://www.twilio.com/docs/sync/limits#sync-payload-limits), before the Map Item expires (time-to-live) and is deleted.
	ItemTtl *int `json:"ItemTtl,omitempty"`
	// An alias for `item_ttl`. If both parameters are provided, this value is ignored.
	Ttl *int `json:"Ttl,omitempty"`
}

func (params *UpdateSyncMapItemParams) SetIfMatch(IfMatch string) *UpdateSyncMapItemParams {
	params.IfMatch = &IfMatch
	return params
}
func (params *UpdateSyncMapItemParams) SetCollectionTtl(CollectionTtl int) *UpdateSyncMapItemParams {
	params.CollectionTtl = &CollectionTtl
	return params
}
func (params *UpdateSyncMapItemParams) SetData(Data map[string]interface{}) *UpdateSyncMapItemParams {
	params.Data = &Data
	return params
}
func (params *UpdateSyncMapItemParams) SetItemTtl(ItemTtl int) *UpdateSyncMapItemParams {
	params.ItemTtl = &ItemTtl
	return params
}
func (params *UpdateSyncMapItemParams) SetTtl(Ttl int) *UpdateSyncMapItemParams {
	params.Ttl = &Ttl
	return params
}

func (c *ApiService) UpdateSyncMapItem(ServiceSid string, MapSid string, Key string, params *UpdateSyncMapItemParams) (*SyncV1SyncMapItem, error) {
	path := "/v1/Services/{ServiceSid}/Maps/{MapSid}/Items/{Key}"
	path = strings.Replace(path, "{"+"ServiceSid"+"}", ServiceSid, -1)
	path = strings.Replace(path, "{"+"MapSid"+"}", MapSid, -1)
	path = strings.Replace(path, "{"+"Key"+"}", Key, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.CollectionTtl != nil {
		data.Set("CollectionTtl", fmt.Sprint(*params.CollectionTtl))
	}
	if params != nil && params.Data != nil {
		v, err := json.Marshal(params.Data)

		if err != nil {
			return nil, err
		}

		data.Set("Data", string(v))
	}
	if params != nil && params.ItemTtl != nil {
		data.Set("ItemTtl", fmt.Sprint(*params.ItemTtl))
	}
	if params != nil && params.Ttl != nil {
		data.Set("Ttl", fmt.Sprint(*params.Ttl))
	}

	if params != nil && params.IfMatch != nil {
		headers["If-Match"] = *params.IfMatch
	}

	resp, err := c.requestHandler.Post(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &SyncV1SyncMapItem{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}
