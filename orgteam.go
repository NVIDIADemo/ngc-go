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

// OrgTeamService contains methods and other services that help with interacting
// with the ngc API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewOrgTeamService] method instead.
type OrgTeamService struct {
	Options        []option.RequestOption
	Users          *OrgTeamUserService
	StarfleetIDs   *OrgTeamStarfleetIDService
	NcaInvitations *OrgTeamNcaInvitationService
}

// NewOrgTeamService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewOrgTeamService(opts ...option.RequestOption) (r *OrgTeamService) {
	r = &OrgTeamService{}
	r.Options = opts
	r.Users = NewOrgTeamUserService(opts...)
	r.StarfleetIDs = NewOrgTeamStarfleetIDService(opts...)
	r.NcaInvitations = NewOrgTeamNcaInvitationService(opts...)
	return
}

// List all Teams
func (r *OrgTeamService) List(ctx context.Context, orgName string, query OrgTeamListParams, opts ...option.RequestOption) (res *pagination.PageNumberTeams[shared.OrgTeamListResponse], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if orgName == "" {
		err = errors.New("missing required org-name parameter")
		return
	}
	path := fmt.Sprintf("v2/org/%s/teams", orgName)
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

// List all Teams
func (r *OrgTeamService) ListAutoPaging(ctx context.Context, orgName string, query OrgTeamListParams, opts ...option.RequestOption) *pagination.PageNumberTeamsAutoPager[shared.OrgTeamListResponse] {
	return pagination.NewPageNumberTeamsAutoPager(r.List(ctx, orgName, query, opts...))
}

// details about one team
type TeamResponse struct {
	RequestStatus TeamResponseRequestStatus `json:"requestStatus"`
	// Information about the team
	Team TeamResponseTeam `json:"team"`
	JSON teamResponseJSON `json:"-"`
}

// teamResponseJSON contains the JSON metadata for the struct [TeamResponse]
type teamResponseJSON struct {
	RequestStatus apijson.Field
	Team          apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *TeamResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r teamResponseJSON) RawJSON() string {
	return r.raw
}

type TeamResponseRequestStatus struct {
	RequestID string `json:"requestId"`
	ServerID  string `json:"serverId"`
	// Describes response status reported by the server.
	StatusCode        TeamResponseRequestStatusStatusCode `json:"statusCode"`
	StatusDescription string                              `json:"statusDescription"`
	JSON              teamResponseRequestStatusJSON       `json:"-"`
}

// teamResponseRequestStatusJSON contains the JSON metadata for the struct
// [TeamResponseRequestStatus]
type teamResponseRequestStatusJSON struct {
	RequestID         apijson.Field
	ServerID          apijson.Field
	StatusCode        apijson.Field
	StatusDescription apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *TeamResponseRequestStatus) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r teamResponseRequestStatusJSON) RawJSON() string {
	return r.raw
}

// Describes response status reported by the server.
type TeamResponseRequestStatusStatusCode string

const (
	TeamResponseRequestStatusStatusCodeUnknown                    TeamResponseRequestStatusStatusCode = "UNKNOWN"
	TeamResponseRequestStatusStatusCodeSuccess                    TeamResponseRequestStatusStatusCode = "SUCCESS"
	TeamResponseRequestStatusStatusCodeUnauthorized               TeamResponseRequestStatusStatusCode = "UNAUTHORIZED"
	TeamResponseRequestStatusStatusCodePaymentRequired            TeamResponseRequestStatusStatusCode = "PAYMENT_REQUIRED"
	TeamResponseRequestStatusStatusCodeForbidden                  TeamResponseRequestStatusStatusCode = "FORBIDDEN"
	TeamResponseRequestStatusStatusCodeTimeout                    TeamResponseRequestStatusStatusCode = "TIMEOUT"
	TeamResponseRequestStatusStatusCodeExists                     TeamResponseRequestStatusStatusCode = "EXISTS"
	TeamResponseRequestStatusStatusCodeNotFound                   TeamResponseRequestStatusStatusCode = "NOT_FOUND"
	TeamResponseRequestStatusStatusCodeInternalError              TeamResponseRequestStatusStatusCode = "INTERNAL_ERROR"
	TeamResponseRequestStatusStatusCodeInvalidRequest             TeamResponseRequestStatusStatusCode = "INVALID_REQUEST"
	TeamResponseRequestStatusStatusCodeInvalidRequestVersion      TeamResponseRequestStatusStatusCode = "INVALID_REQUEST_VERSION"
	TeamResponseRequestStatusStatusCodeInvalidRequestData         TeamResponseRequestStatusStatusCode = "INVALID_REQUEST_DATA"
	TeamResponseRequestStatusStatusCodeMethodNotAllowed           TeamResponseRequestStatusStatusCode = "METHOD_NOT_ALLOWED"
	TeamResponseRequestStatusStatusCodeConflict                   TeamResponseRequestStatusStatusCode = "CONFLICT"
	TeamResponseRequestStatusStatusCodeUnprocessableEntity        TeamResponseRequestStatusStatusCode = "UNPROCESSABLE_ENTITY"
	TeamResponseRequestStatusStatusCodeTooManyRequests            TeamResponseRequestStatusStatusCode = "TOO_MANY_REQUESTS"
	TeamResponseRequestStatusStatusCodeInsufficientStorage        TeamResponseRequestStatusStatusCode = "INSUFFICIENT_STORAGE"
	TeamResponseRequestStatusStatusCodeServiceUnavailable         TeamResponseRequestStatusStatusCode = "SERVICE_UNAVAILABLE"
	TeamResponseRequestStatusStatusCodePayloadTooLarge            TeamResponseRequestStatusStatusCode = "PAYLOAD_TOO_LARGE"
	TeamResponseRequestStatusStatusCodeNotAcceptable              TeamResponseRequestStatusStatusCode = "NOT_ACCEPTABLE"
	TeamResponseRequestStatusStatusCodeUnavailableForLegalReasons TeamResponseRequestStatusStatusCode = "UNAVAILABLE_FOR_LEGAL_REASONS"
	TeamResponseRequestStatusStatusCodeBadGateway                 TeamResponseRequestStatusStatusCode = "BAD_GATEWAY"
)

