package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/Michael0x2a/macstadium_orka/examples/utils"
	"github.com/Michael0x2a/macstadium_orka/pkg/orka"
	"net/http"
)

func main() {
	configPath := flag.String(
		"api-config-path",
		"",
		"Path to JSON config describing how to interact with the Macstadium Orka API",
	)
	flag.Parse()

	config, err := utils.ParseAPIConfig(*configPath)
	if err != nil {
		panic(err.Error())
	}

	err = orkaTokenRevoke(context.Background(), config)
	if err != nil {
		panic(err.Error())
	}
}

func orkaTokenRevoke(ctx context.Context, config utils.APIConfig) error {
	auth := orka.Authentication{
		BearerToken: config.BearerToken,
	}

	client, err := orka.NewOrkaClient(config.ServerURL, auth)
	if err != nil {
		return fmt.Errorf("could not construct client: %w", err)
	}
	
	resp, err := client.RevokeTokenWithResponse(ctx, func(ctx context.Context, req *http.Request) error {
		req.Header.Set("Token", config.BearerToken)
		return nil
	})
	if err != nil {
		return fmt.Errorf("api call to RevokeToken failed: %w", err)
	}
	if resp.JSON200 == nil {
		return fmt.Errorf("api call to RevokeToken failed: %w", orka.ExtractError(resp.Body))
	}

	fmt.Println("Num tokens revoked:", *resp.JSON200.TokensRevoked)

	return nil
}
