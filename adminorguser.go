// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package ngc

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"

	"github.com/NVIDIADemo/ngc-go/internal/apijson"
	"github.com/NVIDIADemo/ngc-go/internal/apiquery"
	"github.com/NVIDIADemo/ngc-go/internal/param"
	"github.com/NVIDIADemo/ngc-go/internal/requestconfig"
	"github.com/NVIDIADemo/ngc-go/option"
	"github.com/NVIDIADemo/ngc-go/shared"
)

// AdminOrgUserService contains methods and other services that help with
// interacting with the ngc API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAdminOrgUserService] method instead.
type AdminOrgUserService struct {
	Options []option.RequestOption
}

// NewAdminOrgUserService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewAdminOrgUserService(opts ...option.RequestOption) (r *AdminOrgUserService) {
	r = &AdminOrgUserService{}
	r.Options = opts
	return
}

// Create an user in an Organization (Super Admin privileges required)
func (r *AdminOrgUserService) New(ctx context.Context, orgName string, params AdminOrgUserNewParams, opts ...option.RequestOption) (res *shared.User, err error) {
	opts = append(r.Options[:], opts...)
	if orgName == "" {
		err = errors.New("missing required org-name parameter")
		return
	}
	path := fmt.Sprintf("v2/admin/org/%s/users", orgName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Get info and role/invitation in an org by email or id
func (r *AdminOrgUserService) Get(ctx context.Context, orgName string, userEmailOrID string, opts ...option.RequestOption) (res *shared.User, err error) {
	opts = append(r.Options[:], opts...)
	if orgName == "" {
		err = errors.New("missing required org-name parameter")
		return
	}
	if userEmailOrID == "" {
		err = errors.New("missing required user-email-or-id parameter")
		return
	}
	path := fmt.Sprintf("v3/orgs/%s/users/%s", orgName, userEmailOrID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Add existing User to an Org
func (r *AdminOrgUserService) Add(ctx context.Context, orgName string, id string, body AdminOrgUserAddParams, opts ...option.RequestOption) (res *shared.User, err error) {
	opts = append(r.Options[:], opts...)
	if orgName == "" {
		err = errors.New("missing required org-name parameter")
		return
	}
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("v2/admin/org/%s/users/%s", orgName, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Add user role in org.
func (r *AdminOrgUserService) AddRole(ctx context.Context, orgName string, id string, body AdminOrgUserAddRoleParams, opts ...option.RequestOption) (res *shared.User, err error) {
	opts = append(r.Options[:], opts...)
	if orgName == "" {
		err = errors.New("missing required org-name parameter")
		return
	}
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("v2/admin/org/%s/users/%s/add-role", orgName, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// List all organizations with entitlements. (SuperAdmin, ECM and Billing
// privileges required)
func (r *AdminOrgUserService) GetEntitlements(ctx context.Context, orgName string, query AdminOrgUserGetEntitlementsParams, opts ...option.RequestOption) (res *http.Response, err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if orgName == "" {
		err = errors.New("missing required org-name parameter")
		return
	}
	path := fmt.Sprintf("v2/admin/org/%s/entitlements", orgName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Remove User from org.
func (r *AdminOrgUserService) Remove(ctx context.Context, orgName string, id string, opts ...option.RequestOption) (res *AdminOrgUserRemoveResponse, err error) {
	opts = append(r.Options[:], opts...)
	if orgName == "" {
		err = errors.New("missing required org-name parameter")
		return
	}
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("v2/admin/org/%s/users/%s", orgName, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

type AdminOrgUserRemoveResponse struct {
	RequestStatus AdminOrgUserRemoveResponseRequestStatus `json:"requestStatus"`
	JSON          adminOrgUserRemoveResponseJSON          `json:"-"`
}

// adminOrgUserRemoveResponseJSON contains the JSON metadata for the struct
// [AdminOrgUserRemoveResponse]
type adminOrgUserRemoveResponseJSON struct {
	RequestStatus apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *AdminOrgUserRemoveResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r adminOrgUserRemoveResponseJSON) RawJSON() string {
	return r.raw
}

type AdminOrgUserRemoveResponseRequestStatus struct {
	RequestID string `json:"requestId"`
	ServerID  string `json:"serverId"`
	// Describes response status reported by the server.
	StatusCode        AdminOrgUserRemoveResponseRequestStatusStatusCode `json:"statusCode"`
	StatusDescription string                                            `json:"statusDescription"`
	JSON              adminOrgUserRemoveResponseRequestStatusJSON       `json:"-"`
}

// adminOrgUserRemoveResponseRequestStatusJSON contains the JSON metadata for the
// struct [AdminOrgUserRemoveResponseRequestStatus]
type adminOrgUserRemoveResponseRequestStatusJSON struct {
	RequestID         apijson.Field
	ServerID          apijson.Field
	StatusCode        apijson.Field
	StatusDescription apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *AdminOrgUserRemoveResponseRequestStatus) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r adminOrgUserRemoveResponseRequestStatusJSON) RawJSON() string {
	return r.raw
}

// Describes response status reported by the server.
type AdminOrgUserRemoveResponseRequestStatusStatusCode string

const (
	AdminOrgUserRemoveResponseRequestStatusStatusCodeUnknown                    AdminOrgUserRemoveResponseRequestStatusStatusCode = "UNKNOWN"
	AdminOrgUserRemoveResponseRequestStatusStatusCodeSuccess                    AdminOrgUserRemoveResponseRequestStatusStatusCode = "SUCCESS"
	AdminOrgUserRemoveResponseRequestStatusStatusCodeUnauthorized               AdminOrgUserRemoveResponseRequestStatusStatusCode = "UNAUTHORIZED"
	AdminOrgUserRemoveResponseRequestStatusStatusCodePaymentRequired            AdminOrgUserRemoveResponseRequestStatusStatusCode = "PAYMENT_REQUIRED"
	AdminOrgUserRemoveResponseRequestStatusStatusCodeForbidden                  AdminOrgUserRemoveResponseRequestStatusStatusCode = "FORBIDDEN"
	AdminOrgUserRemoveResponseRequestStatusStatusCodeTimeout                    AdminOrgUserRemoveResponseRequestStatusStatusCode = "TIMEOUT"
	AdminOrgUserRemoveResponseRequestStatusStatusCodeExists                     AdminOrgUserRemoveResponseRequestStatusStatusCode = "EXISTS"
	AdminOrgUserRemoveResponseRequestStatusStatusCodeNotFound                   AdminOrgUserRemoveResponseRequestStatusStatusCode = "NOT_FOUND"
	AdminOrgUserRemoveResponseRequestStatusStatusCodeInternalError              AdminOrgUserRemoveResponseRequestStatusStatusCode = "INTERNAL_ERROR"
	AdminOrgUserRemoveResponseRequestStatusStatusCodeInvalidRequest             AdminOrgUserRemoveResponseRequestStatusStatusCode = "INVALID_REQUEST"
	AdminOrgUserRemoveResponseRequestStatusStatusCodeInvalidRequestVersion      AdminOrgUserRemoveResponseRequestStatusStatusCode = "INVALID_REQUEST_VERSION"
	AdminOrgUserRemoveResponseRequestStatusStatusCodeInvalidRequestData         AdminOrgUserRemoveResponseRequestStatusStatusCode = "INVALID_REQUEST_DATA"
	AdminOrgUserRemoveResponseRequestStatusStatusCodeMethodNotAllowed           AdminOrgUserRemoveResponseRequestStatusStatusCode = "METHOD_NOT_ALLOWED"
	AdminOrgUserRemoveResponseRequestStatusStatusCodeConflict                   AdminOrgUserRemoveResponseRequestStatusStatusCode = "CONFLICT"
	AdminOrgUserRemoveResponseRequestStatusStatusCodeUnprocessableEntity        AdminOrgUserRemoveResponseRequestStatusStatusCode = "UNPROCESSABLE_ENTITY"
	AdminOrgUserRemoveResponseRequestStatusStatusCodeTooManyRequests            AdminOrgUserRemoveResponseRequestStatusStatusCode = "TOO_MANY_REQUESTS"
	AdminOrgUserRemoveResponseRequestStatusStatusCodeInsufficientStorage        AdminOrgUserRemoveResponseRequestStatusStatusCode = "INSUFFICIENT_STORAGE"
	AdminOrgUserRemoveResponseRequestStatusStatusCodeServiceUnavailable         AdminOrgUserRemoveResponseRequestStatusStatusCode = "SERVICE_UNAVAILABLE"
	AdminOrgUserRemoveResponseRequestStatusStatusCodePayloadTooLarge            AdminOrgUserRemoveResponseRequestStatusStatusCode = "PAYLOAD_TOO_LARGE"
	AdminOrgUserRemoveResponseRequestStatusStatusCodeNotAcceptable              AdminOrgUserRemoveResponseRequestStatusStatusCode = "NOT_ACCEPTABLE"
	AdminOrgUserRemoveResponseRequestStatusStatusCodeUnavailableForLegalReasons AdminOrgUserRemoveResponseRequestStatusStatusCode = "UNAVAILABLE_FOR_LEGAL_REASONS"
	AdminOrgUserRemoveResponseRequestStatusStatusCodeBadGateway                 AdminOrgUserRemoveResponseRequestStatusStatusCode = "BAD_GATEWAY"
)

func (r AdminOrgUserRemoveResponseRequestStatusStatusCode) IsKnown() bool {
	switch r {
	case AdminOrgUserRemoveResponseRequestStatusStatusCodeUnknown, AdminOrgUserRemoveResponseRequestStatusStatusCodeSuccess, AdminOrgUserRemoveResponseRequestStatusStatusCodeUnauthorized, AdminOrgUserRemoveResponseRequestStatusStatusCodePaymentRequired, AdminOrgUserRemoveResponseRequestStatusStatusCodeForbidden, AdminOrgUserRemoveResponseRequestStatusStatusCodeTimeout, AdminOrgUserRemoveResponseRequestStatusStatusCodeExists, AdminOrgUserRemoveResponseRequestStatusStatusCodeNotFound, AdminOrgUserRemoveResponseRequestStatusStatusCodeInternalError, AdminOrgUserRemoveResponseRequestStatusStatusCodeInvalidRequest, AdminOrgUserRemoveResponseRequestStatusStatusCodeInvalidRequestVersion, AdminOrgUserRemoveResponseRequestStatusStatusCodeInvalidRequestData, AdminOrgUserRemoveResponseRequestStatusStatusCodeMethodNotAllowed, AdminOrgUserRemoveResponseRequestStatusStatusCodeConflict, AdminOrgUserRemoveResponseRequestStatusStatusCodeUnprocessableEntity, AdminOrgUserRemoveResponseRequestStatusStatusCodeTooManyRequests, AdminOrgUserRemoveResponseRequestStatusStatusCodeInsufficientStorage, AdminOrgUserRemoveResponseRequestStatusStatusCodeServiceUnavailable, AdminOrgUserRemoveResponseRequestStatusStatusCodePayloadTooLarge, AdminOrgUserRemoveResponseRequestStatusStatusCodeNotAcceptable, AdminOrgUserRemoveResponseRequestStatusStatusCodeUnavailableForLegalReasons, AdminOrgUserRemoveResponseRequestStatusStatusCodeBadGateway:
		return true
	}
	return false
}

type AdminOrgUserNewParams struct {
	// Email address of the user. This should be unique.
	Email param.Field[string] `json:"email,required"`
	// If the IDP ID is provided then it is used instead of the one configured for the
	// organization
	IdpID param.Field[string] `query:"idp-id"`
	// Boolean to send email notification, default is true
	SendEmail param.Field[bool] `query:"send-email"`
	// indicates if user has opt in to nvidia emails
	EmailOptIn param.Field[bool] `json:"emailOptIn"`
	// indicates if user has accepted EULA
	EulaAccepted param.Field[bool] `json:"eulaAccepted"`
	// user name
	Name param.Field[string] `json:"name"`
	// DEPRECATED - use roleTypes which allows multiple roles
	RoleType param.Field[string] `json:"roleType"`
	// feature roles to give to the user
	RoleTypes param.Field[[]string] `json:"roleTypes"`
	// user job role
	SalesforceContactJobRole param.Field[string] `json:"salesforceContactJobRole"`
	// Metadata information about the user.
	UserMetadata param.Field[AdminOrgUserNewParamsUserMetadata] `json:"userMetadata"`
}

func (r AdminOrgUserNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// URLQuery serializes [AdminOrgUserNewParams]'s query parameters as `url.Values`.
func (r AdminOrgUserNewParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Metadata information about the user.
type AdminOrgUserNewParamsUserMetadata struct {
	// Name of the company
	Company param.Field[string] `json:"company"`
	// Company URL
	CompanyURL param.Field[string] `json:"companyUrl"`
	// Country of the user
	Country param.Field[string] `json:"country"`
	// User's first name
	FirstName param.Field[string] `json:"firstName"`
	// Industry segment
	Industry param.Field[string] `json:"industry"`
	// List of development areas that user has interest
	Interest param.Field[[]string] `json:"interest"`
	// User's last name
	LastName param.Field[string] `json:"lastName"`
	// Role of the user in the company
	Role param.Field[string] `json:"role"`
}

func (r AdminOrgUserNewParamsUserMetadata) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AdminOrgUserAddParams struct {
	UserRoleDefaultsToRegistryRead param.Field[string] `query:"user role, defaults to REGISTRY_READ"`
}

// URLQuery serializes [AdminOrgUserAddParams]'s query parameters as `url.Values`.
func (r AdminOrgUserAddParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type AdminOrgUserAddRoleParams struct {
	Roles param.Field[[]string] `query:"roles"`
}

// URLQuery serializes [AdminOrgUserAddRoleParams]'s query parameters as
// `url.Values`.
func (r AdminOrgUserAddRoleParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type AdminOrgUserGetEntitlementsParams struct {
	// get is paid subscription entitlements
	IsPaidSubscription param.Field[bool] `query:"is-paid-subscription"`
	// filter by product-name
	ProductName param.Field[string] `query:"product-name"`
}

// URLQuery serializes [AdminOrgUserGetEntitlementsParams]'s query parameters as
// `url.Values`.
func (r AdminOrgUserGetEntitlementsParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
