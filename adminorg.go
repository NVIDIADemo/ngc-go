// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package ngc

import (
	"context"
	"net/http"
	"net/url"

	"github.com/NVIDIADemo/ngc-go/internal/apijson"
	"github.com/NVIDIADemo/ngc-go/internal/apiquery"
	"github.com/NVIDIADemo/ngc-go/internal/pagination"
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
	Options    []option.RequestOption
	NcaIDs     *AdminOrgNcaIDService
	Offboarded *AdminOrgOffboardedService
	Teams      *AdminOrgTeamService
	Users      *AdminOrgUserService
}

// NewAdminOrgService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewAdminOrgService(opts ...option.RequestOption) (r *AdminOrgService) {
	r = &AdminOrgService{}
	r.Options = opts
	r.NcaIDs = NewAdminOrgNcaIDService(opts...)
	r.Offboarded = NewAdminOrgOffboardedService(opts...)
	r.Teams = NewAdminOrgTeamService(opts...)
	r.Users = NewAdminOrgUserService(opts...)
	return
}

// OrgCreateRequest is used to create the organization or when no nca_id is
// provided upfront, the OrgCreateRequest is stored as proto org, and proto org
// flow initiates (SuperAdmin privileges required)
func (r *AdminOrgService) New(ctx context.Context, params AdminOrgNewParams, opts ...option.RequestOption) (res *http.Response, err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := "v3/admin/orgs"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// List all organizations. (SuperAdmin privileges required)
func (r *AdminOrgService) List(ctx context.Context, query AdminOrgListParams, opts ...option.RequestOption) (res *pagination.PageNumberOrganizations[AdminOrgListResponse], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "v2/admin/org"
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodGet, path, query, &res, opts...)
	if err != nil {
		return nil, err
	}
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

// List all organizations. (SuperAdmin privileges required)
func (r *AdminOrgService) ListAutoPaging(ctx context.Context, query AdminOrgListParams, opts ...option.RequestOption) *pagination.PageNumberOrganizationsAutoPager[AdminOrgListResponse] {
	return pagination.NewPageNumberOrganizationsAutoPager(r.List(ctx, query, opts...))
}

// Information about the Organization
type AdminOrgListResponse struct {
	// Unique Id of this team.
	ID int64 `json:"id"`
	// Org Owner Alternate Contact
	AlternateContact AdminOrgListResponseAlternateContact `json:"alternateContact"`
	// Billing account ID.
	BillingAccountID string `json:"billingAccountId"`
	// Identifies if the org can be reused.
	CanAddOn bool `json:"canAddOn"`
	// ISO country code of the organization.
	Country string `json:"country"`
	// Optional description of the organization.
	Description string `json:"description"`
	// Name of the organization that will be shown to users.
	DisplayName string `json:"displayName"`
	// Identity Provider ID.
	IdpID string `json:"idpId"`
	// Industry of the organization.
	Industry string `json:"industry"`
	// Infinity manager setting definition
	InfinityManagerSettings AdminOrgListResponseInfinityManagerSettings `json:"infinityManagerSettings"`
	// Dataset Service enable flag for an organization
	IsDatasetServiceEnabled bool `json:"isDatasetServiceEnabled"`
	// Is NVIDIA internal org or not
	IsInternal bool `json:"isInternal"`
	// Indicates when the org is a proto org
	IsProto bool `json:"isProto"`
	// Quick Start enable flag for an organization
	IsQuickStartEnabled bool `json:"isQuickStartEnabled"`
	// If a server side encryption is enabled for private registry (models, resources)
	IsRegistrySseEnabled bool `json:"isRegistrySSEEnabled"`
	// Secrets Manager Service enable flag for an organization
	IsSecretsManagerServiceEnabled bool `json:"isSecretsManagerServiceEnabled"`
	// Secure Credential Sharing Service enable flag for an organization
	IsSecureCredentialSharingServiceEnabled bool `json:"isSecureCredentialSharingServiceEnabled"`
	// If a separate influx db used for an organization in BCP for job telemetry
	IsSeparateInfluxDBUsed bool `json:"isSeparateInfluxDbUsed"`
	// Organization name.
	Name string `json:"name"`
	// NVIDIA Cloud Account Number.
	Nan string `json:"nan"`
	// Org owner.
	OrgOwner AdminOrgListResponseOrgOwner `json:"orgOwner"`
	// Org owners
	OrgOwners []AdminOrgListResponseOrgOwner `json:"orgOwners"`
	// Product end customer salesforce.com Id (external customer Id). pecSfdcId is for
	// EMS (entitlement management service) to track external paid customer.
	PecSfdcID            string                                    `json:"pecSfdcId"`
	ProductEnablements   []AdminOrgListResponseProductEnablement   `json:"productEnablements"`
	ProductSubscriptions []AdminOrgListResponseProductSubscription `json:"productSubscriptions"`
	// Repo scan setting definition
	RepoScanSettings AdminOrgListResponseRepoScanSettings `json:"repoScanSettings"`
	Type             AdminOrgListResponseType             `json:"type"`
	// Users information.
	UsersInfo AdminOrgListResponseUsersInfo `json:"usersInfo"`
	JSON      adminOrgListResponseJSON      `json:"-"`
}

// adminOrgListResponseJSON contains the JSON metadata for the struct
// [AdminOrgListResponse]
type adminOrgListResponseJSON struct {
	ID                                      apijson.Field
	AlternateContact                        apijson.Field
	BillingAccountID                        apijson.Field
	CanAddOn                                apijson.Field
	Country                                 apijson.Field
	Description                             apijson.Field
	DisplayName                             apijson.Field
	IdpID                                   apijson.Field
	Industry                                apijson.Field
	InfinityManagerSettings                 apijson.Field
	IsDatasetServiceEnabled                 apijson.Field
	IsInternal                              apijson.Field
	IsProto                                 apijson.Field
	IsQuickStartEnabled                     apijson.Field
	IsRegistrySseEnabled                    apijson.Field
	IsSecretsManagerServiceEnabled          apijson.Field
	IsSecureCredentialSharingServiceEnabled apijson.Field
	IsSeparateInfluxDBUsed                  apijson.Field
	Name                                    apijson.Field
	Nan                                     apijson.Field
	OrgOwner                                apijson.Field
	OrgOwners                               apijson.Field
	PecSfdcID                               apijson.Field
	ProductEnablements                      apijson.Field
	ProductSubscriptions                    apijson.Field
	RepoScanSettings                        apijson.Field
	Type                                    apijson.Field
	UsersInfo                               apijson.Field
	raw                                     string
	ExtraFields                             map[string]apijson.Field
}

func (r *AdminOrgListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r adminOrgListResponseJSON) RawJSON() string {
	return r.raw
}

// Org Owner Alternate Contact
type AdminOrgListResponseAlternateContact struct {
	// Alternate contact's email.
	Email string `json:"email"`
	// Full name of the alternate contact.
	FullName string                                   `json:"fullName"`
	JSON     adminOrgListResponseAlternateContactJSON `json:"-"`
}

// adminOrgListResponseAlternateContactJSON contains the JSON metadata for the
// struct [AdminOrgListResponseAlternateContact]
type adminOrgListResponseAlternateContactJSON struct {
	Email       apijson.Field
	FullName    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AdminOrgListResponseAlternateContact) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r adminOrgListResponseAlternateContactJSON) RawJSON() string {
	return r.raw
}

