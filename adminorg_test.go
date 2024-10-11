// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package ngc_test

import (
	"bytes"
	"context"
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/NVIDIADemo/ngc-go"
	"github.com/NVIDIADemo/ngc-go/internal/testutil"
	"github.com/NVIDIADemo/ngc-go/option"
)

func TestAdminOrgNewWithOptionalParams(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		w.Write([]byte("abc"))
	}))
	defer server.Close()
	baseURL := server.URL
	client := ngc.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAuthToken("My Auth Token"),
	)
	resp, err := client.Admin.Orgs.New(context.TODO(), ngc.AdminOrgNewParams{
		OrgOwner: ngc.F(ngc.AdminOrgNewParamsOrgOwner{
			Email:         ngc.F("email"),
			FullName:      ngc.F("x"),
			LastLoginDate: ngc.F("lastLoginDate"),
		}),
		Country:     ngc.F("country"),
		Description: ngc.F("description"),
		DisplayName: ngc.F("x"),
		IdpID:       ngc.F("idpId"),
		IsInternal:  ngc.F(true),
		Name:        ngc.F("xx"),
		PecName:     ngc.F("pecName"),
		PecSfdcID:   ngc.F("pecSfdcId"),
		ProductEnablements: ngc.F([]ngc.AdminOrgNewParamsProductEnablement{{
			ProductName:    ngc.F("productName"),
			Type:           ngc.F(ngc.AdminOrgNewParamsProductEnablementsTypeNgcAdminEval),
			ExpirationDate: ngc.F("expirationDate"),
			PoDetails: ngc.F([]ngc.AdminOrgNewParamsProductEnablementsPoDetail{{
				EntitlementID: ngc.F("entitlementId"),
				PkID:          ngc.F("pkId"),
			}, {
				EntitlementID: ngc.F("entitlementId"),
				PkID:          ngc.F("pkId"),
			}, {
				EntitlementID: ngc.F("entitlementId"),
				PkID:          ngc.F("pkId"),
			}}),
		}, {
			ProductName:    ngc.F("productName"),
			Type:           ngc.F(ngc.AdminOrgNewParamsProductEnablementsTypeNgcAdminEval),
			ExpirationDate: ngc.F("expirationDate"),
			PoDetails: ngc.F([]ngc.AdminOrgNewParamsProductEnablementsPoDetail{{
				EntitlementID: ngc.F("entitlementId"),
				PkID:          ngc.F("pkId"),
			}, {
				EntitlementID: ngc.F("entitlementId"),
				PkID:          ngc.F("pkId"),
			}, {
				EntitlementID: ngc.F("entitlementId"),
				PkID:          ngc.F("pkId"),
			}}),
		}, {
			ProductName:    ngc.F("productName"),
			Type:           ngc.F(ngc.AdminOrgNewParamsProductEnablementsTypeNgcAdminEval),
			ExpirationDate: ngc.F("expirationDate"),
			PoDetails: ngc.F([]ngc.AdminOrgNewParamsProductEnablementsPoDetail{{
				EntitlementID: ngc.F("entitlementId"),
				PkID:          ngc.F("pkId"),
			}, {
				EntitlementID: ngc.F("entitlementId"),
				PkID:          ngc.F("pkId"),
			}, {
				EntitlementID: ngc.F("entitlementId"),
				PkID:          ngc.F("pkId"),
			}}),
		}}),
		ProductSubscriptions: ngc.F([]ngc.AdminOrgNewParamsProductSubscription{{
			ProductName:        ngc.F("productName"),
			ID:                 ngc.F("id"),
			EmsEntitlementType: ngc.F(ngc.AdminOrgNewParamsProductSubscriptionsEmsEntitlementTypeEmsEval),
			ExpirationDate:     ngc.F("expirationDate"),
			StartDate:          ngc.F("startDate"),
			Type:               ngc.F(ngc.AdminOrgNewParamsProductSubscriptionsTypeNgcAdminEval),
		}, {
			ProductName:        ngc.F("productName"),
			ID:                 ngc.F("id"),
			EmsEntitlementType: ngc.F(ngc.AdminOrgNewParamsProductSubscriptionsEmsEntitlementTypeEmsEval),
			ExpirationDate:     ngc.F("expirationDate"),
			StartDate:          ngc.F("startDate"),
			Type:               ngc.F(ngc.AdminOrgNewParamsProductSubscriptionsTypeNgcAdminEval),
		}, {
			ProductName:        ngc.F("productName"),
			ID:                 ngc.F("id"),
			EmsEntitlementType: ngc.F(ngc.AdminOrgNewParamsProductSubscriptionsEmsEntitlementTypeEmsEval),
			ExpirationDate:     ngc.F("expirationDate"),
			StartDate:          ngc.F("startDate"),
			Type:               ngc.F(ngc.AdminOrgNewParamsProductSubscriptionsTypeNgcAdminEval),
		}}),
		SalesforceAccountIndustry: ngc.F("salesforceAccountIndustry"),
		SendEmail:                 ngc.F(true),
		Type:                      ngc.F(ngc.AdminOrgNewParamsTypeUnknown),
	})
	if err != nil {
		var apierr *ngc.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
	defer resp.Body.Close()

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		var apierr *ngc.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
	if !bytes.Equal(b, []byte("abc")) {
		t.Fatalf("return value not %s: %s", "abc", b)
	}
}

