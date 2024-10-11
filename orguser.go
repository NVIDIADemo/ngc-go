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
	"github.com/NVIDIADemo/ngc-go/internal/pagination"
	"github.com/NVIDIADemo/ngc-go/internal/param"
	"github.com/NVIDIADemo/ngc-go/internal/requestconfig"
	"github.com/NVIDIADemo/ngc-go/option"
	"github.com/NVIDIADemo/ngc-go/shared"
)

// OrgUserService contains methods and other services that help with interacting
// with the ngc API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewOrgUserService] method instead.
type OrgUserService struct {
	Options        []option.RequestOption
	NcaInvitations *OrgUserNcaInvitationService
}

// NewOrgUserService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewOrgUserService(opts ...option.RequestOption) (r *OrgUserService) {
	r = &OrgUserService{}
	r.Options = opts
	r.NcaInvitations = NewOrgUserNcaInvitationService(opts...)
	return
}

// Creates a User
func (r *OrgUserService) New(ctx context.Context, orgName string, params OrgUserNewParams, opts ...option.RequestOption) (res *shared.User, err error) {
	opts = append(r.Options[:], opts...)
	if orgName == "" {
		err = errors.New("missing required org-name parameter")
		return
	}
	path := fmt.Sprintf("v2/org/%s/users", orgName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Get list of users in organization. (User Admin in org privileges required)
func (r *OrgUserService) List(ctx context.Context, orgName string, query OrgUserListParams, opts ...option.RequestOption) (res *pagination.PageNumberUsers[shared.OrgUserListResponse], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if orgName == "" {
		err = errors.New("missing required org-name parameter")
		return
	}
	path := fmt.Sprintf("v2/org/%s/users", orgName)
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

// Get list of users in organization. (User Admin in org privileges required)
func (r *OrgUserService) ListAutoPaging(ctx context.Context, orgName string, query OrgUserListParams, opts ...option.RequestOption) *pagination.PageNumberUsersAutoPager[shared.OrgUserListResponse] {
	return pagination.NewPageNumberUsersAutoPager(r.List(ctx, orgName, query, opts...))
}

// Remove User from org.
func (r *OrgUserService) Delete(ctx context.Context, orgName string, id string, body OrgUserDeleteParams, opts ...option.RequestOption) (res *OrgUserDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	if orgName == "" {
		err = errors.New("missing required org-name parameter")
		return
	}
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("v3/orgs/%s/users/%s", orgName, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, body, &res, opts...)
	return
}

// Invite if user does not exist, otherwise add role in org
func (r *OrgUserService) AddRole(ctx context.Context, orgName string, userEmailOrID string, params OrgUserAddRoleParams, opts ...option.RequestOption) (res *shared.User, err error) {
	opts = append(r.Options[:], opts...)
	if orgName == "" {
		err = errors.New("missing required org-name parameter")
		return
	}
	if userEmailOrID == "" {
		err = errors.New("missing required user-email-or-id parameter")
		return
	}
	path := fmt.Sprintf("v3/orgs/%s/users/%s/add-role", orgName, userEmailOrID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &res, opts...)
	return
}

// Remove role in org if user exists, otherwise remove invitation
func (r *OrgUserService) RemoveRole(ctx context.Context, orgName string, userEmailOrID string, body OrgUserRemoveRoleParams, opts ...option.RequestOption) (res *shared.User, err error) {
	opts = append(r.Options[:], opts...)
	if orgName == "" {
		err = errors.New("missing required org-name parameter")
		return
	}
	if userEmailOrID == "" {
		err = errors.New("missing required user-email-or-id parameter")
		return
	}
	path := fmt.Sprintf("v3/orgs/%s/users/%s/remove-role", orgName, userEmailOrID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, body, &res, opts...)
	return
}

// Update User Role in org
func (r *OrgUserService) UpdateRole(ctx context.Context, orgName string, id string, body OrgUserUpdateRoleParams, opts ...option.RequestOption) (res *shared.User, err error) {
	opts = append(r.Options[:], opts...)
	if orgName == "" {
		err = errors.New("missing required org-name parameter")
		return
	}
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("v2/org/%s/users/%s/update-role", orgName, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// information about the user
type OrgUserListResponse struct {
	// unique Id of this user.
	ID int64 `json:"id"`
	// unique auth client id of this user.
	ClientID string `json:"clientId"`
	// Created date for this user
	CreatedDate string `json:"createdDate"`
	// Email address of the user. This should be unique.
	Email string `json:"email"`
	// Last time the user logged in
	FirstLoginDate string `json:"firstLoginDate"`
	// Determines if the user has beta access
	HasBetaAccess bool `json:"hasBetaAccess"`
	// indicate if user profile has been completed.
	HasProfile bool `json:"hasProfile"`
	// indicates if user has accepted AI Foundry Partnerships eula
	HasSignedAIFoundryPartnershipsEula bool `json:"hasSignedAiFoundryPartnershipsEULA"`
	// indicates if user has accepted Base Command End User License Agreement.
	HasSignedBaseCommandEula bool `json:"hasSignedBaseCommandEULA"`
	// indicates if user has accepted Base Command Manager End User License Agreement.
	HasSignedBaseCommandManagerEula bool `json:"hasSignedBaseCommandManagerEULA"`
	// indicates if user has accepted BioNeMo End User License Agreement.
	HasSignedBioNeMoEula bool `json:"hasSignedBioNeMoEULA"`
	// indicates if user has accepted container publishing eula
	HasSignedContainerPublishingEula bool `json:"hasSignedContainerPublishingEULA"`
	// indicates if user has accepted CuOpt eula
	HasSignedCuOptEula bool `json:"hasSignedCuOptEULA"`
	// indicates if user has accepted Earth-2 eula
	HasSignedEarth2Eula bool `json:"hasSignedEarth2EULA"`
	// [Deprecated] indicates if user has accepted EGX End User License Agreement.
	HasSignedEgxEula bool `json:"hasSignedEgxEULA"`
	// Determines if the user has signed the NGC End User License Agreement.
	HasSignedEula bool `json:"hasSignedEULA"`
	// indicates if user has accepted Fleet Command End User License Agreement.
	HasSignedFleetCommandEula bool `json:"hasSignedFleetCommandEULA"`
	// indicates if user has accepted LLM End User License Agreement.
	HasSignedLlmEula bool `json:"hasSignedLlmEULA"`
	// indicates if user has accepted Fleet Command End User License Agreement.
	HasSignedNvaieeula bool `json:"hasSignedNVAIEEULA"`
	// Determines if the user has signed the NVIDIA End User License Agreement.
	HasSignedNvidiaEula bool `json:"hasSignedNvidiaEULA"`
	// indicates if user has accepted Nvidia Quantum Cloud End User License Agreement.
	HasSignedNvqceula bool `json:"hasSignedNVQCEULA"`
	// indicates if user has accepted Omniverse End User License Agreement.
	HasSignedOmniverseEula bool `json:"hasSignedOmniverseEULA"`
	// Determines if the user has signed the Privacy Policy.
	HasSignedPrivacyPolicy bool `json:"hasSignedPrivacyPolicy"`
	// indicates if user has consented to share their registration info with other
	// parties
	HasSignedThirdPartyRegistryShareEula bool `json:"hasSignedThirdPartyRegistryShareEULA"`
	// Determines if the user has opted in email subscription.
	HasSubscribedToEmail bool `json:"hasSubscribedToEmail"`
	// Type of IDP, Identity Provider. Used for login.
	IdpType OrgUserListResponseIdpType `json:"idpType"`
	// Determines if the user is active or not.
	IsActive bool `json:"isActive"`
	// Indicates if user was deleted from the system.
	IsDeleted bool `json:"isDeleted"`
	// Determines if the user is a SAML account or not.
	IsSAML bool `json:"isSAML"`
	// Title of user's job position.
	JobPositionTitle string `json:"jobPositionTitle"`
	// Last time the user logged in
	LastLoginDate string `json:"lastLoginDate"`
	// user name
	Name string `json:"name"`
	// List of roles that the user have
	Roles []OrgUserListResponseRole `json:"roles"`
	// unique starfleet id of this user.
	StarfleetID string `json:"starfleetId"`
	// Storage quota for this user.
	StorageQuota []OrgUserListResponseStorageQuota `json:"storageQuota"`
	// Updated date for this user
	UpdatedDate string `json:"updatedDate"`
	// Metadata information about the user.
	UserMetadata OrgUserListResponseUserMetadata `json:"userMetadata"`
	JSON         orgUserListResponseJSON         `json:"-"`
}

// orgUserListResponseJSON contains the JSON metadata for the struct
// [OrgUserListResponse]
type orgUserListResponseJSON struct {
	ID                                   apijson.Field
	ClientID                             apijson.Field
	CreatedDate                          apijson.Field
	Email                                apijson.Field
	FirstLoginDate                       apijson.Field
	HasBetaAccess                        apijson.Field
	HasProfile                           apijson.Field
	HasSignedAIFoundryPartnershipsEula   apijson.Field
	HasSignedBaseCommandEula             apijson.Field
	HasSignedBaseCommandManagerEula      apijson.Field
	HasSignedBioNeMoEula                 apijson.Field
	HasSignedContainerPublishingEula     apijson.Field
	HasSignedCuOptEula                   apijson.Field
	HasSignedEarth2Eula                  apijson.Field
	HasSignedEgxEula                     apijson.Field
	HasSignedEula                        apijson.Field
	HasSignedFleetCommandEula            apijson.Field
	HasSignedLlmEula                     apijson.Field
	HasSignedNvaieeula                   apijson.Field
	HasSignedNvidiaEula                  apijson.Field
	HasSignedNvqceula                    apijson.Field
	HasSignedOmniverseEula               apijson.Field
	HasSignedPrivacyPolicy               apijson.Field
	HasSignedThirdPartyRegistryShareEula apijson.Field
	HasSubscribedToEmail                 apijson.Field
	IdpType                              apijson.Field
	IsActive                             apijson.Field
	IsDeleted                            apijson.Field
	IsSAML                               apijson.Field
	JobPositionTitle                     apijson.Field
	LastLoginDate                        apijson.Field
	Name                                 apijson.Field
	Roles                                apijson.Field
	StarfleetID                          apijson.Field
	StorageQuota                         apijson.Field
	UpdatedDate                          apijson.Field
	UserMetadata                         apijson.Field
	raw                                  string
	ExtraFields                          map[string]apijson.Field
}

func (r *OrgUserListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r orgUserListResponseJSON) RawJSON() string {
	return r.raw
}

// Type of IDP, Identity Provider. Used for login.
type OrgUserListResponseIdpType string

const (
	OrgUserListResponseIdpTypeNvidia     OrgUserListResponseIdpType = "NVIDIA"
	OrgUserListResponseIdpTypeEnterprise OrgUserListResponseIdpType = "ENTERPRISE"
)

func (r OrgUserListResponseIdpType) IsKnown() bool {
	switch r {
	case OrgUserListResponseIdpTypeNvidia, OrgUserListResponseIdpTypeEnterprise:
		return true
	}
	return false
}

// List of roles that the user have
type OrgUserListResponseRole struct {
	// Information about the Organization
	Org OrgUserListResponseRolesOrg `json:"org"`
	// List of org role types that the user have
	OrgRoles []string `json:"orgRoles"`
	// DEPRECATED - List of role types that the user have
	RoleTypes []string `json:"roleTypes"`
	// Information about the user who is attempting to run the job
	TargetSystemUserIdentifier OrgUserListResponseRolesTargetSystemUserIdentifier `json:"targetSystemUserIdentifier"`
	// Information about the team
	Team OrgUserListResponseRolesTeam `json:"team"`
	// List of team role types that the user have
	TeamRoles []string                    `json:"teamRoles"`
	JSON      orgUserListResponseRoleJSON `json:"-"`
}

// orgUserListResponseRoleJSON contains the JSON metadata for the struct
// [OrgUserListResponseRole]
type orgUserListResponseRoleJSON struct {
	Org                        apijson.Field
	OrgRoles                   apijson.Field
	RoleTypes                  apijson.Field
	TargetSystemUserIdentifier apijson.Field
	Team                       apijson.Field
	TeamRoles                  apijson.Field
	raw                        string
	ExtraFields                map[string]apijson.Field
}

func (r *OrgUserListResponseRole) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r orgUserListResponseRoleJSON) RawJSON() string {
	return r.raw
}

// Information about the Organization
type OrgUserListResponseRolesOrg struct {
	// Unique Id of this team.
	ID int64 `json:"id"`
	// Org Owner Alternate Contact
	AlternateContact OrgUserListResponseRolesOrgAlternateContact `json:"alternateContact"`
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
	InfinityManagerSettings OrgUserListResponseRolesOrgInfinityManagerSettings `json:"infinityManagerSettings"`
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
	OrgOwner OrgUserListResponseRolesOrgOrgOwner `json:"orgOwner"`
	// Org owners
	OrgOwners []OrgUserListResponseRolesOrgOrgOwner `json:"orgOwners"`
	// Product end customer salesforce.com Id (external customer Id). pecSfdcId is for
	// EMS (entitlement management service) to track external paid customer.
	PecSfdcID            string                                           `json:"pecSfdcId"`
	ProductEnablements   []OrgUserListResponseRolesOrgProductEnablement   `json:"productEnablements"`
	ProductSubscriptions []OrgUserListResponseRolesOrgProductSubscription `json:"productSubscriptions"`
	// Repo scan setting definition
	RepoScanSettings OrgUserListResponseRolesOrgRepoScanSettings `json:"repoScanSettings"`
	Type             OrgUserListResponseRolesOrgType             `json:"type"`
	// Users information.
	UsersInfo OrgUserListResponseRolesOrgUsersInfo `json:"usersInfo"`
	JSON      orgUserListResponseRolesOrgJSON      `json:"-"`
}

// orgUserListResponseRolesOrgJSON contains the JSON metadata for the struct
// [OrgUserListResponseRolesOrg]
type orgUserListResponseRolesOrgJSON struct {
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

func (r *OrgUserListResponseRolesOrg) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r orgUserListResponseRolesOrgJSON) RawJSON() string {
	return r.raw
}

// Org Owner Alternate Contact
type OrgUserListResponseRolesOrgAlternateContact struct {
	// Alternate contact's email.
	Email string `json:"email"`
	// Full name of the alternate contact.
	FullName string                                          `json:"fullName"`
	JSON     orgUserListResponseRolesOrgAlternateContactJSON `json:"-"`
}

// orgUserListResponseRolesOrgAlternateContactJSON contains the JSON metadata for
// the struct [OrgUserListResponseRolesOrgAlternateContact]
type orgUserListResponseRolesOrgAlternateContactJSON struct {
	Email       apijson.Field
	FullName    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OrgUserListResponseRolesOrgAlternateContact) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r orgUserListResponseRolesOrgAlternateContactJSON) RawJSON() string {
	return r.raw
}

// Infinity manager setting definition
type OrgUserListResponseRolesOrgInfinityManagerSettings struct {
	// Enable the infinity manager or not. Used both in org and team level object
	InfinityManagerEnabled bool `json:"infinityManagerEnabled"`
	// Allow override settings at team level. Only used in org level object
	InfinityManagerEnableTeamOverride bool                                                   `json:"infinityManagerEnableTeamOverride"`
	JSON                              orgUserListResponseRolesOrgInfinityManagerSettingsJSON `json:"-"`
}

// orgUserListResponseRolesOrgInfinityManagerSettingsJSON contains the JSON
// metadata for the struct [OrgUserListResponseRolesOrgInfinityManagerSettings]
type orgUserListResponseRolesOrgInfinityManagerSettingsJSON struct {
	InfinityManagerEnabled            apijson.Field
	InfinityManagerEnableTeamOverride apijson.Field
	raw                               string
	ExtraFields                       map[string]apijson.Field
}

func (r *OrgUserListResponseRolesOrgInfinityManagerSettings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r orgUserListResponseRolesOrgInfinityManagerSettingsJSON) RawJSON() string {
	return r.raw
}

// Org owner.
type OrgUserListResponseRolesOrgOrgOwner struct {
	// Email address of the org owner.
	Email string `json:"email,required"`
	// Org owner name.
	FullName string `json:"fullName,required"`
	// Last time the org owner logged in.
	LastLoginDate string                                  `json:"lastLoginDate"`
	JSON          orgUserListResponseRolesOrgOrgOwnerJSON `json:"-"`
}

// orgUserListResponseRolesOrgOrgOwnerJSON contains the JSON metadata for the
// struct [OrgUserListResponseRolesOrgOrgOwner]
type orgUserListResponseRolesOrgOrgOwnerJSON struct {
	Email         apijson.Field
	FullName      apijson.Field
	LastLoginDate apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *OrgUserListResponseRolesOrgOrgOwner) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r orgUserListResponseRolesOrgOrgOwnerJSON) RawJSON() string {
	return r.raw
}

// Product Enablement
type OrgUserListResponseRolesOrgProductEnablement struct {
	// Product Name (NVAIE, BASE_COMMAND, REGISTRY, etc)
	ProductName string `json:"productName,required"`
	// Product Enablement Types
	Type OrgUserListResponseRolesOrgProductEnablementsType `json:"type,required"`
	// Date on which the subscription expires. The subscription is invalid after this
	// date. (yyyy-MM-dd)
	ExpirationDate string                                                  `json:"expirationDate"`
	PoDetails      []OrgUserListResponseRolesOrgProductEnablementsPoDetail `json:"poDetails"`
	JSON           orgUserListResponseRolesOrgProductEnablementJSON        `json:"-"`
}

// orgUserListResponseRolesOrgProductEnablementJSON contains the JSON metadata for
// the struct [OrgUserListResponseRolesOrgProductEnablement]
type orgUserListResponseRolesOrgProductEnablementJSON struct {
	ProductName    apijson.Field
	Type           apijson.Field
	ExpirationDate apijson.Field
	PoDetails      apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *OrgUserListResponseRolesOrgProductEnablement) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r orgUserListResponseRolesOrgProductEnablementJSON) RawJSON() string {
	return r.raw
}

// Product Enablement Types
type OrgUserListResponseRolesOrgProductEnablementsType string

const (
	OrgUserListResponseRolesOrgProductEnablementsTypeNgcAdminEval       OrgUserListResponseRolesOrgProductEnablementsType = "NGC_ADMIN_EVAL"
	OrgUserListResponseRolesOrgProductEnablementsTypeNgcAdminNfr        OrgUserListResponseRolesOrgProductEnablementsType = "NGC_ADMIN_NFR"
	OrgUserListResponseRolesOrgProductEnablementsTypeNgcAdminCommercial OrgUserListResponseRolesOrgProductEnablementsType = "NGC_ADMIN_COMMERCIAL"
	OrgUserListResponseRolesOrgProductEnablementsTypeEmsEval            OrgUserListResponseRolesOrgProductEnablementsType = "EMS_EVAL"
	OrgUserListResponseRolesOrgProductEnablementsTypeEmsNfr             OrgUserListResponseRolesOrgProductEnablementsType = "EMS_NFR"
	OrgUserListResponseRolesOrgProductEnablementsTypeEmsCommercial      OrgUserListResponseRolesOrgProductEnablementsType = "EMS_COMMERCIAL"
	OrgUserListResponseRolesOrgProductEnablementsTypeNgcAdminDeveloper  OrgUserListResponseRolesOrgProductEnablementsType = "NGC_ADMIN_DEVELOPER"
)

func (r OrgUserListResponseRolesOrgProductEnablementsType) IsKnown() bool {
	switch r {
	case OrgUserListResponseRolesOrgProductEnablementsTypeNgcAdminEval, OrgUserListResponseRolesOrgProductEnablementsTypeNgcAdminNfr, OrgUserListResponseRolesOrgProductEnablementsTypeNgcAdminCommercial, OrgUserListResponseRolesOrgProductEnablementsTypeEmsEval, OrgUserListResponseRolesOrgProductEnablementsTypeEmsNfr, OrgUserListResponseRolesOrgProductEnablementsTypeEmsCommercial, OrgUserListResponseRolesOrgProductEnablementsTypeNgcAdminDeveloper:
		return true
	}
	return false
}

// Purchase Order.
type OrgUserListResponseRolesOrgProductEnablementsPoDetail struct {
	// Entitlement identifier.
	EntitlementID string `json:"entitlementId"`
	// PAK (Product Activation Key) identifier.
	PkID string                                                    `json:"pkId"`
	JSON orgUserListResponseRolesOrgProductEnablementsPoDetailJSON `json:"-"`
}

// orgUserListResponseRolesOrgProductEnablementsPoDetailJSON contains the JSON
// metadata for the struct [OrgUserListResponseRolesOrgProductEnablementsPoDetail]
type orgUserListResponseRolesOrgProductEnablementsPoDetailJSON struct {
	EntitlementID apijson.Field
	PkID          apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *OrgUserListResponseRolesOrgProductEnablementsPoDetail) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r orgUserListResponseRolesOrgProductEnablementsPoDetailJSON) RawJSON() string {
	return r.raw
}