// Infinity manager setting definition
type AdminOrgListResponseInfinityManagerSettings struct {
	// Enable the infinity manager or not. Used both in org and team level object
	InfinityManagerEnabled bool `json:"infinityManagerEnabled"`
	// Allow override settings at team level. Only used in org level object
	InfinityManagerEnableTeamOverride bool                                            `json:"infinityManagerEnableTeamOverride"`
	JSON                              adminOrgListResponseInfinityManagerSettingsJSON `json:"-"`
}

// adminOrgListResponseInfinityManagerSettingsJSON contains the JSON metadata for
// the struct [AdminOrgListResponseInfinityManagerSettings]
type adminOrgListResponseInfinityManagerSettingsJSON struct {
	InfinityManagerEnabled            apijson.Field
	InfinityManagerEnableTeamOverride apijson.Field
	raw                               string
	ExtraFields                       map[string]apijson.Field
}

func (r *AdminOrgListResponseInfinityManagerSettings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r adminOrgListResponseInfinityManagerSettingsJSON) RawJSON() string {
	return r.raw
}

// Org owner.
type AdminOrgListResponseOrgOwner struct {
	// Email address of the org owner.
	Email string `json:"email,required"`
	// Org owner name.
	FullName string `json:"fullName,required"`
	// Last time the org owner logged in.
	LastLoginDate string                           `json:"lastLoginDate"`
	JSON          adminOrgListResponseOrgOwnerJSON `json:"-"`
}

