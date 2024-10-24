// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package shared

import (
	"github.com/NVIDIADemo/ngc-go/internal/apijson"
)

// This API is invoked by monitoring tools, other services and infrastructure to
// retrieve health status the targeted service, this is unprotected method
type Health struct {
	// object that describes health of the service
	Health        HealthHealth        `json:"health"`
	RequestStatus HealthRequestStatus `json:"requestStatus"`
	JSON          healthJSON          `json:"-"`
}

// healthJSON contains the JSON metadata for the struct [Health]
type healthJSON struct {
	Health        apijson.Field
	RequestStatus apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *Health) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r healthJSON) RawJSON() string {
	return r.raw
}

// object that describes health of the service
type HealthHealth struct {
	// Enum that describes health of the service
	HealthCode HealthHealthHealthCode `json:"healthCode"`
	// Human readable description
	HealthCodeDescription string                 `json:"healthCodeDescription"`
	MetaData              []HealthHealthMetaData `json:"metaData"`
	JSON                  healthHealthJSON       `json:"-"`
}

// healthHealthJSON contains the JSON metadata for the struct [HealthHealth]
type healthHealthJSON struct {
	HealthCode            apijson.Field
	HealthCodeDescription apijson.Field
	MetaData              apijson.Field
	raw                   string
	ExtraFields           map[string]apijson.Field
}

func (r *HealthHealth) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r healthHealthJSON) RawJSON() string {
	return r.raw
}

// Enum that describes health of the service
type HealthHealthHealthCode string

const (
	HealthHealthHealthCodeUnknown          HealthHealthHealthCode = "UNKNOWN"
	HealthHealthHealthCodeOk               HealthHealthHealthCode = "OK"
	HealthHealthHealthCodeUnderMaintenance HealthHealthHealthCode = "UNDER_MAINTENANCE"
	HealthHealthHealthCodeWarning          HealthHealthHealthCode = "WARNING"
	HealthHealthHealthCodeFailed           HealthHealthHealthCode = "FAILED"
)

func (r HealthHealthHealthCode) IsKnown() bool {
	switch r {
	case HealthHealthHealthCodeUnknown, HealthHealthHealthCodeOk, HealthHealthHealthCodeUnderMaintenance, HealthHealthHealthCodeWarning, HealthHealthHealthCodeFailed:
		return true
	}
	return false
}

type HealthHealthMetaData struct {
	Key   string                   `json:"key"`
	Value string                   `json:"value"`
	JSON  healthHealthMetaDataJSON `json:"-"`
}

// healthHealthMetaDataJSON contains the JSON metadata for the struct
// [HealthHealthMetaData]
type healthHealthMetaDataJSON struct {
	Key         apijson.Field
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *HealthHealthMetaData) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r healthHealthMetaDataJSON) RawJSON() string {
	return r.raw
}

type HealthRequestStatus struct {
	RequestID string `json:"requestId"`
	ServerID  string `json:"serverId"`
	// Describes response status reported by the server.
	StatusCode        HealthRequestStatusStatusCode `json:"statusCode"`
	StatusDescription string                        `json:"statusDescription"`
	JSON              healthRequestStatusJSON       `json:"-"`
}

// healthRequestStatusJSON contains the JSON metadata for the struct
// [HealthRequestStatus]
type healthRequestStatusJSON struct {
	RequestID         apijson.Field
	ServerID          apijson.Field
	StatusCode        apijson.Field
	StatusDescription apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *HealthRequestStatus) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r healthRequestStatusJSON) RawJSON() string {
	return r.raw
}

// Describes response status reported by the server.
type HealthRequestStatusStatusCode string

const (
	HealthRequestStatusStatusCodeUnknown                    HealthRequestStatusStatusCode = "UNKNOWN"
	HealthRequestStatusStatusCodeSuccess                    HealthRequestStatusStatusCode = "SUCCESS"
	HealthRequestStatusStatusCodeUnauthorized               HealthRequestStatusStatusCode = "UNAUTHORIZED"
	HealthRequestStatusStatusCodePaymentRequired            HealthRequestStatusStatusCode = "PAYMENT_REQUIRED"
	HealthRequestStatusStatusCodeForbidden                  HealthRequestStatusStatusCode = "FORBIDDEN"
	HealthRequestStatusStatusCodeTimeout                    HealthRequestStatusStatusCode = "TIMEOUT"
	HealthRequestStatusStatusCodeExists                     HealthRequestStatusStatusCode = "EXISTS"
	HealthRequestStatusStatusCodeNotFound                   HealthRequestStatusStatusCode = "NOT_FOUND"
	HealthRequestStatusStatusCodeInternalError              HealthRequestStatusStatusCode = "INTERNAL_ERROR"
	HealthRequestStatusStatusCodeInvalidRequest             HealthRequestStatusStatusCode = "INVALID_REQUEST"
	HealthRequestStatusStatusCodeInvalidRequestVersion      HealthRequestStatusStatusCode = "INVALID_REQUEST_VERSION"
	HealthRequestStatusStatusCodeInvalidRequestData         HealthRequestStatusStatusCode = "INVALID_REQUEST_DATA"
	HealthRequestStatusStatusCodeMethodNotAllowed           HealthRequestStatusStatusCode = "METHOD_NOT_ALLOWED"
	HealthRequestStatusStatusCodeConflict                   HealthRequestStatusStatusCode = "CONFLICT"
	HealthRequestStatusStatusCodeUnprocessableEntity        HealthRequestStatusStatusCode = "UNPROCESSABLE_ENTITY"
	HealthRequestStatusStatusCodeTooManyRequests            HealthRequestStatusStatusCode = "TOO_MANY_REQUESTS"
	HealthRequestStatusStatusCodeInsufficientStorage        HealthRequestStatusStatusCode = "INSUFFICIENT_STORAGE"
	HealthRequestStatusStatusCodeServiceUnavailable         HealthRequestStatusStatusCode = "SERVICE_UNAVAILABLE"
	HealthRequestStatusStatusCodePayloadTooLarge            HealthRequestStatusStatusCode = "PAYLOAD_TOO_LARGE"
	HealthRequestStatusStatusCodeNotAcceptable              HealthRequestStatusStatusCode = "NOT_ACCEPTABLE"
	HealthRequestStatusStatusCodeUnavailableForLegalReasons HealthRequestStatusStatusCode = "UNAVAILABLE_FOR_LEGAL_REASONS"
	HealthRequestStatusStatusCodeBadGateway                 HealthRequestStatusStatusCode = "BAD_GATEWAY"
)

func (r HealthRequestStatusStatusCode) IsKnown() bool {
	switch r {
	case HealthRequestStatusStatusCodeUnknown, HealthRequestStatusStatusCodeSuccess, HealthRequestStatusStatusCodeUnauthorized, HealthRequestStatusStatusCodePaymentRequired, HealthRequestStatusStatusCodeForbidden, HealthRequestStatusStatusCodeTimeout, HealthRequestStatusStatusCodeExists, HealthRequestStatusStatusCodeNotFound, HealthRequestStatusStatusCodeInternalError, HealthRequestStatusStatusCodeInvalidRequest, HealthRequestStatusStatusCodeInvalidRequestVersion, HealthRequestStatusStatusCodeInvalidRequestData, HealthRequestStatusStatusCodeMethodNotAllowed, HealthRequestStatusStatusCodeConflict, HealthRequestStatusStatusCodeUnprocessableEntity, HealthRequestStatusStatusCodeTooManyRequests, HealthRequestStatusStatusCodeInsufficientStorage, HealthRequestStatusStatusCodeServiceUnavailable, HealthRequestStatusStatusCodePayloadTooLarge, HealthRequestStatusStatusCodeNotAcceptable, HealthRequestStatusStatusCodeUnavailableForLegalReasons, HealthRequestStatusStatusCodeBadGateway:
		return true
	}
	return false
}

// response containing a list of all metering queries results
type MeteringResultList struct {
	Measurements  []MeteringResultListMeasurement `json:"measurements"`
	RequestStatus MeteringResultListRequestStatus `json:"requestStatus"`
	JSON          meteringResultListJSON          `json:"-"`
}

// meteringResultListJSON contains the JSON metadata for the struct
// [MeteringResultList]
type meteringResultListJSON struct {
	Measurements  apijson.Field
	RequestStatus apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *MeteringResultList) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r meteringResultListJSON) RawJSON() string {
	return r.raw
}

// result of a single measurement query
type MeteringResultListMeasurement struct {
	// array of series within a measurement
	Series []MeteringResultListMeasurementsSery `json:"series"`
	JSON   meteringResultListMeasurementJSON    `json:"-"`
}

// meteringResultListMeasurementJSON contains the JSON metadata for the struct
// [MeteringResultListMeasurement]
type meteringResultListMeasurementJSON struct {
	Series      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MeteringResultListMeasurement) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r meteringResultListMeasurementJSON) RawJSON() string {
	return r.raw
}

// object for a single series in the measurement
type MeteringResultListMeasurementsSery struct {
	// list of columns, in order, for the series.
	Columns []string `json:"columns"`
	// name for the measurement
	Name string `json:"name"`
	// list of tags identifying the series.
	Tags []MeteringResultListMeasurementsSeriesTag `json:"tags"`
	// array of values, in the same order as the columns, for the series.
	Values []MeteringResultListMeasurementsSeriesValue `json:"values"`
	JSON   meteringResultListMeasurementsSeryJSON      `json:"-"`
}

// meteringResultListMeasurementsSeryJSON contains the JSON metadata for the struct
// [MeteringResultListMeasurementsSery]
type meteringResultListMeasurementsSeryJSON struct {
	Columns     apijson.Field
	Name        apijson.Field
	Tags        apijson.Field
	Values      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MeteringResultListMeasurementsSery) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r meteringResultListMeasurementsSeryJSON) RawJSON() string {
	return r.raw
}

// object for measurement tags which identifies a measuurement series
type MeteringResultListMeasurementsSeriesTag struct {
	// key for the tag, ie)host, job_id, gpu_id
	TagKey string `json:"tagKey"`
	// value for the tag, ie)host=foo, job_id=bar, gpu_id=racecar
	TagValue string                                      `json:"tagValue"`
	JSON     meteringResultListMeasurementsSeriesTagJSON `json:"-"`
}

// meteringResultListMeasurementsSeriesTagJSON contains the JSON metadata for the
// struct [MeteringResultListMeasurementsSeriesTag]
type meteringResultListMeasurementsSeriesTagJSON struct {
	TagKey      apijson.Field
	TagValue    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MeteringResultListMeasurementsSeriesTag) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r meteringResultListMeasurementsSeriesTagJSON) RawJSON() string {
	return r.raw
}

// object for the measurement values
type MeteringResultListMeasurementsSeriesValue struct {
	// list of values, in the same order as columns
	Value []string                                      `json:"value"`
	JSON  meteringResultListMeasurementsSeriesValueJSON `json:"-"`
}

// meteringResultListMeasurementsSeriesValueJSON contains the JSON metadata for the
// struct [MeteringResultListMeasurementsSeriesValue]
type meteringResultListMeasurementsSeriesValueJSON struct {
	Value       apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *MeteringResultListMeasurementsSeriesValue) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r meteringResultListMeasurementsSeriesValueJSON) RawJSON() string {
	return r.raw
}

type MeteringResultListRequestStatus struct {
	RequestID string `json:"requestId"`
	ServerID  string `json:"serverId"`
	// Describes response status reported by the server.
	StatusCode        MeteringResultListRequestStatusStatusCode `json:"statusCode"`
	StatusDescription string                                    `json:"statusDescription"`
	JSON              meteringResultListRequestStatusJSON       `json:"-"`
}

// meteringResultListRequestStatusJSON contains the JSON metadata for the struct
// [MeteringResultListRequestStatus]
type meteringResultListRequestStatusJSON struct {
	RequestID         apijson.Field
	ServerID          apijson.Field
	StatusCode        apijson.Field
	StatusDescription apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *MeteringResultListRequestStatus) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r meteringResultListRequestStatusJSON) RawJSON() string {
	return r.raw
}

// Describes response status reported by the server.
type MeteringResultListRequestStatusStatusCode string

