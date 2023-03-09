package generator

import (
	"os"
	"os/exec"
	"path/filepath"
	"testing"

	"github.com/denisdubovitskiy/ifacemaker/gopath"
	"github.com/stretchr/testify/require"
)

var vaultAPIFiles = []string{
	"github.com/hashicorp/vault/api@v1.8.2/auth.go",
	"github.com/hashicorp/vault/api@v1.8.2/auth_token.go",
	"github.com/hashicorp/vault/api@v1.8.2/client.go",
	"github.com/hashicorp/vault/api@v1.8.2/help.go",
	"github.com/hashicorp/vault/api@v1.8.2/kv.go",
	"github.com/hashicorp/vault/api@v1.8.2/kv_v1.go",
	"github.com/hashicorp/vault/api@v1.8.2/kv_v2.go",
	"github.com/hashicorp/vault/api@v1.8.2/lifetime_watcher.go",
	"github.com/hashicorp/vault/api@v1.8.2/logical.go",
	"github.com/hashicorp/vault/api@v1.8.2/output_policy.go",
	"github.com/hashicorp/vault/api@v1.8.2/output_string.go",
	"github.com/hashicorp/vault/api@v1.8.2/plugin_helpers.go",
	"github.com/hashicorp/vault/api@v1.8.2/request.go",
	"github.com/hashicorp/vault/api@v1.8.2/response.go",
	"github.com/hashicorp/vault/api@v1.8.2/secret.go",
	"github.com/hashicorp/vault/api@v1.8.2/ssh.go",
	"github.com/hashicorp/vault/api@v1.8.2/ssh_agent.go",
	"github.com/hashicorp/vault/api@v1.8.2/sys.go",
	"github.com/hashicorp/vault/api@v1.8.2/sys_audit.go",
	"github.com/hashicorp/vault/api@v1.8.2/sys_auth.go",
	"github.com/hashicorp/vault/api@v1.8.2/sys_capabilities.go",
	"github.com/hashicorp/vault/api@v1.8.2/sys_config_cors.go",
	"github.com/hashicorp/vault/api@v1.8.2/sys_generate_root.go",
	"github.com/hashicorp/vault/api@v1.8.2/sys_hastatus.go",
	"github.com/hashicorp/vault/api@v1.8.2/sys_health.go",
	"github.com/hashicorp/vault/api@v1.8.2/sys_init.go",
	"github.com/hashicorp/vault/api@v1.8.2/sys_leader.go",
	"github.com/hashicorp/vault/api@v1.8.2/sys_leases.go",
	"github.com/hashicorp/vault/api@v1.8.2/sys_mfa.go",
	"github.com/hashicorp/vault/api@v1.8.2/sys_monitor.go",
	"github.com/hashicorp/vault/api@v1.8.2/sys_mounts.go",
	"github.com/hashicorp/vault/api@v1.8.2/sys_plugins.go",
	"github.com/hashicorp/vault/api@v1.8.2/sys_policy.go",
	"github.com/hashicorp/vault/api@v1.8.2/sys_raft.go",
	"github.com/hashicorp/vault/api@v1.8.2/sys_rekey.go",
	"github.com/hashicorp/vault/api@v1.8.2/sys_rotate.go",
	"github.com/hashicorp/vault/api@v1.8.2/sys_seal.go",
	"github.com/hashicorp/vault/api@v1.8.2/sys_stepdown.go",
}