func (r TeamResponseRequestStatusStatusCode) IsKnown() bool {
	switch r {
	case TeamResponseRequestStatusStatusCodeUnknown, TeamResponseRequestStatusStatusCodeSuccess, TeamResponseRequestStatusStatusCodeUnauthorized, TeamResponseRequestStatusStatusCodePaymentRequired, TeamResponseRequestStatusStatusCodeForbidden, TeamResponseRequestStatusStatusCodeTimeout, TeamResponseRequestStatusStatusCodeExists, TeamResponseRequestStatusStatusCodeNotFound, TeamResponseRequestStatusStatusCodeInternalError, TeamResponseRequestStatusStatusCodeInvalidRequest, TeamResponseRequestStatusStatusCodeInvalidRequestVersion, TeamResponseRequestStatusStatusCodeInvalidRequestData, TeamResponseRequestStatusStatusCodeMethodNotAllowed, TeamResponseRequestStatusStatusCodeConflict, TeamResponseRequestStatusStatusCodeUnprocessableEntity, TeamResponseRequestStatusStatusCodeTooManyRequests, TeamResponseRequestStatusStatusCodeInsufficientStorage, TeamResponseRequestStatusStatusCodeServiceUnavailable, TeamResponseRequestStatusStatusCodePayloadTooLarge, TeamResponseRequestStatusStatusCodeNotAcceptable, TeamResponseRequestStatusStatusCodeUnavailableForLegalReasons, TeamResponseRequestStatusStatusCodeBadGateway:
		return true
	}
	return false
}

// Information about the team
type TeamResponseTeam struct {
	// unique Id of this team.
	ID int64 `json:"id"`
	// description of the team
	Description string `json:"description"`
	// Infinity manager setting definition
	InfinityManagerSettings TeamResponseTeamInfinityManagerSettings `json:"infinityManagerSettings"`
	// indicates if the team is deleted or not
	IsDeleted bool `json:"isDeleted"`
	// team name
	Name string `json:"name"`
	// Repo scan setting definition
	RepoScanSettings TeamResponseTeamRepoScanSettings `json:"repoScanSettings"`
	JSON             teamResponseTeamJSON             `json:"-"`
}

// teamResponseTeamJSON contains the JSON metadata for the struct
// [TeamResponseTeam]
type teamResponseTeamJSON struct {
	ID                      apijson.Field
	Description             apijson.Field
	InfinityManagerSettings apijson.Field
	IsDeleted               apijson.Field
	Name                    apijson.Field
	RepoScanSettings        apijson.Field
	raw                     string
	ExtraFields             map[string]apijson.Field
}

func (r *TeamResponseTeam) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r teamResponseTeamJSON) RawJSON() string {
	return r.raw
}

// Infinity manager setting definition
type TeamResponseTeamInfinityManagerSettings struct {
	// Enable the infinity manager or not. Used both in org and team level object
	InfinityManagerEnabled bool `json:"infinityManagerEnabled"`
	// Allow override settings at team level. Only used in org level object
	InfinityManagerEnableTeamOverride bool                                        `json:"infinityManagerEnableTeamOverride"`
	JSON                              teamResponseTeamInfinityManagerSettingsJSON `json:"-"`
}

// teamResponseTeamInfinityManagerSettingsJSON contains the JSON metadata for the
// struct [TeamResponseTeamInfinityManagerSettings]
type teamResponseTeamInfinityManagerSettingsJSON struct {
	InfinityManagerEnabled            apijson.Field
	InfinityManagerEnableTeamOverride apijson.Field
	raw                               string
	ExtraFields                       map[string]apijson.Field
}