// Product Subscription
type OrgUserListResponseRolesOrgProductSubscription struct {
	// Product Name (NVAIE, BASE_COMMAND, FleetCommand, REGISTRY, etc).
	ProductName string `json:"productName,required"`
	// Unique entitlement identifier
	ID string `json:"id"`
	// EMS Subscription type. (options: EMS_EVAL, EMS_NFR and EMS_COMMERCIAL)
	EmsEntitlementType OrgUserListResponseRolesOrgProductSubscriptionsEmsEntitlementType `json:"emsEntitlementType"`
	// Date on which the subscription expires. The subscription is invalid after this
	// date. (yyyy-MM-dd)
	ExpirationDate string `json:"expirationDate"`
	// Date on which the subscription becomes active. (yyyy-MM-dd)
	StartDate string `json:"startDate"`
	// Subscription type. (options: NGC_ADMIN_EVAL, NGC_ADMIN_NFR,
	// NGC_ADMIN_COMMERCIAL)
	Type OrgUserListResponseRolesOrgProductSubscriptionsType `json:"type"`
	JSON orgUserListResponseRolesOrgProductSubscriptionJSON  `json:"-"`
}

// orgUserListResponseRolesOrgProductSubscriptionJSON contains the JSON metadata
// for the struct [OrgUserListResponseRolesOrgProductSubscription]
type orgUserListResponseRolesOrgProductSubscriptionJSON struct {
	ProductName        apijson.Field
	ID                 apijson.Field
	EmsEntitlementType apijson.Field
	ExpirationDate     apijson.Field
	StartDate          apijson.Field
	Type               apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *OrgUserListResponseRolesOrgProductSubscription) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r orgUserListResponseRolesOrgProductSubscriptionJSON) RawJSON() string {
	return r.raw
}

