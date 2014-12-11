// THIS FILE IS AUTOMATICALLY GENERATED. DO NOT EDIT.

// Package cognitosync provides a client for Amazon Cognito Sync.
package cognitosync

import (
	"encoding/xml"
	"net/http"
	"time"

	"github.com/stripe/aws-go/aws"
	"github.com/stripe/aws-go/aws/gen/endpoints"
)

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/url"
	"strconv"
	"strings"
)

// CognitoSync is a client for Amazon Cognito Sync.
type CognitoSync struct {
	client *aws.RestClient
}

// New returns a new CognitoSync client.
func New(creds aws.Credentials, region string, client *http.Client) *CognitoSync {
	if client == nil {
		client = http.DefaultClient
	}

	service := "cognito-sync"
	endpoint, service, region := endpoints.Lookup("cognito-sync", region)

	return &CognitoSync{
		client: &aws.RestClient{
			Context: aws.Context{
				Credentials: creds,
				Service:     service,
				Region:      region,
			},
			Client:     client,
			Endpoint:   endpoint,
			APIVersion: "2014-06-30",
		},
	}
}

// DeleteDataset deletes the specific dataset. The dataset will be deleted
// permanently, and the action can't be undone. Datasets that this dataset
// was merged with will no longer report the merge. Any consequent
// operation on this dataset will result in a ResourceNotFoundException.
func (c *CognitoSync) DeleteDataset(req *DeleteDatasetRequest) (resp *DeleteDatasetResponse, err error) {
	resp = &DeleteDatasetResponse{}

	var body io.Reader
	var contentType string

	uri := c.client.Endpoint + "/identitypools/{IdentityPoolId}/identities/{IdentityId}/datasets/{DatasetName}"

	if req.DatasetName != nil {
		uri = strings.Replace(uri, "{"+"DatasetName"+"}", *req.DatasetName, -1)
		uri = strings.Replace(uri, "{"+"DatasetName+"+"}", *req.DatasetName, -1)
	}

	if req.IdentityID != nil {
		uri = strings.Replace(uri, "{"+"IdentityId"+"}", *req.IdentityID, -1)
		uri = strings.Replace(uri, "{"+"IdentityId+"+"}", *req.IdentityID, -1)
	}

	if req.IdentityPoolID != nil {
		uri = strings.Replace(uri, "{"+"IdentityPoolId"+"}", *req.IdentityPoolID, -1)
		uri = strings.Replace(uri, "{"+"IdentityPoolId+"+"}", *req.IdentityPoolID, -1)
	}

	q := url.Values{}

	if len(q) > 0 {
		uri += "?" + q.Encode()
	}

	httpReq, err := http.NewRequest("DELETE", uri, body)
	if err != nil {
		return
	}

	if contentType != "" {
		httpReq.Header.Set("Content-Type", contentType)
	}

	httpResp, err := c.client.Do(httpReq)
	if err != nil {
		return
	}

	defer httpResp.Body.Close()
	err = xml.NewDecoder(httpResp.Body).Decode(resp)

	return
}

// DescribeDataset gets metadata about a dataset by identity and dataset
// name. The credentials used to make this API call need to have access to
// the identity data. With Amazon Cognito Sync, each identity has access
// only to its own data. You should use Amazon Cognito Identity service to
// retrieve the credentials necessary to make this API call.
func (c *CognitoSync) DescribeDataset(req *DescribeDatasetRequest) (resp *DescribeDatasetResponse, err error) {
	resp = &DescribeDatasetResponse{}

	var body io.Reader
	var contentType string

	uri := c.client.Endpoint + "/identitypools/{IdentityPoolId}/identities/{IdentityId}/datasets/{DatasetName}"

	if req.DatasetName != nil {
		uri = strings.Replace(uri, "{"+"DatasetName"+"}", *req.DatasetName, -1)
		uri = strings.Replace(uri, "{"+"DatasetName+"+"}", *req.DatasetName, -1)
	}

	if req.IdentityID != nil {
		uri = strings.Replace(uri, "{"+"IdentityId"+"}", *req.IdentityID, -1)
		uri = strings.Replace(uri, "{"+"IdentityId+"+"}", *req.IdentityID, -1)
	}

	if req.IdentityPoolID != nil {
		uri = strings.Replace(uri, "{"+"IdentityPoolId"+"}", *req.IdentityPoolID, -1)
		uri = strings.Replace(uri, "{"+"IdentityPoolId+"+"}", *req.IdentityPoolID, -1)
	}

	q := url.Values{}

	if len(q) > 0 {
		uri += "?" + q.Encode()
	}

	httpReq, err := http.NewRequest("GET", uri, body)
	if err != nil {
		return
	}

	if contentType != "" {
		httpReq.Header.Set("Content-Type", contentType)
	}

	httpResp, err := c.client.Do(httpReq)
	if err != nil {
		return
	}

	defer httpResp.Body.Close()
	err = xml.NewDecoder(httpResp.Body).Decode(resp)

	return
}

