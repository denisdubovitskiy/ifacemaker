// Package client generated with github.com/denisdubovitskiy/ifacemaker, DO NOT EDIT.
package client

import (
	"io"
	"net/http"
	"net/url"
	"time"

	"github.com/mattermost/mattermost-server/v5/model"
)

//go:generate ifacemaker --source-pkg github.com/mattermost/mattermost-server/v5@v5.39.3 --module-path model --result-pkg client --struct-name Client4 --interface-name Client4
type Client4 interface {
	// SetBoolString is a helper method for overriding how true and false query string parameters are
	// sent to the server.
	//
	// This method is only exposed for testing. It is never necessary to configure these values
	// in production.
	SetBoolString(value bool, valueStr string)
	// Must is a convenience function used for testing.
	Must(result interface{}, resp *model.Response) interface{}
	SetToken(token string)
	// MockSession is deprecated in favour of SetToken
	MockSession(token string)
	SetOAuthToken(token string)
	ClearOAuthToken()
	GetUsersRoute() string
	GetUserRoute(userId string) string
	GetUserThreadsRoute(userID string, teamID string) string
	GetUserThreadRoute(userId string, teamId string, threadId string) string
	GetUserCategoryRoute(userID string, teamID string) string
	GetUserAccessTokensRoute() string
	GetUserAccessTokenRoute(tokenId string) string
	GetUserByUsernameRoute(userName string) string
	GetUserByEmailRoute(email string) string
	GetBotsRoute() string
	GetBotRoute(botUserId string) string
	GetTeamsRoute() string
	GetTeamRoute(teamId string) string
	GetTeamAutoCompleteCommandsRoute(teamId string) string
	GetTeamByNameRoute(teamName string) string
	GetTeamMemberRoute(teamId string, userId string) string
	GetTeamMembersRoute(teamId string) string
	GetTeamStatsRoute(teamId string) string
	GetTeamImportRoute(teamId string) string
	GetChannelsRoute() string
	GetChannelsForTeamRoute(teamId string) string
	GetChannelRoute(channelId string) string
	GetChannelByNameRoute(channelName string, teamId string) string
	GetChannelsForTeamForUserRoute(teamId string, userId string, includeDeleted bool) string
	GetChannelByNameForTeamNameRoute(channelName string, teamName string) string
	GetChannelMembersRoute(channelId string) string
	GetChannelMemberRoute(channelId string, userId string) string
	GetPostsRoute() string
	GetPostsEphemeralRoute() string
	GetConfigRoute() string
	GetLicenseRoute() string
	GetPostRoute(postId string) string
	GetFilesRoute() string
	GetFileRoute(fileId string) string
	GetUploadsRoute() string
	GetUploadRoute(uploadId string) string
	GetPluginsRoute() string
	GetPluginRoute(pluginId string) string
	GetSystemRoute() string
	GetCloudRoute() string
	GetTestEmailRoute() string
	GetTestSiteURLRoute() string
	GetTestS3Route() string
	GetDatabaseRoute() string
	GetCacheRoute() string
	GetClusterRoute() string
	GetIncomingWebhooksRoute() string
	GetIncomingWebhookRoute(hookID string) string
	GetComplianceReportsRoute() string
	GetComplianceReportRoute(reportId string) string
	GetComplianceReportDownloadRoute(reportId string) string
	GetOutgoingWebhooksRoute() string
	GetOutgoingWebhookRoute(hookID string) string
	GetPreferencesRoute(userId string) string
	GetUserStatusRoute(userId string) string
	GetUserStatusesRoute() string
	GetSamlRoute() string
	GetLdapRoute() string
	GetBrandRoute() string
	GetDataRetentionRoute() string
	GetDataRetentionPolicyRoute(policyID string) string
	GetElasticsearchRoute() string
	GetBleveRoute() string
	GetCommandsRoute() string
	GetCommandRoute(commandId string) string
	GetCommandMoveRoute(commandId string) string
	GetEmojisRoute() string
	GetEmojiRoute(emojiId string) string
	GetEmojiByNameRoute(name string) string
	GetReactionsRoute() string
	GetOAuthAppsRoute() string
	GetOAuthAppRoute(appId string) string
	GetOpenGraphRoute() string
	GetJobsRoute() string
	GetRolesRoute() string
	GetSchemesRoute() string
	GetSchemeRoute(id string) string
	GetAnalyticsRoute() string
	GetTimezonesRoute() string
	GetChannelSchemeRoute(channelId string) string
	GetTeamSchemeRoute(teamId string) string
	GetTotalUsersStatsRoute() string
	GetRedirectLocationRoute() string
	GetServerBusyRoute() string
	GetUserTermsOfServiceRoute(userId string) string
	GetTermsOfServiceRoute() string
	GetGroupsRoute() string
	GetPublishUserTypingRoute(userId string) string
	GetGroupRoute(groupID string) string
	GetGroupSyncableRoute(groupID string, syncableID string, syncableType model.GroupSyncableType) string
	GetGroupSyncablesRoute(groupID string, syncableType model.GroupSyncableType) string
	GetImportsRoute() string
	GetExportsRoute() string
	GetExportRoute(name string) string
	GetRemoteClusterRoute() string
	GetSharedChannelsRoute() string
	GetPermissionsRoute() string
	DoApiGet(url string, etag string) (*http.Response, *model.AppError)
	DoApiPost(url string, data string) (*http.Response, *model.AppError)
	DoApiPut(url string, data string) (*http.Response, *model.AppError)
	DoApiDelete(url string) (*http.Response, *model.AppError)
	DoApiRequest(method string, url string, data string, etag string) (*http.Response, *model.AppError)
	DoApiRequestWithHeaders(method string, url string, data string, headers map[string]string) (*http.Response, *model.AppError)
	DoUploadFile(url string, data []byte, contentType string) (*model.FileUploadResponse, *model.Response)
	DoEmojiUploadFile(url string, data []byte, contentType string) (*model.Emoji, *model.Response)
	DoUploadImportTeam(url string, data []byte, contentType string) (map[string]string, *model.Response)
	// LoginById authenticates a user by user id and password.
	LoginById(id string, password string) (*model.User, *model.Response)
	// Login authenticates a user by login id, which can be username, email or some sort
	// of SSO identifier based on server configuration, and a password.
	Login(loginId string, password string) (*model.User, *model.Response)
	// LoginByLdap authenticates a user by LDAP id and password.
	LoginByLdap(loginId string, password string) (*model.User, *model.Response)
	// LoginWithDevice authenticates a user by login id (username, email or some sort
	// of SSO identifier based on configuration), password and attaches a device id to
	// the session.
	LoginWithDevice(loginId string, password string, deviceId string) (*model.User, *model.Response)
	// LoginWithMFA logs a user in with a MFA token
	LoginWithMFA(loginId string, password string, mfaToken string) (*model.User, *model.Response)
	// Logout terminates the current user's session.
	Logout() (bool, *model.Response)
	// SwitchAccountType changes a user's login type from one type to another.
	SwitchAccountType(switchRequest *model.SwitchRequest) (string, *model.Response)
	// CreateUser creates a user in the system based on the provided user struct.
	CreateUser(user *model.User) (*model.User, *model.Response)
	// CreateUserWithToken creates a user in the system based on the provided tokenId.
	CreateUserWithToken(user *model.User, tokenId string) (*model.User, *model.Response)
	// CreateUserWithInviteId creates a user in the system based on the provided invited id.
	CreateUserWithInviteId(user *model.User, inviteId string) (*model.User, *model.Response)
	// GetMe returns the logged in user.
	GetMe(etag string) (*model.User, *model.Response)
	// GetUser returns a user based on the provided user id string.
	GetUser(userId string, etag string) (*model.User, *model.Response)
	// GetUserByUsername returns a user based on the provided user name string.
	GetUserByUsername(userName string, etag string) (*model.User, *model.Response)
	// GetUserByEmail returns a user based on the provided user email string.
	GetUserByEmail(email string, etag string) (*model.User, *model.Response)
	// AutocompleteUsersInTeam returns the users on a team based on search term.
	AutocompleteUsersInTeam(teamId string, username string, limit int, etag string) (*model.UserAutocomplete, *model.Response)
	// AutocompleteUsersInChannel returns the users in a channel based on search term.
	AutocompleteUsersInChannel(teamId string, channelId string, username string, limit int, etag string) (*model.UserAutocomplete, *model.Response)
	// AutocompleteUsers returns the users in the system based on search term.
	AutocompleteUsers(username string, limit int, etag string) (*model.UserAutocomplete, *model.Response)
	// GetDefaultProfileImage gets the default user's profile image. Must be logged in.
	GetDefaultProfileImage(userId string) ([]byte, *model.Response)
	// GetProfileImage gets user's profile image. Must be logged in.
	GetProfileImage(userId string, etag string) ([]byte, *model.Response)
	// GetUsers returns a page of users on the system. Page counting starts at 0.
	GetUsers(page int, perPage int, etag string) ([]*model.User, *model.Response)
	// GetUsersInTeam returns a page of users on a team. Page counting starts at 0.
	GetUsersInTeam(teamId string, page int, perPage int, etag string) ([]*model.User, *model.Response)
	// GetNewUsersInTeam returns a page of users on a team. Page counting starts at 0.
	GetNewUsersInTeam(teamId string, page int, perPage int, etag string) ([]*model.User, *model.Response)
	// GetRecentlyActiveUsersInTeam returns a page of users on a team. Page counting starts at 0.
	GetRecentlyActiveUsersInTeam(teamId string, page int, perPage int, etag string) ([]*model.User, *model.Response)
	// GetActiveUsersInTeam returns a page of users on a team. Page counting starts at 0.
	GetActiveUsersInTeam(teamId string, page int, perPage int, etag string) ([]*model.User, *model.Response)
	// GetUsersNotInTeam returns a page of users who are not in a team. Page counting starts at 0.
	GetUsersNotInTeam(teamId string, page int, perPage int, etag string) ([]*model.User, *model.Response)
	// GetUsersInChannel returns a page of users in a channel. Page counting starts at 0.
	GetUsersInChannel(channelId string, page int, perPage int, etag string) ([]*model.User, *model.Response)
	// GetUsersInChannelByStatus returns a page of users in a channel. Page counting starts at 0. Sorted by Status
	GetUsersInChannelByStatus(channelId string, page int, perPage int, etag string) ([]*model.User, *model.Response)
	// GetUsersNotInChannel returns a page of users not in a channel. Page counting starts at 0.
	GetUsersNotInChannel(teamId string, channelId string, page int, perPage int, etag string) ([]*model.User, *model.Response)
	// GetUsersWithoutTeam returns a page of users on the system that aren't on any teams. Page counting starts at 0.
	GetUsersWithoutTeam(page int, perPage int, etag string) ([]*model.User, *model.Response)
	// GetUsersInGroup returns a page of users in a group. Page counting starts at 0.
	GetUsersInGroup(groupID string, page int, perPage int, etag string) ([]*model.User, *model.Response)
	// GetUsersByIds returns a list of users based on the provided user ids.
	GetUsersByIds(userIds []string) ([]*model.User, *model.Response)
	// GetUsersByIds returns a list of users based on the provided user ids.
	GetUsersByIdsWithOptions(userIds []string, options *model.UserGetByIdsOptions) ([]*model.User, *model.Response)
	// GetUsersByUsernames returns a list of users based on the provided usernames.
	GetUsersByUsernames(usernames []string) ([]*model.User, *model.Response)
	// GetUsersByGroupChannelIds returns a map with channel ids as keys
	// and a list of users as values based on the provided user ids.
	GetUsersByGroupChannelIds(groupChannelIds []string) (map[string][]*model.User, *model.Response)
	// SearchUsers returns a list of users based on some search criteria.
	SearchUsers(search *model.UserSearch) ([]*model.User, *model.Response)
	// UpdateUser updates a user in the system based on the provided user struct.
	UpdateUser(user *model.User) (*model.User, *model.Response)
	// PatchUser partially updates a user in the system. Any missing fields are not updated.
	PatchUser(userId string, patch *model.UserPatch) (*model.User, *model.Response)
	// UpdateUserAuth updates a user AuthData (uthData, authService and password) in the system.
	UpdateUserAuth(userId string, userAuth *model.UserAuth) (*model.UserAuth, *model.Response)
	// UpdateUserMfa activates multi-factor authentication for a user if activate
	// is true and a valid code is provided. If activate is false, then code is not
	// required and multi-factor authentication is disabled for the user.
	UpdateUserMfa(userId string, code string, activate bool) (bool, *model.Response)
	// CheckUserMfa checks whether a user has MFA active on their account or not based on the
	// provided login id.
	// Deprecated: Clients should use Login method and check for MFA Error
	CheckUserMfa(loginId string) (bool, *model.Response)
	// GenerateMfaSecret will generate a new MFA secret for a user and return it as a string and
	// as a base64 encoded image QR code.
	GenerateMfaSecret(userId string) (*model.MfaSecret, *model.Response)
	// UpdateUserPassword updates a user's password. Must be logged in as the user or be a system administrator.
	UpdateUserPassword(userId string, currentPassword string, newPassword string) (bool, *model.Response)
	// UpdateUserHashedPassword updates a user's password with an already-hashed password. Must be a system administrator.
	UpdateUserHashedPassword(userId string, newHashedPassword string) (bool, *model.Response)
	// PromoteGuestToUser convert a guest into a regular user
	PromoteGuestToUser(guestId string) (bool, *model.Response)
	// DemoteUserToGuest convert a regular user into a guest
	DemoteUserToGuest(guestId string) (bool, *model.Response)
	// UpdateUserRoles updates a user's roles in the system. A user can have "system_user" and "system_admin" roles.
	UpdateUserRoles(userId string, roles string) (bool, *model.Response)
	// UpdateUserActive updates status of a user whether active or not.
	UpdateUserActive(userId string, active bool) (bool, *model.Response)
	// DeleteUser deactivates a user in the system based on the provided user id string.
	DeleteUser(userId string) (bool, *model.Response)
	// PermanentDeleteUser deletes a user in the system based on the provided user id string.
	PermanentDeleteUser(userId string) (bool, *model.Response)
	// ConvertUserToBot converts a user to a bot user.
	ConvertUserToBot(userId string) (*model.Bot, *model.Response)
	// ConvertBotToUser converts a bot user to a user.
	ConvertBotToUser(userId string, userPatch *model.UserPatch, setSystemAdmin bool) (*model.User, *model.Response)
	// PermanentDeleteAll permanently deletes all users in the system. This is a local only endpoint
	PermanentDeleteAllUsers() (bool, *model.Response)
	// SendPasswordResetEmail will send a link for password resetting to a user with the
	// provided email.
	SendPasswordResetEmail(email string) (bool, *model.Response)
	// ResetPassword uses a recovery code to update reset a user's password.
	ResetPassword(token string, newPassword string) (bool, *model.Response)
	// GetSessions returns a list of sessions based on the provided user id string.
	GetSessions(userId string, etag string) ([]*model.Session, *model.Response)
	// RevokeSession revokes a user session based on the provided user id and session id strings.
	RevokeSession(userId string, sessionId string) (bool, *model.Response)
	// RevokeAllSessions revokes all sessions for the provided user id string.
	RevokeAllSessions(userId string) (bool, *model.Response)
	// RevokeAllSessions revokes all sessions for all the users.
	RevokeSessionsFromAllUsers() (bool, *model.Response)
	// AttachDeviceId attaches a mobile device ID to the current session.
	AttachDeviceId(deviceId string) (bool, *model.Response)
	// GetTeamsUnreadForUser will return an array with TeamUnread objects that contain the amount
	// of unread messages and mentions the current user has for the teams it belongs to.
	// An optional team ID can be set to exclude that team from the results.
	// An optional boolean can be set to include collapsed thread unreads. Must be authenticated.
	GetTeamsUnreadForUser(userId string, teamIdToExclude string, includeCollapsedThreads bool) ([]*model.TeamUnread, *model.Response)
	// GetUserAudits returns a list of audit based on the provided user id string.
	GetUserAudits(userId string, page int, perPage int, etag string) (model.Audits, *model.Response)
	// VerifyUserEmail will verify a user's email using the supplied token.
	VerifyUserEmail(token string) (bool, *model.Response)
	// VerifyUserEmailWithoutToken will verify a user's email by its Id. (Requires manage system role)
	VerifyUserEmailWithoutToken(userId string) (*model.User, *model.Response)
	// SendVerificationEmail will send an email to the user with the provided email address, if
	// that user exists. The email will contain a link that can be used to verify the user's
	// email address.
	SendVerificationEmail(email string) (bool, *model.Response)
	// SetDefaultProfileImage resets the profile image to a default generated one.
	SetDefaultProfileImage(userId string) (bool, *model.Response)
	// SetProfileImage sets profile image of the user.
	SetProfileImage(userId string, data []byte) (bool, *model.Response)
	// CreateUserAccessToken will generate a user access token that can be used in place
	// of a session token to access the REST API. Must have the 'create_user_access_token'
	// permission and if generating for another user, must have the 'edit_other_users'
	// permission. A non-blank description is required.
	CreateUserAccessToken(userId string, description string) (*model.UserAccessToken, *model.Response)
	// GetUserAccessTokens will get a page of access tokens' id, description, is_active
	// and the user_id in the system. The actual token will not be returned. Must have
	// the 'manage_system' permission.
	GetUserAccessTokens(page int, perPage int) ([]*model.UserAccessToken, *model.Response)
	// GetUserAccessToken will get a user access tokens' id, description, is_active
	// and the user_id of the user it is for. The actual token will not be returned.
	// Must have the 'read_user_access_token' permission and if getting for another
	// user, must have the 'edit_other_users' permission.
	GetUserAccessToken(tokenId string) (*model.UserAccessToken, *model.Response)
	// GetUserAccessTokensForUser will get a paged list of user access tokens showing id,
	// description and user_id for each. The actual tokens will not be returned. Must have
	// the 'read_user_access_token' permission and if getting for another user, must have the
	// 'edit_other_users' permission.
	GetUserAccessTokensForUser(userId string, page int, perPage int) ([]*model.UserAccessToken, *model.Response)
	// RevokeUserAccessToken will revoke a user access token by id. Must have the
	// 'revoke_user_access_token' permission and if revoking for another user, must have the
	// 'edit_other_users' permission.
	RevokeUserAccessToken(tokenId string) (bool, *model.Response)
	// SearchUserAccessTokens returns user access tokens matching the provided search term.
	SearchUserAccessTokens(search *model.UserAccessTokenSearch) ([]*model.UserAccessToken, *model.Response)
	// DisableUserAccessToken will disable a user access token by id. Must have the
	// 'revoke_user_access_token' permission and if disabling for another user, must have the
	// 'edit_other_users' permission.
	DisableUserAccessToken(tokenId string) (bool, *model.Response)
	// EnableUserAccessToken will enable a user access token by id. Must have the
	// 'create_user_access_token' permission and if enabling for another user, must have the
	// 'edit_other_users' permission.
	EnableUserAccessToken(tokenId string) (bool, *model.Response)
	// CreateBot creates a bot in the system based on the provided bot struct.
	CreateBot(bot *model.Bot) (*model.Bot, *model.Response)
	// PatchBot partially updates a bot. Any missing fields are not updated.
	PatchBot(userId string, patch *model.BotPatch) (*model.Bot, *model.Response)
	// GetBot fetches the given, undeleted bot.
	GetBot(userId string, etag string) (*model.Bot, *model.Response)
	// GetBot fetches the given bot, even if it is deleted.
	GetBotIncludeDeleted(userId string, etag string) (*model.Bot, *model.Response)
	// GetBots fetches the given page of bots, excluding deleted.
	GetBots(page int, perPage int, etag string) ([]*model.Bot, *model.Response)
	// GetBotsIncludeDeleted fetches the given page of bots, including deleted.
	GetBotsIncludeDeleted(page int, perPage int, etag string) ([]*model.Bot, *model.Response)
	// GetBotsOrphaned fetches the given page of bots, only including orphanded bots.
	GetBotsOrphaned(page int, perPage int, etag string) ([]*model.Bot, *model.Response)
	// DisableBot disables the given bot in the system.
	DisableBot(botUserId string) (*model.Bot, *model.Response)
	// EnableBot disables the given bot in the system.
	EnableBot(botUserId string) (*model.Bot, *model.Response)
	// AssignBot assigns the given bot to the given user
	AssignBot(botUserId string, newOwnerId string) (*model.Bot, *model.Response)
	// SetBotIconImage sets LHS bot icon image.
	SetBotIconImage(botUserId string, data []byte) (bool, *model.Response)
	// GetBotIconImage gets LHS bot icon image. Must be logged in.
	GetBotIconImage(botUserId string) ([]byte, *model.Response)
	// DeleteBotIconImage deletes LHS bot icon image. Must be logged in.
	DeleteBotIconImage(botUserId string) (bool, *model.Response)
	// CreateTeam creates a team in the system based on the provided team struct.
	CreateTeam(team *model.Team) (*model.Team, *model.Response)
	// GetTeam returns a team based on the provided team id string.
	GetTeam(teamId string, etag string) (*model.Team, *model.Response)
	// GetAllTeams returns all teams based on permissions.
	GetAllTeams(etag string, page int, perPage int) ([]*model.Team, *model.Response)
	// GetAllTeamsWithTotalCount returns all teams based on permissions.
	GetAllTeamsWithTotalCount(etag string, page int, perPage int) ([]*model.Team, int64, *model.Response)
	// GetAllTeamsExcludePolicyConstrained returns all teams which are not part of a data retention policy.
	// Must be a system administrator.
	GetAllTeamsExcludePolicyConstrained(etag string, page int, perPage int) ([]*model.Team, *model.Response)
	// GetTeamByName returns a team based on the provided team name string.
	GetTeamByName(name string, etag string) (*model.Team, *model.Response)
	// SearchTeams returns teams matching the provided search term.
	SearchTeams(search *model.TeamSearch) ([]*model.Team, *model.Response)
	// SearchTeamsPaged returns a page of teams and the total count matching the provided search term.
	SearchTeamsPaged(search *model.TeamSearch) ([]*model.Team, int64, *model.Response)
	// TeamExists returns true or false if the team exist or not.
	TeamExists(name string, etag string) (bool, *model.Response)
	// GetTeamsForUser returns a list of teams a user is on. Must be logged in as the user
	// or be a system administrator.
	GetTeamsForUser(userId string, etag string) ([]*model.Team, *model.Response)
	// GetTeamMember returns a team member based on the provided team and user id strings.
	GetTeamMember(teamId string, userId string, etag string) (*model.TeamMember, *model.Response)
	// UpdateTeamMemberRoles will update the roles on a team for a user.
	UpdateTeamMemberRoles(teamId string, userId string, newRoles string) (bool, *model.Response)
	// UpdateTeamMemberSchemeRoles will update the scheme-derived roles on a team for a user.
	UpdateTeamMemberSchemeRoles(teamId string, userId string, schemeRoles *model.SchemeRoles) (bool, *model.Response)
	// UpdateTeam will update a team.
	UpdateTeam(team *model.Team) (*model.Team, *model.Response)
	// PatchTeam partially updates a team. Any missing fields are not updated.
	PatchTeam(teamId string, patch *model.TeamPatch) (*model.Team, *model.Response)
	// RestoreTeam restores a previously deleted team.
	RestoreTeam(teamId string) (*model.Team, *model.Response)
	// RegenerateTeamInviteId requests a new invite ID to be generated.
	RegenerateTeamInviteId(teamId string) (*model.Team, *model.Response)
	// SoftDeleteTeam deletes the team softly (archive only, not permanent delete).
	SoftDeleteTeam(teamId string) (bool, *model.Response)
	// PermanentDeleteTeam deletes the team, should only be used when needed for
	// compliance and the like.
	PermanentDeleteTeam(teamId string) (bool, *model.Response)
	// UpdateTeamPrivacy modifies the team type (model.TEAM_OPEN <--> model.TEAM_INVITE) and sets
	// the corresponding AllowOpenInvite appropriately.
	UpdateTeamPrivacy(teamId string, privacy string) (*model.Team, *model.Response)
	// GetTeamMembers returns team members based on the provided team id string.
	GetTeamMembers(teamId string, page int, perPage int, etag string) ([]*model.TeamMember, *model.Response)
	// GetTeamMembersWithoutDeletedUsers returns team members based on the provided team id string. Additional parameters of sort and exclude_deleted_users accepted as well
	// Could not add it to above function due to it be a breaking change.
	GetTeamMembersSortAndWithoutDeletedUsers(teamId string, page int, perPage int, sort string, exclude_deleted_users bool, etag string) ([]*model.TeamMember, *model.Response)
	// GetTeamMembersForUser returns the team members for a user.
	GetTeamMembersForUser(userId string, etag string) ([]*model.TeamMember, *model.Response)
	// GetTeamMembersByIds will return an array of team members based on the
	// team id and a list of user ids provided. Must be authenticated.
	GetTeamMembersByIds(teamId string, userIds []string) ([]*model.TeamMember, *model.Response)
	// AddTeamMember adds user to a team and return a team member.
	AddTeamMember(teamId string, userId string) (*model.TeamMember, *model.Response)
	// AddTeamMemberFromInvite adds a user to a team and return a team member using an invite id
	// or an invite token/data pair.
	AddTeamMemberFromInvite(token string, inviteId string) (*model.TeamMember, *model.Response)
	// AddTeamMembers adds a number of users to a team and returns the team members.
	AddTeamMembers(teamId string, userIds []string) ([]*model.TeamMember, *model.Response)
	// AddTeamMembers adds a number of users to a team and returns the team members.
	AddTeamMembersGracefully(teamId string, userIds []string) ([]*model.TeamMemberWithError, *model.Response)
	// RemoveTeamMember will remove a user from a team.
	RemoveTeamMember(teamId string, userId string) (bool, *model.Response)
	// GetTeamStats returns a team stats based on the team id string.
	// Must be authenticated.
	GetTeamStats(teamId string, etag string) (*model.TeamStats, *model.Response)
	// GetTotalUsersStats returns a total system user stats.
	// Must be authenticated.
	GetTotalUsersStats(etag string) (*model.UsersStats, *model.Response)
	// GetTeamUnread will return a TeamUnread object that contains the amount of
	// unread messages and mentions the user has for the specified team.
	// Must be authenticated.
	GetTeamUnread(teamId string, userId string) (*model.TeamUnread, *model.Response)
	// ImportTeam will import an exported team from other app into a existing team.
	ImportTeam(data []byte, filesize int, importFrom string, filename string, teamId string) (map[string]string, *model.Response)
	// InviteUsersToTeam invite users by email to the team.
	InviteUsersToTeam(teamId string, userEmails []string) (bool, *model.Response)
	// InviteGuestsToTeam invite guest by email to some channels in a team.
	InviteGuestsToTeam(teamId string, userEmails []string, channels []string, message string) (bool, *model.Response)
	// InviteUsersToTeam invite users by email to the team.
	InviteUsersToTeamGracefully(teamId string, userEmails []string) ([]*model.EmailInviteWithError, *model.Response)
	// InviteGuestsToTeam invite guest by email to some channels in a team.
	InviteGuestsToTeamGracefully(teamId string, userEmails []string, channels []string, message string) ([]*model.EmailInviteWithError, *model.Response)
	// InvalidateEmailInvites will invalidate active email invitations that have not been accepted by the user.
	InvalidateEmailInvites() (bool, *model.Response)
	// GetTeamInviteInfo returns a team object from an invite id containing sanitized information.
	GetTeamInviteInfo(inviteId string) (*model.Team, *model.Response)
	// SetTeamIcon sets team icon of the team.
	SetTeamIcon(teamId string, data []byte) (bool, *model.Response)
	// GetTeamIcon gets the team icon of the team.
	GetTeamIcon(teamId string, etag string) ([]byte, *model.Response)
	// RemoveTeamIcon updates LastTeamIconUpdate to 0 which indicates team icon is removed.
	RemoveTeamIcon(teamId string) (bool, *model.Response)
	// GetAllChannels get all the channels. Must be a system administrator.
	GetAllChannels(page int, perPage int, etag string) (*model.ChannelListWithTeamData, *model.Response)
	// GetAllChannelsIncludeDeleted get all the channels. Must be a system administrator.
	GetAllChannelsIncludeDeleted(page int, perPage int, etag string) (*model.ChannelListWithTeamData, *model.Response)
	// GetAllChannelsExcludePolicyConstrained gets all channels which are not part of a data retention policy.
	// Must be a system administrator.
	GetAllChannelsExcludePolicyConstrained(page int, perPage int, etag string) (*model.ChannelListWithTeamData, *model.Response)
	// GetAllChannelsWithCount get all the channels including the total count. Must be a system administrator.
	GetAllChannelsWithCount(page int, perPage int, etag string) (*model.ChannelListWithTeamData, int64, *model.Response)
	// CreateChannel creates a channel based on the provided channel struct.
	CreateChannel(channel *model.Channel) (*model.Channel, *model.Response)
	// UpdateChannel updates a channel based on the provided channel struct.
	UpdateChannel(channel *model.Channel) (*model.Channel, *model.Response)
	// PatchChannel partially updates a channel. Any missing fields are not updated.
	PatchChannel(channelId string, patch *model.ChannelPatch) (*model.Channel, *model.Response)
	// ConvertChannelToPrivate converts public to private channel.
	ConvertChannelToPrivate(channelId string) (*model.Channel, *model.Response)
	// UpdateChannelPrivacy updates channel privacy
	UpdateChannelPrivacy(channelId string, privacy string) (*model.Channel, *model.Response)
	// RestoreChannel restores a previously deleted channel. Any missing fields are not updated.
	RestoreChannel(channelId string) (*model.Channel, *model.Response)
	// CreateDirectChannel creates a direct message channel based on the two user
	// ids provided.
	CreateDirectChannel(userId1 string, userId2 string) (*model.Channel, *model.Response)
	// CreateGroupChannel creates a group message channel based on userIds provided.
	CreateGroupChannel(userIds []string) (*model.Channel, *model.Response)
	// GetChannel returns a channel based on the provided channel id string.
	GetChannel(channelId string, etag string) (*model.Channel, *model.Response)
	// GetChannelStats returns statistics for a channel.
	GetChannelStats(channelId string, etag string) (*model.ChannelStats, *model.Response)
	// GetChannelMembersTimezones gets a list of timezones for a channel.
	GetChannelMembersTimezones(channelId string) ([]string, *model.Response)
	// GetPinnedPosts gets a list of pinned posts.
	GetPinnedPosts(channelId string, etag string) (*model.PostList, *model.Response)
	// GetPrivateChannelsForTeam returns a list of private channels based on the provided team id string.
	GetPrivateChannelsForTeam(teamId string, page int, perPage int, etag string) ([]*model.Channel, *model.Response)
	// GetPublicChannelsForTeam returns a list of public channels based on the provided team id string.
	GetPublicChannelsForTeam(teamId string, page int, perPage int, etag string) ([]*model.Channel, *model.Response)
	// GetDeletedChannelsForTeam returns a list of public channels based on the provided team id string.
	GetDeletedChannelsForTeam(teamId string, page int, perPage int, etag string) ([]*model.Channel, *model.Response)
	// GetPublicChannelsByIdsForTeam returns a list of public channels based on provided team id string.
	GetPublicChannelsByIdsForTeam(teamId string, channelIds []string) ([]*model.Channel, *model.Response)
	// GetChannelsForTeamForUser returns a list channels of on a team for a user.
	GetChannelsForTeamForUser(teamId string, userId string, includeDeleted bool, etag string) ([]*model.Channel, *model.Response)
	// GetChannelsForTeamAndUserWithLastDeleteAt returns a list channels of a team for a user, additionally filtered with lastDeleteAt. This does not have any effect if includeDeleted is set to false.
	GetChannelsForTeamAndUserWithLastDeleteAt(teamId string, userId string, includeDeleted bool, lastDeleteAt int, etag string) ([]*model.Channel, *model.Response)
	// SearchChannels returns the channels on a team matching the provided search term.
	SearchChannels(teamId string, search *model.ChannelSearch) ([]*model.Channel, *model.Response)
	// SearchArchivedChannels returns the archived channels on a team matching the provided search term.
	SearchArchivedChannels(teamId string, search *model.ChannelSearch) ([]*model.Channel, *model.Response)
	// SearchAllChannels search in all the channels. Must be a system administrator.
	SearchAllChannels(search *model.ChannelSearch) (*model.ChannelListWithTeamData, *model.Response)
	// SearchAllChannelsPaged searches all the channels and returns the results paged with the total count.
	SearchAllChannelsPaged(search *model.ChannelSearch) (*model.ChannelsWithCount, *model.Response)
	// SearchGroupChannels returns the group channels of the user whose members' usernames match the search term.
	SearchGroupChannels(search *model.ChannelSearch) ([]*model.Channel, *model.Response)
	// DeleteChannel deletes channel based on the provided channel id string.
	DeleteChannel(channelId string) (bool, *model.Response)
	// PermanentDeleteChannel deletes a channel based on the provided channel id string.
	PermanentDeleteChannel(channelId string) (bool, *model.Response)
	// MoveChannel moves the channel to the destination team.
	MoveChannel(channelId string, teamId string, force bool) (*model.Channel, *model.Response)
	// GetChannelByName returns a channel based on the provided channel name and team id strings.
	GetChannelByName(channelName string, teamId string, etag string) (*model.Channel, *model.Response)
	// GetChannelByNameIncludeDeleted returns a channel based on the provided channel name and team id strings. Other then GetChannelByName it will also return deleted channels.
	GetChannelByNameIncludeDeleted(channelName string, teamId string, etag string) (*model.Channel, *model.Response)
	// GetChannelByNameForTeamName returns a channel based on the provided channel name and team name strings.
	GetChannelByNameForTeamName(channelName string, teamName string, etag string) (*model.Channel, *model.Response)
	// GetChannelByNameForTeamNameIncludeDeleted returns a channel based on the provided channel name and team name strings. Other then GetChannelByNameForTeamName it will also return deleted channels.
	GetChannelByNameForTeamNameIncludeDeleted(channelName string, teamName string, etag string) (*model.Channel, *model.Response)
	// GetChannelMembers gets a page of channel members.
	GetChannelMembers(channelId string, page int, perPage int, etag string) (*model.ChannelMembers, *model.Response)
	// GetChannelMembersByIds gets the channel members in a channel for a list of user ids.
	GetChannelMembersByIds(channelId string, userIds []string) (*model.ChannelMembers, *model.Response)
	// GetChannelMember gets a channel member.
	GetChannelMember(channelId string, userId string, etag string) (*model.ChannelMember, *model.Response)
	// GetChannelMembersForUser gets all the channel members for a user on a team.
	GetChannelMembersForUser(userId string, teamId string, etag string) (*model.ChannelMembers, *model.Response)
	// ViewChannel performs a view action for a user. Synonymous with switching channels or marking channels as read by a user.
	ViewChannel(userId string, view *model.ChannelView) (*model.ChannelViewResponse, *model.Response)
	// GetChannelUnread will return a ChannelUnread object that contains the number of
	// unread messages and mentions for a user.
	GetChannelUnread(channelId string, userId string) (*model.ChannelUnread, *model.Response)
	// UpdateChannelRoles will update the roles on a channel for a user.
	UpdateChannelRoles(channelId string, userId string, roles string) (bool, *model.Response)
	// UpdateChannelMemberSchemeRoles will update the scheme-derived roles on a channel for a user.
	UpdateChannelMemberSchemeRoles(channelId string, userId string, schemeRoles *model.SchemeRoles) (bool, *model.Response)
	// UpdateChannelNotifyProps will update the notification properties on a channel for a user.
	UpdateChannelNotifyProps(channelId string, userId string, props map[string]string) (bool, *model.Response)
	// AddChannelMember adds user to channel and return a channel member.
	AddChannelMember(channelId string, userId string) (*model.ChannelMember, *model.Response)
	// AddChannelMemberWithRootId adds user to channel and return a channel member. Post add to channel message has the postRootId.
	AddChannelMemberWithRootId(channelId string, userId string, postRootId string) (*model.ChannelMember, *model.Response)
	// RemoveUserFromChannel will delete the channel member object for a user, effectively removing the user from a channel.
	RemoveUserFromChannel(channelId string, userId string) (bool, *model.Response)
	// AutocompleteChannelsForTeam will return an ordered list of channels autocomplete suggestions.
	AutocompleteChannelsForTeam(teamId string, name string) (*model.ChannelList, *model.Response)
	// AutocompleteChannelsForTeamForSearch will return an ordered list of your channels autocomplete suggestions.
	AutocompleteChannelsForTeamForSearch(teamId string, name string) (*model.ChannelList, *model.Response)
	// CreatePost creates a post based on the provided post struct.
	CreatePost(post *model.Post) (*model.Post, *model.Response)
	// CreatePostEphemeral creates a ephemeral post based on the provided post struct which is send to the given user id.
	CreatePostEphemeral(post *model.PostEphemeral) (*model.Post, *model.Response)
	// UpdatePost updates a post based on the provided post struct.
	UpdatePost(postId string, post *model.Post) (*model.Post, *model.Response)
	// PatchPost partially updates a post. Any missing fields are not updated.
	PatchPost(postId string, patch *model.PostPatch) (*model.Post, *model.Response)
	// SetPostUnread marks channel where post belongs as unread on the time of the provided post.
	SetPostUnread(userId string, postId string, collapsedThreadsSupported bool) *model.Response
	// PinPost pin a post based on provided post id string.
	PinPost(postId string) (bool, *model.Response)
	// UnpinPost unpin a post based on provided post id string.
	UnpinPost(postId string) (bool, *model.Response)
	// GetPost gets a single post.
	GetPost(postId string, etag string) (*model.Post, *model.Response)
	// DeletePost deletes a post from the provided post id string.
	DeletePost(postId string) (bool, *model.Response)
	// GetPostThread gets a post with all the other posts in the same thread.
	GetPostThread(postId string, etag string, collapsedThreads bool) (*model.PostList, *model.Response)
	// GetPostsForChannel gets a page of posts with an array for ordering for a channel.
	GetPostsForChannel(channelId string, page int, perPage int, etag string, collapsedThreads bool) (*model.PostList, *model.Response)
	// GetFlaggedPostsForUser returns flagged posts of a user based on user id string.
	GetFlaggedPostsForUser(userId string, page int, perPage int) (*model.PostList, *model.Response)
	// GetFlaggedPostsForUserInTeam returns flagged posts in team of a user based on user id string.
	GetFlaggedPostsForUserInTeam(userId string, teamId string, page int, perPage int) (*model.PostList, *model.Response)
	// GetFlaggedPostsForUserInChannel returns flagged posts in channel of a user based on user id string.
	GetFlaggedPostsForUserInChannel(userId string, channelId string, page int, perPage int) (*model.PostList, *model.Response)
	// GetPostsSince gets posts created after a specified time as Unix time in milliseconds.
	GetPostsSince(channelId string, time int64, collapsedThreads bool) (*model.PostList, *model.Response)
	// GetPostsAfter gets a page of posts that were posted after the post provided.
	GetPostsAfter(channelId string, postId string, page int, perPage int, etag string, collapsedThreads bool) (*model.PostList, *model.Response)
	// GetPostsBefore gets a page of posts that were posted before the post provided.
	GetPostsBefore(channelId string, postId string, page int, perPage int, etag string, collapsedThreads bool) (*model.PostList, *model.Response)
	// GetPostsAroundLastUnread gets a list of posts around last unread post by a user in a channel.
	GetPostsAroundLastUnread(userId string, channelId string, limitBefore int, limitAfter int, collapsedThreads bool) (*model.PostList, *model.Response)
	// SearchFiles returns any posts with matching terms string.
	SearchFiles(teamId string, terms string, isOrSearch bool) (*model.FileInfoList, *model.Response)
	// SearchFilesWithParams returns any posts with matching terms string.
	SearchFilesWithParams(teamId string, params *model.SearchParameter) (*model.FileInfoList, *model.Response)
	// SearchPosts returns any posts with matching terms string.
	SearchPosts(teamId string, terms string, isOrSearch bool) (*model.PostList, *model.Response)
	// SearchPostsWithParams returns any posts with matching terms string.
	SearchPostsWithParams(teamId string, params *model.SearchParameter) (*model.PostList, *model.Response)
	// SearchPostsWithMatches returns any posts with matching terms string, including.
	SearchPostsWithMatches(teamId string, terms string, isOrSearch bool) (*model.PostSearchResults, *model.Response)
	// DoPostAction performs a post action.
	DoPostAction(postId string, actionId string) (bool, *model.Response)
	// DoPostActionWithCookie performs a post action with extra arguments
	DoPostActionWithCookie(postId string, actionId string, selected string, cookieStr string) (bool, *model.Response)
	// OpenInteractiveDialog sends a WebSocket event to a user's clients to
	// open interactive dialogs, based on the provided trigger ID and other
	// provided data. Used with interactive message buttons, menus and
	// slash commands.
	OpenInteractiveDialog(request model.OpenDialogRequest) (bool, *model.Response)
	// SubmitInteractiveDialog will submit the provided dialog data to the integration
	// configured by the URL. Used with the interactive dialogs integration feature.
	SubmitInteractiveDialog(request model.SubmitDialogRequest) (*model.SubmitDialogResponse, *model.Response)
	// UploadFile will upload a file to a channel using a multipart request, to be later attached to a post.
	// This method is functionally equivalent to Client4.UploadFileAsRequestBody.
	UploadFile(data []byte, channelId string, filename string) (*model.FileUploadResponse, *model.Response)
	// UploadFileAsRequestBody will upload a file to a channel as the body of a request, to be later attached
	// to a post. This method is functionally equivalent to Client4.UploadFile.
	UploadFileAsRequestBody(data []byte, channelId string, filename string) (*model.FileUploadResponse, *model.Response)
	// GetFile gets the bytes for a file by id.
	GetFile(fileId string) ([]byte, *model.Response)
	// DownloadFile gets the bytes for a file by id, optionally adding headers to force the browser to download it.
	DownloadFile(fileId string, download bool) ([]byte, *model.Response)
	// GetFileThumbnail gets the bytes for a file by id.
	GetFileThumbnail(fileId string) ([]byte, *model.Response)
	// DownloadFileThumbnail gets the bytes for a file by id, optionally adding headers to force the browser to download it.
	DownloadFileThumbnail(fileId string, download bool) ([]byte, *model.Response)
	// GetFileLink gets the public link of a file by id.
	GetFileLink(fileId string) (string, *model.Response)
	// GetFilePreview gets the bytes for a file by id.
	GetFilePreview(fileId string) ([]byte, *model.Response)
	// DownloadFilePreview gets the bytes for a file by id.
	DownloadFilePreview(fileId string, download bool) ([]byte, *model.Response)
	// GetFileInfo gets all the file info objects.
	GetFileInfo(fileId string) (*model.FileInfo, *model.Response)
	// GetFileInfosForPost gets all the file info objects attached to a post.
	GetFileInfosForPost(postId string, etag string) ([]*model.FileInfo, *model.Response)
	// GenerateSupportPacket downloads the generated support packet
	GenerateSupportPacket() ([]byte, *model.Response)
	// GetPing will return ok if the running goRoutines are below the threshold and unhealthy for above.
	GetPing() (string, *model.Response)
	// GetPingWithServerStatus will return ok if several basic server health checks
	// all pass successfully.
	GetPingWithServerStatus() (string, *model.Response)
	// GetPingWithFullServerStatus will return the full status if several basic server
	// health checks all pass successfully.
	GetPingWithFullServerStatus() (map[string]string, *model.Response)
	// TestEmail will attempt to connect to the configured SMTP server.
	TestEmail(config *model.Config) (bool, *model.Response)
	// TestSiteURL will test the validity of a site URL.
	TestSiteURL(siteURL string) (bool, *model.Response)
	// TestS3Connection will attempt to connect to the AWS S3.
	TestS3Connection(config *model.Config) (bool, *model.Response)
	// GetConfig will retrieve the server config with some sanitized items.
	GetConfig() (*model.Config, *model.Response)
	// ReloadConfig will reload the server configuration.
	ReloadConfig() (bool, *model.Response)
	// GetOldClientConfig will retrieve the parts of the server configuration needed by the
	// client, formatted in the old format.
	GetOldClientConfig(etag string) (map[string]string, *model.Response)
	// GetEnvironmentConfig will retrieve a map mirroring the server configuration where fields
	// are set to true if the corresponding config setting is set through an environment variable.
	// Settings that haven't been set through environment variables will be missing from the map.
	GetEnvironmentConfig() (map[string]interface{}, *model.Response)
	// GetOldClientLicense will retrieve the parts of the server license needed by the
	// client, formatted in the old format.
	GetOldClientLicense(etag string) (map[string]string, *model.Response)
	// DatabaseRecycle will recycle the connections. Discard current connection and get new one.
	DatabaseRecycle() (bool, *model.Response)
	// InvalidateCaches will purge the cache and can affect the performance while is cleaning.
	InvalidateCaches() (bool, *model.Response)
	// UpdateConfig will update the server configuration.
	UpdateConfig(config *model.Config) (*model.Config, *model.Response)
	// MigrateConfig will migrate existing config to the new one.
	MigrateConfig(from string, to string) (bool, *model.Response)
	// UploadLicenseFile will add a license file to the system.
	UploadLicenseFile(data []byte) (bool, *model.Response)
	// RemoveLicenseFile will remove the server license it exists. Note that this will
	// disable all enterprise features.
	RemoveLicenseFile() (bool, *model.Response)
	// GetAnalyticsOld will retrieve analytics using the old format. New format is not
	// available but the "/analytics" endpoint is reserved for it. The "name" argument is optional
	// and defaults to "standard". The "teamId" argument is optional and will limit results
	// to a specific team.
	GetAnalyticsOld(name string, teamId string) (model.AnalyticsRows, *model.Response)
	// CreateIncomingWebhook creates an incoming webhook for a channel.
	CreateIncomingWebhook(hook *model.IncomingWebhook) (*model.IncomingWebhook, *model.Response)
	// UpdateIncomingWebhook updates an incoming webhook for a channel.
	UpdateIncomingWebhook(hook *model.IncomingWebhook) (*model.IncomingWebhook, *model.Response)
	// GetIncomingWebhooks returns a page of incoming webhooks on the system. Page counting starts at 0.
	GetIncomingWebhooks(page int, perPage int, etag string) ([]*model.IncomingWebhook, *model.Response)
	// GetIncomingWebhooksForTeam returns a page of incoming webhooks for a team. Page counting starts at 0.
	GetIncomingWebhooksForTeam(teamId string, page int, perPage int, etag string) ([]*model.IncomingWebhook, *model.Response)
	// GetIncomingWebhook returns an Incoming webhook given the hook ID.
	GetIncomingWebhook(hookID string, etag string) (*model.IncomingWebhook, *model.Response)
	// DeleteIncomingWebhook deletes and Incoming Webhook given the hook ID.
	DeleteIncomingWebhook(hookID string) (bool, *model.Response)
	// CreateOutgoingWebhook creates an outgoing webhook for a team or channel.
	CreateOutgoingWebhook(hook *model.OutgoingWebhook) (*model.OutgoingWebhook, *model.Response)
	// UpdateOutgoingWebhook creates an outgoing webhook for a team or channel.
	UpdateOutgoingWebhook(hook *model.OutgoingWebhook) (*model.OutgoingWebhook, *model.Response)
	// GetOutgoingWebhooks returns a page of outgoing webhooks on the system. Page counting starts at 0.
	GetOutgoingWebhooks(page int, perPage int, etag string) ([]*model.OutgoingWebhook, *model.Response)
	// GetOutgoingWebhook outgoing webhooks on the system requested by Hook Id.
	GetOutgoingWebhook(hookId string) (*model.OutgoingWebhook, *model.Response)
	// GetOutgoingWebhooksForChannel returns a page of outgoing webhooks for a channel. Page counting starts at 0.
	GetOutgoingWebhooksForChannel(channelId string, page int, perPage int, etag string) ([]*model.OutgoingWebhook, *model.Response)
	// GetOutgoingWebhooksForTeam returns a page of outgoing webhooks for a team. Page counting starts at 0.
	GetOutgoingWebhooksForTeam(teamId string, page int, perPage int, etag string) ([]*model.OutgoingWebhook, *model.Response)
	// RegenOutgoingHookToken regenerate the outgoing webhook token.
	RegenOutgoingHookToken(hookId string) (*model.OutgoingWebhook, *model.Response)
	// DeleteOutgoingWebhook delete the outgoing webhook on the system requested by Hook Id.
	DeleteOutgoingWebhook(hookId string) (bool, *model.Response)
	// GetPreferences returns the user's preferences.
	GetPreferences(userId string) (model.Preferences, *model.Response)
	// UpdatePreferences saves the user's preferences.
	UpdatePreferences(userId string, preferences *model.Preferences) (bool, *model.Response)
	// DeletePreferences deletes the user's preferences.
	DeletePreferences(userId string, preferences *model.Preferences) (bool, *model.Response)
	// GetPreferencesByCategory returns the user's preferences from the provided category string.
	GetPreferencesByCategory(userId string, category string) (model.Preferences, *model.Response)
	// GetPreferenceByCategoryAndName returns the user's preferences from the provided category and preference name string.
	GetPreferenceByCategoryAndName(userId string, category string, preferenceName string) (*model.Preference, *model.Response)
	// GetSamlMetadata returns metadata for the SAML configuration.
	GetSamlMetadata() (string, *model.Response)
	// UploadSamlIdpCertificate will upload an IDP certificate for SAML and set the config to use it.
	// The filename parameter is deprecated and ignored: the server will pick a hard-coded filename when writing to disk.
	UploadSamlIdpCertificate(data []byte, filename string) (bool, *model.Response)
	// UploadSamlPublicCertificate will upload a public certificate for SAML and set the config to use it.
	// The filename parameter is deprecated and ignored: the server will pick a hard-coded filename when writing to disk.
	UploadSamlPublicCertificate(data []byte, filename string) (bool, *model.Response)
	// UploadSamlPrivateCertificate will upload a private key for SAML and set the config to use it.
	// The filename parameter is deprecated and ignored: the server will pick a hard-coded filename when writing to disk.
	UploadSamlPrivateCertificate(data []byte, filename string) (bool, *model.Response)
	// DeleteSamlIdpCertificate deletes the SAML IDP certificate from the server and updates the config to not use it and disable SAML.
	DeleteSamlIdpCertificate() (bool, *model.Response)
	// DeleteSamlPublicCertificate deletes the SAML IDP certificate from the server and updates the config to not use it and disable SAML.
	DeleteSamlPublicCertificate() (bool, *model.Response)
	// DeleteSamlPrivateCertificate deletes the SAML IDP certificate from the server and updates the config to not use it and disable SAML.
	DeleteSamlPrivateCertificate() (bool, *model.Response)
	// GetSamlCertificateStatus returns metadata for the SAML configuration.
	GetSamlCertificateStatus() (*model.SamlCertificateStatus, *model.Response)
	GetSamlMetadataFromIdp(samlMetadataURL string) (*model.SamlMetadataResponse, *model.Response)
	// ResetSamlAuthDataToEmail resets the AuthData field of SAML users to their Email.
	ResetSamlAuthDataToEmail(includeDeleted bool, dryRun bool, userIDs []string) (int64, *model.Response)
	// CreateComplianceReport creates an incoming webhook for a channel.
	CreateComplianceReport(report *model.Compliance) (*model.Compliance, *model.Response)
	// GetComplianceReports returns list of compliance reports.
	GetComplianceReports(page int, perPage int) (model.Compliances, *model.Response)
	// GetComplianceReport returns a compliance report.
	GetComplianceReport(reportId string) (*model.Compliance, *model.Response)
	// DownloadComplianceReport returns a full compliance report as a file.
	DownloadComplianceReport(reportId string) ([]byte, *model.Response)
	// GetClusterStatus returns the status of all the configured cluster nodes.
	GetClusterStatus() ([]*model.ClusterInfo, *model.Response)
	// SyncLdap will force a sync with the configured LDAP server.
	// If includeRemovedMembers is true, then group members who left or were removed from a
	// synced team/channel will be re-joined; otherwise, they will be excluded.
	SyncLdap(includeRemovedMembers bool) (bool, *model.Response)
	// TestLdap will attempt to connect to the configured LDAP server and return OK if configured
	// correctly.
	TestLdap() (bool, *model.Response)
	// GetLdapGroups retrieves the immediate child groups of the given parent group.
	GetLdapGroups() ([]*model.Group, *model.Response)
	// LinkLdapGroup creates or undeletes a Mattermost group and associates it to the given LDAP group DN.
	LinkLdapGroup(dn string) (*model.Group, *model.Response)
	// UnlinkLdapGroup deletes the Mattermost group associated with the given LDAP group DN.
	UnlinkLdapGroup(dn string) (*model.Group, *model.Response)
	// MigrateIdLdap migrates the LDAP enabled users to given attribute
	MigrateIdLdap(toAttribute string) (bool, *model.Response)
	// GetGroupsByChannel retrieves the Mattermost Groups associated with a given channel
	GetGroupsByChannel(channelId string, opts model.GroupSearchOpts) ([]*model.GroupWithSchemeAdmin, int, *model.Response)
	// GetGroupsByTeam retrieves the Mattermost Groups associated with a given team
	GetGroupsByTeam(teamId string, opts model.GroupSearchOpts) ([]*model.GroupWithSchemeAdmin, int, *model.Response)
	// GetGroupsAssociatedToChannelsByTeam retrieves the Mattermost Groups associated with channels in a given team
	GetGroupsAssociatedToChannelsByTeam(teamId string, opts model.GroupSearchOpts) (map[string][]*model.GroupWithSchemeAdmin, *model.Response)
	// GetGroups retrieves Mattermost Groups
	GetGroups(opts model.GroupSearchOpts) ([]*model.Group, *model.Response)
	// GetGroupsByUserId retrieves Mattermost Groups for a user
	GetGroupsByUserId(userId string) ([]*model.Group, *model.Response)
	MigrateAuthToLdap(fromAuthService string, matchField string, force bool) (bool, *model.Response)
	MigrateAuthToSaml(fromAuthService string, usersMap map[string]string, auto bool) (bool, *model.Response)
	// UploadLdapPublicCertificate will upload a public certificate for LDAP and set the config to use it.
	UploadLdapPublicCertificate(data []byte) (bool, *model.Response)
	// UploadLdapPrivateCertificate will upload a private key for LDAP and set the config to use it.
	UploadLdapPrivateCertificate(data []byte) (bool, *model.Response)
	// DeleteLdapPublicCertificate deletes the LDAP IDP certificate from the server and updates the config to not use it and disable LDAP.
	DeleteLdapPublicCertificate() (bool, *model.Response)
	// DeleteLDAPPrivateCertificate deletes the LDAP IDP certificate from the server and updates the config to not use it and disable LDAP.
	DeleteLdapPrivateCertificate() (bool, *model.Response)
	// GetAudits returns a list of audits for the whole system.
	GetAudits(page int, perPage int, etag string) (model.Audits, *model.Response)
	// GetBrandImage retrieves the previously uploaded brand image.
	GetBrandImage() ([]byte, *model.Response)
	// DeleteBrandImage deletes the brand image for the system.
	DeleteBrandImage() *model.Response
	// UploadBrandImage sets the brand image for the system.
	UploadBrandImage(data []byte) (bool, *model.Response)
	// GetLogs page of logs as a string array.
	GetLogs(page int, perPage int) ([]string, *model.Response)
	// PostLog is a convenience Web Service call so clients can log messages into
	// the server-side logs. For example we typically log javascript error messages
	// into the server-side. It returns the log message if the logging was successful.
	PostLog(message map[string]string) (map[string]string, *model.Response)
	// CreateOAuthApp will register a new OAuth 2.0 client application with Mattermost acting as an OAuth 2.0 service provider.
	CreateOAuthApp(app *model.OAuthApp) (*model.OAuthApp, *model.Response)
	// UpdateOAuthApp updates a page of registered OAuth 2.0 client applications with Mattermost acting as an OAuth 2.0 service provider.
	UpdateOAuthApp(app *model.OAuthApp) (*model.OAuthApp, *model.Response)
	// GetOAuthApps gets a page of registered OAuth 2.0 client applications with Mattermost acting as an OAuth 2.0 service provider.
	GetOAuthApps(page int, perPage int) ([]*model.OAuthApp, *model.Response)
	// GetOAuthApp gets a registered OAuth 2.0 client application with Mattermost acting as an OAuth 2.0 service provider.
	GetOAuthApp(appId string) (*model.OAuthApp, *model.Response)
	// GetOAuthAppInfo gets a sanitized version of a registered OAuth 2.0 client application with Mattermost acting as an OAuth 2.0 service provider.
	GetOAuthAppInfo(appId string) (*model.OAuthApp, *model.Response)
	// DeleteOAuthApp deletes a registered OAuth 2.0 client application.
	DeleteOAuthApp(appId string) (bool, *model.Response)
	// RegenerateOAuthAppSecret regenerates the client secret for a registered OAuth 2.0 client application.
	RegenerateOAuthAppSecret(appId string) (*model.OAuthApp, *model.Response)
	// GetAuthorizedOAuthAppsForUser gets a page of OAuth 2.0 client applications the user has authorized to use access their account.
	GetAuthorizedOAuthAppsForUser(userId string, page int, perPage int) ([]*model.OAuthApp, *model.Response)
	// AuthorizeOAuthApp will authorize an OAuth 2.0 client application to access a user's account and provide a redirect link to follow.
	AuthorizeOAuthApp(authRequest *model.AuthorizeRequest) (string, *model.Response)
	// DeauthorizeOAuthApp will deauthorize an OAuth 2.0 client application from accessing a user's account.
	DeauthorizeOAuthApp(appId string) (bool, *model.Response)
	// GetOAuthAccessToken is a test helper function for the OAuth access token endpoint.
	GetOAuthAccessToken(data url.Values) (*model.AccessResponse, *model.Response)
	// TestElasticsearch will attempt to connect to the configured Elasticsearch server and return OK if configured.
	// correctly.
	TestElasticsearch() (bool, *model.Response)
	// PurgeElasticsearchIndexes immediately deletes all Elasticsearch indexes.
	PurgeElasticsearchIndexes() (bool, *model.Response)
	// PurgeBleveIndexes immediately deletes all Bleve indexes.
	PurgeBleveIndexes() (bool, *model.Response)
	// GetDataRetentionPolicy will get the current global data retention policy details.
	GetDataRetentionPolicy() (*model.GlobalRetentionPolicy, *model.Response)
	// GetDataRetentionPolicyByID will get the details for the granular data retention policy with the specified ID.
	GetDataRetentionPolicyByID(policyID string) (*model.RetentionPolicyWithTeamAndChannelCounts, *model.Response)
	// GetDataRetentionPoliciesCount will get the total number of granular data retention policies.
	GetDataRetentionPoliciesCount() (int64, *model.Response)
	// GetDataRetentionPolicies will get the current granular data retention policies' details.
	GetDataRetentionPolicies(page int, perPage int) (*model.RetentionPolicyWithTeamAndChannelCountsList, *model.Response)
	// CreateDataRetentionPolicy will create a new granular data retention policy which will be applied to
	// the specified teams and channels. The Id field of `policy` must be empty.
	CreateDataRetentionPolicy(policy *model.RetentionPolicyWithTeamAndChannelIDs) (*model.RetentionPolicyWithTeamAndChannelCounts, *model.Response)
	// DeleteDataRetentionPolicy will delete the granular data retention policy with the specified ID.
	DeleteDataRetentionPolicy(policyID string) *model.Response
	// PatchDataRetentionPolicy will patch the granular data retention policy with the specified ID.
	// The Id field of `patch` must be non-empty.
	PatchDataRetentionPolicy(patch *model.RetentionPolicyWithTeamAndChannelIDs) (*model.RetentionPolicyWithTeamAndChannelCounts, *model.Response)
	// GetTeamsForRetentionPolicy will get the teams to which the specified policy is currently applied.
	GetTeamsForRetentionPolicy(policyID string, page int, perPage int) (*model.TeamsWithCount, *model.Response)
	// SearchTeamsForRetentionPolicy will search the teams to which the specified policy is currently applied.
	SearchTeamsForRetentionPolicy(policyID string, term string) ([]*model.Team, *model.Response)
	// AddTeamsToRetentionPolicy will add the specified teams to the granular data retention policy
	// with the specified ID.
	AddTeamsToRetentionPolicy(policyID string, teamIDs []string) *model.Response
	// RemoveTeamsFromRetentionPolicy will remove the specified teams from the granular data retention policy
	// with the specified ID.
	RemoveTeamsFromRetentionPolicy(policyID string, teamIDs []string) *model.Response
	// GetChannelsForRetentionPolicy will get the channels to which the specified policy is currently applied.
	GetChannelsForRetentionPolicy(policyID string, page int, perPage int) (*model.ChannelsWithCount, *model.Response)
	// SearchChannelsForRetentionPolicy will search the channels to which the specified policy is currently applied.
	SearchChannelsForRetentionPolicy(policyID string, term string) (model.ChannelListWithTeamData, *model.Response)
	// AddChannelsToRetentionPolicy will add the specified channels to the granular data retention policy
	// with the specified ID.
	AddChannelsToRetentionPolicy(policyID string, channelIDs []string) *model.Response
	// RemoveChannelsFromRetentionPolicy will remove the specified channels from the granular data retention policy
	// with the specified ID.
	RemoveChannelsFromRetentionPolicy(policyID string, channelIDs []string) *model.Response
	// GetTeamPoliciesForUser will get the data retention policies for the teams to which a user belongs.
	GetTeamPoliciesForUser(userID string, offset int, limit int) (*model.RetentionPolicyForTeamList, *model.Response)
	// GetChannelPoliciesForUser will get the data retention policies for the channels to which a user belongs.
	GetChannelPoliciesForUser(userID string, offset int, limit int) (*model.RetentionPolicyForChannelList, *model.Response)
	// CreateCommand will create a new command if the user have the right permissions.
	CreateCommand(cmd *model.Command) (*model.Command, *model.Response)
	// UpdateCommand updates a command based on the provided Command struct.
	UpdateCommand(cmd *model.Command) (*model.Command, *model.Response)
	// MoveCommand moves a command to a different team.
	MoveCommand(teamId string, commandId string) (bool, *model.Response)
	// DeleteCommand deletes a command based on the provided command id string.
	DeleteCommand(commandId string) (bool, *model.Response)
	// ListCommands will retrieve a list of commands available in the team.
	ListCommands(teamId string, customOnly bool) ([]*model.Command, *model.Response)
	// ListCommandAutocompleteSuggestions will retrieve a list of suggestions for a userInput.
	ListCommandAutocompleteSuggestions(userInput string, teamId string) ([]model.AutocompleteSuggestion, *model.Response)
	// GetCommandById will retrieve a command by id.
	GetCommandById(cmdId string) (*model.Command, *model.Response)
	// ExecuteCommand executes a given slash command.
	ExecuteCommand(channelId string, command string) (*model.CommandResponse, *model.Response)
	// ExecuteCommandWithTeam executes a given slash command against the specified team.
	// Use this when executing slash commands in a DM/GM, since the team id cannot be inferred in that case.
	ExecuteCommandWithTeam(channelId string, teamId string, command string) (*model.CommandResponse, *model.Response)
	// ListAutocompleteCommands will retrieve a list of commands available in the team.
	ListAutocompleteCommands(teamId string) ([]*model.Command, *model.Response)
	// RegenCommandToken will create a new token if the user have the right permissions.
	RegenCommandToken(commandId string) (string, *model.Response)
	// GetUserStatus returns a user based on the provided user id string.
	GetUserStatus(userId string, etag string) (*model.Status, *model.Response)
	// GetUsersStatusesByIds returns a list of users status based on the provided user ids.
	GetUsersStatusesByIds(userIds []string) ([]*model.Status, *model.Response)
	// UpdateUserStatus sets a user's status based on the provided user id string.
	UpdateUserStatus(userId string, userStatus *model.Status) (*model.Status, *model.Response)
	// CreateEmoji will save an emoji to the server if the current user has permission
	// to do so. If successful, the provided emoji will be returned with its Id field
	// filled in. Otherwise, an error will be returned.
	CreateEmoji(emoji *model.Emoji, image []byte, filename string) (*model.Emoji, *model.Response)
	// GetEmojiList returns a page of custom emoji on the system.
	GetEmojiList(page int, perPage int) ([]*model.Emoji, *model.Response)
	// GetSortedEmojiList returns a page of custom emoji on the system sorted based on the sort
	// parameter, blank for no sorting and "name" to sort by emoji names.
	GetSortedEmojiList(page int, perPage int, sort string) ([]*model.Emoji, *model.Response)
	// DeleteEmoji delete an custom emoji on the provided emoji id string.
	DeleteEmoji(emojiId string) (bool, *model.Response)
	// GetEmoji returns a custom emoji based on the emojiId string.
	GetEmoji(emojiId string) (*model.Emoji, *model.Response)
	// GetEmojiByName returns a custom emoji based on the name string.
	GetEmojiByName(name string) (*model.Emoji, *model.Response)
	// GetEmojiImage returns the emoji image.
	GetEmojiImage(emojiId string) ([]byte, *model.Response)
	// SearchEmoji returns a list of emoji matching some search criteria.
	SearchEmoji(search *model.EmojiSearch) ([]*model.Emoji, *model.Response)
	// AutocompleteEmoji returns a list of emoji starting with or matching name.
	AutocompleteEmoji(name string, etag string) ([]*model.Emoji, *model.Response)
	// SaveReaction saves an emoji reaction for a post. Returns the saved reaction if successful, otherwise an error will be returned.
	SaveReaction(reaction *model.Reaction) (*model.Reaction, *model.Response)
	// GetReactions returns a list of reactions to a post.
	GetReactions(postId string) ([]*model.Reaction, *model.Response)
	// DeleteReaction deletes reaction of a user in a post.
	DeleteReaction(reaction *model.Reaction) (bool, *model.Response)
	// FetchBulkReactions returns a map of postIds and corresponding reactions
	GetBulkReactions(postIds []string) (map[string][]*model.Reaction, *model.Response)
	// GetSupportedTimezone returns a page of supported timezones on the system.
	GetSupportedTimezone() ([]string, *model.Response)
	// OpenGraph return the open graph metadata for a particular url if the site have the metadata.
	OpenGraph(url string) (map[string]string, *model.Response)
	// GetJob gets a single job.
	GetJob(id string) (*model.Job, *model.Response)
	// GetJobs gets all jobs, sorted with the job that was created most recently first.
	GetJobs(page int, perPage int) ([]*model.Job, *model.Response)
	// GetJobsByType gets all jobs of a given type, sorted with the job that was created most recently first.
	GetJobsByType(jobType string, page int, perPage int) ([]*model.Job, *model.Response)
	// CreateJob creates a job based on the provided job struct.
	CreateJob(job *model.Job) (*model.Job, *model.Response)
	// CancelJob requests the cancellation of the job with the provided Id.
	CancelJob(jobId string) (bool, *model.Response)
	// DownloadJob downloads the results of the job
	DownloadJob(jobId string) ([]byte, *model.Response)
	// GetRole gets a single role by ID.
	GetRole(id string) (*model.Role, *model.Response)
	// GetRoleByName gets a single role by Name.
	GetRoleByName(name string) (*model.Role, *model.Response)
	// GetRolesByNames returns a list of roles based on the provided role names.
	GetRolesByNames(roleNames []string) ([]*model.Role, *model.Response)
	// PatchRole partially updates a role in the system. Any missing fields are not updated.
	PatchRole(roleId string, patch *model.RolePatch) (*model.Role, *model.Response)
	// CreateScheme creates a new Scheme.
	CreateScheme(scheme *model.Scheme) (*model.Scheme, *model.Response)
	// GetScheme gets a single scheme by ID.
	GetScheme(id string) (*model.Scheme, *model.Response)
	// GetSchemes gets all schemes, sorted with the most recently created first, optionally filtered by scope.
	GetSchemes(scope string, page int, perPage int) ([]*model.Scheme, *model.Response)
	// DeleteScheme deletes a single scheme by ID.
	DeleteScheme(id string) (bool, *model.Response)
	// PatchScheme partially updates a scheme in the system. Any missing fields are not updated.
	PatchScheme(id string, patch *model.SchemePatch) (*model.Scheme, *model.Response)
	// GetTeamsForScheme gets the teams using this scheme, sorted alphabetically by display name.
	GetTeamsForScheme(schemeId string, page int, perPage int) ([]*model.Team, *model.Response)
	// GetChannelsForScheme gets the channels using this scheme, sorted alphabetically by display name.
	GetChannelsForScheme(schemeId string, page int, perPage int) (model.ChannelList, *model.Response)
	// UploadPlugin takes an io.Reader stream pointing to the contents of a .tar.gz plugin.
	// WARNING: PLUGINS ARE STILL EXPERIMENTAL. THIS FUNCTION IS SUBJECT TO CHANGE.
	UploadPlugin(file io.Reader) (*model.Manifest, *model.Response)
	UploadPluginForced(file io.Reader) (*model.Manifest, *model.Response)
	InstallPluginFromUrl(downloadUrl string, force bool) (*model.Manifest, *model.Response)
	// InstallMarketplacePlugin will install marketplace plugin.
	// WARNING: PLUGINS ARE STILL EXPERIMENTAL. THIS FUNCTION IS SUBJECT TO CHANGE.
	InstallMarketplacePlugin(request *model.InstallMarketplacePluginRequest) (*model.Manifest, *model.Response)
	// GetPlugins will return a list of plugin manifests for currently active plugins.
	// WARNING: PLUGINS ARE STILL EXPERIMENTAL. THIS FUNCTION IS SUBJECT TO CHANGE.
	GetPlugins() (*model.PluginsResponse, *model.Response)
	// GetPluginStatuses will return the plugins installed on any server in the cluster, for reporting
	// to the administrator via the system console.
	// WARNING: PLUGINS ARE STILL EXPERIMENTAL. THIS FUNCTION IS SUBJECT TO CHANGE.
	GetPluginStatuses() (model.PluginStatuses, *model.Response)
	// RemovePlugin will disable and delete a plugin.
	// WARNING: PLUGINS ARE STILL EXPERIMENTAL. THIS FUNCTION IS SUBJECT TO CHANGE.
	RemovePlugin(id string) (bool, *model.Response)
	// GetWebappPlugins will return a list of plugins that the webapp should download.
	// WARNING: PLUGINS ARE STILL EXPERIMENTAL. THIS FUNCTION IS SUBJECT TO CHANGE.
	GetWebappPlugins() ([]*model.Manifest, *model.Response)
	// EnablePlugin will enable an plugin installed.
	// WARNING: PLUGINS ARE STILL EXPERIMENTAL. THIS FUNCTION IS SUBJECT TO CHANGE.
	EnablePlugin(id string) (bool, *model.Response)
	// DisablePlugin will disable an enabled plugin.
	// WARNING: PLUGINS ARE STILL EXPERIMENTAL. THIS FUNCTION IS SUBJECT TO CHANGE.
	DisablePlugin(id string) (bool, *model.Response)
	// GetMarketplacePlugins will return a list of plugins that an admin can install.
	// WARNING: PLUGINS ARE STILL EXPERIMENTAL. THIS FUNCTION IS SUBJECT TO CHANGE.
	GetMarketplacePlugins(filter *model.MarketplacePluginFilter) ([]*model.MarketplacePlugin, *model.Response)
	// UpdateChannelScheme will update a channel's scheme.
	UpdateChannelScheme(channelId string, schemeId string) (bool, *model.Response)
	// UpdateTeamScheme will update a team's scheme.
	UpdateTeamScheme(teamId string, schemeId string) (bool, *model.Response)
	// GetRedirectLocation retrieves the value of the 'Location' header of an HTTP response for a given URL.
	GetRedirectLocation(urlParam string, etag string) (string, *model.Response)
	// SetServerBusy will mark the server as busy, which disables non-critical services for `secs` seconds.
	SetServerBusy(secs int) (bool, *model.Response)
	// ClearServerBusy will mark the server as not busy.
	ClearServerBusy() (bool, *model.Response)
	// GetServerBusy returns the current ServerBusyState including the time when a server marked busy
	// will automatically have the flag cleared.
	GetServerBusy() (*model.ServerBusyState, *model.Response)
	// GetServerBusyExpires returns the time when a server marked busy
	// will automatically have the flag cleared.
	//
	// Deprecated: Use GetServerBusy instead.
	GetServerBusyExpires() (*time.Time, *model.Response)
	// RegisterTermsOfServiceAction saves action performed by a user against a specific terms of service.
	RegisterTermsOfServiceAction(userId string, termsOfServiceId string, accepted bool) (*bool, *model.Response)
	// GetTermsOfService fetches the latest terms of service
	GetTermsOfService(etag string) (*model.TermsOfService, *model.Response)
	// GetUserTermsOfService fetches user's latest terms of service action if the latest action was for acceptance.
	GetUserTermsOfService(userId string, etag string) (*model.UserTermsOfService, *model.Response)
	// CreateTermsOfService creates new terms of service.
	CreateTermsOfService(text string, userId string) (*model.TermsOfService, *model.Response)
	GetGroup(groupID string, etag string) (*model.Group, *model.Response)
	PatchGroup(groupID string, patch *model.GroupPatch) (*model.Group, *model.Response)
	LinkGroupSyncable(groupID string, syncableID string, syncableType model.GroupSyncableType, patch *model.GroupSyncablePatch) (*model.GroupSyncable, *model.Response)
	UnlinkGroupSyncable(groupID string, syncableID string, syncableType model.GroupSyncableType) *model.Response
	GetGroupSyncable(groupID string, syncableID string, syncableType model.GroupSyncableType, etag string) (*model.GroupSyncable, *model.Response)
	GetGroupSyncables(groupID string, syncableType model.GroupSyncableType, etag string) ([]*model.GroupSyncable, *model.Response)
	PatchGroupSyncable(groupID string, syncableID string, syncableType model.GroupSyncableType, patch *model.GroupSyncablePatch) (*model.GroupSyncable, *model.Response)
	TeamMembersMinusGroupMembers(teamID string, groupIDs []string, page int, perPage int, etag string) ([]*model.UserWithGroups, int64, *model.Response)
	ChannelMembersMinusGroupMembers(channelID string, groupIDs []string, page int, perPage int, etag string) ([]*model.UserWithGroups, int64, *model.Response)
	PatchConfig(config *model.Config) (*model.Config, *model.Response)
	GetChannelModerations(channelID string, etag string) ([]*model.ChannelModeration, *model.Response)
	PatchChannelModerations(channelID string, patch []*model.ChannelModerationPatch) ([]*model.ChannelModeration, *model.Response)
	GetKnownUsers() ([]string, *model.Response)
	// PublishUserTyping publishes a user is typing websocket event based on the provided TypingRequest.
	PublishUserTyping(userID string, typingRequest model.TypingRequest) (bool, *model.Response)
	GetChannelMemberCountsByGroup(channelID string, includeTimezones bool, etag string) ([]*model.ChannelMemberCountByGroup, *model.Response)
	// RequestTrialLicense will request a trial license and install it in the server
	RequestTrialLicense(users int) (bool, *model.Response)
	// GetGroupStats retrieves stats for a Mattermost Group
	GetGroupStats(groupID string) (*model.GroupStats, *model.Response)
	GetSidebarCategoriesForTeamForUser(userID string, teamID string, etag string) (*model.OrderedSidebarCategories, *model.Response)
	CreateSidebarCategoryForTeamForUser(userID string, teamID string, category *model.SidebarCategoryWithChannels) (*model.SidebarCategoryWithChannels, *model.Response)
	UpdateSidebarCategoriesForTeamForUser(userID string, teamID string, categories []*model.SidebarCategoryWithChannels) ([]*model.SidebarCategoryWithChannels, *model.Response)
	GetSidebarCategoryOrderForTeamForUser(userID string, teamID string, etag string) ([]string, *model.Response)
	UpdateSidebarCategoryOrderForTeamForUser(userID string, teamID string, order []string) ([]string, *model.Response)
	GetSidebarCategoryForTeamForUser(userID string, teamID string, categoryID string, etag string) (*model.SidebarCategoryWithChannels, *model.Response)
	UpdateSidebarCategoryForTeamForUser(userID string, teamID string, categoryID string, category *model.SidebarCategoryWithChannels) (*model.SidebarCategoryWithChannels, *model.Response)
	// CheckIntegrity performs a database integrity check.
	CheckIntegrity() ([]model.IntegrityCheckResult, *model.Response)
	GetNotices(lastViewed int64, teamId string, client model.NoticeClientType, clientVersion string, locale string, etag string) (model.NoticeMessages, *model.Response)
	MarkNoticesViewed(ids []string) *model.Response
	// CreateUpload creates a new upload session.
	CreateUpload(us *model.UploadSession) (*model.UploadSession, *model.Response)
	// GetUpload returns the upload session for the specified uploadId.
	GetUpload(uploadId string) (*model.UploadSession, *model.Response)
	// GetUploadsForUser returns the upload sessions created by the specified
	// userId.
	GetUploadsForUser(userId string) ([]*model.UploadSession, *model.Response)
	// UploadData performs an upload. On success it returns
	// a FileInfo object.
	UploadData(uploadId string, data io.Reader) (*model.FileInfo, *model.Response)
	UpdatePassword(userId string, currentPassword string, newPassword string) *model.Response
	GetCloudProducts() ([]*model.Product, *model.Response)
	CreateCustomerPayment() (*model.StripeSetupIntent, *model.Response)
	ConfirmCustomerPayment(confirmRequest *model.ConfirmPaymentMethodRequest) *model.Response
	GetCloudCustomer() (*model.CloudCustomer, *model.Response)
	GetSubscription() (*model.Subscription, *model.Response)
	GetSubscriptionStats() (*model.SubscriptionStats, *model.Response)
	GetInvoicesForSubscription() ([]*model.Invoice, *model.Response)
	UpdateCloudCustomer(customerInfo *model.CloudCustomerInfo) (*model.CloudCustomer, *model.Response)
	UpdateCloudCustomerAddress(address *model.Address) (*model.CloudCustomer, *model.Response)
	ListImports() ([]string, *model.Response)
	ListExports() ([]string, *model.Response)
	DeleteExport(name string) (bool, *model.Response)
	DownloadExport(name string, wr io.Writer, offset int64) (int64, *model.Response)
	GetUserThreads(userId string, teamId string, options model.GetUserThreadsOpts) (*model.Threads, *model.Response)
	GetUserThread(userId string, teamId string, threadId string, extended bool) (*model.ThreadResponse, *model.Response)
	UpdateThreadsReadForUser(userId string, teamId string) *model.Response
	UpdateThreadReadForUser(userId string, teamId string, threadId string, timestamp int64) (*model.ThreadResponse, *model.Response)
	UpdateThreadFollowForUser(userId string, teamId string, threadId string, state bool) *model.Response
	SendAdminUpgradeRequestEmail() *model.Response
	SendAdminUpgradeRequestEmailOnJoin() *model.Response
	GetAllSharedChannels(teamID string, page int, perPage int) ([]*model.SharedChannel, *model.Response)
	GetRemoteClusterInfo(remoteID string) (model.RemoteClusterInfo, *model.Response)
	GetAncillaryPermissions(subsectionPermissions []string) ([]string, *model.Response)
}
