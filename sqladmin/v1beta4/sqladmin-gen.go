// Package sqladmin provides access to the Cloud SQL Administration API.
//
// See https://developers.google.com/cloud-sql/docs/admin-api/
//
// Usage example:
//
//   import "google.golang.org/api/sqladmin/v1beta4"
//   ...
//   sqladminService, err := sqladmin.New(oauthHttpClient)
package sqladmin

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"golang.org/x/net/context"
	"google.golang.org/api/googleapi"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

// Always reference these packages, just in case the auto-generated code
// below doesn't.
var _ = bytes.NewBuffer
var _ = strconv.Itoa
var _ = fmt.Sprintf
var _ = json.NewDecoder
var _ = io.Copy
var _ = url.Parse
var _ = googleapi.Version
var _ = errors.New
var _ = strings.Replace
var _ = context.Background

const apiId = "sqladmin:v1beta4"
const apiName = "sqladmin"
const apiVersion = "v1beta4"
const basePath = "https://www.googleapis.com/sql/v1beta4/"

// OAuth2 scopes used by this API.
const (
	// View and manage your data across Google Cloud Platform services
	CloudPlatformScope = "https://www.googleapis.com/auth/cloud-platform"

	// Manage your Google SQL Service instances
	SqlserviceAdminScope = "https://www.googleapis.com/auth/sqlservice.admin"
)

func New(client *http.Client) (*Service, error) {
	if client == nil {
		return nil, errors.New("client is nil")
	}
	s := &Service{client: client, BasePath: basePath}
	s.BackupRuns = NewBackupRunsService(s)
	s.Databases = NewDatabasesService(s)
	s.Flags = NewFlagsService(s)
	s.Instances = NewInstancesService(s)
	s.Operations = NewOperationsService(s)
	s.SslCerts = NewSslCertsService(s)
	s.Tiers = NewTiersService(s)
	s.Users = NewUsersService(s)
	return s, nil
}

type Service struct {
	client    *http.Client
	BasePath  string // API endpoint base URL
	UserAgent string // optional additional User-Agent fragment

	BackupRuns *BackupRunsService

	Databases *DatabasesService

	Flags *FlagsService

	Instances *InstancesService

	Operations *OperationsService

	SslCerts *SslCertsService

	Tiers *TiersService

	Users *UsersService
}

func (s *Service) userAgent() string {
	if s.UserAgent == "" {
		return googleapi.UserAgent
	}
	return googleapi.UserAgent + " " + s.UserAgent
}

func NewBackupRunsService(s *Service) *BackupRunsService {
	rs := &BackupRunsService{s: s}
	return rs
}

type BackupRunsService struct {
	s *Service
}

func NewDatabasesService(s *Service) *DatabasesService {
	rs := &DatabasesService{s: s}
	return rs
}

type DatabasesService struct {
	s *Service
}

func NewFlagsService(s *Service) *FlagsService {
	rs := &FlagsService{s: s}
	return rs
}

type FlagsService struct {
	s *Service
}

func NewInstancesService(s *Service) *InstancesService {
	rs := &InstancesService{s: s}
	return rs
}

type InstancesService struct {
	s *Service
}

func NewOperationsService(s *Service) *OperationsService {
	rs := &OperationsService{s: s}
	return rs
}

type OperationsService struct {
	s *Service
}

func NewSslCertsService(s *Service) *SslCertsService {
	rs := &SslCertsService{s: s}
	return rs
}

type SslCertsService struct {
	s *Service
}

func NewTiersService(s *Service) *TiersService {
	rs := &TiersService{s: s}
	return rs
}

type TiersService struct {
	s *Service
}

func NewUsersService(s *Service) *UsersService {
	rs := &UsersService{s: s}
	return rs
}

type UsersService struct {
	s *Service
}

type AclEntry struct {
	// ExpirationTime: The time when this access control entry expires in
	// RFC 3339 format, for example 2012-11-15T16:19:00.094Z.
	ExpirationTime string `json:"expirationTime,omitempty"`

	// Kind: This is always sql#aclEntry.
	Kind string `json:"kind,omitempty"`

	// Name: An optional label to identify this entry.
	Name string `json:"name,omitempty"`

	// Value: The whitelisted value for the access control list.
	Value string `json:"value,omitempty"`
}

type BackupConfiguration struct {
	// BinaryLogEnabled: Whether binary log is enabled. If backup
	// configuration is disabled, binary log must be disabled as well.
	BinaryLogEnabled bool `json:"binaryLogEnabled,omitempty"`

	// Enabled: Whether this configuration is enabled.
	Enabled bool `json:"enabled,omitempty"`

	// Kind: This is always sql#backupConfiguration.
	Kind string `json:"kind,omitempty"`

	// StartTime: Start time for the daily backup configuration in UTC
	// timezone in the 24 hour format - HH:MM.
	StartTime string `json:"startTime,omitempty"`
}

type BackupRun struct {
	// EndTime: The time the backup operation completed in UTC timezone in
	// RFC 3339 format, for example 2012-11-15T16:19:00.094Z.
	EndTime string `json:"endTime,omitempty"`

	// EnqueuedTime: The time the run was enqueued in UTC timezone in RFC
	// 3339 format, for example 2012-11-15T16:19:00.094Z.
	EnqueuedTime string `json:"enqueuedTime,omitempty"`

	// Error: Information about why the backup operation failed. This is
	// only present if the run has the FAILED status.
	Error *OperationError `json:"error,omitempty"`

	// Id: A unique identifier for this backup run. Note that this is unique
	// only within the scope of a particular Cloud SQL instance.
	Id int64 `json:"id,omitempty,string"`

	// Instance: Name of the database instance.
	Instance string `json:"instance,omitempty"`

	// Kind: This is always sql#backupRun.
	Kind string `json:"kind,omitempty"`

	// SelfLink: The URI of this resource.
	SelfLink string `json:"selfLink,omitempty"`

	// StartTime: The time the backup operation actually started in UTC
	// timezone in RFC 3339 format, for example 2012-11-15T16:19:00.094Z.
	StartTime string `json:"startTime,omitempty"`

	// Status: The status of this run.
	Status string `json:"status,omitempty"`

	// WindowStartTime: The start time of the backup window during which
	// this the backup was attempted in RFC 3339 format, for example
	// 2012-11-15T16:19:00.094Z.
	WindowStartTime string `json:"windowStartTime,omitempty"`
}

type BackupRunsListResponse struct {
	// Items: A list of backup runs in reverse chronological order of the
	// enqueued time.
	Items []*BackupRun `json:"items,omitempty"`

	// Kind: This is always sql#backupRunsList.
	Kind string `json:"kind,omitempty"`

	// NextPageToken: The continuation token, used to page through large
	// result sets. Provide this value in a subsequent request to return the
	// next page of results.
	NextPageToken string `json:"nextPageToken,omitempty"`
}

type BinLogCoordinates struct {
	// BinLogFileName: Name of the binary log file for a Cloud SQL instance.
	BinLogFileName string `json:"binLogFileName,omitempty"`

	// BinLogPosition: Position (offset) within the binary log file.
	BinLogPosition int64 `json:"binLogPosition,omitempty,string"`

	// Kind: This is always sql#binLogCoordinates.
	Kind string `json:"kind,omitempty"`
}

type CloneContext struct {
	// BinLogCoordinates: Binary log coordinates, if specified, indentify
	// the the position up to which the source instance should be cloned. If
	// not specified, the source instance is cloned up to the most recent
	// binary log coordintes.
	BinLogCoordinates *BinLogCoordinates `json:"binLogCoordinates,omitempty"`

	// DestinationInstanceName: Name of the Cloud SQL instance to be created
	// as a clone.
	DestinationInstanceName string `json:"destinationInstanceName,omitempty"`

	// Kind: This is always sql#cloneContext.
	Kind string `json:"kind,omitempty"`
}

type Database struct {
	// Charset: The MySQL charset value.
	Charset string `json:"charset,omitempty"`

	// Collation: The MySQL collation value.
	Collation string `json:"collation,omitempty"`

	// Etag: HTTP 1.1 Entity tag for the resource.
	Etag string `json:"etag,omitempty"`

	// Instance: The name of the Cloud SQL instance. This does not include
	// the project ID.
	Instance string `json:"instance,omitempty"`

	// Kind: This is always sql#database.
	Kind string `json:"kind,omitempty"`

	// Name: The name of the database in the Cloud SQL instance. This does
	// not include the project ID or instance name.
	Name string `json:"name,omitempty"`

	// Project: The project ID of the project containing the Cloud SQL
	// database. The Google apps domain is prefixed if applicable.
	Project string `json:"project,omitempty"`

	// SelfLink: The URI of this resource.
	SelfLink string `json:"selfLink,omitempty"`
}

type DatabaseFlags struct {
	// Name: The name of the flag. These flags are passed at instance
	// startup, so include both MySQL server options and MySQL system
	// variables. Flags should be specified with underscores, not hyphens.
	// Refer to the official MySQL documentation on server options and
	// system variables for descriptions of what these flags do. Acceptable
	// values are:  character_set_server utf8 or utf8mb4 event_scheduler on
	// or off (Note: The event scheduler will only work reliably if the
	// instance activationPolicy is set to ALWAYS) general_log on or off
	// group_concat_max_len 4..17179869184 innodb_flush_log_at_trx_commit
	// 0..2 innodb_lock_wait_timeout 1..1073741824
	// log_bin_trust_function_creators on or off log_output Can be either
	// TABLE or NONE, FILE is not supported log_queries_not_using_indexes on
	// or off long_query_time 0..30000000 lower_case_table_names 0..2
	// max_allowed_packet 16384..1073741824 read_only on or off
	// skip_show_database on or off slow_query_log on or off. If set to on,
	// you must also set the log_output flag to TABLE to receive logs.
	// wait_timeout 1..31536000
	Name string `json:"name,omitempty"`

	// Value: The value of the flag. Booleans should be set using 1 for
	// true, and 0 for false. This field must be omitted if the flag doesn't
	// take a value.
	Value string `json:"value,omitempty"`
}

