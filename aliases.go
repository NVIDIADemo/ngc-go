// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package ngc

import (
	"github.com/NVIDIADemo/ngc-go/internal/apierror"
	"github.com/NVIDIADemo/ngc-go/shared"
)

type Error = apierror.Error

// This API is invoked by monitoring tools, other services and infrastructure to
// retrieve health status the targeted service, this is unprotected method
//
// This is an alias to an internal type.
type Health = shared.Health

// object that describes health of the service
//
// This is an alias to an internal type.
type HealthHealth = shared.HealthHealth

// Enum that describes health of the service
//
// This is an alias to an internal type.
type HealthHealthHealthCode = shared.HealthHealthHealthCode

// This is an alias to an internal value.
const HealthHealthHealthCodeUnknown = shared.HealthHealthHealthCodeUnknown

// This is an alias to an internal value.
const HealthHealthHealthCodeOk = shared.HealthHealthHealthCodeOk

// This is an alias to an internal value.
const HealthHealthHealthCodeUnderMaintenance = shared.HealthHealthHealthCodeUnderMaintenance

// This is an alias to an internal value.
const HealthHealthHealthCodeWarning = shared.HealthHealthHealthCodeWarning

// This is an alias to an internal value.
const HealthHealthHealthCodeFailed = shared.HealthHealthHealthCodeFailed

// This is an alias to an internal type.
type HealthHealthMetaData = shared.HealthHealthMetaData

// This is an alias to an internal type.
type HealthRequestStatus = shared.HealthRequestStatus

// Describes response status reported by the server.
//
// This is an alias to an internal type.
type HealthRequestStatusStatusCode = shared.HealthRequestStatusStatusCode

// This is an alias to an internal value.
const HealthRequestStatusStatusCodeUnknown = shared.HealthRequestStatusStatusCodeUnknown

// This is an alias to an internal value.
const HealthRequestStatusStatusCodeSuccess = shared.HealthRequestStatusStatusCodeSuccess

// This is an alias to an internal value.
const HealthRequestStatusStatusCodeUnauthorized = shared.HealthRequestStatusStatusCodeUnauthorized

// This is an alias to an internal value.
const HealthRequestStatusStatusCodePaymentRequired = shared.HealthRequestStatusStatusCodePaymentRequired

// This is an alias to an internal value.
const HealthRequestStatusStatusCodeForbidden = shared.HealthRequestStatusStatusCodeForbidden

// This is an alias to an internal value.
const HealthRequestStatusStatusCodeTimeout = shared.HealthRequestStatusStatusCodeTimeout

// This is an alias to an internal value.
const HealthRequestStatusStatusCodeExists = shared.HealthRequestStatusStatusCodeExists

// This is an alias to an internal value.
const HealthRequestStatusStatusCodeNotFound = shared.HealthRequestStatusStatusCodeNotFound

// This is an alias to an internal value.
const HealthRequestStatusStatusCodeInternalError = shared.HealthRequestStatusStatusCodeInternalError

// This is an alias to an internal value.
const HealthRequestStatusStatusCodeInvalidRequest = shared.HealthRequestStatusStatusCodeInvalidRequest

// This is an alias to an internal value.
const HealthRequestStatusStatusCodeInvalidRequestVersion = shared.HealthRequestStatusStatusCodeInvalidRequestVersion

// This is an alias to an internal value.
const HealthRequestStatusStatusCodeInvalidRequestData = shared.HealthRequestStatusStatusCodeInvalidRequestData

// This is an alias to an internal value.
const HealthRequestStatusStatusCodeMethodNotAllowed = shared.HealthRequestStatusStatusCodeMethodNotAllowed

// This is an alias to an internal value.
const HealthRequestStatusStatusCodeConflict = shared.HealthRequestStatusStatusCodeConflict

// This is an alias to an internal value.
const HealthRequestStatusStatusCodeUnprocessableEntity = shared.HealthRequestStatusStatusCodeUnprocessableEntity

// This is an alias to an internal value.
const HealthRequestStatusStatusCodeTooManyRequests = shared.HealthRequestStatusStatusCodeTooManyRequests

// This is an alias to an internal value.
const HealthRequestStatusStatusCodeInsufficientStorage = shared.HealthRequestStatusStatusCodeInsufficientStorage

// This is an alias to an internal value.
const HealthRequestStatusStatusCodeServiceUnavailable = shared.HealthRequestStatusStatusCodeServiceUnavailable

// This is an alias to an internal value.
const HealthRequestStatusStatusCodePayloadTooLarge = shared.HealthRequestStatusStatusCodePayloadTooLarge

// This is an alias to an internal value.
const HealthRequestStatusStatusCodeNotAcceptable = shared.HealthRequestStatusStatusCodeNotAcceptable

