# Shared Response Types

- <a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go/shared#Health">Health</a>
- <a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go/shared#MeteringResultList">MeteringResultList</a>
- <a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go/shared#Team">Team</a>
- <a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go/shared#TeamList">TeamList</a>
- <a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go/shared#User">User</a>
- <a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go/shared#UserList">UserList</a>

# Orgs

Response Types:

- <a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go">ngc</a>.<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go#OrgList">OrgList</a>
- <a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go">ngc</a>.<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go#OrgResponse">OrgResponse</a>
- <a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go">ngc</a>.<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go#OrgListResponse">OrgListResponse</a>

Methods:

- <code title="post /v3/orgs">client.Orgs.<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go#OrgService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go">ngc</a>.<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go#OrgNewParams">OrgNewParams</a>) (<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go">ngc</a>.<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go#OrgResponse">OrgResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v2/orgs/{org-name}">client.Orgs.<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go#OrgService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, orgName <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go">ngc</a>.<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go#OrgResponse">OrgResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="patch /v2/orgs/{org-name}">client.Orgs.<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go#OrgService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, orgName <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go">ngc</a>.<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go#OrgUpdateParams">OrgUpdateParams</a>) (<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go">ngc</a>.<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go#OrgResponse">OrgResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v2/orgs">client.Orgs.<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go#OrgService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go">ngc</a>.<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go#OrgListParams">OrgListParams</a>) (<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go/packages/pagination">pagination</a>.<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go/packages/pagination#PageNumberOrganizations">PageNumberOrganizations</a>[<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go">ngc</a>.<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go#OrgListResponse">OrgListResponse</a>], <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Users

Response Types:

- <a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go">ngc</a>.<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go#OrgUserListResponse">OrgUserListResponse</a>
- <a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go">ngc</a>.<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go#OrgUserDeleteResponse">OrgUserDeleteResponse</a>

Methods:

- <code title="post /v2/org/{org-name}/users">client.Orgs.Users.<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go#OrgUserService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, orgName <a href="https://pkg.go.dev/builtin#string">string</a>, params <a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go">ngc</a>.<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go#OrgUserNewParams">OrgUserNewParams</a>) (<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go/shared#User">User</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v2/org/{org-name}/users">client.Orgs.Users.<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go#OrgUserService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, orgName <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go">ngc</a>.<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go#OrgUserListParams">OrgUserListParams</a>) (<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go/packages/pagination">pagination</a>.<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go/packages/pagination#PageNumberUsers">PageNumberUsers</a>[<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go/shared#OrgUserListResponse">OrgUserListResponse</a>], <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /v3/orgs/{org-name}/users/{id}">client.Orgs.Users.<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go#OrgUserService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, orgName <a href="https://pkg.go.dev/builtin#string">string</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go">ngc</a>.<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go#OrgUserDeleteParams">OrgUserDeleteParams</a>) (<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go">ngc</a>.<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go#OrgUserDeleteResponse">OrgUserDeleteResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="patch /v3/orgs/{org-name}/users/{user-email-or-id}/add-role">client.Orgs.Users.<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go#OrgUserService.AddRole">AddRole</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, orgName <a href="https://pkg.go.dev/builtin#string">string</a>, userEmailOrID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go">ngc</a>.<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go#OrgUserAddRoleParams">OrgUserAddRoleParams</a>) (<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go/shared#User">User</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /v3/orgs/{org-name}/users/{user-email-or-id}/remove-role">client.Orgs.Users.<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go#OrgUserService.RemoveRole">RemoveRole</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, orgName <a href="https://pkg.go.dev/builtin#string">string</a>, userEmailOrID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go">ngc</a>.<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go#OrgUserRemoveRoleParams">OrgUserRemoveRoleParams</a>) (<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go/shared#User">User</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="patch /v2/org/{org-name}/users/{id}/update-role">client.Orgs.Users.<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go#OrgUserService.UpdateRole">UpdateRole</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, orgName <a href="https://pkg.go.dev/builtin#string">string</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go">ngc</a>.<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go#OrgUserUpdateRoleParams">OrgUserUpdateRoleParams</a>) (<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go/shared#User">User</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### NcaInvitations

Methods:

- <code title="post /v3/orgs/{org-name}/users/nca-invitations">client.Orgs.Users.NcaInvitations.<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go#OrgUserNcaInvitationService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, orgName <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go">ngc</a>.<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go#OrgUserNcaInvitationNewParams">OrgUserNcaInvitationNewParams</a>) (<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go/shared#User">User</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Teams

Methods:

- <code title="get /v2/org/{org-name}/teams">client.Orgs.Teams.<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go#OrgTeamService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, orgName <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go">ngc</a>.<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go#OrgTeamListParams">OrgTeamListParams</a>) (<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go/packages/pagination">pagination</a>.<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go/packages/pagination#PageNumberTeams">PageNumberTeams</a>[<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go/shared#Team">Team</a>], <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### Users

Response Types:

- <a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go">ngc</a>.<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go#OrgTeamUserDeleteResponse">OrgTeamUserDeleteResponse</a>

Methods:

- <code title="delete /v3/orgs/{org-name}/teams/{team-name}/users/{id}">client.Orgs.Teams.Users.<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go#OrgTeamUserService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, orgName <a href="https://pkg.go.dev/builtin#string">string</a>, teamName <a href="https://pkg.go.dev/builtin#string">string</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go">ngc</a>.<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go#OrgTeamUserDeleteParams">OrgTeamUserDeleteParams</a>) (<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go">ngc</a>.<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go#OrgTeamUserDeleteResponse">OrgTeamUserDeleteResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="patch /v3/orgs/{org-name}/teams/{team-name}/users/{user-email-or-id}/add-role">client.Orgs.Teams.Users.<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go#OrgTeamUserService.AddRole">AddRole</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, orgName <a href="https://pkg.go.dev/builtin#string">string</a>, teamName <a href="https://pkg.go.dev/builtin#string">string</a>, userEmailOrID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go">ngc</a>.<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go#OrgTeamUserAddRoleParams">OrgTeamUserAddRoleParams</a>) (<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go/shared#User">User</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /v3/orgs/{org-name}/teams/{team-name}/users/{user-email-or-id}/remove-role">client.Orgs.Teams.Users.<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go#OrgTeamUserService.RemoveRole">RemoveRole</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, orgName <a href="https://pkg.go.dev/builtin#string">string</a>, teamName <a href="https://pkg.go.dev/builtin#string">string</a>, userEmailOrID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go">ngc</a>.<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go#OrgTeamUserRemoveRoleParams">OrgTeamUserRemoveRoleParams</a>) (<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go/shared#User">User</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### StarfleetIDs

Methods:

- <code title="get /v2/org/{org-name}/team/{team-name}/starfleetIds/{starfleet-id}">client.Orgs.Teams.StarfleetIDs.<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go#OrgTeamStarfleetIDService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, orgName <a href="https://pkg.go.dev/builtin#string">string</a>, teamName <a href="https://pkg.go.dev/builtin#string">string</a>, starfleetID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go/shared#User">User</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### NcaInvitations

Methods:

- <code title="post /v3/orgs/{org-name}/teams/{team-name}/users/nca-invitations">client.Orgs.Teams.NcaInvitations.<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go#OrgTeamNcaInvitationService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, orgName <a href="https://pkg.go.dev/builtin#string">string</a>, teamName <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go">ngc</a>.<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go#OrgTeamNcaInvitationNewParams">OrgTeamNcaInvitationNewParams</a>) (<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go/shared#User">User</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## ProtoOrg

Methods:

- <code title="post /v3/orgs/proto-org">client.Orgs.ProtoOrg.<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go#OrgProtoOrgService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go">ngc</a>.<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go#OrgProtoOrgNewParams">OrgProtoOrgNewParams</a>) (<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go">ngc</a>.<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go#OrgResponse">OrgResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

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

# Me

## APIKey

Response Types:

- <a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go">ngc</a>.<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go#UserKeyResponse">UserKeyResponse</a>

Methods:

- <code title="get /v2/users/me">client.Me.APIKey.<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go#MeAPIKeyService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go">ngc</a>.<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go#MeAPIKeyGetParams">MeAPIKeyGetParams</a>) (<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go/shared#User">User</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="patch /v2/users/me">client.Me.APIKey.<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go#MeAPIKeyService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go">ngc</a>.<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go#MeAPIKeyUpdateParams">MeAPIKeyUpdateParams</a>) (<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go/shared#User">User</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /v2/users/me/api-key">client.Me.APIKey.<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go#MeAPIKeyService.NewAPIKey">NewAPIKey</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go">ngc</a>.<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go#UserKeyResponse">UserKeyResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Admin