type DatabaseInstance struct {
	// CurrentDiskSize: The current disk usage of the instance in bytes.
	CurrentDiskSize int64 `json:"currentDiskSize,omitempty,string"`

	// DatabaseVersion: The database engine type and version. Can be
	// MYSQL_5_5 or MYSQL_5_6. Defaults to MYSQL_5_5. The databaseVersion
	// can not be changed after instance creation.
	DatabaseVersion string `json:"databaseVersion,omitempty"`

	// Etag: HTTP 1.1 Entity tag for the resource.
	Etag string `json:"etag,omitempty"`

	// InstanceType: The instance type. This can be one of the
	// following.
	// CLOUD_SQL_INSTANCE: A Cloud SQL instance that is not
	// replicating from a master.
	// READ_REPLICA_INSTANCE: A Cloud SQL
	// instance configured as a read-replica.
	InstanceType string `json:"instanceType,omitempty"`

	// IpAddresses: The assigned IP addresses for the instance.
	IpAddresses []*IpMapping `json:"ipAddresses,omitempty"`

	// Ipv6Address: The IPv6 address assigned to the instance.
	Ipv6Address string `json:"ipv6Address,omitempty"`

	// Kind: This is always sql#instance.
	Kind string `json:"kind,omitempty"`

	// MasterInstanceName: The name of the instance which will act as master
	// in the replication setup.
	MasterInstanceName string `json:"masterInstanceName,omitempty"`

	// MaxDiskSize: The maximum disk size of the instance in bytes.
	MaxDiskSize int64 `json:"maxDiskSize,omitempty,string"`

	// Name: Name of the Cloud SQL instance. This does not include the
	// project ID.
	Name string `json:"name,omitempty"`

	// Project: The project ID of the project containing the Cloud SQL
	// instance. The Google apps domain is prefixed if applicable.
	Project string `json:"project,omitempty"`

	// Region: The geographical region. Can be us-central, asia-east1 or
	// europe-west1. Defaults to us-central. The region can not be changed
	// after instance creation.
	Region string `json:"region,omitempty"`

	// ReplicaNames: The replicas of the instance.
	ReplicaNames []string `json:"replicaNames,omitempty"`

	// SelfLink: The URI of this resource.
	SelfLink string `json:"selfLink,omitempty"`

	// ServerCaCert: SSL configuration.
	ServerCaCert *SslCert `json:"serverCaCert,omitempty"`

	// ServiceAccountEmailAddress: The service account email address
	// assigned to the instance.
	ServiceAccountEmailAddress string `json:"serviceAccountEmailAddress,omitempty"`

	// Settings: The user settings.
	Settings *Settings `json:"settings,omitempty"`

	// State: The current serving state of the Cloud SQL instance. This can
	// be one of the following.
	// RUNNABLE: The instance is running, or is
	// ready to run when accessed.
	// SUSPENDED: The instance is not available,
	// for example due to problems with billing.
	// PENDING_CREATE: The
	// instance is being created.
	// MAINTENANCE: The instance is down for
	// maintenance.
	// UNKNOWN_STATE: The state of the instance is unknown.
	State string `json:"state,omitempty"`
}

type DatabasesListResponse struct {
	// Items: List of database resources in the instance.
	Items []*Database `json:"items,omitempty"`

	// Kind: This is always sql#databasesList.
	Kind string `json:"kind,omitempty"`
}

type ExportContext struct {
	// CsvExportOptions: Options for exporting data as CSV.
	CsvExportOptions *ExportContextCsvExportOptions `json:"csvExportOptions,omitempty"`

	// Databases: Databases (for example, guestbook) from which the export
	// is made. If fileType is SQL and no database is specified, all
	// databases are exported. If fileType is CSV, you can optionally
	// specify at most one database to export. If
	// csvExportOptions.selectQuery also specifies the database, this field
	// will be ignored.
	Databases []string `json:"databases,omitempty"`

	// FileType: The file type for the specified uri.
	// SQL: The file contains
	// SQL statements.
	// CSV: The file contains CSV data.
	FileType string `json:"fileType,omitempty"`

	// Kind: This is always sql#exportContext.
	Kind string `json:"kind,omitempty"`

	// SqlExportOptions: Options for exporting data as SQL statements.
	SqlExportOptions *ExportContextSqlExportOptions `json:"sqlExportOptions,omitempty"`

	// Uri: The path to the file in Google Cloud Storage where the export
	// will be stored. The URI is in the form gs://bucketName/fileName. If
	// the file already exists, the operation fails. If fileType is SQL and
	// the filename ends with .gz, the contents are compressed.
	Uri string `json:"uri,omitempty"`
}

type ExportContextCsvExportOptions struct {
	// SelectQuery: The select query used to extract the data.
	SelectQuery string `json:"selectQuery,omitempty"`
}

type ExportContextSqlExportOptions struct {
	// Tables: Tables to export, or that were exported, from the specified
	// database. If you specify tables, specify one and only one database.
	Tables []string `json:"tables,omitempty"`
}

type Flag struct {
	// AllowedStringValues: For STRING flags, a list of strings that the
	// value can be set to.
	AllowedStringValues []string `json:"allowedStringValues,omitempty"`

	// AppliesTo: The database version this flag applies to. Currently this
	// can only be [MYSQL_5_5].
	AppliesTo []string `json:"appliesTo,omitempty"`

	// Kind: This is always sql#flag.
	Kind string `json:"kind,omitempty"`

	// MaxValue: For INTEGER flags, the maximum allowed value.
	MaxValue int64 `json:"maxValue,omitempty,string"`

	// MinValue: For INTEGER flags, the minimum allowed value.
	MinValue int64 `json:"minValue,omitempty,string"`

	// Name: This is the name of the flag. Flag names always use
	// underscores, not hyphens, e.g. max_allowed_packet
	Name string `json:"name,omitempty"`

	// Type: The type of the flag. Flags are typed to being BOOLEAN, STRING,
	// INTEGER or NONE. NONE is used for flags which do not take a value,
	// such as skip_grant_tables.
	Type string `json:"type,omitempty"`
}

type FlagsListResponse struct {
	// Items: List of flags.
	Items []*Flag `json:"items,omitempty"`

	// Kind: This is always sql#flagsList.
	Kind string `json:"kind,omitempty"`
}

type ImportContext struct {
	// CsvImportOptions: Options for importing data as CSV.
	CsvImportOptions *ImportContextCsvImportOptions `json:"csvImportOptions,omitempty"`

	// Database: The database (for example, guestbook) to which the import
	// is made. If fileType is SQL and no database is specified, it is
	// assumed that the database is specified in the file to be imported. If
	// fileType is CSV, it must be specified.
	Database string `json:"database,omitempty"`

	// FileType: The file type for the specified uri.
	// SQL: The file contains
	// SQL statements.
	// CSV: The file contains CSV data.
	FileType string `json:"fileType,omitempty"`

	// Kind: This is always sql#importContext.
	Kind string `json:"kind,omitempty"`

	// Uri: A path to the file in Google Cloud Storage from which the import
	// is made. The URI is in the form gs://bucketName/fileName. Compressed
	// gzip files (.gz) are supported when fileType is SQL.
	Uri string `json:"uri,omitempty"`
}

type ImportContextCsvImportOptions struct {
	// Columns: The columns to which CSV data is imported. If not specified,
	// all columns of the database table are loaded with CSV data.
	Columns []string `json:"columns,omitempty"`

	// Table: The table to which CSV data is imported.
	Table string `json:"table,omitempty"`
}

type InstancesCloneRequest struct {
	// CloneContext: Contains details about the clone operation.
	CloneContext *CloneContext `json:"cloneContext,omitempty"`
}

type InstancesExportRequest struct {
	// ExportContext: Contains details about the export operation.
	ExportContext *ExportContext `json:"exportContext,omitempty"`
}

type InstancesImportRequest struct {
	// ImportContext: Contains details about the import operation.
	ImportContext *ImportContext `json:"importContext,omitempty"`
}

type InstancesListResponse struct {
	// Items: List of database instance resources.
	Items []*DatabaseInstance `json:"items,omitempty"`

	// Kind: This is always sql#instancesList.
	Kind string `json:"kind,omitempty"`

	// NextPageToken: The continuation token, used to page through large
	// result sets. Provide this value in a subsequent request to return the
	// next page of results.
	NextPageToken string `json:"nextPageToken,omitempty"`
}

type InstancesRestoreBackupRequest struct {
	// RestoreBackupContext: Parameters required to perform the restore
	// backup operation.
	RestoreBackupContext *RestoreBackupContext `json:"restoreBackupContext,omitempty"`
}

type IpConfiguration struct {
	// AuthorizedNetworks: The list of external networks that are allowed to
	// connect to the instance using the IP. In CIDR notation, also known as
	// 'slash' notation (e.g. 192.168.100.0/24).
	AuthorizedNetworks []*AclEntry `json:"authorizedNetworks,omitempty"`

	// Ipv4Enabled: Whether the instance should be assigned an IP address or
	// not.
	Ipv4Enabled bool `json:"ipv4Enabled,omitempty"`

	// RequireSsl: Whether the mysqld should default to 'REQUIRE X509' for
	// users connecting over IP.
	RequireSsl bool `json:"requireSsl,omitempty"`
}

type IpMapping struct {
	// IpAddress: The IP address assigned.
	IpAddress string `json:"ipAddress,omitempty"`

	// TimeToRetire: The due time for this IP to be retired in RFC 3339
	// format, for example 2012-11-15T16:19:00.094Z. This field is only
	// available when the IP is scheduled to be retired.
	TimeToRetire string `json:"timeToRetire,omitempty"`
}

type LocationPreference struct {
	// FollowGaeApplication: The AppEngine application to follow, it must be
	// in the same region as the Cloud SQL instance.
	FollowGaeApplication string `json:"followGaeApplication,omitempty"`

	// Kind: This is always sql#locationPreference.
	Kind string `json:"kind,omitempty"`

	// Zone: The preferred Compute Engine zone (e.g. us-centra1-a,
	// us-central1-b, etc.).
	Zone string `json:"zone,omitempty"`
}

type Operation struct {
	// EndTime: The time this operation finished in UTC timezone in RFC 3339
	// format, for example 2012-11-15T16:19:00.094Z.
	EndTime string `json:"endTime,omitempty"`

	// Error: If errors occurred during processing of this operation, this
	// field will be populated.
	Error *OperationError1 `json:"error,omitempty"`

	// ExportContext: The context for export operation, if applicable.
	ExportContext *ExportContext `json:"exportContext,omitempty"`

	// ImportContext: The context for import operation, if applicable.
	ImportContext *ImportContext `json:"importContext,omitempty"`

	// InsertTime: The time this operation was enqueued in UTC timezone in
	// RFC 3339 format, for example 2012-11-15T16:19:00.094Z.
	InsertTime string `json:"insertTime,omitempty"`

	// Kind: This is always sql#instanceOperation.
	Kind string `json:"kind,omitempty"`

	// Name: An identifier that uniquely identifies the operation. You can
	// use this identifier to retrieve the Operations resource that has
	// information about the operation.
	Name string `json:"name,omitempty"`

	// OperationType: TODO(b/18431310): update this list to reflect current
	// values. The type of the operation. Valid values are CREATE, DELETE,
	// UPDATE, RESTART, IMPORT, EXPORT, BACKUP_VOLUME, RESTORE_VOLUME.
	OperationType string `json:"operationType,omitempty"`

	// SelfLink: The URI of this resource.
	SelfLink string `json:"selfLink,omitempty"`

	// StartTime: The time this operation actually started in UTC timezone
	// in RFC 3339 format, for example 2012-11-15T16:19:00.094Z.
	StartTime string `json:"startTime,omitempty"`

	// Status: The status of an operation. Valid values are PENDING,
	// RUNNING, DONE, UNKNOWN.
	Status string `json:"status,omitempty"`

	// TargetId: Name of the database instance related to this operation.
	TargetId string `json:"targetId,omitempty"`

	// TargetLink: The URI of the instance related to the operation.
	TargetLink string `json:"targetLink,omitempty"`

	// TargetProject: The project ID of the target instance related to this
	// operation.
	TargetProject string `json:"targetProject,omitempty"`

	// User: The email address of the user who initiated this operation.
	User string `json:"user,omitempty"`
}

type OperationError1 struct {
	// Errors: The list of errors encountered while processing this
	// operation.
	Errors []*OperationError `json:"errors,omitempty"`
}