// EMS Subscription type. (options: EMS_EVAL, EMS_NFR and EMS_COMMERCIAL)
type OrgUserListResponseRolesOrgProductSubscriptionsEmsEntitlementType string

const (
	OrgUserListResponseRolesOrgProductSubscriptionsEmsEntitlementTypeEmsEval       OrgUserListResponseRolesOrgProductSubscriptionsEmsEntitlementType = "EMS_EVAL"
	OrgUserListResponseRolesOrgProductSubscriptionsEmsEntitlementTypeEmsNfr        OrgUserListResponseRolesOrgProductSubscriptionsEmsEntitlementType = "EMS_NFR"
	OrgUserListResponseRolesOrgProductSubscriptionsEmsEntitlementTypeEmsCommerical OrgUserListResponseRolesOrgProductSubscriptionsEmsEntitlementType = "EMS_COMMERICAL"
	OrgUserListResponseRolesOrgProductSubscriptionsEmsEntitlementTypeEmsCommercial OrgUserListResponseRolesOrgProductSubscriptionsEmsEntitlementType = "EMS_COMMERCIAL"
)

func (r OrgUserListResponseRolesOrgProductSubscriptionsEmsEntitlementType) IsKnown() bool {
	switch r {
	case OrgUserListResponseRolesOrgProductSubscriptionsEmsEntitlementTypeEmsEval, OrgUserListResponseRolesOrgProductSubscriptionsEmsEntitlementTypeEmsNfr, OrgUserListResponseRolesOrgProductSubscriptionsEmsEntitlementTypeEmsCommerical, OrgUserListResponseRolesOrgProductSubscriptionsEmsEntitlementTypeEmsCommercial:
		return true
	}
	return false
}

