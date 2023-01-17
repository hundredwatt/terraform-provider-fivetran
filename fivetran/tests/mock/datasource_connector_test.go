package mock

import (
	"net/http"
	"testing"

	"github.com/fivetran/go-fivetran/tests/mock"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

var (
	connectorDataSourceMockGetHandler *mock.Handler

	connectorDataSourceMockData map[string]interface{}
)

func setupMockClientConnectorDataSourceConfigMapping(t *testing.T) {
	mockClient.Reset()

	connectorDataSourceMockGetHandler = mockClient.When(http.MethodGet, "/v1/connectors/connector_id").ThenCall(
		func(req *http.Request) (*http.Response, error) {
			connectorDataSourceMockData = createMapFromJsonString(t, connectorMappingResponse)
			return fivetranSuccessResponse(t, req, http.StatusOK, "Success", connectorDataSourceMockData), nil
		},
	)
}

func TestDataSourceConnectorConfigMappingMock(t *testing.T) {
	// NOTE: the config is totally inconsistent and contains all possible values for mapping test
	step1 := resource.TestStep{
		Config: `
		data "fivetran_connector" "test_connector" {
			provider = fivetran-provider
			id = "connector_id"
		}`,

		Check: resource.ComposeAggregateTestCheckFunc(
			func(s *terraform.State) error {
				assertEqual(t, connectorDataSourceMockGetHandler.Interactions, 2)
				assertNotEmpty(t, connectorDataSourceMockData)
				return nil
			},
			resource.TestCheckResourceAttr("data.fivetran_connector.test_connector", "service", "google_sheets"),
			resource.TestCheckResourceAttr("data.fivetran_connector.test_connector", "service", "google_sheets"),
			resource.TestCheckResourceAttr("data.fivetran_connector.test_connector", "service_version", "1"),
			resource.TestCheckResourceAttr("data.fivetran_connector.test_connector", "schedule_type", "auto"),

			resource.TestCheckResourceAttr("data.fivetran_connector.test_connector", "status.0.is_historical_sync", "true"),
			resource.TestCheckResourceAttr("data.fivetran_connector.test_connector", "status.0.update_state", "on_schedule"),
			resource.TestCheckResourceAttr("data.fivetran_connector.test_connector", "status.0.setup_state", "incomplete"),
			resource.TestCheckResourceAttr("data.fivetran_connector.test_connector", "status.0.sync_state", "paused"),

			resource.TestCheckResourceAttr("data.fivetran_connector.test_connector", "status.0.tasks.0.code", "task_code"),
			resource.TestCheckResourceAttr("data.fivetran_connector.test_connector", "status.0.tasks.0.message", "task_message"),

			resource.TestCheckResourceAttr("data.fivetran_connector.test_connector", "status.0.warnings.0.code", "warning_code"),
			resource.TestCheckResourceAttr("data.fivetran_connector.test_connector", "status.0.warnings.0.message", "warning_message"),

			resource.TestCheckResourceAttr("data.fivetran_connector.test_connector", "name", "google_sheets_schema.table"),
			resource.TestCheckResourceAttr("data.fivetran_connector.test_connector", "sync_frequency", "5"),
			resource.TestCheckResourceAttr("data.fivetran_connector.test_connector", "paused", "true"),
			resource.TestCheckResourceAttr("data.fivetran_connector.test_connector", "pause_after_trial", "true"),

			// check sensitive fields are have original values
			resource.TestCheckResourceAttr("data.fivetran_connector.test_connector", "config.0.oauth_token", "******"),
			resource.TestCheckResourceAttr("data.fivetran_connector.test_connector", "config.0.oauth_token_secret", "******"),
			resource.TestCheckResourceAttr("data.fivetran_connector.test_connector", "config.0.consumer_key", "******"),
			resource.TestCheckResourceAttr("data.fivetran_connector.test_connector", "config.0.client_secret", "******"),
			resource.TestCheckResourceAttr("data.fivetran_connector.test_connector", "config.0.private_key", "******"),
			resource.TestCheckResourceAttr("data.fivetran_connector.test_connector", "config.0.s3role_arn", "******"),
			resource.TestCheckResourceAttr("data.fivetran_connector.test_connector", "config.0.ftp_password", "******"),
			resource.TestCheckResourceAttr("data.fivetran_connector.test_connector", "config.0.sftp_password", "******"),
			resource.TestCheckResourceAttr("data.fivetran_connector.test_connector", "config.0.api_key", "******"),
			resource.TestCheckResourceAttr("data.fivetran_connector.test_connector", "config.0.role_arn", "******"),
			resource.TestCheckResourceAttr("data.fivetran_connector.test_connector", "config.0.password", "******"),
			resource.TestCheckResourceAttr("data.fivetran_connector.test_connector", "config.0.secret_key", "******"),
			resource.TestCheckResourceAttr("data.fivetran_connector.test_connector", "config.0.pem_certificate", "******"),
			resource.TestCheckResourceAttr("data.fivetran_connector.test_connector", "config.0.access_token", "******"),
			resource.TestCheckResourceAttr("data.fivetran_connector.test_connector", "config.0.api_secret", "******"),
			resource.TestCheckResourceAttr("data.fivetran_connector.test_connector", "config.0.api_access_token", "******"),
			resource.TestCheckResourceAttr("data.fivetran_connector.test_connector", "config.0.secret", "******"),
			resource.TestCheckResourceAttr("data.fivetran_connector.test_connector", "config.0.consumer_secret", "******"),
			resource.TestCheckResourceAttr("data.fivetran_connector.test_connector", "config.0.secrets", "******"),
			resource.TestCheckResourceAttr("data.fivetran_connector.test_connector", "config.0.api_token", "******"),
			resource.TestCheckResourceAttr("data.fivetran_connector.test_connector", "config.0.encryption_key", "******"),
			resource.TestCheckResourceAttr("data.fivetran_connector.test_connector", "config.0.pat", "******"),
			resource.TestCheckResourceAttr("data.fivetran_connector.test_connector", "config.0.function_trigger", "******"),
			resource.TestCheckResourceAttr("data.fivetran_connector.test_connector", "config.0.token_key", "******"),
			resource.TestCheckResourceAttr("data.fivetran_connector.test_connector", "config.0.token_secret", "******"),

			resource.TestCheckResourceAttr("data.fivetran_connector.test_connector", "config.0.sheet_id", "sheet_id"),
			resource.TestCheckResourceAttr("data.fivetran_connector.test_connector", "config.0.auth_type", "OAuth"),
			resource.TestCheckResourceAttr("data.fivetran_connector.test_connector", "config.0.named_range", "range"),
			resource.TestCheckResourceAttr("data.fivetran_connector.test_connector", "config.0.sync_method", "sync_method"),
			resource.TestCheckResourceAttr("data.fivetran_connector.test_connector", "config.0.is_ftps", "false"),

			resource.TestCheckResourceAttr("data.fivetran_connector.test_connector", "config.0.sftp_is_key_pair", "false"),
			resource.TestCheckResourceAttr("data.fivetran_connector.test_connector", "config.0.sync_data_locker", "false"),
			resource.TestCheckResourceAttr("data.fivetran_connector.test_connector", "config.0.enable_all_dimension_combinations", "false"),
			resource.TestCheckResourceAttr("data.fivetran_connector.test_connector", "config.0.update_config_on_each_sync", "false"),
			resource.TestCheckResourceAttr("data.fivetran_connector.test_connector", "config.0.on_premise", "false"),
			resource.TestCheckResourceAttr("data.fivetran_connector.test_connector", "config.0.is_new_package", "false"),
			resource.TestCheckResourceAttr("data.fivetran_connector.test_connector", "config.0.is_multi_entity_feature_enabled", "false"),
			resource.TestCheckResourceAttr("data.fivetran_connector.test_connector", "config.0.always_encrypted", "false"),
			resource.TestCheckResourceAttr("data.fivetran_connector.test_connector", "config.0.is_secure", "false"),
			resource.TestCheckResourceAttr("data.fivetran_connector.test_connector", "config.0.use_api_keys", "false"),
			resource.TestCheckResourceAttr("data.fivetran_connector.test_connector", "config.0.use_webhooks", "false"),
			resource.TestCheckResourceAttr("data.fivetran_connector.test_connector", "config.0.eu_region", "false"),
			resource.TestCheckResourceAttr("data.fivetran_connector.test_connector", "config.0.is_keypair", "false"),
			resource.TestCheckResourceAttr("data.fivetran_connector.test_connector", "config.0.is_account_level_connector", "true"),

			resource.TestCheckResourceAttr("data.fivetran_connector.test_connector", "config.0.conversion_window_size", "0"),
			resource.TestCheckResourceAttr("data.fivetran_connector.test_connector", "config.0.skip_before", "0"),
			resource.TestCheckResourceAttr("data.fivetran_connector.test_connector", "config.0.skip_after", "0"),
			resource.TestCheckResourceAttr("data.fivetran_connector.test_connector", "config.0.ftp_port", "0"),
			resource.TestCheckResourceAttr("data.fivetran_connector.test_connector", "config.0.sftp_port", "0"),
			resource.TestCheckResourceAttr("data.fivetran_connector.test_connector", "config.0.port", "0"),
			resource.TestCheckResourceAttr("data.fivetran_connector.test_connector", "config.0.tunnel_port", "0"),
			resource.TestCheckResourceAttr("data.fivetran_connector.test_connector", "config.0.api_quota", "0"),
			resource.TestCheckResourceAttr("data.fivetran_connector.test_connector", "config.0.daily_api_call_limit", "0"),

			resource.TestCheckResourceAttr("data.fivetran_connector.test_connector", "config.0.connection_type", "connection_type"),
			resource.TestCheckResourceAttr("data.fivetran_connector.test_connector", "config.0.sync_method", "sync_method"),
			resource.TestCheckResourceAttr("data.fivetran_connector.test_connector", "config.0.sync_mode", "sync_mode"),
			resource.TestCheckResourceAttr("data.fivetran_connector.test_connector", "config.0.date_granularity", "date_granularity"),
			resource.TestCheckResourceAttr("data.fivetran_connector.test_connector", "config.0.timeframe_months", "timeframe_months"),
			resource.TestCheckResourceAttr("data.fivetran_connector.test_connector", "config.0.report_type", "report_type"),
			resource.TestCheckResourceAttr("data.fivetran_connector.test_connector", "config.0.aggregation", "aggregation"),
			resource.TestCheckResourceAttr("data.fivetran_connector.test_connector", "config.0.config_type", "config_type"),
			resource.TestCheckResourceAttr("data.fivetran_connector.test_connector", "config.0.prebuilt_report", "prebuilt_report"),
			resource.TestCheckResourceAttr("data.fivetran_connector.test_connector", "config.0.action_report_time", "action_report_time"),
			resource.TestCheckResourceAttr("data.fivetran_connector.test_connector", "config.0.click_attribution_window", "click_attribution_window"),
			resource.TestCheckResourceAttr("data.fivetran_connector.test_connector", "config.0.view_attribution_window", "view_attribution_window"),
			resource.TestCheckResourceAttr("data.fivetran_connector.test_connector", "config.0.view_through_attribution_window_size", "view_through_attribution_window_size"),
			resource.TestCheckResourceAttr("data.fivetran_connector.test_connector", "config.0.post_click_attribution_window_size", "post_click_attribution_window_size"),
			resource.TestCheckResourceAttr("data.fivetran_connector.test_connector", "config.0.update_method", "update_method"),
			resource.TestCheckResourceAttr("data.fivetran_connector.test_connector", "config.0.swipe_attribution_window", "swipe_attribution_window"),
			resource.TestCheckResourceAttr("data.fivetran_connector.test_connector", "config.0.api_type", "api_type"),
			resource.TestCheckResourceAttr("data.fivetran_connector.test_connector", "config.0.sync_format", "sync_format"),
			resource.TestCheckResourceAttr("data.fivetran_connector.test_connector", "config.0.app_sync_mode", "app_sync_mode"),
			resource.TestCheckResourceAttr("data.fivetran_connector.test_connector", "config.0.sales_account_sync_mode", "sales_account_sync_mode"),
			resource.TestCheckResourceAttr("data.fivetran_connector.test_connector", "config.0.finance_account_sync_mode", "finance_account_sync_mode"),
			resource.TestCheckResourceAttr("data.fivetran_connector.test_connector", "config.0.source", "source"),
			resource.TestCheckResourceAttr("data.fivetran_connector.test_connector", "config.0.file_type", "file_type"),
			resource.TestCheckResourceAttr("data.fivetran_connector.test_connector", "config.0.compression", "compression"),
			resource.TestCheckResourceAttr("data.fivetran_connector.test_connector", "config.0.on_error", "on_error"),
			resource.TestCheckResourceAttr("data.fivetran_connector.test_connector", "config.0.append_file_option", "append_file_option"),
			resource.TestCheckResourceAttr("data.fivetran_connector.test_connector", "config.0.engagement_attribution_window", "engagement_attribution_window"),
			resource.TestCheckResourceAttr("data.fivetran_connector.test_connector", "config.0.conversion_report_time", "conversion_report_time"),
			resource.TestCheckResourceAttr("data.fivetran_connector.test_connector", "config.0.external_id", "external_id"),
			resource.TestCheckResourceAttr("data.fivetran_connector.test_connector", "config.0.public_key", "public_key"),
			resource.TestCheckResourceAttr("data.fivetran_connector.test_connector", "config.0.client_id", "client_id"),
			resource.TestCheckResourceAttr("data.fivetran_connector.test_connector", "config.0.technical_account_id", "technical_account_id"),
			resource.TestCheckResourceAttr("data.fivetran_connector.test_connector", "config.0.organization_id", "organization_id"),
			resource.TestCheckResourceAttr("data.fivetran_connector.test_connector", "config.0.s3bucket", "s3bucket"),
			resource.TestCheckResourceAttr("data.fivetran_connector.test_connector", "config.0.abs_connection_string", "abs_connection_string"),
			resource.TestCheckResourceAttr("data.fivetran_connector.test_connector", "config.0.abs_container_name", "abs_container_name"),
			resource.TestCheckResourceAttr("data.fivetran_connector.test_connector", "config.0.folder_id", "folder_id"),
			resource.TestCheckResourceAttr("data.fivetran_connector.test_connector", "config.0.ftp_host", "ftp_host"),
			resource.TestCheckResourceAttr("data.fivetran_connector.test_connector", "config.0.ftp_user", "ftp_user"),
			resource.TestCheckResourceAttr("data.fivetran_connector.test_connector", "config.0.sftp_host", "sftp_host"),
			resource.TestCheckResourceAttr("data.fivetran_connector.test_connector", "config.0.sftp_user", "sftp_user"),
			resource.TestCheckResourceAttr("data.fivetran_connector.test_connector", "config.0.bucket", "bucket"),
			resource.TestCheckResourceAttr("data.fivetran_connector.test_connector", "config.0.prefix", "prefix"),
			resource.TestCheckResourceAttr("data.fivetran_connector.test_connector", "config.0.pattern", "pattern"),
			resource.TestCheckResourceAttr("data.fivetran_connector.test_connector", "config.0.archive_pattern", "archive_pattern"),
			resource.TestCheckResourceAttr("data.fivetran_connector.test_connector", "config.0.null_sequence", "null_sequence"),
			resource.TestCheckResourceAttr("data.fivetran_connector.test_connector", "config.0.delimiter", "delimiter"),
			resource.TestCheckResourceAttr("data.fivetran_connector.test_connector", "config.0.escape_char", "escape_char"),
			resource.TestCheckResourceAttr("data.fivetran_connector.test_connector", "config.0.auth_mode", "auth_mode"),
			resource.TestCheckResourceAttr("data.fivetran_connector.test_connector", "config.0.certificate", "certificate"),
			resource.TestCheckResourceAttr("data.fivetran_connector.test_connector", "config.0.consumer_group", "consumer_group"),
			resource.TestCheckResourceAttr("data.fivetran_connector.test_connector", "config.0.servers", "servers"),
			resource.TestCheckResourceAttr("data.fivetran_connector.test_connector", "config.0.message_type", "message_type"),
			resource.TestCheckResourceAttr("data.fivetran_connector.test_connector", "config.0.sync_type", "sync_type"),
			resource.TestCheckResourceAttr("data.fivetran_connector.test_connector", "config.0.security_protocol", "security_protocol"),
			resource.TestCheckResourceAttr("data.fivetran_connector.test_connector", "config.0.access_key_id", "access_key_id"),
			resource.TestCheckResourceAttr("data.fivetran_connector.test_connector", "config.0.home_folder", "home_folder"),
			resource.TestCheckResourceAttr("data.fivetran_connector.test_connector", "config.0.function", "function"),
			resource.TestCheckResourceAttr("data.fivetran_connector.test_connector", "config.0.region", "region"),
			resource.TestCheckResourceAttr("data.fivetran_connector.test_connector", "config.0.container_name", "container_name"),
			resource.TestCheckResourceAttr("data.fivetran_connector.test_connector", "config.0.connection_string", "connection_string"),
			resource.TestCheckResourceAttr("data.fivetran_connector.test_connector", "config.0.function_app", "function_app"),
			resource.TestCheckResourceAttr("data.fivetran_connector.test_connector", "config.0.function_name", "function_name"),
			resource.TestCheckResourceAttr("data.fivetran_connector.test_connector", "config.0.function_key", "function_key"),
			resource.TestCheckResourceAttr("data.fivetran_connector.test_connector", "config.0.merchant_id", "merchant_id"),
			resource.TestCheckResourceAttr("data.fivetran_connector.test_connector", "config.0.api_url", "api_url"),
			resource.TestCheckResourceAttr("data.fivetran_connector.test_connector", "config.0.cloud_storage_type", "cloud_storage_type"),
			resource.TestCheckResourceAttr("data.fivetran_connector.test_connector", "config.0.s3external_id", "s3external_id"),
			resource.TestCheckResourceAttr("data.fivetran_connector.test_connector", "config.0.s3folder", "s3folder"),
			resource.TestCheckResourceAttr("data.fivetran_connector.test_connector", "config.0.gcs_bucket", "gcs_bucket"),
			resource.TestCheckResourceAttr("data.fivetran_connector.test_connector", "config.0.gcs_folder", "gcs_folder"),
			resource.TestCheckResourceAttr("data.fivetran_connector.test_connector", "config.0.instance", "instance"),
			resource.TestCheckResourceAttr("data.fivetran_connector.test_connector", "config.0.aws_region_code", "aws_region_code"),
			resource.TestCheckResourceAttr("data.fivetran_connector.test_connector", "config.0.host", "host"),
			resource.TestCheckResourceAttr("data.fivetran_connector.test_connector", "config.0.user", "user"),
			resource.TestCheckResourceAttr("data.fivetran_connector.test_connector", "config.0.network_code", "network_code"),
			resource.TestCheckResourceAttr("data.fivetran_connector.test_connector", "config.0.customer_id", "customer_id"),
			resource.TestCheckResourceAttr("data.fivetran_connector.test_connector", "config.0.project_id", "project_id"),
			resource.TestCheckResourceAttr("data.fivetran_connector.test_connector", "config.0.dataset_id", "dataset_id"),
			resource.TestCheckResourceAttr("data.fivetran_connector.test_connector", "config.0.bucket_name", "bucket_name"),
			resource.TestCheckResourceAttr("data.fivetran_connector.test_connector", "config.0.config_method", "config_method"),
			resource.TestCheckResourceAttr("data.fivetran_connector.test_connector", "config.0.query_id", "query_id"),
			resource.TestCheckResourceAttr("data.fivetran_connector.test_connector", "config.0.path", "path"),
			resource.TestCheckResourceAttr("data.fivetran_connector.test_connector", "config.0.endpoint", "endpoint"),
			resource.TestCheckResourceAttr("data.fivetran_connector.test_connector", "config.0.identity", "identity"),
			resource.TestCheckResourceAttr("data.fivetran_connector.test_connector", "config.0.domain_name", "domain_name"),
			resource.TestCheckResourceAttr("data.fivetran_connector.test_connector", "config.0.resource_url", "resource_url"),
			resource.TestCheckResourceAttr("data.fivetran_connector.test_connector", "config.0.tunnel_host", "tunnel_host"),
			resource.TestCheckResourceAttr("data.fivetran_connector.test_connector", "config.0.tunnel_user", "tunnel_user"),
			resource.TestCheckResourceAttr("data.fivetran_connector.test_connector", "config.0.database", "database"),
			resource.TestCheckResourceAttr("data.fivetran_connector.test_connector", "config.0.datasource", "datasource"),
			resource.TestCheckResourceAttr("data.fivetran_connector.test_connector", "config.0.account", "account"),
			resource.TestCheckResourceAttr("data.fivetran_connector.test_connector", "config.0.role", "role"),
			resource.TestCheckResourceAttr("data.fivetran_connector.test_connector", "config.0.email", "email"),
			resource.TestCheckResourceAttr("data.fivetran_connector.test_connector", "config.0.account_id", "account_id"),
			resource.TestCheckResourceAttr("data.fivetran_connector.test_connector", "config.0.server_url", "server_url"),
			resource.TestCheckResourceAttr("data.fivetran_connector.test_connector", "config.0.user_key", "user_key"),
			resource.TestCheckResourceAttr("data.fivetran_connector.test_connector", "config.0.api_version", "api_version"),
			resource.TestCheckResourceAttr("data.fivetran_connector.test_connector", "config.0.time_zone", "time_zone"),
			resource.TestCheckResourceAttr("data.fivetran_connector.test_connector", "config.0.integration_key", "integration_key"),
			resource.TestCheckResourceAttr("data.fivetran_connector.test_connector", "config.0.domain", "domain"),
			resource.TestCheckResourceAttr("data.fivetran_connector.test_connector", "config.0.replication_slot", "replication_slot"),
			resource.TestCheckResourceAttr("data.fivetran_connector.test_connector", "config.0.publication_name", "publication_name"),
			resource.TestCheckResourceAttr("data.fivetran_connector.test_connector", "config.0.data_center", "data_center"),
			resource.TestCheckResourceAttr("data.fivetran_connector.test_connector", "config.0.sub_domain", "sub_domain"),
			resource.TestCheckResourceAttr("data.fivetran_connector.test_connector", "config.0.subdomain", "subdomain"),
			resource.TestCheckResourceAttr("data.fivetran_connector.test_connector", "config.0.test_table_name", "test_table_name"),
			resource.TestCheckResourceAttr("data.fivetran_connector.test_connector", "config.0.shop", "shop"),
			resource.TestCheckResourceAttr("data.fivetran_connector.test_connector", "config.0.sid", "sid"),
			resource.TestCheckResourceAttr("data.fivetran_connector.test_connector", "config.0.key", "key"),
			resource.TestCheckResourceAttr("data.fivetran_connector.test_connector", "config.0.bucket_service", "bucket_service"),
			resource.TestCheckResourceAttr("data.fivetran_connector.test_connector", "config.0.user_name", "user_name"),
			resource.TestCheckResourceAttr("data.fivetran_connector.test_connector", "config.0.username", "username"),
			resource.TestCheckResourceAttr("data.fivetran_connector.test_connector", "config.0.report_url", "report_url"),
			resource.TestCheckResourceAttr("data.fivetran_connector.test_connector", "config.0.unique_id", "unique_id"),
			resource.TestCheckResourceAttr("data.fivetran_connector.test_connector", "config.0.base_url", "base_url"),
			resource.TestCheckResourceAttr("data.fivetran_connector.test_connector", "config.0.entity_id", "entity_id"),
			resource.TestCheckResourceAttr("data.fivetran_connector.test_connector", "config.0.soap_uri", "soap_uri"),
			resource.TestCheckResourceAttr("data.fivetran_connector.test_connector", "config.0.user_id", "user_id"),
			resource.TestCheckResourceAttr("data.fivetran_connector.test_connector", "config.0.share_url", "share_url"),
			resource.TestCheckResourceAttr("data.fivetran_connector.test_connector", "config.0.report_suites.0", "report_suite"),
			resource.TestCheckResourceAttr("data.fivetran_connector.test_connector", "config.0.elements.0", "element"),
			resource.TestCheckResourceAttr("data.fivetran_connector.test_connector", "config.0.metrics.0", "metric"),
			resource.TestCheckResourceAttr("data.fivetran_connector.test_connector", "config.0.advertisables.0", "advertisable"),
			resource.TestCheckResourceAttr("data.fivetran_connector.test_connector", "config.0.dimensions.0", "dimension"),
			resource.TestCheckResourceAttr("data.fivetran_connector.test_connector", "config.0.selected_exports.0", "selected_export"),
			resource.TestCheckResourceAttr("data.fivetran_connector.test_connector", "config.0.apps.0", "app"),
			resource.TestCheckResourceAttr("data.fivetran_connector.test_connector", "config.0.sales_accounts.0", "sales_account"),
			resource.TestCheckResourceAttr("data.fivetran_connector.test_connector", "config.0.finance_accounts.0", "finance_account"),
			resource.TestCheckResourceAttr("data.fivetran_connector.test_connector", "config.0.projects.0", "project"),
			resource.TestCheckResourceAttr("data.fivetran_connector.test_connector", "config.0.user_profiles.0", "user_profile"),
			resource.TestCheckResourceAttr("data.fivetran_connector.test_connector", "config.0.report_configuration_ids.0", "report_configuration_id"),
			resource.TestCheckResourceAttr("data.fivetran_connector.test_connector", "config.0.accounts.0", "account"),
			resource.TestCheckResourceAttr("data.fivetran_connector.test_connector", "config.0.fields.0", "field"),
			resource.TestCheckResourceAttr("data.fivetran_connector.test_connector", "config.0.breakdowns.0", "breakdown"),
			resource.TestCheckResourceAttr("data.fivetran_connector.test_connector", "config.0.action_breakdowns.0", "action_breakdown"),
			resource.TestCheckResourceAttr("data.fivetran_connector.test_connector", "config.0.pages.0", "page"),
			resource.TestCheckResourceAttr("data.fivetran_connector.test_connector", "config.0.repositories.0", "repository"),
			resource.TestCheckResourceAttr("data.fivetran_connector.test_connector", "config.0.dimension_attributes.0", "dimension_attribute"),
			resource.TestCheckResourceAttr("data.fivetran_connector.test_connector", "config.0.columns.0", "column"),
			resource.TestCheckResourceAttr("data.fivetran_connector.test_connector", "config.0.manager_accounts.0", "manager_account"),
			resource.TestCheckResourceAttr("data.fivetran_connector.test_connector", "config.0.profiles.0", "profile"),
			resource.TestCheckResourceAttr("data.fivetran_connector.test_connector", "config.0.site_urls.0", "site_url"),
			resource.TestCheckResourceAttr("data.fivetran_connector.test_connector", "config.0.api_keys.0", "******"),
			resource.TestCheckResourceAttr("data.fivetran_connector.test_connector", "config.0.advertisers_id.0", "advertiser_id"),
			resource.TestCheckResourceAttr("data.fivetran_connector.test_connector", "config.0.hosts.0", "host"),
			resource.TestCheckResourceAttr("data.fivetran_connector.test_connector", "config.0.advertisers.0", "advertiser"),
			resource.TestCheckResourceAttr("data.fivetran_connector.test_connector", "config.0.organizations.0", "organization"),
			resource.TestCheckResourceAttr("data.fivetran_connector.test_connector", "config.0.account_ids.0", "account_id"),

			resource.TestCheckResourceAttr("data.fivetran_connector.test_connector", "config.0.adobe_analytics_configurations.0.sync_mode", "sync_mode"),
			resource.TestCheckResourceAttr("data.fivetran_connector.test_connector", "config.0.adobe_analytics_configurations.0.report_suites.0", "report_suite"),
			resource.TestCheckResourceAttr("data.fivetran_connector.test_connector", "config.0.adobe_analytics_configurations.0.elements.0", "element"),
			resource.TestCheckResourceAttr("data.fivetran_connector.test_connector", "config.0.adobe_analytics_configurations.0.metrics.0", "metric"),
			resource.TestCheckResourceAttr("data.fivetran_connector.test_connector", "config.0.adobe_analytics_configurations.0.calculated_metrics.0", "calculated_metric"),
			resource.TestCheckResourceAttr("data.fivetran_connector.test_connector", "config.0.adobe_analytics_configurations.0.segments.0", "segment"),

			resource.TestCheckResourceAttr("data.fivetran_connector.test_connector", "config.0.reports.0.table", "table"),
			resource.TestCheckResourceAttr("data.fivetran_connector.test_connector", "config.0.reports.0.config_type", "config_type"),
			resource.TestCheckResourceAttr("data.fivetran_connector.test_connector", "config.0.reports.0.prebuilt_report", "prebuilt_report"),
			resource.TestCheckResourceAttr("data.fivetran_connector.test_connector", "config.0.reports.0.report_type", "report_type"),
			resource.TestCheckResourceAttr("data.fivetran_connector.test_connector", "config.0.reports.0.fields.0", "field"),
			resource.TestCheckResourceAttr("data.fivetran_connector.test_connector", "config.0.reports.0.dimensions.0", "dimension"),
			resource.TestCheckResourceAttr("data.fivetran_connector.test_connector", "config.0.reports.0.metrics.0", "metric"),
			resource.TestCheckResourceAttr("data.fivetran_connector.test_connector", "config.0.reports.0.segments.0", "segment"),
			resource.TestCheckResourceAttr("data.fivetran_connector.test_connector", "config.0.reports.0.filter", "filter"),

			resource.TestCheckResourceAttr("data.fivetran_connector.test_connector", "config.0.custom_tables.0.table_name", "table_name"),
			resource.TestCheckResourceAttr("data.fivetran_connector.test_connector", "config.0.custom_tables.0.config_type", "config_type"),
			resource.TestCheckResourceAttr("data.fivetran_connector.test_connector", "config.0.custom_tables.0.fields.0", "field"),
			resource.TestCheckResourceAttr("data.fivetran_connector.test_connector", "config.0.custom_tables.0.breakdowns.0", "breakdown"),
			resource.TestCheckResourceAttr("data.fivetran_connector.test_connector", "config.0.custom_tables.0.action_breakdowns.0", "action_breakdown"),
			resource.TestCheckResourceAttr("data.fivetran_connector.test_connector", "config.0.custom_tables.0.aggregation", "aggregation"),
			resource.TestCheckResourceAttr("data.fivetran_connector.test_connector", "config.0.custom_tables.0.action_report_time", "action_report_time"),
			resource.TestCheckResourceAttr("data.fivetran_connector.test_connector", "config.0.custom_tables.0.click_attribution_window", "click_attribution_window"),
			resource.TestCheckResourceAttr("data.fivetran_connector.test_connector", "config.0.custom_tables.0.view_attribution_window", "view_attribution_window"),
			resource.TestCheckResourceAttr("data.fivetran_connector.test_connector", "config.0.custom_tables.0.prebuilt_report_name", "prebuilt_report_name"),

			resource.TestCheckResourceAttr("data.fivetran_connector.test_connector", "config.0.project_credentials.0.project", "project"),
			resource.TestCheckResourceAttr("data.fivetran_connector.test_connector", "config.0.project_credentials.0.api_key", "api_key"),
			resource.TestCheckResourceAttr("data.fivetran_connector.test_connector", "config.0.project_credentials.0.secret_key", "******"),

			resource.TestCheckResourceAttr("data.fivetran_connector.test_connector", "config.0.secrets_list.0.key", "key"),
			resource.TestCheckResourceAttr("data.fivetran_connector.test_connector", "config.0.secrets_list.0.value", "******"),
		),
	}

	resource.Test(
		t,
		resource.TestCase{
			PreCheck: func() {
				setupMockClientConnectorDataSourceConfigMapping(t)
			},
			Providers: testProviders,
			CheckDestroy: func(s *terraform.State) error {
				return nil
			},
			Steps: []resource.TestStep{
				step1,
			},
		},
	)
}
