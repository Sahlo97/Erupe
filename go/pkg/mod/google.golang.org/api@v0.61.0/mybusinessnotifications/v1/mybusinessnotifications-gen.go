// Copyright 2021 Google LLC.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Code generated file. DO NOT EDIT.

// Package mybusinessnotifications provides access to the My Business Notifications API.
//
// For product documentation, see: https://developers.google.com/my-business/
//
// Creating a client
//
// Usage example:
//
//   import "google.golang.org/api/mybusinessnotifications/v1"
//   ...
//   ctx := context.Background()
//   mybusinessnotificationsService, err := mybusinessnotifications.NewService(ctx)
//
// In this example, Google Application Default Credentials are used for authentication.
//
// For information on how to create and obtain Application Default Credentials, see https://developers.google.com/identity/protocols/application-default-credentials.
//
// Other authentication options
//
// To use an API key for authentication (note: some APIs do not support API keys), use option.WithAPIKey:
//
//   mybusinessnotificationsService, err := mybusinessnotifications.NewService(ctx, option.WithAPIKey("AIza..."))
//
// To use an OAuth token (e.g., a user token obtained via a three-legged OAuth flow), use option.WithTokenSource:
//
//   config := &oauth2.Config{...}
//   // ...
//   token, err := config.Exchange(ctx, ...)
//   mybusinessnotificationsService, err := mybusinessnotifications.NewService(ctx, option.WithTokenSource(config.TokenSource(ctx, token)))
//
// See https://godoc.org/google.golang.org/api/option/ for details on options.
package mybusinessnotifications // import "google.golang.org/api/mybusinessnotifications/v1"

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	googleapi "google.golang.org/api/googleapi"
	gensupport "google.golang.org/api/internal/gensupport"
	option "google.golang.org/api/option"
	internaloption "google.golang.org/api/option/internaloption"
	htransport "google.golang.org/api/transport/http"
)

// Always reference these packages, just in case the auto-generated code
// below doesn't.
var _ = bytes.NewBuffer
var _ = strconv.Itoa
var _ = fmt.Sprintf
var _ = json.NewDecoder
var _ = io.Copy
var _ = url.Parse
var _ = gensupport.MarshalJSON
var _ = googleapi.Version
var _ = errors.New
var _ = strings.Replace
var _ = context.Canceled
var _ = internaloption.WithDefaultEndpoint

const apiId = "mybusinessnotifications:v1"
const apiName = "mybusinessnotifications"
const apiVersion = "v1"
const basePath = "https://mybusinessnotifications.googleapis.com/"
const mtlsBasePath = "https://mybusinessnotifications.mtls.googleapis.com/"

// NewService creates a new Service.
func NewService(ctx context.Context, opts ...option.ClientOption) (*Service, error) {
	opts = append(opts, internaloption.WithDefaultEndpoint(basePath))
	opts = append(opts, internaloption.WithDefaultMTLSEndpoint(mtlsBasePath))
	client, endpoint, err := htransport.NewClient(ctx, opts...)
	if err != nil {
		return nil, err
	}
	s, err := New(client)
	if err != nil {
		return nil, err
	}
	if endpoint != "" {
		s.BasePath = endpoint
	}
	return s, nil
}

// New creates a new Service. It uses the provided http.Client for requests.
//
// Deprecated: please use NewService instead.
// To provide a custom HTTP client, use option.WithHTTPClient.
// If you are using google.golang.org/api/googleapis/transport.APIKey, use option.WithAPIKey with NewService instead.
func New(client *http.Client) (*Service, error) {
	if client == nil {
		return nil, errors.New("client is nil")
	}
	s := &Service{client: client, BasePath: basePath}
	s.Accounts = NewAccountsService(s)
	return s, nil
}

type Service struct {
	client    *http.Client
	BasePath  string // API endpoint base URL
	UserAgent string // optional additional User-Agent fragment

	Accounts *AccountsService
}

func (s *Service) userAgent() string {
	if s.UserAgent == "" {
		return googleapi.UserAgent
	}
	return googleapi.UserAgent + " " + s.UserAgent
}

func NewAccountsService(s *Service) *AccountsService {
	rs := &AccountsService{s: s}
	return rs
}

