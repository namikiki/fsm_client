package types

type Config struct {
	Device Device
}

type Device struct {
	ClientID string
	Platform string
}
