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

// AdminOrgOffboardedService contains methods and other services that help with
// interacting with the ngc API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAdminOrgOffboardedService] method instead.
type AdminOrgOffboardedService struct {
	Options []option.RequestOption
}

// NewAdminOrgOffboardedService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAdminOrgOffboardedService(opts ...option.RequestOption) (r *AdminOrgOffboardedService) {
	r = &AdminOrgOffboardedService{}
	r.Options = opts
	return
}

// List all Offboarded organizations.
func (r *AdminOrgOffboardedService) List(ctx context.Context, query AdminOrgOffboardedListParams, opts ...option.RequestOption) (res *pagination.PageNumberOrganizations[AdminOrgOffboardedListResponse], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "v2/admin/orgs/offboarded"
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

// List all Offboarded organizations.
func (r *AdminOrgOffboardedService) ListAutoPaging(ctx context.Context, query AdminOrgOffboardedListParams, opts ...option.RequestOption) *pagination.PageNumberOrganizationsAutoPager[AdminOrgOffboardedListResponse] {
	return pagination.NewPageNumberOrganizationsAutoPager(r.List(ctx, query, opts...))
}

// Information about the Organization
type AdminOrgOffboardedListResponse struct {
	// Unique Id of this team.
	ID int64 `json:"id"`
	// Org Owner Alternate Contact
	AlternateContact AdminOrgOffboardedListResponseAlternateContact `json:"alternateContact"`
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
	InfinityManagerSettings AdminOrgOffboardedListResponseInfinityManagerSettings `json:"infinityManagerSettings"`
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
	OrgOwner AdminOrgOffboardedListResponseOrgOwner `json:"orgOwner"`
	// Org owners
	OrgOwners []AdminOrgOffboardedListResponseOrgOwner `json:"orgOwners"`
	// Product end customer salesforce.com Id (external customer Id). pecSfdcId is for
	// EMS (entitlement management service) to track external paid customer.
	PecSfdcID            string                                              `json:"pecSfdcId"`
	ProductEnablements   []AdminOrgOffboardedListResponseProductEnablement   `json:"productEnablements"`
	ProductSubscriptions []AdminOrgOffboardedListResponseProductSubscription `json:"productSubscriptions"`
	// Repo scan setting definition
	RepoScanSettings AdminOrgOffboardedListResponseRepoScanSettings `json:"repoScanSettings"`
	Type             AdminOrgOffboardedListResponseType             `json:"type"`
	// Users information.
	UsersInfo AdminOrgOffboardedListResponseUsersInfo `json:"usersInfo"`
	JSON      adminOrgOffboardedListResponseJSON      `json:"-"`
}

// adminOrgOffboardedListResponseJSON contains the JSON metadata for the struct
// [AdminOrgOffboardedListResponse]
type adminOrgOffboardedListResponseJSON struct {
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

func (r *AdminOrgOffboardedListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r adminOrgOffboardedListResponseJSON) RawJSON() string {
	return r.raw
}

// Org Owner Alternate Contact
type AdminOrgOffboardedListResponseAlternateContact struct {
	// Alternate contact's email.
	Email string `json:"email"`
	// Full name of the alternate contact.
	FullName string                                             `json:"fullName"`
	JSON     adminOrgOffboardedListResponseAlternateContactJSON `json:"-"`
}

// adminOrgOffboardedListResponseAlternateContactJSON contains the JSON metadata
// for the struct [AdminOrgOffboardedListResponseAlternateContact]
type adminOrgOffboardedListResponseAlternateContactJSON struct {
	Email       apijson.Field
	FullName    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AdminOrgOffboardedListResponseAlternateContact) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r adminOrgOffboardedListResponseAlternateContactJSON) RawJSON() string {
	return r.raw
}

// Infinity manager setting definition
type AdminOrgOffboardedListResponseInfinityManagerSettings struct {
	// Enable the infinity manager or not. Used both in org and team level object
	InfinityManagerEnabled bool `json:"infinityManagerEnabled"`
	// Allow override settings at team level. Only used in org level object
	InfinityManagerEnableTeamOverride bool                                                      `json:"infinityManagerEnableTeamOverride"`
	JSON                              adminOrgOffboardedListResponseInfinityManagerSettingsJSON `json:"-"`
}

// adminOrgOffboardedListResponseInfinityManagerSettingsJSON contains the JSON
// metadata for the struct [AdminOrgOffboardedListResponseInfinityManagerSettings]
type adminOrgOffboardedListResponseInfinityManagerSettingsJSON struct {
	InfinityManagerEnabled            apijson.Field
	InfinityManagerEnableTeamOverride apijson.Field
	raw                               string
	ExtraFields                       map[string]apijson.Field
}

func (r *AdminOrgOffboardedListResponseInfinityManagerSettings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r adminOrgOffboardedListResponseInfinityManagerSettingsJSON) RawJSON() string {
	return r.raw
}

// Org owner.
type AdminOrgOffboardedListResponseOrgOwner struct {
	// Email address of the org owner.
	Email string `json:"email,required"`
	// Org owner name.
	FullName string `json:"fullName,required"`
	// Last time the org owner logged in.
	LastLoginDate string                                     `json:"lastLoginDate"`
	JSON          adminOrgOffboardedListResponseOrgOwnerJSON `json:"-"`
}

// adminOrgOffboardedListResponseOrgOwnerJSON contains the JSON metadata for the
// struct [AdminOrgOffboardedListResponseOrgOwner]
type adminOrgOffboardedListResponseOrgOwnerJSON struct {
	Email         apijson.Field
	FullName      apijson.Field
	LastLoginDate apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *AdminOrgOffboardedListResponseOrgOwner) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r adminOrgOffboardedListResponseOrgOwnerJSON) RawJSON() string {
	return r.raw
}

