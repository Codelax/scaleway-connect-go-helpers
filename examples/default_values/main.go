package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"buf.build/gen/go/scaleway/scaleway-apis/connectrpc/go/scaleway/iam/v1alpha1/iamv1alpha1connect"
	iamv1alpha1 "buf.build/gen/go/scaleway/scaleway-apis/protocolbuffers/go/scaleway/iam/v1alpha1"
	"connectrpc.com/connect"
	"github.com/Codelax/scaleway-connect-go-helpers/connectscw"
	"github.com/scaleway/scaleway-sdk-go/scw"
)

func main() {
	ctx := context.Background()

	cfg, err := scw.LoadConfig()
	if err != nil {
		log.Fatalln(fmt.Errorf("failed to load scw config: %w", err))
	}

	client := iamv1alpha1connect.NewApiClient(&http.Client{},
		"https://api.scaleway.com",
		connect.WithGRPC(),
		connect.WithInterceptors(
			connectscw.NewAuthInterceptorFromProfile(&cfg.Profile),
		),
	)

	resp, err := client.CreateApplication(ctx, &connect.Request[iamv1alpha1.CreateApplicationRequest]{
		Msg: &iamv1alpha1.CreateApplicationRequest{
			Name: "scaleway-connect-go-helpers-application",
		},
	})
	if err != nil {
		log.Fatalln(connectscw.SdkError(err))
	}

	log.Println(resp)
}