## Entitlements

Methods:

- <code title="get /v2/admin/entitlements">client.Admin.Entitlements.<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go#AdminEntitlementService.GetAll">GetAll</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go">ngc</a>.<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go#AdminEntitlementGetAllParams">AdminEntitlementGetAllParams</a>) (http.Response, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Orgs

Response Types:

- <a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go">ngc</a>.<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go#AdminOrgOrgOwnerBackfillResponse">AdminOrgOrgOwnerBackfillResponse</a>
- <a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go">ngc</a>.<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go#AdminOrgValidateResponse">AdminOrgValidateResponse</a>

Methods:

- <code title="post /v2/admin/orgs">client.Admin.Orgs.<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go#AdminOrgService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go">ngc</a>.<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go#AdminOrgNewParams">AdminOrgNewParams</a>) (http.Response, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v3/admin/org/{org-name}">client.Admin.Orgs.<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go#AdminOrgService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, orgName <a href="https://pkg.go.dev/builtin#string">string</a>) (http.Response, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="patch /v3/admin/org/{org-name}">client.Admin.Orgs.<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go#AdminOrgService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, orgName <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go">ngc</a>.<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go#AdminOrgUpdateParams">AdminOrgUpdateParams</a>) (http.Response, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /v2/admin/backfill-orgs-to-kratos">client.Admin.Orgs.<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go#AdminOrgService.BackfillOrgsToKratos">BackfillOrgsToKratos</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (http.Response, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /v2/admin/org/{org-name}/enablement">client.Admin.Orgs.<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go#AdminOrgService.Enable">Enable</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, orgName <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go">ngc</a>.<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go#AdminOrgEnableParams">AdminOrgEnableParams</a>) (http.Response, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /v2/admin/org/{org-name}/org-owner-backfill">client.Admin.Orgs.<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go#AdminOrgService.OrgOwnerBackfill">OrgOwnerBackfill</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, orgName <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go">ngc</a>.<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go#AdminOrgOrgOwnerBackfillResponse">AdminOrgOrgOwnerBackfillResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v3/orgs/proto-org/validate">client.Admin.Orgs.<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go#AdminOrgService.Validate">Validate</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go">ngc</a>.<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go#AdminOrgValidateParams">AdminOrgValidateParams</a>) (<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go">ngc</a>.<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go#AdminOrgValidateResponse">AdminOrgValidateResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### Users

Response Types:

- <a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go">ngc</a>.<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go#AdminOrgUserRemoveResponse">AdminOrgUserRemoveResponse</a>

Methods:

- <code title="post /v2/admin/org/{org-name}/users">client.Admin.Orgs.Users.<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go#AdminOrgUserService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, orgName <a href="https://pkg.go.dev/builtin#string">string</a>, params <a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go">ngc</a>.<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go#AdminOrgUserNewParams">AdminOrgUserNewParams</a>) (<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go/shared#User">User</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v3/orgs/{org-name}/users/{user-email-or-id}">client.Admin.Orgs.Users.<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go#AdminOrgUserService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, orgName <a href="https://pkg.go.dev/builtin#string">string</a>, userEmailOrID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go/shared#User">User</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /v2/admin/org/{org-name}/users/{id}">client.Admin.Orgs.Users.<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go#AdminOrgUserService.Add">Add</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, orgName <a href="https://pkg.go.dev/builtin#string">string</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go">ngc</a>.<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go#AdminOrgUserAddParams">AdminOrgUserAddParams</a>) (<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go/shared#User">User</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="patch /v2/admin/org/{org-name}/users/{id}/add-role">client.Admin.Orgs.Users.<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go#AdminOrgUserService.AddRole">AddRole</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, orgName <a href="https://pkg.go.dev/builtin#string">string</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go">ngc</a>.<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go#AdminOrgUserAddRoleParams">AdminOrgUserAddRoleParams</a>) (<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go/shared#User">User</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v2/admin/org/{org-name}/entitlements">client.Admin.Orgs.Users.<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go#AdminOrgUserService.GetEntitlements">GetEntitlements</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, orgName <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go">ngc</a>.<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go#AdminOrgUserGetEntitlementsParams">AdminOrgUserGetEntitlementsParams</a>) (http.Response, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /v2/admin/org/{org-name}/users/{id}">client.Admin.Orgs.Users.<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go#AdminOrgUserService.Remove">Remove</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, orgName <a href="https://pkg.go.dev/builtin#string">string</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go">ngc</a>.<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go#AdminOrgUserRemoveResponse">AdminOrgUserRemoveResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### Teams