// Subscription type. (options: NGC_ADMIN_EVAL, NGC_ADMIN_NFR,
// NGC_ADMIN_COMMERCIAL)
type OrgUserListResponseRolesOrgProductSubscriptionsType string

const (
	OrgUserListResponseRolesOrgProductSubscriptionsTypeNgcAdminEval       OrgUserListResponseRolesOrgProductSubscriptionsType = "NGC_ADMIN_EVAL"
	OrgUserListResponseRolesOrgProductSubscriptionsTypeNgcAdminNfr        OrgUserListResponseRolesOrgProductSubscriptionsType = "NGC_ADMIN_NFR"
	OrgUserListResponseRolesOrgProductSubscriptionsTypeNgcAdminCommercial OrgUserListResponseRolesOrgProductSubscriptionsType = "NGC_ADMIN_COMMERCIAL"
)

func (r OrgUserListResponseRolesOrgProductSubscriptionsType) IsKnown() bool {
	switch r {
	case OrgUserListResponseRolesOrgProductSubscriptionsTypeNgcAdminEval, OrgUserListResponseRolesOrgProductSubscriptionsTypeNgcAdminNfr, OrgUserListResponseRolesOrgProductSubscriptionsTypeNgcAdminCommercial:
		return true
	}
	return false
}

// Repo scan setting definition
type OrgUserListResponseRolesOrgRepoScanSettings struct {
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
	RepoScanShowResults bool                                            `json:"repoScanShowResults"`
	JSON                orgUserListResponseRolesOrgRepoScanSettingsJSON `json:"-"`
}

