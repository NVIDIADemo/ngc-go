// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package ngc

import (
	"github.com/NVIDIADemo/ngc-go/option"
)

// MeService contains methods and other services that help with interacting with
// the ngc API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewMeService] method instead.
type MeService struct {
	Options []option.RequestOption
	APIKey  *MeAPIKeyService
}

// NewMeService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewMeService(opts ...option.RequestOption) (r *MeService) {
	r = &MeService{}
	r.Options = opts
	r.APIKey = NewMeAPIKeyService(opts...)
	return
}
