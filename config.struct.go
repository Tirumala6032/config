package configuration

type Config struct {
	ServerPort       int              `json:"serverPort"`
	DBHost           string           `json:"dbHost"`
	DBPort           int              `json:"dbPort"`
	DBName           string           `json:"dbName"`
	DBDriver         string           `json:"dbDriver"`
	DBUsername       string           `json:"dbUsername"`
	DBPassword       string           `json:"dbPassword"`
	AuthService      AuthService      `json:"authService"`
	WebSocketService WebSocketService `json:"webSocketService"`
	UserService      UserService      `json:"userService"`
}

type AuthService struct {
	ServerHost string `json:"serverHost"`
	ServerPort int    `json:"serverPort"`
}

type WebSocketService struct {
	ServerHost string `json:"serverHost"`
	ServerPort int    `json:"serverPort"`
}

type UserService struct {
	ServerHost string `json:"serverHost"`
	ServerPort int    `json:"serverPort"`
}
