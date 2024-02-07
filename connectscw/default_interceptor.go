package connectscw

import (
	"context"

	"connectrpc.com/connect"
	"github.com/scaleway/scaleway-sdk-go/scw"
)

type DefaultValuesInterceptor struct {
}

func (d *DefaultValuesInterceptor) WrapUnary(unaryFunc connect.UnaryFunc) connect.UnaryFunc {
	return func(ctx context.Context, request connect.AnyRequest) (connect.AnyResponse, error) {

		return unaryFunc(ctx, request)
	}
}

func (d *DefaultValuesInterceptor) WrapStreamingClient(clientFunc connect.StreamingClientFunc) connect.StreamingClientFunc {
	return clientFunc
}

func (d *DefaultValuesInterceptor) WrapStreamingHandler(handlerFunc connect.StreamingHandlerFunc) connect.StreamingHandlerFunc {
	return handlerFunc
}

func NewDefaultValuesInterceptor(profile *scw.Profile) connect.Interceptor {
	return &DefaultValuesInterceptor{}
}