// DescribeIdentityPoolUsage gets usage details (for example, data storage)
// about a particular identity pool.
func (c *CognitoSync) DescribeIdentityPoolUsage(req *DescribeIdentityPoolUsageRequest) (resp *DescribeIdentityPoolUsageResponse, err error) {
	resp = &DescribeIdentityPoolUsageResponse{}

	var body io.Reader
	var contentType string

	uri := c.client.Endpoint + "/identitypools/{IdentityPoolId}"

	if req.IdentityPoolID != nil {
		uri = strings.Replace(uri, "{"+"IdentityPoolId"+"}", *req.IdentityPoolID, -1)
		uri = strings.Replace(uri, "{"+"IdentityPoolId+"+"}", *req.IdentityPoolID, -1)
	}

	q := url.Values{}

	if len(q) > 0 {
		uri += "?" + q.Encode()
	}

	httpReq, err := http.NewRequest("GET", uri, body)
	if err != nil {
		return
	}

	if contentType != "" {
		httpReq.Header.Set("Content-Type", contentType)
	}

	httpResp, err := c.client.Do(httpReq)
	if err != nil {
		return
	}

	defer httpResp.Body.Close()
	err = xml.NewDecoder(httpResp.Body).Decode(resp)

	return
}

// DescribeIdentityUsage gets usage information for an identity, including
// number of datasets and data usage.
func (c *CognitoSync) DescribeIdentityUsage(req *DescribeIdentityUsageRequest) (resp *DescribeIdentityUsageResponse, err error) {
	resp = &DescribeIdentityUsageResponse{}

	var body io.Reader
	var contentType string

	uri := c.client.Endpoint + "/identitypools/{IdentityPoolId}/identities/{IdentityId}"

	if req.IdentityID != nil {
		uri = strings.Replace(uri, "{"+"IdentityId"+"}", *req.IdentityID, -1)
		uri = strings.Replace(uri, "{"+"IdentityId+"+"}", *req.IdentityID, -1)
	}

	if req.IdentityPoolID != nil {
		uri = strings.Replace(uri, "{"+"IdentityPoolId"+"}", *req.IdentityPoolID, -1)
		uri = strings.Replace(uri, "{"+"IdentityPoolId+"+"}", *req.IdentityPoolID, -1)
	}

	q := url.Values{}

	if len(q) > 0 {
		uri += "?" + q.Encode()
	}

	httpReq, err := http.NewRequest("GET", uri, body)
	if err != nil {
		return
	}

	if contentType != "" {
		httpReq.Header.Set("Content-Type", contentType)
	}

	httpResp, err := c.client.Do(httpReq)
	if err != nil {
		return
	}

	defer httpResp.Body.Close()
	err = xml.NewDecoder(httpResp.Body).Decode(resp)

	return
}

// GetIdentityPoolConfiguration gets the configuration settings of an
// identity pool.
func (c *CognitoSync) GetIdentityPoolConfiguration(req *GetIdentityPoolConfigurationRequest) (resp *GetIdentityPoolConfigurationResponse, err error) {
	resp = &GetIdentityPoolConfigurationResponse{}

	var body io.Reader
	var contentType string

	uri := c.client.Endpoint + "/identitypools/{IdentityPoolId}/configuration"

	if req.IdentityPoolID != nil {
		uri = strings.Replace(uri, "{"+"IdentityPoolId"+"}", *req.IdentityPoolID, -1)
		uri = strings.Replace(uri, "{"+"IdentityPoolId+"+"}", *req.IdentityPoolID, -1)
	}

	q := url.Values{}

	if len(q) > 0 {
		uri += "?" + q.Encode()
	}

	httpReq, err := http.NewRequest("GET", uri, body)
	if err != nil {
		return
	}

	if contentType != "" {
		httpReq.Header.Set("Content-Type", contentType)
	}

	httpResp, err := c.client.Do(httpReq)
	if err != nil {
		return
	}

	defer httpResp.Body.Close()
	err = xml.NewDecoder(httpResp.Body).Decode(resp)

	return
}

