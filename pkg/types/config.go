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

type Ignore struct {
	Regexp   []string
	Filepath []string
}

//type Regexp struct {
//}
//
//type Filepath struct {
//}