const (
	MeteringResultListRequestStatusStatusCodeUnknown                    MeteringResultListRequestStatusStatusCode = "UNKNOWN"
	MeteringResultListRequestStatusStatusCodeSuccess                    MeteringResultListRequestStatusStatusCode = "SUCCESS"
	MeteringResultListRequestStatusStatusCodeUnauthorized               MeteringResultListRequestStatusStatusCode = "UNAUTHORIZED"
	MeteringResultListRequestStatusStatusCodePaymentRequired            MeteringResultListRequestStatusStatusCode = "PAYMENT_REQUIRED"
	MeteringResultListRequestStatusStatusCodeForbidden                  MeteringResultListRequestStatusStatusCode = "FORBIDDEN"
	MeteringResultListRequestStatusStatusCodeTimeout                    MeteringResultListRequestStatusStatusCode = "TIMEOUT"
	MeteringResultListRequestStatusStatusCodeExists                     MeteringResultListRequestStatusStatusCode = "EXISTS"
	MeteringResultListRequestStatusStatusCodeNotFound                   MeteringResultListRequestStatusStatusCode = "NOT_FOUND"
	MeteringResultListRequestStatusStatusCodeInternalError              MeteringResultListRequestStatusStatusCode = "INTERNAL_ERROR"
	MeteringResultListRequestStatusStatusCodeInvalidRequest             MeteringResultListRequestStatusStatusCode = "INVALID_REQUEST"
	MeteringResultListRequestStatusStatusCodeInvalidRequestVersion      MeteringResultListRequestStatusStatusCode = "INVALID_REQUEST_VERSION"
	MeteringResultListRequestStatusStatusCodeInvalidRequestData         MeteringResultListRequestStatusStatusCode = "INVALID_REQUEST_DATA"
	MeteringResultListRequestStatusStatusCodeMethodNotAllowed           MeteringResultListRequestStatusStatusCode = "METHOD_NOT_ALLOWED"
	MeteringResultListRequestStatusStatusCodeConflict                   MeteringResultListRequestStatusStatusCode = "CONFLICT"
	MeteringResultListRequestStatusStatusCodeUnprocessableEntity        MeteringResultListRequestStatusStatusCode = "UNPROCESSABLE_ENTITY"
	MeteringResultListRequestStatusStatusCodeTooManyRequests            MeteringResultListRequestStatusStatusCode = "TOO_MANY_REQUESTS"
	MeteringResultListRequestStatusStatusCodeInsufficientStorage        MeteringResultListRequestStatusStatusCode = "INSUFFICIENT_STORAGE"
	MeteringResultListRequestStatusStatusCodeServiceUnavailable         MeteringResultListRequestStatusStatusCode = "SERVICE_UNAVAILABLE"
	MeteringResultListRequestStatusStatusCodePayloadTooLarge            MeteringResultListRequestStatusStatusCode = "PAYLOAD_TOO_LARGE"
	MeteringResultListRequestStatusStatusCodeNotAcceptable              MeteringResultListRequestStatusStatusCode = "NOT_ACCEPTABLE"
	MeteringResultListRequestStatusStatusCodeUnavailableForLegalReasons MeteringResultListRequestStatusStatusCode = "UNAVAILABLE_FOR_LEGAL_REASONS"
	MeteringResultListRequestStatusStatusCodeBadGateway                 MeteringResultListRequestStatusStatusCode = "BAD_GATEWAY"
)

func (r MeteringResultListRequestStatusStatusCode) IsKnown() bool {
	switch r {
	case MeteringResultListRequestStatusStatusCodeUnknown, MeteringResultListRequestStatusStatusCodeSuccess, MeteringResultListRequestStatusStatusCodeUnauthorized, MeteringResultListRequestStatusStatusCodePaymentRequired, MeteringResultListRequestStatusStatusCodeForbidden, MeteringResultListRequestStatusStatusCodeTimeout, MeteringResultListRequestStatusStatusCodeExists, MeteringResultListRequestStatusStatusCodeNotFound, MeteringResultListRequestStatusStatusCodeInternalError, MeteringResultListRequestStatusStatusCodeInvalidRequest, MeteringResultListRequestStatusStatusCodeInvalidRequestVersion, MeteringResultListRequestStatusStatusCodeInvalidRequestData, MeteringResultListRequestStatusStatusCodeMethodNotAllowed, MeteringResultListRequestStatusStatusCodeConflict, MeteringResultListRequestStatusStatusCodeUnprocessableEntity, MeteringResultListRequestStatusStatusCodeTooManyRequests, MeteringResultListRequestStatusStatusCodeInsufficientStorage, MeteringResultListRequestStatusStatusCodeServiceUnavailable, MeteringResultListRequestStatusStatusCodePayloadTooLarge, MeteringResultListRequestStatusStatusCodeNotAcceptable, MeteringResultListRequestStatusStatusCodeUnavailableForLegalReasons, MeteringResultListRequestStatusStatusCodeBadGateway:
		return true
	}
	return false
}

// Information about the team
type Team struct {
	// unique Id of this team.
	ID int64 `json:"id"`
	// description of the team
	Description string `json:"description"`
	// Infinity manager setting definition
	InfinityManagerSettings TeamInfinityManagerSettings `json:"infinityManagerSettings"`
	// indicates if the team is deleted or not
	IsDeleted bool `json:"isDeleted"`
	// team name
	Name string `json:"name"`
	// Repo scan setting definition
	RepoScanSettings TeamRepoScanSettings `json:"repoScanSettings"`
	JSON             teamJSON             `json:"-"`
}

// teamJSON contains the JSON metadata for the struct [Team]
type teamJSON struct {
	ID                      apijson.Field
	Description             apijson.Field
	InfinityManagerSettings apijson.Field
	IsDeleted               apijson.Field
	Name                    apijson.Field
	RepoScanSettings        apijson.Field
	raw                     string
	ExtraFields             map[string]apijson.Field
}

func (r *Team) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r teamJSON) RawJSON() string {
	return r.raw
}

// Infinity manager setting definition
type TeamInfinityManagerSettings struct {
	// Enable the infinity manager or not. Used both in org and team level object
	InfinityManagerEnabled bool `json:"infinityManagerEnabled"`
	// Allow override settings at team level. Only used in org level object
	InfinityManagerEnableTeamOverride bool                            `json:"infinityManagerEnableTeamOverride"`
	JSON                              teamInfinityManagerSettingsJSON `json:"-"`
}

// teamInfinityManagerSettingsJSON contains the JSON metadata for the struct
// [TeamInfinityManagerSettings]
type teamInfinityManagerSettingsJSON struct {
	InfinityManagerEnabled            apijson.Field
	InfinityManagerEnableTeamOverride apijson.Field
	raw                               string
	ExtraFields                       map[string]apijson.Field
}

func (r *TeamInfinityManagerSettings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r teamInfinityManagerSettingsJSON) RawJSON() string {
	return r.raw
}

// Repo scan setting definition
type TeamRepoScanSettings struct {
	// Allow org admin to override the org level repo scan settings
	RepoScanAllowOverride bool `json:"repoScanAllowOverride"`
	// Allow repository scanning by default
	RepoScanByDefault bool `json:"repoScanByDefault"`
	// Enable the repository scan or not. Only used in org level object
	RepoScanEnabled bool `json:"repoScanEnabled"`
	// Sends notification to end user after scanning is done
	RepoScanEnableNotifications bool `json:"repoScanEnableNotifications"`
	// Allow override settings at team level. Only used in org level object
	RepoScanEnableTeamOverride bool `json:"repoScanEnableTeamOverride"`
	// Allow showing scan results to CLI or UI
	RepoScanShowResults bool                     `json:"repoScanShowResults"`
	JSON                teamRepoScanSettingsJSON `json:"-"`
}

// teamRepoScanSettingsJSON contains the JSON metadata for the struct
// [TeamRepoScanSettings]
type teamRepoScanSettingsJSON struct {
	RepoScanAllowOverride       apijson.Field
	RepoScanByDefault           apijson.Field
	RepoScanEnabled             apijson.Field
	RepoScanEnableNotifications apijson.Field
	RepoScanEnableTeamOverride  apijson.Field
	RepoScanShowResults         apijson.Field
	raw                         string
	ExtraFields                 map[string]apijson.Field
}

func (r *TeamRepoScanSettings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r teamRepoScanSettingsJSON) RawJSON() string {
	return r.raw
}

// listing of all teams
type TeamList struct {
	// object that describes the pagination information
	PaginationInfo TeamListPaginationInfo `json:"paginationInfo"`
	RequestStatus  TeamListRequestStatus  `json:"requestStatus"`
	Teams          []Team                 `json:"teams"`
	JSON           teamListJSON           `json:"-"`
}

// teamListJSON contains the JSON metadata for the struct [TeamList]
type teamListJSON struct {
	PaginationInfo apijson.Field
	RequestStatus  apijson.Field
	Teams          apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *TeamList) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r teamListJSON) RawJSON() string {
	return r.raw
}

// object that describes the pagination information
type TeamListPaginationInfo struct {
	// Page index of results
	Index int64 `json:"index"`
	// Serialized pointer to the next results page. Should be used for fetching next
	// page. Can be empty
	NextPage string `json:"nextPage"`
	// Number of results in page
	Size int64 `json:"size"`
	// Total number of pages available
	TotalPages int64 `json:"totalPages"`
	// Total number of results available
	TotalResults int64                      `json:"totalResults"`
	JSON         teamListPaginationInfoJSON `json:"-"`
}

// teamListPaginationInfoJSON contains the JSON metadata for the struct
// [TeamListPaginationInfo]
type teamListPaginationInfoJSON struct {
	Index        apijson.Field
	NextPage     apijson.Field
	Size         apijson.Field
	TotalPages   apijson.Field
	TotalResults apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *TeamListPaginationInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r teamListPaginationInfoJSON) RawJSON() string {
	return r.raw
}

type TeamListRequestStatus struct {
	RequestID string `json:"requestId"`
	ServerID  string `json:"serverId"`
	// Describes response status reported by the server.
	StatusCode        TeamListRequestStatusStatusCode `json:"statusCode"`
	StatusDescription string                          `json:"statusDescription"`
	JSON              teamListRequestStatusJSON       `json:"-"`
}

// teamListRequestStatusJSON contains the JSON metadata for the struct
// [TeamListRequestStatus]
type teamListRequestStatusJSON struct {
	RequestID         apijson.Field
	ServerID          apijson.Field
	StatusCode        apijson.Field
	StatusDescription apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *TeamListRequestStatus) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r teamListRequestStatusJSON) RawJSON() string {
	return r.raw
}

// Describes response status reported by the server.
type TeamListRequestStatusStatusCode string

const (
	TeamListRequestStatusStatusCodeUnknown                    TeamListRequestStatusStatusCode = "UNKNOWN"
	TeamListRequestStatusStatusCodeSuccess                    TeamListRequestStatusStatusCode = "SUCCESS"
	TeamListRequestStatusStatusCodeUnauthorized               TeamListRequestStatusStatusCode = "UNAUTHORIZED"
	TeamListRequestStatusStatusCodePaymentRequired            TeamListRequestStatusStatusCode = "PAYMENT_REQUIRED"
	TeamListRequestStatusStatusCodeForbidden                  TeamListRequestStatusStatusCode = "FORBIDDEN"
	TeamListRequestStatusStatusCodeTimeout                    TeamListRequestStatusStatusCode = "TIMEOUT"
	TeamListRequestStatusStatusCodeExists                     TeamListRequestStatusStatusCode = "EXISTS"
	TeamListRequestStatusStatusCodeNotFound                   TeamListRequestStatusStatusCode = "NOT_FOUND"
	TeamListRequestStatusStatusCodeInternalError              TeamListRequestStatusStatusCode = "INTERNAL_ERROR"
	TeamListRequestStatusStatusCodeInvalidRequest             TeamListRequestStatusStatusCode = "INVALID_REQUEST"
	TeamListRequestStatusStatusCodeInvalidRequestVersion      TeamListRequestStatusStatusCode = "INVALID_REQUEST_VERSION"
	TeamListRequestStatusStatusCodeInvalidRequestData         TeamListRequestStatusStatusCode = "INVALID_REQUEST_DATA"
	TeamListRequestStatusStatusCodeMethodNotAllowed           TeamListRequestStatusStatusCode = "METHOD_NOT_ALLOWED"
	TeamListRequestStatusStatusCodeConflict                   TeamListRequestStatusStatusCode = "CONFLICT"
	TeamListRequestStatusStatusCodeUnprocessableEntity        TeamListRequestStatusStatusCode = "UNPROCESSABLE_ENTITY"
	TeamListRequestStatusStatusCodeTooManyRequests            TeamListRequestStatusStatusCode = "TOO_MANY_REQUESTS"
	TeamListRequestStatusStatusCodeInsufficientStorage        TeamListRequestStatusStatusCode = "INSUFFICIENT_STORAGE"
	TeamListRequestStatusStatusCodeServiceUnavailable         TeamListRequestStatusStatusCode = "SERVICE_UNAVAILABLE"
	TeamListRequestStatusStatusCodePayloadTooLarge            TeamListRequestStatusStatusCode = "PAYLOAD_TOO_LARGE"
	TeamListRequestStatusStatusCodeNotAcceptable              TeamListRequestStatusStatusCode = "NOT_ACCEPTABLE"
	TeamListRequestStatusStatusCodeUnavailableForLegalReasons TeamListRequestStatusStatusCode = "UNAVAILABLE_FOR_LEGAL_REASONS"
	TeamListRequestStatusStatusCodeBadGateway                 TeamListRequestStatusStatusCode = "BAD_GATEWAY"
)

