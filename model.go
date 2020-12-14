package main

type config struct {
	Routes []*routes `yaml:"routes"`
}

type routes struct {
	Url        string              `yaml:"url"`
	Method     string              `yaml:"method"`
	StatusCode int                 `yaml:"status_code"`
	Response   []map[string]string `yaml:"response"`
	Auth       *auth               `yaml:"auth"`
}

type auth struct {
	Basic       *authBasic        `yaml:"basic"`
	BearerToken *authBearerToken  `yaml:"bearer_token"`
	Custom      map[string]string `yaml:"custom"`
}

type authBasic struct {
	Username string `yaml:"username"`
	Password string `yaml:"password"`
}

type authBearerToken struct {
	Token string `yaml:"token" validate:"required"`
}
