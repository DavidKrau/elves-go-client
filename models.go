package elves

type Rule struct {
	ID         int64  `json:"id,omitempty"`
	RuleType   string `json:"rule_type" validate:"required,oneof=BINARY CERTIFICATE TEAMID SIGNINGID CDHASH"`
	Policy     string `json:"policy" validate:"required,oneof=ALLOWLIST_COMPILER ALLOWLIST BLOCKLIST SILENT_BLOCKLIST CEL"`
	Identifier string `json:"identifier" validate:"required"`
	CelExpr    string `json:"cel_expr,omitempty" validate:"required_if=Policy CEL"`
	CustomMsg  string `json:"custom_msg" validate:"required"`
	IsDefault  bool   `json:"is_default"`
}
