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

func TestMeAPIKeyGetWithOptionalParams(t *testing.T) {
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
	_, err := client.Me.APIKey.Get(context.TODO(), ngc.MeAPIKeyGetParams{
		OrgName: ngc.F("org-name"),
	})
	if err != nil {
		var apierr *ngc.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestMeAPIKeyUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Me.APIKey.Update(context.TODO(), ngc.MeAPIKeyUpdateParams{
		HasEmailOptIn:                        ngc.F(true),
		HasSignedAIFoundryPartnershipsEula:   ngc.F(true),
		HasSignedBaseCommandEula:             ngc.F(true),
		HasSignedBaseCommandManagerEula:      ngc.F(true),
		HasSignedBioNeMoEula:                 ngc.F(true),
		HasSignedContainerPublishingEula:     ngc.F(true),
		HasSignedCuOptEula:                   ngc.F(true),
		HasSignedEarth2Eula:                  ngc.F(true),
		HasSignedEgxEula:                     ngc.F(true),
		HasSignedEula:                        ngc.F(true),
		HasSignedFleetCommandEula:            ngc.F(true),
		HasSignedLlmEula:                     ngc.F(true),
		HasSignedNvaieeula:                   ngc.F(true),
		HasSignedNvidiaEula:                  ngc.F(true),
		HasSignedNvqceula:                    ngc.F(true),
		HasSignedOmniverseEula:               ngc.F(true),
		HasSignedPrivacyPolicy:               ngc.F(true),
		HasSignedThirdPartyRegistryShareEula: ngc.F(true),
		Name:                                 ngc.F("x"),
		UserMetadata: ngc.F(ngc.MeAPIKeyUpdateParamsUserMetadata{
			Company:    ngc.F("company"),
			CompanyURL: ngc.F("companyUrl"),
			Country:    ngc.F("country"),
			FirstName:  ngc.F("firstName"),
			Industry:   ngc.F("industry"),
			Interest:   ngc.F([]string{"string", "string", "string"}),
			LastName:   ngc.F("lastName"),
			Role:       ngc.F("role"),
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

func TestMeAPIKeyNewAPIKey(t *testing.T) {
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
	_, err := client.Me.APIKey.NewAPIKey(context.TODO())
	if err != nil {
		var apierr *ngc.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