// adminOrgListResponseOrgOwnerJSON contains the JSON metadata for the struct
// [AdminOrgListResponseOrgOwner]
type adminOrgListResponseOrgOwnerJSON struct {
	Email         apijson.Field
	FullName      apijson.Field
	LastLoginDate apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *AdminOrgListResponseOrgOwner) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r adminOrgListResponseOrgOwnerJSON) RawJSON() string {
	return r.raw
}

// Product Enablement
type AdminOrgListResponseProductEnablement struct {
	// Product Name (NVAIE, BASE_COMMAND, REGISTRY, etc)
	ProductName string `json:"productName,required"`
	// Product Enablement Types
	Type AdminOrgListResponseProductEnablementsType `json:"type,required"`
	// Date on which the subscription expires. The subscription is invalid after this
	// date. (yyyy-MM-dd)
	ExpirationDate string                                           `json:"expirationDate"`
	PoDetails      []AdminOrgListResponseProductEnablementsPoDetail `json:"poDetails"`
	JSON           adminOrgListResponseProductEnablementJSON        `json:"-"`
}

// adminOrgListResponseProductEnablementJSON contains the JSON metadata for the
// struct [AdminOrgListResponseProductEnablement]
type adminOrgListResponseProductEnablementJSON struct {
	ProductName    apijson.Field
	Type           apijson.Field
	ExpirationDate apijson.Field
	PoDetails      apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *AdminOrgListResponseProductEnablement) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r adminOrgListResponseProductEnablementJSON) RawJSON() string {
	return r.raw
}

// Product Enablement Types
type AdminOrgListResponseProductEnablementsType string

const (
	AdminOrgListResponseProductEnablementsTypeNgcAdminEval       AdminOrgListResponseProductEnablementsType = "NGC_ADMIN_EVAL"
	AdminOrgListResponseProductEnablementsTypeNgcAdminNfr        AdminOrgListResponseProductEnablementsType = "NGC_ADMIN_NFR"
	AdminOrgListResponseProductEnablementsTypeNgcAdminCommercial AdminOrgListResponseProductEnablementsType = "NGC_ADMIN_COMMERCIAL"
	AdminOrgListResponseProductEnablementsTypeEmsEval            AdminOrgListResponseProductEnablementsType = "EMS_EVAL"
	AdminOrgListResponseProductEnablementsTypeEmsNfr             AdminOrgListResponseProductEnablementsType = "EMS_NFR"
	AdminOrgListResponseProductEnablementsTypeEmsCommercial      AdminOrgListResponseProductEnablementsType = "EMS_COMMERCIAL"
	AdminOrgListResponseProductEnablementsTypeNgcAdminDeveloper  AdminOrgListResponseProductEnablementsType = "NGC_ADMIN_DEVELOPER"
)

func (r AdminOrgListResponseProductEnablementsType) IsKnown() bool {
	switch r {
	case AdminOrgListResponseProductEnablementsTypeNgcAdminEval, AdminOrgListResponseProductEnablementsTypeNgcAdminNfr, AdminOrgListResponseProductEnablementsTypeNgcAdminCommercial, AdminOrgListResponseProductEnablementsTypeEmsEval, AdminOrgListResponseProductEnablementsTypeEmsNfr, AdminOrgListResponseProductEnablementsTypeEmsCommercial, AdminOrgListResponseProductEnablementsTypeNgcAdminDeveloper:
		return true
	}
	return false
}

// Purchase Order.
type AdminOrgListResponseProductEnablementsPoDetail struct {
	// Entitlement identifier.
	EntitlementID string `json:"entitlementId"`
	// PAK (Product Activation Key) identifier.
	PkID string                                             `json:"pkId"`
	JSON adminOrgListResponseProductEnablementsPoDetailJSON `json:"-"`
}