type AccountsService struct {
	s *Service
}

// NotificationSetting: A Google Pub/Sub topic where notifications can
// be published when a location is updated or has a new review. There
// will be only one notification setting resource per-account.
type NotificationSetting struct {
	// Name: Required. The resource name this setting is for. This is of the
	// form `accounts/{account_id}/notificationSetting`.
	Name string `json:"name,omitempty"`

	// NotificationTypes: The types of notifications that will be sent to
	// the Pub/Sub topic. To stop receiving notifications entirely, use
	// NotificationSettings.UpdateNotificationSetting with an empty
	// notification_types or set the pubsub_topic to an empty string.
	//
	// Possible values:
	//   "NOTIFICATION_TYPE_UNSPECIFIED" - No notification type. Will not
	// match any notifications.
	//   "GOOGLE_UPDATE" - The location has Google updates for review. The
	// location_name field on the notification will provide the resource
	// name of the location with Google updates.
	//   "NEW_REVIEW" - A new review has been added to the location. The
	// review_name field on the notification will provide the resource name
	// of the review that was added, and location_name will have the
	// location's resource name.
	//   "UPDATED_REVIEW" - A review on the location has been updated. The
	// review_name field on the notification will provide the resource name
	// of the review that was added, and location_name will have the
	// location's resource name.
	//   "NEW_CUSTOMER_MEDIA" - A new media item has been added to the
	// location by a Google Maps user. The notification will provide the
	// resource name of the new media item.
	//   "NEW_QUESTION" - A new question is added to the location. The
	// notification will provide the resource name of question.
	//   "UPDATED_QUESTION" - A question of the location is updated. The
	// notification will provide the resource name of question.
	//   "NEW_ANSWER" - A new answer is added to the location. The
	// notification will provide the resource name of question and answer.
	//   "UPDATED_ANSWER" - An answer of the location is updated. The
	// notification will provide the resource name of question and answer.
	//   "DUPLICATE_LOCATION" - Indicates whether there is a change in
	// location metadata's duplicate location field.
	//   "LOSS_OF_VOICE_OF_MERCHANT" - Indicates whether the location has a
	// loss in voice of merchant status. Call GetVoiceOfMerchantState rpc
	// for more details
	NotificationTypes []string `json:"notificationTypes,omitempty"`

	// PubsubTopic: Optional. The Google Pub/Sub topic that will receive
	// notifications when locations managed by this account are updated. If
	// unset, no notifications will be posted. The account
	// mybusiness-api-pubsub@system.gserviceaccount.com must have at least
	// Publish permissions on the Pub/Sub topic.
	PubsubTopic string `json:"pubsubTopic,omitempty"`

	// ServerResponse contains the HTTP response code and headers from the
	// server.
	googleapi.ServerResponse `json:"-"`

	// ForceSendFields is a list of field names (e.g. "Name") to
	// unconditionally include in API requests. By default, fields with
	// empty or default values are omitted from API requests. However, any
	// non-pointer, non-interface field appearing in ForceSendFields will be
	// sent to the server regardless of whether the field is empty or not.
	// This may be used to include empty fields in Patch requests.
	ForceSendFields []string `json:"-"`

	// NullFields is a list of field names (e.g. "Name") to include in API
	// requests with the JSON null value. By default, fields with empty
	// values are omitted from API requests. However, any field with an
	// empty value appearing in NullFields will be sent to the server as
	// null. It is an error if a field in this list has a non-empty value.
	// This may be used to include null fields in Patch requests.
	NullFields []string `json:"-"`
}

func (s *NotificationSetting) MarshalJSON() ([]byte, error) {
	type NoMethod NotificationSetting
	raw := NoMethod(*s)
	return gensupport.MarshalJSON(raw, s.ForceSendFields, s.NullFields)
}

// method id "mybusinessnotifications.accounts.getNotificationSetting":

type AccountsGetNotificationSettingCall struct {
	s            *Service
	name         string
	urlParams_   gensupport.URLParams
	ifNoneMatch_ string
	ctx_         context.Context
	header_      http.Header
}