type OperationError struct {
	// Code: Identifies the specific error that occurred.
	Code string `json:"code,omitempty"`

	// Kind: This is always sql#operationError.
	Kind string `json:"kind,omitempty"`

	// Message: Additional information about the error encountered.
	Message string `json:"message,omitempty"`
}

type OperationsListResponse struct {
	// Items: List of operation resources.
	Items []*Operation `json:"items,omitempty"`

	// Kind: This is always sql#operationsList.
	Kind string `json:"kind,omitempty"`

	// NextPageToken: The continuation token, used to page through large
	// result sets. Provide this value in a subsequent request to return the
	// next page of results.
	NextPageToken string `json:"nextPageToken,omitempty"`
}

type RestoreBackupContext struct {
	// BackupRunId: The ID of the backup run to restore from.
	BackupRunId int64 `json:"backupRunId,omitempty,string"`

	// Kind: This is always sql#restoreBackupContext.
	Kind string `json:"kind,omitempty"`
}

type Settings struct {
	// ActivationPolicy: The activation policy for this instance. This
	// specifies when the instance should be activated and is applicable
	// only when the instance state is RUNNABLE. This can be one of the
	// following.
	// ALWAYS: The instance should always be active.
	// NEVER: The
	// instance should never be activated.
	// ON_DEMAND: The instance is
	// activated upon receiving requests.
	ActivationPolicy string `json:"activationPolicy,omitempty"`

	// AuthorizedGaeApplications: The App Engine app IDs that can access
	// this instance.
	AuthorizedGaeApplications []string `json:"authorizedGaeApplications,omitempty"`

	// BackupConfiguration: The daily backup configuration for the instance.
	BackupConfiguration *BackupConfiguration `json:"backupConfiguration,omitempty"`

	// DatabaseFlags: The database flags passed to the instance at startup.
	DatabaseFlags []*DatabaseFlags `json:"databaseFlags,omitempty"`

	// DatabaseReplicationEnabled: Configuration specific to read replica
	// instances. Indicates whether replication is enabled or not.
	DatabaseReplicationEnabled bool `json:"databaseReplicationEnabled,omitempty"`

	// IpConfiguration: The settings for IP Management. This allows to
	// enable or disable the instance IP and manage which external networks
	// can connect to the instance.
	IpConfiguration *IpConfiguration `json:"ipConfiguration,omitempty"`

	// Kind: This is always sql#settings.
	Kind string `json:"kind,omitempty"`

	// LocationPreference: The location preference settings. This allows the
	// instance to be located as near as possible to either an App Engine
	// app or GCE zone for better performance.
	LocationPreference *LocationPreference `json:"locationPreference,omitempty"`

	// PricingPlan: The pricing plan for this instance. This can be either
	// PER_USE or PACKAGE.
	PricingPlan string `json:"pricingPlan,omitempty"`

	// ReplicationType: The type of replication this instance uses. This can
	// be either ASYNCHRONOUS or SYNCHRONOUS.
	ReplicationType string `json:"replicationType,omitempty"`

	// SettingsVersion: The version of instance settings. This is a required
	// field for update method to make sure concurrent updates are handled
	// properly. During update, use the most recent settingsVersion value
	// for this instance and do not try to update this value.
	SettingsVersion int64 `json:"settingsVersion,omitempty,string"`

	// Tier: The tier of service for this instance, for example D1, D2. For
	// more information, see pricing.
	Tier string `json:"tier,omitempty"`
}

type SslCert struct {
	// Cert: PEM representation.
	Cert string `json:"cert,omitempty"`

	// CertSerialNumber: Serial number, as extracted from the certificate.
	CertSerialNumber string `json:"certSerialNumber,omitempty"`

	// CommonName: User supplied name. Constrained to [a-zA-Z.-_ ]+.
	CommonName string `json:"commonName,omitempty"`

	// CreateTime: The time when the certificate was created in RFC 3339
	// format, for example 2012-11-15T16:19:00.094Z
	CreateTime string `json:"createTime,omitempty"`

	// ExpirationTime: The time when the certificate expires in RFC 3339
	// format, for example 2012-11-15T16:19:00.094Z.
	ExpirationTime string `json:"expirationTime,omitempty"`

	// Instance: Name of the database instance.
	Instance string `json:"instance,omitempty"`

	// Kind: This is always sql#sslCert.
	Kind string `json:"kind,omitempty"`

	SelfLink string `json:"selfLink,omitempty"`

	// Sha1Fingerprint: Sha1 Fingerprint.
	Sha1Fingerprint string `json:"sha1Fingerprint,omitempty"`
}

type SslCertDetail struct {
	// CertInfo: The public information about the cert.
	CertInfo *SslCert `json:"certInfo,omitempty"`

	// CertPrivateKey: The private key for the client cert, in pem format.
	// Keep private in order to protect your security.
	CertPrivateKey string `json:"certPrivateKey,omitempty"`
}

type SslCertsInsertRequest struct {
	// CommonName: User supplied name. Must be a distinct name from the
	// other certificates for this instance. New certificates will not be
	// usable until the instance is restarted.
	CommonName string `json:"commonName,omitempty"`
}

type SslCertsInsertResponse struct {
	// ClientCert: The new client certificate and private key. The new
	// certificate will not work until the instance is restarted.
	ClientCert *SslCertDetail `json:"clientCert,omitempty"`

	// Kind: This is always sql#sslCertsInsert.
	Kind string `json:"kind,omitempty"`

	// ServerCaCert: The server Certificate Authority's certificate. If this
	// is missing you can force a new one to be generated by calling
	// resetSslConfig method on instances resource.
	ServerCaCert *SslCert `json:"serverCaCert,omitempty"`
}

type SslCertsListResponse struct {
	// Items: List of client certificates for the instance.
	Items []*SslCert `json:"items,omitempty"`

	// Kind: This is always sql#sslCertsList.
	Kind string `json:"kind,omitempty"`
}

type Tier struct {
	// DiskQuota: The maximum disk size of this tier in bytes.
	DiskQuota int64 `json:"DiskQuota,omitempty,string"`

	// RAM: The maximum RAM usage of this tier in bytes.
	RAM int64 `json:"RAM,omitempty,string"`

	// Kind: This is always sql#tier.
	Kind string `json:"kind,omitempty"`

	// Region: The applicable regions for this tier. Can be us-east1,
	// europe-west1 or asia-east1.
	Region []string `json:"region,omitempty"`

	// Tier: An identifier for the service tier, for example D1, D2 etc. For
	// related information, see Pricing.
	Tier string `json:"tier,omitempty"`
}

type TiersListResponse struct {
	// Items: List of tiers.
	Items []*Tier `json:"items,omitempty"`

	// Kind: This is always sql#tiersList.
	Kind string `json:"kind,omitempty"`
}

type User struct {
	// Etag: HTTP 1.1 Entity tag for the resource.
	Etag string `json:"etag,omitempty"`

	// Host: The host name from which the user can connect. For insert
	// operations, host is set to '%'. For update operations, host is
	// specified as part of the request URL. The host name is not mutable
	// with this API.
	Host string `json:"host,omitempty"`

	// Instance: The name of the Cloud SQL instance. This does not include
	// the project ID. Can be omitted for update since it is already
	// specified on the URL.
	Instance string `json:"instance,omitempty"`

	// Kind: This is always sql#user.
	Kind string `json:"kind,omitempty"`

	// Name: The name of the user in the Cloud SQL instance. Can be omitted
	// for update since it is already specified on the URL.
	Name string `json:"name,omitempty"`

	// Password: The password for the user.
	Password string `json:"password,omitempty"`

	// Project: The project ID of the project containing the Cloud SQL
	// database. The Google apps domain is prefixed if applicable. Can be
	// omitted for update since it is already specified on the URL.
	Project string `json:"project,omitempty"`
}

type UsersListResponse struct {
	// Items: List of user resources in the instance.
	Items []*User `json:"items,omitempty"`

	// Kind: This is always sql#usersList.
	Kind string `json:"kind,omitempty"`

	// NextPageToken: An identifier that uniquely identifies the operation.
	// You can use this identifier to retrieve the Operations resource that
	// has information about the operation.
	NextPageToken string `json:"nextPageToken,omitempty"`
}

// method id "sql.backupRuns.get":

type BackupRunsGetCall struct {
	s        *Service
	project  string
	instance string
	id       int64
	opt_     map[string]interface{}
}

// Get: Retrieves a resource containing information about a backup run.
func (r *BackupRunsService) Get(project string, instance string, id int64) *BackupRunsGetCall {
	c := &BackupRunsGetCall{s: r.s, opt_: make(map[string]interface{})}
	c.project = project
	c.instance = instance
	c.id = id
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *BackupRunsGetCall) Fields(s ...googleapi.Field) *BackupRunsGetCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *BackupRunsGetCall) Do() (*BackupRun, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "projects/{project}/instances/{instance}/backupRuns/{id}")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"project":  c.project,
		"instance": c.instance,
		"id":       strconv.FormatInt(c.id, 10),
	})
	req.Header.Set("User-Agent", c.s.userAgent())
	res, err := c.s.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	var ret *BackupRun
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Retrieves a resource containing information about a backup run.",
	//   "httpMethod": "GET",
	//   "id": "sql.backupRuns.get",
	//   "parameterOrder": [
	//     "project",
	//     "instance",
	//     "id"
	//   ],
	//   "parameters": {
	//     "id": {
	//       "description": "The ID of this Backup Run.",
	//       "format": "int64",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "instance": {
	//       "description": "Cloud SQL instance ID. This does not include the project ID.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "project": {
	//       "description": "Project ID of the project that contains the instance.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "projects/{project}/instances/{instance}/backupRuns/{id}",
	//   "response": {
	//     "$ref": "BackupRun"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/sqlservice.admin"
	//   ]
	// }

}

// method id "sql.backupRuns.list":

type BackupRunsListCall struct {
	s        *Service
	project  string
	instance string
	opt_     map[string]interface{}
}

// List: Lists all backup runs associated with a given instance and
// configuration in the reverse chronological order of the enqueued
// time.
func (r *BackupRunsService) List(project string, instance string) *BackupRunsListCall {
	c := &BackupRunsListCall{s: r.s, opt_: make(map[string]interface{})}
	c.project = project
	c.instance = instance
	return c
}

// MaxResults sets the optional parameter "maxResults": Maximum number
// of backup runs per response.
func (c *BackupRunsListCall) MaxResults(maxResults int64) *BackupRunsListCall {
	c.opt_["maxResults"] = maxResults
	return c
}

