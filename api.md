# Shared Response Types

- <a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go/shared#Health">Health</a>
- <a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go/shared#MeteringResultList">MeteringResultList</a>
- <a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go/shared#TeamList">TeamList</a>
- <a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go/shared#User">User</a>
- <a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go/shared#UserInvitationList">UserInvitationList</a>
- <a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go/shared#UserList">UserList</a>

# Orgs

Response Types:

- <a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go">ngc</a>.<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go#OrgList">OrgList</a>
- <a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go">ngc</a>.<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go#OrgResponse">OrgResponse</a>
- <a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go">ngc</a>.<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go#OrgListResponse">OrgListResponse</a>

Methods:

- <code title="post /v3/orgs">client.Orgs.<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go#OrgService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, params <a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go">ngc</a>.<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go#OrgNewParams">OrgNewParams</a>) (<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go">ngc</a>.<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go#OrgResponse">OrgResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v2/orgs/{org-name}">client.Orgs.<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go#OrgService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, orgName <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go">ngc</a>.<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go#OrgResponse">OrgResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="patch /v2/orgs/{org-name}">client.Orgs.<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go#OrgService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, orgName <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go">ngc</a>.<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go#OrgUpdateParams">OrgUpdateParams</a>) (<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go">ngc</a>.<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go#OrgResponse">OrgResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v2/orgs">client.Orgs.<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go#OrgService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go">ngc</a>.<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go#OrgListParams">OrgListParams</a>) (<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go/internal/pagination">pagination</a>.<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go/internal/pagination#PageNumberOrganizations">PageNumberOrganizations</a>[<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go">ngc</a>.<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go#OrgListResponse">OrgListResponse</a>], <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Users

Response Types:

- <a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go">ngc</a>.<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go#OrgUserListResponse">OrgUserListResponse</a>
- <a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go">ngc</a>.<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go#OrgUserDeleteResponse">OrgUserDeleteResponse</a>

Methods:

- <code title="post /v2/org/{org-name}/users">client.Orgs.Users.<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go#OrgUserService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, orgName <a href="https://pkg.go.dev/builtin#string">string</a>, params <a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go">ngc</a>.<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go#OrgUserNewParams">OrgUserNewParams</a>) (<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go/shared#User">User</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v2/org/{org-name}/users">client.Orgs.Users.<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go#OrgUserService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, orgName <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go">ngc</a>.<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go#OrgUserListParams">OrgUserListParams</a>) (<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go/internal/pagination">pagination</a>.<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go/internal/pagination#PageNumberUsers">PageNumberUsers</a>[<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go/shared#OrgUserListResponse">OrgUserListResponse</a>], <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /v3/orgs/{org-name}/users/{id}">client.Orgs.Users.<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go#OrgUserService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, orgName <a href="https://pkg.go.dev/builtin#string">string</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go">ngc</a>.<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go#OrgUserDeleteParams">OrgUserDeleteParams</a>) (<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go">ngc</a>.<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go#OrgUserDeleteResponse">OrgUserDeleteResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="patch /v3/orgs/{org-name}/users/{user-email-or-id}/add-role">client.Orgs.Users.<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go#OrgUserService.AddRole">AddRole</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, orgName <a href="https://pkg.go.dev/builtin#string">string</a>, userEmailOrID <a href="https://pkg.go.dev/builtin#string">string</a>, params <a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go">ngc</a>.<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go#OrgUserAddRoleParams">OrgUserAddRoleParams</a>) (<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go/shared#User">User</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /v3/orgs/{org-name}/users/{user-email-or-id}/remove-role">client.Orgs.Users.<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go#OrgUserService.RemoveRole">RemoveRole</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, orgName <a href="https://pkg.go.dev/builtin#string">string</a>, userEmailOrID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go">ngc</a>.<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go#OrgUserRemoveRoleParams">OrgUserRemoveRoleParams</a>) (<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go/shared#User">User</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="patch /v2/org/{org-name}/users/{id}/update-role">client.Orgs.Users.<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go#OrgUserService.UpdateRole">UpdateRole</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, orgName <a href="https://pkg.go.dev/builtin#string">string</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go">ngc</a>.<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go#OrgUserUpdateRoleParams">OrgUserUpdateRoleParams</a>) (<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go/shared#User">User</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### NcaInvitations

Methods:

- <code title="post /v3/orgs/{org-name}/users/nca-invitations">client.Orgs.Users.NcaInvitations.<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go#OrgUserNcaInvitationService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, orgName <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go">ngc</a>.<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go#OrgUserNcaInvitationNewParams">OrgUserNcaInvitationNewParams</a>) (<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go/shared#User">User</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Teams

Response Types:

- <a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go">ngc</a>.<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go#TeamResponse">TeamResponse</a>
- <a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go">ngc</a>.<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go#OrgTeamListResponse">OrgTeamListResponse</a>

Methods:

- <code title="get /v2/org/{org-name}/teams">client.Orgs.Teams.<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go#OrgTeamService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, orgName <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go">ngc</a>.<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go#OrgTeamListParams">OrgTeamListParams</a>) (<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go/internal/pagination">pagination</a>.<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go/internal/pagination#PageNumberTeams">PageNumberTeams</a>[<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go/shared#OrgTeamListResponse">OrgTeamListResponse</a>], <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### Users