// ListDatasets lists datasets for an identity. The credentials used to
// make this API call need to have access to the identity data. With Amazon
// Cognito Sync, each identity has access only to its own data. You should
// use Amazon Cognito Identity service to retrieve the credentials
// necessary to make this API call.
func (c *CognitoSync) ListDatasets(req *ListDatasetsRequest) (resp *ListDatasetsResponse, err error) {
	resp = &ListDatasetsResponse{}

	var body io.Reader
	var contentType string

	uri := c.client.Endpoint + "/identitypools/{IdentityPoolId}/identities/{IdentityId}/datasets"

	if req.IdentityID != nil {
		uri = strings.Replace(uri, "{"+"IdentityId"+"}", *req.IdentityID, -1)
		uri = strings.Replace(uri, "{"+"IdentityId+"+"}", *req.IdentityID, -1)
	}

	if req.IdentityPoolID != nil {
		uri = strings.Replace(uri, "{"+"IdentityPoolId"+"}", *req.IdentityPoolID, -1)
		uri = strings.Replace(uri, "{"+"IdentityPoolId+"+"}", *req.IdentityPoolID, -1)
	}

	q := url.Values{}

	if req.MaxResults != nil {
		q.Set("maxResults", strconv.Itoa(*req.MaxResults))
	}

	if req.NextToken != nil {
		q.Set("nextToken", *req.NextToken)
	}

	if len(q) > 0 {
		uri += "?" + q.Encode()
	}

	httpReq, err := http.NewRequest("GET", uri, body)
	if err != nil {
		return
	}

	if contentType != "" {
		httpReq.Header.Set("Content-Type", contentType)
	}

	httpResp, err := c.client.Do(httpReq)
	if err != nil {
		return
	}

	defer httpResp.Body.Close()
	err = xml.NewDecoder(httpResp.Body).Decode(resp)

	return
}

// ListIdentityPoolUsage gets a list of identity pools registered with
// Cognito.
func (c *CognitoSync) ListIdentityPoolUsage(req *ListIdentityPoolUsageRequest) (resp *ListIdentityPoolUsageResponse, err error) {
	resp = &ListIdentityPoolUsageResponse{}

	var body io.Reader
	var contentType string

	uri := c.client.Endpoint + "/identitypools"

	q := url.Values{}

	if req.MaxResults != nil {
		q.Set("maxResults", strconv.Itoa(*req.MaxResults))
	}

	if req.NextToken != nil {
		q.Set("nextToken", *req.NextToken)
	}

	if len(q) > 0 {
		uri += "?" + q.Encode()
	}

	httpReq, err := http.NewRequest("GET", uri, body)
	if err != nil {
		return
	}

	if contentType != "" {
		httpReq.Header.Set("Content-Type", contentType)
	}

	httpResp, err := c.client.Do(httpReq)
	if err != nil {
		return
	}

	defer httpResp.Body.Close()
	err = xml.NewDecoder(httpResp.Body).Decode(resp)

	return
}