// This is an alias to an internal value.
const HealthRequestStatusStatusCodeUnavailableForLegalReasons = shared.HealthRequestStatusStatusCodeUnavailableForLegalReasons

// This is an alias to an internal value.
const HealthRequestStatusStatusCodeBadGateway = shared.HealthRequestStatusStatusCodeBadGateway

// response containing a list of all metering queries results
//
// This is an alias to an internal type.
type MeteringResultList = shared.MeteringResultList

// result of a single measurement query
//
// This is an alias to an internal type.
type MeteringResultListMeasurement = shared.MeteringResultListMeasurement

// object for a single series in the measurement
//
// This is an alias to an internal type.
type MeteringResultListMeasurementsSery = shared.MeteringResultListMeasurementsSery

// object for measurement tags which identifies a measuurement series
//
// This is an alias to an internal type.
type MeteringResultListMeasurementsSeriesTag = shared.MeteringResultListMeasurementsSeriesTag

// object for the measurement values
//
// This is an alias to an internal type.
type MeteringResultListMeasurementsSeriesValue = shared.MeteringResultListMeasurementsSeriesValue

// This is an alias to an internal type.
type MeteringResultListRequestStatus = shared.MeteringResultListRequestStatus

// Describes response status reported by the server.
//
// This is an alias to an internal type.
type MeteringResultListRequestStatusStatusCode = shared.MeteringResultListRequestStatusStatusCode

// This is an alias to an internal value.
const MeteringResultListRequestStatusStatusCodeUnknown = shared.MeteringResultListRequestStatusStatusCodeUnknown

// This is an alias to an internal value.
const MeteringResultListRequestStatusStatusCodeSuccess = shared.MeteringResultListRequestStatusStatusCodeSuccess

// This is an alias to an internal value.
const MeteringResultListRequestStatusStatusCodeUnauthorized = shared.MeteringResultListRequestStatusStatusCodeUnauthorized

// This is an alias to an internal value.
const MeteringResultListRequestStatusStatusCodePaymentRequired = shared.MeteringResultListRequestStatusStatusCodePaymentRequired

// This is an alias to an internal value.
const MeteringResultListRequestStatusStatusCodeForbidden = shared.MeteringResultListRequestStatusStatusCodeForbidden

// This is an alias to an internal value.
const MeteringResultListRequestStatusStatusCodeTimeout = shared.MeteringResultListRequestStatusStatusCodeTimeout

// This is an alias to an internal value.
const MeteringResultListRequestStatusStatusCodeExists = shared.MeteringResultListRequestStatusStatusCodeExists

// This is an alias to an internal value.
const MeteringResultListRequestStatusStatusCodeNotFound = shared.MeteringResultListRequestStatusStatusCodeNotFound

// This is an alias to an internal value.
const MeteringResultListRequestStatusStatusCodeInternalError = shared.MeteringResultListRequestStatusStatusCodeInternalError

// This is an alias to an internal value.
const MeteringResultListRequestStatusStatusCodeInvalidRequest = shared.MeteringResultListRequestStatusStatusCodeInvalidRequest

// This is an alias to an internal value.
const MeteringResultListRequestStatusStatusCodeInvalidRequestVersion = shared.MeteringResultListRequestStatusStatusCodeInvalidRequestVersion

// This is an alias to an internal value.
const MeteringResultListRequestStatusStatusCodeInvalidRequestData = shared.MeteringResultListRequestStatusStatusCodeInvalidRequestData

// This is an alias to an internal value.
const MeteringResultListRequestStatusStatusCodeMethodNotAllowed = shared.MeteringResultListRequestStatusStatusCodeMethodNotAllowed

// This is an alias to an internal value.
const MeteringResultListRequestStatusStatusCodeConflict = shared.MeteringResultListRequestStatusStatusCodeConflict

// This is an alias to an internal value.
const MeteringResultListRequestStatusStatusCodeUnprocessableEntity = shared.MeteringResultListRequestStatusStatusCodeUnprocessableEntity

// This is an alias to an internal value.
const MeteringResultListRequestStatusStatusCodeTooManyRequests = shared.MeteringResultListRequestStatusStatusCodeTooManyRequests

// This is an alias to an internal value.
const MeteringResultListRequestStatusStatusCodeInsufficientStorage = shared.MeteringResultListRequestStatusStatusCodeInsufficientStorage

// This is an alias to an internal value.
const MeteringResultListRequestStatusStatusCodeServiceUnavailable = shared.MeteringResultListRequestStatusStatusCodeServiceUnavailable

// This is an alias to an internal value.
const MeteringResultListRequestStatusStatusCodePayloadTooLarge = shared.MeteringResultListRequestStatusStatusCodePayloadTooLarge