// Product Enablement
type AdminOrgOffboardedListResponseProductEnablement struct {
	// Product Name (NVAIE, BASE_COMMAND, REGISTRY, etc)
	ProductName string `json:"productName,required"`
	// Product Enablement Types
	Type AdminOrgOffboardedListResponseProductEnablementsType `json:"type,required"`
	// Date on which the subscription expires. The subscription is invalid after this
	// date. (yyyy-MM-dd)
	ExpirationDate string                                                     `json:"expirationDate"`
	PoDetails      []AdminOrgOffboardedListResponseProductEnablementsPoDetail `json:"poDetails"`
	JSON           adminOrgOffboardedListResponseProductEnablementJSON        `json:"-"`
}

// adminOrgOffboardedListResponseProductEnablementJSON contains the JSON metadata
// for the struct [AdminOrgOffboardedListResponseProductEnablement]
type adminOrgOffboardedListResponseProductEnablementJSON struct {
	ProductName    apijson.Field
	Type           apijson.Field
	ExpirationDate apijson.Field
	PoDetails      apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *AdminOrgOffboardedListResponseProductEnablement) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r adminOrgOffboardedListResponseProductEnablementJSON) RawJSON() string {
	return r.raw
}

// Product Enablement Types
type AdminOrgOffboardedListResponseProductEnablementsType string

const (
	AdminOrgOffboardedListResponseProductEnablementsTypeNgcAdminEval       AdminOrgOffboardedListResponseProductEnablementsType = "NGC_ADMIN_EVAL"
	AdminOrgOffboardedListResponseProductEnablementsTypeNgcAdminNfr        AdminOrgOffboardedListResponseProductEnablementsType = "NGC_ADMIN_NFR"
	AdminOrgOffboardedListResponseProductEnablementsTypeNgcAdminCommercial AdminOrgOffboardedListResponseProductEnablementsType = "NGC_ADMIN_COMMERCIAL"
	AdminOrgOffboardedListResponseProductEnablementsTypeEmsEval            AdminOrgOffboardedListResponseProductEnablementsType = "EMS_EVAL"
	AdminOrgOffboardedListResponseProductEnablementsTypeEmsNfr             AdminOrgOffboardedListResponseProductEnablementsType = "EMS_NFR"
	AdminOrgOffboardedListResponseProductEnablementsTypeEmsCommercial      AdminOrgOffboardedListResponseProductEnablementsType = "EMS_COMMERCIAL"
	AdminOrgOffboardedListResponseProductEnablementsTypeNgcAdminDeveloper  AdminOrgOffboardedListResponseProductEnablementsType = "NGC_ADMIN_DEVELOPER"
)

func (r AdminOrgOffboardedListResponseProductEnablementsType) IsKnown() bool {
	switch r {
	case AdminOrgOffboardedListResponseProductEnablementsTypeNgcAdminEval, AdminOrgOffboardedListResponseProductEnablementsTypeNgcAdminNfr, AdminOrgOffboardedListResponseProductEnablementsTypeNgcAdminCommercial, AdminOrgOffboardedListResponseProductEnablementsTypeEmsEval, AdminOrgOffboardedListResponseProductEnablementsTypeEmsNfr, AdminOrgOffboardedListResponseProductEnablementsTypeEmsCommercial, AdminOrgOffboardedListResponseProductEnablementsTypeNgcAdminDeveloper:
		return true
	}
	return false
}