// PageToken sets the optional parameter "pageToken": A
// previously-returned page token representing part of the larger set of
// results to view.
func (c *BackupRunsListCall) PageToken(pageToken string) *BackupRunsListCall {
	c.opt_["pageToken"] = pageToken
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *BackupRunsListCall) Fields(s ...googleapi.Field) *BackupRunsListCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *BackupRunsListCall) Do() (*BackupRunsListResponse, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["maxResults"]; ok {
		params.Set("maxResults", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["pageToken"]; ok {
		params.Set("pageToken", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "projects/{project}/instances/{instance}/backupRuns")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"project":  c.project,
		"instance": c.instance,
	})
	req.Header.Set("User-Agent", c.s.userAgent())
	res, err := c.s.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	var ret *BackupRunsListResponse
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Lists all backup runs associated with a given instance and configuration in the reverse chronological order of the enqueued time.",
	//   "httpMethod": "GET",
	//   "id": "sql.backupRuns.list",
	//   "parameterOrder": [
	//     "project",
	//     "instance"
	//   ],
	//   "parameters": {
	//     "instance": {
	//       "description": "Cloud SQL instance ID. This does not include the project ID.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "maxResults": {
	//       "description": "Maximum number of backup runs per response.",
	//       "format": "int32",
	//       "location": "query",
	//       "type": "integer"
	//     },
	//     "pageToken": {
	//       "description": "A previously-returned page token representing part of the larger set of results to view.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "project": {
	//       "description": "Project ID of the project that contains the instance.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "projects/{project}/instances/{instance}/backupRuns",
	//   "response": {
	//     "$ref": "BackupRunsListResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/sqlservice.admin"
	//   ]
	// }

}

// method id "sql.databases.delete":

type DatabasesDeleteCall struct {
	s        *Service
	project  string
	instance string
	database string
	opt_     map[string]interface{}
}

// Delete: Deletes a resource containing information about a database
// inside a Cloud SQL instance.
func (r *DatabasesService) Delete(project string, instance string, database string) *DatabasesDeleteCall {
	c := &DatabasesDeleteCall{s: r.s, opt_: make(map[string]interface{})}
	c.project = project
	c.instance = instance
	c.database = database
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *DatabasesDeleteCall) Fields(s ...googleapi.Field) *DatabasesDeleteCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *DatabasesDeleteCall) Do() (*Operation, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "projects/{project}/instances/{instance}/databases/{database}")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("DELETE", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"project":  c.project,
		"instance": c.instance,
		"database": c.database,
	})
	req.Header.Set("User-Agent", c.s.userAgent())
	res, err := c.s.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	var ret *Operation
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Deletes a resource containing information about a database inside a Cloud SQL instance.",
	//   "httpMethod": "DELETE",
	//   "id": "sql.databases.delete",
	//   "parameterOrder": [
	//     "project",
	//     "instance",
	//     "database"
	//   ],
	//   "parameters": {
	//     "database": {
	//       "description": "Name of the database to be deleted in the instance.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "instance": {
	//       "description": "Database instance ID. This does not include the project ID.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "project": {
	//       "description": "Project ID of the project that contains the instance.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "projects/{project}/instances/{instance}/databases/{database}",
	//   "response": {
	//     "$ref": "Operation"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/sqlservice.admin"
	//   ]
	// }

}

// method id "sql.databases.get":

type DatabasesGetCall struct {
	s        *Service
	project  string
	instance string
	database string
	opt_     map[string]interface{}
}

// Get: Retrieves a resource containing information about a database
// inside a Cloud SQL instance.
func (r *DatabasesService) Get(project string, instance string, database string) *DatabasesGetCall {
	c := &DatabasesGetCall{s: r.s, opt_: make(map[string]interface{})}
	c.project = project
	c.instance = instance
	c.database = database
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *DatabasesGetCall) Fields(s ...googleapi.Field) *DatabasesGetCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *DatabasesGetCall) Do() (*Database, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "projects/{project}/instances/{instance}/databases/{database}")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"project":  c.project,
		"instance": c.instance,
		"database": c.database,
	})
	req.Header.Set("User-Agent", c.s.userAgent())
	res, err := c.s.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	var ret *Database
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Retrieves a resource containing information about a database inside a Cloud SQL instance.",
	//   "httpMethod": "GET",
	//   "id": "sql.databases.get",
	//   "parameterOrder": [
	//     "project",
	//     "instance",
	//     "database"
	//   ],
	//   "parameters": {
	//     "database": {
	//       "description": "Name of the database in the instance.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "instance": {
	//       "description": "Database instance ID. This does not include the project ID.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "project": {
	//       "description": "Project ID of the project that contains the instance.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "projects/{project}/instances/{instance}/databases/{database}",
	//   "response": {
	//     "$ref": "Database"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/sqlservice.admin"
	//   ]
	// }

}

// method id "sql.databases.insert":

type DatabasesInsertCall struct {
	s        *Service
	project  string
	instance string
	database *Database
	opt_     map[string]interface{}
}

// Insert: Inserts a resource containing information about a database
// inside a Cloud SQL instance.
func (r *DatabasesService) Insert(project string, instance string, database *Database) *DatabasesInsertCall {
	c := &DatabasesInsertCall{s: r.s, opt_: make(map[string]interface{})}
	c.project = project
	c.instance = instance
	c.database = database
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *DatabasesInsertCall) Fields(s ...googleapi.Field) *DatabasesInsertCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *DatabasesInsertCall) Do() (*Operation, error) {
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.database)
	if err != nil {
		return nil, err
	}
	ctype := "application/json"
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "projects/{project}/instances/{instance}/databases")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"project":  c.project,
		"instance": c.instance,
	})
	req.Header.Set("Content-Type", ctype)
	req.Header.Set("User-Agent", c.s.userAgent())
	res, err := c.s.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	var ret *Operation
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Inserts a resource containing information about a database inside a Cloud SQL instance.",
	//   "httpMethod": "POST",
	//   "id": "sql.databases.insert",
	//   "parameterOrder": [
	//     "project",
	//     "instance"
	//   ],
	//   "parameters": {
	//     "instance": {
	//       "description": "Database instance ID. This does not include the project ID.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "project": {
	//       "description": "Project ID of the project that contains the instance.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "projects/{project}/instances/{instance}/databases",
	//   "request": {
	//     "$ref": "Database"
	//   },
	//   "response": {
	//     "$ref": "Operation"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/sqlservice.admin"
	//   ]
	// }

}

// method id "sql.databases.list":

type DatabasesListCall struct {
	s        *Service
	project  string
	instance string
	opt_     map[string]interface{}
}

// List: Lists databases in the specified Cloud SQL instance.
func (r *DatabasesService) List(project string, instance string) *DatabasesListCall {
	c := &DatabasesListCall{s: r.s, opt_: make(map[string]interface{})}
	c.project = project
	c.instance = instance
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *DatabasesListCall) Fields(s ...googleapi.Field) *DatabasesListCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *DatabasesListCall) Do() (*DatabasesListResponse, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "projects/{project}/instances/{instance}/databases")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"project":  c.project,
		"instance": c.instance,
	})
	req.Header.Set("User-Agent", c.s.userAgent())
	res, err := c.s.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	var ret *DatabasesListResponse
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Lists databases in the specified Cloud SQL instance.",
	//   "httpMethod": "GET",
	//   "id": "sql.databases.list",
	//   "parameterOrder": [
	//     "project",
	//     "instance"
	//   ],
	//   "parameters": {
	//     "instance": {
	//       "description": "Cloud SQL instance ID. This does not include the project ID.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "project": {
	//       "description": "Project ID of the project for which to list Cloud SQL instances.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "projects/{project}/instances/{instance}/databases",
	//   "response": {
	//     "$ref": "DatabasesListResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/sqlservice.admin"
	//   ]
	// }

}

// method id "sql.databases.patch":

type DatabasesPatchCall struct {
	s         *Service
	project   string
	instance  string
	database  string
	database2 *Database
	opt_      map[string]interface{}
}

// Patch: Updates a resource containing information about a database
// inside a Cloud SQL instance. This method supports patch semantics.
func (r *DatabasesService) Patch(project string, instance string, database string, database2 *Database) *DatabasesPatchCall {
	c := &DatabasesPatchCall{s: r.s, opt_: make(map[string]interface{})}
	c.project = project
	c.instance = instance
	c.database = database
	c.database2 = database2
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *DatabasesPatchCall) Fields(s ...googleapi.Field) *DatabasesPatchCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *DatabasesPatchCall) Do() (*Operation, error) {
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.database2)
	if err != nil {
		return nil, err
	}
	ctype := "application/json"
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "projects/{project}/instances/{instance}/databases/{database}")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("PATCH", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"project":  c.project,
		"instance": c.instance,
		"database": c.database,
	})
	req.Header.Set("Content-Type", ctype)
	req.Header.Set("User-Agent", c.s.userAgent())
	res, err := c.s.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	var ret *Operation
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Updates a resource containing information about a database inside a Cloud SQL instance. This method supports patch semantics.",
	//   "httpMethod": "PATCH",
	//   "id": "sql.databases.patch",
	//   "parameterOrder": [
	//     "project",
	//     "instance",
	//     "database"
	//   ],
	//   "parameters": {
	//     "database": {
	//       "description": "Name of the database to be updated in the instance.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "instance": {
	//       "description": "Database instance ID. This does not include the project ID.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "project": {
	//       "description": "Project ID of the project that contains the instance.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "projects/{project}/instances/{instance}/databases/{database}",
	//   "request": {
	//     "$ref": "Database"
	//   },
	//   "response": {
	//     "$ref": "Operation"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/sqlservice.admin"
	//   ]
	// }

}

// method id "sql.databases.update":

type DatabasesUpdateCall struct {
	s         *Service
	project   string
	instance  string
	database  string
	database2 *Database
	opt_      map[string]interface{}
}

// Update: Updates a resource containing information about a database
// inside a Cloud SQL instance.
func (r *DatabasesService) Update(project string, instance string, database string, database2 *Database) *DatabasesUpdateCall {
	c := &DatabasesUpdateCall{s: r.s, opt_: make(map[string]interface{})}
	c.project = project
	c.instance = instance
	c.database = database
	c.database2 = database2
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *DatabasesUpdateCall) Fields(s ...googleapi.Field) *DatabasesUpdateCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *DatabasesUpdateCall) Do() (*Operation, error) {
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.database2)
	if err != nil {
		return nil, err
	}
	ctype := "application/json"
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "projects/{project}/instances/{instance}/databases/{database}")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("PUT", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"project":  c.project,
		"instance": c.instance,
		"database": c.database,
	})
	req.Header.Set("Content-Type", ctype)
	req.Header.Set("User-Agent", c.s.userAgent())
	res, err := c.s.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	var ret *Operation
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Updates a resource containing information about a database inside a Cloud SQL instance.",
	//   "httpMethod": "PUT",
	//   "id": "sql.databases.update",
	//   "parameterOrder": [
	//     "project",
	//     "instance",
	//     "database"
	//   ],
	//   "parameters": {
	//     "database": {
	//       "description": "Name of the database to be updated in the instance.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "instance": {
	//       "description": "Database instance ID. This does not include the project ID.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "project": {
	//       "description": "Project ID of the project that contains the instance.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "projects/{project}/instances/{instance}/databases/{database}",
	//   "request": {
	//     "$ref": "Database"
	//   },
	//   "response": {
	//     "$ref": "Operation"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/sqlservice.admin"
	//   ]
	// }

}

// method id "sql.flags.list":

type FlagsListCall struct {
	s    *Service
	opt_ map[string]interface{}
}

// List: List all available database flags for Google Cloud SQL
// instances.
func (r *FlagsService) List() *FlagsListCall {
	c := &FlagsListCall{s: r.s, opt_: make(map[string]interface{})}
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *FlagsListCall) Fields(s ...googleapi.Field) *FlagsListCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *FlagsListCall) Do() (*FlagsListResponse, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "flags")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	googleapi.SetOpaque(req.URL)
	req.Header.Set("User-Agent", c.s.userAgent())
	res, err := c.s.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	var ret *FlagsListResponse
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "List all available database flags for Google Cloud SQL instances.",
	//   "httpMethod": "GET",
	//   "id": "sql.flags.list",
	//   "path": "flags",
	//   "response": {
	//     "$ref": "FlagsListResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/sqlservice.admin"
	//   ]
	// }

}

