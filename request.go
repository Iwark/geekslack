package geekslack

import (
	"encoding/json"
	"net/url"
)

// Request is the struct
type Request struct {
	Token       String `json:"token"`
	TeamID      String `json:"team_id"`
	TeamDomain  String `json:"team_domain"`
	ChannelID   String `json:"channel_id"`
	ChannelName String `json:"channel_name"`
	Timestamp   Number `json:"timestamp"`
	UserID      String `json:"user_id"`
	UserName    String `json:"user_name"`
	Text        String `json:"text"`
	TriggerWord String `json:"trigger_word"`
}

// String is an alias for type string
type String string

// UnmarshalJSON url-decodes string
func (s *String) UnmarshalJSON(b []byte) error {
	var str string
	err := json.Unmarshal(b, &str)
	if err != nil {
		return err
	}
	unescaped, err := url.QueryUnescape(str)
	if err != nil {
		return err
	}
	*s = String(unescaped)
	return nil
}

// Number is an alias for type float64
type Number float64

// UnmarshalJSON converts json.Number to Number
func (m *Number) UnmarshalJSON(b []byte) error {
	var number json.Number
	if err := json.Unmarshal(b, &number); err != nil {
		return err
	}
	i, err := number.Float64()
	if err != nil {
		return err
	}
	*m = Number(i)
	return nil
}