Response Types:

- <a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go">ngc</a>.<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go#OrgTeamUserDeleteResponse">OrgTeamUserDeleteResponse</a>

Methods:

- <code title="delete /v3/orgs/{org-name}/teams/{team-name}/users/{id}">client.Orgs.Teams.Users.<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go#OrgTeamUserService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, orgName <a href="https://pkg.go.dev/builtin#string">string</a>, teamName <a href="https://pkg.go.dev/builtin#string">string</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go">ngc</a>.<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go#OrgTeamUserDeleteParams">OrgTeamUserDeleteParams</a>) (<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go">ngc</a>.<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go#OrgTeamUserDeleteResponse">OrgTeamUserDeleteResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="patch /v3/orgs/{org-name}/teams/{team-name}/users/{user-email-or-id}/add-role">client.Orgs.Teams.Users.<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go#OrgTeamUserService.AddRole">AddRole</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, orgName <a href="https://pkg.go.dev/builtin#string">string</a>, teamName <a href="https://pkg.go.dev/builtin#string">string</a>, userEmailOrID <a href="https://pkg.go.dev/builtin#string">string</a>, params <a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go">ngc</a>.<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go#OrgTeamUserAddRoleParams">OrgTeamUserAddRoleParams</a>) (<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go/shared#User">User</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /v3/orgs/{org-name}/teams/{team-name}/users/{user-email-or-id}/remove-role">client.Orgs.Teams.Users.<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go#OrgTeamUserService.RemoveRole">RemoveRole</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, orgName <a href="https://pkg.go.dev/builtin#string">string</a>, teamName <a href="https://pkg.go.dev/builtin#string">string</a>, userEmailOrID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go">ngc</a>.<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go#OrgTeamUserRemoveRoleParams">OrgTeamUserRemoveRoleParams</a>) (<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go/shared#User">User</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### StarfleetIDs

Methods:

- <code title="get /v2/org/{org-name}/team/{team-name}/starfleetIds/{starfleet-id}">client.Orgs.Teams.StarfleetIDs.<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go#OrgTeamStarfleetIDService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, orgName <a href="https://pkg.go.dev/builtin#string">string</a>, teamName <a href="https://pkg.go.dev/builtin#string">string</a>, starfleetID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go/shared#User">User</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### NcaInvitations

Methods:

- <code title="post /v3/orgs/{org-name}/teams/{team-name}/users/nca-invitations">client.Orgs.Teams.NcaInvitations.<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go#OrgTeamNcaInvitationService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, orgName <a href="https://pkg.go.dev/builtin#string">string</a>, teamName <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go">ngc</a>.<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go#OrgTeamNcaInvitationNewParams">OrgTeamNcaInvitationNewParams</a>) (<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go/shared#User">User</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## ProtoOrg

Methods:

- <code title="post /v3/orgs/proto-org">client.Orgs.ProtoOrg.<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go#OrgProtoOrgService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, params <a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go">ngc</a>.<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go#OrgProtoOrgNewParams">OrgProtoOrgNewParams</a>) (<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go">ngc</a>.<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go#OrgResponse">OrgResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Credits

Response Types:

- <a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go">ngc</a>.<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go#CreditsHistory">CreditsHistory</a>

Methods:

- <code title="get /v2/orgs/{org-name}/credits">client.Orgs.Credits.<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go#OrgCreditService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, orgName <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go">ngc</a>.<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go#CreditsHistory">CreditsHistory</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## StarfleetIDs

Methods:

- <code title="get /v2/org/{org-name}/starfleetIds/{starfleet-id}">client.Orgs.StarfleetIDs.<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go#OrgStarfleetIDService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, orgName <a href="https://pkg.go.dev/builtin#string">string</a>, starfleetID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go/shared#User">User</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Metering

Methods:

- <code title="get /v2/org/{org-name}/metering">client.Orgs.Metering.<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go#OrgMeteringService.GetAll">GetAll</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, orgName <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go">ngc</a>.<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go#OrgMeteringGetAllParams">OrgMeteringGetAllParams</a>) (<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go/shared#MeteringResultList">MeteringResultList</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### Gpupeak

Methods:

- <code title="get /v2/org/{org-name}/meterings/gpupeak">client.Orgs.Metering.Gpupeak.<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go#OrgMeteringGpupeakService.GetAll">GetAll</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, orgName <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go">ngc</a>.<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go#OrgMeteringGpupeakGetAllParams">OrgMeteringGpupeakGetAllParams</a>) (<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go/shared#MeteringResultList">MeteringResultList</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## AuditLogs

Response Types:

- <a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go">ngc</a>.<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go#AuditLogsPresignedURL">AuditLogsPresignedURL</a>

Methods:

- <code title="get /v2/org/{org-name}/auditLogs/{log-id}">client.Orgs.AuditLogs.<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go#OrgAuditLogService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, orgName <a href="https://pkg.go.dev/builtin#string">string</a>, logID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go">ngc</a>.<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go#AuditLogsPresignedURL">AuditLogsPresignedURL</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Admin

## Orgs

Response Types:

- <a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go">ngc</a>.<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go#AdminOrgListResponse">AdminOrgListResponse</a>

Methods:

