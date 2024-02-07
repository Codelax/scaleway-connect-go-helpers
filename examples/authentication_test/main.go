package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"buf.build/gen/go/scaleway/scaleway-apis/connectrpc/go/scaleway/test/v1/testv1connect"
	testv1 "buf.build/gen/go/scaleway/scaleway-apis/protocolbuffers/go/scaleway/test/v1"
	"connectrpc.com/connect"
	"github.com/Codelax/scaleway-connect-go-helpers/connectscw"
)

func main() {
	ctx := context.Background()

	// Create a tmpClient that will be used to get a secret-key for the test API
	tmpClient := testv1connect.NewApiClient(&http.Client{}, "https://api.scaleway.com", connect.WithGRPC())

	registerResp, err := tmpClient.Register(ctx, &connect.Request[testv1.RegisterRequest]{
		Msg: &testv1.RegisterRequest{
			Username: "connect-go-helpers-authentication_test-example",
		},
	})
	if err != nil {
		log.Fatalln(fmt.Errorf("failed to register: %w\n", err))
	}

	client := testv1connect.NewApiClient(&http.Client{}, "https://api.scaleway.com", connect.WithGRPC(),
		connect.WithInterceptors(connectscw.NewAuthInterceptor(registerResp.Msg.SecretKey)),
	)

	resp, err := client.ListHumans(ctx, &connect.Request[testv1.ListHumansRequest]{})
	if err != nil {
		log.Fatalln(fmt.Errorf("failed to request with authentication_test: %w", err))
	}

	log.Println(resp)
}
