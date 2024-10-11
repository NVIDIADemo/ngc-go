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

func TestSuperAdminUserOrgTeamUserNewWithOptionalParams(t *testing.T) {
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
	_, err := client.SuperAdminUser.Orgs.Teams.Users.New(
		context.TODO(),
		"org-name",
		"team-name",
		ngc.SuperAdminUserOrgTeamUserNewParams{
			Email:                    ngc.F("xxxxxx"),
			IdpID:                    ngc.F("idp-id"),
			SendEmail:                ngc.F(true),
			EmailOptIn:               ngc.F(true),
			EulaAccepted:             ngc.F(true),
			Name:                     ngc.F("x"),
			RoleType:                 ngc.F("roleType"),
			RoleTypes:                ngc.F([]string{"string", "string", "string"}),
			SalesforceContactJobRole: ngc.F("salesforceContactJobRole"),
			UserMetadata: ngc.F(ngc.SuperAdminUserOrgTeamUserNewParamsUserMetadata{
				Company:    ngc.F("company"),
				CompanyURL: ngc.F("companyUrl"),
				Country:    ngc.F("country"),
				FirstName:  ngc.F("firstName"),
				Industry:   ngc.F("industry"),
				Interest:   ngc.F([]string{"string", "string", "string"}),
				LastName:   ngc.F("lastName"),
				Role:       ngc.F("role"),
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

func TestSuperAdminUserOrgTeamUserAddWithOptionalParams(t *testing.T) {
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
	_, err := client.SuperAdminUser.Orgs.Teams.Users.Add(
		context.TODO(),
		"org-name",
		"team-name",
		"id",
		ngc.SuperAdminUserOrgTeamUserAddParams{
			UserRoleDefaultsToRegistryRead: ngc.F("user role, defaults to REGISTRY_READ"),
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
