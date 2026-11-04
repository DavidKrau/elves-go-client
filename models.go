package elves

type Rule struct {
	ID         int64  `json:"id,omitempty"`
	RuleType   string `json:"rule_type" validate:"required,oneof=BINARY CERTIFICATE TEAMID SIGNINGID CDHASH"`
	Policy     string `json:"policy" validate:"required,oneof=ALLOWLIST_COMPILER ALLOWLIST BLOCKLIST SILENT_BLOCKLIST CEL"`
	Identifier string `json:"identifier" validate:"required"`
	CelExpr    string `json:"cel_expr,omitempty" validate:"required_if=Policy CEL"`
	CustomMsg  string `json:"custom_msg" validate:"required"`
	IsDefault  bool   `json:"is_default"`

	// Shown to the user on the Elves rules page when this rule blocks something
	Description             string `json:"description"`
	RemediationInstructions string `json:"remediation_instructions"`
	// Elves only lets a logged-in admin change the script, so API key clients
	// must send back the value they read
	RemediationScript string `json:"remediation_script"`
}

// AssignedRule is a rule as assigned to a SCIM group or a client, with
// whether it is included or excluded there.
type AssignedRule struct {
	ID         int64  `json:"id"`
	RuleType   string `json:"rule_type"`
	Policy     string `json:"policy"`
	Identifier string `json:"identifier"`
	CustomMsg  string `json:"custom_msg"`
	IsDefault  bool   `json:"is_default"`
	Action     string `json:"action"`
}

// ScimGroup is a group provisioned into Elves by SCIM (for example from Okta).
type ScimGroup struct {
	ID          int64  `json:"id"`
	DisplayName string `json:"display_name"`
	MemberCount int    `json:"member_count"`
}

// ScimUser is a user provisioned into Elves by SCIM (for example from Okta).
// Rules for a user apply to Macs whose primary user equals its UserName.
type ScimUser struct {
	ID       int64  `json:"id"`
	UserName string `json:"user_name"`
	Email    string `json:"email,omitempty"`
	Active   bool   `json:"active"`
}

// RuleScimUser is a SCIM user a rule is assigned to, with the action.
type RuleScimUser struct {
	UserID   int64  `json:"user_id"`
	UserName string `json:"user_name"`
	Email    string `json:"email,omitempty"`
	Active   bool   `json:"active"`
	Action   string `json:"action"`
}

// RuleConfig is a configuration a rule is assigned to, with the action.
type RuleConfig struct {
	ConfigID int64  `json:"config_id"`
	Name     string `json:"name"`
	Action   string `json:"action"`
}

// Configuration is a Santa configuration that Elves sends to the Macs using it.
type Configuration struct {
	ID                       int64  `json:"id,omitempty"`
	Name                     string `json:"name"`
	ClientMode               string `json:"client_mode"`
	BlockedPathRegex         string `json:"blocked_path_regex"`
	AllowedPathRegex         string `json:"allowed_path_regex"`
	BatchSize                int64  `json:"batch_size"`
	EnableBundles            bool   `json:"enable_bundles"`
	EnableTransitiveRules    bool   `json:"enable_transitive_rules"`
	FullSyncInterval         int64  `json:"full_sync_interval"`
	BlockUSBMount            bool   `json:"block_usb_mount"`
	RemountUSBMode           string `json:"remount_usb_mode"`
	SyncType                 string `json:"sync_type"`
	OverrideFileAccessAction string `json:"override_file_access_action"`
}