func (r TeamListRequestStatusStatusCode) IsKnown() bool {
	switch r {
	case TeamListRequestStatusStatusCodeUnknown, TeamListRequestStatusStatusCodeSuccess, TeamListRequestStatusStatusCodeUnauthorized, TeamListRequestStatusStatusCodePaymentRequired, TeamListRequestStatusStatusCodeForbidden, TeamListRequestStatusStatusCodeTimeout, TeamListRequestStatusStatusCodeExists, TeamListRequestStatusStatusCodeNotFound, TeamListRequestStatusStatusCodeInternalError, TeamListRequestStatusStatusCodeInvalidRequest, TeamListRequestStatusStatusCodeInvalidRequestVersion, TeamListRequestStatusStatusCodeInvalidRequestData, TeamListRequestStatusStatusCodeMethodNotAllowed, TeamListRequestStatusStatusCodeConflict, TeamListRequestStatusStatusCodeUnprocessableEntity, TeamListRequestStatusStatusCodeTooManyRequests, TeamListRequestStatusStatusCodeInsufficientStorage, TeamListRequestStatusStatusCodeServiceUnavailable, TeamListRequestStatusStatusCodePayloadTooLarge, TeamListRequestStatusStatusCodeNotAcceptable, TeamListRequestStatusStatusCodeUnavailableForLegalReasons, TeamListRequestStatusStatusCodeBadGateway:
		return true
	}
	return false
}

// about one user
type User struct {
	// token needed to activate the user to enable login and other features
	ActivationToken string `json:"activationToken"`
	// NCA role
	NcaRole       UserNcaRole       `json:"ncaRole"`
	RequestStatus UserRequestStatus `json:"requestStatus"`
	// information about the user
	User UserUser `json:"user"`
	// DEPRECATED - Please use roles inside user
	UserRoles []UserUserRole `json:"userRoles"`
	JSON      userJSON       `json:"-"`
}

// userJSON contains the JSON metadata for the struct [User]
type userJSON struct {
	ActivationToken apijson.Field
	NcaRole         apijson.Field
	RequestStatus   apijson.Field
	User            apijson.Field
	UserRoles       apijson.Field
	raw             string
	ExtraFields     map[string]apijson.Field
}

func (r *User) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userJSON) RawJSON() string {
	return r.raw
}

// NCA role
type UserNcaRole string

const (
	UserNcaRoleUnknown       UserNcaRole = "UNKNOWN"
	UserNcaRoleAdministrator UserNcaRole = "ADMINISTRATOR"
	UserNcaRoleMember        UserNcaRole = "MEMBER"
	UserNcaRoleOwner         UserNcaRole = "OWNER"
	UserNcaRolePending       UserNcaRole = "PENDING"
)

func (r UserNcaRole) IsKnown() bool {
	switch r {
	case UserNcaRoleUnknown, UserNcaRoleAdministrator, UserNcaRoleMember, UserNcaRoleOwner, UserNcaRolePending:
		return true
	}
	return false
}

type UserRequestStatus struct {
	RequestID string `json:"requestId"`
	ServerID  string `json:"serverId"`
	// Describes response status reported by the server.
	StatusCode        UserRequestStatusStatusCode `json:"statusCode"`
	StatusDescription string                      `json:"statusDescription"`
	JSON              userRequestStatusJSON       `json:"-"`
}

// userRequestStatusJSON contains the JSON metadata for the struct
// [UserRequestStatus]
type userRequestStatusJSON struct {
	RequestID         apijson.Field
	ServerID          apijson.Field
	StatusCode        apijson.Field
	StatusDescription apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *UserRequestStatus) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userRequestStatusJSON) RawJSON() string {
	return r.raw
}

// Describes response status reported by the server.
type UserRequestStatusStatusCode string

const (
	UserRequestStatusStatusCodeUnknown                    UserRequestStatusStatusCode = "UNKNOWN"
	UserRequestStatusStatusCodeSuccess                    UserRequestStatusStatusCode = "SUCCESS"
	UserRequestStatusStatusCodeUnauthorized               UserRequestStatusStatusCode = "UNAUTHORIZED"
	UserRequestStatusStatusCodePaymentRequired            UserRequestStatusStatusCode = "PAYMENT_REQUIRED"
	UserRequestStatusStatusCodeForbidden                  UserRequestStatusStatusCode = "FORBIDDEN"
	UserRequestStatusStatusCodeTimeout                    UserRequestStatusStatusCode = "TIMEOUT"
	UserRequestStatusStatusCodeExists                     UserRequestStatusStatusCode = "EXISTS"
	UserRequestStatusStatusCodeNotFound                   UserRequestStatusStatusCode = "NOT_FOUND"
	UserRequestStatusStatusCodeInternalError              UserRequestStatusStatusCode = "INTERNAL_ERROR"
	UserRequestStatusStatusCodeInvalidRequest             UserRequestStatusStatusCode = "INVALID_REQUEST"
	UserRequestStatusStatusCodeInvalidRequestVersion      UserRequestStatusStatusCode = "INVALID_REQUEST_VERSION"
	UserRequestStatusStatusCodeInvalidRequestData         UserRequestStatusStatusCode = "INVALID_REQUEST_DATA"
	UserRequestStatusStatusCodeMethodNotAllowed           UserRequestStatusStatusCode = "METHOD_NOT_ALLOWED"
	UserRequestStatusStatusCodeConflict                   UserRequestStatusStatusCode = "CONFLICT"
	UserRequestStatusStatusCodeUnprocessableEntity        UserRequestStatusStatusCode = "UNPROCESSABLE_ENTITY"
	UserRequestStatusStatusCodeTooManyRequests            UserRequestStatusStatusCode = "TOO_MANY_REQUESTS"
	UserRequestStatusStatusCodeInsufficientStorage        UserRequestStatusStatusCode = "INSUFFICIENT_STORAGE"
	UserRequestStatusStatusCodeServiceUnavailable         UserRequestStatusStatusCode = "SERVICE_UNAVAILABLE"
	UserRequestStatusStatusCodePayloadTooLarge            UserRequestStatusStatusCode = "PAYLOAD_TOO_LARGE"
	UserRequestStatusStatusCodeNotAcceptable              UserRequestStatusStatusCode = "NOT_ACCEPTABLE"
	UserRequestStatusStatusCodeUnavailableForLegalReasons UserRequestStatusStatusCode = "UNAVAILABLE_FOR_LEGAL_REASONS"
	UserRequestStatusStatusCodeBadGateway                 UserRequestStatusStatusCode = "BAD_GATEWAY"
)

func (r UserRequestStatusStatusCode) IsKnown() bool {
	switch r {
	case UserRequestStatusStatusCodeUnknown, UserRequestStatusStatusCodeSuccess, UserRequestStatusStatusCodeUnauthorized, UserRequestStatusStatusCodePaymentRequired, UserRequestStatusStatusCodeForbidden, UserRequestStatusStatusCodeTimeout, UserRequestStatusStatusCodeExists, UserRequestStatusStatusCodeNotFound, UserRequestStatusStatusCodeInternalError, UserRequestStatusStatusCodeInvalidRequest, UserRequestStatusStatusCodeInvalidRequestVersion, UserRequestStatusStatusCodeInvalidRequestData, UserRequestStatusStatusCodeMethodNotAllowed, UserRequestStatusStatusCodeConflict, UserRequestStatusStatusCodeUnprocessableEntity, UserRequestStatusStatusCodeTooManyRequests, UserRequestStatusStatusCodeInsufficientStorage, UserRequestStatusStatusCodeServiceUnavailable, UserRequestStatusStatusCodePayloadTooLarge, UserRequestStatusStatusCodeNotAcceptable, UserRequestStatusStatusCodeUnavailableForLegalReasons, UserRequestStatusStatusCodeBadGateway:
		return true
	}
	return false
}

// information about the user
type UserUser struct {
	// unique Id of this user.
	ID int64 `json:"id"`
	// unique auth client id of this user.
	ClientID string `json:"clientId"`
	// Created date for this user
	CreatedDate string `json:"createdDate"`
	// Email address of the user. This should be unique.
	Email string `json:"email"`
	// Last time the user logged in
	FirstLoginDate string `json:"firstLoginDate"`
	// Determines if the user has beta access
	HasBetaAccess bool `json:"hasBetaAccess"`
	// indicate if user profile has been completed.
	HasProfile bool `json:"hasProfile"`
	// indicates if user has accepted AI Foundry Partnerships eula
	HasSignedAIFoundryPartnershipsEula bool `json:"hasSignedAiFoundryPartnershipsEULA"`
	// indicates if user has accepted Base Command End User License Agreement.
	HasSignedBaseCommandEula bool `json:"hasSignedBaseCommandEULA"`
	// indicates if user has accepted Base Command Manager End User License Agreement.
	HasSignedBaseCommandManagerEula bool `json:"hasSignedBaseCommandManagerEULA"`
	// indicates if user has accepted BioNeMo End User License Agreement.
	HasSignedBioNeMoEula bool `json:"hasSignedBioNeMoEULA"`
	// indicates if user has accepted container publishing eula
	HasSignedContainerPublishingEula bool `json:"hasSignedContainerPublishingEULA"`
	// indicates if user has accepted CuOpt eula
	HasSignedCuOptEula bool `json:"hasSignedCuOptEULA"`
	// indicates if user has accepted Earth-2 eula
	HasSignedEarth2Eula bool `json:"hasSignedEarth2EULA"`
	// [Deprecated] indicates if user has accepted EGX End User License Agreement.
	HasSignedEgxEula bool `json:"hasSignedEgxEULA"`
	// Determines if the user has signed the NGC End User License Agreement.
	HasSignedEula bool `json:"hasSignedEULA"`
	// indicates if user has accepted Fleet Command End User License Agreement.
	HasSignedFleetCommandEula bool `json:"hasSignedFleetCommandEULA"`
	// indicates if user has accepted LLM End User License Agreement.
	HasSignedLlmEula bool `json:"hasSignedLlmEULA"`
	// indicates if user has accepted Fleet Command End User License Agreement.
	HasSignedNvaieeula bool `json:"hasSignedNVAIEEULA"`
	// Determines if the user has signed the NVIDIA End User License Agreement.
	HasSignedNvidiaEula bool `json:"hasSignedNvidiaEULA"`
	// indicates if user has accepted Nvidia Quantum Cloud End User License Agreement.
	HasSignedNvqceula bool `json:"hasSignedNVQCEULA"`
	// indicates if user has accepted Omniverse End User License Agreement.
	HasSignedOmniverseEula bool `json:"hasSignedOmniverseEULA"`
	// Determines if the user has signed the Privacy Policy.
	HasSignedPrivacyPolicy bool `json:"hasSignedPrivacyPolicy"`
	// indicates if user has consented to share their registration info with other
	// parties
	HasSignedThirdPartyRegistryShareEula bool `json:"hasSignedThirdPartyRegistryShareEULA"`
	// Determines if the user has opted in email subscription.
	HasSubscribedToEmail bool `json:"hasSubscribedToEmail"`
	// Type of IDP, Identity Provider. Used for login.
	IdpType UserUserIdpType `json:"idpType"`
	// Determines if the user is active or not.
	IsActive bool `json:"isActive"`
	// Indicates if user was deleted from the system.
	IsDeleted bool `json:"isDeleted"`
	// Determines if the user is a SAML account or not.
	IsSAML bool `json:"isSAML"`
	// Title of user's job position.
	JobPositionTitle string `json:"jobPositionTitle"`
	// Last time the user logged in
	LastLoginDate string `json:"lastLoginDate"`
	// user name
	Name string `json:"name"`
	// List of roles that the user have
	Roles []UserUserRole `json:"roles"`
	// unique starfleet id of this user.
	StarfleetID string `json:"starfleetId"`
	// Storage quota for this user.
	StorageQuota []UserUserStorageQuota `json:"storageQuota"`
	// Updated date for this user
	UpdatedDate string `json:"updatedDate"`
	// Metadata information about the user.
	UserMetadata UserUserUserMetadata `json:"userMetadata"`
	JSON         userUserJSON         `json:"-"`
}

