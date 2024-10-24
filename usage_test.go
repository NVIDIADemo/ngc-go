// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package ngc_test

import (
	"context"
	"os"
	"testing"

	"github.com/NVIDIADemo/ngc-go"
	"github.com/NVIDIADemo/ngc-go/internal/testutil"
	"github.com/NVIDIADemo/ngc-go/option"
)

func TestUsage(t *testing.T) {
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
	orgResponse, err := client.Orgs.New(context.TODO(), ngc.OrgNewParams{})
	if err != nil {
		t.Error(err)
	}
	t.Logf("%+v\n", orgResponse.Organizations)
}