// adminOrgListResponseProductEnablementsPoDetailJSON contains the JSON metadata
// for the struct [AdminOrgListResponseProductEnablementsPoDetail]
type adminOrgListResponseProductEnablementsPoDetailJSON struct {
	EntitlementID apijson.Field
	PkID          apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *AdminOrgListResponseProductEnablementsPoDetail) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r adminOrgListResponseProductEnablementsPoDetailJSON) RawJSON() string {
	return r.raw
}

// Product Subscription
type AdminOrgListResponseProductSubscription struct {
	// Product Name (NVAIE, BASE_COMMAND, FleetCommand, REGISTRY, etc).
	ProductName string `json:"productName,required"`
	// Unique entitlement identifier
	ID string `json:"id"`
	// EMS Subscription type. (options: EMS_EVAL, EMS_NFR and EMS_COMMERCIAL)
	EmsEntitlementType AdminOrgListResponseProductSubscriptionsEmsEntitlementType `json:"emsEntitlementType"`
	// Date on which the subscription expires. The subscription is invalid after this
	// date. (yyyy-MM-dd)
	ExpirationDate string `json:"expirationDate"`
	// Date on which the subscription becomes active. (yyyy-MM-dd)
	StartDate string `json:"startDate"`
	// Subscription type. (options: NGC_ADMIN_EVAL, NGC_ADMIN_NFR,
	// NGC_ADMIN_COMMERCIAL)
	Type AdminOrgListResponseProductSubscriptionsType `json:"type"`
	JSON adminOrgListResponseProductSubscriptionJSON  `json:"-"`
}

// adminOrgListResponseProductSubscriptionJSON contains the JSON metadata for the
// struct [AdminOrgListResponseProductSubscription]
type adminOrgListResponseProductSubscriptionJSON struct {
	ProductName        apijson.Field
	ID                 apijson.Field
	EmsEntitlementType apijson.Field
	ExpirationDate     apijson.Field
	StartDate          apijson.Field
	Type               apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *AdminOrgListResponseProductSubscription) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r adminOrgListResponseProductSubscriptionJSON) RawJSON() string {
	return r.raw
}

// EMS Subscription type. (options: EMS_EVAL, EMS_NFR and EMS_COMMERCIAL)
type AdminOrgListResponseProductSubscriptionsEmsEntitlementType string

const (
	AdminOrgListResponseProductSubscriptionsEmsEntitlementTypeEmsEval       AdminOrgListResponseProductSubscriptionsEmsEntitlementType = "EMS_EVAL"
	AdminOrgListResponseProductSubscriptionsEmsEntitlementTypeEmsNfr        AdminOrgListResponseProductSubscriptionsEmsEntitlementType = "EMS_NFR"
	AdminOrgListResponseProductSubscriptionsEmsEntitlementTypeEmsCommerical AdminOrgListResponseProductSubscriptionsEmsEntitlementType = "EMS_COMMERICAL"
	AdminOrgListResponseProductSubscriptionsEmsEntitlementTypeEmsCommercial AdminOrgListResponseProductSubscriptionsEmsEntitlementType = "EMS_COMMERCIAL"
)

func (r AdminOrgListResponseProductSubscriptionsEmsEntitlementType) IsKnown() bool {
	switch r {
	case AdminOrgListResponseProductSubscriptionsEmsEntitlementTypeEmsEval, AdminOrgListResponseProductSubscriptionsEmsEntitlementTypeEmsNfr, AdminOrgListResponseProductSubscriptionsEmsEntitlementTypeEmsCommerical, AdminOrgListResponseProductSubscriptionsEmsEntitlementTypeEmsCommercial:
		return true
	}
	return false
}

// Subscription type. (options: NGC_ADMIN_EVAL, NGC_ADMIN_NFR,
// NGC_ADMIN_COMMERCIAL)
type AdminOrgListResponseProductSubscriptionsType string

