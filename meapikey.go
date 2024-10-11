// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package ngc

import (
	"context"
	"net/http"
	"net/url"

	"github.com/NVIDIADemo/ngc-go/internal/apijson"
	"github.com/NVIDIADemo/ngc-go/internal/apiquery"
	"github.com/NVIDIADemo/ngc-go/internal/param"
	"github.com/NVIDIADemo/ngc-go/internal/requestconfig"
	"github.com/NVIDIADemo/ngc-go/option"
	"github.com/NVIDIADemo/ngc-go/shared"
)

// MeAPIKeyService contains methods and other services that help with interacting
// with the ngc API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewMeAPIKeyService] method instead.
type MeAPIKeyService struct {
	Options []option.RequestOption
}

// NewMeAPIKeyService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewMeAPIKeyService(opts ...option.RequestOption) (r *MeAPIKeyService) {
	r = &MeAPIKeyService{}
	r.Options = opts
	return
}

// What am I?
func (r *MeAPIKeyService) Get(ctx context.Context, query MeAPIKeyGetParams, opts ...option.RequestOption) (res *shared.User, err error) {
	opts = append(r.Options[:], opts...)
	path := "v2/users/me"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Edit current user profile
func (r *MeAPIKeyService) Update(ctx context.Context, body MeAPIKeyUpdateParams, opts ...option.RequestOption) (res *shared.User, err error) {
	opts = append(r.Options[:], opts...)
	path := "v2/users/me"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// Generate API Key
func (r *MeAPIKeyService) NewAPIKey(ctx context.Context, opts ...option.RequestOption) (res *UserKeyResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v2/users/me/api-key"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
}

// response to a request to access key such as docker token
type UserKeyResponse struct {
	Key              string                       `json:"key,required"`
	CloudNfsKey      string                       `json:"cloudNfsKey"`
	CloudNfsKeyPwd   string                       `json:"cloudNfsKeyPwd"`
	CloudNfsUserName string                       `json:"cloudNfsUserName"`
	RequestStatus    UserKeyResponseRequestStatus `json:"requestStatus"`
	JSON             userKeyResponseJSON          `json:"-"`
}

// userKeyResponseJSON contains the JSON metadata for the struct [UserKeyResponse]
type userKeyResponseJSON struct {
	Key              apijson.Field
	CloudNfsKey      apijson.Field
	CloudNfsKeyPwd   apijson.Field
	CloudNfsUserName apijson.Field
	RequestStatus    apijson.Field
	raw              string
	ExtraFields      map[string]apijson.Field
}

func (r *UserKeyResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userKeyResponseJSON) RawJSON() string {
	return r.raw
}

type UserKeyResponseRequestStatus struct {
	RequestID string `json:"requestId"`
	ServerID  string `json:"serverId"`
	// Describes response status reported by the server.
	StatusCode        UserKeyResponseRequestStatusStatusCode `json:"statusCode"`
	StatusDescription string                                 `json:"statusDescription"`
	JSON              userKeyResponseRequestStatusJSON       `json:"-"`
}

// userKeyResponseRequestStatusJSON contains the JSON metadata for the struct
// [UserKeyResponseRequestStatus]
type userKeyResponseRequestStatusJSON struct {
	RequestID         apijson.Field
	ServerID          apijson.Field
	StatusCode        apijson.Field
	StatusDescription apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *UserKeyResponseRequestStatus) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userKeyResponseRequestStatusJSON) RawJSON() string {
	return r.raw
}

// Describes response status reported by the server.
type UserKeyResponseRequestStatusStatusCode string

const (
	UserKeyResponseRequestStatusStatusCodeUnknown                    UserKeyResponseRequestStatusStatusCode = "UNKNOWN"
	UserKeyResponseRequestStatusStatusCodeSuccess                    UserKeyResponseRequestStatusStatusCode = "SUCCESS"
	UserKeyResponseRequestStatusStatusCodeUnauthorized               UserKeyResponseRequestStatusStatusCode = "UNAUTHORIZED"
	UserKeyResponseRequestStatusStatusCodePaymentRequired            UserKeyResponseRequestStatusStatusCode = "PAYMENT_REQUIRED"
	UserKeyResponseRequestStatusStatusCodeForbidden                  UserKeyResponseRequestStatusStatusCode = "FORBIDDEN"
	UserKeyResponseRequestStatusStatusCodeTimeout                    UserKeyResponseRequestStatusStatusCode = "TIMEOUT"
	UserKeyResponseRequestStatusStatusCodeExists                     UserKeyResponseRequestStatusStatusCode = "EXISTS"
	UserKeyResponseRequestStatusStatusCodeNotFound                   UserKeyResponseRequestStatusStatusCode = "NOT_FOUND"
	UserKeyResponseRequestStatusStatusCodeInternalError              UserKeyResponseRequestStatusStatusCode = "INTERNAL_ERROR"
	UserKeyResponseRequestStatusStatusCodeInvalidRequest             UserKeyResponseRequestStatusStatusCode = "INVALID_REQUEST"
	UserKeyResponseRequestStatusStatusCodeInvalidRequestVersion      UserKeyResponseRequestStatusStatusCode = "INVALID_REQUEST_VERSION"
	UserKeyResponseRequestStatusStatusCodeInvalidRequestData         UserKeyResponseRequestStatusStatusCode = "INVALID_REQUEST_DATA"
	UserKeyResponseRequestStatusStatusCodeMethodNotAllowed           UserKeyResponseRequestStatusStatusCode = "METHOD_NOT_ALLOWED"
	UserKeyResponseRequestStatusStatusCodeConflict                   UserKeyResponseRequestStatusStatusCode = "CONFLICT"
	UserKeyResponseRequestStatusStatusCodeUnprocessableEntity        UserKeyResponseRequestStatusStatusCode = "UNPROCESSABLE_ENTITY"
	UserKeyResponseRequestStatusStatusCodeTooManyRequests            UserKeyResponseRequestStatusStatusCode = "TOO_MANY_REQUESTS"
	UserKeyResponseRequestStatusStatusCodeInsufficientStorage        UserKeyResponseRequestStatusStatusCode = "INSUFFICIENT_STORAGE"
	UserKeyResponseRequestStatusStatusCodeServiceUnavailable         UserKeyResponseRequestStatusStatusCode = "SERVICE_UNAVAILABLE"
	UserKeyResponseRequestStatusStatusCodePayloadTooLarge            UserKeyResponseRequestStatusStatusCode = "PAYLOAD_TOO_LARGE"
	UserKeyResponseRequestStatusStatusCodeNotAcceptable              UserKeyResponseRequestStatusStatusCode = "NOT_ACCEPTABLE"
	UserKeyResponseRequestStatusStatusCodeUnavailableForLegalReasons UserKeyResponseRequestStatusStatusCode = "UNAVAILABLE_FOR_LEGAL_REASONS"
	UserKeyResponseRequestStatusStatusCodeBadGateway                 UserKeyResponseRequestStatusStatusCode = "BAD_GATEWAY"
)

func (r UserKeyResponseRequestStatusStatusCode) IsKnown() bool {
	switch r {
	case UserKeyResponseRequestStatusStatusCodeUnknown, UserKeyResponseRequestStatusStatusCodeSuccess, UserKeyResponseRequestStatusStatusCodeUnauthorized, UserKeyResponseRequestStatusStatusCodePaymentRequired, UserKeyResponseRequestStatusStatusCodeForbidden, UserKeyResponseRequestStatusStatusCodeTimeout, UserKeyResponseRequestStatusStatusCodeExists, UserKeyResponseRequestStatusStatusCodeNotFound, UserKeyResponseRequestStatusStatusCodeInternalError, UserKeyResponseRequestStatusStatusCodeInvalidRequest, UserKeyResponseRequestStatusStatusCodeInvalidRequestVersion, UserKeyResponseRequestStatusStatusCodeInvalidRequestData, UserKeyResponseRequestStatusStatusCodeMethodNotAllowed, UserKeyResponseRequestStatusStatusCodeConflict, UserKeyResponseRequestStatusStatusCodeUnprocessableEntity, UserKeyResponseRequestStatusStatusCodeTooManyRequests, UserKeyResponseRequestStatusStatusCodeInsufficientStorage, UserKeyResponseRequestStatusStatusCodeServiceUnavailable, UserKeyResponseRequestStatusStatusCodePayloadTooLarge, UserKeyResponseRequestStatusStatusCodeNotAcceptable, UserKeyResponseRequestStatusStatusCodeUnavailableForLegalReasons, UserKeyResponseRequestStatusStatusCodeBadGateway:
		return true
	}
	return false
}

type MeAPIKeyGetParams struct {
	OrgName param.Field[string] `query:"org-name"`
}

// URLQuery serializes [MeAPIKeyGetParams]'s query parameters as `url.Values`.
func (r MeAPIKeyGetParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type MeAPIKeyUpdateParams struct {
	// indicates if user has opt in to nvidia emails
	HasEmailOptIn param.Field[bool] `json:"hasEmailOptIn"`
	// indicates if user has accepted AI Foundry Partnerships End User License
	// Agreement.
	HasSignedAIFoundryPartnershipsEula param.Field[bool] `json:"hasSignedAiFoundryPartnershipsEULA"`
	// indicates if user has accepted Base Command EULA
	HasSignedBaseCommandEula param.Field[bool] `json:"hasSignedBaseCommandEULA"`
	// indicates if user has accepted Base Command Manager End User License Agreement.
	HasSignedBaseCommandManagerEula param.Field[bool] `json:"hasSignedBaseCommandManagerEULA"`
	// indicates if user has accepted BioNeMo End User License Agreement.
	HasSignedBioNeMoEula param.Field[bool] `json:"hasSignedBioNeMoEULA"`
	// indicates if user has accepted container publishing eula
	HasSignedContainerPublishingEula param.Field[bool] `json:"hasSignedContainerPublishingEULA"`
	// indicates if user has accepted CuOpt End User License Agreement.
	HasSignedCuOptEula param.Field[bool] `json:"hasSignedCuOptEULA"`
	// indicates if user has accepted Earth-2 End User License Agreement.
	HasSignedEarth2Eula param.Field[bool] `json:"hasSignedEarth2EULA"`
	// indicates if user has accepted EGX EULA
	HasSignedEgxEula param.Field[bool] `json:"hasSignedEgxEULA"`
	// indicates if user has accepted NGC EULA
	HasSignedEula param.Field[bool] `json:"hasSignedEULA"`
	// indicates if user has accepted Fleet Command End User License Agreement.
	HasSignedFleetCommandEula param.Field[bool] `json:"hasSignedFleetCommandEULA"`
	// indicates if user has accepted LLM End User License Agreement.
	HasSignedLlmEula param.Field[bool] `json:"hasSignedLlmEULA"`
	// indicates if user has accepted Fleet Command End User License Agreement.
	HasSignedNvaieeula param.Field[bool] `json:"hasSignedNVAIEEULA"`
	// indicates if user has accepted NVIDIA EULA
	HasSignedNvidiaEula param.Field[bool] `json:"hasSignedNvidiaEULA"`
	// indicates if user has accepted Nvidia Quantum Cloud End User License Agreement.
	HasSignedNvqceula param.Field[bool] `json:"hasSignedNVQCEULA"`
	// indicates if user has accepted Omniverse End User License Agreement.
	HasSignedOmniverseEula param.Field[bool] `json:"hasSignedOmniverseEULA"`
	// indicates if the user has signed the Privacy Policy
	HasSignedPrivacyPolicy param.Field[bool] `json:"hasSignedPrivacyPolicy"`
	// indicates if user has consented to share their registration info with other
	// parties
	HasSignedThirdPartyRegistryShareEula param.Field[bool] `json:"hasSignedThirdPartyRegistryShareEULA"`
	// user name
	Name param.Field[string] `json:"name"`
	// Metadata information about the user.
	UserMetadata param.Field[MeAPIKeyUpdateParamsUserMetadata] `json:"userMetadata"`
}

func (r MeAPIKeyUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Metadata information about the user.
type MeAPIKeyUpdateParamsUserMetadata struct {
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

func (r MeAPIKeyUpdateParamsUserMetadata) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
