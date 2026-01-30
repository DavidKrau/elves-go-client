package elves

type Rule struct {
	ID         int    `json:"id,omitempty"`
	RuleType   string `json:"rule_type"`
	Policy     string `json:"policy"`
	Identifier string `json:"identifier"`
	CustomMsg  string `json:"custom_msg"`
	IsDefault  bool   `json:"is_default"`
}