// Purchase Order.
type AdminOrgOffboardedListResponseProductEnablementsPoDetail struct {
	// Entitlement identifier.
	EntitlementID string `json:"entitlementId"`
	// PAK (Product Activation Key) identifier.
	PkID string                                                       `json:"pkId"`
	JSON adminOrgOffboardedListResponseProductEnablementsPoDetailJSON `json:"-"`
}

// adminOrgOffboardedListResponseProductEnablementsPoDetailJSON contains the JSON
// metadata for the struct
// [AdminOrgOffboardedListResponseProductEnablementsPoDetail]
type adminOrgOffboardedListResponseProductEnablementsPoDetailJSON struct {
	EntitlementID apijson.Field
	PkID          apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *AdminOrgOffboardedListResponseProductEnablementsPoDetail) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r adminOrgOffboardedListResponseProductEnablementsPoDetailJSON) RawJSON() string {
	return r.raw
}

// Product Subscription
type AdminOrgOffboardedListResponseProductSubscription struct {
	// Product Name (NVAIE, BASE_COMMAND, FleetCommand, REGISTRY, etc).
	ProductName string `json:"productName,required"`
	// Unique entitlement identifier
	ID string `json:"id"`
	// EMS Subscription type. (options: EMS_EVAL, EMS_NFR and EMS_COMMERCIAL)
	EmsEntitlementType AdminOrgOffboardedListResponseProductSubscriptionsEmsEntitlementType `json:"emsEntitlementType"`
	// Date on which the subscription expires. The subscription is invalid after this
	// date. (yyyy-MM-dd)
	ExpirationDate string `json:"expirationDate"`
	// Date on which the subscription becomes active. (yyyy-MM-dd)
	StartDate string `json:"startDate"`
	// Subscription type. (options: NGC_ADMIN_EVAL, NGC_ADMIN_NFR,
	// NGC_ADMIN_COMMERCIAL)
	Type AdminOrgOffboardedListResponseProductSubscriptionsType `json:"type"`
	JSON adminOrgOffboardedListResponseProductSubscriptionJSON  `json:"-"`
}

// adminOrgOffboardedListResponseProductSubscriptionJSON contains the JSON metadata
// for the struct [AdminOrgOffboardedListResponseProductSubscription]
type adminOrgOffboardedListResponseProductSubscriptionJSON struct {
	ProductName        apijson.Field
	ID                 apijson.Field
	EmsEntitlementType apijson.Field
	ExpirationDate     apijson.Field
	StartDate          apijson.Field
	Type               apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *AdminOrgOffboardedListResponseProductSubscription) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r adminOrgOffboardedListResponseProductSubscriptionJSON) RawJSON() string {
	return r.raw
}

// EMS Subscription type. (options: EMS_EVAL, EMS_NFR and EMS_COMMERCIAL)
type AdminOrgOffboardedListResponseProductSubscriptionsEmsEntitlementType string

const (
	AdminOrgOffboardedListResponseProductSubscriptionsEmsEntitlementTypeEmsEval       AdminOrgOffboardedListResponseProductSubscriptionsEmsEntitlementType = "EMS_EVAL"
	AdminOrgOffboardedListResponseProductSubscriptionsEmsEntitlementTypeEmsNfr        AdminOrgOffboardedListResponseProductSubscriptionsEmsEntitlementType = "EMS_NFR"
	AdminOrgOffboardedListResponseProductSubscriptionsEmsEntitlementTypeEmsCommerical AdminOrgOffboardedListResponseProductSubscriptionsEmsEntitlementType = "EMS_COMMERICAL"
	AdminOrgOffboardedListResponseProductSubscriptionsEmsEntitlementTypeEmsCommercial AdminOrgOffboardedListResponseProductSubscriptionsEmsEntitlementType = "EMS_COMMERCIAL"
)

func (r AdminOrgOffboardedListResponseProductSubscriptionsEmsEntitlementType) IsKnown() bool {
	switch r {
	case AdminOrgOffboardedListResponseProductSubscriptionsEmsEntitlementTypeEmsEval, AdminOrgOffboardedListResponseProductSubscriptionsEmsEntitlementTypeEmsNfr, AdminOrgOffboardedListResponseProductSubscriptionsEmsEntitlementTypeEmsCommerical, AdminOrgOffboardedListResponseProductSubscriptionsEmsEntitlementTypeEmsCommercial:
		return true
	}
	return false
}