- <code title="post /v3/admin/orgs">client.Admin.Orgs.<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go#AdminOrgService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, params <a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go">ngc</a>.<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go#AdminOrgNewParams">AdminOrgNewParams</a>) (http.Response, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v2/admin/org">client.Admin.Orgs.<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go#AdminOrgService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go">ngc</a>.<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go#AdminOrgListParams">AdminOrgListParams</a>) (<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go/internal/pagination">pagination</a>.<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go/internal/pagination#PageNumberOrganizations">PageNumberOrganizations</a>[<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go">ngc</a>.<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go#AdminOrgListResponse">AdminOrgListResponse</a>], <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### NcaIDs

Methods:

- <code title="post /v3/admin/orgs/ncaIds">client.Admin.Orgs.NcaIDs.<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go#AdminOrgNcaIDService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go">ngc</a>.<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go#AdminOrgNcaIDNewParams">AdminOrgNcaIDNewParams</a>) (http.Response, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### Offboarded

Response Types:

- <a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go">ngc</a>.<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go#AdminOrgOffboardedListResponse">AdminOrgOffboardedListResponse</a>

Methods:

- <code title="get /v2/admin/orgs/offboarded">client.Admin.Orgs.Offboarded.<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go#AdminOrgOffboardedService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go">ngc</a>.<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go#AdminOrgOffboardedListParams">AdminOrgOffboardedListParams</a>) (<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go/internal/pagination">pagination</a>.<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go/internal/pagination#PageNumberOrganizations">PageNumberOrganizations</a>[<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go">ngc</a>.<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go#AdminOrgOffboardedListResponse">AdminOrgOffboardedListResponse</a>], <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### Teams

Response Types:

- <a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go">ngc</a>.<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go#AdminOrgTeamListResponse">AdminOrgTeamListResponse</a>

Methods:

- <code title="get /v2/admin/org/{org-name}/teams">client.Admin.Orgs.Teams.<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go#AdminOrgTeamService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, orgName <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go">ngc</a>.<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go#AdminOrgTeamListParams">AdminOrgTeamListParams</a>) (<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go/internal/pagination">pagination</a>.<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go/internal/pagination#PageNumberTeams">PageNumberTeams</a>[<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go/shared#AdminOrgTeamListResponse">AdminOrgTeamListResponse</a>], <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### Users

## Users

Response Types:

- <a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go">ngc</a>.<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go#AdminUserOrgOwnerBackfillResponse">AdminUserOrgOwnerBackfillResponse</a>

Methods:

- <code title="get /v2/admin/users/{id}">client.Admin.Users.<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go#AdminUserService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go/shared#User">User</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="patch /v2/admin/users/invite">client.Admin.Users.<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go#AdminUserService.Invite">Invite</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go">ngc</a>.<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go#AdminUserInviteParams">AdminUserInviteParams</a>) (<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go/shared#User">User</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v2/admin/users/me">client.Admin.Users.<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go#AdminUserService.Me">Me</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go">ngc</a>.<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go#AdminUserMeParams">AdminUserMeParams</a>) (<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go/shared#User">User</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /v2/admin/users/{user-id}/org-owner-backfill">client.Admin.Users.<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go#AdminUserService.OrgOwnerBackfill">OrgOwnerBackfill</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, userID <a href="https://pkg.go.dev/builtin#int64">int64</a>) (<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go">ngc</a>.<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go#AdminUserOrgOwnerBackfillResponse">AdminUserOrgOwnerBackfillResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Org

Response Types:

- <a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go">ngc</a>.<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go#AdminOrgListResponse">AdminOrgListResponse</a>

Methods:

- <code title="get /v2/admin/org/{org-name}">client.Admin.Org.<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go#AdminOrgService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, orgName <a href="https://pkg.go.dev/builtin#string">string</a>) (http.Response, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="patch /v2/admin/org/{org-name}">client.Admin.Org.<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go#AdminOrgService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, orgName <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go">ngc</a>.<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go#AdminOrgUpdateParams">AdminOrgUpdateParams</a>) (http.Response, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v3/admin/org">client.Admin.Org.<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go#AdminOrgService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go">ngc</a>.<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go#AdminOrgListParams">AdminOrgListParams</a>) (<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go/internal/pagination">pagination</a>.<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go/internal/pagination#PageNumberOrganizations">PageNumberOrganizations</a>[<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go">ngc</a>.<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go#AdminOrgListResponse">AdminOrgListResponse</a>], <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### Users

Methods:

- <code title="delete /v2/admin/org/{org-name}/users/{id}/remove-role">client.Admin.Org.Users.<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go#AdminOrgUserService.RemoveRole">RemoveRole</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, orgName <a href="https://pkg.go.dev/builtin#string">string</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go">ngc</a>.<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go#AdminOrgUserRemoveRoleParams">AdminOrgUserRemoveRoleParams</a>) (<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go/shared#User">User</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="patch /v2/admin/org/{org-name}/users/{id}/update-role">client.Admin.Org.Users.<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go#AdminOrgUserService.UpdateRole">UpdateRole</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, orgName <a href="https://pkg.go.dev/builtin#string">string</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go">ngc</a>.<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go#AdminOrgUserUpdateRoleParams">AdminOrgUserUpdateRoleParams</a>) (<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go/shared#User">User</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### Registry

#### Metering