func TestAdminOrgGet(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		w.Write([]byte("abc"))
	}))
	defer server.Close()
	baseURL := server.URL
	client := ngc.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAuthToken("My Auth Token"),
	)
	resp, err := client.Admin.Orgs.Get(context.TODO(), "org-name")
	if err != nil {
		var apierr *ngc.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
	defer resp.Body.Close()

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		var apierr *ngc.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
	if !bytes.Equal(b, []byte("abc")) {
		t.Fatalf("return value not %s: %s", "abc", b)
	}
}

func TestAdminOrgUpdateWithOptionalParams(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		w.Write([]byte("abc"))
	}))
	defer server.Close()
	baseURL := server.URL
	client := ngc.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAuthToken("My Auth Token"),
	)
	resp, err := client.Admin.Orgs.Update(
		context.TODO(),
		"org-name",
		ngc.AdminOrgUpdateParams{
			AlternateContact: ngc.F(ngc.AdminOrgUpdateParamsAlternateContact{
				Email:    ngc.F("xxxxxx"),
				FullName: ngc.F("fullName"),
			}),
			CompanyName: ngc.F("companyName"),
			Description: ngc.F("description"),
			DisplayName: ngc.F("x"),
			IdpID:       ngc.F("idpId"),
			InfinityManagerSettings: ngc.F(ngc.AdminOrgUpdateParamsInfinityManagerSettings{
				InfinityManagerEnabled:            ngc.F(true),
				InfinityManagerEnableTeamOverride: ngc.F(true),
			}),
			IsDatasetServiceEnabled:                 ngc.F(true),
			IsInternal:                              ngc.F(true),
			IsQuickStartEnabled:                     ngc.F(true),
			IsRegistrySseEnabled:                    ngc.F(true),
			IsSecretsManagerServiceEnabled:          ngc.F(true),
			IsSecureCredentialSharingServiceEnabled: ngc.F(true),
			IsSeparateInfluxDBUsed:                  ngc.F(true),
			OrgOwner: ngc.F(ngc.AdminOrgUpdateParamsOrgOwner{
				Email:         ngc.F("email"),
				FullName:      ngc.F("x"),
				LastLoginDate: ngc.F("lastLoginDate"),
			}),
			OrgOwners: ngc.F([]ngc.AdminOrgUpdateParamsOrgOwner{{
				Email:         ngc.F("email"),
				FullName:      ngc.F("x"),
				LastLoginDate: ngc.F("lastLoginDate"),
			}, {
				Email:         ngc.F("email"),
				FullName:      ngc.F("x"),
				LastLoginDate: ngc.F("lastLoginDate"),
			}, {
				Email:         ngc.F("email"),
				FullName:      ngc.F("x"),
				LastLoginDate: ngc.F("lastLoginDate"),
			}}),
			PecName:   ngc.F("pecName"),
			PecSfdcID: ngc.F("pecSfdcId"),
			ProductEnablements: ngc.F([]ngc.AdminOrgUpdateParamsProductEnablement{{
				ProductName:    ngc.F("productName"),
				Type:           ngc.F(ngc.AdminOrgUpdateParamsProductEnablementsTypeNgcAdminEval),
				ExpirationDate: ngc.F("expirationDate"),
				PoDetails: ngc.F([]ngc.AdminOrgUpdateParamsProductEnablementsPoDetail{{
					EntitlementID: ngc.F("entitlementId"),
					PkID:          ngc.F("pkId"),
				}, {
					EntitlementID: ngc.F("entitlementId"),
					PkID:          ngc.F("pkId"),
				}, {
					EntitlementID: ngc.F("entitlementId"),
					PkID:          ngc.F("pkId"),
				}}),
			}, {
				ProductName:    ngc.F("productName"),
				Type:           ngc.F(ngc.AdminOrgUpdateParamsProductEnablementsTypeNgcAdminEval),
				ExpirationDate: ngc.F("expirationDate"),
				PoDetails: ngc.F([]ngc.AdminOrgUpdateParamsProductEnablementsPoDetail{{
					EntitlementID: ngc.F("entitlementId"),
					PkID:          ngc.F("pkId"),
				}, {
					EntitlementID: ngc.F("entitlementId"),
					PkID:          ngc.F("pkId"),
				}, {
					EntitlementID: ngc.F("entitlementId"),
					PkID:          ngc.F("pkId"),
				}}),
			}, {
				ProductName:    ngc.F("productName"),
				Type:           ngc.F(ngc.AdminOrgUpdateParamsProductEnablementsTypeNgcAdminEval),
				ExpirationDate: ngc.F("expirationDate"),
				PoDetails: ngc.F([]ngc.AdminOrgUpdateParamsProductEnablementsPoDetail{{
					EntitlementID: ngc.F("entitlementId"),
					PkID:          ngc.F("pkId"),
				}, {
					EntitlementID: ngc.F("entitlementId"),
					PkID:          ngc.F("pkId"),
				}, {
					EntitlementID: ngc.F("entitlementId"),
					PkID:          ngc.F("pkId"),
				}}),
			}}),
			ProductSubscriptions: ngc.F([]ngc.AdminOrgUpdateParamsProductSubscription{{
				ProductName:        ngc.F("productName"),
				ID:                 ngc.F("id"),
				EmsEntitlementType: ngc.F(ngc.AdminOrgUpdateParamsProductSubscriptionsEmsEntitlementTypeEmsEval),
				ExpirationDate:     ngc.F("expirationDate"),
				StartDate:          ngc.F("startDate"),
				Type:               ngc.F(ngc.AdminOrgUpdateParamsProductSubscriptionsTypeNgcAdminEval),
			}, {
				ProductName:        ngc.F("productName"),
				ID:                 ngc.F("id"),
				EmsEntitlementType: ngc.F(ngc.AdminOrgUpdateParamsProductSubscriptionsEmsEntitlementTypeEmsEval),
				ExpirationDate:     ngc.F("expirationDate"),
				StartDate:          ngc.F("startDate"),
				Type:               ngc.F(ngc.AdminOrgUpdateParamsProductSubscriptionsTypeNgcAdminEval),
			}, {
				ProductName:        ngc.F("productName"),
				ID:                 ngc.F("id"),
				EmsEntitlementType: ngc.F(ngc.AdminOrgUpdateParamsProductSubscriptionsEmsEntitlementTypeEmsEval),
				ExpirationDate:     ngc.F("expirationDate"),
				StartDate:          ngc.F("startDate"),
				Type:               ngc.F(ngc.AdminOrgUpdateParamsProductSubscriptionsTypeNgcAdminEval),
			}}),
			RepoScanSettings: ngc.F(ngc.AdminOrgUpdateParamsRepoScanSettings{
				RepoScanAllowOverride:       ngc.F(true),
				RepoScanByDefault:           ngc.F(true),
				RepoScanEnabled:             ngc.F(true),
				RepoScanEnableNotifications: ngc.F(true),
				RepoScanEnableTeamOverride:  ngc.F(true),
				RepoScanShowResults:         ngc.F(true),
			}),
			Type: ngc.F(ngc.AdminOrgUpdateParamsTypeUnknown),
		},
	)
	if err != nil {
		var apierr *ngc.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
	defer resp.Body.Close()

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		var apierr *ngc.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
	if !bytes.Equal(b, []byte("abc")) {
		t.Fatalf("return value not %s: %s", "abc", b)
	}
}

