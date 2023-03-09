package generator

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"testing"

	"github.com/denisdubovitskiy/ifacemaker/gopath"
	"github.com/stretchr/testify/require"
)

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

			g := string(got)
			fmt.Println(g)

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
