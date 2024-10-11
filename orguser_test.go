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

func TestOrgUserNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Orgs.Users.New(
		context.TODO(),
		"org-name",
		ngc.OrgUserNewParams{
			Email:                    ngc.F("xxxxxx"),
			IdpID:                    ngc.F("idp-id"),
			SendEmail:                ngc.F(true),
			EmailOptIn:               ngc.F(true),
			EulaAccepted:             ngc.F(true),
			Name:                     ngc.F("x"),
			RoleType:                 ngc.F("roleType"),
			RoleTypes:                ngc.F([]string{"string", "string", "string"}),
			SalesforceContactJobRole: ngc.F("salesforceContactJobRole"),
			UserMetadata: ngc.F(ngc.OrgUserNewParamsUserMetadata{
				Company:    ngc.F("company"),
				CompanyURL: ngc.F("companyUrl"),
				Country:    ngc.F("country"),
				FirstName:  ngc.F("firstName"),
				Industry:   ngc.F("industry"),
				Interest:   ngc.F([]string{"string", "string", "string"}),
				LastName:   ngc.F("lastName"),
				Role:       ngc.F("role"),
			}),
			Ncid:      ngc.F("ncid"),
			VisitorID: ngc.F("VisitorID"),
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

func TestOrgUserListWithOptionalParams(t *testing.T) {
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
	_, err := client.Orgs.Users.List(
		context.TODO(),
		"org-name",
		ngc.OrgUserListParams{
			ExcludeFromTeam: ngc.F("exclude-from-team"),
			PageNumber:      ngc.F(int64(0)),
			PageSize:        ngc.F(int64(0)),
			Q: ngc.F(ngc.OrgUserListParamsQ{
				Fields: ngc.F([]string{"string", "string", "string"}),
				Filters: ngc.F([]ngc.OrgUserListParamsQFilter{{
					Field: ngc.F("field"),
					Value: ngc.F("value"),
				}, {
					Field: ngc.F("field"),
					Value: ngc.F("value"),
				}, {
					Field: ngc.F("field"),
					Value: ngc.F("value"),
				}}),
				GroupBy: ngc.F("groupBy"),
				OrderBy: ngc.F([]ngc.OrgUserListParamsQOrderBy{{
					Field: ngc.F("field"),
					Value: ngc.F(ngc.OrgUserListParamsQOrderByValueAsc),
				}, {
					Field: ngc.F("field"),
					Value: ngc.F(ngc.OrgUserListParamsQOrderByValueAsc),
				}, {
					Field: ngc.F("field"),
					Value: ngc.F(ngc.OrgUserListParamsQOrderByValueAsc),
				}}),
				Page:        ngc.F(int64(0)),
				PageSize:    ngc.F(int64(0)),
				Query:       ngc.F("query"),
				QueryFields: ngc.F([]string{"string", "string", "string"}),
				ScoredSize:  ngc.F(int64(0)),
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

func TestOrgUserDeleteWithOptionalParams(t *testing.T) {
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
	_, err := client.Orgs.Users.Delete(
		context.TODO(),
		"org-name",
		"id",
		ngc.OrgUserDeleteParams{
			Anonymize: ngc.F(true),
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

func TestOrgUserAddRoleWithOptionalParams(t *testing.T) {
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
	_, err := client.Orgs.Users.AddRole(
		context.TODO(),
		"org-name",
		"user-email-or-id",
		ngc.OrgUserAddRoleParams{
			Roles:     ngc.F([]string{"string", "string", "string"}),
			Ncid:      ngc.F("ncid"),
			VisitorID: ngc.F("VisitorID"),
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

func TestOrgUserRemoveRoleWithOptionalParams(t *testing.T) {
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
	_, err := client.Orgs.Users.RemoveRole(
		context.TODO(),
		"org-name",
		"user-email-or-id",
		ngc.OrgUserRemoveRoleParams{
			Roles: ngc.F([]string{"string", "string", "string"}),
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

func TestOrgUserUpdateRoleWithOptionalParams(t *testing.T) {
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
	_, err := client.Orgs.Users.UpdateRole(
		context.TODO(),
		"org-name",
		"id",
		ngc.OrgUserUpdateRoleParams{
			Roles: ngc.F([]string{"string", "string", "string"}),
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