func TestAdminOrgBackfillOrgsToKratos(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		w.Write([]byte("abc"))
	}))
	defer server.Close()
	baseURL := server.URL
	client := ngc.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAuthToken("My Auth Token"),
	)
	resp, err := client.Admin.Orgs.BackfillOrgsToKratos(context.TODO())
	if err != nil {
		var apierr *ngc.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
	defer resp.Body.Close()

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		var apierr *ngc.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
	if !bytes.Equal(b, []byte("abc")) {
		t.Fatalf("return value not %s: %s", "abc", b)
	}
}

func TestAdminOrgEnableWithOptionalParams(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		w.Write([]byte("abc"))
	}))
	defer server.Close()
	baseURL := server.URL
	client := ngc.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAuthToken("My Auth Token"),
	)
	resp, err := client.Admin.Orgs.Enable(
		context.TODO(),
		"org-name",
		ngc.AdminOrgEnableParams{
			CreateSubscription: ngc.F(true),
			ProductEnablement: ngc.F(ngc.AdminOrgEnableParamsProductEnablement{
				ProductName:    ngc.F("productName"),
				Type:           ngc.F(ngc.AdminOrgEnableParamsProductEnablementTypeNgcAdminEval),
				ExpirationDate: ngc.F("expirationDate"),
				PoDetails: ngc.F([]ngc.AdminOrgEnableParamsProductEnablementPoDetail{{
					EntitlementID: ngc.F("entitlementId"),
					PkID:          ngc.F("pkId"),
				}, {
					EntitlementID: ngc.F("entitlementId"),
					PkID:          ngc.F("pkId"),
				}, {
					EntitlementID: ngc.F("entitlementId"),
					PkID:          ngc.F("pkId"),
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
	defer resp.Body.Close()

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		var apierr *ngc.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
	if !bytes.Equal(b, []byte("abc")) {
		t.Fatalf("return value not %s: %s", "abc", b)
	}
}

func TestAdminOrgOrgOwnerBackfill(t *testing.T) {
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
	_, err := client.Admin.Orgs.OrgOwnerBackfill(context.TODO(), "org-name")
	if err != nil {
		var apierr *ngc.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
