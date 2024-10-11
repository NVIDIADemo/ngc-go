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

func TestOrgTeamNcaInvitationNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Orgs.Teams.NcaInvitations.New(
		context.TODO(),
		"org-name",
		"team-name",
		ngc.OrgTeamNcaInvitationNewParams{
			Email:                  ngc.F("xxxxxxx"),
			InvitationExpirationIn: ngc.F(int64(0)),
			InviteAs:               ngc.F(ngc.OrgTeamNcaInvitationNewParamsInviteAsAdmin),
			Message:                ngc.F("message"),
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