// userUserJSON contains the JSON metadata for the struct [UserUser]
type userUserJSON struct {
	ID                                   apijson.Field
	ClientID                             apijson.Field
	CreatedDate                          apijson.Field
	Email                                apijson.Field
	FirstLoginDate                       apijson.Field
	HasBetaAccess                        apijson.Field
	HasProfile                           apijson.Field
	HasSignedAIFoundryPartnershipsEula   apijson.Field
	HasSignedBaseCommandEula             apijson.Field
	HasSignedBaseCommandManagerEula      apijson.Field
	HasSignedBioNeMoEula                 apijson.Field
	HasSignedContainerPublishingEula     apijson.Field
	HasSignedCuOptEula                   apijson.Field
	HasSignedEarth2Eula                  apijson.Field
	HasSignedEgxEula                     apijson.Field
	HasSignedEula                        apijson.Field
	HasSignedFleetCommandEula            apijson.Field
	HasSignedLlmEula                     apijson.Field
	HasSignedNvaieeula                   apijson.Field
	HasSignedNvidiaEula                  apijson.Field
	HasSignedNvqceula                    apijson.Field
	HasSignedOmniverseEula               apijson.Field
	HasSignedPrivacyPolicy               apijson.Field
	HasSignedThirdPartyRegistryShareEula apijson.Field
	HasSubscribedToEmail                 apijson.Field
	IdpType                              apijson.Field
	IsActive                             apijson.Field
	IsDeleted                            apijson.Field
	IsSAML                               apijson.Field
	JobPositionTitle                     apijson.Field
	LastLoginDate                        apijson.Field
	Name                                 apijson.Field
	Roles                                apijson.Field
	StarfleetID                          apijson.Field
	StorageQuota                         apijson.Field
	UpdatedDate                          apijson.Field
	UserMetadata                         apijson.Field
	raw                                  string
	ExtraFields                          map[string]apijson.Field
}

func (r *UserUser) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userUserJSON) RawJSON() string {
	return r.raw
}

// Type of IDP, Identity Provider. Used for login.
type UserUserIdpType string

const (
	UserUserIdpTypeNvidia     UserUserIdpType = "NVIDIA"
	UserUserIdpTypeEnterprise UserUserIdpType = "ENTERPRISE"
)

func (r UserUserIdpType) IsKnown() bool {
	switch r {
	case UserUserIdpTypeNvidia, UserUserIdpTypeEnterprise:
		return true
	}
	return false
}

// List of roles that the user have
type UserUserRole struct {
	// Information about the Organization
	Org UserUserRolesOrg `json:"org"`
	// List of org role types that the user have
	OrgRoles []string `json:"orgRoles"`
	// DEPRECATED - List of role types that the user have
	RoleTypes []string `json:"roleTypes"`
	// Information about the user who is attempting to run the job
	TargetSystemUserIdentifier UserUserRolesTargetSystemUserIdentifier `json:"targetSystemUserIdentifier"`
	// Information about the team
	Team Team `json:"team"`
	// List of team role types that the user have
	TeamRoles []string         `json:"teamRoles"`
	JSON      userUserRoleJSON `json:"-"`
}

// userUserRoleJSON contains the JSON metadata for the struct [UserUserRole]
type userUserRoleJSON struct {
	Org                        apijson.Field
	OrgRoles                   apijson.Field
	RoleTypes                  apijson.Field
	TargetSystemUserIdentifier apijson.Field
	Team                       apijson.Field
	TeamRoles                  apijson.Field
	raw                        string
	ExtraFields                map[string]apijson.Field
}

func (r *UserUserRole) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userUserRoleJSON) RawJSON() string {
	return r.raw
}

// Information about the Organization
type UserUserRolesOrg struct {
	// Unique Id of this team.
	ID int64 `json:"id"`
	// Org Owner Alternate Contact
	AlternateContact UserUserRolesOrgAlternateContact `json:"alternateContact"`
	// Billing account ID.
	BillingAccountID string `json:"billingAccountId"`
	// Identifies if the org can be reused.
	CanAddOn bool `json:"canAddOn"`
	// ISO country code of the organization.
	Country string `json:"country"`
	// Optional description of the organization.
	Description string `json:"description"`
	// Name of the organization that will be shown to users.
	DisplayName string `json:"displayName"`
	// Identity Provider ID.
	IdpID string `json:"idpId"`
	// Industry of the organization.
	Industry string `json:"industry"`
	// Infinity manager setting definition
	InfinityManagerSettings UserUserRolesOrgInfinityManagerSettings `json:"infinityManagerSettings"`
	// Dataset Service enable flag for an organization
	IsDatasetServiceEnabled bool `json:"isDatasetServiceEnabled"`
	// Is NVIDIA internal org or not
	IsInternal bool `json:"isInternal"`
	// Indicates when the org is a proto org
	IsProto bool `json:"isProto"`
	// Quick Start enable flag for an organization
	IsQuickStartEnabled bool `json:"isQuickStartEnabled"`
	// If a server side encryption is enabled for private registry (models, resources)
	IsRegistrySseEnabled bool `json:"isRegistrySSEEnabled"`
	// Secrets Manager Service enable flag for an organization
	IsSecretsManagerServiceEnabled bool `json:"isSecretsManagerServiceEnabled"`
	// Secure Credential Sharing Service enable flag for an organization
	IsSecureCredentialSharingServiceEnabled bool `json:"isSecureCredentialSharingServiceEnabled"`
	// If a separate influx db used for an organization in BCP for job telemetry
	IsSeparateInfluxDBUsed bool `json:"isSeparateInfluxDbUsed"`
	// Organization name.
	Name string `json:"name"`
	// NVIDIA Cloud Account Number.
	Nan string `json:"nan"`
	// Org owner.
	OrgOwner UserUserRolesOrgOrgOwner `json:"orgOwner"`
	// Org owners
	OrgOwners []UserUserRolesOrgOrgOwner `json:"orgOwners"`
	// Product end customer salesforce.com Id (external customer Id). pecSfdcId is for
	// EMS (entitlement management service) to track external paid customer.
	PecSfdcID            string                                `json:"pecSfdcId"`
	ProductEnablements   []UserUserRolesOrgProductEnablement   `json:"productEnablements"`
	ProductSubscriptions []UserUserRolesOrgProductSubscription `json:"productSubscriptions"`
	// Repo scan setting definition
	RepoScanSettings UserUserRolesOrgRepoScanSettings `json:"repoScanSettings"`
	Type             UserUserRolesOrgType             `json:"type"`
	// Users information.
	UsersInfo UserUserRolesOrgUsersInfo `json:"usersInfo"`
	JSON      userUserRolesOrgJSON      `json:"-"`
}

// userUserRolesOrgJSON contains the JSON metadata for the struct
// [UserUserRolesOrg]
type userUserRolesOrgJSON struct {
	ID                                      apijson.Field
	AlternateContact                        apijson.Field
	BillingAccountID                        apijson.Field
	CanAddOn                                apijson.Field
	Country                                 apijson.Field
	Description                             apijson.Field
	DisplayName                             apijson.Field
	IdpID                                   apijson.Field
	Industry                                apijson.Field
	InfinityManagerSettings                 apijson.Field
	IsDatasetServiceEnabled                 apijson.Field
	IsInternal                              apijson.Field
	IsProto                                 apijson.Field
	IsQuickStartEnabled                     apijson.Field
	IsRegistrySseEnabled                    apijson.Field
	IsSecretsManagerServiceEnabled          apijson.Field
	IsSecureCredentialSharingServiceEnabled apijson.Field
	IsSeparateInfluxDBUsed                  apijson.Field
	Name                                    apijson.Field
	Nan                                     apijson.Field
	OrgOwner                                apijson.Field
	OrgOwners                               apijson.Field
	PecSfdcID                               apijson.Field
	ProductEnablements                      apijson.Field
	ProductSubscriptions                    apijson.Field
	RepoScanSettings                        apijson.Field
	Type                                    apijson.Field
	UsersInfo                               apijson.Field
	raw                                     string
	ExtraFields                             map[string]apijson.Field
}

func (r *UserUserRolesOrg) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userUserRolesOrgJSON) RawJSON() string {
	return r.raw
}

// Org Owner Alternate Contact
type UserUserRolesOrgAlternateContact struct {
	// Alternate contact's email.
	Email string `json:"email"`
	// Full name of the alternate contact.
	FullName string                               `json:"fullName"`
	JSON     userUserRolesOrgAlternateContactJSON `json:"-"`
}

// userUserRolesOrgAlternateContactJSON contains the JSON metadata for the struct
// [UserUserRolesOrgAlternateContact]
type userUserRolesOrgAlternateContactJSON struct {
	Email       apijson.Field
	FullName    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserUserRolesOrgAlternateContact) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userUserRolesOrgAlternateContactJSON) RawJSON() string {
	return r.raw
}

// Infinity manager setting definition
type UserUserRolesOrgInfinityManagerSettings struct {
	// Enable the infinity manager or not. Used both in org and team level object
	InfinityManagerEnabled bool `json:"infinityManagerEnabled"`
	// Allow override settings at team level. Only used in org level object
	InfinityManagerEnableTeamOverride bool                                        `json:"infinityManagerEnableTeamOverride"`
	JSON                              userUserRolesOrgInfinityManagerSettingsJSON `json:"-"`
}

// userUserRolesOrgInfinityManagerSettingsJSON contains the JSON metadata for the
// struct [UserUserRolesOrgInfinityManagerSettings]
type userUserRolesOrgInfinityManagerSettingsJSON struct {
	InfinityManagerEnabled            apijson.Field
	InfinityManagerEnableTeamOverride apijson.Field
	raw                               string
	ExtraFields                       map[string]apijson.Field
}

func (r *UserUserRolesOrgInfinityManagerSettings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userUserRolesOrgInfinityManagerSettingsJSON) RawJSON() string {
	return r.raw
}

// Org owner.
type UserUserRolesOrgOrgOwner struct {
	// Email address of the org owner.
	Email string `json:"email,required"`
	// Org owner name.
	FullName string `json:"fullName,required"`
	// Last time the org owner logged in.
	LastLoginDate string                       `json:"lastLoginDate"`
	JSON          userUserRolesOrgOrgOwnerJSON `json:"-"`
}

// userUserRolesOrgOrgOwnerJSON contains the JSON metadata for the struct
// [UserUserRolesOrgOrgOwner]
type userUserRolesOrgOrgOwnerJSON struct {
	Email         apijson.Field
	FullName      apijson.Field
	LastLoginDate apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *UserUserRolesOrgOrgOwner) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userUserRolesOrgOrgOwnerJSON) RawJSON() string {
	return r.raw
}

// Product Enablement
type UserUserRolesOrgProductEnablement struct {
	// Product Name (NVAIE, BASE_COMMAND, REGISTRY, etc)
	ProductName string `json:"productName,required"`
	// Product Enablement Types
	Type UserUserRolesOrgProductEnablementsType `json:"type,required"`
	// Date on which the subscription expires. The subscription is invalid after this
	// date. (yyyy-MM-dd)
	ExpirationDate string                                       `json:"expirationDate"`
	PoDetails      []UserUserRolesOrgProductEnablementsPoDetail `json:"poDetails"`
	JSON           userUserRolesOrgProductEnablementJSON        `json:"-"`
}

// userUserRolesOrgProductEnablementJSON contains the JSON metadata for the struct
// [UserUserRolesOrgProductEnablement]
type userUserRolesOrgProductEnablementJSON struct {
	ProductName    apijson.Field
	Type           apijson.Field
	ExpirationDate apijson.Field
	PoDetails      apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *UserUserRolesOrgProductEnablement) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userUserRolesOrgProductEnablementJSON) RawJSON() string {
	return r.raw
}

// Product Enablement Types
type UserUserRolesOrgProductEnablementsType string

const (
	UserUserRolesOrgProductEnablementsTypeNgcAdminEval       UserUserRolesOrgProductEnablementsType = "NGC_ADMIN_EVAL"
	UserUserRolesOrgProductEnablementsTypeNgcAdminNfr        UserUserRolesOrgProductEnablementsType = "NGC_ADMIN_NFR"
	UserUserRolesOrgProductEnablementsTypeNgcAdminCommercial UserUserRolesOrgProductEnablementsType = "NGC_ADMIN_COMMERCIAL"
	UserUserRolesOrgProductEnablementsTypeEmsEval            UserUserRolesOrgProductEnablementsType = "EMS_EVAL"
	UserUserRolesOrgProductEnablementsTypeEmsNfr             UserUserRolesOrgProductEnablementsType = "EMS_NFR"
	UserUserRolesOrgProductEnablementsTypeEmsCommercial      UserUserRolesOrgProductEnablementsType = "EMS_COMMERCIAL"
	UserUserRolesOrgProductEnablementsTypeNgcAdminDeveloper  UserUserRolesOrgProductEnablementsType = "NGC_ADMIN_DEVELOPER"
)

func (r UserUserRolesOrgProductEnablementsType) IsKnown() bool {
	switch r {
	case UserUserRolesOrgProductEnablementsTypeNgcAdminEval, UserUserRolesOrgProductEnablementsTypeNgcAdminNfr, UserUserRolesOrgProductEnablementsTypeNgcAdminCommercial, UserUserRolesOrgProductEnablementsTypeEmsEval, UserUserRolesOrgProductEnablementsTypeEmsNfr, UserUserRolesOrgProductEnablementsTypeEmsCommercial, UserUserRolesOrgProductEnablementsTypeNgcAdminDeveloper:
		return true
	}
	return false
}