// ListRecords gets paginated records, optionally changed after a
// particular sync count for a dataset and identity. The credentials used
// to make this API call need to have access to the identity data. With
// Amazon Cognito Sync, each identity has access only to its own data. You
// should use Amazon Cognito Identity service to retrieve the credentials
// necessary to make this API call.
func (c *CognitoSync) ListRecords(req *ListRecordsRequest) (resp *ListRecordsResponse, err error) {
	resp = &ListRecordsResponse{}

	var body io.Reader
	var contentType string

	uri := c.client.Endpoint + "/identitypools/{IdentityPoolId}/identities/{IdentityId}/datasets/{DatasetName}/records"

	if req.DatasetName != nil {
		uri = strings.Replace(uri, "{"+"DatasetName"+"}", *req.DatasetName, -1)
		uri = strings.Replace(uri, "{"+"DatasetName+"+"}", *req.DatasetName, -1)
	}

	if req.IdentityID != nil {
		uri = strings.Replace(uri, "{"+"IdentityId"+"}", *req.IdentityID, -1)
		uri = strings.Replace(uri, "{"+"IdentityId+"+"}", *req.IdentityID, -1)
	}

	if req.IdentityPoolID != nil {
		uri = strings.Replace(uri, "{"+"IdentityPoolId"+"}", *req.IdentityPoolID, -1)
		uri = strings.Replace(uri, "{"+"IdentityPoolId+"+"}", *req.IdentityPoolID, -1)
	}

	q := url.Values{}

	if req.LastSyncCount != nil {
		q.Set("lastSyncCount", fmt.Sprintf("%v", req.LastSyncCount))
	}

	if req.MaxResults != nil {
		q.Set("maxResults", strconv.Itoa(*req.MaxResults))
	}

	if req.NextToken != nil {
		q.Set("nextToken", *req.NextToken)
	}

	if req.SyncSessionToken != nil {
		q.Set("syncSessionToken", *req.SyncSessionToken)
	}

	if len(q) > 0 {
		uri += "?" + q.Encode()
	}

	httpReq, err := http.NewRequest("GET", uri, body)
	if err != nil {
		return
	}

	if contentType != "" {
		httpReq.Header.Set("Content-Type", contentType)
	}

	httpResp, err := c.client.Do(httpReq)
	if err != nil {
		return
	}

	defer httpResp.Body.Close()
	err = xml.NewDecoder(httpResp.Body).Decode(resp)

	return
}

// RegisterDevice registers a device to receive push sync notifications.
func (c *CognitoSync) RegisterDevice(req *RegisterDeviceRequest) (resp *RegisterDeviceResponse, err error) {
	resp = &RegisterDeviceResponse{}

	var body io.Reader
	var contentType string

	uri := c.client.Endpoint + "/identitypools/{IdentityPoolId}/identity/{IdentityId}/device"

	if req.IdentityID != nil {
		uri = strings.Replace(uri, "{"+"IdentityId"+"}", *req.IdentityID, -1)
		uri = strings.Replace(uri, "{"+"IdentityId+"+"}", *req.IdentityID, -1)
	}

	if req.IdentityPoolID != nil {
		uri = strings.Replace(uri, "{"+"IdentityPoolId"+"}", *req.IdentityPoolID, -1)
		uri = strings.Replace(uri, "{"+"IdentityPoolId+"+"}", *req.IdentityPoolID, -1)
	}

	q := url.Values{}

	if len(q) > 0 {
		uri += "?" + q.Encode()
	}

	httpReq, err := http.NewRequest("POST", uri, body)
	if err != nil {
		return
	}

	if contentType != "" {
		httpReq.Header.Set("Content-Type", contentType)
	}

	httpResp, err := c.client.Do(httpReq)
	if err != nil {
		return
	}

	defer httpResp.Body.Close()
	err = xml.NewDecoder(httpResp.Body).Decode(resp)

	return
}

// SetIdentityPoolConfiguration is undocumented.
func (c *CognitoSync) SetIdentityPoolConfiguration(req *SetIdentityPoolConfigurationRequest) (resp *SetIdentityPoolConfigurationResponse, err error) {
	resp = &SetIdentityPoolConfigurationResponse{}

	var body io.Reader
	var contentType string

	uri := c.client.Endpoint + "/identitypools/{IdentityPoolId}/configuration"

	if req.IdentityPoolID != nil {
		uri = strings.Replace(uri, "{"+"IdentityPoolId"+"}", *req.IdentityPoolID, -1)
		uri = strings.Replace(uri, "{"+"IdentityPoolId+"+"}", *req.IdentityPoolID, -1)
	}

	q := url.Values{}

	if len(q) > 0 {
		uri += "?" + q.Encode()
	}

	httpReq, err := http.NewRequest("POST", uri, body)
	if err != nil {
		return
	}

	if contentType != "" {
		httpReq.Header.Set("Content-Type", contentType)
	}

	httpResp, err := c.client.Do(httpReq)
	if err != nil {
		return
	}

	defer httpResp.Body.Close()
	err = xml.NewDecoder(httpResp.Body).Decode(resp)

	return
}

