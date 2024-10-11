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
	"github.com/NVIDIADemo/ngc-go/shared"
)

// OrgTeamNcaInvitationService contains methods and other services that help with
// interacting with the ngc API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewOrgTeamNcaInvitationService] method instead.
type OrgTeamNcaInvitationService struct {
	Options []option.RequestOption
}

// NewOrgTeamNcaInvitationService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewOrgTeamNcaInvitationService(opts ...option.RequestOption) (r *OrgTeamNcaInvitationService) {
	r = &OrgTeamNcaInvitationService{}
	r.Options = opts
	return
}

// Invites and creates a User in team
func (r *OrgTeamNcaInvitationService) New(ctx context.Context, orgName string, teamName string, body OrgTeamNcaInvitationNewParams, opts ...option.RequestOption) (res *shared.User, err error) {
	opts = append(r.Options[:], opts...)
	if orgName == "" {
		err = errors.New("missing required org-name parameter")
		return
	}
	if teamName == "" {
		err = errors.New("missing required team-name parameter")
		return
	}
	path := fmt.Sprintf("v3/orgs/%s/teams/%s/users/nca-invitations", orgName, teamName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type OrgTeamNcaInvitationNewParams struct {
	// Is the user email
	Email param.Field[string] `json:"email"`
	// Is the numbers of days the invitation will expire
	InvitationExpirationIn param.Field[int64] `json:"invitationExpirationIn"`
	// Nca allow users to be invited as Admin and as Member
	InviteAs param.Field[OrgTeamNcaInvitationNewParamsInviteAs] `json:"inviteAs"`
	// Is a message to the new user
	Message param.Field[string] `json:"message"`
}

func (r OrgTeamNcaInvitationNewParams) MarshalJSON() (data []byte, err error) {
	return apijson.MarshalRoot(r)
}

// Nca allow users to be invited as Admin and as Member
type OrgTeamNcaInvitationNewParamsInviteAs string

const (
	OrgTeamNcaInvitationNewParamsInviteAsAdmin  OrgTeamNcaInvitationNewParamsInviteAs = "ADMIN"
	OrgTeamNcaInvitationNewParamsInviteAsMember OrgTeamNcaInvitationNewParamsInviteAs = "MEMBER"
)

func (r OrgTeamNcaInvitationNewParamsInviteAs) IsKnown() bool {
	switch r {
	case OrgTeamNcaInvitationNewParamsInviteAsAdmin, OrgTeamNcaInvitationNewParamsInviteAsMember:
		return true
	}
	return false
}
