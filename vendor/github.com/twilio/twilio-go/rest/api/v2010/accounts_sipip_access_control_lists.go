/*
 * Twilio - Api
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

// Optional parameters for the method 'CreateSipIpAccessControlList'
type CreateSipIpAccessControlListParams struct {
	// The unique id of the [Account](https://www.twilio.com/docs/iam/api/account) responsible for this resource.
	PathAccountSid *string `json:"PathAccountSid,omitempty"`
	// A human readable descriptive text that describes the IpAccessControlList, up to 64 characters long.
	FriendlyName *string `json:"FriendlyName,omitempty"`
}

func (params *CreateSipIpAccessControlListParams) SetPathAccountSid(PathAccountSid string) *CreateSipIpAccessControlListParams {
	params.PathAccountSid = &PathAccountSid
	return params
}
func (params *CreateSipIpAccessControlListParams) SetFriendlyName(FriendlyName string) *CreateSipIpAccessControlListParams {
	params.FriendlyName = &FriendlyName
	return params
}

// Create a new IpAccessControlList resource
func (c *ApiService) CreateSipIpAccessControlList(params *CreateSipIpAccessControlListParams) (*ApiV2010SipIpAccessControlList, error) {
	path := "/2010-04-01/Accounts/{AccountSid}/SIP/IpAccessControlLists.json"
	if params != nil && params.PathAccountSid != nil {
		path = strings.Replace(path, "{"+"AccountSid"+"}", *params.PathAccountSid, -1)
	} else {
		path = strings.Replace(path, "{"+"AccountSid"+"}", c.requestHandler.Client.AccountSid(), -1)
	}

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.FriendlyName != nil {
		data.Set("FriendlyName", *params.FriendlyName)
	}

	resp, err := c.requestHandler.Post(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ApiV2010SipIpAccessControlList{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Optional parameters for the method 'DeleteSipIpAccessControlList'
type DeleteSipIpAccessControlListParams struct {
	// The unique id of the [Account](https://www.twilio.com/docs/iam/api/account) responsible for this resource.
	PathAccountSid *string `json:"PathAccountSid,omitempty"`
}

func (params *DeleteSipIpAccessControlListParams) SetPathAccountSid(PathAccountSid string) *DeleteSipIpAccessControlListParams {
	params.PathAccountSid = &PathAccountSid
	return params
}

// Delete an IpAccessControlList from the requested account
func (c *ApiService) DeleteSipIpAccessControlList(Sid string, params *DeleteSipIpAccessControlListParams) error {
	path := "/2010-04-01/Accounts/{AccountSid}/SIP/IpAccessControlLists/{Sid}.json"
	if params != nil && params.PathAccountSid != nil {
		path = strings.Replace(path, "{"+"AccountSid"+"}", *params.PathAccountSid, -1)
	} else {
		path = strings.Replace(path, "{"+"AccountSid"+"}", c.requestHandler.Client.AccountSid(), -1)
	}
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	resp, err := c.requestHandler.Delete(c.baseURL+path, data, headers)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	return nil
}

// Optional parameters for the method 'FetchSipIpAccessControlList'
type FetchSipIpAccessControlListParams struct {
	// The unique id of the [Account](https://www.twilio.com/docs/iam/api/account) responsible for this resource.
	PathAccountSid *string `json:"PathAccountSid,omitempty"`
}

func (params *FetchSipIpAccessControlListParams) SetPathAccountSid(PathAccountSid string) *FetchSipIpAccessControlListParams {
	params.PathAccountSid = &PathAccountSid
	return params
}

// Fetch a specific instance of an IpAccessControlList
func (c *ApiService) FetchSipIpAccessControlList(Sid string, params *FetchSipIpAccessControlListParams) (*ApiV2010SipIpAccessControlList, error) {
	path := "/2010-04-01/Accounts/{AccountSid}/SIP/IpAccessControlLists/{Sid}.json"
	if params != nil && params.PathAccountSid != nil {
		path = strings.Replace(path, "{"+"AccountSid"+"}", *params.PathAccountSid, -1)
	} else {
		path = strings.Replace(path, "{"+"AccountSid"+"}", c.requestHandler.Client.AccountSid(), -1)
	}
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	resp, err := c.requestHandler.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ApiV2010SipIpAccessControlList{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Optional parameters for the method 'ListSipIpAccessControlList'
type ListSipIpAccessControlListParams struct {
	// The unique id of the [Account](https://www.twilio.com/docs/iam/api/account) responsible for this resource.
	PathAccountSid *string `json:"PathAccountSid,omitempty"`
	// How many resources to return in each list page. The default is 50, and the maximum is 1000.
	PageSize *int `json:"PageSize,omitempty"`
	// Max number of records to return.
	Limit *int `json:"limit,omitempty"`
}

func (params *ListSipIpAccessControlListParams) SetPathAccountSid(PathAccountSid string) *ListSipIpAccessControlListParams {
	params.PathAccountSid = &PathAccountSid
	return params
}
func (params *ListSipIpAccessControlListParams) SetPageSize(PageSize int) *ListSipIpAccessControlListParams {
	params.PageSize = &PageSize
	return params
}
func (params *ListSipIpAccessControlListParams) SetLimit(Limit int) *ListSipIpAccessControlListParams {
	params.Limit = &Limit
	return params
}

// Retrieve a single page of SipIpAccessControlList records from the API. Request is executed immediately.
func (c *ApiService) PageSipIpAccessControlList(params *ListSipIpAccessControlListParams, pageToken, pageNumber string) (*ListSipIpAccessControlListResponse, error) {
	path := "/2010-04-01/Accounts/{AccountSid}/SIP/IpAccessControlLists.json"

	if params != nil && params.PathAccountSid != nil {
		path = strings.Replace(path, "{"+"AccountSid"+"}", *params.PathAccountSid, -1)
	} else {
		path = strings.Replace(path, "{"+"AccountSid"+"}", c.requestHandler.Client.AccountSid(), -1)
	}

	data := url.Values{}
	headers := make(map[string]interface{})

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

	ps := &ListSipIpAccessControlListResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Lists SipIpAccessControlList records from the API as a list. Unlike stream, this operation is eager and loads 'limit' records into memory before returning.
func (c *ApiService) ListSipIpAccessControlList(params *ListSipIpAccessControlListParams) ([]ApiV2010SipIpAccessControlList, error) {
	if params == nil {
		params = &ListSipIpAccessControlListParams{}
	}
	params.SetPageSize(client.ReadLimits(params.PageSize, params.Limit))

	response, err := c.PageSipIpAccessControlList(params, "", "")
	if err != nil {
		return nil, err
	}

	curRecord := 0
	var records []ApiV2010SipIpAccessControlList

	for response != nil {
		records = append(records, response.IpAccessControlLists...)

		var record interface{}
		if record, err = client.GetNext(c.baseURL, response, &curRecord, params.Limit, c.getNextListSipIpAccessControlListResponse); record == nil || err != nil {
			return records, err
		}

		response = record.(*ListSipIpAccessControlListResponse)
	}

	return records, err
}

// Streams SipIpAccessControlList records from the API as a channel stream. This operation lazily loads records as efficiently as possible until the limit is reached.
func (c *ApiService) StreamSipIpAccessControlList(params *ListSipIpAccessControlListParams) (chan ApiV2010SipIpAccessControlList, error) {
	if params == nil {
		params = &ListSipIpAccessControlListParams{}
	}
	params.SetPageSize(client.ReadLimits(params.PageSize, params.Limit))

	response, err := c.PageSipIpAccessControlList(params, "", "")
	if err != nil {
		return nil, err
	}

	curRecord := 0
	//set buffer size of the channel to 1
	channel := make(chan ApiV2010SipIpAccessControlList, 1)

	go func() {
		for response != nil {
			for item := range response.IpAccessControlLists {
				channel <- response.IpAccessControlLists[item]
			}

			var record interface{}
			if record, err = client.GetNext(c.baseURL, response, &curRecord, params.Limit, c.getNextListSipIpAccessControlListResponse); record == nil || err != nil {
				close(channel)
				return
			}

			response = record.(*ListSipIpAccessControlListResponse)
		}
		close(channel)
	}()

	return channel, err
}

func (c *ApiService) getNextListSipIpAccessControlListResponse(nextPageUrl string) (interface{}, error) {
	if nextPageUrl == "" {
		return nil, nil
	}
	resp, err := c.requestHandler.Get(nextPageUrl, nil, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ListSipIpAccessControlListResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}
	return ps, nil
}

// Optional parameters for the method 'UpdateSipIpAccessControlList'
type UpdateSipIpAccessControlListParams struct {
	// The unique id of the [Account](https://www.twilio.com/docs/iam/api/account) responsible for this resource.
	PathAccountSid *string `json:"PathAccountSid,omitempty"`
	// A human readable descriptive text, up to 64 characters long.
	FriendlyName *string `json:"FriendlyName,omitempty"`
}

func (params *UpdateSipIpAccessControlListParams) SetPathAccountSid(PathAccountSid string) *UpdateSipIpAccessControlListParams {
	params.PathAccountSid = &PathAccountSid
	return params
}
func (params *UpdateSipIpAccessControlListParams) SetFriendlyName(FriendlyName string) *UpdateSipIpAccessControlListParams {
	params.FriendlyName = &FriendlyName
	return params
}

// Rename an IpAccessControlList
func (c *ApiService) UpdateSipIpAccessControlList(Sid string, params *UpdateSipIpAccessControlListParams) (*ApiV2010SipIpAccessControlList, error) {
	path := "/2010-04-01/Accounts/{AccountSid}/SIP/IpAccessControlLists/{Sid}.json"
	if params != nil && params.PathAccountSid != nil {
		path = strings.Replace(path, "{"+"AccountSid"+"}", *params.PathAccountSid, -1)
	} else {
		path = strings.Replace(path, "{"+"AccountSid"+"}", c.requestHandler.Client.AccountSid(), -1)
	}
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.FriendlyName != nil {
		data.Set("FriendlyName", *params.FriendlyName)
	}

	resp, err := c.requestHandler.Post(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ApiV2010SipIpAccessControlList{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}