// SubscribeToDataset subscribes to receive notifications when a dataset is
// modified by another device.
func (c *CognitoSync) SubscribeToDataset(req *SubscribeToDatasetRequest) (resp *SubscribeToDatasetResponse, err error) {
	resp = &SubscribeToDatasetResponse{}

	var body io.Reader
	var contentType string

	uri := c.client.Endpoint + "/identitypools/{IdentityPoolId}/identities/{IdentityId}/datasets/{DatasetName}/subscriptions/{DeviceId}"

	if req.DatasetName != nil {
		uri = strings.Replace(uri, "{"+"DatasetName"+"}", *req.DatasetName, -1)
		uri = strings.Replace(uri, "{"+"DatasetName+"+"}", *req.DatasetName, -1)
	}

	if req.DeviceID != nil {
		uri = strings.Replace(uri, "{"+"DeviceId"+"}", *req.DeviceID, -1)
		uri = strings.Replace(uri, "{"+"DeviceId+"+"}", *req.DeviceID, -1)
	}

	if req.IdentityID != nil {
		uri = strings.Replace(uri, "{"+"IdentityId"+"}", *req.IdentityID, -1)
		uri = strings.Replace(uri, "{"+"IdentityId+"+"}", *req.IdentityID, -1)
	}

	if req.IdentityPoolID != nil {
		uri = strings.Replace(uri, "{"+"IdentityPoolId"+"}", *req.IdentityPoolID, -1)
		uri = strings.Replace(uri, "{"+"IdentityPoolId+"+"}", *req.IdentityPoolID, -1)
	}

	q := url.Values{}

	if len(q) > 0 {
		uri += "?" + q.Encode()
	}

	httpReq, err := http.NewRequest("POST", uri, body)
	if err != nil {
		return
	}

	if contentType != "" {
		httpReq.Header.Set("Content-Type", contentType)
	}

	httpResp, err := c.client.Do(httpReq)
	if err != nil {
		return
	}

	defer httpResp.Body.Close()
	err = xml.NewDecoder(httpResp.Body).Decode(resp)

	return
}

// UnsubscribeFromDataset unsubscribe from receiving notifications when a
// dataset is modified by another device.
func (c *CognitoSync) UnsubscribeFromDataset(req *UnsubscribeFromDatasetRequest) (resp *UnsubscribeFromDatasetResponse, err error) {
	resp = &UnsubscribeFromDatasetResponse{}

	var body io.Reader
	var contentType string

	uri := c.client.Endpoint + "/identitypools/{IdentityPoolId}/identities/{IdentityId}/datasets/{DatasetName}/subscriptions/{DeviceId}"

	if req.DatasetName != nil {
		uri = strings.Replace(uri, "{"+"DatasetName"+"}", *req.DatasetName, -1)
		uri = strings.Replace(uri, "{"+"DatasetName+"+"}", *req.DatasetName, -1)
	}

	if req.DeviceID != nil {
		uri = strings.Replace(uri, "{"+"DeviceId"+"}", *req.DeviceID, -1)
		uri = strings.Replace(uri, "{"+"DeviceId+"+"}", *req.DeviceID, -1)
	}

	if req.IdentityID != nil {
		uri = strings.Replace(uri, "{"+"IdentityId"+"}", *req.IdentityID, -1)
		uri = strings.Replace(uri, "{"+"IdentityId+"+"}", *req.IdentityID, -1)
	}

	if req.IdentityPoolID != nil {
		uri = strings.Replace(uri, "{"+"IdentityPoolId"+"}", *req.IdentityPoolID, -1)
		uri = strings.Replace(uri, "{"+"IdentityPoolId+"+"}", *req.IdentityPoolID, -1)
	}

	q := url.Values{}

	if len(q) > 0 {
		uri += "?" + q.Encode()
	}

	httpReq, err := http.NewRequest("DELETE", uri, body)
	if err != nil {
		return
	}

	if contentType != "" {
		httpReq.Header.Set("Content-Type", contentType)
	}

	httpResp, err := c.client.Do(httpReq)
	if err != nil {
		return
	}

	defer httpResp.Body.Close()
	err = xml.NewDecoder(httpResp.Body).Decode(resp)

	return
}