// This is an alias to an internal value.
const MeteringResultListRequestStatusStatusCodeNotAcceptable = shared.MeteringResultListRequestStatusStatusCodeNotAcceptable

// This is an alias to an internal value.
const MeteringResultListRequestStatusStatusCodeUnavailableForLegalReasons = shared.MeteringResultListRequestStatusStatusCodeUnavailableForLegalReasons

// This is an alias to an internal value.
const MeteringResultListRequestStatusStatusCodeBadGateway = shared.MeteringResultListRequestStatusStatusCodeBadGateway

// listing of all teams
//
// This is an alias to an internal type.
type TeamList = shared.TeamList

// object that describes the pagination information
//
// This is an alias to an internal type.
type TeamListPaginationInfo = shared.TeamListPaginationInfo

// This is an alias to an internal type.
type TeamListRequestStatus = shared.TeamListRequestStatus

// Describes response status reported by the server.
//
// This is an alias to an internal type.
type TeamListRequestStatusStatusCode = shared.TeamListRequestStatusStatusCode

// This is an alias to an internal value.
const TeamListRequestStatusStatusCodeUnknown = shared.TeamListRequestStatusStatusCodeUnknown

// This is an alias to an internal value.
const TeamListRequestStatusStatusCodeSuccess = shared.TeamListRequestStatusStatusCodeSuccess

// This is an alias to an internal value.
const TeamListRequestStatusStatusCodeUnauthorized = shared.TeamListRequestStatusStatusCodeUnauthorized

// This is an alias to an internal value.
const TeamListRequestStatusStatusCodePaymentRequired = shared.TeamListRequestStatusStatusCodePaymentRequired

// This is an alias to an internal value.
const TeamListRequestStatusStatusCodeForbidden = shared.TeamListRequestStatusStatusCodeForbidden

// This is an alias to an internal value.
const TeamListRequestStatusStatusCodeTimeout = shared.TeamListRequestStatusStatusCodeTimeout

// This is an alias to an internal value.
const TeamListRequestStatusStatusCodeExists = shared.TeamListRequestStatusStatusCodeExists

// This is an alias to an internal value.
const TeamListRequestStatusStatusCodeNotFound = shared.TeamListRequestStatusStatusCodeNotFound

// This is an alias to an internal value.
const TeamListRequestStatusStatusCodeInternalError = shared.TeamListRequestStatusStatusCodeInternalError

// This is an alias to an internal value.
const TeamListRequestStatusStatusCodeInvalidRequest = shared.TeamListRequestStatusStatusCodeInvalidRequest

// This is an alias to an internal value.
const TeamListRequestStatusStatusCodeInvalidRequestVersion = shared.TeamListRequestStatusStatusCodeInvalidRequestVersion

// This is an alias to an internal value.
const TeamListRequestStatusStatusCodeInvalidRequestData = shared.TeamListRequestStatusStatusCodeInvalidRequestData

// This is an alias to an internal value.
const TeamListRequestStatusStatusCodeMethodNotAllowed = shared.TeamListRequestStatusStatusCodeMethodNotAllowed

// This is an alias to an internal value.
const TeamListRequestStatusStatusCodeConflict = shared.TeamListRequestStatusStatusCodeConflict

// This is an alias to an internal value.
const TeamListRequestStatusStatusCodeUnprocessableEntity = shared.TeamListRequestStatusStatusCodeUnprocessableEntity

// This is an alias to an internal value.
const TeamListRequestStatusStatusCodeTooManyRequests = shared.TeamListRequestStatusStatusCodeTooManyRequests

// This is an alias to an internal value.
const TeamListRequestStatusStatusCodeInsufficientStorage = shared.TeamListRequestStatusStatusCodeInsufficientStorage

// This is an alias to an internal value.
const TeamListRequestStatusStatusCodeServiceUnavailable = shared.TeamListRequestStatusStatusCodeServiceUnavailable

// This is an alias to an internal value.
const TeamListRequestStatusStatusCodePayloadTooLarge = shared.TeamListRequestStatusStatusCodePayloadTooLarge

// This is an alias to an internal value.
const TeamListRequestStatusStatusCodeNotAcceptable = shared.TeamListRequestStatusStatusCodeNotAcceptable

// This is an alias to an internal value.
const TeamListRequestStatusStatusCodeUnavailableForLegalReasons = shared.TeamListRequestStatusStatusCodeUnavailableForLegalReasons

// This is an alias to an internal value.
const TeamListRequestStatusStatusCodeBadGateway = shared.TeamListRequestStatusStatusCodeBadGateway