// orgUserListResponseRolesOrgRepoScanSettingsJSON contains the JSON metadata for
// the struct [OrgUserListResponseRolesOrgRepoScanSettings]
type orgUserListResponseRolesOrgRepoScanSettingsJSON struct {
	RepoScanAllowOverride       apijson.Field
	RepoScanByDefault           apijson.Field
	RepoScanEnabled             apijson.Field
	RepoScanEnableNotifications apijson.Field
	RepoScanEnableTeamOverride  apijson.Field
	RepoScanShowResults         apijson.Field
	raw                         string
	ExtraFields                 map[string]apijson.Field
}

func (r *OrgUserListResponseRolesOrgRepoScanSettings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r orgUserListResponseRolesOrgRepoScanSettingsJSON) RawJSON() string {
	return r.raw
}

type OrgUserListResponseRolesOrgType string

const (
	OrgUserListResponseRolesOrgTypeUnknown    OrgUserListResponseRolesOrgType = "UNKNOWN"
	OrgUserListResponseRolesOrgTypeCloud      OrgUserListResponseRolesOrgType = "CLOUD"
	OrgUserListResponseRolesOrgTypeEnterprise OrgUserListResponseRolesOrgType = "ENTERPRISE"
	OrgUserListResponseRolesOrgTypeIndividual OrgUserListResponseRolesOrgType = "INDIVIDUAL"
)

func (r OrgUserListResponseRolesOrgType) IsKnown() bool {
	switch r {
	case OrgUserListResponseRolesOrgTypeUnknown, OrgUserListResponseRolesOrgTypeCloud, OrgUserListResponseRolesOrgTypeEnterprise, OrgUserListResponseRolesOrgTypeIndividual:
		return true
	}
	return false
}

// Users information.
type OrgUserListResponseRolesOrgUsersInfo struct {
	// Total number of users.
	TotalUsers int64                                    `json:"totalUsers"`
	JSON       orgUserListResponseRolesOrgUsersInfoJSON `json:"-"`
}

// orgUserListResponseRolesOrgUsersInfoJSON contains the JSON metadata for the
// struct [OrgUserListResponseRolesOrgUsersInfo]
type orgUserListResponseRolesOrgUsersInfoJSON struct {
	TotalUsers  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OrgUserListResponseRolesOrgUsersInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r orgUserListResponseRolesOrgUsersInfoJSON) RawJSON() string {
	return r.raw
}

// Information about the user who is attempting to run the job
type OrgUserListResponseRolesTargetSystemUserIdentifier struct {
	// gid of the user on this team
	Gid int64 `json:"gid"`
	// Org context for the job
	OrgName string `json:"orgName"`
	// Starfleet ID of the user creating the job.
	StarfleetID string `json:"starfleetId"`
	// Team context for the job
	TeamName string `json:"teamName"`
	// uid of the user on this team
	Uid int64 `json:"uid"`
	// Unique ID of the user who submitted the job
	UserID int64                                                  `json:"userId"`
	JSON   orgUserListResponseRolesTargetSystemUserIdentifierJSON `json:"-"`
}

// orgUserListResponseRolesTargetSystemUserIdentifierJSON contains the JSON
// metadata for the struct [OrgUserListResponseRolesTargetSystemUserIdentifier]
type orgUserListResponseRolesTargetSystemUserIdentifierJSON struct {
	Gid         apijson.Field
	OrgName     apijson.Field
	StarfleetID apijson.Field
	TeamName    apijson.Field
	Uid         apijson.Field
	UserID      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OrgUserListResponseRolesTargetSystemUserIdentifier) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r orgUserListResponseRolesTargetSystemUserIdentifierJSON) RawJSON() string {
	return r.raw
}

// Information about the team
type OrgUserListResponseRolesTeam struct {
	// unique Id of this team.
	ID int64 `json:"id"`
	// description of the team
	Description string `json:"description"`
	// Infinity manager setting definition
	InfinityManagerSettings OrgUserListResponseRolesTeamInfinityManagerSettings `json:"infinityManagerSettings"`
	// indicates if the team is deleted or not
	IsDeleted bool `json:"isDeleted"`
	// team name
	Name string `json:"name"`
	// Repo scan setting definition
	RepoScanSettings OrgUserListResponseRolesTeamRepoScanSettings `json:"repoScanSettings"`
	JSON             orgUserListResponseRolesTeamJSON             `json:"-"`
}