// UpdateRecords posts updates to records and add and delete records for a
// dataset and user. The credentials used to make this API call need to
// have access to the identity data. With Amazon Cognito Sync, each
// identity has access only to its own data. You should use Amazon Cognito
// Identity service to retrieve the credentials necessary to make this API
// call.
func (c *CognitoSync) UpdateRecords(req *UpdateRecordsRequest) (resp *UpdateRecordsResponse, err error) {
	resp = &UpdateRecordsResponse{}

	var body io.Reader
	var contentType string

	uri := c.client.Endpoint + "/identitypools/{IdentityPoolId}/identities/{IdentityId}/datasets/{DatasetName}"

	if req.DatasetName != nil {
		uri = strings.Replace(uri, "{"+"DatasetName"+"}", *req.DatasetName, -1)
		uri = strings.Replace(uri, "{"+"DatasetName+"+"}", *req.DatasetName, -1)
	}

	if req.IdentityID != nil {
		uri = strings.Replace(uri, "{"+"IdentityId"+"}", *req.IdentityID, -1)
		uri = strings.Replace(uri, "{"+"IdentityId+"+"}", *req.IdentityID, -1)
	}

	if req.IdentityPoolID != nil {
		uri = strings.Replace(uri, "{"+"IdentityPoolId"+"}", *req.IdentityPoolID, -1)
		uri = strings.Replace(uri, "{"+"IdentityPoolId+"+"}", *req.IdentityPoolID, -1)
	}

	q := url.Values{}

	if len(q) > 0 {
		uri += "?" + q.Encode()
	}

	httpReq, err := http.NewRequest("POST", uri, body)
	if err != nil {
		return
	}

	if contentType != "" {
		httpReq.Header.Set("Content-Type", contentType)
	}

	if req.ClientContext != nil {
		httpReq.Header.Set("x-amz-Client-Context", *req.ClientContext)
	}

	httpResp, err := c.client.Do(httpReq)
	if err != nil {
		return
	}

	defer httpResp.Body.Close()
	err = xml.NewDecoder(httpResp.Body).Decode(resp)

	return
}

// Dataset is undocumented.
type Dataset struct {
	CreationDate     time.Time       `json:"CreationDate,omitempty"`
	DataStorage      aws.LongValue   `json:"DataStorage,omitempty"`
	DatasetName      aws.StringValue `json:"DatasetName,omitempty"`
	IdentityID       aws.StringValue `json:"IdentityId,omitempty"`
	LastModifiedBy   aws.StringValue `json:"LastModifiedBy,omitempty"`
	LastModifiedDate time.Time       `json:"LastModifiedDate,omitempty"`
	NumRecords       aws.LongValue   `json:"NumRecords,omitempty"`
}

// DeleteDatasetRequest is undocumented.
type DeleteDatasetRequest struct {
	DatasetName    aws.StringValue `json:"DatasetName"`
	IdentityID     aws.StringValue `json:"IdentityId"`
	IdentityPoolID aws.StringValue `json:"IdentityPoolId"`
}

// DeleteDatasetResponse is undocumented.
type DeleteDatasetResponse struct {
	Dataset *Dataset `json:"Dataset,omitempty"`
}

// DescribeDatasetRequest is undocumented.
type DescribeDatasetRequest struct {
	DatasetName    aws.StringValue `json:"DatasetName"`
	IdentityID     aws.StringValue `json:"IdentityId"`
	IdentityPoolID aws.StringValue `json:"IdentityPoolId"`
}

// DescribeDatasetResponse is undocumented.
type DescribeDatasetResponse struct {
	Dataset *Dataset `json:"Dataset,omitempty"`
}

// DescribeIdentityPoolUsageRequest is undocumented.
type DescribeIdentityPoolUsageRequest struct {
	IdentityPoolID aws.StringValue `json:"IdentityPoolId"`
}

// DescribeIdentityPoolUsageResponse is undocumented.
type DescribeIdentityPoolUsageResponse struct {
	IdentityPoolUsage *IdentityPoolUsage `json:"IdentityPoolUsage,omitempty"`
}

// DescribeIdentityUsageRequest is undocumented.
type DescribeIdentityUsageRequest struct {
	IdentityID     aws.StringValue `json:"IdentityId"`
	IdentityPoolID aws.StringValue `json:"IdentityPoolId"`
}

// DescribeIdentityUsageResponse is undocumented.
type DescribeIdentityUsageResponse struct {
	IdentityUsage *IdentityUsage `json:"IdentityUsage,omitempty"`
}