// Information about the team
//
// This is an alias to an internal type.
type TeamListTeam = shared.TeamListTeam

// Infinity manager setting definition
//
// This is an alias to an internal type.
type TeamListTeamsInfinityManagerSettings = shared.TeamListTeamsInfinityManagerSettings

// Repo scan setting definition
//
// This is an alias to an internal type.
type TeamListTeamsRepoScanSettings = shared.TeamListTeamsRepoScanSettings

// about one user
//
// This is an alias to an internal type.
type User = shared.User

// NCA role
//
// This is an alias to an internal type.
type UserNcaRole = shared.UserNcaRole

// This is an alias to an internal value.
const UserNcaRoleUnknown = shared.UserNcaRoleUnknown

// This is an alias to an internal value.
const UserNcaRoleAdministrator = shared.UserNcaRoleAdministrator

// This is an alias to an internal value.
const UserNcaRoleMember = shared.UserNcaRoleMember

// This is an alias to an internal value.
const UserNcaRoleOwner = shared.UserNcaRoleOwner

// This is an alias to an internal value.
const UserNcaRolePending = shared.UserNcaRolePending

// This is an alias to an internal type.
type UserRequestStatus = shared.UserRequestStatus

// Describes response status reported by the server.
//
// This is an alias to an internal type.
type UserRequestStatusStatusCode = shared.UserRequestStatusStatusCode

// This is an alias to an internal value.
const UserRequestStatusStatusCodeUnknown = shared.UserRequestStatusStatusCodeUnknown

// This is an alias to an internal value.
const UserRequestStatusStatusCodeSuccess = shared.UserRequestStatusStatusCodeSuccess

// This is an alias to an internal value.
const UserRequestStatusStatusCodeUnauthorized = shared.UserRequestStatusStatusCodeUnauthorized

// This is an alias to an internal value.
const UserRequestStatusStatusCodePaymentRequired = shared.UserRequestStatusStatusCodePaymentRequired

// This is an alias to an internal value.
const UserRequestStatusStatusCodeForbidden = shared.UserRequestStatusStatusCodeForbidden

// This is an alias to an internal value.
const UserRequestStatusStatusCodeTimeout = shared.UserRequestStatusStatusCodeTimeout

// This is an alias to an internal value.
const UserRequestStatusStatusCodeExists = shared.UserRequestStatusStatusCodeExists

// This is an alias to an internal value.
const UserRequestStatusStatusCodeNotFound = shared.UserRequestStatusStatusCodeNotFound

// This is an alias to an internal value.
const UserRequestStatusStatusCodeInternalError = shared.UserRequestStatusStatusCodeInternalError

// This is an alias to an internal value.
const UserRequestStatusStatusCodeInvalidRequest = shared.UserRequestStatusStatusCodeInvalidRequest

// This is an alias to an internal value.
const UserRequestStatusStatusCodeInvalidRequestVersion = shared.UserRequestStatusStatusCodeInvalidRequestVersion

// This is an alias to an internal value.
const UserRequestStatusStatusCodeInvalidRequestData = shared.UserRequestStatusStatusCodeInvalidRequestData

// This is an alias to an internal value.
const UserRequestStatusStatusCodeMethodNotAllowed = shared.UserRequestStatusStatusCodeMethodNotAllowed

// This is an alias to an internal value.
const UserRequestStatusStatusCodeConflict = shared.UserRequestStatusStatusCodeConflict

// This is an alias to an internal value.
const UserRequestStatusStatusCodeUnprocessableEntity = shared.UserRequestStatusStatusCodeUnprocessableEntity

// This is an alias to an internal value.
const UserRequestStatusStatusCodeTooManyRequests = shared.UserRequestStatusStatusCodeTooManyRequests

// This is an alias to an internal value.
const UserRequestStatusStatusCodeInsufficientStorage = shared.UserRequestStatusStatusCodeInsufficientStorage

// This is an alias to an internal value.
const UserRequestStatusStatusCodeServiceUnavailable = shared.UserRequestStatusStatusCodeServiceUnavailable

// This is an alias to an internal value.
const UserRequestStatusStatusCodePayloadTooLarge = shared.UserRequestStatusStatusCodePayloadTooLarge

// This is an alias to an internal value.
const UserRequestStatusStatusCodeNotAcceptable = shared.UserRequestStatusStatusCodeNotAcceptable

// This is an alias to an internal value.
const UserRequestStatusStatusCodeUnavailableForLegalReasons = shared.UserRequestStatusStatusCodeUnavailableForLegalReasons

// This is an alias to an internal value.
const UserRequestStatusStatusCodeBadGateway = shared.UserRequestStatusStatusCodeBadGateway

// information about the user
//
// This is an alias to an internal type.
type UserUser = shared.UserUser

