package main

type AdminUiConfig struct {
	CloudConfig `yaml:"cloud_config,inline"`
	BindAddress string `yaml:"bind_address"`
	DbURI       string `yaml:"db_uri"`
	Port        int    `yaml:"port"`
	StatsFile   string `cloud:"stats_file" yaml:"stats_file,omitempty"`
}
type UaaClientConfig struct {
	ID     string `cloud:"id" yaml:"id"`
	Secret string `cloud:"secret" yaml:"secret"`
}
type SenderEmailConfig struct {
	Server  string `cloud:"server,default=10.10.10.10" yaml:"server"`
	Account string `cloud:"account,default=system@10.10.10.10" yaml:"account"`
}
type CloudConfig struct {
	UaaClient                        UaaClientConfig   `cloud:"uaa_client" yaml:"uaa_client"`
	CcdbURI                          string            `cloud:"ccdb_uri" yaml:"ccdb_uri"`
	CloudControllerDiscoveryInterval int               `cloud:"cloud_controller_discovery_interval,default=300" yaml:"cloud_controller_discovery_interval"`
	CloudControllerSslVerifyNone     bool              `cloud:"cloud_controller_ssl_verify_none" yaml:"cloud_controller_ssl_verify_none"`
	CloudControllerURI               string            `cloud:"cloud_controller_uri" yaml:"cloud_controller_uri"`
	ComponentConnectionRetries       int               `cloud:"component_connection_retries,default=2" yaml:"component_connection_retries"`
	DataFile                         string            `cloud:"data_file,default=data/data.json" yaml:"data_file"`
	DisplayEncryptedValues           bool              `cloud:"display_encrypted_values,default=true" yaml:"display_encrypted_values"`
	DopplerDataFile                  string            `cloud:"doppler_data_file,default=data/data.json" yaml:"doppler_data_file"`
	DopplerReconnectDelay            int               `cloud:"doppler_reconnect_delay,default=300" yaml:"doppler_reconnect_delay"`
	DopplerRollupInterval            int               `cloud:"doppler_rollup_interval,default=30" yaml:"doppler_rollup_interval"`
	EventDays                        int               `cloud:"event_days,default=7" yaml:"event_days"`
	LogFile                          string            `cloud:"log_file,default=admin_ui.log" yaml:"log_file"`
	LogFilePageSize                  int               `cloud:"log_file_page_size,default=51200" yaml:"log_file_page_size"`
	LogFileSftpKeys                  []interface{}     `cloud:"log_file_sftp_keys" yaml:"log_file_sftp_keys"`
	LogFiles                         []string          `cloud:"log_files,default=admin_ui.log" yaml:"log_files"`
	Mbus                             string            `cloud:"mbus" yaml:"mbus"`
	MonitoredComponents              []string          `cloud:"monitored_components,default=NATS,CloudController,Diego,HealthManager,Router,-Provisioner,ALL" yaml:"monitored_components"`
	NatsDiscoveryInterval            int               `cloud:"nats_discovery_interval,default=30" yaml:"nats_discovery_interval"`
	NatsDiscoveryTimeout             int               `cloud:"nats_discovery_timeout,default=10" yaml:"nats_discovery_timeout"`
	ReceiverEmails                   []interface{}     `cloud:"receiver_emails" yaml:"receiver_emails"`
	SenderEmail                      SenderEmailConfig `cloud:"sender_email" yaml:"sender_email"`
	StatsRefreshSchedules            []string          `cloud:"stats_refresh_schedules,default=0 5 * * *" yaml:"stats_refresh_schedules"`
	StatsRetries                     int               `cloud:"stats_retries,default=5" yaml:"stats_retries"`
	StatsRetryInterval               int               `cloud:"stats_retry_interval,default=300" yaml:"stats_retry_interval"`
	UaadbURI                         string            `cloud:"uaadb_uri" yaml:"uaadb_uri"`
	UaaGroupsAdmin                   []string          `cloud:"uaa_groups_admin,default=admin_ui.admin" yaml:"uaa_groups_admin"`
	UaaGroupsUser                    []string          `cloud:"uaa_groups_user,default=admin_ui.user" yaml:"uaa_groups_user"`
	VarzDiscoveryInterval            int               `cloud:"varz_discovery_interval,default=30" yaml:"varz_discovery_interval"`
}