var mattermostModelFiles = []string{
	"github.com/mattermost/mattermost-server/v5@v5.39.3/model/access.go",
	"github.com/mattermost/mattermost-server/v5@v5.39.3/model/analytics_row.go",
	"github.com/mattermost/mattermost-server/v5@v5.39.3/model/at_mentions.go",
	"github.com/mattermost/mattermost-server/v5@v5.39.3/model/audit.go",
	"github.com/mattermost/mattermost-server/v5@v5.39.3/model/auditconv.go",
	"github.com/mattermost/mattermost-server/v5@v5.39.3/model/audits.go",
	"github.com/mattermost/mattermost-server/v5@v5.39.3/model/authorize.go",
	"github.com/mattermost/mattermost-server/v5@v5.39.3/model/bot.go",
	"github.com/mattermost/mattermost-server/v5@v5.39.3/model/builtin.go",
	"github.com/mattermost/mattermost-server/v5@v5.39.3/model/bundle_info.go",
	"github.com/mattermost/mattermost-server/v5@v5.39.3/model/channel.go",
	"github.com/mattermost/mattermost-server/v5@v5.39.3/model/channel_count.go",
	"github.com/mattermost/mattermost-server/v5@v5.39.3/model/channel_data.go",
	"github.com/mattermost/mattermost-server/v5@v5.39.3/model/channel_list.go",
	"github.com/mattermost/mattermost-server/v5@v5.39.3/model/channel_member.go",
	"github.com/mattermost/mattermost-server/v5@v5.39.3/model/channel_member_history.go",
	"github.com/mattermost/mattermost-server/v5@v5.39.3/model/channel_member_history_result.go",
	"github.com/mattermost/mattermost-server/v5@v5.39.3/model/channel_mentions.go",
	"github.com/mattermost/mattermost-server/v5@v5.39.3/model/channel_search.go",
	"github.com/mattermost/mattermost-server/v5@v5.39.3/model/channel_sidebar.go",
	"github.com/mattermost/mattermost-server/v5@v5.39.3/model/channel_stats.go",
	"github.com/mattermost/mattermost-server/v5@v5.39.3/model/channel_view.go",
	"github.com/mattermost/mattermost-server/v5@v5.39.3/model/client4.go",
	"github.com/mattermost/mattermost-server/v5@v5.39.3/model/cloud.go",
	"github.com/mattermost/mattermost-server/v5@v5.39.3/model/cluster_discovery.go",
	"github.com/mattermost/mattermost-server/v5@v5.39.3/model/cluster_info.go",
	"github.com/mattermost/mattermost-server/v5@v5.39.3/model/cluster_message.go",
	"github.com/mattermost/mattermost-server/v5@v5.39.3/model/cluster_stats.go",
	"github.com/mattermost/mattermost-server/v5@v5.39.3/model/command.go",
	"github.com/mattermost/mattermost-server/v5@v5.39.3/model/command_args.go",
	"github.com/mattermost/mattermost-server/v5@v5.39.3/model/command_autocomplete.go",
	"github.com/mattermost/mattermost-server/v5@v5.39.3/model/command_request.go",
	"github.com/mattermost/mattermost-server/v5@v5.39.3/model/command_response.go",
	"github.com/mattermost/mattermost-server/v5@v5.39.3/model/command_webhook.go",
	"github.com/mattermost/mattermost-server/v5@v5.39.3/model/compliance.go",
	"github.com/mattermost/mattermost-server/v5@v5.39.3/model/compliance_post.go",
	"github.com/mattermost/mattermost-server/v5@v5.39.3/model/config.go",
	"github.com/mattermost/mattermost-server/v5@v5.39.3/model/custom_status.go",
	"github.com/mattermost/mattermost-server/v5@v5.39.3/model/data_retention_policy.go",
	"github.com/mattermost/mattermost-server/v5@v5.39.3/model/emoji.go",
	"github.com/mattermost/mattermost-server/v5@v5.39.3/model/emoji_data.go",
	"github.com/mattermost/mattermost-server/v5@v5.39.3/model/emoji_search.go",
	"github.com/mattermost/mattermost-server/v5@v5.39.3/model/feature_flags.go",
	"github.com/mattermost/mattermost-server/v5@v5.39.3/model/file.go",
	"github.com/mattermost/mattermost-server/v5@v5.39.3/model/file_info.go",
	"github.com/mattermost/mattermost-server/v5@v5.39.3/model/file_info_list.go",
	"github.com/mattermost/mattermost-server/v5@v5.39.3/model/file_info_search_results.go",
	"github.com/mattermost/mattermost-server/v5@v5.39.3/model/gitlab.go",
	"github.com/mattermost/mattermost-server/v5@v5.39.3/model/group.go",
	"github.com/mattermost/mattermost-server/v5@v5.39.3/model/group_member.go",
	"github.com/mattermost/mattermost-server/v5@v5.39.3/model/group_syncable.go",
	"github.com/mattermost/mattermost-server/v5@v5.39.3/model/guest_invite.go",
	"github.com/mattermost/mattermost-server/v5@v5.39.3/model/incoming_webhook.go",
	"github.com/mattermost/mattermost-server/v5@v5.39.3/model/initial_load.go",
	"github.com/mattermost/mattermost-server/v5@v5.39.3/model/integration_action.go",
	"github.com/mattermost/mattermost-server/v5@v5.39.3/model/integrity.go",
	"github.com/mattermost/mattermost-server/v5@v5.39.3/model/job.go",
	"github.com/mattermost/mattermost-server/v5@v5.39.3/model/ldap.go",
	"github.com/mattermost/mattermost-server/v5@v5.39.3/model/license.go",
	"github.com/mattermost/mattermost-server/v5@v5.39.3/model/link_metadata.go",
	"github.com/mattermost/mattermost-server/v5@v5.39.3/model/manifest.go",
	"github.com/mattermost/mattermost-server/v5@v5.39.3/model/marketplace_plugin.go",
	"github.com/mattermost/mattermost-server/v5@v5.39.3/model/mention_map.go",
	"github.com/mattermost/mattermost-server/v5@v5.39.3/model/message_export.go",
	"github.com/mattermost/mattermost-server/v5@v5.39.3/model/mfa_secret.go",
	"github.com/mattermost/mattermost-server/v5@v5.39.3/model/migration.go",
	"github.com/mattermost/mattermost-server/v5@v5.39.3/model/oauth.go",
	"github.com/mattermost/mattermost-server/v5@v5.39.3/model/outgoing_webhook.go",
	"github.com/mattermost/mattermost-server/v5@v5.39.3/model/permission.go",
	"github.com/mattermost/mattermost-server/v5@v5.39.3/model/plugin_cluster_event.go",
	"github.com/mattermost/mattermost-server/v5@v5.39.3/model/plugin_event_data.go",
	"github.com/mattermost/mattermost-server/v5@v5.39.3/model/plugin_key_value.go",
	"github.com/mattermost/mattermost-server/v5@v5.39.3/model/plugin_kvset_options.go",
	"github.com/mattermost/mattermost-server/v5@v5.39.3/model/plugin_status.go",
	"github.com/mattermost/mattermost-server/v5@v5.39.3/model/plugin_valid.go",
	"github.com/mattermost/mattermost-server/v5@v5.39.3/model/plugins_response.go",
	"github.com/mattermost/mattermost-server/v5@v5.39.3/model/post.go",
	"github.com/mattermost/mattermost-server/v5@v5.39.3/model/post_embed.go",
	"github.com/mattermost/mattermost-server/v5@v5.39.3/model/post_list.go",
	"github.com/mattermost/mattermost-server/v5@v5.39.3/model/post_metadata.go",
	"github.com/mattermost/mattermost-server/v5@v5.39.3/model/post_search_results.go",
	"github.com/mattermost/mattermost-server/v5@v5.39.3/model/preference.go",
	"github.com/mattermost/mattermost-server/v5@v5.39.3/model/preferences.go",
	"github.com/mattermost/mattermost-server/v5@v5.39.3/model/product_notices.go",
	"github.com/mattermost/mattermost-server/v5@v5.39.3/model/push_notification.go",
	"github.com/mattermost/mattermost-server/v5@v5.39.3/model/push_response.go",
	"github.com/mattermost/mattermost-server/v5@v5.39.3/model/reaction.go",
	"github.com/mattermost/mattermost-server/v5@v5.39.3/model/remote_cluster.go",
	"github.com/mattermost/mattermost-server/v5@v5.39.3/model/role.go",
	"github.com/mattermost/mattermost-server/v5@v5.39.3/model/saml.go",
	"github.com/mattermost/mattermost-server/v5@v5.39.3/model/scheduled_task.go",
	"github.com/mattermost/mattermost-server/v5@v5.39.3/model/scheme.go",
	"github.com/mattermost/mattermost-server/v5@v5.39.3/model/search_params.go",
	"github.com/mattermost/mattermost-server/v5@v5.39.3/model/security_bulletin.go",
	"github.com/mattermost/mattermost-server/v5@v5.39.3/model/session.go",
	"github.com/mattermost/mattermost-server/v5@v5.39.3/model/session_serial_gen.go",
	"github.com/mattermost/mattermost-server/v5@v5.39.3/model/shared_channel.go",
	"github.com/mattermost/mattermost-server/v5@v5.39.3/model/slack_attachment.go",
	"github.com/mattermost/mattermost-server/v5@v5.39.3/model/slack_compatibility.go",
	"github.com/mattermost/mattermost-server/v5@v5.39.3/model/status.go",
	"github.com/mattermost/mattermost-server/v5@v5.39.3/model/suggest_command.go",
	"github.com/mattermost/mattermost-server/v5@v5.39.3/model/switch_request.go",
	"github.com/mattermost/mattermost-server/v5@v5.39.3/model/system.go",
	"github.com/mattermost/mattermost-server/v5@v5.39.3/model/team.go",
	"github.com/mattermost/mattermost-server/v5@v5.39.3/model/team_member.go",
	"github.com/mattermost/mattermost-server/v5@v5.39.3/model/team_member_serial_gen.go",
	"github.com/mattermost/mattermost-server/v5@v5.39.3/model/team_search.go",
	"github.com/mattermost/mattermost-server/v5@v5.39.3/model/team_stats.go",
	"github.com/mattermost/mattermost-server/v5@v5.39.3/model/terms_of_service.go",
	"github.com/mattermost/mattermost-server/v5@v5.39.3/model/thread.go",
	"github.com/mattermost/mattermost-server/v5@v5.39.3/model/token.go",
	"github.com/mattermost/mattermost-server/v5@v5.39.3/model/typing_request.go",
	"github.com/mattermost/mattermost-server/v5@v5.39.3/model/upload_session.go",
	"github.com/mattermost/mattermost-server/v5@v5.39.3/model/user.go",
	"github.com/mattermost/mattermost-server/v5@v5.39.3/model/user_access_token.go",
	"github.com/mattermost/mattermost-server/v5@v5.39.3/model/user_access_token_search.go",
	"github.com/mattermost/mattermost-server/v5@v5.39.3/model/user_autocomplete.go",
	"github.com/mattermost/mattermost-server/v5@v5.39.3/model/user_count.go",
	"github.com/mattermost/mattermost-server/v5@v5.39.3/model/user_get.go",
	"github.com/mattermost/mattermost-server/v5@v5.39.3/model/user_search.go",
	"github.com/mattermost/mattermost-server/v5@v5.39.3/model/user_serial_gen.go",
	"github.com/mattermost/mattermost-server/v5@v5.39.3/model/user_terms_of_service.go",
	"github.com/mattermost/mattermost-server/v5@v5.39.3/model/users_stats.go",
	"github.com/mattermost/mattermost-server/v5@v5.39.3/model/utils.go",
	"github.com/mattermost/mattermost-server/v5@v5.39.3/model/version.go",
	"github.com/mattermost/mattermost-server/v5@v5.39.3/model/websocket_client.go",
	"github.com/mattermost/mattermost-server/v5@v5.39.3/model/websocket_message.go",
	"github.com/mattermost/mattermost-server/v5@v5.39.3/model/websocket_request.go",
}