Methods:

- <code title="get /v2/admin/org/{org-name}/registry/metering/downsample">client.Admin.Org.Registry.Metering.<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go#AdminOrgRegistryMeteringService.Downsample">Downsample</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, orgName <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go">ngc</a>.<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go#AdminOrgRegistryMeteringDownsampleParams">AdminOrgRegistryMeteringDownsampleParams</a>) (<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go/shared#MeteringResultList">MeteringResultList</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### Validate

Methods:

- <code title="get /v2/admin/org/validate">client.Admin.Org.Validate.<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go#AdminOrgValidateService.GetAll">GetAll</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go">ngc</a>.<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go#AdminOrgValidateGetAllParams">AdminOrgValidateGetAllParams</a>) (<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go">ngc</a>.<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go#OrgList">OrgList</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### Team

#### Users

Methods:

- <code title="delete /v2/admin/org/{org-name}/team/{team-name}/users/{id}/remove-role">client.Admin.Org.Team.Users.<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go#AdminOrgTeamUserService.RemoveRole">RemoveRole</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, orgName <a href="https://pkg.go.dev/builtin#string">string</a>, teamName <a href="https://pkg.go.dev/builtin#string">string</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go">ngc</a>.<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go#AdminOrgTeamUserRemoveRoleParams">AdminOrgTeamUserRemoveRoleParams</a>) (<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go/shared#User">User</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Users

## V2

### APIKey

Response Types:

- <a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go">ngc</a>.<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go#UserKeyResponse">UserKeyResponse</a>

Methods:

- <code title="post /v2/users/me/api-key">client.Users.V2.APIKey.<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go#UserV2APIKeyService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go">ngc</a>.<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go#UserKeyResponse">UserKeyResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Organizations

## Teams

Response Types:

- <a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go">ngc</a>.<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go#TeamCreateResponse">TeamCreateResponse</a>

Methods:

- <code title="post /v2/org/{org-name}/teams">client.Organizations.Teams.<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go#OrganizationTeamService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, orgName <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go">ngc</a>.<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go#OrganizationTeamNewParams">OrganizationTeamNewParams</a>) (<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go">ngc</a>.<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go#TeamCreateResponse">TeamCreateResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### Users

Response Types:

- <a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go">ngc</a>.<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go#OrganizationTeamUserListResponse">OrganizationTeamUserListResponse</a>
- <a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go">ngc</a>.<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go#OrganizationTeamUserDeleteResponse">OrganizationTeamUserDeleteResponse</a>

Methods:

- <code title="post /v2/org/{org-name}/team/{team-name}/users">client.Organizations.Teams.Users.<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go#OrganizationTeamUserService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, orgName <a href="https://pkg.go.dev/builtin#string">string</a>, teamName <a href="https://pkg.go.dev/builtin#string">string</a>, params <a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go">ngc</a>.<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go#OrganizationTeamUserNewParams">OrganizationTeamUserNewParams</a>) (<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go/shared#User">User</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v2/org/{org-name}/team/{team-name}/users/{id}">client.Organizations.Teams.Users.<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go#OrganizationTeamUserService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, orgName <a href="https://pkg.go.dev/builtin#string">string</a>, teamName <a href="https://pkg.go.dev/builtin#string">string</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go">ngc</a>.<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go#OrganizationTeamUserGetParams">OrganizationTeamUserGetParams</a>) (<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go/shared#User">User</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v2/org/{org-name}/team/{team-name}/users">client.Organizations.Teams.Users.<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go#OrganizationTeamUserService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, orgName <a href="https://pkg.go.dev/builtin#string">string</a>, teamName <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go">ngc</a>.<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go#OrganizationTeamUserListParams">OrganizationTeamUserListParams</a>) (<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go/internal/pagination">pagination</a>.<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go/internal/pagination#PageNumberUsers">PageNumberUsers</a>[<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go/shared#OrganizationTeamUserListResponse">OrganizationTeamUserListResponse</a>], <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /v2/org/{org-name}/team/{team-name}/users/{id}">client.Organizations.Teams.Users.<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go#OrganizationTeamUserService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, orgName <a href="https://pkg.go.dev/builtin#string">string</a>, teamName <a href="https://pkg.go.dev/builtin#string">string</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go">ngc</a>.<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go#OrganizationTeamUserDeleteResponse">OrganizationTeamUserDeleteResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## AuditLogs

Response Types:

- <a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go">ngc</a>.<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go#AuditLogsResponse">AuditLogsResponse</a>
- <a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go">ngc</a>.<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go#OrganizationAuditLogDeleteResponse">OrganizationAuditLogDeleteResponse</a>
- <a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go">ngc</a>.<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go#OrganizationAuditLogRequestResponse">OrganizationAuditLogRequestResponse</a>

Methods:

