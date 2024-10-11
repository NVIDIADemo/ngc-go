// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package ngc

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"

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
func (r *OrgTeamService) List(ctx context.Context, orgName string, query OrgTeamListParams, opts ...option.RequestOption) (res *pagination.PageNumberTeams[shared.Team], err error) {
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
func (r *OrgTeamService) ListAutoPaging(ctx context.Context, orgName string, query OrgTeamListParams, opts ...option.RequestOption) *pagination.PageNumberTeamsAutoPager[shared.Team] {
	return pagination.NewPageNumberTeamsAutoPager(r.List(ctx, orgName, query, opts...))
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
