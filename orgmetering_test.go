// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package ngc_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/NVIDIADemo/ngc-go"
	"github.com/NVIDIADemo/ngc-go/internal/testutil"
	"github.com/NVIDIADemo/ngc-go/option"
)

func TestOrgMeteringGetAllWithOptionalParams(t *testing.T) {
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := ngc.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAuthToken("My Auth Token"),
	)
	_, err := client.Orgs.Metering.GetAll(
		context.TODO(),
		"org-name",
		ngc.OrgMeteringGetAllParams{
			Q: ngc.F(ngc.OrgMeteringGetAllParamsQ{
				Measurements: ngc.F([]ngc.OrgMeteringGetAllParamsQMeasurement{{
					Fill:            ngc.F(0.000000),
					FromDate:        ngc.F("fromDate"),
					GroupBy:         ngc.F([]string{"string", "string", "string"}),
					PeriodInSeconds: ngc.F(0.000000),
					ToDate:          ngc.F("toDate"),
					Type:            ngc.F(ngc.OrgMeteringGetAllParamsQMeasurementsTypeEgxGPUUtilizationDaily),
				}, {
					Fill:            ngc.F(0.000000),
					FromDate:        ngc.F("fromDate"),
					GroupBy:         ngc.F([]string{"string", "string", "string"}),
					PeriodInSeconds: ngc.F(0.000000),
					ToDate:          ngc.F("toDate"),
					Type:            ngc.F(ngc.OrgMeteringGetAllParamsQMeasurementsTypeEgxGPUUtilizationDaily),
				}, {
					Fill:            ngc.F(0.000000),
					FromDate:        ngc.F("fromDate"),
					GroupBy:         ngc.F([]string{"string", "string", "string"}),
					PeriodInSeconds: ngc.F(0.000000),
					ToDate:          ngc.F("toDate"),
					Type:            ngc.F(ngc.OrgMeteringGetAllParamsQMeasurementsTypeEgxGPUUtilizationDaily),
				}}),
			}),
		},
	)
	if err != nil {
		var apierr *ngc.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
