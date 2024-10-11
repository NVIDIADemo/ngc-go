// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package ngc

import (
	"context"
	"net/http"
	"net/url"

	"github.com/NVIDIADemo/ngc-go/internal/apiquery"
	"github.com/NVIDIADemo/ngc-go/internal/param"
	"github.com/NVIDIADemo/ngc-go/internal/requestconfig"
	"github.com/NVIDIADemo/ngc-go/option"
	"github.com/NVIDIADemo/ngc-go/shared"
)

// HealthAllService contains methods and other services that help with interacting
// with the ngc API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewHealthAllService] method instead.
type HealthAllService struct {
	Options []option.RequestOption
}

// NewHealthAllService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewHealthAllService(opts ...option.RequestOption) (r *HealthAllService) {
	r = &HealthAllService{}
	r.Options = opts
	return
}

// Used to get health status of all services
func (r *HealthAllService) GetAll(ctx context.Context, query HealthAllGetAllParams, opts ...option.RequestOption) (res *shared.Health, err error) {
	opts = append(r.Options[:], opts...)
	path := "health/all"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type HealthAllGetAllParams struct {
	// secret value that validates the call in order to show details
	Secret param.Field[string] `query:"secret"`
	// boolean value to indicating to show details or not
	ShowDetails param.Field[bool] `query:"showDetails"`
}

// URLQuery serializes [HealthAllGetAllParams]'s query parameters as `url.Values`.
func (r HealthAllGetAllParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