// Type of IDP, Identity Provider. Used for login.
//
// This is an alias to an internal type.
type UserUserIdpType = shared.UserUserIdpType

// This is an alias to an internal value.
const UserUserIdpTypeNvidia = shared.UserUserIdpTypeNvidia

// This is an alias to an internal value.
const UserUserIdpTypeEnterprise = shared.UserUserIdpTypeEnterprise

// List of roles that the user have
//
// This is an alias to an internal type.
type UserUserRole = shared.UserUserRole

// Information about the Organization
//
// This is an alias to an internal type.
type UserUserRolesOrg = shared.UserUserRolesOrg

// Org Owner Alternate Contact
//
// This is an alias to an internal type.
type UserUserRolesOrgAlternateContact = shared.UserUserRolesOrgAlternateContact

// Infinity manager setting definition
//
// This is an alias to an internal type.
type UserUserRolesOrgInfinityManagerSettings = shared.UserUserRolesOrgInfinityManagerSettings

// Org owner.
//
// This is an alias to an internal type.
type UserUserRolesOrgOrgOwner = shared.UserUserRolesOrgOrgOwner

// Product Enablement
//
// This is an alias to an internal type.
type UserUserRolesOrgProductEnablement = shared.UserUserRolesOrgProductEnablement

// Product Enablement Types
//
// This is an alias to an internal type.
type UserUserRolesOrgProductEnablementsType = shared.UserUserRolesOrgProductEnablementsType

// This is an alias to an internal value.
const UserUserRolesOrgProductEnablementsTypeNgcAdminEval = shared.UserUserRolesOrgProductEnablementsTypeNgcAdminEval

// This is an alias to an internal value.
const UserUserRolesOrgProductEnablementsTypeNgcAdminNfr = shared.UserUserRolesOrgProductEnablementsTypeNgcAdminNfr

// This is an alias to an internal value.
const UserUserRolesOrgProductEnablementsTypeNgcAdminCommercial = shared.UserUserRolesOrgProductEnablementsTypeNgcAdminCommercial

// This is an alias to an internal value.
const UserUserRolesOrgProductEnablementsTypeEmsEval = shared.UserUserRolesOrgProductEnablementsTypeEmsEval

// This is an alias to an internal value.
const UserUserRolesOrgProductEnablementsTypeEmsNfr = shared.UserUserRolesOrgProductEnablementsTypeEmsNfr

// This is an alias to an internal value.
const UserUserRolesOrgProductEnablementsTypeEmsCommercial = shared.UserUserRolesOrgProductEnablementsTypeEmsCommercial

// This is an alias to an internal value.
const UserUserRolesOrgProductEnablementsTypeNgcAdminDeveloper = shared.UserUserRolesOrgProductEnablementsTypeNgcAdminDeveloper

// Purchase Order.
//
// This is an alias to an internal type.
type UserUserRolesOrgProductEnablementsPoDetail = shared.UserUserRolesOrgProductEnablementsPoDetail

// Product Subscription
//
// This is an alias to an internal type.
type UserUserRolesOrgProductSubscription = shared.UserUserRolesOrgProductSubscription

// EMS Subscription type. (options: EMS_EVAL, EMS_NFR and EMS_COMMERCIAL)
//
// This is an alias to an internal type.
type UserUserRolesOrgProductSubscriptionsEmsEntitlementType = shared.UserUserRolesOrgProductSubscriptionsEmsEntitlementType

// This is an alias to an internal value.
const UserUserRolesOrgProductSubscriptionsEmsEntitlementTypeEmsEval = shared.UserUserRolesOrgProductSubscriptionsEmsEntitlementTypeEmsEval

// This is an alias to an internal value.
const UserUserRolesOrgProductSubscriptionsEmsEntitlementTypeEmsNfr = shared.UserUserRolesOrgProductSubscriptionsEmsEntitlementTypeEmsNfr

// This is an alias to an internal value.
const UserUserRolesOrgProductSubscriptionsEmsEntitlementTypeEmsCommerical = shared.UserUserRolesOrgProductSubscriptionsEmsEntitlementTypeEmsCommerical

// This is an alias to an internal value.
const UserUserRolesOrgProductSubscriptionsEmsEntitlementTypeEmsCommercial = shared.UserUserRolesOrgProductSubscriptionsEmsEntitlementTypeEmsCommercial

// Subscription type. (options: NGC_ADMIN_EVAL, NGC_ADMIN_NFR,
// NGC_ADMIN_COMMERCIAL)
//
// This is an alias to an internal type.
type UserUserRolesOrgProductSubscriptionsType = shared.UserUserRolesOrgProductSubscriptionsType