// orgUserListResponseRolesTeamJSON contains the JSON metadata for the struct
// [OrgUserListResponseRolesTeam]
type orgUserListResponseRolesTeamJSON struct {
	ID                      apijson.Field
	Description             apijson.Field
	InfinityManagerSettings apijson.Field
	IsDeleted               apijson.Field
	Name                    apijson.Field
	RepoScanSettings        apijson.Field
	raw                     string
	ExtraFields             map[string]apijson.Field
}

func (r *OrgUserListResponseRolesTeam) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r orgUserListResponseRolesTeamJSON) RawJSON() string {
	return r.raw
}

// Infinity manager setting definition
type OrgUserListResponseRolesTeamInfinityManagerSettings struct {
	// Enable the infinity manager or not. Used both in org and team level object
	InfinityManagerEnabled bool `json:"infinityManagerEnabled"`
	// Allow override settings at team level. Only used in org level object
	InfinityManagerEnableTeamOverride bool                                                    `json:"infinityManagerEnableTeamOverride"`
	JSON                              orgUserListResponseRolesTeamInfinityManagerSettingsJSON `json:"-"`
}

// orgUserListResponseRolesTeamInfinityManagerSettingsJSON contains the JSON
// metadata for the struct [OrgUserListResponseRolesTeamInfinityManagerSettings]
type orgUserListResponseRolesTeamInfinityManagerSettingsJSON struct {
	InfinityManagerEnabled            apijson.Field
	InfinityManagerEnableTeamOverride apijson.Field
	raw                               string
	ExtraFields                       map[string]apijson.Field
}

func (r *OrgUserListResponseRolesTeamInfinityManagerSettings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r orgUserListResponseRolesTeamInfinityManagerSettingsJSON) RawJSON() string {
	return r.raw
}

// Repo scan setting definition
type OrgUserListResponseRolesTeamRepoScanSettings struct {
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
	RepoScanShowResults bool                                             `json:"repoScanShowResults"`
	JSON                orgUserListResponseRolesTeamRepoScanSettingsJSON `json:"-"`
}

// orgUserListResponseRolesTeamRepoScanSettingsJSON contains the JSON metadata for
// the struct [OrgUserListResponseRolesTeamRepoScanSettings]
type orgUserListResponseRolesTeamRepoScanSettingsJSON struct {
	RepoScanAllowOverride       apijson.Field
	RepoScanByDefault           apijson.Field
	RepoScanEnabled             apijson.Field
	RepoScanEnableNotifications apijson.Field
	RepoScanEnableTeamOverride  apijson.Field
	RepoScanShowResults         apijson.Field
	raw                         string
	ExtraFields                 map[string]apijson.Field
}

func (r *OrgUserListResponseRolesTeamRepoScanSettings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r orgUserListResponseRolesTeamRepoScanSettingsJSON) RawJSON() string {
	return r.raw
}

// represents user storage quota for a given ace and available unused storage
type OrgUserListResponseStorageQuota struct {
	// id of the ace
	AceID int64 `json:"aceId"`
	// name of the ace
	AceName string `json:"aceName"`
	// Available space in bytes
	Available int64 `json:"available"`
	// Number of datasets that are a part of user's used storage
	DatasetCount int64 `json:"datasetCount"`
	// Space used by datasets in bytes
	DatasetsUsage int64 `json:"datasetsUsage"`
	// The org name that this user quota tied to. This is needed for analytics
	OrgName string `json:"orgName"`
	// Assigned quota in bytes
	Quota int64 `json:"quota"`
	// Number of resultsets that are a part of user's used storage
	ResultsetCount int64 `json:"resultsetCount"`
	// Space used by resultsets in bytes
	ResultsetsUsage int64 `json:"resultsetsUsage"`
	// Description of this storage cluster
	StorageClusterDescription string `json:"storageClusterDescription"`
	// Name of storage cluster
	StorageClusterName string `json:"storageClusterName"`
	// Identifier to this storage cluster
	StorageClusterUuid string `json:"storageClusterUuid"`
	// Number of workspaces that are a part of user's used storage
	WorkspacesCount int64 `json:"workspacesCount"`
	// Space used by workspaces in bytes
	WorkspacesUsage int64                               `json:"workspacesUsage"`
	JSON            orgUserListResponseStorageQuotaJSON `json:"-"`
}

// orgUserListResponseStorageQuotaJSON contains the JSON metadata for the struct
// [OrgUserListResponseStorageQuota]
type orgUserListResponseStorageQuotaJSON struct {
	AceID                     apijson.Field
	AceName                   apijson.Field
	Available                 apijson.Field
	DatasetCount              apijson.Field
	DatasetsUsage             apijson.Field
	OrgName                   apijson.Field
	Quota                     apijson.Field
	ResultsetCount            apijson.Field
	ResultsetsUsage           apijson.Field
	StorageClusterDescription apijson.Field
	StorageClusterName        apijson.Field
	StorageClusterUuid        apijson.Field
	WorkspacesCount           apijson.Field
	WorkspacesUsage           apijson.Field
	raw                       string
	ExtraFields               map[string]apijson.Field
}

func (r *OrgUserListResponseStorageQuota) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r orgUserListResponseStorageQuotaJSON) RawJSON() string {
	return r.raw
}

// Metadata information about the user.
type OrgUserListResponseUserMetadata struct {
	// Name of the company
	Company string `json:"company"`
	// Company URL
	CompanyURL string `json:"companyUrl"`
	// Country of the user
	Country string `json:"country"`
	// User's first name
	FirstName string `json:"firstName"`
	// Industry segment
	Industry string `json:"industry"`
	// List of development areas that user has interest
	Interest []string `json:"interest"`
	// User's last name
	LastName string `json:"lastName"`
	// Role of the user in the company
	Role string                              `json:"role"`
	JSON orgUserListResponseUserMetadataJSON `json:"-"`
}