// GetNotificationSetting: Returns the pubsub notification settings for
// the account.
//
// - name: The resource name of the notification setting we are trying
//   to fetch.
func (r *AccountsService) GetNotificationSetting(name string) *AccountsGetNotificationSettingCall {
	c := &AccountsGetNotificationSettingCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.name = name
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *AccountsGetNotificationSettingCall) Fields(s ...googleapi.Field) *AccountsGetNotificationSettingCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// IfNoneMatch sets the optional parameter which makes the operation
// fail if the object's ETag matches the given value. This is useful for
// getting updates only after the object has changed since the last
// request. Use googleapi.IsNotModified to check whether the response
// error from Do is the result of In-None-Match.
func (c *AccountsGetNotificationSettingCall) IfNoneMatch(entityTag string) *AccountsGetNotificationSettingCall {
	c.ifNoneMatch_ = entityTag
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *AccountsGetNotificationSettingCall) Context(ctx context.Context) *AccountsGetNotificationSettingCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *AccountsGetNotificationSettingCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *AccountsGetNotificationSettingCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	reqHeaders.Set("x-goog-api-client", "gl-go/"+gensupport.GoVersion()+" gdcl/20211201")
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	if c.ifNoneMatch_ != "" {
		reqHeaders.Set("If-None-Match", c.ifNoneMatch_)
	}
	var body io.Reader = nil
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1/{+name}")
	urls += "?" + c.urlParams_.Encode()
	req, err := http.NewRequest("GET", urls, body)
	if err != nil {
		return nil, err
	}
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"name": c.name,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "mybusinessnotifications.accounts.getNotificationSetting" call.
// Exactly one of *NotificationSetting or error will be non-nil. Any
// non-2xx status code is an error. Response headers are in either
// *NotificationSetting.ServerResponse.Header or (if a response was
// returned at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *AccountsGetNotificationSettingCall) Do(opts ...googleapi.CallOption) (*NotificationSetting, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &NotificationSetting{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Returns the pubsub notification settings for the account.",
	//   "flatPath": "v1/accounts/{accountsId}/notificationSetting",
	//   "httpMethod": "GET",
	//   "id": "mybusinessnotifications.accounts.getNotificationSetting",
	//   "parameterOrder": [
	//     "name"
	//   ],
	//   "parameters": {
	//     "name": {
	//       "description": "Required. The resource name of the notification setting we are trying to fetch.",
	//       "location": "path",
	//       "pattern": "^accounts/[^/]+/notificationSetting$",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1/{+name}",
	//   "response": {
	//     "$ref": "NotificationSetting"
	//   }
	// }

}

// method id "mybusinessnotifications.accounts.updateNotificationSetting":

type AccountsUpdateNotificationSettingCall struct {
	s                   *Service
	name                string
	notificationsetting *NotificationSetting
	urlParams_          gensupport.URLParams
	ctx_                context.Context
	header_             http.Header
}

// UpdateNotificationSetting: Sets the pubsub notification setting for
// the account informing Google which topic to send pubsub notifications
// for. Use the notification_types field within notification_setting to
// manipulate the events an account wants to subscribe to. An account
// will only have one notification setting resource, and only one pubsub
// topic can be set. To delete the setting, update with an empty
// notification_types
//
// - name: The resource name this setting is for. This is of the form
//   `accounts/{account_id}/notificationSetting`.
func (r *AccountsService) UpdateNotificationSetting(name string, notificationsetting *NotificationSetting) *AccountsUpdateNotificationSettingCall {
	c := &AccountsUpdateNotificationSettingCall{s: r.s, urlParams_: make(gensupport.URLParams)}
	c.name = name
	c.notificationsetting = notificationsetting
	return c
}

// UpdateMask sets the optional parameter "updateMask": Required. The
// specific fields that should be updated. The only editable field is
// notification_setting.
func (c *AccountsUpdateNotificationSettingCall) UpdateMask(updateMask string) *AccountsUpdateNotificationSettingCall {
	c.urlParams_.Set("updateMask", updateMask)
	return c
}

// Fields allows partial responses to be retrieved. See
// https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *AccountsUpdateNotificationSettingCall) Fields(s ...googleapi.Field) *AccountsUpdateNotificationSettingCall {
	c.urlParams_.Set("fields", googleapi.CombineFields(s))
	return c
}

