package esme

type config struct {
	RouteGroups []*group `json:"route_groups" validate:"gte=1,dive"`
}

type group struct {
	Auth      *auth    `json:"auth"`
	Endpoints []*route `json:"endpoints" validate:"gte=1,dive"`
}

type route struct {
	Url        string            `json:"url" validate:"required"`
	Method     string            `json:"method" validate:"required"`
	StatusCode int               `json:"status_code" validate:"required"`
	Body       map[string]string `json:"body"`
	Response   interface{}       `json:"response"`
	Auth       *auth             `json:"-"`
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