- <code title="delete /v2/org/{org-name}/auditLogs">client.Organizations.AuditLogs.<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go#OrganizationAuditLogService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, orgName <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go">ngc</a>.<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go#OrganizationAuditLogDeleteParams">OrganizationAuditLogDeleteParams</a>) (<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go">ngc</a>.<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go#OrganizationAuditLogDeleteResponse">OrganizationAuditLogDeleteResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /v2/org/{org-name}/auditLogs">client.Organizations.AuditLogs.<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go#OrganizationAuditLogService.Request">Request</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, orgName <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go">ngc</a>.<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go#OrganizationAuditLogRequestParams">OrganizationAuditLogRequestParams</a>) (<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go">ngc</a>.<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go#OrganizationAuditLogRequestResponse">OrganizationAuditLogRequestResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v2/org/{org-name}/auditLogs">client.Organizations.AuditLogs.<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go#OrganizationAuditLogService.GetAll">GetAll</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, orgName <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go">ngc</a>.<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go#AuditLogsResponse">AuditLogsResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# SuperAdminUser

Response Types:

- <a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go">ngc</a>.<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go#SuperAdminUserCRMSyncResponse">SuperAdminUserCRMSyncResponse</a>

Methods:

- <code title="post /v2/admin/users/{id}/crm-sync">client.SuperAdminUser.<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go#SuperAdminUserService.CRMSync">CRMSync</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#int64">int64</a>) (<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go">ngc</a>.<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go#SuperAdminUserCRMSyncResponse">SuperAdminUserCRMSyncResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /v2/admin/users/{id}/migrate-deprecated-roles">client.SuperAdminUser.<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go#SuperAdminUserService.MigrateDeprecatedRoles">MigrateDeprecatedRoles</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go/shared#User">User</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Orgs

Response Types:

- <a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go">ngc</a>.<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go#SuperAdminUserOrgOrgOwnerBackfillResponse">SuperAdminUserOrgOrgOwnerBackfillResponse</a>

Methods:

- <code title="post /v2/admin/org/{org-name}/org-owner-backfill">client.SuperAdminUser.Orgs.<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go#SuperAdminUserOrgService.OrgOwnerBackfill">OrgOwnerBackfill</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, orgName <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go">ngc</a>.<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go#SuperAdminUserOrgOrgOwnerBackfillResponse">SuperAdminUserOrgOrgOwnerBackfillResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### Users

Response Types:

- <a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go">ngc</a>.<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go#SuperAdminUserOrgUserRemoveResponse">SuperAdminUserOrgUserRemoveResponse</a>

Methods:

- <code title="post /v2/admin/org/{org-name}/users">client.SuperAdminUser.Orgs.Users.<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go#SuperAdminUserOrgUserService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, orgName <a href="https://pkg.go.dev/builtin#string">string</a>, params <a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go">ngc</a>.<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go#SuperAdminUserOrgUserNewParams">SuperAdminUserOrgUserNewParams</a>) (<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go/shared#User">User</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /v2/admin/org/{org-name}/users/{id}">client.SuperAdminUser.Orgs.Users.<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go#SuperAdminUserOrgUserService.Add">Add</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, orgName <a href="https://pkg.go.dev/builtin#string">string</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, params <a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go">ngc</a>.<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go#SuperAdminUserOrgUserAddParams">SuperAdminUserOrgUserAddParams</a>) (<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go/shared#User">User</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /v2/admin/org/{org-name}/users/{id}">client.SuperAdminUser.Orgs.Users.<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go#SuperAdminUserOrgUserService.Remove">Remove</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, orgName <a href="https://pkg.go.dev/builtin#string">string</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go">ngc</a>.<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go#SuperAdminUserOrgUserRemoveResponse">SuperAdminUserOrgUserRemoveResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### Teams

#### Users

Methods:

- <code title="post /v2/admin/org/{org-name}/team/{team-name}/users">client.SuperAdminUser.Orgs.Teams.Users.<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go#SuperAdminUserOrgTeamUserService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, orgName <a href="https://pkg.go.dev/builtin#string">string</a>, teamName <a href="https://pkg.go.dev/builtin#string">string</a>, params <a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go">ngc</a>.<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go#SuperAdminUserOrgTeamUserNewParams">SuperAdminUserOrgTeamUserNewParams</a>) (<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go/shared#User">User</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /v2/admin/org/{org-name}/team/{team-name}/users/{id}">client.SuperAdminUser.Orgs.Teams.Users.<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go#SuperAdminUserOrgTeamUserService.Add">Add</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, orgName <a href="https://pkg.go.dev/builtin#string">string</a>, teamName <a href="https://pkg.go.dev/builtin#string">string</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, params <a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go">ngc</a>.<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go#SuperAdminUserOrgTeamUserAddParams">SuperAdminUserOrgTeamUserAddParams</a>) (<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go/shared#User">User</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# SuperAdminOrg

Methods:

- <code title="post /v2/admin/orgs">client.SuperAdminOrg.<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go#SuperAdminOrgService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, params <a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go">ngc</a>.<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go#SuperAdminOrgNewParams">SuperAdminOrgNewParams</a>) (http.Response, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## OrgStatus

Methods:

- <code title="post /v2/admin/org/{org-name}/enablement">client.SuperAdminOrg.OrgStatus.<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go#SuperAdminOrgOrgStatusService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, orgName <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go">ngc</a>.<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go#SuperAdminOrgOrgStatusNewParams">SuperAdminOrgOrgStatusNewParams</a>) (http.Response, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# SuperAdminOrgControllers

Methods:

- <code title="post /v2/admin/backfill-orgs-to-kratos">client.SuperAdminOrgControllers.<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go#SuperAdminOrgControllerService.BackfillOrgsToKratos">BackfillOrgsToKratos</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (http.Response, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Org