// Purchase Order.
type UserUserRolesOrgProductEnablementsPoDetail struct {
	// Entitlement identifier.
	EntitlementID string `json:"entitlementId"`
	// PAK (Product Activation Key) identifier.
	PkID string                                         `json:"pkId"`
	JSON userUserRolesOrgProductEnablementsPoDetailJSON `json:"-"`
}

// userUserRolesOrgProductEnablementsPoDetailJSON contains the JSON metadata for
// the struct [UserUserRolesOrgProductEnablementsPoDetail]
type userUserRolesOrgProductEnablementsPoDetailJSON struct {
	EntitlementID apijson.Field
	PkID          apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *UserUserRolesOrgProductEnablementsPoDetail) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userUserRolesOrgProductEnablementsPoDetailJSON) RawJSON() string {
	return r.raw
}

// Product Subscription
type UserUserRolesOrgProductSubscription struct {
	// Product Name (NVAIE, BASE_COMMAND, FleetCommand, REGISTRY, etc).
	ProductName string `json:"productName,required"`
	// Unique entitlement identifier
	ID string `json:"id"`
	// EMS Subscription type. (options: EMS_EVAL, EMS_NFR and EMS_COMMERCIAL)
	EmsEntitlementType UserUserRolesOrgProductSubscriptionsEmsEntitlementType `json:"emsEntitlementType"`
	// Date on which the subscription expires. The subscription is invalid after this
	// date. (yyyy-MM-dd)
	ExpirationDate string `json:"expirationDate"`
	// Date on which the subscription becomes active. (yyyy-MM-dd)
	StartDate string `json:"startDate"`
	// Subscription type. (options: NGC_ADMIN_EVAL, NGC_ADMIN_NFR,
	// NGC_ADMIN_COMMERCIAL)
	Type UserUserRolesOrgProductSubscriptionsType `json:"type"`
	JSON userUserRolesOrgProductSubscriptionJSON  `json:"-"`
}

// userUserRolesOrgProductSubscriptionJSON contains the JSON metadata for the
// struct [UserUserRolesOrgProductSubscription]
type userUserRolesOrgProductSubscriptionJSON struct {
	ProductName        apijson.Field
	ID                 apijson.Field
	EmsEntitlementType apijson.Field
	ExpirationDate     apijson.Field
	StartDate          apijson.Field
	Type               apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *UserUserRolesOrgProductSubscription) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userUserRolesOrgProductSubscriptionJSON) RawJSON() string {
	return r.raw
}

// EMS Subscription type. (options: EMS_EVAL, EMS_NFR and EMS_COMMERCIAL)
type UserUserRolesOrgProductSubscriptionsEmsEntitlementType string

const (
	UserUserRolesOrgProductSubscriptionsEmsEntitlementTypeEmsEval       UserUserRolesOrgProductSubscriptionsEmsEntitlementType = "EMS_EVAL"
	UserUserRolesOrgProductSubscriptionsEmsEntitlementTypeEmsNfr        UserUserRolesOrgProductSubscriptionsEmsEntitlementType = "EMS_NFR"
	UserUserRolesOrgProductSubscriptionsEmsEntitlementTypeEmsCommerical UserUserRolesOrgProductSubscriptionsEmsEntitlementType = "EMS_COMMERICAL"
	UserUserRolesOrgProductSubscriptionsEmsEntitlementTypeEmsCommercial UserUserRolesOrgProductSubscriptionsEmsEntitlementType = "EMS_COMMERCIAL"
)

func (r UserUserRolesOrgProductSubscriptionsEmsEntitlementType) IsKnown() bool {
	switch r {
	case UserUserRolesOrgProductSubscriptionsEmsEntitlementTypeEmsEval, UserUserRolesOrgProductSubscriptionsEmsEntitlementTypeEmsNfr, UserUserRolesOrgProductSubscriptionsEmsEntitlementTypeEmsCommerical, UserUserRolesOrgProductSubscriptionsEmsEntitlementTypeEmsCommercial:
		return true
	}
	return false
}

// Subscription type. (options: NGC_ADMIN_EVAL, NGC_ADMIN_NFR,
// NGC_ADMIN_COMMERCIAL)
type UserUserRolesOrgProductSubscriptionsType string

const (
	UserUserRolesOrgProductSubscriptionsTypeNgcAdminEval       UserUserRolesOrgProductSubscriptionsType = "NGC_ADMIN_EVAL"
	UserUserRolesOrgProductSubscriptionsTypeNgcAdminNfr        UserUserRolesOrgProductSubscriptionsType = "NGC_ADMIN_NFR"
	UserUserRolesOrgProductSubscriptionsTypeNgcAdminCommercial UserUserRolesOrgProductSubscriptionsType = "NGC_ADMIN_COMMERCIAL"
)

func (r UserUserRolesOrgProductSubscriptionsType) IsKnown() bool {
	switch r {
	case UserUserRolesOrgProductSubscriptionsTypeNgcAdminEval, UserUserRolesOrgProductSubscriptionsTypeNgcAdminNfr, UserUserRolesOrgProductSubscriptionsTypeNgcAdminCommercial:
		return true
	}
	return false
}

// Repo scan setting definition
type UserUserRolesOrgRepoScanSettings struct {
	// Allow org admin to override the org level repo scan settings
	RepoScanAllowOverride bool `json:"repoScanAllowOverride"`
	// Allow repository scanning by default
	RepoScanByDefault bool `json:"repoScanByDefault"`
	// Enable the repository scan or not. Only used in org level object
	RepoScanEnabled bool `json:"repoScanEnabled"`
	// Sends notification to end user after scanning is done
	RepoScanEnableNotifications bool `json:"repoScanEnableNotifications"`
	// Allow override settings at team level. Only used in org level object
	RepoScanEnableTeamOverride bool `json:"repoScanEnableTeamOverride"`
	// Allow showing scan results to CLI or UI
	RepoScanShowResults bool                                 `json:"repoScanShowResults"`
	JSON                userUserRolesOrgRepoScanSettingsJSON `json:"-"`
}

// userUserRolesOrgRepoScanSettingsJSON contains the JSON metadata for the struct
// [UserUserRolesOrgRepoScanSettings]
type userUserRolesOrgRepoScanSettingsJSON struct {
	RepoScanAllowOverride       apijson.Field
	RepoScanByDefault           apijson.Field
	RepoScanEnabled             apijson.Field
	RepoScanEnableNotifications apijson.Field
	RepoScanEnableTeamOverride  apijson.Field
	RepoScanShowResults         apijson.Field
	raw                         string
	ExtraFields                 map[string]apijson.Field
}

func (r *UserUserRolesOrgRepoScanSettings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userUserRolesOrgRepoScanSettingsJSON) RawJSON() string {
	return r.raw
}

type UserUserRolesOrgType string

const (
	UserUserRolesOrgTypeUnknown    UserUserRolesOrgType = "UNKNOWN"
	UserUserRolesOrgTypeCloud      UserUserRolesOrgType = "CLOUD"
	UserUserRolesOrgTypeEnterprise UserUserRolesOrgType = "ENTERPRISE"
	UserUserRolesOrgTypeIndividual UserUserRolesOrgType = "INDIVIDUAL"
)

func (r UserUserRolesOrgType) IsKnown() bool {
	switch r {
	case UserUserRolesOrgTypeUnknown, UserUserRolesOrgTypeCloud, UserUserRolesOrgTypeEnterprise, UserUserRolesOrgTypeIndividual:
		return true
	}
	return false
}

// Users information.
type UserUserRolesOrgUsersInfo struct {
	// Total number of users.
	TotalUsers int64                         `json:"totalUsers"`
	JSON       userUserRolesOrgUsersInfoJSON `json:"-"`
}

// userUserRolesOrgUsersInfoJSON contains the JSON metadata for the struct
// [UserUserRolesOrgUsersInfo]
type userUserRolesOrgUsersInfoJSON struct {
	TotalUsers  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserUserRolesOrgUsersInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userUserRolesOrgUsersInfoJSON) RawJSON() string {
	return r.raw
}

// Information about the user who is attempting to run the job
type UserUserRolesTargetSystemUserIdentifier struct {
	// gid of the user on this team
	Gid int64 `json:"gid"`
	// Org context for the job
	OrgName string `json:"orgName"`
	// Starfleet ID of the user creating the job.
	StarfleetID string `json:"starfleetId"`
	// Team context for the job
	TeamName string `json:"teamName"`
	// uid of the user on this team
	Uid int64 `json:"uid"`
	// Unique ID of the user who submitted the job
	UserID int64                                       `json:"userId"`
	JSON   userUserRolesTargetSystemUserIdentifierJSON `json:"-"`
}

// userUserRolesTargetSystemUserIdentifierJSON contains the JSON metadata for the
// struct [UserUserRolesTargetSystemUserIdentifier]
type userUserRolesTargetSystemUserIdentifierJSON struct {
	Gid         apijson.Field
	OrgName     apijson.Field
	StarfleetID apijson.Field
	TeamName    apijson.Field
	Uid         apijson.Field
	UserID      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserUserRolesTargetSystemUserIdentifier) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userUserRolesTargetSystemUserIdentifierJSON) RawJSON() string {
	return r.raw
}

// represents user storage quota for a given ace and available unused storage
type UserUserStorageQuota struct {
	// id of the ace
	AceID int64 `json:"aceId"`
	// name of the ace
	AceName string `json:"aceName"`
	// Available space in bytes
	Available int64 `json:"available"`
	// Number of datasets that are a part of user's used storage
	DatasetCount int64 `json:"datasetCount"`
	// Space used by datasets in bytes
	DatasetsUsage int64 `json:"datasetsUsage"`
	// The org name that this user quota tied to. This is needed for analytics
	OrgName string `json:"orgName"`
	// Assigned quota in bytes
	Quota int64 `json:"quota"`
	// Number of resultsets that are a part of user's used storage
	ResultsetCount int64 `json:"resultsetCount"`
	// Space used by resultsets in bytes
	ResultsetsUsage int64 `json:"resultsetsUsage"`
	// Description of this storage cluster
	StorageClusterDescription string `json:"storageClusterDescription"`
	// Name of storage cluster
	StorageClusterName string `json:"storageClusterName"`
	// Identifier to this storage cluster
	StorageClusterUuid string `json:"storageClusterUuid"`
	// Number of workspaces that are a part of user's used storage
	WorkspacesCount int64 `json:"workspacesCount"`
	// Space used by workspaces in bytes
	WorkspacesUsage int64                    `json:"workspacesUsage"`
	JSON            userUserStorageQuotaJSON `json:"-"`
}

// userUserStorageQuotaJSON contains the JSON metadata for the struct
// [UserUserStorageQuota]
type userUserStorageQuotaJSON struct {
	AceID                     apijson.Field
	AceName                   apijson.Field
	Available                 apijson.Field
	DatasetCount              apijson.Field
	DatasetsUsage             apijson.Field
	OrgName                   apijson.Field
	Quota                     apijson.Field
	ResultsetCount            apijson.Field
	ResultsetsUsage           apijson.Field
	StorageClusterDescription apijson.Field
	StorageClusterName        apijson.Field
	StorageClusterUuid        apijson.Field
	WorkspacesCount           apijson.Field
	WorkspacesUsage           apijson.Field
	raw                       string
	ExtraFields               map[string]apijson.Field
}

func (r *UserUserStorageQuota) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userUserStorageQuotaJSON) RawJSON() string {
	return r.raw
}

// Metadata information about the user.
type UserUserUserMetadata struct {
	// Name of the company
	Company string `json:"company"`
	// Company URL
	CompanyURL string `json:"companyUrl"`
	// Country of the user
	Country string `json:"country"`
	// User's first name
	FirstName string `json:"firstName"`
	// Industry segment
	Industry string `json:"industry"`
	// List of development areas that user has interest
	Interest []string `json:"interest"`
	// User's last name
	LastName string `json:"lastName"`
	// Role of the user in the company
	Role string                   `json:"role"`
	JSON userUserUserMetadataJSON `json:"-"`
}

// userUserUserMetadataJSON contains the JSON metadata for the struct
// [UserUserUserMetadata]
type userUserUserMetadataJSON struct {
	Company     apijson.Field
	CompanyURL  apijson.Field
	Country     apijson.Field
	FirstName   apijson.Field
	Industry    apijson.Field
	Interest    apijson.Field
	LastName    apijson.Field
	Role        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserUserUserMetadata) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userUserUserMetadataJSON) RawJSON() string {
	return r.raw
}