// method id "sql.instances.clone":

type InstancesCloneCall struct {
	s                     *Service
	project               string
	instance              string
	instancesclonerequest *InstancesCloneRequest
	opt_                  map[string]interface{}
}

// Clone: Creates a Cloud SQL instance as a clone of the source
// instance.
func (r *InstancesService) Clone(project string, instance string, instancesclonerequest *InstancesCloneRequest) *InstancesCloneCall {
	c := &InstancesCloneCall{s: r.s, opt_: make(map[string]interface{})}
	c.project = project
	c.instance = instance
	c.instancesclonerequest = instancesclonerequest
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *InstancesCloneCall) Fields(s ...googleapi.Field) *InstancesCloneCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *InstancesCloneCall) Do() (*Operation, error) {
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.instancesclonerequest)
	if err != nil {
		return nil, err
	}
	ctype := "application/json"
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "projects/{project}/instances/{instance}/clone")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"project":  c.project,
		"instance": c.instance,
	})
	req.Header.Set("Content-Type", ctype)
	req.Header.Set("User-Agent", c.s.userAgent())
	res, err := c.s.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	var ret *Operation
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Creates a Cloud SQL instance as a clone of the source instance.",
	//   "httpMethod": "POST",
	//   "id": "sql.instances.clone",
	//   "parameterOrder": [
	//     "project",
	//     "instance"
	//   ],
	//   "parameters": {
	//     "instance": {
	//       "description": "The ID of the Cloud SQL instance to be cloned (source). This does not include the project ID.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "project": {
	//       "description": "Project ID of the source as well as the clone Cloud SQL instance.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "projects/{project}/instances/{instance}/clone",
	//   "request": {
	//     "$ref": "InstancesCloneRequest"
	//   },
	//   "response": {
	//     "$ref": "Operation"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/sqlservice.admin"
	//   ]
	// }

}

// method id "sql.instances.delete":

type InstancesDeleteCall struct {
	s        *Service
	project  string
	instance string
	opt_     map[string]interface{}
}

// Delete: Deletes a Cloud SQL instance.
func (r *InstancesService) Delete(project string, instance string) *InstancesDeleteCall {
	c := &InstancesDeleteCall{s: r.s, opt_: make(map[string]interface{})}
	c.project = project
	c.instance = instance
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *InstancesDeleteCall) Fields(s ...googleapi.Field) *InstancesDeleteCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *InstancesDeleteCall) Do() (*Operation, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "projects/{project}/instances/{instance}")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("DELETE", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"project":  c.project,
		"instance": c.instance,
	})
	req.Header.Set("User-Agent", c.s.userAgent())
	res, err := c.s.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	var ret *Operation
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Deletes a Cloud SQL instance.",
	//   "httpMethod": "DELETE",
	//   "id": "sql.instances.delete",
	//   "parameterOrder": [
	//     "project",
	//     "instance"
	//   ],
	//   "parameters": {
	//     "instance": {
	//       "description": "Cloud SQL instance ID. This does not include the project ID.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "project": {
	//       "description": "Project ID of the project that contains the instance to be deleted.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "projects/{project}/instances/{instance}",
	//   "response": {
	//     "$ref": "Operation"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/sqlservice.admin"
	//   ]
	// }

}

// method id "sql.instances.export":

type InstancesExportCall struct {
	s                      *Service
	project                string
	instance               string
	instancesexportrequest *InstancesExportRequest
	opt_                   map[string]interface{}
}

// Export: Exports data from a Cloud SQL instance to a Google Cloud
// Storage bucket as a MySQL dump file.
func (r *InstancesService) Export(project string, instance string, instancesexportrequest *InstancesExportRequest) *InstancesExportCall {
	c := &InstancesExportCall{s: r.s, opt_: make(map[string]interface{})}
	c.project = project
	c.instance = instance
	c.instancesexportrequest = instancesexportrequest
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *InstancesExportCall) Fields(s ...googleapi.Field) *InstancesExportCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *InstancesExportCall) Do() (*Operation, error) {
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.instancesexportrequest)
	if err != nil {
		return nil, err
	}
	ctype := "application/json"
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "projects/{project}/instances/{instance}/export")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"project":  c.project,
		"instance": c.instance,
	})
	req.Header.Set("Content-Type", ctype)
	req.Header.Set("User-Agent", c.s.userAgent())
	res, err := c.s.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	var ret *Operation
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Exports data from a Cloud SQL instance to a Google Cloud Storage bucket as a MySQL dump file.",
	//   "httpMethod": "POST",
	//   "id": "sql.instances.export",
	//   "parameterOrder": [
	//     "project",
	//     "instance"
	//   ],
	//   "parameters": {
	//     "instance": {
	//       "description": "Cloud SQL instance ID. This does not include the project ID.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "project": {
	//       "description": "Project ID of the project that contains the instance to be exported.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "projects/{project}/instances/{instance}/export",
	//   "request": {
	//     "$ref": "InstancesExportRequest"
	//   },
	//   "response": {
	//     "$ref": "Operation"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "sql.instances.get":

type InstancesGetCall struct {
	s        *Service
	project  string
	instance string
	opt_     map[string]interface{}
}

// Get: Retrieves a resource containing information about a Cloud SQL
// instance.
func (r *InstancesService) Get(project string, instance string) *InstancesGetCall {
	c := &InstancesGetCall{s: r.s, opt_: make(map[string]interface{})}
	c.project = project
	c.instance = instance
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *InstancesGetCall) Fields(s ...googleapi.Field) *InstancesGetCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *InstancesGetCall) Do() (*DatabaseInstance, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "projects/{project}/instances/{instance}")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"project":  c.project,
		"instance": c.instance,
	})
	req.Header.Set("User-Agent", c.s.userAgent())
	res, err := c.s.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	var ret *DatabaseInstance
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Retrieves a resource containing information about a Cloud SQL instance.",
	//   "httpMethod": "GET",
	//   "id": "sql.instances.get",
	//   "parameterOrder": [
	//     "project",
	//     "instance"
	//   ],
	//   "parameters": {
	//     "instance": {
	//       "description": "Database instance ID. This does not include the project ID.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "project": {
	//       "description": "Project ID of the project that contains the instance.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "projects/{project}/instances/{instance}",
	//   "response": {
	//     "$ref": "DatabaseInstance"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/sqlservice.admin"
	//   ]
	// }

}

// method id "sql.instances.import":

type InstancesImportCall struct {
	s                      *Service
	project                string
	instance               string
	instancesimportrequest *InstancesImportRequest
	opt_                   map[string]interface{}
}

// Import: Imports data into a Cloud SQL instance from a MySQL dump file
// in Google Cloud Storage.
func (r *InstancesService) Import(project string, instance string, instancesimportrequest *InstancesImportRequest) *InstancesImportCall {
	c := &InstancesImportCall{s: r.s, opt_: make(map[string]interface{})}
	c.project = project
	c.instance = instance
	c.instancesimportrequest = instancesimportrequest
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *InstancesImportCall) Fields(s ...googleapi.Field) *InstancesImportCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *InstancesImportCall) Do() (*Operation, error) {
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.instancesimportrequest)
	if err != nil {
		return nil, err
	}
	ctype := "application/json"
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "projects/{project}/instances/{instance}/import")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"project":  c.project,
		"instance": c.instance,
	})
	req.Header.Set("Content-Type", ctype)
	req.Header.Set("User-Agent", c.s.userAgent())
	res, err := c.s.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	var ret *Operation
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Imports data into a Cloud SQL instance from a MySQL dump file in Google Cloud Storage.",
	//   "httpMethod": "POST",
	//   "id": "sql.instances.import",
	//   "parameterOrder": [
	//     "project",
	//     "instance"
	//   ],
	//   "parameters": {
	//     "instance": {
	//       "description": "Cloud SQL instance ID. This does not include the project ID.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "project": {
	//       "description": "Project ID of the project that contains the instance.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "projects/{project}/instances/{instance}/import",
	//   "request": {
	//     "$ref": "InstancesImportRequest"
	//   },
	//   "response": {
	//     "$ref": "Operation"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform"
	//   ]
	// }

}

// method id "sql.instances.insert":

type InstancesInsertCall struct {
	s                *Service
	project          string
	databaseinstance *DatabaseInstance
	opt_             map[string]interface{}
}

// Insert: Creates a new Cloud SQL instance.
func (r *InstancesService) Insert(project string, databaseinstance *DatabaseInstance) *InstancesInsertCall {
	c := &InstancesInsertCall{s: r.s, opt_: make(map[string]interface{})}
	c.project = project
	c.databaseinstance = databaseinstance
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *InstancesInsertCall) Fields(s ...googleapi.Field) *InstancesInsertCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *InstancesInsertCall) Do() (*Operation, error) {
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.databaseinstance)
	if err != nil {
		return nil, err
	}
	ctype := "application/json"
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "projects/{project}/instances")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"project": c.project,
	})
	req.Header.Set("Content-Type", ctype)
	req.Header.Set("User-Agent", c.s.userAgent())
	res, err := c.s.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	var ret *Operation
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Creates a new Cloud SQL instance.",
	//   "httpMethod": "POST",
	//   "id": "sql.instances.insert",
	//   "parameterOrder": [
	//     "project"
	//   ],
	//   "parameters": {
	//     "project": {
	//       "description": "Project ID of the project to which the newly created Cloud SQL instances should belong.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "projects/{project}/instances",
	//   "request": {
	//     "$ref": "DatabaseInstance"
	//   },
	//   "response": {
	//     "$ref": "Operation"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/sqlservice.admin"
	//   ]
	// }

}

// method id "sql.instances.list":

type InstancesListCall struct {
	s       *Service
	project string
	opt_    map[string]interface{}
}

// List: Lists instances under a given project in the alphabetical order
// of the instance name.
func (r *InstancesService) List(project string) *InstancesListCall {
	c := &InstancesListCall{s: r.s, opt_: make(map[string]interface{})}
	c.project = project
	return c
}

// MaxResults sets the optional parameter "maxResults": The maximum
// number of results to return per response.
func (c *InstancesListCall) MaxResults(maxResults int64) *InstancesListCall {
	c.opt_["maxResults"] = maxResults
	return c
}