const (
	AdminOrgListResponseProductSubscriptionsTypeNgcAdminEval       AdminOrgListResponseProductSubscriptionsType = "NGC_ADMIN_EVAL"
	AdminOrgListResponseProductSubscriptionsTypeNgcAdminNfr        AdminOrgListResponseProductSubscriptionsType = "NGC_ADMIN_NFR"
	AdminOrgListResponseProductSubscriptionsTypeNgcAdminCommercial AdminOrgListResponseProductSubscriptionsType = "NGC_ADMIN_COMMERCIAL"
)

func (r AdminOrgListResponseProductSubscriptionsType) IsKnown() bool {
	switch r {
	case AdminOrgListResponseProductSubscriptionsTypeNgcAdminEval, AdminOrgListResponseProductSubscriptionsTypeNgcAdminNfr, AdminOrgListResponseProductSubscriptionsTypeNgcAdminCommercial:
		return true
	}
	return false
}

// Repo scan setting definition
type AdminOrgListResponseRepoScanSettings struct {
	// Allow org admin to override the org level repo scan settings
	RepoScanAllowOverride bool `json:"repoScanAllowOverride"`
	// Allow repository scanning by default
	RepoScanByDefault bool `json:"repoScanByDefault"`
	// Enable the repository scan or not. Only used in org level object
	RepoScanEnabled bool `json:"repoScanEnabled"`
	// Sends notification to end user after scanning is done
	RepoScanEnableNotifications bool `json:"repoScanEnableNotifications"`
	// Allow override settings at team level. Only used in org level object
	RepoScanEnableTeamOverride bool `json:"repoScanEnableTeamOverride"`
	// Allow showing scan results to CLI or UI
	RepoScanShowResults bool                                     `json:"repoScanShowResults"`
	JSON                adminOrgListResponseRepoScanSettingsJSON `json:"-"`
}

// adminOrgListResponseRepoScanSettingsJSON contains the JSON metadata for the
// struct [AdminOrgListResponseRepoScanSettings]
type adminOrgListResponseRepoScanSettingsJSON struct {
	RepoScanAllowOverride       apijson.Field
	RepoScanByDefault           apijson.Field
	RepoScanEnabled             apijson.Field
	RepoScanEnableNotifications apijson.Field
	RepoScanEnableTeamOverride  apijson.Field
	RepoScanShowResults         apijson.Field
	raw                         string
	ExtraFields                 map[string]apijson.Field
}

func (r *AdminOrgListResponseRepoScanSettings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r adminOrgListResponseRepoScanSettingsJSON) RawJSON() string {
	return r.raw
}

type AdminOrgListResponseType string

const (
	AdminOrgListResponseTypeUnknown    AdminOrgListResponseType = "UNKNOWN"
	AdminOrgListResponseTypeCloud      AdminOrgListResponseType = "CLOUD"
	AdminOrgListResponseTypeEnterprise AdminOrgListResponseType = "ENTERPRISE"
	AdminOrgListResponseTypeIndividual AdminOrgListResponseType = "INDIVIDUAL"
)

func (r AdminOrgListResponseType) IsKnown() bool {
	switch r {
	case AdminOrgListResponseTypeUnknown, AdminOrgListResponseTypeCloud, AdminOrgListResponseTypeEnterprise, AdminOrgListResponseTypeIndividual:
		return true
	}
	return false
}

// Users information.
type AdminOrgListResponseUsersInfo struct {
	// Total number of users.
	TotalUsers int64                             `json:"totalUsers"`
	JSON       adminOrgListResponseUsersInfoJSON `json:"-"`
}

// adminOrgListResponseUsersInfoJSON contains the JSON metadata for the struct
// [AdminOrgListResponseUsersInfo]
type adminOrgListResponseUsersInfoJSON struct {
	TotalUsers  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AdminOrgListResponseUsersInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r adminOrgListResponseUsersInfoJSON) RawJSON() string {
	return r.raw
}

