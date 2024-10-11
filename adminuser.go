// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package ngc

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/NVIDIADemo/ngc-go/internal/apijson"
	"github.com/NVIDIADemo/ngc-go/internal/requestconfig"
	"github.com/NVIDIADemo/ngc-go/option"
	"github.com/NVIDIADemo/ngc-go/shared"
)

// AdminUserService contains methods and other services that help with interacting
// with the ngc API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAdminUserService] method instead.
type AdminUserService struct {
	Options []option.RequestOption
}

// NewAdminUserService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewAdminUserService(opts ...option.RequestOption) (r *AdminUserService) {
	r = &AdminUserService{}
	r.Options = opts
	return
}

// Sync crm id with user email (Super Admin privileges required)
func (r *AdminUserService) CRMSync(ctx context.Context, id int64, opts ...option.RequestOption) (res *AdminUserCRMSyncResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("v2/admin/users/%v/crm-sync", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
}

// Migrate User Deprecated Roles.
func (r *AdminUserService) MigrateDeprecatedRoles(ctx context.Context, id string, opts ...option.RequestOption) (res *shared.User, err error) {
	opts = append(r.Options[:], opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("v2/admin/users/%s/migrate-deprecated-roles", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
}

type AdminUserCRMSyncResponse struct {
	RequestStatus AdminUserCRMSyncResponseRequestStatus `json:"requestStatus"`
	JSON          adminUserCRMSyncResponseJSON          `json:"-"`
}

// adminUserCRMSyncResponseJSON contains the JSON metadata for the struct
// [AdminUserCRMSyncResponse]
type adminUserCRMSyncResponseJSON struct {
	RequestStatus apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *AdminUserCRMSyncResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r adminUserCRMSyncResponseJSON) RawJSON() string {
	return r.raw
}

type AdminUserCRMSyncResponseRequestStatus struct {
	RequestID string `json:"requestId"`
	ServerID  string `json:"serverId"`
	// Describes response status reported by the server.
	StatusCode        AdminUserCRMSyncResponseRequestStatusStatusCode `json:"statusCode"`
	StatusDescription string                                          `json:"statusDescription"`
	JSON              adminUserCRMSyncResponseRequestStatusJSON       `json:"-"`
}

// adminUserCRMSyncResponseRequestStatusJSON contains the JSON metadata for the
// struct [AdminUserCRMSyncResponseRequestStatus]
type adminUserCRMSyncResponseRequestStatusJSON struct {
	RequestID         apijson.Field
	ServerID          apijson.Field
	StatusCode        apijson.Field
	StatusDescription apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *AdminUserCRMSyncResponseRequestStatus) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r adminUserCRMSyncResponseRequestStatusJSON) RawJSON() string {
	return r.raw
}

// Describes response status reported by the server.
type AdminUserCRMSyncResponseRequestStatusStatusCode string

const (
	AdminUserCRMSyncResponseRequestStatusStatusCodeUnknown                    AdminUserCRMSyncResponseRequestStatusStatusCode = "UNKNOWN"
	AdminUserCRMSyncResponseRequestStatusStatusCodeSuccess                    AdminUserCRMSyncResponseRequestStatusStatusCode = "SUCCESS"
	AdminUserCRMSyncResponseRequestStatusStatusCodeUnauthorized               AdminUserCRMSyncResponseRequestStatusStatusCode = "UNAUTHORIZED"
	AdminUserCRMSyncResponseRequestStatusStatusCodePaymentRequired            AdminUserCRMSyncResponseRequestStatusStatusCode = "PAYMENT_REQUIRED"
	AdminUserCRMSyncResponseRequestStatusStatusCodeForbidden                  AdminUserCRMSyncResponseRequestStatusStatusCode = "FORBIDDEN"
	AdminUserCRMSyncResponseRequestStatusStatusCodeTimeout                    AdminUserCRMSyncResponseRequestStatusStatusCode = "TIMEOUT"
	AdminUserCRMSyncResponseRequestStatusStatusCodeExists                     AdminUserCRMSyncResponseRequestStatusStatusCode = "EXISTS"
	AdminUserCRMSyncResponseRequestStatusStatusCodeNotFound                   AdminUserCRMSyncResponseRequestStatusStatusCode = "NOT_FOUND"
	AdminUserCRMSyncResponseRequestStatusStatusCodeInternalError              AdminUserCRMSyncResponseRequestStatusStatusCode = "INTERNAL_ERROR"
	AdminUserCRMSyncResponseRequestStatusStatusCodeInvalidRequest             AdminUserCRMSyncResponseRequestStatusStatusCode = "INVALID_REQUEST"
	AdminUserCRMSyncResponseRequestStatusStatusCodeInvalidRequestVersion      AdminUserCRMSyncResponseRequestStatusStatusCode = "INVALID_REQUEST_VERSION"
	AdminUserCRMSyncResponseRequestStatusStatusCodeInvalidRequestData         AdminUserCRMSyncResponseRequestStatusStatusCode = "INVALID_REQUEST_DATA"
	AdminUserCRMSyncResponseRequestStatusStatusCodeMethodNotAllowed           AdminUserCRMSyncResponseRequestStatusStatusCode = "METHOD_NOT_ALLOWED"
	AdminUserCRMSyncResponseRequestStatusStatusCodeConflict                   AdminUserCRMSyncResponseRequestStatusStatusCode = "CONFLICT"
	AdminUserCRMSyncResponseRequestStatusStatusCodeUnprocessableEntity        AdminUserCRMSyncResponseRequestStatusStatusCode = "UNPROCESSABLE_ENTITY"
	AdminUserCRMSyncResponseRequestStatusStatusCodeTooManyRequests            AdminUserCRMSyncResponseRequestStatusStatusCode = "TOO_MANY_REQUESTS"
	AdminUserCRMSyncResponseRequestStatusStatusCodeInsufficientStorage        AdminUserCRMSyncResponseRequestStatusStatusCode = "INSUFFICIENT_STORAGE"
	AdminUserCRMSyncResponseRequestStatusStatusCodeServiceUnavailable         AdminUserCRMSyncResponseRequestStatusStatusCode = "SERVICE_UNAVAILABLE"
	AdminUserCRMSyncResponseRequestStatusStatusCodePayloadTooLarge            AdminUserCRMSyncResponseRequestStatusStatusCode = "PAYLOAD_TOO_LARGE"
	AdminUserCRMSyncResponseRequestStatusStatusCodeNotAcceptable              AdminUserCRMSyncResponseRequestStatusStatusCode = "NOT_ACCEPTABLE"
	AdminUserCRMSyncResponseRequestStatusStatusCodeUnavailableForLegalReasons AdminUserCRMSyncResponseRequestStatusStatusCode = "UNAVAILABLE_FOR_LEGAL_REASONS"
	AdminUserCRMSyncResponseRequestStatusStatusCodeBadGateway                 AdminUserCRMSyncResponseRequestStatusStatusCode = "BAD_GATEWAY"
)

func (r AdminUserCRMSyncResponseRequestStatusStatusCode) IsKnown() bool {
	switch r {
	case AdminUserCRMSyncResponseRequestStatusStatusCodeUnknown, AdminUserCRMSyncResponseRequestStatusStatusCodeSuccess, AdminUserCRMSyncResponseRequestStatusStatusCodeUnauthorized, AdminUserCRMSyncResponseRequestStatusStatusCodePaymentRequired, AdminUserCRMSyncResponseRequestStatusStatusCodeForbidden, AdminUserCRMSyncResponseRequestStatusStatusCodeTimeout, AdminUserCRMSyncResponseRequestStatusStatusCodeExists, AdminUserCRMSyncResponseRequestStatusStatusCodeNotFound, AdminUserCRMSyncResponseRequestStatusStatusCodeInternalError, AdminUserCRMSyncResponseRequestStatusStatusCodeInvalidRequest, AdminUserCRMSyncResponseRequestStatusStatusCodeInvalidRequestVersion, AdminUserCRMSyncResponseRequestStatusStatusCodeInvalidRequestData, AdminUserCRMSyncResponseRequestStatusStatusCodeMethodNotAllowed, AdminUserCRMSyncResponseRequestStatusStatusCodeConflict, AdminUserCRMSyncResponseRequestStatusStatusCodeUnprocessableEntity, AdminUserCRMSyncResponseRequestStatusStatusCodeTooManyRequests, AdminUserCRMSyncResponseRequestStatusStatusCodeInsufficientStorage, AdminUserCRMSyncResponseRequestStatusStatusCodeServiceUnavailable, AdminUserCRMSyncResponseRequestStatusStatusCodePayloadTooLarge, AdminUserCRMSyncResponseRequestStatusStatusCodeNotAcceptable, AdminUserCRMSyncResponseRequestStatusStatusCodeUnavailableForLegalReasons, AdminUserCRMSyncResponseRequestStatusStatusCodeBadGateway:
		return true
	}
	return false
}