#### Users

Methods:

- <code title="post /v2/admin/org/{org-name}/team/{team-name}/users">client.Admin.Orgs.Teams.Users.<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go#AdminOrgTeamUserService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, orgName <a href="https://pkg.go.dev/builtin#string">string</a>, teamName <a href="https://pkg.go.dev/builtin#string">string</a>, params <a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go">ngc</a>.<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go#AdminOrgTeamUserNewParams">AdminOrgTeamUserNewParams</a>) (<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go/shared#User">User</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v2/admin/org/{org-name}/teams/{team-name}">client.Admin.Orgs.Teams.Users.<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go#AdminOrgTeamUserService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, orgName <a href="https://pkg.go.dev/builtin#string">string</a>, teamName <a href="https://pkg.go.dev/builtin#string">string</a>) (http.Response, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="patch /v2/admin/org/{org-name}/teams/{team-name}">client.Admin.Orgs.Teams.Users.<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go#AdminOrgTeamUserService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, orgName <a href="https://pkg.go.dev/builtin#string">string</a>, teamName <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go">ngc</a>.<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go#AdminOrgTeamUserUpdateParams">AdminOrgTeamUserUpdateParams</a>) (http.Response, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /v2/admin/org/{org-name}/team/{team-name}/users/{id}">client.Admin.Orgs.Teams.Users.<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go#AdminOrgTeamUserService.Add">Add</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, orgName <a href="https://pkg.go.dev/builtin#string">string</a>, teamName <a href="https://pkg.go.dev/builtin#string">string</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go">ngc</a>.<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go#AdminOrgTeamUserAddParams">AdminOrgTeamUserAddParams</a>) (<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go/shared#User">User</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="patch /v2/admin/org/{org-name}/team/{team-name}/users/{id}/add-role">client.Admin.Orgs.Teams.Users.<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go#AdminOrgTeamUserService.AddRole">AddRole</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, orgName <a href="https://pkg.go.dev/builtin#string">string</a>, teamName <a href="https://pkg.go.dev/builtin#string">string</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go">ngc</a>.<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go#AdminOrgTeamUserAddRoleParams">AdminOrgTeamUserAddRoleParams</a>) (<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go/shared#User">User</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v3/orgs/{org-name}/teams/{team-name}/users/{user-email-or-id}">client.Admin.Orgs.Teams.Users.<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go#AdminOrgTeamUserService.GetUser">GetUser</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, orgName <a href="https://pkg.go.dev/builtin#string">string</a>, teamName <a href="https://pkg.go.dev/builtin#string">string</a>, userEmailOrID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go/shared#User">User</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Users

Response Types:

- <a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go">ngc</a>.<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go#AdminUserCRMSyncResponse">AdminUserCRMSyncResponse</a>

Methods:

- <code title="post /v2/admin/users/{id}/crm-sync">client.Admin.Users.<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go#AdminUserService.CRMSync">CRMSync</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#int64">int64</a>) (<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go">ngc</a>.<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go#AdminUserCRMSyncResponse">AdminUserCRMSyncResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /v2/admin/users/{id}/migrate-deprecated-roles">client.Admin.Users.<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go#AdminUserService.MigrateDeprecatedRoles">MigrateDeprecatedRoles</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go/shared">shared</a>.<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go/shared#User">User</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Services

Response Types:

- <a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go">ngc</a>.<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go#ServiceVersionResponse">ServiceVersionResponse</a>

Methods:

- <code title="get /version">client.Services.<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go#ServiceService.Version">Version</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go">ngc</a>.<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go#ServiceVersionParams">ServiceVersionParams</a>) (<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go">ngc</a>.<a href="https://pkg.go.dev/github.com/NVIDIADemo/ngc-go#ServiceVersionResponse">ServiceVersionResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

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