// GetIdentityPoolConfigurationRequest is undocumented.
type GetIdentityPoolConfigurationRequest struct {
	IdentityPoolID aws.StringValue `json:"IdentityPoolId"`
}

// GetIdentityPoolConfigurationResponse is undocumented.
type GetIdentityPoolConfigurationResponse struct {
	IdentityPoolID aws.StringValue `json:"IdentityPoolId,omitempty"`
	PushSync       *PushSync       `json:"PushSync,omitempty"`
}

// IdentityPoolUsage is undocumented.
type IdentityPoolUsage struct {
	DataStorage       aws.LongValue   `json:"DataStorage,omitempty"`
	IdentityPoolID    aws.StringValue `json:"IdentityPoolId,omitempty"`
	LastModifiedDate  time.Time       `json:"LastModifiedDate,omitempty"`
	SyncSessionsCount aws.LongValue   `json:"SyncSessionsCount,omitempty"`
}

// IdentityUsage is undocumented.
type IdentityUsage struct {
	DataStorage      aws.LongValue    `json:"DataStorage,omitempty"`
	DatasetCount     aws.IntegerValue `json:"DatasetCount,omitempty"`
	IdentityID       aws.StringValue  `json:"IdentityId,omitempty"`
	IdentityPoolID   aws.StringValue  `json:"IdentityPoolId,omitempty"`
	LastModifiedDate time.Time        `json:"LastModifiedDate,omitempty"`
}

// ListDatasetsRequest is undocumented.
type ListDatasetsRequest struct {
	IdentityID     aws.StringValue  `json:"IdentityId"`
	IdentityPoolID aws.StringValue  `json:"IdentityPoolId"`
	MaxResults     aws.IntegerValue `json:"MaxResults,omitempty"`
	NextToken      aws.StringValue  `json:"NextToken,omitempty"`
}

// ListDatasetsResponse is undocumented.
type ListDatasetsResponse struct {
	Count     aws.IntegerValue `json:"Count,omitempty"`
	Datasets  []Dataset        `json:"Datasets,omitempty"`
	NextToken aws.StringValue  `json:"NextToken,omitempty"`
}

// ListIdentityPoolUsageRequest is undocumented.
type ListIdentityPoolUsageRequest struct {
	MaxResults aws.IntegerValue `json:"MaxResults,omitempty"`
	NextToken  aws.StringValue  `json:"NextToken,omitempty"`
}

// ListIdentityPoolUsageResponse is undocumented.
type ListIdentityPoolUsageResponse struct {
	Count              aws.IntegerValue    `json:"Count,omitempty"`
	IdentityPoolUsages []IdentityPoolUsage `json:"IdentityPoolUsages,omitempty"`
	MaxResults         aws.IntegerValue    `json:"MaxResults,omitempty"`
	NextToken          aws.StringValue     `json:"NextToken,omitempty"`
}

// ListRecordsRequest is undocumented.
type ListRecordsRequest struct {
	DatasetName      aws.StringValue  `json:"DatasetName"`
	IdentityID       aws.StringValue  `json:"IdentityId"`
	IdentityPoolID   aws.StringValue  `json:"IdentityPoolId"`
	LastSyncCount    aws.LongValue    `json:"LastSyncCount,omitempty"`
	MaxResults       aws.IntegerValue `json:"MaxResults,omitempty"`
	NextToken        aws.StringValue  `json:"NextToken,omitempty"`
	SyncSessionToken aws.StringValue  `json:"SyncSessionToken,omitempty"`
}

// ListRecordsResponse is undocumented.
type ListRecordsResponse struct {
	Count                                 aws.IntegerValue `json:"Count,omitempty"`
	DatasetDeletedAfterRequestedSyncCount aws.BooleanValue `json:"DatasetDeletedAfterRequestedSyncCount,omitempty"`
	DatasetExists                         aws.BooleanValue `json:"DatasetExists,omitempty"`
	DatasetSyncCount                      aws.LongValue    `json:"DatasetSyncCount,omitempty"`
	LastModifiedBy                        aws.StringValue  `json:"LastModifiedBy,omitempty"`
	MergedDatasetNames                    []string         `json:"MergedDatasetNames,omitempty"`
	NextToken                             aws.StringValue  `json:"NextToken,omitempty"`
	Records                               []Record         `json:"Records,omitempty"`
	SyncSessionToken                      aws.StringValue  `json:"SyncSessionToken,omitempty"`
}