func (r *TeamResponseTeamInfinityManagerSettings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r teamResponseTeamInfinityManagerSettingsJSON) RawJSON() string {
	return r.raw
}

// Repo scan setting definition
type TeamResponseTeamRepoScanSettings struct {
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
	RepoScanShowResults bool                                 `json:"repoScanShowResults"`
	JSON                teamResponseTeamRepoScanSettingsJSON `json:"-"`
}

// teamResponseTeamRepoScanSettingsJSON contains the JSON metadata for the struct
// [TeamResponseTeamRepoScanSettings]
type teamResponseTeamRepoScanSettingsJSON struct {
	RepoScanAllowOverride       apijson.Field
	RepoScanByDefault           apijson.Field
	RepoScanEnabled             apijson.Field
	RepoScanEnableNotifications apijson.Field
	RepoScanEnableTeamOverride  apijson.Field
	RepoScanShowResults         apijson.Field
	raw                         string
	ExtraFields                 map[string]apijson.Field
}

func (r *TeamResponseTeamRepoScanSettings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r teamResponseTeamRepoScanSettingsJSON) RawJSON() string {
	return r.raw
}

// Information about the team
type OrgTeamListResponse struct {
	// unique Id of this team.
	ID int64 `json:"id"`
	// description of the team
	Description string `json:"description"`
	// Infinity manager setting definition
	InfinityManagerSettings OrgTeamListResponseInfinityManagerSettings `json:"infinityManagerSettings"`
	// indicates if the team is deleted or not
	IsDeleted bool `json:"isDeleted"`
	// team name
	Name string `json:"name"`
	// Repo scan setting definition
	RepoScanSettings OrgTeamListResponseRepoScanSettings `json:"repoScanSettings"`
	JSON             orgTeamListResponseJSON             `json:"-"`
}

// orgTeamListResponseJSON contains the JSON metadata for the struct
// [OrgTeamListResponse]
type orgTeamListResponseJSON struct {
	ID                      apijson.Field
	Description             apijson.Field
	InfinityManagerSettings apijson.Field
	IsDeleted               apijson.Field
	Name                    apijson.Field
	RepoScanSettings        apijson.Field
	raw                     string
	ExtraFields             map[string]apijson.Field
}

func (r *OrgTeamListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r orgTeamListResponseJSON) RawJSON() string {
	return r.raw
}

// Infinity manager setting definition
type OrgTeamListResponseInfinityManagerSettings struct {
	// Enable the infinity manager or not. Used both in org and team level object
	InfinityManagerEnabled bool `json:"infinityManagerEnabled"`
	// Allow override settings at team level. Only used in org level object
	InfinityManagerEnableTeamOverride bool                                           `json:"infinityManagerEnableTeamOverride"`
	JSON                              orgTeamListResponseInfinityManagerSettingsJSON `json:"-"`
}

// orgTeamListResponseInfinityManagerSettingsJSON contains the JSON metadata for
// the struct [OrgTeamListResponseInfinityManagerSettings]
type orgTeamListResponseInfinityManagerSettingsJSON struct {
	InfinityManagerEnabled            apijson.Field
	InfinityManagerEnableTeamOverride apijson.Field
	raw                               string
	ExtraFields                       map[string]apijson.Field
}

func (r *OrgTeamListResponseInfinityManagerSettings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r orgTeamListResponseInfinityManagerSettingsJSON) RawJSON() string {
	return r.raw
}

// Repo scan setting definition
type OrgTeamListResponseRepoScanSettings struct {
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
	RepoScanShowResults bool                                    `json:"repoScanShowResults"`
	JSON                orgTeamListResponseRepoScanSettingsJSON `json:"-"`
}

// orgTeamListResponseRepoScanSettingsJSON contains the JSON metadata for the
// struct [OrgTeamListResponseRepoScanSettings]
type orgTeamListResponseRepoScanSettingsJSON struct {
	RepoScanAllowOverride       apijson.Field
	RepoScanByDefault           apijson.Field
	RepoScanEnabled             apijson.Field
	RepoScanEnableNotifications apijson.Field
	RepoScanEnableTeamOverride  apijson.Field
	RepoScanShowResults         apijson.Field
	raw                         string
	ExtraFields                 map[string]apijson.Field
}

func (r *OrgTeamListResponseRepoScanSettings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r orgTeamListResponseRepoScanSettingsJSON) RawJSON() string {
	return r.raw
}

type OrgTeamListParams struct {
	// The page number of result
	PageNumber param.Field[int64] `query:"page-number"`
	// The page size of result
	PageSize param.Field[int64] `query:"page-size"`
}

// URLQuery serializes [OrgTeamListParams]'s query parameters as `url.Values`.
func (r OrgTeamListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