Methods:

- <code title="get /v3/admin/org/{org-name}">client.SuperAdminOrgControllers.Org.<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go#SuperAdminOrgControllerOrgService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, orgName <a href="https://pkg.go.dev/builtin#string">string</a>) (http.Response, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="patch /v3/admin/org/{org-name}">client.SuperAdminOrgControllers.Org.<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go#SuperAdminOrgControllerOrgService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, orgName <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go">ngc</a>.<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go#SuperAdminOrgControllerOrgUpdateParams">SuperAdminOrgControllerOrgUpdateParams</a>) (http.Response, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# UsersManagement

## Me

Methods:

- <code title="get /v2/users/me">client.UsersManagement.Me.<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go#UsersManagementMeService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go">ngc</a>.<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go#UsersManagementMeGetParams">UsersManagementMeGetParams</a>) (<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go/shared#User">User</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="patch /v2/users/me">client.UsersManagement.Me.<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go#UsersManagementMeService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go">ngc</a>.<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go#UsersManagementMeUpdateParams">UsersManagementMeUpdateParams</a>) (<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go/shared#User">User</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Org

## Users

Response Types:

- <a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go">ngc</a>.<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go#OrgUserRemoveResponse">OrgUserRemoveResponse</a>

Methods:

- <code title="get /v2/org/{org-name}/users/{id}">client.Org.Users.<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go#OrgUserService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, orgName <a href="https://pkg.go.dev/builtin#string">string</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go">ngc</a>.<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go#OrgUserGetParams">OrgUserGetParams</a>) (<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go/shared#User">User</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="patch /v2/org/{org-name}/users/{id}/add-role">client.Org.Users.<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go#OrgUserService.AddRole">AddRole</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, orgName <a href="https://pkg.go.dev/builtin#string">string</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go">ngc</a>.<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go#OrgUserAddRoleParams">OrgUserAddRoleParams</a>) (<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go/shared#User">User</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v2/org/{org-name}/users/{id}/invite">client.Org.Users.<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go#OrgUserService.Invite">Invite</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, orgName <a href="https://pkg.go.dev/builtin#string">string</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go">ngc</a>.<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go#OrgUserInviteParams">OrgUserInviteParams</a>) (<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go/shared#User">User</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /v2/org/{org-name}/users/{id}">client.Org.Users.<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go#OrgUserService.Remove">Remove</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, orgName <a href="https://pkg.go.dev/builtin#string">string</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go">ngc</a>.<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go#OrgUserRemoveParams">OrgUserRemoveParams</a>) (<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go">ngc</a>.<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go#OrgUserRemoveResponse">OrgUserRemoveResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /v2/org/{org-name}/users/{id}/remove-role">client.Org.Users.<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go#OrgUserService.RemoveRole">RemoveRole</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, orgName <a href="https://pkg.go.dev/builtin#string">string</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go">ngc</a>.<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go#OrgUserRemoveRoleParams">OrgUserRemoveRoleParams</a>) (<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go/shared#User">User</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### Invitations

Response Types:

- <a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go">ngc</a>.<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go#OrgUserInvitationListResponse">OrgUserInvitationListResponse</a>
- <a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go">ngc</a>.<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go#OrgUserInvitationDeleteResponse">OrgUserInvitationDeleteResponse</a>

Methods:

- <code title="get /v2/org/{org-name}/users/invitations">client.Org.Users.Invitations.<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go#OrgUserInvitationService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, orgName <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go">ngc</a>.<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go#OrgUserInvitationListParams">OrgUserInvitationListParams</a>) (<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go/internal/pagination">pagination</a>.<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go/internal/pagination#PageNumberInvitations">PageNumberInvitations</a>[<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go/shared#OrgUserInvitationListResponse">OrgUserInvitationListResponse</a>], <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /v2/org/{org-name}/users/invitations/{id}">client.Org.Users.Invitations.<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go#OrgUserInvitationService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, orgName <a href="https://pkg.go.dev/builtin#string">string</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go">ngc</a>.<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go#OrgUserInvitationDeleteResponse">OrgUserInvitationDeleteResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v2/org/{org-name}/users/invitations/{id}/resend-invitation-email">client.Org.Users.Invitations.<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go#OrgUserInvitationService.InviteResend">InviteResend</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, orgName <a href="https://pkg.go.dev/builtin#string">string</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go/shared#User">User</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Teams

Response Types:

- <a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go">ngc</a>.<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go#OrgTeamDeleteResponse">OrgTeamDeleteResponse</a>

Methods:

- <code title="get /v2/org/{org-name}/teams/{team-name}">client.Org.Teams.<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go#OrgTeamService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, orgName <a href="https://pkg.go.dev/builtin#string">string</a>, teamName <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go">ngc</a>.<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go#TeamResponse">TeamResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="patch /v2/org/{org-name}/teams/{team-name}">client.Org.Teams.<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go#OrgTeamService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, orgName <a href="https://pkg.go.dev/builtin#string">string</a>, teamName <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go">ngc</a>.<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go#OrgTeamUpdateParams">OrgTeamUpdateParams</a>) (<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go">ngc</a>.<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go#TeamResponse">TeamResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /v2/org/{org-name}/teams/{team-name}">client.Org.Teams.<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go#OrgTeamService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, orgName <a href="https://pkg.go.dev/builtin#string">string</a>, teamName <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go">ngc</a>.<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go#OrgTeamDeleteResponse">OrgTeamDeleteResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### Users

