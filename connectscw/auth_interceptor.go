package connectscw

import (
	"context"

	"connectrpc.com/connect"

	scwsdkgo "github.com/scaleway/scaleway-sdk-go/scw"
)

type AuthInterceptor struct {
	secretKey string
}

func (a *AuthInterceptor) WrapUnary(unaryFunc connect.UnaryFunc) connect.UnaryFunc {
	return func(ctx context.Context, request connect.AnyRequest) (connect.AnyResponse, error) {
		request.Header().Set("X-Auth-Token", a.secretKey)

		return unaryFunc(ctx, request)
	}
}

func (a *AuthInterceptor) WrapStreamingClient(clientFunc connect.StreamingClientFunc) connect.StreamingClientFunc {
	return clientFunc
}

func (a *AuthInterceptor) WrapStreamingHandler(handlerFunc connect.StreamingHandlerFunc) connect.StreamingHandlerFunc {
	return handlerFunc
}

func NewAuthInterceptor(secretKey string) connect.Interceptor {
	return &AuthInterceptor{
		secretKey: secretKey,
	}
}

func NewAuthInterceptorFromProfile(profile *scwsdkgo.Profile) connect.Interceptor {
	secretKey := ""
	if profile.SecretKey != nil {
		secretKey = *profile.SecretKey
	}

	return &AuthInterceptor{
		secretKey: secretKey,
	}
}
