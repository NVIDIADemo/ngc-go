// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package ngc

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"

	"github.com/NVIDIADemo/ngc-go/internal/apijson"
	"github.com/NVIDIADemo/ngc-go/internal/apiquery"
	"github.com/NVIDIADemo/ngc-go/internal/pagination"
	"github.com/NVIDIADemo/ngc-go/internal/param"
	"github.com/NVIDIADemo/ngc-go/internal/requestconfig"
	"github.com/NVIDIADemo/ngc-go/option"
	"github.com/NVIDIADemo/ngc-go/shared"
)

// OrgUserInvitationService contains methods and other services that help with
// interacting with the ngc API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewOrgUserInvitationService] method instead.
type OrgUserInvitationService struct {
	Options []option.RequestOption
}

// NewOrgUserInvitationService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewOrgUserInvitationService(opts ...option.RequestOption) (r *OrgUserInvitationService) {
	r = &OrgUserInvitationService{}
	r.Options = opts
	return
}

// List invitations in an org. (Org User Admin privileges required)
func (r *OrgUserInvitationService) List(ctx context.Context, orgName string, query OrgUserInvitationListParams, opts ...option.RequestOption) (res *pagination.PageNumberInvitations[shared.OrgUserInvitationListResponse], err error) {
	var raw *http.Response
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if orgName == "" {
		err = errors.New("missing required org-name parameter")
		return
	}
	path := fmt.Sprintf("v2/org/%s/users/invitations", orgName)
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodGet, path, query, &res, opts...)
	if err != nil {
		return nil, err
	}
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

// List invitations in an org. (Org User Admin privileges required)
func (r *OrgUserInvitationService) ListAutoPaging(ctx context.Context, orgName string, query OrgUserInvitationListParams, opts ...option.RequestOption) *pagination.PageNumberInvitationsAutoPager[shared.OrgUserInvitationListResponse] {
	return pagination.NewPageNumberInvitationsAutoPager(r.List(ctx, orgName, query, opts...))
}