// This is an alias to an internal value.
const UserUserRolesOrgProductSubscriptionsTypeNgcAdminEval = shared.UserUserRolesOrgProductSubscriptionsTypeNgcAdminEval

// This is an alias to an internal value.
const UserUserRolesOrgProductSubscriptionsTypeNgcAdminNfr = shared.UserUserRolesOrgProductSubscriptionsTypeNgcAdminNfr

// This is an alias to an internal value.
const UserUserRolesOrgProductSubscriptionsTypeNgcAdminCommercial = shared.UserUserRolesOrgProductSubscriptionsTypeNgcAdminCommercial

// Repo scan setting definition
//
// This is an alias to an internal type.
type UserUserRolesOrgRepoScanSettings = shared.UserUserRolesOrgRepoScanSettings

// This is an alias to an internal type.
type UserUserRolesOrgType = shared.UserUserRolesOrgType

// This is an alias to an internal value.
const UserUserRolesOrgTypeUnknown = shared.UserUserRolesOrgTypeUnknown

// This is an alias to an internal value.
const UserUserRolesOrgTypeCloud = shared.UserUserRolesOrgTypeCloud

// This is an alias to an internal value.
const UserUserRolesOrgTypeEnterprise = shared.UserUserRolesOrgTypeEnterprise

// This is an alias to an internal value.
const UserUserRolesOrgTypeIndividual = shared.UserUserRolesOrgTypeIndividual

// Users information.
//
// This is an alias to an internal type.
type UserUserRolesOrgUsersInfo = shared.UserUserRolesOrgUsersInfo

// Information about the user who is attempting to run the job
//
// This is an alias to an internal type.
type UserUserRolesTargetSystemUserIdentifier = shared.UserUserRolesTargetSystemUserIdentifier

// Information about the team
//
// This is an alias to an internal type.
type UserUserRolesTeam = shared.UserUserRolesTeam

// Infinity manager setting definition
//
// This is an alias to an internal type.
type UserUserRolesTeamInfinityManagerSettings = shared.UserUserRolesTeamInfinityManagerSettings

// Repo scan setting definition
//
// This is an alias to an internal type.
type UserUserRolesTeamRepoScanSettings = shared.UserUserRolesTeamRepoScanSettings

// represents user storage quota for a given ace and available unused storage
//
// This is an alias to an internal type.
type UserUserStorageQuota = shared.UserUserStorageQuota

// Metadata information about the user.
//
// This is an alias to an internal type.
type UserUserUserMetadata = shared.UserUserUserMetadata

// Response for List User reponse
//
// This is an alias to an internal type.
type UserList = shared.UserList

// object that describes the pagination information
//
// This is an alias to an internal type.
type UserListPaginationInfo = shared.UserListPaginationInfo

// This is an alias to an internal type.
type UserListRequestStatus = shared.UserListRequestStatus

// Describes response status reported by the server.
//
// This is an alias to an internal type.
type UserListRequestStatusStatusCode = shared.UserListRequestStatusStatusCode

// This is an alias to an internal value.
const UserListRequestStatusStatusCodeUnknown = shared.UserListRequestStatusStatusCodeUnknown

// This is an alias to an internal value.
const UserListRequestStatusStatusCodeSuccess = shared.UserListRequestStatusStatusCodeSuccess

// This is an alias to an internal value.
const UserListRequestStatusStatusCodeUnauthorized = shared.UserListRequestStatusStatusCodeUnauthorized

// This is an alias to an internal value.
const UserListRequestStatusStatusCodePaymentRequired = shared.UserListRequestStatusStatusCodePaymentRequired

// This is an alias to an internal value.
const UserListRequestStatusStatusCodeForbidden = shared.UserListRequestStatusStatusCodeForbidden

// This is an alias to an internal value.
const UserListRequestStatusStatusCodeTimeout = shared.UserListRequestStatusStatusCodeTimeout

// This is an alias to an internal value.
const UserListRequestStatusStatusCodeExists = shared.UserListRequestStatusStatusCodeExists

// This is an alias to an internal value.
const UserListRequestStatusStatusCodeNotFound = shared.UserListRequestStatusStatusCodeNotFound

// This is an alias to an internal value.
const UserListRequestStatusStatusCodeInternalError = shared.UserListRequestStatusStatusCodeInternalError

// This is an alias to an internal value.
const UserListRequestStatusStatusCodeInvalidRequest = shared.UserListRequestStatusStatusCodeInvalidRequest

// This is an alias to an internal value.
const UserListRequestStatusStatusCodeInvalidRequestVersion = shared.UserListRequestStatusStatusCodeInvalidRequestVersion

