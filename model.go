package esme

type config struct {
	Routes []*route `yaml:"routes" validate:"gte=1,dive"`
}

type route struct {
	Url         string                 `yaml:"url" validate:"required"`
	Method      string                 `yaml:"method" validate:"required"`
	StatusCode  int                    `yaml:"status_code" validate:"required"`
	Body        map[string]string      `yaml:"body"`
	ResponseObj map[string]interface{} `yaml:"response_obj"`
	ResponseArr []map[string]string    `yaml:"response_arr"`
	ResponseStr string                 `yaml:"response_str"`
	Auth        *auth                  `yaml:"auth"`
}

type auth struct {
	Basic       *authBasic        `yaml:"basic"`
	BearerToken *authBearerToken  `yaml:"bearer_token"`
	Custom      map[string]string `yaml:"custom"`
}

type authBasic struct {
	Username string `yaml:"username" validate:"required"`
	Password string `yaml:"password" validate:"required"`
}

type authBearerToken struct {
	Token string `yaml:"token" validate:"required"`
}
