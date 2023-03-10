package generator

import (
	"os"
	"os/exec"
	"path/filepath"
	"testing"

	"github.com/denisdubovitskiy/ifacemaker/internal/golang"
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
			module:         "github.com/mattermost/mattermost-server/v5@v5.39.3",
			name:           "mattermost model audit",
			want:           "01_audit",
			files:          mattermostModelFiles,
			structName:     "Audit",
			interfaceName:  "Audit",
			outPackageName: "audit",
		},
		{
			module:         "github.com/mattermost/mattermost-server/v5@v5.39.3",
			name:           "mattermost model audit interface rename",
			want:           "02_audit_rename",
			files:          mattermostModelFiles,
			structName:     "Audit",
			interfaceName:  "Audit2",
			outPackageName: "audit",
		},
		{
			module:         "github.com/mattermost/mattermost-server/v5@v5.39.3",
			name:           "mattermost model audit package rename",
			want:           "03_audit_package_rename",
			files:          mattermostModelFiles,
			structName:     "Audit",
			interfaceName:  "Audit",
			outPackageName: "testpackage",
		},
		{
			module:         "github.com/mattermost/mattermost-server/v5@v5.39.3",
			name:           "mattermost model client4",
			want:           "04_client4",
			files:          mattermostModelFiles,
			structName:     "Client4",
			interfaceName:  "Client4",
			outPackageName: "client",
		},
		{
			module:         "github.com/mattermost/mattermost-server/v5@v5.39.3",
			name:           "mattermost model user",
			want:           "05_user",
			files:          mattermostModelFiles,
			structName:     "User",
			interfaceName:  "User",
			outPackageName: "user",
		},
		{
			module:         "github.com/hashicorp/vault/api@v1.8.2",
			name:           "vault api client",
			want:           "05_vault_api_client",
			files:          vaultAPIFiles,
			structName:     "Client",
			interfaceName:  "Client",
			outPackageName: "vault",
		},
	}

	modpath := filepath.Join(golang.GOPATH(), "pkg", "mod")

	for _, tc := range cases {
		tc := tc

		t.Run(tc.name, func(t *testing.T) {
			cmd := exec.Command("go", "get", tc.module)
			cmd.Env = append(os.Environ(), "GOPATH="+golang.GOPATH())
			out, err := cmd.CombinedOutput()
			require.NoErrorf(t, err, "cmd output: %s", string(out))

			content, err := os.ReadFile(filepath.Join("testdata", tc.want, "out.txt"))
			require.NoError(t, err)

			want := string(content)

			got, err := Generate(Options{
				Files:             encodeFiles(tc.files, modpath),
				StructName:        tc.structName,
				InterfaceName:     tc.interfaceName,
				OutputPackageName: tc.outPackageName,
			})

			require.NoError(t, err)
			require.Equal(t, want, string(got))
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
