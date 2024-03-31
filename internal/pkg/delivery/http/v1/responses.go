package http

type responseError struct {
	Status  string `json:"status" example:"error"`
	Code    string `json:"code" example:"auth"`
	Message string `json:"message" example:"the authorization token has not been transferred"`
}

type responseOk struct {
	Status string `json:"status" example:"ok"`
	Body   any    `json:"body"`
}