// PushSync is undocumented.
type PushSync struct {
	ApplicationARNs []string        `json:"ApplicationArns,omitempty"`
	RoleARN         aws.StringValue `json:"RoleArn,omitempty"`
}

// Record is undocumented.
type Record struct {
	DeviceLastModifiedDate time.Time       `json:"DeviceLastModifiedDate,omitempty"`
	Key                    aws.StringValue `json:"Key,omitempty"`
	LastModifiedBy         aws.StringValue `json:"LastModifiedBy,omitempty"`
	LastModifiedDate       time.Time       `json:"LastModifiedDate,omitempty"`
	SyncCount              aws.LongValue   `json:"SyncCount,omitempty"`
	Value                  aws.StringValue `json:"Value,omitempty"`
}

// RecordPatch is undocumented.
type RecordPatch struct {
	DeviceLastModifiedDate time.Time       `json:"DeviceLastModifiedDate,omitempty"`
	Key                    aws.StringValue `json:"Key"`
	Op                     aws.StringValue `json:"Op"`
	SyncCount              aws.LongValue   `json:"SyncCount"`
	Value                  aws.StringValue `json:"Value,omitempty"`
}

// RegisterDeviceRequest is undocumented.
type RegisterDeviceRequest struct {
	IdentityID     aws.StringValue `json:"IdentityId"`
	IdentityPoolID aws.StringValue `json:"IdentityPoolId"`
	Platform       aws.StringValue `json:"Platform"`
	Token          aws.StringValue `json:"Token"`
}

// RegisterDeviceResponse is undocumented.
type RegisterDeviceResponse struct {
	DeviceID aws.StringValue `json:"DeviceId,omitempty"`
}

// SetIdentityPoolConfigurationRequest is undocumented.
type SetIdentityPoolConfigurationRequest struct {
	IdentityPoolID aws.StringValue `json:"IdentityPoolId"`
	PushSync       *PushSync       `json:"PushSync,omitempty"`
}

// SetIdentityPoolConfigurationResponse is undocumented.
type SetIdentityPoolConfigurationResponse struct {
	IdentityPoolID aws.StringValue `json:"IdentityPoolId,omitempty"`
	PushSync       *PushSync       `json:"PushSync,omitempty"`
}

// SubscribeToDatasetRequest is undocumented.
type SubscribeToDatasetRequest struct {
	DatasetName    aws.StringValue `json:"DatasetName"`
	DeviceID       aws.StringValue `json:"DeviceId"`
	IdentityID     aws.StringValue `json:"IdentityId"`
	IdentityPoolID aws.StringValue `json:"IdentityPoolId"`
}

// SubscribeToDatasetResponse is undocumented.
type SubscribeToDatasetResponse struct {
}

// UnsubscribeFromDatasetRequest is undocumented.
type UnsubscribeFromDatasetRequest struct {
	DatasetName    aws.StringValue `json:"DatasetName"`
	DeviceID       aws.StringValue `json:"DeviceId"`
	IdentityID     aws.StringValue `json:"IdentityId"`
	IdentityPoolID aws.StringValue `json:"IdentityPoolId"`
}

// UnsubscribeFromDatasetResponse is undocumented.
type UnsubscribeFromDatasetResponse struct {
}

// UpdateRecordsRequest is undocumented.
type UpdateRecordsRequest struct {
	ClientContext    aws.StringValue `json:"ClientContext,omitempty"`
	DatasetName      aws.StringValue `json:"DatasetName"`
	DeviceID         aws.StringValue `json:"DeviceId,omitempty"`
	IdentityID       aws.StringValue `json:"IdentityId"`
	IdentityPoolID   aws.StringValue `json:"IdentityPoolId"`
	RecordPatches    []RecordPatch   `json:"RecordPatches,omitempty"`
	SyncSessionToken aws.StringValue `json:"SyncSessionToken"`
}

// UpdateRecordsResponse is undocumented.
type UpdateRecordsResponse struct {
	Records []Record `json:"Records,omitempty"`
}

// avoid errors if the packages aren't referenced
var _ time.Time
var _ xml.Name

var _ bytes.Reader
var _ url.URL
var _ fmt.Stringer
var _ strings.Reader
var _ strconv.NumError
var _ = ioutil.Discard
var _ json.RawMessage