// This is an alias to an internal value.
const UserListRequestStatusStatusCodeInvalidRequestData = shared.UserListRequestStatusStatusCodeInvalidRequestData

// This is an alias to an internal value.
const UserListRequestStatusStatusCodeMethodNotAllowed = shared.UserListRequestStatusStatusCodeMethodNotAllowed

// This is an alias to an internal value.
const UserListRequestStatusStatusCodeConflict = shared.UserListRequestStatusStatusCodeConflict

// This is an alias to an internal value.
const UserListRequestStatusStatusCodeUnprocessableEntity = shared.UserListRequestStatusStatusCodeUnprocessableEntity

// This is an alias to an internal value.
const UserListRequestStatusStatusCodeTooManyRequests = shared.UserListRequestStatusStatusCodeTooManyRequests

// This is an alias to an internal value.
const UserListRequestStatusStatusCodeInsufficientStorage = shared.UserListRequestStatusStatusCodeInsufficientStorage

// This is an alias to an internal value.
const UserListRequestStatusStatusCodeServiceUnavailable = shared.UserListRequestStatusStatusCodeServiceUnavailable

// This is an alias to an internal value.
const UserListRequestStatusStatusCodePayloadTooLarge = shared.UserListRequestStatusStatusCodePayloadTooLarge

// This is an alias to an internal value.
const UserListRequestStatusStatusCodeNotAcceptable = shared.UserListRequestStatusStatusCodeNotAcceptable

// This is an alias to an internal value.
const UserListRequestStatusStatusCodeUnavailableForLegalReasons = shared.UserListRequestStatusStatusCodeUnavailableForLegalReasons

// This is an alias to an internal value.
const UserListRequestStatusStatusCodeBadGateway = shared.UserListRequestStatusStatusCodeBadGateway

// information about the user
//
// This is an alias to an internal type.
type UserListUser = shared.UserListUser

// Type of IDP, Identity Provider. Used for login.
//
// This is an alias to an internal type.
type UserListUsersIdpType = shared.UserListUsersIdpType

// This is an alias to an internal value.
const UserListUsersIdpTypeNvidia = shared.UserListUsersIdpTypeNvidia

// This is an alias to an internal value.
const UserListUsersIdpTypeEnterprise = shared.UserListUsersIdpTypeEnterprise

// List of roles that the user have
//
// This is an alias to an internal type.
type UserListUsersRole = shared.UserListUsersRole

// Information about the Organization
//
// This is an alias to an internal type.
type UserListUsersRolesOrg = shared.UserListUsersRolesOrg

// Org Owner Alternate Contact
//
// This is an alias to an internal type.
type UserListUsersRolesOrgAlternateContact = shared.UserListUsersRolesOrgAlternateContact

// Infinity manager setting definition
//
// This is an alias to an internal type.
type UserListUsersRolesOrgInfinityManagerSettings = shared.UserListUsersRolesOrgInfinityManagerSettings

// Org owner.
//
// This is an alias to an internal type.
type UserListUsersRolesOrgOrgOwner = shared.UserListUsersRolesOrgOrgOwner

// Product Enablement
//
// This is an alias to an internal type.
type UserListUsersRolesOrgProductEnablement = shared.UserListUsersRolesOrgProductEnablement

// Product Enablement Types
//
// This is an alias to an internal type.
type UserListUsersRolesOrgProductEnablementsType = shared.UserListUsersRolesOrgProductEnablementsType

// This is an alias to an internal value.
const UserListUsersRolesOrgProductEnablementsTypeNgcAdminEval = shared.UserListUsersRolesOrgProductEnablementsTypeNgcAdminEval

// This is an alias to an internal value.
const UserListUsersRolesOrgProductEnablementsTypeNgcAdminNfr = shared.UserListUsersRolesOrgProductEnablementsTypeNgcAdminNfr

// This is an alias to an internal value.
const UserListUsersRolesOrgProductEnablementsTypeNgcAdminCommercial = shared.UserListUsersRolesOrgProductEnablementsTypeNgcAdminCommercial

// This is an alias to an internal value.
const UserListUsersRolesOrgProductEnablementsTypeEmsEval = shared.UserListUsersRolesOrgProductEnablementsTypeEmsEval

// This is an alias to an internal value.
const UserListUsersRolesOrgProductEnablementsTypeEmsNfr = shared.UserListUsersRolesOrgProductEnablementsTypeEmsNfr

// This is an alias to an internal value.
const UserListUsersRolesOrgProductEnablementsTypeEmsCommercial = shared.UserListUsersRolesOrgProductEnablementsTypeEmsCommercial

// This is an alias to an internal value.
const UserListUsersRolesOrgProductEnablementsTypeNgcAdminDeveloper = shared.UserListUsersRolesOrgProductEnablementsTypeNgcAdminDeveloper