// PageToken sets the optional parameter "pageToken": A
// previously-returned page token representing part of the larger set of
// results to view.
func (c *InstancesListCall) PageToken(pageToken string) *InstancesListCall {
	c.opt_["pageToken"] = pageToken
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *InstancesListCall) Fields(s ...googleapi.Field) *InstancesListCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *InstancesListCall) Do() (*InstancesListResponse, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["maxResults"]; ok {
		params.Set("maxResults", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["pageToken"]; ok {
		params.Set("pageToken", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "projects/{project}/instances")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"project": c.project,
	})
	req.Header.Set("User-Agent", c.s.userAgent())
	res, err := c.s.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	var ret *InstancesListResponse
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Lists instances under a given project in the alphabetical order of the instance name.",
	//   "httpMethod": "GET",
	//   "id": "sql.instances.list",
	//   "parameterOrder": [
	//     "project"
	//   ],
	//   "parameters": {
	//     "maxResults": {
	//       "description": "The maximum number of results to return per response.",
	//       "format": "uint32",
	//       "location": "query",
	//       "type": "integer"
	//     },
	//     "pageToken": {
	//       "description": "A previously-returned page token representing part of the larger set of results to view.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "project": {
	//       "description": "Project ID of the project for which to list Cloud SQL instances.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "projects/{project}/instances",
	//   "response": {
	//     "$ref": "InstancesListResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/sqlservice.admin"
	//   ]
	// }

}

// method id "sql.instances.patch":

type InstancesPatchCall struct {
	s                *Service
	project          string
	instance         string
	databaseinstance *DatabaseInstance
	opt_             map[string]interface{}
}

// Patch: Updates settings of a Cloud SQL instance. Caution: This is not
// a partial update, so you must include values for all the settings
// that you want to retain. For partial updates, use patch.. This method
// supports patch semantics.
func (r *InstancesService) Patch(project string, instance string, databaseinstance *DatabaseInstance) *InstancesPatchCall {
	c := &InstancesPatchCall{s: r.s, opt_: make(map[string]interface{})}
	c.project = project
	c.instance = instance
	c.databaseinstance = databaseinstance
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *InstancesPatchCall) Fields(s ...googleapi.Field) *InstancesPatchCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *InstancesPatchCall) Do() (*Operation, error) {
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.databaseinstance)
	if err != nil {
		return nil, err
	}
	ctype := "application/json"
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "projects/{project}/instances/{instance}")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("PATCH", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"project":  c.project,
		"instance": c.instance,
	})
	req.Header.Set("Content-Type", ctype)
	req.Header.Set("User-Agent", c.s.userAgent())
	res, err := c.s.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	var ret *Operation
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Updates settings of a Cloud SQL instance. Caution: This is not a partial update, so you must include values for all the settings that you want to retain. For partial updates, use patch.. This method supports patch semantics.",
	//   "httpMethod": "PATCH",
	//   "id": "sql.instances.patch",
	//   "parameterOrder": [
	//     "project",
	//     "instance"
	//   ],
	//   "parameters": {
	//     "instance": {
	//       "description": "Cloud SQL instance ID. This does not include the project ID.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "project": {
	//       "description": "Project ID of the project that contains the instance.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "projects/{project}/instances/{instance}",
	//   "request": {
	//     "$ref": "DatabaseInstance"
	//   },
	//   "response": {
	//     "$ref": "Operation"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/sqlservice.admin"
	//   ]
	// }

}

// method id "sql.instances.promoteReplica":

type InstancesPromoteReplicaCall struct {
	s        *Service
	project  string
	instance string
	opt_     map[string]interface{}
}

// PromoteReplica: Promotes the read replica instance to be a
// stand-alone Cloud SQL instance.
func (r *InstancesService) PromoteReplica(project string, instance string) *InstancesPromoteReplicaCall {
	c := &InstancesPromoteReplicaCall{s: r.s, opt_: make(map[string]interface{})}
	c.project = project
	c.instance = instance
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *InstancesPromoteReplicaCall) Fields(s ...googleapi.Field) *InstancesPromoteReplicaCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *InstancesPromoteReplicaCall) Do() (*Operation, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "projects/{project}/instances/{instance}/promoteReplica")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"project":  c.project,
		"instance": c.instance,
	})
	req.Header.Set("User-Agent", c.s.userAgent())
	res, err := c.s.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	var ret *Operation
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Promotes the read replica instance to be a stand-alone Cloud SQL instance.",
	//   "httpMethod": "POST",
	//   "id": "sql.instances.promoteReplica",
	//   "parameterOrder": [
	//     "project",
	//     "instance"
	//   ],
	//   "parameters": {
	//     "instance": {
	//       "description": "Cloud SQL read replica instance name.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "project": {
	//       "description": "ID of the project that contains the read replica.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "projects/{project}/instances/{instance}/promoteReplica",
	//   "response": {
	//     "$ref": "Operation"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/sqlservice.admin"
	//   ]
	// }

}

// method id "sql.instances.resetSslConfig":

type InstancesResetSslConfigCall struct {
	s        *Service
	project  string
	instance string
	opt_     map[string]interface{}
}

// ResetSslConfig: Deletes all client certificates and generates a new
// server SSL certificate for the instance. The changes will not take
// effect until the instance is restarted. Existing instances without a
// server certificate will need to call this once to set a server
// certificate.
func (r *InstancesService) ResetSslConfig(project string, instance string) *InstancesResetSslConfigCall {
	c := &InstancesResetSslConfigCall{s: r.s, opt_: make(map[string]interface{})}
	c.project = project
	c.instance = instance
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *InstancesResetSslConfigCall) Fields(s ...googleapi.Field) *InstancesResetSslConfigCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *InstancesResetSslConfigCall) Do() (*Operation, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "projects/{project}/instances/{instance}/resetSslConfig")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"project":  c.project,
		"instance": c.instance,
	})
	req.Header.Set("User-Agent", c.s.userAgent())
	res, err := c.s.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	var ret *Operation
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Deletes all client certificates and generates a new server SSL certificate for the instance. The changes will not take effect until the instance is restarted. Existing instances without a server certificate will need to call this once to set a server certificate.",
	//   "httpMethod": "POST",
	//   "id": "sql.instances.resetSslConfig",
	//   "parameterOrder": [
	//     "project",
	//     "instance"
	//   ],
	//   "parameters": {
	//     "instance": {
	//       "description": "Cloud SQL instance ID. This does not include the project ID.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "project": {
	//       "description": "Project ID of the project that contains the instance.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "projects/{project}/instances/{instance}/resetSslConfig",
	//   "response": {
	//     "$ref": "Operation"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/sqlservice.admin"
	//   ]
	// }

}

// method id "sql.instances.restart":

type InstancesRestartCall struct {
	s        *Service
	project  string
	instance string
	opt_     map[string]interface{}
}

// Restart: Restarts a Cloud SQL instance.
func (r *InstancesService) Restart(project string, instance string) *InstancesRestartCall {
	c := &InstancesRestartCall{s: r.s, opt_: make(map[string]interface{})}
	c.project = project
	c.instance = instance
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *InstancesRestartCall) Fields(s ...googleapi.Field) *InstancesRestartCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *InstancesRestartCall) Do() (*Operation, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "projects/{project}/instances/{instance}/restart")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"project":  c.project,
		"instance": c.instance,
	})
	req.Header.Set("User-Agent", c.s.userAgent())
	res, err := c.s.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	var ret *Operation
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Restarts a Cloud SQL instance.",
	//   "httpMethod": "POST",
	//   "id": "sql.instances.restart",
	//   "parameterOrder": [
	//     "project",
	//     "instance"
	//   ],
	//   "parameters": {
	//     "instance": {
	//       "description": "Cloud SQL instance ID. This does not include the project ID.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "project": {
	//       "description": "Project ID of the project that contains the instance to be restarted.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "projects/{project}/instances/{instance}/restart",
	//   "response": {
	//     "$ref": "Operation"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/sqlservice.admin"
	//   ]
	// }

}

// method id "sql.instances.restoreBackup":

type InstancesRestoreBackupCall struct {
	s                             *Service
	project                       string
	instance                      string
	instancesrestorebackuprequest *InstancesRestoreBackupRequest
	opt_                          map[string]interface{}
}

// RestoreBackup: Restores a backup of a Cloud SQL instance.
func (r *InstancesService) RestoreBackup(project string, instance string, instancesrestorebackuprequest *InstancesRestoreBackupRequest) *InstancesRestoreBackupCall {
	c := &InstancesRestoreBackupCall{s: r.s, opt_: make(map[string]interface{})}
	c.project = project
	c.instance = instance
	c.instancesrestorebackuprequest = instancesrestorebackuprequest
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *InstancesRestoreBackupCall) Fields(s ...googleapi.Field) *InstancesRestoreBackupCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *InstancesRestoreBackupCall) Do() (*Operation, error) {
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.instancesrestorebackuprequest)
	if err != nil {
		return nil, err
	}
	ctype := "application/json"
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "projects/{project}/instances/{instance}/restoreBackup")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"project":  c.project,
		"instance": c.instance,
	})
	req.Header.Set("Content-Type", ctype)
	req.Header.Set("User-Agent", c.s.userAgent())
	res, err := c.s.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	var ret *Operation
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Restores a backup of a Cloud SQL instance.",
	//   "httpMethod": "POST",
	//   "id": "sql.instances.restoreBackup",
	//   "parameterOrder": [
	//     "project",
	//     "instance"
	//   ],
	//   "parameters": {
	//     "instance": {
	//       "description": "Cloud SQL instance ID. This does not include the project ID.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "project": {
	//       "description": "Project ID of the project that contains the instance.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "projects/{project}/instances/{instance}/restoreBackup",
	//   "request": {
	//     "$ref": "InstancesRestoreBackupRequest"
	//   },
	//   "response": {
	//     "$ref": "Operation"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/sqlservice.admin"
	//   ]
	// }

}

// method id "sql.instances.startReplica":

type InstancesStartReplicaCall struct {
	s        *Service
	project  string
	instance string
	opt_     map[string]interface{}
}

// StartReplica: Starts the replication in the read replica instance.
func (r *InstancesService) StartReplica(project string, instance string) *InstancesStartReplicaCall {
	c := &InstancesStartReplicaCall{s: r.s, opt_: make(map[string]interface{})}
	c.project = project
	c.instance = instance
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *InstancesStartReplicaCall) Fields(s ...googleapi.Field) *InstancesStartReplicaCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *InstancesStartReplicaCall) Do() (*Operation, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "projects/{project}/instances/{instance}/startReplica")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"project":  c.project,
		"instance": c.instance,
	})
	req.Header.Set("User-Agent", c.s.userAgent())
	res, err := c.s.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	var ret *Operation
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Starts the replication in the read replica instance.",
	//   "httpMethod": "POST",
	//   "id": "sql.instances.startReplica",
	//   "parameterOrder": [
	//     "project",
	//     "instance"
	//   ],
	//   "parameters": {
	//     "instance": {
	//       "description": "Cloud SQL read replica instance name.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "project": {
	//       "description": "ID of the project that contains the read replica.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "projects/{project}/instances/{instance}/startReplica",
	//   "response": {
	//     "$ref": "Operation"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/sqlservice.admin"
	//   ]
	// }

}

// method id "sql.instances.stopReplica":

type InstancesStopReplicaCall struct {
	s        *Service
	project  string
	instance string
	opt_     map[string]interface{}
}

// StopReplica: Stops the replication in the read replica instance.
func (r *InstancesService) StopReplica(project string, instance string) *InstancesStopReplicaCall {
	c := &InstancesStopReplicaCall{s: r.s, opt_: make(map[string]interface{})}
	c.project = project
	c.instance = instance
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *InstancesStopReplicaCall) Fields(s ...googleapi.Field) *InstancesStopReplicaCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *InstancesStopReplicaCall) Do() (*Operation, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "projects/{project}/instances/{instance}/stopReplica")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"project":  c.project,
		"instance": c.instance,
	})
	req.Header.Set("User-Agent", c.s.userAgent())
	res, err := c.s.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	var ret *Operation
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Stops the replication in the read replica instance.",
	//   "httpMethod": "POST",
	//   "id": "sql.instances.stopReplica",
	//   "parameterOrder": [
	//     "project",
	//     "instance"
	//   ],
	//   "parameters": {
	//     "instance": {
	//       "description": "Cloud SQL read replica instance name.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "project": {
	//       "description": "ID of the project that contains the read replica.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "projects/{project}/instances/{instance}/stopReplica",
	//   "response": {
	//     "$ref": "Operation"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/sqlservice.admin"
	//   ]
	// }

}

