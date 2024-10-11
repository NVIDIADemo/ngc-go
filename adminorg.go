// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package ngc

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/NVIDIADemo/ngc-go/internal/apijson"
	"github.com/NVIDIADemo/ngc-go/internal/param"
	"github.com/NVIDIADemo/ngc-go/internal/requestconfig"
	"github.com/NVIDIADemo/ngc-go/option"
)

// AdminOrgService contains methods and other services that help with interacting
// with the ngc API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAdminOrgService] method instead.
type AdminOrgService struct {
	Options []option.RequestOption
	Users   *AdminOrgUserService
	Teams   *AdminOrgTeamService
}

// NewAdminOrgService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewAdminOrgService(opts ...option.RequestOption) (r *AdminOrgService) {
	r = &AdminOrgService{}
	r.Options = opts
	r.Users = NewAdminOrgUserService(opts...)
	r.Teams = NewAdminOrgTeamService(opts...)
	return
}

// Create a new organization. (SuperAdmin privileges required)
func (r *AdminOrgService) New(ctx context.Context, body AdminOrgNewParams, opts ...option.RequestOption) (res *http.Response, err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := "v2/admin/orgs"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Get organization or proto organization info. (SuperAdmin privileges required)
func (r *AdminOrgService) Get(ctx context.Context, orgName string, opts ...option.RequestOption) (res *http.Response, err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if orgName == "" {
		err = errors.New("missing required org-name parameter")
		return
	}
	path := fmt.Sprintf("v3/admin/org/%s", orgName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Update org information and settings. Superadmin privileges required
func (r *AdminOrgService) Update(ctx context.Context, orgName string, body AdminOrgUpdateParams, opts ...option.RequestOption) (res *http.Response, err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if orgName == "" {
		err = errors.New("missing required org-name parameter")
		return
	}
	path := fmt.Sprintf("v3/admin/org/%s", orgName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// Backfill Orgs to Kratos
func (r *AdminOrgService) BackfillOrgsToKratos(ctx context.Context, opts ...option.RequestOption) (res *http.Response, err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := "v2/admin/backfill-orgs-to-kratos"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
}

// Create org product enablement
func (r *AdminOrgService) Enable(ctx context.Context, orgName string, body AdminOrgEnableParams, opts ...option.RequestOption) (res *http.Response, err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if orgName == "" {
		err = errors.New("missing required org-name parameter")
		return
	}
	path := fmt.Sprintf("v2/admin/org/%s/enablement", orgName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Backfill the org owner for users
func (r *AdminOrgService) OrgOwnerBackfill(ctx context.Context, orgName string, opts ...option.RequestOption) (res *AdminOrgOrgOwnerBackfillResponse, err error) {
	opts = append(r.Options[:], opts...)
	if orgName == "" {
		err = errors.New("missing required org-name parameter")
		return
	}
	path := fmt.Sprintf("v2/admin/org/%s/org-owner-backfill", orgName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
}

type AdminOrgOrgOwnerBackfillResponse struct {
	RequestStatus AdminOrgOrgOwnerBackfillResponseRequestStatus `json:"requestStatus"`
	JSON          adminOrgOrgOwnerBackfillResponseJSON          `json:"-"`
}

// adminOrgOrgOwnerBackfillResponseJSON contains the JSON metadata for the struct
// [AdminOrgOrgOwnerBackfillResponse]
type adminOrgOrgOwnerBackfillResponseJSON struct {
	RequestStatus apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *AdminOrgOrgOwnerBackfillResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r adminOrgOrgOwnerBackfillResponseJSON) RawJSON() string {
	return r.raw
}

type AdminOrgOrgOwnerBackfillResponseRequestStatus struct {
	RequestID string `json:"requestId"`
	ServerID  string `json:"serverId"`
	// Describes response status reported by the server.
	StatusCode        AdminOrgOrgOwnerBackfillResponseRequestStatusStatusCode `json:"statusCode"`
	StatusDescription string                                                  `json:"statusDescription"`
	JSON              adminOrgOrgOwnerBackfillResponseRequestStatusJSON       `json:"-"`
}

// adminOrgOrgOwnerBackfillResponseRequestStatusJSON contains the JSON metadata for
// the struct [AdminOrgOrgOwnerBackfillResponseRequestStatus]
type adminOrgOrgOwnerBackfillResponseRequestStatusJSON struct {
	RequestID         apijson.Field
	ServerID          apijson.Field
	StatusCode        apijson.Field
	StatusDescription apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *AdminOrgOrgOwnerBackfillResponseRequestStatus) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r adminOrgOrgOwnerBackfillResponseRequestStatusJSON) RawJSON() string {
	return r.raw
}

// Describes response status reported by the server.
type AdminOrgOrgOwnerBackfillResponseRequestStatusStatusCode string

const (
	AdminOrgOrgOwnerBackfillResponseRequestStatusStatusCodeUnknown                    AdminOrgOrgOwnerBackfillResponseRequestStatusStatusCode = "UNKNOWN"
	AdminOrgOrgOwnerBackfillResponseRequestStatusStatusCodeSuccess                    AdminOrgOrgOwnerBackfillResponseRequestStatusStatusCode = "SUCCESS"
	AdminOrgOrgOwnerBackfillResponseRequestStatusStatusCodeUnauthorized               AdminOrgOrgOwnerBackfillResponseRequestStatusStatusCode = "UNAUTHORIZED"
	AdminOrgOrgOwnerBackfillResponseRequestStatusStatusCodePaymentRequired            AdminOrgOrgOwnerBackfillResponseRequestStatusStatusCode = "PAYMENT_REQUIRED"
	AdminOrgOrgOwnerBackfillResponseRequestStatusStatusCodeForbidden                  AdminOrgOrgOwnerBackfillResponseRequestStatusStatusCode = "FORBIDDEN"
	AdminOrgOrgOwnerBackfillResponseRequestStatusStatusCodeTimeout                    AdminOrgOrgOwnerBackfillResponseRequestStatusStatusCode = "TIMEOUT"
	AdminOrgOrgOwnerBackfillResponseRequestStatusStatusCodeExists                     AdminOrgOrgOwnerBackfillResponseRequestStatusStatusCode = "EXISTS"
	AdminOrgOrgOwnerBackfillResponseRequestStatusStatusCodeNotFound                   AdminOrgOrgOwnerBackfillResponseRequestStatusStatusCode = "NOT_FOUND"
	AdminOrgOrgOwnerBackfillResponseRequestStatusStatusCodeInternalError              AdminOrgOrgOwnerBackfillResponseRequestStatusStatusCode = "INTERNAL_ERROR"
	AdminOrgOrgOwnerBackfillResponseRequestStatusStatusCodeInvalidRequest             AdminOrgOrgOwnerBackfillResponseRequestStatusStatusCode = "INVALID_REQUEST"
	AdminOrgOrgOwnerBackfillResponseRequestStatusStatusCodeInvalidRequestVersion      AdminOrgOrgOwnerBackfillResponseRequestStatusStatusCode = "INVALID_REQUEST_VERSION"
	AdminOrgOrgOwnerBackfillResponseRequestStatusStatusCodeInvalidRequestData         AdminOrgOrgOwnerBackfillResponseRequestStatusStatusCode = "INVALID_REQUEST_DATA"
	AdminOrgOrgOwnerBackfillResponseRequestStatusStatusCodeMethodNotAllowed           AdminOrgOrgOwnerBackfillResponseRequestStatusStatusCode = "METHOD_NOT_ALLOWED"
	AdminOrgOrgOwnerBackfillResponseRequestStatusStatusCodeConflict                   AdminOrgOrgOwnerBackfillResponseRequestStatusStatusCode = "CONFLICT"
	AdminOrgOrgOwnerBackfillResponseRequestStatusStatusCodeUnprocessableEntity        AdminOrgOrgOwnerBackfillResponseRequestStatusStatusCode = "UNPROCESSABLE_ENTITY"
	AdminOrgOrgOwnerBackfillResponseRequestStatusStatusCodeTooManyRequests            AdminOrgOrgOwnerBackfillResponseRequestStatusStatusCode = "TOO_MANY_REQUESTS"
	AdminOrgOrgOwnerBackfillResponseRequestStatusStatusCodeInsufficientStorage        AdminOrgOrgOwnerBackfillResponseRequestStatusStatusCode = "INSUFFICIENT_STORAGE"
	AdminOrgOrgOwnerBackfillResponseRequestStatusStatusCodeServiceUnavailable         AdminOrgOrgOwnerBackfillResponseRequestStatusStatusCode = "SERVICE_UNAVAILABLE"
	AdminOrgOrgOwnerBackfillResponseRequestStatusStatusCodePayloadTooLarge            AdminOrgOrgOwnerBackfillResponseRequestStatusStatusCode = "PAYLOAD_TOO_LARGE"
	AdminOrgOrgOwnerBackfillResponseRequestStatusStatusCodeNotAcceptable              AdminOrgOrgOwnerBackfillResponseRequestStatusStatusCode = "NOT_ACCEPTABLE"
	AdminOrgOrgOwnerBackfillResponseRequestStatusStatusCodeUnavailableForLegalReasons AdminOrgOrgOwnerBackfillResponseRequestStatusStatusCode = "UNAVAILABLE_FOR_LEGAL_REASONS"
	AdminOrgOrgOwnerBackfillResponseRequestStatusStatusCodeBadGateway                 AdminOrgOrgOwnerBackfillResponseRequestStatusStatusCode = "BAD_GATEWAY"
)

func (r AdminOrgOrgOwnerBackfillResponseRequestStatusStatusCode) IsKnown() bool {
	switch r {
	case AdminOrgOrgOwnerBackfillResponseRequestStatusStatusCodeUnknown, AdminOrgOrgOwnerBackfillResponseRequestStatusStatusCodeSuccess, AdminOrgOrgOwnerBackfillResponseRequestStatusStatusCodeUnauthorized, AdminOrgOrgOwnerBackfillResponseRequestStatusStatusCodePaymentRequired, AdminOrgOrgOwnerBackfillResponseRequestStatusStatusCodeForbidden, AdminOrgOrgOwnerBackfillResponseRequestStatusStatusCodeTimeout, AdminOrgOrgOwnerBackfillResponseRequestStatusStatusCodeExists, AdminOrgOrgOwnerBackfillResponseRequestStatusStatusCodeNotFound, AdminOrgOrgOwnerBackfillResponseRequestStatusStatusCodeInternalError, AdminOrgOrgOwnerBackfillResponseRequestStatusStatusCodeInvalidRequest, AdminOrgOrgOwnerBackfillResponseRequestStatusStatusCodeInvalidRequestVersion, AdminOrgOrgOwnerBackfillResponseRequestStatusStatusCodeInvalidRequestData, AdminOrgOrgOwnerBackfillResponseRequestStatusStatusCodeMethodNotAllowed, AdminOrgOrgOwnerBackfillResponseRequestStatusStatusCodeConflict, AdminOrgOrgOwnerBackfillResponseRequestStatusStatusCodeUnprocessableEntity, AdminOrgOrgOwnerBackfillResponseRequestStatusStatusCodeTooManyRequests, AdminOrgOrgOwnerBackfillResponseRequestStatusStatusCodeInsufficientStorage, AdminOrgOrgOwnerBackfillResponseRequestStatusStatusCodeServiceUnavailable, AdminOrgOrgOwnerBackfillResponseRequestStatusStatusCodePayloadTooLarge, AdminOrgOrgOwnerBackfillResponseRequestStatusStatusCodeNotAcceptable, AdminOrgOrgOwnerBackfillResponseRequestStatusStatusCodeUnavailableForLegalReasons, AdminOrgOrgOwnerBackfillResponseRequestStatusStatusCodeBadGateway:
		return true
	}
	return false
}

type AdminOrgNewParams struct {
	// Org owner.
	OrgOwner param.Field[AdminOrgNewParamsOrgOwner] `json:"orgOwner,required"`
	// user country
	Country param.Field[string] `json:"country"`
	// optional description of the organization
	Description param.Field[string] `json:"description"`
	// Name of the organization that will be shown to users.
	DisplayName param.Field[string] `json:"displayName"`
	// Identity Provider ID.
	IdpID param.Field[string] `json:"idpId"`
	// Is NVIDIA internal org or not
	IsInternal param.Field[bool] `json:"isInternal"`
	// Organization name
	Name param.Field[string] `json:"name"`
	// product end customer name for enterprise(Fleet Command) product
	PecName param.Field[string] `json:"pecName"`
	// product end customer salesforce.com Id (external customer Id) for
	// enterprise(Fleet Command) product
	PecSfdcID          param.Field[string]                               `json:"pecSfdcId"`
	ProductEnablements param.Field[[]AdminOrgNewParamsProductEnablement] `json:"productEnablements"`
	// This should be deprecated, use productEnablements instead
	ProductSubscriptions param.Field[[]AdminOrgNewParamsProductSubscription] `json:"productSubscriptions"`
	// Company or organization industry
	SalesforceAccountIndustry param.Field[string] `json:"salesforceAccountIndustry"`
	// Send email to org owner or not. Default is true
	SendEmail param.Field[bool]                  `json:"sendEmail"`
	Type      param.Field[AdminOrgNewParamsType] `json:"type"`
}

func (r AdminOrgNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Org owner.
type AdminOrgNewParamsOrgOwner struct {
	// Email address of the org owner.
	Email param.Field[string] `json:"email,required"`
	// Org owner name.
	FullName param.Field[string] `json:"fullName,required"`
	// Last time the org owner logged in.
	LastLoginDate param.Field[string] `json:"lastLoginDate"`
}

func (r AdminOrgNewParamsOrgOwner) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Product Enablement
type AdminOrgNewParamsProductEnablement struct {
	// Product Name (NVAIE, BASE_COMMAND, REGISTRY, etc)
	ProductName param.Field[string] `json:"productName,required"`
	// Product Enablement Types
	Type param.Field[AdminOrgNewParamsProductEnablementsType] `json:"type,required"`
	// Date on which the subscription expires. The subscription is invalid after this
	// date. (yyyy-MM-dd)
	ExpirationDate param.Field[string]                                        `json:"expirationDate"`
	PoDetails      param.Field[[]AdminOrgNewParamsProductEnablementsPoDetail] `json:"poDetails"`
}

func (r AdminOrgNewParamsProductEnablement) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Product Enablement Types
type AdminOrgNewParamsProductEnablementsType string

const (
	AdminOrgNewParamsProductEnablementsTypeNgcAdminEval       AdminOrgNewParamsProductEnablementsType = "NGC_ADMIN_EVAL"
	AdminOrgNewParamsProductEnablementsTypeNgcAdminNfr        AdminOrgNewParamsProductEnablementsType = "NGC_ADMIN_NFR"
	AdminOrgNewParamsProductEnablementsTypeNgcAdminCommercial AdminOrgNewParamsProductEnablementsType = "NGC_ADMIN_COMMERCIAL"
	AdminOrgNewParamsProductEnablementsTypeEmsEval            AdminOrgNewParamsProductEnablementsType = "EMS_EVAL"
	AdminOrgNewParamsProductEnablementsTypeEmsNfr             AdminOrgNewParamsProductEnablementsType = "EMS_NFR"
	AdminOrgNewParamsProductEnablementsTypeEmsCommercial      AdminOrgNewParamsProductEnablementsType = "EMS_COMMERCIAL"
	AdminOrgNewParamsProductEnablementsTypeNgcAdminDeveloper  AdminOrgNewParamsProductEnablementsType = "NGC_ADMIN_DEVELOPER"
)

func (r AdminOrgNewParamsProductEnablementsType) IsKnown() bool {
	switch r {
	case AdminOrgNewParamsProductEnablementsTypeNgcAdminEval, AdminOrgNewParamsProductEnablementsTypeNgcAdminNfr, AdminOrgNewParamsProductEnablementsTypeNgcAdminCommercial, AdminOrgNewParamsProductEnablementsTypeEmsEval, AdminOrgNewParamsProductEnablementsTypeEmsNfr, AdminOrgNewParamsProductEnablementsTypeEmsCommercial, AdminOrgNewParamsProductEnablementsTypeNgcAdminDeveloper:
		return true
	}
	return false
}

// Purchase Order.
type AdminOrgNewParamsProductEnablementsPoDetail struct {
	// Entitlement identifier.
	EntitlementID param.Field[string] `json:"entitlementId"`
	// PAK (Product Activation Key) identifier.
	PkID param.Field[string] `json:"pkId"`
}

func (r AdminOrgNewParamsProductEnablementsPoDetail) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Product Subscription
type AdminOrgNewParamsProductSubscription struct {
	// Product Name (NVAIE, BASE_COMMAND, FleetCommand, REGISTRY, etc).
	ProductName param.Field[string] `json:"productName,required"`
	// Unique entitlement identifier
	ID param.Field[string] `json:"id"`
	// EMS Subscription type. (options: EMS_EVAL, EMS_NFR and EMS_COMMERCIAL)
	EmsEntitlementType param.Field[AdminOrgNewParamsProductSubscriptionsEmsEntitlementType] `json:"emsEntitlementType"`
	// Date on which the subscription expires. The subscription is invalid after this
	// date. (yyyy-MM-dd)
	ExpirationDate param.Field[string] `json:"expirationDate"`
	// Date on which the subscription becomes active. (yyyy-MM-dd)
	StartDate param.Field[string] `json:"startDate"`
	// Subscription type. (options: NGC_ADMIN_EVAL, NGC_ADMIN_NFR,
	// NGC_ADMIN_COMMERCIAL)
	Type param.Field[AdminOrgNewParamsProductSubscriptionsType] `json:"type"`
}

func (r AdminOrgNewParamsProductSubscription) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// EMS Subscription type. (options: EMS_EVAL, EMS_NFR and EMS_COMMERCIAL)
type AdminOrgNewParamsProductSubscriptionsEmsEntitlementType string

const (
	AdminOrgNewParamsProductSubscriptionsEmsEntitlementTypeEmsEval       AdminOrgNewParamsProductSubscriptionsEmsEntitlementType = "EMS_EVAL"
	AdminOrgNewParamsProductSubscriptionsEmsEntitlementTypeEmsNfr        AdminOrgNewParamsProductSubscriptionsEmsEntitlementType = "EMS_NFR"
	AdminOrgNewParamsProductSubscriptionsEmsEntitlementTypeEmsCommerical AdminOrgNewParamsProductSubscriptionsEmsEntitlementType = "EMS_COMMERICAL"
	AdminOrgNewParamsProductSubscriptionsEmsEntitlementTypeEmsCommercial AdminOrgNewParamsProductSubscriptionsEmsEntitlementType = "EMS_COMMERCIAL"
)

func (r AdminOrgNewParamsProductSubscriptionsEmsEntitlementType) IsKnown() bool {
	switch r {
	case AdminOrgNewParamsProductSubscriptionsEmsEntitlementTypeEmsEval, AdminOrgNewParamsProductSubscriptionsEmsEntitlementTypeEmsNfr, AdminOrgNewParamsProductSubscriptionsEmsEntitlementTypeEmsCommerical, AdminOrgNewParamsProductSubscriptionsEmsEntitlementTypeEmsCommercial:
		return true
	}
	return false
}

// Subscription type. (options: NGC_ADMIN_EVAL, NGC_ADMIN_NFR,
// NGC_ADMIN_COMMERCIAL)
type AdminOrgNewParamsProductSubscriptionsType string

const (
	AdminOrgNewParamsProductSubscriptionsTypeNgcAdminEval       AdminOrgNewParamsProductSubscriptionsType = "NGC_ADMIN_EVAL"
	AdminOrgNewParamsProductSubscriptionsTypeNgcAdminNfr        AdminOrgNewParamsProductSubscriptionsType = "NGC_ADMIN_NFR"
	AdminOrgNewParamsProductSubscriptionsTypeNgcAdminCommercial AdminOrgNewParamsProductSubscriptionsType = "NGC_ADMIN_COMMERCIAL"
)

func (r AdminOrgNewParamsProductSubscriptionsType) IsKnown() bool {
	switch r {
	case AdminOrgNewParamsProductSubscriptionsTypeNgcAdminEval, AdminOrgNewParamsProductSubscriptionsTypeNgcAdminNfr, AdminOrgNewParamsProductSubscriptionsTypeNgcAdminCommercial:
		return true
	}
	return false
}

type AdminOrgNewParamsType string

const (
	AdminOrgNewParamsTypeUnknown    AdminOrgNewParamsType = "UNKNOWN"
	AdminOrgNewParamsTypeCloud      AdminOrgNewParamsType = "CLOUD"
	AdminOrgNewParamsTypeEnterprise AdminOrgNewParamsType = "ENTERPRISE"
	AdminOrgNewParamsTypeIndividual AdminOrgNewParamsType = "INDIVIDUAL"
)

func (r AdminOrgNewParamsType) IsKnown() bool {
	switch r {
	case AdminOrgNewParamsTypeUnknown, AdminOrgNewParamsTypeCloud, AdminOrgNewParamsTypeEnterprise, AdminOrgNewParamsTypeIndividual:
		return true
	}
	return false
}

type AdminOrgUpdateParams struct {
	// Org Owner Alternate Contact
	AlternateContact param.Field[AdminOrgUpdateParamsAlternateContact] `json:"alternateContact"`
	// Name of the company
	CompanyName param.Field[string] `json:"companyName"`
	// optional description of the organization
	Description param.Field[string] `json:"description"`
	// Name of the organization that will be shown to users.
	DisplayName param.Field[string] `json:"displayName"`
	// Identity Provider ID.
	IdpID param.Field[string] `json:"idpId"`
	// Infinity manager setting definition
	InfinityManagerSettings param.Field[AdminOrgUpdateParamsInfinityManagerSettings] `json:"infinityManagerSettings"`
	// Dataset Service enable flag for an organization
	IsDatasetServiceEnabled param.Field[bool] `json:"isDatasetServiceEnabled"`
	// Is NVIDIA internal org or not
	IsInternal param.Field[bool] `json:"isInternal"`
	// Quick Start enable flag for an organization
	IsQuickStartEnabled param.Field[bool] `json:"isQuickStartEnabled"`
	// If a server side encryption is enabled for private registry (models, resources)
	IsRegistrySseEnabled param.Field[bool] `json:"isRegistrySSEEnabled"`
	// Secrets Manager Service enable flag for an organization
	IsSecretsManagerServiceEnabled param.Field[bool] `json:"isSecretsManagerServiceEnabled"`
	// Secure Credential Sharing Service enable flag for an organization
	IsSecureCredentialSharingServiceEnabled param.Field[bool] `json:"isSecureCredentialSharingServiceEnabled"`
	// If a separate influx db used for an organization in Base Command Platform job
	// telemetry
	IsSeparateInfluxDBUsed param.Field[bool] `json:"isSeparateInfluxDbUsed"`
	// Org owner.
	OrgOwner param.Field[AdminOrgUpdateParamsOrgOwner] `json:"orgOwner"`
	// Org owners
	OrgOwners param.Field[[]AdminOrgUpdateParamsOrgOwner] `json:"orgOwners"`
	// product end customer name for enterprise(Fleet Command) product
	PecName param.Field[string] `json:"pecName"`
	// product end customer salesforce.com Id (external customer Id) for
	// enterprise(Fleet Command) product
	PecSfdcID            param.Field[string]                                    `json:"pecSfdcId"`
	ProductEnablements   param.Field[[]AdminOrgUpdateParamsProductEnablement]   `json:"productEnablements"`
	ProductSubscriptions param.Field[[]AdminOrgUpdateParamsProductSubscription] `json:"productSubscriptions"`
	// Repo scan setting definition
	RepoScanSettings param.Field[AdminOrgUpdateParamsRepoScanSettings] `json:"repoScanSettings"`
	Type             param.Field[AdminOrgUpdateParamsType]             `json:"type"`
}

func (r AdminOrgUpdateParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Org Owner Alternate Contact
type AdminOrgUpdateParamsAlternateContact struct {
	// Alternate contact's email.
	Email param.Field[string] `json:"email"`
	// Full name of the alternate contact.
	FullName param.Field[string] `json:"fullName"`
}

func (r AdminOrgUpdateParamsAlternateContact) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Infinity manager setting definition
type AdminOrgUpdateParamsInfinityManagerSettings struct {
	// Enable the infinity manager or not. Used both in org and team level object
	InfinityManagerEnabled param.Field[bool] `json:"infinityManagerEnabled"`
	// Allow override settings at team level. Only used in org level object
	InfinityManagerEnableTeamOverride param.Field[bool] `json:"infinityManagerEnableTeamOverride"`
}

func (r AdminOrgUpdateParamsInfinityManagerSettings) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Org owner.
type AdminOrgUpdateParamsOrgOwner struct {
	// Email address of the org owner.
	Email param.Field[string] `json:"email,required"`
	// Org owner name.
	FullName param.Field[string] `json:"fullName,required"`
	// Last time the org owner logged in.
	LastLoginDate param.Field[string] `json:"lastLoginDate"`
}

func (r AdminOrgUpdateParamsOrgOwner) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Product Enablement
type AdminOrgUpdateParamsProductEnablement struct {
	// Product Name (NVAIE, BASE_COMMAND, REGISTRY, etc)
	ProductName param.Field[string] `json:"productName,required"`
	// Product Enablement Types
	Type param.Field[AdminOrgUpdateParamsProductEnablementsType] `json:"type,required"`
	// Date on which the subscription expires. The subscription is invalid after this
	// date. (yyyy-MM-dd)
	ExpirationDate param.Field[string]                                           `json:"expirationDate"`
	PoDetails      param.Field[[]AdminOrgUpdateParamsProductEnablementsPoDetail] `json:"poDetails"`
}

func (r AdminOrgUpdateParamsProductEnablement) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Product Enablement Types
type AdminOrgUpdateParamsProductEnablementsType string

const (
	AdminOrgUpdateParamsProductEnablementsTypeNgcAdminEval       AdminOrgUpdateParamsProductEnablementsType = "NGC_ADMIN_EVAL"
	AdminOrgUpdateParamsProductEnablementsTypeNgcAdminNfr        AdminOrgUpdateParamsProductEnablementsType = "NGC_ADMIN_NFR"
	AdminOrgUpdateParamsProductEnablementsTypeNgcAdminCommercial AdminOrgUpdateParamsProductEnablementsType = "NGC_ADMIN_COMMERCIAL"
	AdminOrgUpdateParamsProductEnablementsTypeEmsEval            AdminOrgUpdateParamsProductEnablementsType = "EMS_EVAL"
	AdminOrgUpdateParamsProductEnablementsTypeEmsNfr             AdminOrgUpdateParamsProductEnablementsType = "EMS_NFR"
	AdminOrgUpdateParamsProductEnablementsTypeEmsCommercial      AdminOrgUpdateParamsProductEnablementsType = "EMS_COMMERCIAL"
	AdminOrgUpdateParamsProductEnablementsTypeNgcAdminDeveloper  AdminOrgUpdateParamsProductEnablementsType = "NGC_ADMIN_DEVELOPER"
)

func (r AdminOrgUpdateParamsProductEnablementsType) IsKnown() bool {
	switch r {
	case AdminOrgUpdateParamsProductEnablementsTypeNgcAdminEval, AdminOrgUpdateParamsProductEnablementsTypeNgcAdminNfr, AdminOrgUpdateParamsProductEnablementsTypeNgcAdminCommercial, AdminOrgUpdateParamsProductEnablementsTypeEmsEval, AdminOrgUpdateParamsProductEnablementsTypeEmsNfr, AdminOrgUpdateParamsProductEnablementsTypeEmsCommercial, AdminOrgUpdateParamsProductEnablementsTypeNgcAdminDeveloper:
		return true
	}
	return false
}

// Purchase Order.
type AdminOrgUpdateParamsProductEnablementsPoDetail struct {
	// Entitlement identifier.
	EntitlementID param.Field[string] `json:"entitlementId"`
	// PAK (Product Activation Key) identifier.
	PkID param.Field[string] `json:"pkId"`
}

func (r AdminOrgUpdateParamsProductEnablementsPoDetail) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Product Subscription
type AdminOrgUpdateParamsProductSubscription struct {
	// Product Name (NVAIE, BASE_COMMAND, FleetCommand, REGISTRY, etc).
	ProductName param.Field[string] `json:"productName,required"`
	// Unique entitlement identifier
	ID param.Field[string] `json:"id"`
	// EMS Subscription type. (options: EMS_EVAL, EMS_NFR and EMS_COMMERCIAL)
	EmsEntitlementType param.Field[AdminOrgUpdateParamsProductSubscriptionsEmsEntitlementType] `json:"emsEntitlementType"`
	// Date on which the subscription expires. The subscription is invalid after this
	// date. (yyyy-MM-dd)
	ExpirationDate param.Field[string] `json:"expirationDate"`
	// Date on which the subscription becomes active. (yyyy-MM-dd)
	StartDate param.Field[string] `json:"startDate"`
	// Subscription type. (options: NGC_ADMIN_EVAL, NGC_ADMIN_NFR,
	// NGC_ADMIN_COMMERCIAL)
	Type param.Field[AdminOrgUpdateParamsProductSubscriptionsType] `json:"type"`
}

func (r AdminOrgUpdateParamsProductSubscription) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// EMS Subscription type. (options: EMS_EVAL, EMS_NFR and EMS_COMMERCIAL)
type AdminOrgUpdateParamsProductSubscriptionsEmsEntitlementType string

const (
	AdminOrgUpdateParamsProductSubscriptionsEmsEntitlementTypeEmsEval       AdminOrgUpdateParamsProductSubscriptionsEmsEntitlementType = "EMS_EVAL"
	AdminOrgUpdateParamsProductSubscriptionsEmsEntitlementTypeEmsNfr        AdminOrgUpdateParamsProductSubscriptionsEmsEntitlementType = "EMS_NFR"
	AdminOrgUpdateParamsProductSubscriptionsEmsEntitlementTypeEmsCommerical AdminOrgUpdateParamsProductSubscriptionsEmsEntitlementType = "EMS_COMMERICAL"
	AdminOrgUpdateParamsProductSubscriptionsEmsEntitlementTypeEmsCommercial AdminOrgUpdateParamsProductSubscriptionsEmsEntitlementType = "EMS_COMMERCIAL"
)

func (r AdminOrgUpdateParamsProductSubscriptionsEmsEntitlementType) IsKnown() bool {
	switch r {
	case AdminOrgUpdateParamsProductSubscriptionsEmsEntitlementTypeEmsEval, AdminOrgUpdateParamsProductSubscriptionsEmsEntitlementTypeEmsNfr, AdminOrgUpdateParamsProductSubscriptionsEmsEntitlementTypeEmsCommerical, AdminOrgUpdateParamsProductSubscriptionsEmsEntitlementTypeEmsCommercial:
		return true
	}
	return false
}

// Subscription type. (options: NGC_ADMIN_EVAL, NGC_ADMIN_NFR,
// NGC_ADMIN_COMMERCIAL)
type AdminOrgUpdateParamsProductSubscriptionsType string

const (
	AdminOrgUpdateParamsProductSubscriptionsTypeNgcAdminEval       AdminOrgUpdateParamsProductSubscriptionsType = "NGC_ADMIN_EVAL"
	AdminOrgUpdateParamsProductSubscriptionsTypeNgcAdminNfr        AdminOrgUpdateParamsProductSubscriptionsType = "NGC_ADMIN_NFR"
	AdminOrgUpdateParamsProductSubscriptionsTypeNgcAdminCommercial AdminOrgUpdateParamsProductSubscriptionsType = "NGC_ADMIN_COMMERCIAL"
)

func (r AdminOrgUpdateParamsProductSubscriptionsType) IsKnown() bool {
	switch r {
	case AdminOrgUpdateParamsProductSubscriptionsTypeNgcAdminEval, AdminOrgUpdateParamsProductSubscriptionsTypeNgcAdminNfr, AdminOrgUpdateParamsProductSubscriptionsTypeNgcAdminCommercial:
		return true
	}
	return false
}

// Repo scan setting definition
type AdminOrgUpdateParamsRepoScanSettings struct {
	// Allow org admin to override the org level repo scan settings
	RepoScanAllowOverride param.Field[bool] `json:"repoScanAllowOverride"`
	// Allow repository scanning by default
	RepoScanByDefault param.Field[bool] `json:"repoScanByDefault"`
	// Enable the repository scan or not. Only used in org level object
	RepoScanEnabled param.Field[bool] `json:"repoScanEnabled"`
	// Sends notification to end user after scanning is done
	RepoScanEnableNotifications param.Field[bool] `json:"repoScanEnableNotifications"`
	// Allow override settings at team level. Only used in org level object
	RepoScanEnableTeamOverride param.Field[bool] `json:"repoScanEnableTeamOverride"`
	// Allow showing scan results to CLI or UI
	RepoScanShowResults param.Field[bool] `json:"repoScanShowResults"`
}

func (r AdminOrgUpdateParamsRepoScanSettings) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type AdminOrgUpdateParamsType string

const (
	AdminOrgUpdateParamsTypeUnknown    AdminOrgUpdateParamsType = "UNKNOWN"
	AdminOrgUpdateParamsTypeCloud      AdminOrgUpdateParamsType = "CLOUD"
	AdminOrgUpdateParamsTypeEnterprise AdminOrgUpdateParamsType = "ENTERPRISE"
	AdminOrgUpdateParamsTypeIndividual AdminOrgUpdateParamsType = "INDIVIDUAL"
)

func (r AdminOrgUpdateParamsType) IsKnown() bool {
	switch r {
	case AdminOrgUpdateParamsTypeUnknown, AdminOrgUpdateParamsTypeCloud, AdminOrgUpdateParamsTypeEnterprise, AdminOrgUpdateParamsTypeIndividual:
		return true
	}
	return false
}

type AdminOrgEnableParams struct {
	// False only if called by SbMS.
	CreateSubscription param.Field[bool] `json:"createSubscription,required"`
	// Product Enablement
	ProductEnablement param.Field[AdminOrgEnableParamsProductEnablement] `json:"productEnablement,required"`
}

func (r AdminOrgEnableParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Product Enablement
type AdminOrgEnableParamsProductEnablement struct {
	// Product Name (NVAIE, BASE_COMMAND, REGISTRY, etc)
	ProductName param.Field[string] `json:"productName,required"`
	// Product Enablement Types
	Type param.Field[AdminOrgEnableParamsProductEnablementType] `json:"type,required"`
	// Date on which the subscription expires. The subscription is invalid after this
	// date. (yyyy-MM-dd)
	ExpirationDate param.Field[string]                                          `json:"expirationDate"`
	PoDetails      param.Field[[]AdminOrgEnableParamsProductEnablementPoDetail] `json:"poDetails"`
}

func (r AdminOrgEnableParamsProductEnablement) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Product Enablement Types
type AdminOrgEnableParamsProductEnablementType string

const (
	AdminOrgEnableParamsProductEnablementTypeNgcAdminEval       AdminOrgEnableParamsProductEnablementType = "NGC_ADMIN_EVAL"
	AdminOrgEnableParamsProductEnablementTypeNgcAdminNfr        AdminOrgEnableParamsProductEnablementType = "NGC_ADMIN_NFR"
	AdminOrgEnableParamsProductEnablementTypeNgcAdminCommercial AdminOrgEnableParamsProductEnablementType = "NGC_ADMIN_COMMERCIAL"
	AdminOrgEnableParamsProductEnablementTypeEmsEval            AdminOrgEnableParamsProductEnablementType = "EMS_EVAL"
	AdminOrgEnableParamsProductEnablementTypeEmsNfr             AdminOrgEnableParamsProductEnablementType = "EMS_NFR"
	AdminOrgEnableParamsProductEnablementTypeEmsCommercial      AdminOrgEnableParamsProductEnablementType = "EMS_COMMERCIAL"
	AdminOrgEnableParamsProductEnablementTypeNgcAdminDeveloper  AdminOrgEnableParamsProductEnablementType = "NGC_ADMIN_DEVELOPER"
)

func (r AdminOrgEnableParamsProductEnablementType) IsKnown() bool {
	switch r {
	case AdminOrgEnableParamsProductEnablementTypeNgcAdminEval, AdminOrgEnableParamsProductEnablementTypeNgcAdminNfr, AdminOrgEnableParamsProductEnablementTypeNgcAdminCommercial, AdminOrgEnableParamsProductEnablementTypeEmsEval, AdminOrgEnableParamsProductEnablementTypeEmsNfr, AdminOrgEnableParamsProductEnablementTypeEmsCommercial, AdminOrgEnableParamsProductEnablementTypeNgcAdminDeveloper:
		return true
	}
	return false
}

// Purchase Order.
type AdminOrgEnableParamsProductEnablementPoDetail struct {
	// Entitlement identifier.
	EntitlementID param.Field[string] `json:"entitlementId"`
	// PAK (Product Activation Key) identifier.
	PkID param.Field[string] `json:"pkId"`
}

func (r AdminOrgEnableParamsProductEnablementPoDetail) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}
