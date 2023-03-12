package types

type Config struct {
	Device Device
	Server Server
}

type Device struct {
	ClientID string
	Platform string
}

type Server struct {
	BaseUrl      string
	WebSocketUrl string
}