// method id "sql.instances.update":

type InstancesUpdateCall struct {
	s                *Service
	project          string
	instance         string
	databaseinstance *DatabaseInstance
	opt_             map[string]interface{}
}

// Update: Updates settings of a Cloud SQL instance. Caution: This is
// not a partial update, so you must include values for all the settings
// that you want to retain. For partial updates, use patch.
func (r *InstancesService) Update(project string, instance string, databaseinstance *DatabaseInstance) *InstancesUpdateCall {
	c := &InstancesUpdateCall{s: r.s, opt_: make(map[string]interface{})}
	c.project = project
	c.instance = instance
	c.databaseinstance = databaseinstance
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *InstancesUpdateCall) Fields(s ...googleapi.Field) *InstancesUpdateCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *InstancesUpdateCall) Do() (*Operation, error) {
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.databaseinstance)
	if err != nil {
		return nil, err
	}
	ctype := "application/json"
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "projects/{project}/instances/{instance}")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("PUT", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"project":  c.project,
		"instance": c.instance,
	})
	req.Header.Set("Content-Type", ctype)
	req.Header.Set("User-Agent", c.s.userAgent())
	res, err := c.s.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	var ret *Operation
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Updates settings of a Cloud SQL instance. Caution: This is not a partial update, so you must include values for all the settings that you want to retain. For partial updates, use patch.",
	//   "etagRequired": true,
	//   "httpMethod": "PUT",
	//   "id": "sql.instances.update",
	//   "parameterOrder": [
	//     "project",
	//     "instance"
	//   ],
	//   "parameters": {
	//     "instance": {
	//       "description": "Cloud SQL instance ID. This does not include the project ID.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "project": {
	//       "description": "Project ID of the project that contains the instance.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "projects/{project}/instances/{instance}",
	//   "request": {
	//     "$ref": "DatabaseInstance"
	//   },
	//   "response": {
	//     "$ref": "Operation"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/sqlservice.admin"
	//   ]
	// }

}

// method id "sql.operations.get":

type OperationsGetCall struct {
	s         *Service
	project   string
	operation string
	opt_      map[string]interface{}
}

// Get: Retrieves an instance operation that has been performed on an
// instance.
func (r *OperationsService) Get(project string, operation string) *OperationsGetCall {
	c := &OperationsGetCall{s: r.s, opt_: make(map[string]interface{})}
	c.project = project
	c.operation = operation
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *OperationsGetCall) Fields(s ...googleapi.Field) *OperationsGetCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *OperationsGetCall) Do() (*Operation, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "projects/{project}/operations/{operation}")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"project":   c.project,
		"operation": c.operation,
	})
	req.Header.Set("User-Agent", c.s.userAgent())
	res, err := c.s.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	var ret *Operation
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Retrieves an instance operation that has been performed on an instance.",
	//   "httpMethod": "GET",
	//   "id": "sql.operations.get",
	//   "parameterOrder": [
	//     "project",
	//     "operation"
	//   ],
	//   "parameters": {
	//     "operation": {
	//       "description": "Instance operation ID.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "project": {
	//       "description": "Project ID of the project that contains the instance.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "projects/{project}/operations/{operation}",
	//   "response": {
	//     "$ref": "Operation"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/sqlservice.admin"
	//   ]
	// }

}

// method id "sql.operations.list":

type OperationsListCall struct {
	s        *Service
	project  string
	instance string
	opt_     map[string]interface{}
}

// List: Lists all instance operations that have been performed on the
// given Cloud SQL instance in the reverse chronological order of the
// start time.
func (r *OperationsService) List(project string, instance string) *OperationsListCall {
	c := &OperationsListCall{s: r.s, opt_: make(map[string]interface{})}
	c.project = project
	c.instance = instance
	return c
}

// MaxResults sets the optional parameter "maxResults": Maximum number
// of operations per response.
func (c *OperationsListCall) MaxResults(maxResults int64) *OperationsListCall {
	c.opt_["maxResults"] = maxResults
	return c
}

// PageToken sets the optional parameter "pageToken": A
// previously-returned page token representing part of the larger set of
// results to view.
func (c *OperationsListCall) PageToken(pageToken string) *OperationsListCall {
	c.opt_["pageToken"] = pageToken
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *OperationsListCall) Fields(s ...googleapi.Field) *OperationsListCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *OperationsListCall) Do() (*OperationsListResponse, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	params.Set("instance", fmt.Sprintf("%v", c.instance))
	if v, ok := c.opt_["maxResults"]; ok {
		params.Set("maxResults", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["pageToken"]; ok {
		params.Set("pageToken", fmt.Sprintf("%v", v))
	}
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "projects/{project}/operations")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"project": c.project,
	})
	req.Header.Set("User-Agent", c.s.userAgent())
	res, err := c.s.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	var ret *OperationsListResponse
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Lists all instance operations that have been performed on the given Cloud SQL instance in the reverse chronological order of the start time.",
	//   "httpMethod": "GET",
	//   "id": "sql.operations.list",
	//   "parameterOrder": [
	//     "project",
	//     "instance"
	//   ],
	//   "parameters": {
	//     "instance": {
	//       "description": "Cloud SQL instance ID. This does not include the project ID.",
	//       "location": "query",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "maxResults": {
	//       "description": "Maximum number of operations per response.",
	//       "format": "uint32",
	//       "location": "query",
	//       "type": "integer"
	//     },
	//     "pageToken": {
	//       "description": "A previously-returned page token representing part of the larger set of results to view.",
	//       "location": "query",
	//       "type": "string"
	//     },
	//     "project": {
	//       "description": "Project ID of the project that contains the instance.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "projects/{project}/operations",
	//   "response": {
	//     "$ref": "OperationsListResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/sqlservice.admin"
	//   ]
	// }

}

// method id "sql.sslCerts.delete":

type SslCertsDeleteCall struct {
	s               *Service
	project         string
	instance        string
	sha1Fingerprint string
	opt_            map[string]interface{}
}

// Delete: Deletes the SSL certificate. The change will not take effect
// until the instance is restarted.
func (r *SslCertsService) Delete(project string, instance string, sha1Fingerprint string) *SslCertsDeleteCall {
	c := &SslCertsDeleteCall{s: r.s, opt_: make(map[string]interface{})}
	c.project = project
	c.instance = instance
	c.sha1Fingerprint = sha1Fingerprint
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *SslCertsDeleteCall) Fields(s ...googleapi.Field) *SslCertsDeleteCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *SslCertsDeleteCall) Do() (*Operation, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "projects/{project}/instances/{instance}/sslCerts/{sha1Fingerprint}")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("DELETE", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"project":         c.project,
		"instance":        c.instance,
		"sha1Fingerprint": c.sha1Fingerprint,
	})
	req.Header.Set("User-Agent", c.s.userAgent())
	res, err := c.s.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	var ret *Operation
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Deletes the SSL certificate. The change will not take effect until the instance is restarted.",
	//   "httpMethod": "DELETE",
	//   "id": "sql.sslCerts.delete",
	//   "parameterOrder": [
	//     "project",
	//     "instance",
	//     "sha1Fingerprint"
	//   ],
	//   "parameters": {
	//     "instance": {
	//       "description": "Cloud SQL instance ID. This does not include the project ID.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "project": {
	//       "description": "Project ID of the project that contains the instance to be deleted.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "sha1Fingerprint": {
	//       "description": "Sha1 FingerPrint.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "projects/{project}/instances/{instance}/sslCerts/{sha1Fingerprint}",
	//   "response": {
	//     "$ref": "Operation"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/sqlservice.admin"
	//   ]
	// }

}

// method id "sql.sslCerts.get":

type SslCertsGetCall struct {
	s               *Service
	project         string
	instance        string
	sha1Fingerprint string
	opt_            map[string]interface{}
}

// Get: Retrieves a particular SSL certificate. Does not include the
// private key (required for usage). The private key must be saved from
// the response to initial creation.
func (r *SslCertsService) Get(project string, instance string, sha1Fingerprint string) *SslCertsGetCall {
	c := &SslCertsGetCall{s: r.s, opt_: make(map[string]interface{})}
	c.project = project
	c.instance = instance
	c.sha1Fingerprint = sha1Fingerprint
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *SslCertsGetCall) Fields(s ...googleapi.Field) *SslCertsGetCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *SslCertsGetCall) Do() (*SslCert, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "projects/{project}/instances/{instance}/sslCerts/{sha1Fingerprint}")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"project":         c.project,
		"instance":        c.instance,
		"sha1Fingerprint": c.sha1Fingerprint,
	})
	req.Header.Set("User-Agent", c.s.userAgent())
	res, err := c.s.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	var ret *SslCert
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Retrieves a particular SSL certificate. Does not include the private key (required for usage). The private key must be saved from the response to initial creation.",
	//   "httpMethod": "GET",
	//   "id": "sql.sslCerts.get",
	//   "parameterOrder": [
	//     "project",
	//     "instance",
	//     "sha1Fingerprint"
	//   ],
	//   "parameters": {
	//     "instance": {
	//       "description": "Cloud SQL instance ID. This does not include the project ID.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "project": {
	//       "description": "Project ID of the project that contains the instance.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "sha1Fingerprint": {
	//       "description": "Sha1 FingerPrint.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "projects/{project}/instances/{instance}/sslCerts/{sha1Fingerprint}",
	//   "response": {
	//     "$ref": "SslCert"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/sqlservice.admin"
	//   ]
	// }

}

// method id "sql.sslCerts.insert":

type SslCertsInsertCall struct {
	s                     *Service
	project               string
	instance              string
	sslcertsinsertrequest *SslCertsInsertRequest
	opt_                  map[string]interface{}
}

// Insert: Creates an SSL certificate and returns it along with the
// private key and server certificate authority. The new certificate
// will not be usable until the instance is restarted.
func (r *SslCertsService) Insert(project string, instance string, sslcertsinsertrequest *SslCertsInsertRequest) *SslCertsInsertCall {
	c := &SslCertsInsertCall{s: r.s, opt_: make(map[string]interface{})}
	c.project = project
	c.instance = instance
	c.sslcertsinsertrequest = sslcertsinsertrequest
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *SslCertsInsertCall) Fields(s ...googleapi.Field) *SslCertsInsertCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *SslCertsInsertCall) Do() (*SslCertsInsertResponse, error) {
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.sslcertsinsertrequest)
	if err != nil {
		return nil, err
	}
	ctype := "application/json"
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "projects/{project}/instances/{instance}/sslCerts")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"project":  c.project,
		"instance": c.instance,
	})
	req.Header.Set("Content-Type", ctype)
	req.Header.Set("User-Agent", c.s.userAgent())
	res, err := c.s.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	var ret *SslCertsInsertResponse
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Creates an SSL certificate and returns it along with the private key and server certificate authority. The new certificate will not be usable until the instance is restarted.",
	//   "httpMethod": "POST",
	//   "id": "sql.sslCerts.insert",
	//   "parameterOrder": [
	//     "project",
	//     "instance"
	//   ],
	//   "parameters": {
	//     "instance": {
	//       "description": "Cloud SQL instance ID. This does not include the project ID.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "project": {
	//       "description": "Project ID of the project to which the newly created Cloud SQL instances should belong.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "projects/{project}/instances/{instance}/sslCerts",
	//   "request": {
	//     "$ref": "SslCertsInsertRequest"
	//   },
	//   "response": {
	//     "$ref": "SslCertsInsertResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/sqlservice.admin"
	//   ]
	// }

}