func TestGenerate(t *testing.T) {
	cases := []struct {
		name           string
		module         string
		files          []string
		want           string
		structName     string
		interfaceName  string
		outPackageName string
	}{
		{
			module: "github.com/mattermost/mattermost-server/v5@v5.39.3",
			name:   "mattermost model audit",
			want: `// Package audit generated with github.com/denisdubovitskiy/ifacemaker, DO NOT EDIT.
package audit

type Audit interface {
	ToJson() string
}
`,
			files:          mattermostModelFiles,
			structName:     "Audit",
			interfaceName:  "Audit",
			outPackageName: "audit",
		},
		{
			module: "github.com/mattermost/mattermost-server/v5@v5.39.3",
			name:   "mattermost model audit interface rename",
			want: `// Package audit generated with github.com/denisdubovitskiy/ifacemaker, DO NOT EDIT.
package audit

type Audit2 interface {
	ToJson() string
}
`,
			files:          mattermostModelFiles,
			structName:     "Audit",
			interfaceName:  "Audit2",
			outPackageName: "audit",
		},
		{
			module: "github.com/mattermost/mattermost-server/v5@v5.39.3",
			name:   "mattermost model audit package rename",
			want: `// Package testpackage generated with github.com/denisdubovitskiy/ifacemaker, DO NOT EDIT.
package testpackage

type Audit interface {
	ToJson() string
}
`,
			files:          mattermostModelFiles,
			structName:     "Audit",
			interfaceName:  "Audit",
			outPackageName: "testpackage",
		},
		{
			module: "github.com/mattermost/mattermost-server/v5@v5.39.3",
			name:   "mattermost model user",
			want: `// Package user generated with github.com/denisdubovitskiy/ifacemaker, DO NOT EDIT.
package user

import "github.com/mattermost/mattermost-server/v5/model"

type User interface {
	DeepCopy() *model.User
	// IsValid validates the user and returns an error if it isn't configured
	// correctly.
	IsValid() *model.AppError
	// PreSave will set the Id and Username if missing.  It will also fill
	// in the CreateAt, UpdateAt times.  It will also hash the password.  It should
	// be run before saving the user to the db.
	PreSave()
	// PreUpdate should be run before updating the user in the db.
	PreUpdate()
	SetDefaultNotifications()
	UpdateMentionKeysFromUsername(oldUsername string)
	GetMentionKeys() []string
	Patch(patch *model.UserPatch)
	// ToJson convert a User to a json string
	ToJson() string
	// Generate a valid strong etag so the browser can cache the results
	Etag(showFullName bool, showEmail bool) string
	// Remove any private data from the user object
	Sanitize(options map[string]bool)
	// Remove any input data from the user object that is not user controlled
	SanitizeInput(isAdmin bool)
	ClearNonProfileFields()
	SanitizeProfile(options map[string]bool)
	MakeNonNil()
	AddNotifyProp(key string, value string)
	SetCustomStatus(cs *model.CustomStatus)
	ClearCustomStatus()
	GetFullName() string
	GetDisplayName(nameFormat string) string
	GetDisplayNameWithPrefix(nameFormat string, prefix string) string
	GetRoles() []string
	GetRawRoles() string
	// Make sure you acually want to use this function. In context.go there are functions to check permissions
	// This function should not be used to check permissions.
	IsGuest() bool
	IsSystemAdmin() bool
	// Make sure you acually want to use this function. In context.go there are functions to check permissions
	// This function should not be used to check permissions.
	IsInRole(inRole string) bool
	IsSSOUser() bool
	IsOAuthUser() bool
	IsLDAPUser() bool
	IsSAMLUser() bool
	GetPreferredTimezone() string
	// IsRemote returns true if the user belongs to a remote cluster (has RemoteId).
	IsRemote() bool
	// GetRemoteID returns the remote id for this user or "" if not a remote user.
	GetRemoteID() string
	// GetProp fetches a prop value by name.
	GetProp(name string) (string, bool)
	// SetProp sets a prop value by name, creating the map if nil.
	// Not thread safe.
	SetProp(name string, value string)
	ToPatch() *model.UserPatch
	// DecodeMsg implements msgp.Decodable
	DecodeMsg(dc *msgp.Reader) (err error)
	// EncodeMsg implements msgp.Encodable
	EncodeMsg(en *msgp.Writer) (err error)
	// MarshalMsg implements msgp.Marshaler
	MarshalMsg(b []byte) (o []byte, err error)
	// UnmarshalMsg implements msgp.Unmarshaler
	UnmarshalMsg(bts []byte) (o []byte, err error)
	// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
	Msgsize() (s int)
}
`,
			files:          mattermostModelFiles,
			structName:     "User",
			interfaceName:  "User",
			outPackageName: "user",
		},
		{
			module: "github.com/hashicorp/vault/api@v1.8.2",
			name:   "vault api client",
			want: `// Package vault generated with github.com/denisdubovitskiy/ifacemaker, DO NOT EDIT.
package vault

import (
	"context"
	"net/http"
	"time"

	"github.com/hashicorp/go-retryablehttp"
	"github.com/hashicorp/vault/api"
	"golang.org/x/time/rate"
)

type Client interface {
	// Auth is used to return the client for credential-backend API calls.
	Auth() *api.Auth
	CloneConfig() *api.Config
	// SetAddress sets the address of Vault in the client. The format of address should be
	// "<Scheme>://<Host>:<Port>". Setting this on a client will override the
	// value of VAULT_ADDR environment variable.
	SetAddress(addr string) error
	// Address returns the Vault URL the client is configured to connect to
	Address() string
	SetCheckRedirect(f func(*http.Request, []*http.Request) error)
	// SetLimiter will set the rate limiter for this client.
	// This method is thread-safe.
	// rateLimit and burst are specified according to https://godoc.org/golang.org/x/time/rate#NewLimiter
	SetLimiter(rateLimit float64, burst int)
	Limiter() *rate.Limiter
	// SetMinRetryWait sets the minimum time to wait before retrying in the case of certain errors.
	SetMinRetryWait(retryWait time.Duration)
	MinRetryWait() time.Duration
	// SetMaxRetryWait sets the maximum time to wait before retrying in the case of certain errors.
	SetMaxRetryWait(retryWait time.Duration)
	MaxRetryWait() time.Duration
	// SetMaxRetries sets the number of retries that will be used in the case of certain errors
	SetMaxRetries(retries int)
	SetMaxIdleConnections(idle int)
	MaxIdleConnections() int
	SetDisableKeepAlives(disable bool)
	DisableKeepAlives() bool
	MaxRetries() int
	SetSRVLookup(srv bool)
	SRVLookup() bool
	// SetCheckRetry sets the CheckRetry function to be used for future requests.
	SetCheckRetry(checkRetry retryablehttp.CheckRetry)
	CheckRetry() retryablehttp.CheckRetry
	// SetClientTimeout sets the client request timeout
	SetClientTimeout(timeout time.Duration)
	ClientTimeout() time.Duration
	OutputCurlString() bool
	SetOutputCurlString(curl bool)
	OutputPolicy() bool
	SetOutputPolicy(isSet bool)
	// CurrentWrappingLookupFunc sets a lookup function that returns desired wrap TTLs
	// for a given operation and path.
	CurrentWrappingLookupFunc() api.WrappingLookupFunc
	// SetWrappingLookupFunc sets a lookup function that returns desired wrap TTLs
	// for a given operation and path.
	SetWrappingLookupFunc(lookupFunc api.WrappingLookupFunc)
	// SetMFACreds sets the MFA credentials supplied either via the environment
	// variable or via the command line.
	SetMFACreds(creds []string)
	// SetNamespace sets the namespace supplied either via the environment
	// variable or via the command line.
	SetNamespace(namespace string)
	// ClearNamespace removes the namespace header if set.
	ClearNamespace()
	// Namespace returns the namespace currently set in this client. It will
	// return an empty string if there is no namespace set.
	Namespace() string
	// WithNamespace makes a shallow copy of Client, modifies it to use
	// the given namespace, and returns it. Passing an empty string will
	// temporarily unset the namespace.
	WithNamespace(namespace string) *api.Client
	// Token returns the access token being used by this client. It will
	// return the empty string if there is no token set.
	Token() string
	// SetToken sets the token directly. This won't perform any auth
	// verification, it simply sets the token properly for future requests.
	SetToken(v string)
	// ClearToken deletes the token if it is set or does nothing otherwise.
	ClearToken()
	// Headers gets the current set of headers used for requests. This returns a
	// copy; to modify it call AddHeader or SetHeaders.
	Headers() http.Header
	// AddHeader allows a single header key/value pair to be added
	// in a race-safe fashion.
	AddHeader(key string, value string)
	// SetHeaders clears all previous headers and uses only the given
	// ones going forward.
	SetHeaders(headers http.Header)
	// SetBackoff sets the backoff function to be used for future requests.
	SetBackoff(backoff retryablehttp.Backoff)
	SetLogger(logger retryablehttp.LeveledLogger)
	// SetCloneHeaders to allow headers to be copied whenever the client is cloned.
	SetCloneHeaders(cloneHeaders bool)
	// CloneHeaders gets the configured CloneHeaders value.
	CloneHeaders() bool
	// SetCloneToken from parent
	SetCloneToken(cloneToken bool)
	// CloneToken gets the configured CloneToken value.
	CloneToken() bool
	// SetReadYourWrites to prevent reading stale cluster replication state.
	SetReadYourWrites(preventStaleReads bool)
	// ReadYourWrites gets the configured value of ReadYourWrites
	ReadYourWrites() bool
	// Clone creates a new client with the same configuration. Note that the same
	// underlying http.Client is used; modifying the client from more than one
	// goroutine at once may not be safe, so modify the client as needed and then
	// clone. The headers are cloned based on the CloneHeaders property of the
	// source config
	//
	// Also, only the client's config is currently copied; this means items not in
	// the api.Config struct, such as policy override and wrapping function
	// behavior, must currently then be set as desired on the new client.
	Clone() (*api.Client, error)
	// CloneWithHeaders creates a new client similar to Clone, with the difference
	// being that the  headers are always cloned
	CloneWithHeaders() (*api.Client, error)
	// SetPolicyOverride sets whether requests should be sent with the policy
	// override flag to request overriding soft-mandatory Sentinel policies (both
	// RGPs and EGPs)
	SetPolicyOverride(override bool)
	// NewRequest creates a new raw request object to query the Vault server
	// configured for this client. This is an advanced method and generally
	// doesn't need to be called externally.
	NewRequest(method string, requestPath string) *api.Request
	// RawRequest performs the raw request given. This request may be against
	// a Vault server not configured with this client. This is an advanced operation
	// that generally won't need to be called externally.
	//
	// Deprecated: This method should not be used directly. Use higher level
	// methods instead.
	RawRequest(r *api.Request) (*api.Response, error)
	// RawRequestWithContext performs the raw request given. This request may be against
	// a Vault server not configured with this client. This is an advanced operation
	// that generally won't need to be called externally.
	//
	// Deprecated: This method should not be used directly. Use higher level
	// methods instead.
	RawRequestWithContext(ctx context.Context, r *api.Request) (*api.Response, error)
	// WithRequestCallbacks makes a shallow clone of Client, modifies it to use
	// the given callbacks, and returns it.  Each of the callbacks will be invoked
	// on every outgoing request.  A client may be used to issue requests
	// concurrently; any locking needed by callbacks invoked concurrently is the
	// callback's responsibility.
	WithRequestCallbacks(callbacks ...api.RequestCallback) *api.Client
	// WithResponseCallbacks makes a shallow clone of Client, modifies it to use
	// the given callbacks, and returns it.  Each of the callbacks will be invoked
	// on every received response.  A client may be used to issue requests
	// concurrently; any locking needed by callbacks invoked concurrently is the
	// callback's responsibility.
	WithResponseCallbacks(callbacks ...api.ResponseCallback) *api.Client
	// Help wraps HelpWithContext using context.Background.
	Help(path string) (*api.Help, error)
	// HelpWithContext reads the help information for the given path.
	HelpWithContext(ctx context.Context, path string) (*api.Help, error)
	// KVv1 is used to return a client for reads and writes against
	// a KV v1 secrets engine in Vault.
	//
	// The mount path is the location where the target KV secrets engine resides
	// in Vault.
	//
	// While v1 is not necessarily deprecated, Vault development servers tend to
	// use v2 as the version of the KV secrets engine, as this is what's mounted
	// by default when a server is started in -dev mode. See the kvv2 struct.
	//
	// Learn more about the KV secrets engine here:
	// https://www.vaultproject.io/docs/secrets/kv
	KVv1(mountPath string) *api.KVv1
	// KVv2 is used to return a client for reads and writes against
	// a KV v2 secrets engine in Vault.
	//
	// The mount path is the location where the target KV secrets engine resides
	// in Vault.
	//
	// Vault development servers tend to have "secret" as the mount path,
	// as these are the default settings when a server is started in -dev mode.
	//
	// Learn more about the KV secrets engine here:
	// https://www.vaultproject.io/docs/secrets/kv
	KVv2(mountPath string) *api.KVv2
	// NewLifetimeWatcher creates a new renewer from the given input.
	NewLifetimeWatcher(i *api.LifetimeWatcherInput) (*api.LifetimeWatcher, error)
	// Deprecated: exists only for backwards compatibility. Calls
	// NewLifetimeWatcher, and sets compatibility flags.
	NewRenewer(i *api.LifetimeWatcherInput) (*api.LifetimeWatcher, error)
	// Logical is used to return the client for logical-backend API calls.
	Logical() *api.Logical
	// SSH returns the client for logical-backend API calls.
	SSH() *api.SSH
	// SSHWithMountPoint returns the client with specific SSH mount point.
	SSHWithMountPoint(mountPoint string) *api.SSH
	// SSHHelper creates an SSHHelper object which can talk to Vault server with SSH backend
	// mounted at default path ("ssh").
	SSHHelper() *api.SSHHelper
	// SSHHelperWithMountPoint creates an SSHHelper object which can talk to Vault server with SSH backend
	// mounted at a specific mount point.
	SSHHelperWithMountPoint(mountPoint string) *api.SSHHelper
	// Sys is used to return the client for sys-related API calls.
	Sys() *api.Sys
}
`,
			files:          vaultAPIFiles,
			structName:     "Client",
			interfaceName:  "Client",
			outPackageName: "vault",
		},
	}

	modpath := filepath.Join(gopath.Find(), "pkg", "mod")

	for _, tc := range cases {
		tc := tc

		t.Run(tc.name, func(t *testing.T) {
			cmd := exec.Command("go", "get", tc.module)
			cmd.Env = append(os.Environ(), "GOPATH="+gopath.Find())
			out, err := cmd.CombinedOutput()
			require.NoErrorf(t, err, "cmd output: %s", string(out))

			got, err := Generate(Options{
				Files:             encodeFiles(tc.files, modpath),
				StructName:        tc.structName,
				InterfaceName:     tc.interfaceName,
				OutputPackageName: tc.outPackageName,
			})

			require.NoError(t, err)
			require.Equal(t, tc.want, string(got))
		})
	}
}

func encodeFiles(files []string, modpath string) []string {
	result := make([]string, len(files))
	for i, f := range files {
		result[i] = filepath.Join(modpath, f)
	}
	return result
}