Methods:

- <code title="patch /v2/org/{org-name}/team/{team-name}/users/{id}/add-role">client.Org.Teams.Users.<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go#OrgTeamUserService.AddRole">AddRole</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, orgName <a href="https://pkg.go.dev/builtin#string">string</a>, teamName <a href="https://pkg.go.dev/builtin#string">string</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go">ngc</a>.<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go#OrgTeamUserAddRoleParams">OrgTeamUserAddRoleParams</a>) (<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go/shared#User">User</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="patch /v2/org/{org-name}/team/{team-name}/users/{id}/update-role">client.Org.Teams.Users.<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go#OrgTeamUserService.UpdateRole">UpdateRole</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, orgName <a href="https://pkg.go.dev/builtin#string">string</a>, teamName <a href="https://pkg.go.dev/builtin#string">string</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go">ngc</a>.<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go#OrgTeamUserUpdateRoleParams">OrgTeamUserUpdateRoleParams</a>) (<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go/shared#User">User</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Team

### Users

Methods:

- <code title="delete /v2/org/{org-name}/team/{team-name}/users/{id}/remove-role">client.Org.Team.Users.<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go#OrgTeamUserService.RemoveRole">RemoveRole</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, orgName <a href="https://pkg.go.dev/builtin#string">string</a>, teamName <a href="https://pkg.go.dev/builtin#string">string</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go">ngc</a>.<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go#OrgTeamUserRemoveRoleParams">OrgTeamUserRemoveRoleParams</a>) (<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go/shared#User">User</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

#### Invitations

Response Types:

- <a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go">ngc</a>.<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go#OrgTeamUserInvitationListResponse">OrgTeamUserInvitationListResponse</a>
- <a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go">ngc</a>.<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go#OrgTeamUserInvitationDeleteResponse">OrgTeamUserInvitationDeleteResponse</a>

Methods:

- <code title="get /v2/org/{org-name}/team/{team-name}/users/invitations">client.Org.Team.Users.Invitations.<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go#OrgTeamUserInvitationService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, orgName <a href="https://pkg.go.dev/builtin#string">string</a>, teamName <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go">ngc</a>.<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go#OrgTeamUserInvitationListParams">OrgTeamUserInvitationListParams</a>) (<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go/internal/pagination">pagination</a>.<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go/internal/pagination#PageNumberInvitations">PageNumberInvitations</a>[<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go/shared#OrgTeamUserInvitationListResponse">OrgTeamUserInvitationListResponse</a>], <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /v2/org/{org-name}/team/{team-name}/users/invitations/{id}">client.Org.Team.Users.Invitations.<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go#OrgTeamUserInvitationService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, orgName <a href="https://pkg.go.dev/builtin#string">string</a>, teamName <a href="https://pkg.go.dev/builtin#string">string</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go">ngc</a>.<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go#OrgTeamUserInvitationDeleteResponse">OrgTeamUserInvitationDeleteResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v2/org/{org-name}/team/{team-name}/users/invitations/{id}/resend-invitation-email">client.Org.Team.Users.Invitations.<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go#OrgTeamUserInvitationService.InviteResend">InviteResend</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, orgName <a href="https://pkg.go.dev/builtin#string">string</a>, teamName <a href="https://pkg.go.dev/builtin#string">string</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go/shared#User">User</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# V2AdminOrgUsers

Methods:

- <code title="patch /v2/admin/org/{org-name}/users/{id}/add-role">client.V2AdminOrgUsers.<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go#V2AdminOrgUserService.AddRole">AddRole</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, orgName <a href="https://pkg.go.dev/builtin#string">string</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go">ngc</a>.<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go#V2AdminOrgUserAddRoleParams">V2AdminOrgUserAddRoleParams</a>) (<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go/shared#User">User</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# V2AdminOrgTeams

Methods:

- <code title="get /v2/admin/org/{org-name}/teams/{team-name}">client.V2AdminOrgTeams.<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go#V2AdminOrgTeamService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, orgName <a href="https://pkg.go.dev/builtin#string">string</a>, teamName <a href="https://pkg.go.dev/builtin#string">string</a>) (http.Response, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="patch /v2/admin/org/{org-name}/teams/{team-name}">client.V2AdminOrgTeams.<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go#V2AdminOrgTeamService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, orgName <a href="https://pkg.go.dev/builtin#string">string</a>, teamName <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go">ngc</a>.<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go#V2AdminOrgTeamUpdateParams">V2AdminOrgTeamUpdateParams</a>) (http.Response, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# V2AdminOrgTeamUsers

Methods:

- <code title="patch /v2/admin/org/{org-name}/team/{team-name}/users/{id}/add-role">client.V2AdminOrgTeamUsers.<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go#V2AdminOrgTeamUserService.AddRole">AddRole</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, orgName <a href="https://pkg.go.dev/builtin#string">string</a>, teamName <a href="https://pkg.go.dev/builtin#string">string</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go">ngc</a>.<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go#V2AdminOrgTeamUserAddRoleParams">V2AdminOrgTeamUserAddRoleParams</a>) (<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go/shared#User">User</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# V2AdminOrgEntitlements

