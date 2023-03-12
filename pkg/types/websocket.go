package types

type PubSubMessage struct {
	Type     string `json:"type"`
	Action   string `json:"action"`
	ClientID string `json:"client_id"`
	Data     []byte `json:"data"`
}