// Purchase Order.
//
// This is an alias to an internal type.
type UserListUsersRolesOrgProductEnablementsPoDetail = shared.UserListUsersRolesOrgProductEnablementsPoDetail

// Product Subscription
//
// This is an alias to an internal type.
type UserListUsersRolesOrgProductSubscription = shared.UserListUsersRolesOrgProductSubscription

// EMS Subscription type. (options: EMS_EVAL, EMS_NFR and EMS_COMMERCIAL)
//
// This is an alias to an internal type.
type UserListUsersRolesOrgProductSubscriptionsEmsEntitlementType = shared.UserListUsersRolesOrgProductSubscriptionsEmsEntitlementType

// This is an alias to an internal value.
const UserListUsersRolesOrgProductSubscriptionsEmsEntitlementTypeEmsEval = shared.UserListUsersRolesOrgProductSubscriptionsEmsEntitlementTypeEmsEval

// This is an alias to an internal value.
const UserListUsersRolesOrgProductSubscriptionsEmsEntitlementTypeEmsNfr = shared.UserListUsersRolesOrgProductSubscriptionsEmsEntitlementTypeEmsNfr

// This is an alias to an internal value.
const UserListUsersRolesOrgProductSubscriptionsEmsEntitlementTypeEmsCommerical = shared.UserListUsersRolesOrgProductSubscriptionsEmsEntitlementTypeEmsCommerical

// This is an alias to an internal value.
const UserListUsersRolesOrgProductSubscriptionsEmsEntitlementTypeEmsCommercial = shared.UserListUsersRolesOrgProductSubscriptionsEmsEntitlementTypeEmsCommercial

// Subscription type. (options: NGC_ADMIN_EVAL, NGC_ADMIN_NFR,
// NGC_ADMIN_COMMERCIAL)
//
// This is an alias to an internal type.
type UserListUsersRolesOrgProductSubscriptionsType = shared.UserListUsersRolesOrgProductSubscriptionsType

// This is an alias to an internal value.
const UserListUsersRolesOrgProductSubscriptionsTypeNgcAdminEval = shared.UserListUsersRolesOrgProductSubscriptionsTypeNgcAdminEval

// This is an alias to an internal value.
const UserListUsersRolesOrgProductSubscriptionsTypeNgcAdminNfr = shared.UserListUsersRolesOrgProductSubscriptionsTypeNgcAdminNfr

// This is an alias to an internal value.
const UserListUsersRolesOrgProductSubscriptionsTypeNgcAdminCommercial = shared.UserListUsersRolesOrgProductSubscriptionsTypeNgcAdminCommercial

// Repo scan setting definition
//
// This is an alias to an internal type.
type UserListUsersRolesOrgRepoScanSettings = shared.UserListUsersRolesOrgRepoScanSettings

// This is an alias to an internal type.
type UserListUsersRolesOrgType = shared.UserListUsersRolesOrgType

// This is an alias to an internal value.
const UserListUsersRolesOrgTypeUnknown = shared.UserListUsersRolesOrgTypeUnknown

// This is an alias to an internal value.
const UserListUsersRolesOrgTypeCloud = shared.UserListUsersRolesOrgTypeCloud

// This is an alias to an internal value.
const UserListUsersRolesOrgTypeEnterprise = shared.UserListUsersRolesOrgTypeEnterprise

// This is an alias to an internal value.
const UserListUsersRolesOrgTypeIndividual = shared.UserListUsersRolesOrgTypeIndividual

// Users information.
//
// This is an alias to an internal type.
type UserListUsersRolesOrgUsersInfo = shared.UserListUsersRolesOrgUsersInfo

// Information about the user who is attempting to run the job
//
// This is an alias to an internal type.
type UserListUsersRolesTargetSystemUserIdentifier = shared.UserListUsersRolesTargetSystemUserIdentifier

// Information about the team
//
// This is an alias to an internal type.
type UserListUsersRolesTeam = shared.UserListUsersRolesTeam

// Infinity manager setting definition
//
// This is an alias to an internal type.
type UserListUsersRolesTeamInfinityManagerSettings = shared.UserListUsersRolesTeamInfinityManagerSettings

// Repo scan setting definition
//
// This is an alias to an internal type.
type UserListUsersRolesTeamRepoScanSettings = shared.UserListUsersRolesTeamRepoScanSettings

// represents user storage quota for a given ace and available unused storage
//
// This is an alias to an internal type.
type UserListUsersStorageQuota = shared.UserListUsersStorageQuota

// Metadata information about the user.
//
// This is an alias to an internal type.
type UserListUsersUserMetadata = shared.UserListUsersUserMetadata