// Context sets the context to be used in this call's Do method. Any
// pending HTTP request will be aborted if the provided context is
// canceled.
func (c *AccountsUpdateNotificationSettingCall) Context(ctx context.Context) *AccountsUpdateNotificationSettingCall {
	c.ctx_ = ctx
	return c
}

// Header returns an http.Header that can be modified by the caller to
// add HTTP headers to the request.
func (c *AccountsUpdateNotificationSettingCall) Header() http.Header {
	if c.header_ == nil {
		c.header_ = make(http.Header)
	}
	return c.header_
}

func (c *AccountsUpdateNotificationSettingCall) doRequest(alt string) (*http.Response, error) {
	reqHeaders := make(http.Header)
	reqHeaders.Set("x-goog-api-client", "gl-go/"+gensupport.GoVersion()+" gdcl/20211201")
	for k, v := range c.header_ {
		reqHeaders[k] = v
	}
	reqHeaders.Set("User-Agent", c.s.userAgent())
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.notificationsetting)
	if err != nil {
		return nil, err
	}
	reqHeaders.Set("Content-Type", "application/json")
	c.urlParams_.Set("alt", alt)
	c.urlParams_.Set("prettyPrint", "false")
	urls := googleapi.ResolveRelative(c.s.BasePath, "v1/{+name}")
	urls += "?" + c.urlParams_.Encode()
	req, err := http.NewRequest("PATCH", urls, body)
	if err != nil {
		return nil, err
	}
	req.Header = reqHeaders
	googleapi.Expand(req.URL, map[string]string{
		"name": c.name,
	})
	return gensupport.SendRequest(c.ctx_, c.s.client, req)
}

// Do executes the "mybusinessnotifications.accounts.updateNotificationSetting" call.
// Exactly one of *NotificationSetting or error will be non-nil. Any
// non-2xx status code is an error. Response headers are in either
// *NotificationSetting.ServerResponse.Header or (if a response was
// returned at all) in error.(*googleapi.Error).Header. Use
// googleapi.IsNotModified to check whether the returned error was
// because http.StatusNotModified was returned.
func (c *AccountsUpdateNotificationSettingCall) Do(opts ...googleapi.CallOption) (*NotificationSetting, error) {
	gensupport.SetOptions(c.urlParams_, opts...)
	res, err := c.doRequest("json")
	if res != nil && res.StatusCode == http.StatusNotModified {
		if res.Body != nil {
			res.Body.Close()
		}
		return nil, &googleapi.Error{
			Code:   res.StatusCode,
			Header: res.Header,
		}
	}
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	ret := &NotificationSetting{
		ServerResponse: googleapi.ServerResponse{
			Header:         res.Header,
			HTTPStatusCode: res.StatusCode,
		},
	}
	target := &ret
	if err := gensupport.DecodeResponse(target, res); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Sets the pubsub notification setting for the account informing Google which topic to send pubsub notifications for. Use the notification_types field within notification_setting to manipulate the events an account wants to subscribe to. An account will only have one notification setting resource, and only one pubsub topic can be set. To delete the setting, update with an empty notification_types",
	//   "flatPath": "v1/accounts/{accountsId}/notificationSetting",
	//   "httpMethod": "PATCH",
	//   "id": "mybusinessnotifications.accounts.updateNotificationSetting",
	//   "parameterOrder": [
	//     "name"
	//   ],
	//   "parameters": {
	//     "name": {
	//       "description": "Required. The resource name this setting is for. This is of the form `accounts/{account_id}/notificationSetting`.",
	//       "location": "path",
	//       "pattern": "^accounts/[^/]+/notificationSetting$",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "updateMask": {
	//       "description": "Required. The specific fields that should be updated. The only editable field is notification_setting.",
	//       "format": "google-fieldmask",
	//       "location": "query",
	//       "type": "string"
	//     }
	//   },
	//   "path": "v1/{+name}",
	//   "request": {
	//     "$ref": "NotificationSetting"
	//   },
	//   "response": {
	//     "$ref": "NotificationSetting"
	//   }
	// }

}
