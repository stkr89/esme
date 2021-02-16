package esme

type config struct {
	Routes []*route `json:"routes" validate:"gte=1,dive"`
}

type route struct {
	Url        string      `json:"url" validate:"required"`
	Method     string      `json:"method" validate:"required"`
	StatusCode int         `json:"status_code" validate:"required"`
	Body       interface{} `json:"body"`
	Response   interface{} `json:"response"`
	Auth       *auth       `json:"auth"`
}

type auth struct {
	Basic       *authBasic        `json:"basic"`
	BearerToken *authBearerToken  `json:"bearer_token"`
	Custom      map[string]string `json:"custom"`
}

type authBasic struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type authBearerToken struct {
	Token string `json:"token" validate:"required"`
}