// method id "sql.sslCerts.list":

type SslCertsListCall struct {
	s        *Service
	project  string
	instance string
	opt_     map[string]interface{}
}

// List: Lists all of the current SSL certificates for the instance.
func (r *SslCertsService) List(project string, instance string) *SslCertsListCall {
	c := &SslCertsListCall{s: r.s, opt_: make(map[string]interface{})}
	c.project = project
	c.instance = instance
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *SslCertsListCall) Fields(s ...googleapi.Field) *SslCertsListCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *SslCertsListCall) Do() (*SslCertsListResponse, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "projects/{project}/instances/{instance}/sslCerts")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"project":  c.project,
		"instance": c.instance,
	})
	req.Header.Set("User-Agent", c.s.userAgent())
	res, err := c.s.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	var ret *SslCertsListResponse
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Lists all of the current SSL certificates for the instance.",
	//   "httpMethod": "GET",
	//   "id": "sql.sslCerts.list",
	//   "parameterOrder": [
	//     "project",
	//     "instance"
	//   ],
	//   "parameters": {
	//     "instance": {
	//       "description": "Cloud SQL instance ID. This does not include the project ID.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "project": {
	//       "description": "Project ID of the project for which to list Cloud SQL instances.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "projects/{project}/instances/{instance}/sslCerts",
	//   "response": {
	//     "$ref": "SslCertsListResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/sqlservice.admin"
	//   ]
	// }

}

// method id "sql.tiers.list":

type TiersListCall struct {
	s       *Service
	project string
	opt_    map[string]interface{}
}

// List: Lists all available service tiers for Google Cloud SQL, for
// example D1, D2. For related information, see Pricing.
func (r *TiersService) List(project string) *TiersListCall {
	c := &TiersListCall{s: r.s, opt_: make(map[string]interface{})}
	c.project = project
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *TiersListCall) Fields(s ...googleapi.Field) *TiersListCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *TiersListCall) Do() (*TiersListResponse, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "projects/{project}/tiers")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"project": c.project,
	})
	req.Header.Set("User-Agent", c.s.userAgent())
	res, err := c.s.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	var ret *TiersListResponse
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Lists all available service tiers for Google Cloud SQL, for example D1, D2. For related information, see Pricing.",
	//   "httpMethod": "GET",
	//   "id": "sql.tiers.list",
	//   "parameterOrder": [
	//     "project"
	//   ],
	//   "parameters": {
	//     "project": {
	//       "description": "Project ID of the project for which to list tiers.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "projects/{project}/tiers",
	//   "response": {
	//     "$ref": "TiersListResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/sqlservice.admin"
	//   ]
	// }

}

// method id "sql.users.delete":

type UsersDeleteCall struct {
	s        *Service
	project  string
	instance string
	host     string
	name     string
	opt_     map[string]interface{}
}

// Delete: Deletes a user from a Cloud SQL instance.
func (r *UsersService) Delete(project string, instance string, host string, name string) *UsersDeleteCall {
	c := &UsersDeleteCall{s: r.s, opt_: make(map[string]interface{})}
	c.project = project
	c.instance = instance
	c.host = host
	c.name = name
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *UsersDeleteCall) Fields(s ...googleapi.Field) *UsersDeleteCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *UsersDeleteCall) Do() (*Operation, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	params.Set("host", fmt.Sprintf("%v", c.host))
	params.Set("name", fmt.Sprintf("%v", c.name))
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "projects/{project}/instances/{instance}/users")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("DELETE", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"project":  c.project,
		"instance": c.instance,
	})
	req.Header.Set("User-Agent", c.s.userAgent())
	res, err := c.s.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	var ret *Operation
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Deletes a user from a Cloud SQL instance.",
	//   "httpMethod": "DELETE",
	//   "id": "sql.users.delete",
	//   "parameterOrder": [
	//     "project",
	//     "instance",
	//     "host",
	//     "name"
	//   ],
	//   "parameters": {
	//     "host": {
	//       "description": "Host of the user in the instance.",
	//       "location": "query",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "instance": {
	//       "description": "Database instance ID. This does not include the project ID.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "name": {
	//       "description": "Name of the user in the instance.",
	//       "location": "query",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "project": {
	//       "description": "Project ID of the project that contains the instance.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "projects/{project}/instances/{instance}/users",
	//   "response": {
	//     "$ref": "Operation"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/sqlservice.admin"
	//   ]
	// }

}

// method id "sql.users.insert":

type UsersInsertCall struct {
	s        *Service
	project  string
	instance string
	user     *User
	opt_     map[string]interface{}
}

// Insert: Creates a new user in a Cloud SQL instance.
func (r *UsersService) Insert(project string, instance string, user *User) *UsersInsertCall {
	c := &UsersInsertCall{s: r.s, opt_: make(map[string]interface{})}
	c.project = project
	c.instance = instance
	c.user = user
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *UsersInsertCall) Fields(s ...googleapi.Field) *UsersInsertCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *UsersInsertCall) Do() (*Operation, error) {
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.user)
	if err != nil {
		return nil, err
	}
	ctype := "application/json"
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "projects/{project}/instances/{instance}/users")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("POST", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"project":  c.project,
		"instance": c.instance,
	})
	req.Header.Set("Content-Type", ctype)
	req.Header.Set("User-Agent", c.s.userAgent())
	res, err := c.s.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	var ret *Operation
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Creates a new user in a Cloud SQL instance.",
	//   "httpMethod": "POST",
	//   "id": "sql.users.insert",
	//   "parameterOrder": [
	//     "project",
	//     "instance"
	//   ],
	//   "parameters": {
	//     "instance": {
	//       "description": "Database instance ID. This does not include the project ID.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "project": {
	//       "description": "Project ID of the project that contains the instance.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "projects/{project}/instances/{instance}/users",
	//   "request": {
	//     "$ref": "User"
	//   },
	//   "response": {
	//     "$ref": "Operation"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/sqlservice.admin"
	//   ]
	// }

}

// method id "sql.users.list":

type UsersListCall struct {
	s        *Service
	project  string
	instance string
	opt_     map[string]interface{}
}

// List: Lists users in the specified Cloud SQL instance.
func (r *UsersService) List(project string, instance string) *UsersListCall {
	c := &UsersListCall{s: r.s, opt_: make(map[string]interface{})}
	c.project = project
	c.instance = instance
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *UsersListCall) Fields(s ...googleapi.Field) *UsersListCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *UsersListCall) Do() (*UsersListResponse, error) {
	var body io.Reader = nil
	params := make(url.Values)
	params.Set("alt", "json")
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "projects/{project}/instances/{instance}/users")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("GET", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"project":  c.project,
		"instance": c.instance,
	})
	req.Header.Set("User-Agent", c.s.userAgent())
	res, err := c.s.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	var ret *UsersListResponse
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Lists users in the specified Cloud SQL instance.",
	//   "httpMethod": "GET",
	//   "id": "sql.users.list",
	//   "parameterOrder": [
	//     "project",
	//     "instance"
	//   ],
	//   "parameters": {
	//     "instance": {
	//       "description": "Database instance ID. This does not include the project ID.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "project": {
	//       "description": "Project ID of the project that contains the instance.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "projects/{project}/instances/{instance}/users",
	//   "response": {
	//     "$ref": "UsersListResponse"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/sqlservice.admin"
	//   ]
	// }

}

// method id "sql.users.update":

type UsersUpdateCall struct {
	s        *Service
	project  string
	instance string
	host     string
	name     string
	user     *User
	opt_     map[string]interface{}
}

// Update: Updates an existing user in a Cloud SQL instance.
func (r *UsersService) Update(project string, instance string, host string, name string, user *User) *UsersUpdateCall {
	c := &UsersUpdateCall{s: r.s, opt_: make(map[string]interface{})}
	c.project = project
	c.instance = instance
	c.host = host
	c.name = name
	c.user = user
	return c
}

// Fields allows partial responses to be retrieved.
// See https://developers.google.com/gdata/docs/2.0/basics#PartialResponse
// for more information.
func (c *UsersUpdateCall) Fields(s ...googleapi.Field) *UsersUpdateCall {
	c.opt_["fields"] = googleapi.CombineFields(s)
	return c
}

func (c *UsersUpdateCall) Do() (*Operation, error) {
	var body io.Reader = nil
	body, err := googleapi.WithoutDataWrapper.JSONReader(c.user)
	if err != nil {
		return nil, err
	}
	ctype := "application/json"
	params := make(url.Values)
	params.Set("alt", "json")
	params.Set("host", fmt.Sprintf("%v", c.host))
	params.Set("name", fmt.Sprintf("%v", c.name))
	if v, ok := c.opt_["fields"]; ok {
		params.Set("fields", fmt.Sprintf("%v", v))
	}
	urls := googleapi.ResolveRelative(c.s.BasePath, "projects/{project}/instances/{instance}/users")
	urls += "?" + params.Encode()
	req, _ := http.NewRequest("PUT", urls, body)
	googleapi.Expand(req.URL, map[string]string{
		"project":  c.project,
		"instance": c.instance,
	})
	req.Header.Set("Content-Type", ctype)
	req.Header.Set("User-Agent", c.s.userAgent())
	res, err := c.s.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer googleapi.CloseBody(res)
	if err := googleapi.CheckResponse(res); err != nil {
		return nil, err
	}
	var ret *Operation
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return ret, nil
	// {
	//   "description": "Updates an existing user in a Cloud SQL instance.",
	//   "httpMethod": "PUT",
	//   "id": "sql.users.update",
	//   "parameterOrder": [
	//     "project",
	//     "instance",
	//     "host",
	//     "name"
	//   ],
	//   "parameters": {
	//     "host": {
	//       "description": "Host of the user in the instance.",
	//       "location": "query",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "instance": {
	//       "description": "Database instance ID. This does not include the project ID.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "name": {
	//       "description": "Name of the user in the instance.",
	//       "location": "query",
	//       "required": true,
	//       "type": "string"
	//     },
	//     "project": {
	//       "description": "Project ID of the project that contains the instance.",
	//       "location": "path",
	//       "required": true,
	//       "type": "string"
	//     }
	//   },
	//   "path": "projects/{project}/instances/{instance}/users",
	//   "request": {
	//     "$ref": "User"
	//   },
	//   "response": {
	//     "$ref": "Operation"
	//   },
	//   "scopes": [
	//     "https://www.googleapis.com/auth/cloud-platform",
	//     "https://www.googleapis.com/auth/sqlservice.admin"
	//   ]
	// }

}