// Subscription type. (options: NGC_ADMIN_EVAL, NGC_ADMIN_NFR,
// NGC_ADMIN_COMMERCIAL)
type AdminOrgOffboardedListResponseProductSubscriptionsType string

const (
	AdminOrgOffboardedListResponseProductSubscriptionsTypeNgcAdminEval       AdminOrgOffboardedListResponseProductSubscriptionsType = "NGC_ADMIN_EVAL"
	AdminOrgOffboardedListResponseProductSubscriptionsTypeNgcAdminNfr        AdminOrgOffboardedListResponseProductSubscriptionsType = "NGC_ADMIN_NFR"
	AdminOrgOffboardedListResponseProductSubscriptionsTypeNgcAdminCommercial AdminOrgOffboardedListResponseProductSubscriptionsType = "NGC_ADMIN_COMMERCIAL"
)

func (r AdminOrgOffboardedListResponseProductSubscriptionsType) IsKnown() bool {
	switch r {
	case AdminOrgOffboardedListResponseProductSubscriptionsTypeNgcAdminEval, AdminOrgOffboardedListResponseProductSubscriptionsTypeNgcAdminNfr, AdminOrgOffboardedListResponseProductSubscriptionsTypeNgcAdminCommercial:
		return true
	}
	return false
}

// Repo scan setting definition
type AdminOrgOffboardedListResponseRepoScanSettings struct {
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
	RepoScanShowResults bool                                               `json:"repoScanShowResults"`
	JSON                adminOrgOffboardedListResponseRepoScanSettingsJSON `json:"-"`
}

// adminOrgOffboardedListResponseRepoScanSettingsJSON contains the JSON metadata
// for the struct [AdminOrgOffboardedListResponseRepoScanSettings]
type adminOrgOffboardedListResponseRepoScanSettingsJSON struct {
	RepoScanAllowOverride       apijson.Field
	RepoScanByDefault           apijson.Field
	RepoScanEnabled             apijson.Field
	RepoScanEnableNotifications apijson.Field
	RepoScanEnableTeamOverride  apijson.Field
	RepoScanShowResults         apijson.Field
	raw                         string
	ExtraFields                 map[string]apijson.Field
}

func (r *AdminOrgOffboardedListResponseRepoScanSettings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r adminOrgOffboardedListResponseRepoScanSettingsJSON) RawJSON() string {
	return r.raw
}

type AdminOrgOffboardedListResponseType string

const (
	AdminOrgOffboardedListResponseTypeUnknown    AdminOrgOffboardedListResponseType = "UNKNOWN"
	AdminOrgOffboardedListResponseTypeCloud      AdminOrgOffboardedListResponseType = "CLOUD"
	AdminOrgOffboardedListResponseTypeEnterprise AdminOrgOffboardedListResponseType = "ENTERPRISE"
	AdminOrgOffboardedListResponseTypeIndividual AdminOrgOffboardedListResponseType = "INDIVIDUAL"
)

func (r AdminOrgOffboardedListResponseType) IsKnown() bool {
	switch r {
	case AdminOrgOffboardedListResponseTypeUnknown, AdminOrgOffboardedListResponseTypeCloud, AdminOrgOffboardedListResponseTypeEnterprise, AdminOrgOffboardedListResponseTypeIndividual:
		return true
	}
	return false
}

// Users information.
type AdminOrgOffboardedListResponseUsersInfo struct {
	// Total number of users.
	TotalUsers int64                                       `json:"totalUsers"`
	JSON       adminOrgOffboardedListResponseUsersInfoJSON `json:"-"`
}

// adminOrgOffboardedListResponseUsersInfoJSON contains the JSON metadata for the
// struct [AdminOrgOffboardedListResponseUsersInfo]
type adminOrgOffboardedListResponseUsersInfoJSON struct {
	TotalUsers  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *AdminOrgOffboardedListResponseUsersInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r adminOrgOffboardedListResponseUsersInfoJSON) RawJSON() string {
	return r.raw
}

type AdminOrgOffboardedListParams struct {
	// The page number of result
	PageNumber param.Field[int64] `query:"page-number"`
	// The page size of result
	PageSize param.Field[int64] `query:"page-size"`
}

// URLQuery serializes [AdminOrgOffboardedListParams]'s query parameters as
// `url.Values`.
func (r AdminOrgOffboardedListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