// Response for List User reponse
type UserList struct {
	// object that describes the pagination information
	PaginationInfo UserListPaginationInfo `json:"paginationInfo"`
	RequestStatus  UserListRequestStatus  `json:"requestStatus"`
	// information about the user
	Users []UserListUser `json:"users"`
	JSON  userListJSON   `json:"-"`
}

// userListJSON contains the JSON metadata for the struct [UserList]
type userListJSON struct {
	PaginationInfo apijson.Field
	RequestStatus  apijson.Field
	Users          apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *UserList) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userListJSON) RawJSON() string {
	return r.raw
}

// object that describes the pagination information
type UserListPaginationInfo struct {
	// Page index of results
	Index int64 `json:"index"`
	// Serialized pointer to the next results page. Should be used for fetching next
	// page. Can be empty
	NextPage string `json:"nextPage"`
	// Number of results in page
	Size int64 `json:"size"`
	// Total number of pages available
	TotalPages int64 `json:"totalPages"`
	// Total number of results available
	TotalResults int64                      `json:"totalResults"`
	JSON         userListPaginationInfoJSON `json:"-"`
}

// userListPaginationInfoJSON contains the JSON metadata for the struct
// [UserListPaginationInfo]
type userListPaginationInfoJSON struct {
	Index        apijson.Field
	NextPage     apijson.Field
	Size         apijson.Field
	TotalPages   apijson.Field
	TotalResults apijson.Field
	raw          string
	ExtraFields  map[string]apijson.Field
}

func (r *UserListPaginationInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userListPaginationInfoJSON) RawJSON() string {
	return r.raw
}

type UserListRequestStatus struct {
	RequestID string `json:"requestId"`
	ServerID  string `json:"serverId"`
	// Describes response status reported by the server.
	StatusCode        UserListRequestStatusStatusCode `json:"statusCode"`
	StatusDescription string                          `json:"statusDescription"`
	JSON              userListRequestStatusJSON       `json:"-"`
}

// userListRequestStatusJSON contains the JSON metadata for the struct
// [UserListRequestStatus]
type userListRequestStatusJSON struct {
	RequestID         apijson.Field
	ServerID          apijson.Field
	StatusCode        apijson.Field
	StatusDescription apijson.Field
	raw               string
	ExtraFields       map[string]apijson.Field
}

func (r *UserListRequestStatus) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userListRequestStatusJSON) RawJSON() string {
	return r.raw
}

// Describes response status reported by the server.
type UserListRequestStatusStatusCode string

const (
	UserListRequestStatusStatusCodeUnknown                    UserListRequestStatusStatusCode = "UNKNOWN"
	UserListRequestStatusStatusCodeSuccess                    UserListRequestStatusStatusCode = "SUCCESS"
	UserListRequestStatusStatusCodeUnauthorized               UserListRequestStatusStatusCode = "UNAUTHORIZED"
	UserListRequestStatusStatusCodePaymentRequired            UserListRequestStatusStatusCode = "PAYMENT_REQUIRED"
	UserListRequestStatusStatusCodeForbidden                  UserListRequestStatusStatusCode = "FORBIDDEN"
	UserListRequestStatusStatusCodeTimeout                    UserListRequestStatusStatusCode = "TIMEOUT"
	UserListRequestStatusStatusCodeExists                     UserListRequestStatusStatusCode = "EXISTS"
	UserListRequestStatusStatusCodeNotFound                   UserListRequestStatusStatusCode = "NOT_FOUND"
	UserListRequestStatusStatusCodeInternalError              UserListRequestStatusStatusCode = "INTERNAL_ERROR"
	UserListRequestStatusStatusCodeInvalidRequest             UserListRequestStatusStatusCode = "INVALID_REQUEST"
	UserListRequestStatusStatusCodeInvalidRequestVersion      UserListRequestStatusStatusCode = "INVALID_REQUEST_VERSION"
	UserListRequestStatusStatusCodeInvalidRequestData         UserListRequestStatusStatusCode = "INVALID_REQUEST_DATA"
	UserListRequestStatusStatusCodeMethodNotAllowed           UserListRequestStatusStatusCode = "METHOD_NOT_ALLOWED"
	UserListRequestStatusStatusCodeConflict                   UserListRequestStatusStatusCode = "CONFLICT"
	UserListRequestStatusStatusCodeUnprocessableEntity        UserListRequestStatusStatusCode = "UNPROCESSABLE_ENTITY"
	UserListRequestStatusStatusCodeTooManyRequests            UserListRequestStatusStatusCode = "TOO_MANY_REQUESTS"
	UserListRequestStatusStatusCodeInsufficientStorage        UserListRequestStatusStatusCode = "INSUFFICIENT_STORAGE"
	UserListRequestStatusStatusCodeServiceUnavailable         UserListRequestStatusStatusCode = "SERVICE_UNAVAILABLE"
	UserListRequestStatusStatusCodePayloadTooLarge            UserListRequestStatusStatusCode = "PAYLOAD_TOO_LARGE"
	UserListRequestStatusStatusCodeNotAcceptable              UserListRequestStatusStatusCode = "NOT_ACCEPTABLE"
	UserListRequestStatusStatusCodeUnavailableForLegalReasons UserListRequestStatusStatusCode = "UNAVAILABLE_FOR_LEGAL_REASONS"
	UserListRequestStatusStatusCodeBadGateway                 UserListRequestStatusStatusCode = "BAD_GATEWAY"
)

func (r UserListRequestStatusStatusCode) IsKnown() bool {
	switch r {
	case UserListRequestStatusStatusCodeUnknown, UserListRequestStatusStatusCodeSuccess, UserListRequestStatusStatusCodeUnauthorized, UserListRequestStatusStatusCodePaymentRequired, UserListRequestStatusStatusCodeForbidden, UserListRequestStatusStatusCodeTimeout, UserListRequestStatusStatusCodeExists, UserListRequestStatusStatusCodeNotFound, UserListRequestStatusStatusCodeInternalError, UserListRequestStatusStatusCodeInvalidRequest, UserListRequestStatusStatusCodeInvalidRequestVersion, UserListRequestStatusStatusCodeInvalidRequestData, UserListRequestStatusStatusCodeMethodNotAllowed, UserListRequestStatusStatusCodeConflict, UserListRequestStatusStatusCodeUnprocessableEntity, UserListRequestStatusStatusCodeTooManyRequests, UserListRequestStatusStatusCodeInsufficientStorage, UserListRequestStatusStatusCodeServiceUnavailable, UserListRequestStatusStatusCodePayloadTooLarge, UserListRequestStatusStatusCodeNotAcceptable, UserListRequestStatusStatusCodeUnavailableForLegalReasons, UserListRequestStatusStatusCodeBadGateway:
		return true
	}
	return false
}

// information about the user
type UserListUser struct {
	// unique Id of this user.
	ID int64 `json:"id"`
	// unique auth client id of this user.
	ClientID string `json:"clientId"`
	// Created date for this user
	CreatedDate string `json:"createdDate"`
	// Email address of the user. This should be unique.
	Email string `json:"email"`
	// Last time the user logged in
	FirstLoginDate string `json:"firstLoginDate"`
	// Determines if the user has beta access
	HasBetaAccess bool `json:"hasBetaAccess"`
	// indicate if user profile has been completed.
	HasProfile bool `json:"hasProfile"`
	// indicates if user has accepted AI Foundry Partnerships eula
	HasSignedAIFoundryPartnershipsEula bool `json:"hasSignedAiFoundryPartnershipsEULA"`
	// indicates if user has accepted Base Command End User License Agreement.
	HasSignedBaseCommandEula bool `json:"hasSignedBaseCommandEULA"`
	// indicates if user has accepted Base Command Manager End User License Agreement.
	HasSignedBaseCommandManagerEula bool `json:"hasSignedBaseCommandManagerEULA"`
	// indicates if user has accepted BioNeMo End User License Agreement.
	HasSignedBioNeMoEula bool `json:"hasSignedBioNeMoEULA"`
	// indicates if user has accepted container publishing eula
	HasSignedContainerPublishingEula bool `json:"hasSignedContainerPublishingEULA"`
	// indicates if user has accepted CuOpt eula
	HasSignedCuOptEula bool `json:"hasSignedCuOptEULA"`
	// indicates if user has accepted Earth-2 eula
	HasSignedEarth2Eula bool `json:"hasSignedEarth2EULA"`
	// [Deprecated] indicates if user has accepted EGX End User License Agreement.
	HasSignedEgxEula bool `json:"hasSignedEgxEULA"`
	// Determines if the user has signed the NGC End User License Agreement.
	HasSignedEula bool `json:"hasSignedEULA"`
	// indicates if user has accepted Fleet Command End User License Agreement.
	HasSignedFleetCommandEula bool `json:"hasSignedFleetCommandEULA"`
	// indicates if user has accepted LLM End User License Agreement.
	HasSignedLlmEula bool `json:"hasSignedLlmEULA"`
	// indicates if user has accepted Fleet Command End User License Agreement.
	HasSignedNvaieeula bool `json:"hasSignedNVAIEEULA"`
	// Determines if the user has signed the NVIDIA End User License Agreement.
	HasSignedNvidiaEula bool `json:"hasSignedNvidiaEULA"`
	// indicates if user has accepted Nvidia Quantum Cloud End User License Agreement.
	HasSignedNvqceula bool `json:"hasSignedNVQCEULA"`
	// indicates if user has accepted Omniverse End User License Agreement.
	HasSignedOmniverseEula bool `json:"hasSignedOmniverseEULA"`
	// Determines if the user has signed the Privacy Policy.
	HasSignedPrivacyPolicy bool `json:"hasSignedPrivacyPolicy"`
	// indicates if user has consented to share their registration info with other
	// parties
	HasSignedThirdPartyRegistryShareEula bool `json:"hasSignedThirdPartyRegistryShareEULA"`
	// Determines if the user has opted in email subscription.
	HasSubscribedToEmail bool `json:"hasSubscribedToEmail"`
	// Type of IDP, Identity Provider. Used for login.
	IdpType UserListUsersIdpType `json:"idpType"`
	// Determines if the user is active or not.
	IsActive bool `json:"isActive"`
	// Indicates if user was deleted from the system.
	IsDeleted bool `json:"isDeleted"`
	// Determines if the user is a SAML account or not.
	IsSAML bool `json:"isSAML"`
	// Title of user's job position.
	JobPositionTitle string `json:"jobPositionTitle"`
	// Last time the user logged in
	LastLoginDate string `json:"lastLoginDate"`
	// user name
	Name string `json:"name"`
	// List of roles that the user have
	Roles []UserListUsersRole `json:"roles"`
	// unique starfleet id of this user.
	StarfleetID string `json:"starfleetId"`
	// Storage quota for this user.
	StorageQuota []UserListUsersStorageQuota `json:"storageQuota"`
	// Updated date for this user
	UpdatedDate string `json:"updatedDate"`
	// Metadata information about the user.
	UserMetadata UserListUsersUserMetadata `json:"userMetadata"`
	JSON         userListUserJSON          `json:"-"`
}

// userListUserJSON contains the JSON metadata for the struct [UserListUser]
type userListUserJSON struct {
	ID                                   apijson.Field
	ClientID                             apijson.Field
	CreatedDate                          apijson.Field
	Email                                apijson.Field
	FirstLoginDate                       apijson.Field
	HasBetaAccess                        apijson.Field
	HasProfile                           apijson.Field
	HasSignedAIFoundryPartnershipsEula   apijson.Field
	HasSignedBaseCommandEula             apijson.Field
	HasSignedBaseCommandManagerEula      apijson.Field
	HasSignedBioNeMoEula                 apijson.Field
	HasSignedContainerPublishingEula     apijson.Field
	HasSignedCuOptEula                   apijson.Field
	HasSignedEarth2Eula                  apijson.Field
	HasSignedEgxEula                     apijson.Field
	HasSignedEula                        apijson.Field
	HasSignedFleetCommandEula            apijson.Field
	HasSignedLlmEula                     apijson.Field
	HasSignedNvaieeula                   apijson.Field
	HasSignedNvidiaEula                  apijson.Field
	HasSignedNvqceula                    apijson.Field
	HasSignedOmniverseEula               apijson.Field
	HasSignedPrivacyPolicy               apijson.Field
	HasSignedThirdPartyRegistryShareEula apijson.Field
	HasSubscribedToEmail                 apijson.Field
	IdpType                              apijson.Field
	IsActive                             apijson.Field
	IsDeleted                            apijson.Field
	IsSAML                               apijson.Field
	JobPositionTitle                     apijson.Field
	LastLoginDate                        apijson.Field
	Name                                 apijson.Field
	Roles                                apijson.Field
	StarfleetID                          apijson.Field
	StorageQuota                         apijson.Field
	UpdatedDate                          apijson.Field
	UserMetadata                         apijson.Field
	raw                                  string
	ExtraFields                          map[string]apijson.Field
}

