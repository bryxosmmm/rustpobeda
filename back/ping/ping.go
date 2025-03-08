package ping

import "context"

type PingResponse struct {
	Message string
}

// encore:api public method=GET path=/api/v1/ping
func Ping(ctx context.Context) (*PingResponse, error) {
	return &PingResponse{Message: "Hello world!"}, nil
}
