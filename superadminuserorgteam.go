// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package ngc

import (
	"github.com/NVIDIADemo/ngc-go/option"
)

// SuperAdminUserOrgTeamService contains methods and other services that help with
// interacting with the ngc API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSuperAdminUserOrgTeamService] method instead.
type SuperAdminUserOrgTeamService struct {
	Options []option.RequestOption
	Users   *SuperAdminUserOrgTeamUserService
}

// NewSuperAdminUserOrgTeamService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewSuperAdminUserOrgTeamService(opts ...option.RequestOption) (r *SuperAdminUserOrgTeamService) {
	r = &SuperAdminUserOrgTeamService{}
	r.Options = opts
	r.Users = NewSuperAdminUserOrgTeamUserService(opts...)
	return
}