func (r *UserListUser) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userListUserJSON) RawJSON() string {
	return r.raw
}

// Type of IDP, Identity Provider. Used for login.
type UserListUsersIdpType string

const (
	UserListUsersIdpTypeNvidia     UserListUsersIdpType = "NVIDIA"
	UserListUsersIdpTypeEnterprise UserListUsersIdpType = "ENTERPRISE"
)

func (r UserListUsersIdpType) IsKnown() bool {
	switch r {
	case UserListUsersIdpTypeNvidia, UserListUsersIdpTypeEnterprise:
		return true
	}
	return false
}

// List of roles that the user have
type UserListUsersRole struct {
	// Information about the Organization
	Org UserListUsersRolesOrg `json:"org"`
	// List of org role types that the user have
	OrgRoles []string `json:"orgRoles"`
	// DEPRECATED - List of role types that the user have
	RoleTypes []string `json:"roleTypes"`
	// Information about the user who is attempting to run the job
	TargetSystemUserIdentifier UserListUsersRolesTargetSystemUserIdentifier `json:"targetSystemUserIdentifier"`
	// Information about the team
	Team Team `json:"team"`
	// List of team role types that the user have
	TeamRoles []string              `json:"teamRoles"`
	JSON      userListUsersRoleJSON `json:"-"`
}

// userListUsersRoleJSON contains the JSON metadata for the struct
// [UserListUsersRole]
type userListUsersRoleJSON struct {
	Org                        apijson.Field
	OrgRoles                   apijson.Field
	RoleTypes                  apijson.Field
	TargetSystemUserIdentifier apijson.Field
	Team                       apijson.Field
	TeamRoles                  apijson.Field
	raw                        string
	ExtraFields                map[string]apijson.Field
}

func (r *UserListUsersRole) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userListUsersRoleJSON) RawJSON() string {
	return r.raw
}

// Information about the Organization
type UserListUsersRolesOrg struct {
	// Unique Id of this team.
	ID int64 `json:"id"`
	// Org Owner Alternate Contact
	AlternateContact UserListUsersRolesOrgAlternateContact `json:"alternateContact"`
	// Billing account ID.
	BillingAccountID string `json:"billingAccountId"`
	// Identifies if the org can be reused.
	CanAddOn bool `json:"canAddOn"`
	// ISO country code of the organization.
	Country string `json:"country"`
	// Optional description of the organization.
	Description string `json:"description"`
	// Name of the organization that will be shown to users.
	DisplayName string `json:"displayName"`
	// Identity Provider ID.
	IdpID string `json:"idpId"`
	// Industry of the organization.
	Industry string `json:"industry"`
	// Infinity manager setting definition
	InfinityManagerSettings UserListUsersRolesOrgInfinityManagerSettings `json:"infinityManagerSettings"`
	// Dataset Service enable flag for an organization
	IsDatasetServiceEnabled bool `json:"isDatasetServiceEnabled"`
	// Is NVIDIA internal org or not
	IsInternal bool `json:"isInternal"`
	// Indicates when the org is a proto org
	IsProto bool `json:"isProto"`
	// Quick Start enable flag for an organization
	IsQuickStartEnabled bool `json:"isQuickStartEnabled"`
	// If a server side encryption is enabled for private registry (models, resources)
	IsRegistrySseEnabled bool `json:"isRegistrySSEEnabled"`
	// Secrets Manager Service enable flag for an organization
	IsSecretsManagerServiceEnabled bool `json:"isSecretsManagerServiceEnabled"`
	// Secure Credential Sharing Service enable flag for an organization
	IsSecureCredentialSharingServiceEnabled bool `json:"isSecureCredentialSharingServiceEnabled"`
	// If a separate influx db used for an organization in BCP for job telemetry
	IsSeparateInfluxDBUsed bool `json:"isSeparateInfluxDbUsed"`
	// Organization name.
	Name string `json:"name"`
	// NVIDIA Cloud Account Number.
	Nan string `json:"nan"`
	// Org owner.
	OrgOwner UserListUsersRolesOrgOrgOwner `json:"orgOwner"`
	// Org owners
	OrgOwners []UserListUsersRolesOrgOrgOwner `json:"orgOwners"`
	// Product end customer salesforce.com Id (external customer Id). pecSfdcId is for
	// EMS (entitlement management service) to track external paid customer.
	PecSfdcID            string                                     `json:"pecSfdcId"`
	ProductEnablements   []UserListUsersRolesOrgProductEnablement   `json:"productEnablements"`
	ProductSubscriptions []UserListUsersRolesOrgProductSubscription `json:"productSubscriptions"`
	// Repo scan setting definition
	RepoScanSettings UserListUsersRolesOrgRepoScanSettings `json:"repoScanSettings"`
	Type             UserListUsersRolesOrgType             `json:"type"`
	// Users information.
	UsersInfo UserListUsersRolesOrgUsersInfo `json:"usersInfo"`
	JSON      userListUsersRolesOrgJSON      `json:"-"`
}

// userListUsersRolesOrgJSON contains the JSON metadata for the struct
// [UserListUsersRolesOrg]
type userListUsersRolesOrgJSON struct {
	ID                                      apijson.Field
	AlternateContact                        apijson.Field
	BillingAccountID                        apijson.Field
	CanAddOn                                apijson.Field
	Country                                 apijson.Field
	Description                             apijson.Field
	DisplayName                             apijson.Field
	IdpID                                   apijson.Field
	Industry                                apijson.Field
	InfinityManagerSettings                 apijson.Field
	IsDatasetServiceEnabled                 apijson.Field
	IsInternal                              apijson.Field
	IsProto                                 apijson.Field
	IsQuickStartEnabled                     apijson.Field
	IsRegistrySseEnabled                    apijson.Field
	IsSecretsManagerServiceEnabled          apijson.Field
	IsSecureCredentialSharingServiceEnabled apijson.Field
	IsSeparateInfluxDBUsed                  apijson.Field
	Name                                    apijson.Field
	Nan                                     apijson.Field
	OrgOwner                                apijson.Field
	OrgOwners                               apijson.Field
	PecSfdcID                               apijson.Field
	ProductEnablements                      apijson.Field
	ProductSubscriptions                    apijson.Field
	RepoScanSettings                        apijson.Field
	Type                                    apijson.Field
	UsersInfo                               apijson.Field
	raw                                     string
	ExtraFields                             map[string]apijson.Field
}

func (r *UserListUsersRolesOrg) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userListUsersRolesOrgJSON) RawJSON() string {
	return r.raw
}

// Org Owner Alternate Contact
type UserListUsersRolesOrgAlternateContact struct {
	// Alternate contact's email.
	Email string `json:"email"`
	// Full name of the alternate contact.
	FullName string                                    `json:"fullName"`
	JSON     userListUsersRolesOrgAlternateContactJSON `json:"-"`
}

// userListUsersRolesOrgAlternateContactJSON contains the JSON metadata for the
// struct [UserListUsersRolesOrgAlternateContact]
type userListUsersRolesOrgAlternateContactJSON struct {
	Email       apijson.Field
	FullName    apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserListUsersRolesOrgAlternateContact) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userListUsersRolesOrgAlternateContactJSON) RawJSON() string {
	return r.raw
}

// Infinity manager setting definition
type UserListUsersRolesOrgInfinityManagerSettings struct {
	// Enable the infinity manager or not. Used both in org and team level object
	InfinityManagerEnabled bool `json:"infinityManagerEnabled"`
	// Allow override settings at team level. Only used in org level object
	InfinityManagerEnableTeamOverride bool                                             `json:"infinityManagerEnableTeamOverride"`
	JSON                              userListUsersRolesOrgInfinityManagerSettingsJSON `json:"-"`
}

// userListUsersRolesOrgInfinityManagerSettingsJSON contains the JSON metadata for
// the struct [UserListUsersRolesOrgInfinityManagerSettings]
type userListUsersRolesOrgInfinityManagerSettingsJSON struct {
	InfinityManagerEnabled            apijson.Field
	InfinityManagerEnableTeamOverride apijson.Field
	raw                               string
	ExtraFields                       map[string]apijson.Field
}

func (r *UserListUsersRolesOrgInfinityManagerSettings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userListUsersRolesOrgInfinityManagerSettingsJSON) RawJSON() string {
	return r.raw
}

// Org owner.
type UserListUsersRolesOrgOrgOwner struct {
	// Email address of the org owner.
	Email string `json:"email,required"`
	// Org owner name.
	FullName string `json:"fullName,required"`
	// Last time the org owner logged in.
	LastLoginDate string                            `json:"lastLoginDate"`
	JSON          userListUsersRolesOrgOrgOwnerJSON `json:"-"`
}

// userListUsersRolesOrgOrgOwnerJSON contains the JSON metadata for the struct
// [UserListUsersRolesOrgOrgOwner]
type userListUsersRolesOrgOrgOwnerJSON struct {
	Email         apijson.Field
	FullName      apijson.Field
	LastLoginDate apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *UserListUsersRolesOrgOrgOwner) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userListUsersRolesOrgOrgOwnerJSON) RawJSON() string {
	return r.raw
}

// Product Enablement
type UserListUsersRolesOrgProductEnablement struct {
	// Product Name (NVAIE, BASE_COMMAND, REGISTRY, etc)
	ProductName string `json:"productName,required"`
	// Product Enablement Types
	Type UserListUsersRolesOrgProductEnablementsType `json:"type,required"`
	// Date on which the subscription expires. The subscription is invalid after this
	// date. (yyyy-MM-dd)
	ExpirationDate string                                            `json:"expirationDate"`
	PoDetails      []UserListUsersRolesOrgProductEnablementsPoDetail `json:"poDetails"`
	JSON           userListUsersRolesOrgProductEnablementJSON        `json:"-"`
}

// userListUsersRolesOrgProductEnablementJSON contains the JSON metadata for the
// struct [UserListUsersRolesOrgProductEnablement]
type userListUsersRolesOrgProductEnablementJSON struct {
	ProductName    apijson.Field
	Type           apijson.Field
	ExpirationDate apijson.Field
	PoDetails      apijson.Field
	raw            string
	ExtraFields    map[string]apijson.Field
}

func (r *UserListUsersRolesOrgProductEnablement) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userListUsersRolesOrgProductEnablementJSON) RawJSON() string {
	return r.raw
}

// Product Enablement Types
type UserListUsersRolesOrgProductEnablementsType string

const (
	UserListUsersRolesOrgProductEnablementsTypeNgcAdminEval       UserListUsersRolesOrgProductEnablementsType = "NGC_ADMIN_EVAL"
	UserListUsersRolesOrgProductEnablementsTypeNgcAdminNfr        UserListUsersRolesOrgProductEnablementsType = "NGC_ADMIN_NFR"
	UserListUsersRolesOrgProductEnablementsTypeNgcAdminCommercial UserListUsersRolesOrgProductEnablementsType = "NGC_ADMIN_COMMERCIAL"
	UserListUsersRolesOrgProductEnablementsTypeEmsEval            UserListUsersRolesOrgProductEnablementsType = "EMS_EVAL"
	UserListUsersRolesOrgProductEnablementsTypeEmsNfr             UserListUsersRolesOrgProductEnablementsType = "EMS_NFR"
	UserListUsersRolesOrgProductEnablementsTypeEmsCommercial      UserListUsersRolesOrgProductEnablementsType = "EMS_COMMERCIAL"
	UserListUsersRolesOrgProductEnablementsTypeNgcAdminDeveloper  UserListUsersRolesOrgProductEnablementsType = "NGC_ADMIN_DEVELOPER"
)

func (r UserListUsersRolesOrgProductEnablementsType) IsKnown() bool {
	switch r {
	case UserListUsersRolesOrgProductEnablementsTypeNgcAdminEval, UserListUsersRolesOrgProductEnablementsTypeNgcAdminNfr, UserListUsersRolesOrgProductEnablementsTypeNgcAdminCommercial, UserListUsersRolesOrgProductEnablementsTypeEmsEval, UserListUsersRolesOrgProductEnablementsTypeEmsNfr, UserListUsersRolesOrgProductEnablementsTypeEmsCommercial, UserListUsersRolesOrgProductEnablementsTypeNgcAdminDeveloper:
		return true
	}
	return false
}

// Purchase Order.
type UserListUsersRolesOrgProductEnablementsPoDetail struct {
	// Entitlement identifier.
	EntitlementID string `json:"entitlementId"`
	// PAK (Product Activation Key) identifier.
	PkID string                                              `json:"pkId"`
	JSON userListUsersRolesOrgProductEnablementsPoDetailJSON `json:"-"`
}

