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

func TestAdminOrgValidateGetAll(t *testing.T) {
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
	_, err := client.Admin.Org.Validate.GetAll(context.TODO(), ngc.AdminOrgValidateGetAllParams{
		Q: ngc.F(ngc.AdminOrgValidateGetAllParamsQ{
			OrgOwner: ngc.F(ngc.AdminOrgValidateGetAllParamsQOrgOwner{
				Email:         ngc.F("email"),
				FullName:      ngc.F("x"),
				LastLoginDate: ngc.F("lastLoginDate"),
			}),
			PecSfdcID: ngc.F("pecSfdcId"),
			ProductSubscriptions: ngc.F([]ngc.AdminOrgValidateGetAllParamsQProductSubscription{{
				ProductName:        ngc.F("productName"),
				ID:                 ngc.F("id"),
				EmsEntitlementType: ngc.F(ngc.AdminOrgValidateGetAllParamsQProductSubscriptionsEmsEntitlementTypeEmsEval),
				ExpirationDate:     ngc.F("expirationDate"),
				StartDate:          ngc.F("startDate"),
				Type:               ngc.F(ngc.AdminOrgValidateGetAllParamsQProductSubscriptionsTypeNgcAdminEval),
			}, {
				ProductName:        ngc.F("productName"),
				ID:                 ngc.F("id"),
				EmsEntitlementType: ngc.F(ngc.AdminOrgValidateGetAllParamsQProductSubscriptionsEmsEntitlementTypeEmsEval),
				ExpirationDate:     ngc.F("expirationDate"),
				StartDate:          ngc.F("startDate"),
				Type:               ngc.F(ngc.AdminOrgValidateGetAllParamsQProductSubscriptionsTypeNgcAdminEval),
			}, {
				ProductName:        ngc.F("productName"),
				ID:                 ngc.F("id"),
				EmsEntitlementType: ngc.F(ngc.AdminOrgValidateGetAllParamsQProductSubscriptionsEmsEntitlementTypeEmsEval),
				ExpirationDate:     ngc.F("expirationDate"),
				StartDate:          ngc.F("startDate"),
				Type:               ngc.F(ngc.AdminOrgValidateGetAllParamsQProductSubscriptionsTypeNgcAdminEval),
			}}),
		}),
	})
	if err != nil {
		var apierr *ngc.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