// orgUserListResponseUserMetadataJSON contains the JSON metadata for the struct
// [OrgUserListResponseUserMetadata]
type orgUserListResponseUserMetadataJSON struct {
	Company     apijson.Field
	CompanyURL  apijson.Field
	Country     apijson.Field
	FirstName   apijson.Field
	Industry    apijson.Field
	Interest    apijson.Field
	LastName    apijson.Field
	Role        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OrgUserListResponseUserMetadata) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r orgUserListResponseUserMetadataJSON) RawJSON() string {
	return r.raw
}

type OrgUserDeleteResponse struct {
	RequestStatus OrgUserDeleteResponseRequestStatus `json:"requestStatus"`
	JSON          orgUserDeleteResponseJSON          `json:"-"`
}

// orgUserDeleteResponseJSON contains the JSON metadata for the struct
// [OrgUserDeleteResponse]
type orgUserDeleteResponseJSON struct {
	RequestStatus apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *OrgUserDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r orgUserDeleteResponseJSON) RawJSON() string {
	return r.raw
}

type OrgUserDeleteResponseRequestStatus struct {
	RequestID string `json:"requestId"`
	ServerID  string `json:"serverId"`
	// Describes response status reported by the server.
	StatusCode        OrgUserDeleteResponseRequestStatusStatusCode `json:"statusCode"`
	StatusDescription string                                       `json:"statusDescription"`
	JSON              orgUserDeleteResponseRequestStatusJSON       `json:"-"`
}

// orgUserDeleteResponseRequestStatusJSON contains the JSON metadata for the struct
// [OrgUserDeleteResponseRequestStatus]
type orgUserDeleteResponseRequestStatusJSON struct {
	RequestID         apijson.Field
	ServerID          apijson.Field
	StatusCode        apijson.Field
	StatusDescription apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *OrgUserDeleteResponseRequestStatus) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r orgUserDeleteResponseRequestStatusJSON) RawJSON() string {
	return r.raw
}

// Describes response status reported by the server.
type OrgUserDeleteResponseRequestStatusStatusCode string

const (
	OrgUserDeleteResponseRequestStatusStatusCodeUnknown                    OrgUserDeleteResponseRequestStatusStatusCode = "UNKNOWN"
	OrgUserDeleteResponseRequestStatusStatusCodeSuccess                    OrgUserDeleteResponseRequestStatusStatusCode = "SUCCESS"
	OrgUserDeleteResponseRequestStatusStatusCodeUnauthorized               OrgUserDeleteResponseRequestStatusStatusCode = "UNAUTHORIZED"
	OrgUserDeleteResponseRequestStatusStatusCodePaymentRequired            OrgUserDeleteResponseRequestStatusStatusCode = "PAYMENT_REQUIRED"
	OrgUserDeleteResponseRequestStatusStatusCodeForbidden                  OrgUserDeleteResponseRequestStatusStatusCode = "FORBIDDEN"
	OrgUserDeleteResponseRequestStatusStatusCodeTimeout                    OrgUserDeleteResponseRequestStatusStatusCode = "TIMEOUT"
	OrgUserDeleteResponseRequestStatusStatusCodeExists                     OrgUserDeleteResponseRequestStatusStatusCode = "EXISTS"
	OrgUserDeleteResponseRequestStatusStatusCodeNotFound                   OrgUserDeleteResponseRequestStatusStatusCode = "NOT_FOUND"
	OrgUserDeleteResponseRequestStatusStatusCodeInternalError              OrgUserDeleteResponseRequestStatusStatusCode = "INTERNAL_ERROR"
	OrgUserDeleteResponseRequestStatusStatusCodeInvalidRequest             OrgUserDeleteResponseRequestStatusStatusCode = "INVALID_REQUEST"
	OrgUserDeleteResponseRequestStatusStatusCodeInvalidRequestVersion      OrgUserDeleteResponseRequestStatusStatusCode = "INVALID_REQUEST_VERSION"
	OrgUserDeleteResponseRequestStatusStatusCodeInvalidRequestData         OrgUserDeleteResponseRequestStatusStatusCode = "INVALID_REQUEST_DATA"
	OrgUserDeleteResponseRequestStatusStatusCodeMethodNotAllowed           OrgUserDeleteResponseRequestStatusStatusCode = "METHOD_NOT_ALLOWED"
	OrgUserDeleteResponseRequestStatusStatusCodeConflict                   OrgUserDeleteResponseRequestStatusStatusCode = "CONFLICT"
	OrgUserDeleteResponseRequestStatusStatusCodeUnprocessableEntity        OrgUserDeleteResponseRequestStatusStatusCode = "UNPROCESSABLE_ENTITY"
	OrgUserDeleteResponseRequestStatusStatusCodeTooManyRequests            OrgUserDeleteResponseRequestStatusStatusCode = "TOO_MANY_REQUESTS"
	OrgUserDeleteResponseRequestStatusStatusCodeInsufficientStorage        OrgUserDeleteResponseRequestStatusStatusCode = "INSUFFICIENT_STORAGE"
	OrgUserDeleteResponseRequestStatusStatusCodeServiceUnavailable         OrgUserDeleteResponseRequestStatusStatusCode = "SERVICE_UNAVAILABLE"
	OrgUserDeleteResponseRequestStatusStatusCodePayloadTooLarge            OrgUserDeleteResponseRequestStatusStatusCode = "PAYLOAD_TOO_LARGE"
	OrgUserDeleteResponseRequestStatusStatusCodeNotAcceptable              OrgUserDeleteResponseRequestStatusStatusCode = "NOT_ACCEPTABLE"
	OrgUserDeleteResponseRequestStatusStatusCodeUnavailableForLegalReasons OrgUserDeleteResponseRequestStatusStatusCode = "UNAVAILABLE_FOR_LEGAL_REASONS"
	OrgUserDeleteResponseRequestStatusStatusCodeBadGateway                 OrgUserDeleteResponseRequestStatusStatusCode = "BAD_GATEWAY"
)

