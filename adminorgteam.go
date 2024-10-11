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

// AdminOrgTeamService contains methods and other services that help with
// interacting with the ngc API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAdminOrgTeamService] method instead.
type AdminOrgTeamService struct {
	Options []option.RequestOption
}

// NewAdminOrgTeamService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewAdminOrgTeamService(opts ...option.RequestOption) (r *AdminOrgTeamService) {
	r = &AdminOrgTeamService{}
	r.Options = opts
	return
}

// List all Teams
func (r *AdminOrgTeamService) List(ctx context.Context, orgName string, query AdminOrgTeamListParams, opts ...option.RequestOption) (res *pagination.PageNumberTeams[shared.AdminOrgTeamListResponse], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if orgName == "" {
		err = errors.New("missing required org-name parameter")
		return
	}
	path := fmt.Sprintf("v2/admin/org/%s/teams", orgName)
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
func (r *AdminOrgTeamService) ListAutoPaging(ctx context.Context, orgName string, query AdminOrgTeamListParams, opts ...option.RequestOption) *pagination.PageNumberTeamsAutoPager[shared.AdminOrgTeamListResponse] {
	return pagination.NewPageNumberTeamsAutoPager(r.List(ctx, orgName, query, opts...))
}

// Information about the team
type AdminOrgTeamListResponse struct {
	// unique Id of this team.
	ID int64 `json:"id"`
	// description of the team
	Description string `json:"description"`
	// Infinity manager setting definition
	InfinityManagerSettings AdminOrgTeamListResponseInfinityManagerSettings `json:"infinityManagerSettings"`
	// indicates if the team is deleted or not
	IsDeleted bool `json:"isDeleted"`
	// team name
	Name string `json:"name"`
	// Repo scan setting definition
	RepoScanSettings AdminOrgTeamListResponseRepoScanSettings `json:"repoScanSettings"`
	JSON             adminOrgTeamListResponseJSON             `json:"-"`
}

// adminOrgTeamListResponseJSON contains the JSON metadata for the struct
// [AdminOrgTeamListResponse]
type adminOrgTeamListResponseJSON struct {
	ID                      apijson.Field
	Description             apijson.Field
	InfinityManagerSettings apijson.Field
	IsDeleted               apijson.Field
	Name                    apijson.Field
	RepoScanSettings        apijson.Field
	raw                     string
	ExtraFields             map[string]apijson.Field
}

func (r *AdminOrgTeamListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r adminOrgTeamListResponseJSON) RawJSON() string {
	return r.raw
}

// Infinity manager setting definition
type AdminOrgTeamListResponseInfinityManagerSettings struct {
	// Enable the infinity manager or not. Used both in org and team level object
	InfinityManagerEnabled bool `json:"infinityManagerEnabled"`
	// Allow override settings at team level. Only used in org level object
	InfinityManagerEnableTeamOverride bool                                                `json:"infinityManagerEnableTeamOverride"`
	JSON                              adminOrgTeamListResponseInfinityManagerSettingsJSON `json:"-"`
}

// adminOrgTeamListResponseInfinityManagerSettingsJSON contains the JSON metadata
// for the struct [AdminOrgTeamListResponseInfinityManagerSettings]
type adminOrgTeamListResponseInfinityManagerSettingsJSON struct {
	InfinityManagerEnabled            apijson.Field
	InfinityManagerEnableTeamOverride apijson.Field
	raw                               string
	ExtraFields                       map[string]apijson.Field
}

func (r *AdminOrgTeamListResponseInfinityManagerSettings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r adminOrgTeamListResponseInfinityManagerSettingsJSON) RawJSON() string {
	return r.raw
}

// Repo scan setting definition
type AdminOrgTeamListResponseRepoScanSettings struct {
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
	RepoScanShowResults bool                                         `json:"repoScanShowResults"`
	JSON                adminOrgTeamListResponseRepoScanSettingsJSON `json:"-"`
}

// adminOrgTeamListResponseRepoScanSettingsJSON contains the JSON metadata for the
// struct [AdminOrgTeamListResponseRepoScanSettings]
type adminOrgTeamListResponseRepoScanSettingsJSON struct {
	RepoScanAllowOverride       apijson.Field
	RepoScanByDefault           apijson.Field
	RepoScanEnabled             apijson.Field
	RepoScanEnableNotifications apijson.Field
	RepoScanEnableTeamOverride  apijson.Field
	RepoScanShowResults         apijson.Field
	raw                         string
	ExtraFields                 map[string]apijson.Field
}

func (r *AdminOrgTeamListResponseRepoScanSettings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r adminOrgTeamListResponseRepoScanSettingsJSON) RawJSON() string {
	return r.raw
}

type AdminOrgTeamListParams struct {
	// The page number of result
	PageNumber param.Field[int64] `query:"page-number"`
	// The page size of result
	PageSize param.Field[int64] `query:"page-size"`
	// Get all team that has scan allow override only
	ScanAllowOverride param.Field[bool] `query:"scan-allow-override"`
	// Get all team that has scan by default only
	ScanByDefault param.Field[bool] `query:"scan-by-default"`
	// Get all team that has scan show results only
	ScanShowResults param.Field[bool] `query:"scan-show-results"`
	// Team name to search
	TeamName param.Field[string] `query:"team-name"`
}

// URLQuery serializes [AdminOrgTeamListParams]'s query parameters as `url.Values`.
func (r AdminOrgTeamListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