// userListUsersRolesOrgProductEnablementsPoDetailJSON contains the JSON metadata
// for the struct [UserListUsersRolesOrgProductEnablementsPoDetail]
type userListUsersRolesOrgProductEnablementsPoDetailJSON struct {
	EntitlementID apijson.Field
	PkID          apijson.Field
	raw           string
	ExtraFields   map[string]apijson.Field
}

func (r *UserListUsersRolesOrgProductEnablementsPoDetail) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userListUsersRolesOrgProductEnablementsPoDetailJSON) RawJSON() string {
	return r.raw
}

// Product Subscription
type UserListUsersRolesOrgProductSubscription struct {
	// Product Name (NVAIE, BASE_COMMAND, FleetCommand, REGISTRY, etc).
	ProductName string `json:"productName,required"`
	// Unique entitlement identifier
	ID string `json:"id"`
	// EMS Subscription type. (options: EMS_EVAL, EMS_NFR and EMS_COMMERCIAL)
	EmsEntitlementType UserListUsersRolesOrgProductSubscriptionsEmsEntitlementType `json:"emsEntitlementType"`
	// Date on which the subscription expires. The subscription is invalid after this
	// date. (yyyy-MM-dd)
	ExpirationDate string `json:"expirationDate"`
	// Date on which the subscription becomes active. (yyyy-MM-dd)
	StartDate string `json:"startDate"`
	// Subscription type. (options: NGC_ADMIN_EVAL, NGC_ADMIN_NFR,
	// NGC_ADMIN_COMMERCIAL)
	Type UserListUsersRolesOrgProductSubscriptionsType `json:"type"`
	JSON userListUsersRolesOrgProductSubscriptionJSON  `json:"-"`
}

// userListUsersRolesOrgProductSubscriptionJSON contains the JSON metadata for the
// struct [UserListUsersRolesOrgProductSubscription]
type userListUsersRolesOrgProductSubscriptionJSON struct {
	ProductName        apijson.Field
	ID                 apijson.Field
	EmsEntitlementType apijson.Field
	ExpirationDate     apijson.Field
	StartDate          apijson.Field
	Type               apijson.Field
	raw                string
	ExtraFields        map[string]apijson.Field
}

func (r *UserListUsersRolesOrgProductSubscription) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userListUsersRolesOrgProductSubscriptionJSON) RawJSON() string {
	return r.raw
}

// EMS Subscription type. (options: EMS_EVAL, EMS_NFR and EMS_COMMERCIAL)
type UserListUsersRolesOrgProductSubscriptionsEmsEntitlementType string

const (
	UserListUsersRolesOrgProductSubscriptionsEmsEntitlementTypeEmsEval       UserListUsersRolesOrgProductSubscriptionsEmsEntitlementType = "EMS_EVAL"
	UserListUsersRolesOrgProductSubscriptionsEmsEntitlementTypeEmsNfr        UserListUsersRolesOrgProductSubscriptionsEmsEntitlementType = "EMS_NFR"
	UserListUsersRolesOrgProductSubscriptionsEmsEntitlementTypeEmsCommerical UserListUsersRolesOrgProductSubscriptionsEmsEntitlementType = "EMS_COMMERICAL"
	UserListUsersRolesOrgProductSubscriptionsEmsEntitlementTypeEmsCommercial UserListUsersRolesOrgProductSubscriptionsEmsEntitlementType = "EMS_COMMERCIAL"
)

func (r UserListUsersRolesOrgProductSubscriptionsEmsEntitlementType) IsKnown() bool {
	switch r {
	case UserListUsersRolesOrgProductSubscriptionsEmsEntitlementTypeEmsEval, UserListUsersRolesOrgProductSubscriptionsEmsEntitlementTypeEmsNfr, UserListUsersRolesOrgProductSubscriptionsEmsEntitlementTypeEmsCommerical, UserListUsersRolesOrgProductSubscriptionsEmsEntitlementTypeEmsCommercial:
		return true
	}
	return false
}

// Subscription type. (options: NGC_ADMIN_EVAL, NGC_ADMIN_NFR,
// NGC_ADMIN_COMMERCIAL)
type UserListUsersRolesOrgProductSubscriptionsType string

const (
	UserListUsersRolesOrgProductSubscriptionsTypeNgcAdminEval       UserListUsersRolesOrgProductSubscriptionsType = "NGC_ADMIN_EVAL"
	UserListUsersRolesOrgProductSubscriptionsTypeNgcAdminNfr        UserListUsersRolesOrgProductSubscriptionsType = "NGC_ADMIN_NFR"
	UserListUsersRolesOrgProductSubscriptionsTypeNgcAdminCommercial UserListUsersRolesOrgProductSubscriptionsType = "NGC_ADMIN_COMMERCIAL"
)

func (r UserListUsersRolesOrgProductSubscriptionsType) IsKnown() bool {
	switch r {
	case UserListUsersRolesOrgProductSubscriptionsTypeNgcAdminEval, UserListUsersRolesOrgProductSubscriptionsTypeNgcAdminNfr, UserListUsersRolesOrgProductSubscriptionsTypeNgcAdminCommercial:
		return true
	}
	return false
}

// Repo scan setting definition
type UserListUsersRolesOrgRepoScanSettings struct {
	// Allow org admin to override the org level repo scan settings
	RepoScanAllowOverride bool `json:"repoScanAllowOverride"`
	// Allow repository scanning by default
	RepoScanByDefault bool `json:"repoScanByDefault"`
	// Enable the repository scan or not. Only used in org level object
	RepoScanEnabled bool `json:"repoScanEnabled"`
	// Sends notification to end user after scanning is done
	RepoScanEnableNotifications bool `json:"repoScanEnableNotifications"`
	// Allow override settings at team level. Only used in org level object
	RepoScanEnableTeamOverride bool `json:"repoScanEnableTeamOverride"`
	// Allow showing scan results to CLI or UI
	RepoScanShowResults bool                                      `json:"repoScanShowResults"`
	JSON                userListUsersRolesOrgRepoScanSettingsJSON `json:"-"`
}

// userListUsersRolesOrgRepoScanSettingsJSON contains the JSON metadata for the
// struct [UserListUsersRolesOrgRepoScanSettings]
type userListUsersRolesOrgRepoScanSettingsJSON struct {
	RepoScanAllowOverride       apijson.Field
	RepoScanByDefault           apijson.Field
	RepoScanEnabled             apijson.Field
	RepoScanEnableNotifications apijson.Field
	RepoScanEnableTeamOverride  apijson.Field
	RepoScanShowResults         apijson.Field
	raw                         string
	ExtraFields                 map[string]apijson.Field
}

func (r *UserListUsersRolesOrgRepoScanSettings) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userListUsersRolesOrgRepoScanSettingsJSON) RawJSON() string {
	return r.raw
}

type UserListUsersRolesOrgType string

const (
	UserListUsersRolesOrgTypeUnknown    UserListUsersRolesOrgType = "UNKNOWN"
	UserListUsersRolesOrgTypeCloud      UserListUsersRolesOrgType = "CLOUD"
	UserListUsersRolesOrgTypeEnterprise UserListUsersRolesOrgType = "ENTERPRISE"
	UserListUsersRolesOrgTypeIndividual UserListUsersRolesOrgType = "INDIVIDUAL"
)

func (r UserListUsersRolesOrgType) IsKnown() bool {
	switch r {
	case UserListUsersRolesOrgTypeUnknown, UserListUsersRolesOrgTypeCloud, UserListUsersRolesOrgTypeEnterprise, UserListUsersRolesOrgTypeIndividual:
		return true
	}
	return false
}

// Users information.
type UserListUsersRolesOrgUsersInfo struct {
	// Total number of users.
	TotalUsers int64                              `json:"totalUsers"`
	JSON       userListUsersRolesOrgUsersInfoJSON `json:"-"`
}

// userListUsersRolesOrgUsersInfoJSON contains the JSON metadata for the struct
// [UserListUsersRolesOrgUsersInfo]
type userListUsersRolesOrgUsersInfoJSON struct {
	TotalUsers  apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserListUsersRolesOrgUsersInfo) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userListUsersRolesOrgUsersInfoJSON) RawJSON() string {
	return r.raw
}

// Information about the user who is attempting to run the job
type UserListUsersRolesTargetSystemUserIdentifier struct {
	// gid of the user on this team
	Gid int64 `json:"gid"`
	// Org context for the job
	OrgName string `json:"orgName"`
	// Starfleet ID of the user creating the job.
	StarfleetID string `json:"starfleetId"`
	// Team context for the job
	TeamName string `json:"teamName"`
	// uid of the user on this team
	Uid int64 `json:"uid"`
	// Unique ID of the user who submitted the job
	UserID int64                                            `json:"userId"`
	JSON   userListUsersRolesTargetSystemUserIdentifierJSON `json:"-"`
}

// userListUsersRolesTargetSystemUserIdentifierJSON contains the JSON metadata for
// the struct [UserListUsersRolesTargetSystemUserIdentifier]
type userListUsersRolesTargetSystemUserIdentifierJSON struct {
	Gid         apijson.Field
	OrgName     apijson.Field
	StarfleetID apijson.Field
	TeamName    apijson.Field
	Uid         apijson.Field
	UserID      apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserListUsersRolesTargetSystemUserIdentifier) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userListUsersRolesTargetSystemUserIdentifierJSON) RawJSON() string {
	return r.raw
}

// represents user storage quota for a given ace and available unused storage
type UserListUsersStorageQuota struct {
	// id of the ace
	AceID int64 `json:"aceId"`
	// name of the ace
	AceName string `json:"aceName"`
	// Available space in bytes
	Available int64 `json:"available"`
	// Number of datasets that are a part of user's used storage
	DatasetCount int64 `json:"datasetCount"`
	// Space used by datasets in bytes
	DatasetsUsage int64 `json:"datasetsUsage"`
	// The org name that this user quota tied to. This is needed for analytics
	OrgName string `json:"orgName"`
	// Assigned quota in bytes
	Quota int64 `json:"quota"`
	// Number of resultsets that are a part of user's used storage
	ResultsetCount int64 `json:"resultsetCount"`
	// Space used by resultsets in bytes
	ResultsetsUsage int64 `json:"resultsetsUsage"`
	// Description of this storage cluster
	StorageClusterDescription string `json:"storageClusterDescription"`
	// Name of storage cluster
	StorageClusterName string `json:"storageClusterName"`
	// Identifier to this storage cluster
	StorageClusterUuid string `json:"storageClusterUuid"`
	// Number of workspaces that are a part of user's used storage
	WorkspacesCount int64 `json:"workspacesCount"`
	// Space used by workspaces in bytes
	WorkspacesUsage int64                         `json:"workspacesUsage"`
	JSON            userListUsersStorageQuotaJSON `json:"-"`
}

// userListUsersStorageQuotaJSON contains the JSON metadata for the struct
// [UserListUsersStorageQuota]
type userListUsersStorageQuotaJSON struct {
	AceID                     apijson.Field
	AceName                   apijson.Field
	Available                 apijson.Field
	DatasetCount              apijson.Field
	DatasetsUsage             apijson.Field
	OrgName                   apijson.Field
	Quota                     apijson.Field
	ResultsetCount            apijson.Field
	ResultsetsUsage           apijson.Field
	StorageClusterDescription apijson.Field
	StorageClusterName        apijson.Field
	StorageClusterUuid        apijson.Field
	WorkspacesCount           apijson.Field
	WorkspacesUsage           apijson.Field
	raw                       string
	ExtraFields               map[string]apijson.Field
}

func (r *UserListUsersStorageQuota) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userListUsersStorageQuotaJSON) RawJSON() string {
	return r.raw
}

// Metadata information about the user.
type UserListUsersUserMetadata struct {
	// Name of the company
	Company string `json:"company"`
	// Company URL
	CompanyURL string `json:"companyUrl"`
	// Country of the user
	Country string `json:"country"`
	// User's first name
	FirstName string `json:"firstName"`
	// Industry segment
	Industry string `json:"industry"`
	// List of development areas that user has interest
	Interest []string `json:"interest"`
	// User's last name
	LastName string `json:"lastName"`
	// Role of the user in the company
	Role string                        `json:"role"`
	JSON userListUsersUserMetadataJSON `json:"-"`
}

// userListUsersUserMetadataJSON contains the JSON metadata for the struct
// [UserListUsersUserMetadata]
type userListUsersUserMetadataJSON struct {
	Company     apijson.Field
	CompanyURL  apijson.Field
	Country     apijson.Field
	FirstName   apijson.Field
	Industry    apijson.Field
	Interest    apijson.Field
	LastName    apijson.Field
	Role        apijson.Field
	raw         string
	ExtraFields map[string]apijson.Field
}

func (r *UserListUsersUserMetadata) UnmarshalJSON(data []byte) (err error) {
	return apijson.UnmarshalRoot(data, r)
}

func (r userListUsersUserMetadataJSON) RawJSON() string {
	return r.raw
}