func (r OrgUserDeleteResponseRequestStatusStatusCode) IsKnown() bool {
	switch r {
	case OrgUserDeleteResponseRequestStatusStatusCodeUnknown, OrgUserDeleteResponseRequestStatusStatusCodeSuccess, OrgUserDeleteResponseRequestStatusStatusCodeUnauthorized, OrgUserDeleteResponseRequestStatusStatusCodePaymentRequired, OrgUserDeleteResponseRequestStatusStatusCodeForbidden, OrgUserDeleteResponseRequestStatusStatusCodeTimeout, OrgUserDeleteResponseRequestStatusStatusCodeExists, OrgUserDeleteResponseRequestStatusStatusCodeNotFound, OrgUserDeleteResponseRequestStatusStatusCodeInternalError, OrgUserDeleteResponseRequestStatusStatusCodeInvalidRequest, OrgUserDeleteResponseRequestStatusStatusCodeInvalidRequestVersion, OrgUserDeleteResponseRequestStatusStatusCodeInvalidRequestData, OrgUserDeleteResponseRequestStatusStatusCodeMethodNotAllowed, OrgUserDeleteResponseRequestStatusStatusCodeConflict, OrgUserDeleteResponseRequestStatusStatusCodeUnprocessableEntity, OrgUserDeleteResponseRequestStatusStatusCodeTooManyRequests, OrgUserDeleteResponseRequestStatusStatusCodeInsufficientStorage, OrgUserDeleteResponseRequestStatusStatusCodeServiceUnavailable, OrgUserDeleteResponseRequestStatusStatusCodePayloadTooLarge, OrgUserDeleteResponseRequestStatusStatusCodeNotAcceptable, OrgUserDeleteResponseRequestStatusStatusCodeUnavailableForLegalReasons, OrgUserDeleteResponseRequestStatusStatusCodeBadGateway:
		return true
	}
	return false
}

type OrgUserNewParams struct {
	// Email address of the user. This should be unique.
	Email param.Field[string] `json:"email,required"`
	// If the IDP ID is provided then it is used instead of the one configured for the
	// organization
	IdpID     param.Field[string] `query:"idp-id"`
	SendEmail param.Field[bool]   `query:"send-email"`
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
	UserMetadata param.Field[OrgUserNewParamsUserMetadata] `json:"userMetadata"`
	Ncid         param.Field[string]                       `cookie:"ncid"`
	VisitorID    param.Field[string]                       `cookie:"VisitorID"`
}

func (r OrgUserNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// URLQuery serializes [OrgUserNewParams]'s query parameters as `url.Values`.
func (r OrgUserNewParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Metadata information about the user.
type OrgUserNewParamsUserMetadata struct {
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

func (r OrgUserNewParamsUserMetadata) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

type OrgUserListParams struct {
	// Name of team to exclude members from
	ExcludeFromTeam param.Field[string] `query:"exclude-from-team"`
	// The page number of result
	PageNumber param.Field[int64] `query:"page-number"`
	// The page size of result
	PageSize param.Field[int64] `query:"page-size"`
	// User Search Parameters. Only 'filters' and 'orderBy' for 'name' and 'email' are
	// implemented
	Q param.Field[OrgUserListParamsQ] `query:"q"`
}

// URLQuery serializes [OrgUserListParams]'s query parameters as `url.Values`.
func (r OrgUserListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// User Search Parameters. Only 'filters' and 'orderBy' for 'name' and 'email' are
// implemented
type OrgUserListParamsQ struct {
	Fields      param.Field[[]string]                    `query:"fields"`
	Filters     param.Field[[]OrgUserListParamsQFilter]  `query:"filters"`
	GroupBy     param.Field[string]                      `query:"groupBy"`
	OrderBy     param.Field[[]OrgUserListParamsQOrderBy] `query:"orderBy"`
	Page        param.Field[int64]                       `query:"page"`
	PageSize    param.Field[int64]                       `query:"pageSize"`
	Query       param.Field[string]                      `query:"query"`
	QueryFields param.Field[[]string]                    `query:"queryFields"`
	ScoredSize  param.Field[int64]                       `query:"scoredSize"`
}

// URLQuery serializes [OrgUserListParamsQ]'s query parameters as `url.Values`.
func (r OrgUserListParamsQ) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type OrgUserListParamsQFilter struct {
	Field param.Field[string] `query:"field"`
	Value param.Field[string] `query:"value"`
}

// URLQuery serializes [OrgUserListParamsQFilter]'s query parameters as
// `url.Values`.
func (r OrgUserListParamsQFilter) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type OrgUserListParamsQOrderBy struct {
	Field param.Field[string]                         `query:"field"`
	Value param.Field[OrgUserListParamsQOrderByValue] `query:"value"`
}

// URLQuery serializes [OrgUserListParamsQOrderBy]'s query parameters as
// `url.Values`.
func (r OrgUserListParamsQOrderBy) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type OrgUserListParamsQOrderByValue string

const (
	OrgUserListParamsQOrderByValueAsc  OrgUserListParamsQOrderByValue = "ASC"
	OrgUserListParamsQOrderByValueDesc OrgUserListParamsQOrderByValue = "DESC"
)

func (r OrgUserListParamsQOrderByValue) IsKnown() bool {
	switch r {
	case OrgUserListParamsQOrderByValueAsc, OrgUserListParamsQOrderByValueDesc:
		return true
	}
	return false
}

type OrgUserDeleteParams struct {
	// If anonymize is true, then org owner permission is required.
	Anonymize param.Field[bool] `query:"anonymize"`
}

// URLQuery serializes [OrgUserDeleteParams]'s query parameters as `url.Values`.
func (r OrgUserDeleteParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type OrgUserAddRoleParams struct {
	Roles     param.Field[[]string] `query:"roles,required"`
	Ncid      param.Field[string]   `cookie:"ncid"`
	VisitorID param.Field[string]   `cookie:"VisitorID"`
}

// URLQuery serializes [OrgUserAddRoleParams]'s query parameters as `url.Values`.
func (r OrgUserAddRoleParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type OrgUserRemoveRoleParams struct {
	Roles param.Field[[]string] `query:"roles"`
}

// URLQuery serializes [OrgUserRemoveRoleParams]'s query parameters as
// `url.Values`.
func (r OrgUserRemoveRoleParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type OrgUserUpdateRoleParams struct {
	Roles param.Field[[]string] `query:"roles"`
}

// URLQuery serializes [OrgUserUpdateRoleParams]'s query parameters as
// `url.Values`.
func (r OrgUserUpdateRoleParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
