// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package ngc

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"

	"github.com/NVIDIADemo/ngc-go/internal/apiquery"
	"github.com/NVIDIADemo/ngc-go/internal/param"
	"github.com/NVIDIADemo/ngc-go/internal/requestconfig"
	"github.com/NVIDIADemo/ngc-go/option"
	"github.com/NVIDIADemo/ngc-go/shared"
)

// OrgMeteringService contains methods and other services that help with
// interacting with the ngc API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewOrgMeteringService] method instead.
type OrgMeteringService struct {
	Options []option.RequestOption
	Gpupeak *OrgMeteringGpupeakService
}

// NewOrgMeteringService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewOrgMeteringService(opts ...option.RequestOption) (r *OrgMeteringService) {
	r = &OrgMeteringService{}
	r.Options = opts
	r.Gpupeak = NewOrgMeteringGpupeakService(opts...)
	return
}

// Returns Private Registry / EGX resources usage metering as measurement series.
// Requires admin privileges for organization.
func (r *OrgMeteringService) GetAll(ctx context.Context, orgName string, query OrgMeteringGetAllParams, opts ...option.RequestOption) (res *shared.MeteringResultList, err error) {
	opts = append(r.Options[:], opts...)
	if orgName == "" {
		err = errors.New("missing required org-name parameter")
		return
	}
	path := fmt.Sprintf("v2/org/%s/metering", orgName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type OrgMeteringGetAllParams struct {
	// request params for getting metering usage
	Q param.Field[OrgMeteringGetAllParamsQ] `query:"q,required"`
}

// URLQuery serializes [OrgMeteringGetAllParams]'s query parameters as
// `url.Values`.
func (r OrgMeteringGetAllParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// request params for getting metering usage
type OrgMeteringGetAllParamsQ struct {
	Measurements param.Field[[]OrgMeteringGetAllParamsQMeasurement] `query:"measurements"`
}

// URLQuery serializes [OrgMeteringGetAllParamsQ]'s query parameters as
// `url.Values`.
func (r OrgMeteringGetAllParamsQ) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// object used for sending metering query parameter request
type OrgMeteringGetAllParamsQMeasurement struct {
	// this replaces all null values in an output stream with a non-null value that is
	// provided.
	Fill param.Field[float64] `query:"fill"`
	// end time range for the data, in ISO formate, yyyy-MM-dd'T'HH:mm:ss.SSS'Z'
	FromDate param.Field[string] `query:"fromDate"`
	// group by specific tags
	GroupBy param.Field[[]string] `query:"groupBy"`
	// time period to aggregate the data over with, in seconds. If none provided, raw
	// data will be returned.
	PeriodInSeconds param.Field[float64] `query:"periodInSeconds"`
	// start time range for the data, in ISO formate, yyyy-MM-dd'T'HH:mm:ss.SSS'Z'
	ToDate param.Field[string]                                   `query:"toDate"`
	Type   param.Field[OrgMeteringGetAllParamsQMeasurementsType] `query:"type"`
}

// URLQuery serializes [OrgMeteringGetAllParamsQMeasurement]'s query parameters as
// `url.Values`.
func (r OrgMeteringGetAllParamsQMeasurement) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type OrgMeteringGetAllParamsQMeasurementsType string

const (
	OrgMeteringGetAllParamsQMeasurementsTypeEgxGPUUtilizationDaily                   OrgMeteringGetAllParamsQMeasurementsType = "EGX_GPU_UTILIZATION_DAILY"
	OrgMeteringGetAllParamsQMeasurementsTypeFleetCommandGPUUtilizationDaily          OrgMeteringGetAllParamsQMeasurementsType = "FLEET_COMMAND_GPU_UTILIZATION_DAILY"
	OrgMeteringGetAllParamsQMeasurementsTypeEgxLogStorageUtilizationDaily            OrgMeteringGetAllParamsQMeasurementsType = "EGX_LOG_STORAGE_UTILIZATION_DAILY"
	OrgMeteringGetAllParamsQMeasurementsTypeFleetCommandLogStorageUtilizationDaily   OrgMeteringGetAllParamsQMeasurementsType = "FLEET_COMMAND_LOG_STORAGE_UTILIZATION_DAILY"
	OrgMeteringGetAllParamsQMeasurementsTypeRegistryStorageUtilizationDaily          OrgMeteringGetAllParamsQMeasurementsType = "REGISTRY_STORAGE_UTILIZATION_DAILY"
	OrgMeteringGetAllParamsQMeasurementsTypeEgxGPUUtilizationMonthly                 OrgMeteringGetAllParamsQMeasurementsType = "EGX_GPU_UTILIZATION_MONTHLY"
	OrgMeteringGetAllParamsQMeasurementsTypeFleetCommandGPUUtilizationMonthly        OrgMeteringGetAllParamsQMeasurementsType = "FLEET_COMMAND_GPU_UTILIZATION_MONTHLY"
	OrgMeteringGetAllParamsQMeasurementsTypeEgxLogStorageUtilizationMonthly          OrgMeteringGetAllParamsQMeasurementsType = "EGX_LOG_STORAGE_UTILIZATION_MONTHLY"
	OrgMeteringGetAllParamsQMeasurementsTypeFleetCommandLogStorageUtilizationMonthly OrgMeteringGetAllParamsQMeasurementsType = "FLEET_COMMAND_LOG_STORAGE_UTILIZATION_MONTHLY"
	OrgMeteringGetAllParamsQMeasurementsTypeRegistryStorageUtilizationMonthly        OrgMeteringGetAllParamsQMeasurementsType = "REGISTRY_STORAGE_UTILIZATION_MONTHLY"
)

func (r OrgMeteringGetAllParamsQMeasurementsType) IsKnown() bool {
	switch r {
	case OrgMeteringGetAllParamsQMeasurementsTypeEgxGPUUtilizationDaily, OrgMeteringGetAllParamsQMeasurementsTypeFleetCommandGPUUtilizationDaily, OrgMeteringGetAllParamsQMeasurementsTypeEgxLogStorageUtilizationDaily, OrgMeteringGetAllParamsQMeasurementsTypeFleetCommandLogStorageUtilizationDaily, OrgMeteringGetAllParamsQMeasurementsTypeRegistryStorageUtilizationDaily, OrgMeteringGetAllParamsQMeasurementsTypeEgxGPUUtilizationMonthly, OrgMeteringGetAllParamsQMeasurementsTypeFleetCommandGPUUtilizationMonthly, OrgMeteringGetAllParamsQMeasurementsTypeEgxLogStorageUtilizationMonthly, OrgMeteringGetAllParamsQMeasurementsTypeFleetCommandLogStorageUtilizationMonthly, OrgMeteringGetAllParamsQMeasurementsTypeRegistryStorageUtilizationMonthly:
		return true
	}
	return false
}