Methods:

- <code title="get /v2/admin/org/{org-name}/entitlements">client.V2AdminOrgEntitlements.<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go#V2AdminOrgEntitlementService.GetAll">GetAll</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, orgName <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go">ngc</a>.<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go#V2AdminOrgEntitlementGetAllParams">V2AdminOrgEntitlementGetAllParams</a>) (http.Response, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# V2AdminEntitlements

Methods:

- <code title="get /v2/admin/entitlements">client.V2AdminEntitlements.<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go#V2AdminEntitlementService.GetAll">GetAll</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go">ngc</a>.<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go#V2AdminEntitlementGetAllParams">V2AdminEntitlementGetAllParams</a>) (http.Response, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Services

Response Types:

- <a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go">ngc</a>.<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go#ServiceVersionResponse">ServiceVersionResponse</a>

Methods:

- <code title="get /version">client.Services.<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go#ServiceService.Version">Version</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go">ngc</a>.<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go#ServiceVersionParams">ServiceVersionParams</a>) (<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go">ngc</a>.<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go#ServiceVersionResponse">ServiceVersionResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# V3OrgsUsers

Methods:

- <code title="get /v3/orgs/{org-name}/users/{user-email-or-id}">client.V3OrgsUsers.<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go#V3OrgsUserService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, orgName <a href="https://pkg.go.dev/builtin#string">string</a>, userEmailOrID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go/shared#User">User</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# V3OrgsTeamsUsers

Methods:

- <code title="get /v3/orgs/{org-name}/teams/{team-name}/users/{user-email-or-id}">client.V3OrgsTeamsUsers.<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go#V3OrgsTeamsUserService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, orgName <a href="https://pkg.go.dev/builtin#string">string</a>, teamName <a href="https://pkg.go.dev/builtin#string">string</a>, userEmailOrID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go/shared#User">User</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# V3Orgs

Response Types:

- <a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go">ngc</a>.<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go#OrgInvitation">OrgInvitation</a>

Methods:

- <code title="get /v3/orgs/proto-org/validate">client.V3Orgs.<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go#V3OrgService.Validate">Validate</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go">ngc</a>.<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go#V3OrgValidateParams">V3OrgValidateParams</a>) (<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go">ngc</a>.<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go#OrgInvitation">OrgInvitation</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Roles

Response Types:

- <a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go">ngc</a>.<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go#UserRoleDefinitions">UserRoleDefinitions</a>

Methods:

- <code title="get /roles">client.Roles.<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go#RoleService.GetAll">GetAll</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go">ngc</a>.<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go#RoleGetAllParams">RoleGetAllParams</a>) (<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go">ngc</a>.<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go#UserRoleDefinitions">UserRoleDefinitions</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# PublicKeys

Methods:

- <code title="get /public-keys">client.PublicKeys.<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go#PublicKeyService.GetAll">GetAll</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) ([]<a href="https://pkg.go.dev/builtin#string">string</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Health

Methods:

- <code title="get /health">client.Health.<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go#HealthService.GetAll">GetAll</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go/shared#Health">Health</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## All

Methods:

- <code title="get /health/all">client.Health.All.<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go#HealthAllService.GetAll">GetAll</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go">ngc</a>.<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go#HealthAllGetAllParams">HealthAllGetAllParams</a>) (<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go/shared#Health">Health</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# SwaggerResources

Methods:

- <code title="post /swagger-resources">client.SwaggerResources.<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go#SwaggerResourceService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (http.Response, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="patch /swagger-resources">client.SwaggerResources.<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go#SwaggerResourceService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (http.Response, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /swagger-resources">client.SwaggerResources.<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go#SwaggerResourceService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (http.Response, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /swagger-resources">client.SwaggerResources.<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go#SwaggerResourceService.GetAll">GetAll</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (http.Response, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Configuration

### Ui

Methods:

- <code title="post /swagger-resources/configuration/ui">client.SwaggerResources.Configuration.Ui.<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go#SwaggerResourceConfigurationUiService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (http.Response, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /swagger-resources/configuration/ui">client.SwaggerResources.Configuration.Ui.<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go#SwaggerResourceConfigurationUiService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (http.Response, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="patch /swagger-resources/configuration/ui">client.SwaggerResources.Configuration.Ui.<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go#SwaggerResourceConfigurationUiService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (http.Response, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /swagger-resources/configuration/ui">client.SwaggerResources.Configuration.Ui.<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go#SwaggerResourceConfigurationUiService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (http.Response, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### Security

Methods:

- <code title="post /swagger-resources/configuration/security">client.SwaggerResources.Configuration.Security.<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go#SwaggerResourceConfigurationSecurityService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (http.Response, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /swagger-resources/configuration/security">client.SwaggerResources.Configuration.Security.<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go#SwaggerResourceConfigurationSecurityService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (http.Response, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="patch /swagger-resources/configuration/security">client.SwaggerResources.Configuration.Security.<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go#SwaggerResourceConfigurationSecurityService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (http.Response, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /swagger-resources/configuration/security">client.SwaggerResources.Configuration.Security.<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go#SwaggerResourceConfigurationSecurityService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (http.Response, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