// Delete a specific invitation in an org. (Org User Admin privileges required)
func (r *OrgUserInvitationService) Delete(ctx context.Context, orgName string, id string, opts ...option.RequestOption) (res *OrgUserInvitationDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	if orgName == "" {
		err = errors.New("missing required org-name parameter")
		return
	}
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("v2/org/%s/users/invitations/%s", orgName, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Resend email of a specific invitation in an org (Org User Admin privileges
// required).
func (r *OrgUserInvitationService) InviteResend(ctx context.Context, orgName string, id string, opts ...option.RequestOption) (res *shared.User, err error) {
	opts = append(r.Options[:], opts...)
	if orgName == "" {
		err = errors.New("missing required org-name parameter")
		return
	}
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("v2/org/%s/users/invitations/%s/resend-invitation-email", orgName, id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// User invitation to an NGC org or team
type OrgUserInvitationListResponse struct {
	// Unique invitation ID
	ID string `json:"id"`
	// Date on which the invitation was created. (ISO-8601 format)
	CreatedDate string `json:"createdDate"`
	// Email address of the user.
	Email string `json:"email"`
	// Flag indicating if the invitation has already been accepted by the user.
	IsProcessed bool `json:"isProcessed"`
	// user name
	Name string `json:"name"`
	// Org to which a user was invited.
	Org string `json:"org"`
	// List of roles that the user have.
	Roles []string `json:"roles"`
	// Team to which a user was invited.
	Team string `json:"team"`
	// Type of invitation. The invitation is either to an organization or to a team
	// within organization.
	Type OrgUserInvitationListResponseType `json:"type"`
	JSON orgUserInvitationListResponseJSON `json:"-"`
}

// orgUserInvitationListResponseJSON contains the JSON metadata for the struct
// [OrgUserInvitationListResponse]
type orgUserInvitationListResponseJSON struct {
	ID          apijson.Field
	CreatedDate apijson.Field
	Email       apijson.Field
	IsProcessed apijson.Field
	Name        apijson.Field
	Org         apijson.Field
	Roles       apijson.Field
	Team        apijson.Field
	Type        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *OrgUserInvitationListResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r orgUserInvitationListResponseJSON) RawJSON() string {
	return r.raw
}

// Type of invitation. The invitation is either to an organization or to a team
// within organization.
type OrgUserInvitationListResponseType string

const (
	OrgUserInvitationListResponseTypeOrganization OrgUserInvitationListResponseType = "ORGANIZATION"
	OrgUserInvitationListResponseTypeTeam         OrgUserInvitationListResponseType = "TEAM"
)

func (r OrgUserInvitationListResponseType) IsKnown() bool {
	switch r {
	case OrgUserInvitationListResponseTypeOrganization, OrgUserInvitationListResponseTypeTeam:
		return true
	}
	return false
}

type OrgUserInvitationDeleteResponse struct {
	RequestStatus OrgUserInvitationDeleteResponseRequestStatus `json:"requestStatus"`
	JSON          orgUserInvitationDeleteResponseJSON          `json:"-"`
}

// orgUserInvitationDeleteResponseJSON contains the JSON metadata for the struct
// [OrgUserInvitationDeleteResponse]
type orgUserInvitationDeleteResponseJSON struct {
	RequestStatus apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *OrgUserInvitationDeleteResponse) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r orgUserInvitationDeleteResponseJSON) RawJSON() string {
	return r.raw
}

type OrgUserInvitationDeleteResponseRequestStatus struct {
	RequestID string `json:"requestId"`
	ServerID  string `json:"serverId"`
	// Describes response status reported by the server.
	StatusCode        OrgUserInvitationDeleteResponseRequestStatusStatusCode `json:"statusCode"`
	StatusDescription string                                                 `json:"statusDescription"`
	JSON              orgUserInvitationDeleteResponseRequestStatusJSON       `json:"-"`
}

// orgUserInvitationDeleteResponseRequestStatusJSON contains the JSON metadata for
// the struct [OrgUserInvitationDeleteResponseRequestStatus]
type orgUserInvitationDeleteResponseRequestStatusJSON struct {
	RequestID         apijson.Field
	ServerID          apijson.Field
	StatusCode        apijson.Field
	StatusDescription apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *OrgUserInvitationDeleteResponseRequestStatus) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r orgUserInvitationDeleteResponseRequestStatusJSON) RawJSON() string {
	return r.raw
}

// Describes response status reported by the server.
type OrgUserInvitationDeleteResponseRequestStatusStatusCode string

const (
	OrgUserInvitationDeleteResponseRequestStatusStatusCodeUnknown                    OrgUserInvitationDeleteResponseRequestStatusStatusCode = "UNKNOWN"
	OrgUserInvitationDeleteResponseRequestStatusStatusCodeSuccess                    OrgUserInvitationDeleteResponseRequestStatusStatusCode = "SUCCESS"
	OrgUserInvitationDeleteResponseRequestStatusStatusCodeUnauthorized               OrgUserInvitationDeleteResponseRequestStatusStatusCode = "UNAUTHORIZED"
	OrgUserInvitationDeleteResponseRequestStatusStatusCodePaymentRequired            OrgUserInvitationDeleteResponseRequestStatusStatusCode = "PAYMENT_REQUIRED"
	OrgUserInvitationDeleteResponseRequestStatusStatusCodeForbidden                  OrgUserInvitationDeleteResponseRequestStatusStatusCode = "FORBIDDEN"
	OrgUserInvitationDeleteResponseRequestStatusStatusCodeTimeout                    OrgUserInvitationDeleteResponseRequestStatusStatusCode = "TIMEOUT"
	OrgUserInvitationDeleteResponseRequestStatusStatusCodeExists                     OrgUserInvitationDeleteResponseRequestStatusStatusCode = "EXISTS"
	OrgUserInvitationDeleteResponseRequestStatusStatusCodeNotFound                   OrgUserInvitationDeleteResponseRequestStatusStatusCode = "NOT_FOUND"
	OrgUserInvitationDeleteResponseRequestStatusStatusCodeInternalError              OrgUserInvitationDeleteResponseRequestStatusStatusCode = "INTERNAL_ERROR"
	OrgUserInvitationDeleteResponseRequestStatusStatusCodeInvalidRequest             OrgUserInvitationDeleteResponseRequestStatusStatusCode = "INVALID_REQUEST"
	OrgUserInvitationDeleteResponseRequestStatusStatusCodeInvalidRequestVersion      OrgUserInvitationDeleteResponseRequestStatusStatusCode = "INVALID_REQUEST_VERSION"
	OrgUserInvitationDeleteResponseRequestStatusStatusCodeInvalidRequestData         OrgUserInvitationDeleteResponseRequestStatusStatusCode = "INVALID_REQUEST_DATA"
	OrgUserInvitationDeleteResponseRequestStatusStatusCodeMethodNotAllowed           OrgUserInvitationDeleteResponseRequestStatusStatusCode = "METHOD_NOT_ALLOWED"
	OrgUserInvitationDeleteResponseRequestStatusStatusCodeConflict                   OrgUserInvitationDeleteResponseRequestStatusStatusCode = "CONFLICT"
	OrgUserInvitationDeleteResponseRequestStatusStatusCodeUnprocessableEntity        OrgUserInvitationDeleteResponseRequestStatusStatusCode = "UNPROCESSABLE_ENTITY"
	OrgUserInvitationDeleteResponseRequestStatusStatusCodeTooManyRequests            OrgUserInvitationDeleteResponseRequestStatusStatusCode = "TOO_MANY_REQUESTS"
	OrgUserInvitationDeleteResponseRequestStatusStatusCodeInsufficientStorage        OrgUserInvitationDeleteResponseRequestStatusStatusCode = "INSUFFICIENT_STORAGE"
	OrgUserInvitationDeleteResponseRequestStatusStatusCodeServiceUnavailable         OrgUserInvitationDeleteResponseRequestStatusStatusCode = "SERVICE_UNAVAILABLE"
	OrgUserInvitationDeleteResponseRequestStatusStatusCodePayloadTooLarge            OrgUserInvitationDeleteResponseRequestStatusStatusCode = "PAYLOAD_TOO_LARGE"
	OrgUserInvitationDeleteResponseRequestStatusStatusCodeNotAcceptable              OrgUserInvitationDeleteResponseRequestStatusStatusCode = "NOT_ACCEPTABLE"
	OrgUserInvitationDeleteResponseRequestStatusStatusCodeUnavailableForLegalReasons OrgUserInvitationDeleteResponseRequestStatusStatusCode = "UNAVAILABLE_FOR_LEGAL_REASONS"
	OrgUserInvitationDeleteResponseRequestStatusStatusCodeBadGateway                 OrgUserInvitationDeleteResponseRequestStatusStatusCode = "BAD_GATEWAY"
)

func (r OrgUserInvitationDeleteResponseRequestStatusStatusCode) IsKnown() bool {
	switch r {
	case OrgUserInvitationDeleteResponseRequestStatusStatusCodeUnknown, OrgUserInvitationDeleteResponseRequestStatusStatusCodeSuccess, OrgUserInvitationDeleteResponseRequestStatusStatusCodeUnauthorized, OrgUserInvitationDeleteResponseRequestStatusStatusCodePaymentRequired, OrgUserInvitationDeleteResponseRequestStatusStatusCodeForbidden, OrgUserInvitationDeleteResponseRequestStatusStatusCodeTimeout, OrgUserInvitationDeleteResponseRequestStatusStatusCodeExists, OrgUserInvitationDeleteResponseRequestStatusStatusCodeNotFound, OrgUserInvitationDeleteResponseRequestStatusStatusCodeInternalError, OrgUserInvitationDeleteResponseRequestStatusStatusCodeInvalidRequest, OrgUserInvitationDeleteResponseRequestStatusStatusCodeInvalidRequestVersion, OrgUserInvitationDeleteResponseRequestStatusStatusCodeInvalidRequestData, OrgUserInvitationDeleteResponseRequestStatusStatusCodeMethodNotAllowed, OrgUserInvitationDeleteResponseRequestStatusStatusCodeConflict, OrgUserInvitationDeleteResponseRequestStatusStatusCodeUnprocessableEntity, OrgUserInvitationDeleteResponseRequestStatusStatusCodeTooManyRequests, OrgUserInvitationDeleteResponseRequestStatusStatusCodeInsufficientStorage, OrgUserInvitationDeleteResponseRequestStatusStatusCodeServiceUnavailable, OrgUserInvitationDeleteResponseRequestStatusStatusCodePayloadTooLarge, OrgUserInvitationDeleteResponseRequestStatusStatusCodeNotAcceptable, OrgUserInvitationDeleteResponseRequestStatusStatusCodeUnavailableForLegalReasons, OrgUserInvitationDeleteResponseRequestStatusStatusCodeBadGateway:
		return true
	}
	return false
}

type OrgUserInvitationListParams struct {
	OrderBy param.Field[OrgUserInvitationListParamsOrderBy] `query:"orderBy"`
	// The page number of result
	PageNumber param.Field[int64] `query:"page-number"`
	// The page size of result
	PageSize param.Field[int64] `query:"page-size"`
	// User Search Parameters
	Q param.Field[OrgUserInvitationListParamsQ] `query:"q"`
}

// URLQuery serializes [OrgUserInvitationListParams]'s query parameters as
// `url.Values`.
func (r OrgUserInvitationListParams) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type OrgUserInvitationListParamsOrderBy string

const (
	OrgUserInvitationListParamsOrderByNameAsc  OrgUserInvitationListParamsOrderBy = "NAME_ASC"
	OrgUserInvitationListParamsOrderByNameDesc OrgUserInvitationListParamsOrderBy = "NAME_DESC"
)

func (r OrgUserInvitationListParamsOrderBy) IsKnown() bool {
	switch r {
	case OrgUserInvitationListParamsOrderByNameAsc, OrgUserInvitationListParamsOrderByNameDesc:
		return true
	}
	return false
}

// User Search Parameters
type OrgUserInvitationListParamsQ struct {
	Fields      param.Field[[]string]                              `query:"fields"`
	Filters     param.Field[[]OrgUserInvitationListParamsQFilter]  `query:"filters"`
	GroupBy     param.Field[string]                                `query:"groupBy"`
	OrderBy     param.Field[[]OrgUserInvitationListParamsQOrderBy] `query:"orderBy"`
	Page        param.Field[int64]                                 `query:"page"`
	PageSize    param.Field[int64]                                 `query:"pageSize"`
	Query       param.Field[string]                                `query:"query"`
	QueryFields param.Field[[]string]                              `query:"queryFields"`
	ScoredSize  param.Field[int64]                                 `query:"scoredSize"`
}

// URLQuery serializes [OrgUserInvitationListParamsQ]'s query parameters as
// `url.Values`.
func (r OrgUserInvitationListParamsQ) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type OrgUserInvitationListParamsQFilter struct {
	Field param.Field[string] `query:"field"`
	Value param.Field[string] `query:"value"`
}

// URLQuery serializes [OrgUserInvitationListParamsQFilter]'s query parameters as
// `url.Values`.
func (r OrgUserInvitationListParamsQFilter) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type OrgUserInvitationListParamsQOrderBy struct {
	Field param.Field[string]                                   `query:"field"`
	Value param.Field[OrgUserInvitationListParamsQOrderByValue] `query:"value"`
}

// URLQuery serializes [OrgUserInvitationListParamsQOrderBy]'s query parameters as
// `url.Values`.
func (r OrgUserInvitationListParamsQOrderBy) URLQuery() (v url.Values) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type OrgUserInvitationListParamsQOrderByValue string

const (
	OrgUserInvitationListParamsQOrderByValueAsc  OrgUserInvitationListParamsQOrderByValue = "ASC"
	OrgUserInvitationListParamsQOrderByValueDesc OrgUserInvitationListParamsQOrderByValue = "DESC"
)

func (r OrgUserInvitationListParamsQOrderByValue) IsKnown() bool {
	switch r {
	case OrgUserInvitationListParamsQOrderByValueAsc, OrgUserInvitationListParamsQOrderByValueDesc:
		return true
	}
	return false
}