type AdminOrgNewParams struct {
	// user country
	Country param.Field[string] `json:"country"`
	// optional description of the organization
	Description param.Field[string] `json:"description"`
	// Name of the organization that will be shown to users.
	DisplayName param.Field[string] `json:"displayName"`
	// Identify the initiator of the org request
	Initiator param.Field[string] `json:"initiator"`
	// Is NVIDIA internal org or not
	IsInternal param.Field[bool] `json:"isInternal"`
	// Organization name
	Name param.Field[string] `json:"name"`
	// NVIDIA Cloud Account Identifier
	NcaID param.Field[string] `json:"ncaId"`
	// NVIDIA Cloud Account Number
	NcaNumber param.Field[string] `json:"ncaNumber"`
	// Org owner.
	OrgOwner param.Field[AdminOrgNewParamsOrgOwner] `json:"orgOwner"`
	// product end customer name for enterprise(Fleet Command) product
	PecName param.Field[string] `json:"pecName"`
	// product end customer salesforce.com Id (external customer Id) for
	// enterprise(Fleet Command) product
	PecSfdcID          param.Field[string]                               `json:"pecSfdcId"`
	ProductEnablements param.Field[[]AdminOrgNewParamsProductEnablement] `json:"productEnablements"`
	// This should be deprecated, use productEnablements instead
	ProductSubscriptions param.Field[[]AdminOrgNewParamsProductSubscription] `json:"productSubscriptions"`
	// Proto org identifier
	ProtoOrgID param.Field[string] `json:"protoOrgId"`
	// Company or organization industry
	SalesforceAccountIndustry param.Field[string] `json:"salesforceAccountIndustry"`
	// Send email to org owner or not. Default is true
	SendEmail param.Field[bool]                  `json:"sendEmail"`
	Type      param.Field[AdminOrgNewParamsType] `json:"type"`
	Ncid      param.Field[string]                `cookie:"ncid"`
	VisitorID param.Field[string]                `cookie:"VisitorID"`
}

func (r AdminOrgNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Org owner.
type AdminOrgNewParamsOrgOwner struct {
	// Email address of the org owner.
	Email param.Field[string] `json:"email,required"`
	// Org owner name.
	FullName param.Field[string] `json:"fullName"`
	// Identity Provider ID of the org owner.
	IdpID param.Field[string] `json:"idpId"`
	// Starfleet ID of the org owner.
	StarfleetID param.Field[string] `json:"starfleetId"`
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

type AdminOrgListParams struct {
	FilterUsingOrgDisplayName param.Field[string]                                     `query:"Filter using org display name"`
	FilterUsingOrgOwnerEmail  param.Field[AdminOrgListParamsFilterUsingOrgOwnerEmail] `query:"Filter using org owner email"`
	FilterUsingOrgOwnerName   param.Field[string]                                     `query:"Filter using org owner name"`
	// Org description to search
	OrgDesc param.Field[string] `query:"org-desc"`
	// Org name to search
	OrgName param.Field[string] `query:"org-name"`
	// Org Type to search
	OrgType param.Field[AdminOrgListParamsOrgType] `query:"org-type"`
	// The page number of result
	PageNumber param.Field[int64] `query:"page-number"`
	// The page size of result
	PageSize param.Field[int64] `query:"page-size"`
	// PEC ID to search
	PecID param.Field[string] `query:"pec-id"`
}

// URLQuery serializes [AdminOrgListParams]'s query parameters as `url.Values`.
func (r AdminOrgListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type AdminOrgListParamsFilterUsingOrgOwnerEmail struct {
	EmailShouldBeBase64Encoded param.Field[string] `query:" Email should be base-64-encoded"`
}

// URLQuery serializes [AdminOrgListParamsFilterUsingOrgOwnerEmail]'s query
// parameters as `url.Values`.
func (r AdminOrgListParamsFilterUsingOrgOwnerEmail) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Org Type to search
type AdminOrgListParamsOrgType string

const (
	AdminOrgListParamsOrgTypeUnknown    AdminOrgListParamsOrgType = "UNKNOWN"
	AdminOrgListParamsOrgTypeCloud      AdminOrgListParamsOrgType = "CLOUD"
	AdminOrgListParamsOrgTypeEnterprise AdminOrgListParamsOrgType = "ENTERPRISE"
	AdminOrgListParamsOrgTypeIndividual AdminOrgListParamsOrgType = "INDIVIDUAL"
)

func (r AdminOrgListParamsOrgType) IsKnown() bool {
	switch r {
	case AdminOrgListParamsOrgTypeUnknown, AdminOrgListParamsOrgTypeCloud, AdminOrgListParamsOrgTypeEnterprise, AdminOrgListParamsOrgTypeIndividual:
		return true
	}
	return false
}
