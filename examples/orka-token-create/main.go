package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/Michael0x2a/macstadium_orka/examples/utils"
	"github.com/Michael0x2a/macstadium_orka/pkg/orka"
)

func main() {
	configPath := flag.String(
		"user-config-path",
		"",
		"Path to JSON config containing user credentials and information needed to make administrative requests",
	)
	flag.Parse()

	config, err := utils.ParseUserConfig(*configPath)
	if err != nil {
		panic(err.Error())
	}

	err = orkaTokenCreate(context.Background(), config)
	if err != nil {
		panic(err.Error())
	}
}

func orkaTokenCreate(ctx context.Context, config utils.UserConfig) error {
	// The license key isn't required, interestingly -- just the user credentials.
	auth := orka.Authentication{}

	client, err := orka.NewOrkaClient(config.ServerURL, auth)
	if err != nil {
		return fmt.Errorf("could not construct client: %w", err)
	}
	
	resp, err := client.CreateTokenWithResponse(ctx, orka.CreateTokenJSONRequestBody{
		Email: &config.Email,
		Password: &config.Password,
	})
	if err != nil {
		return fmt.Errorf("api call to CreateToken failed: %w", err)
	}
	if resp.JSON200 == nil {
		return fmt.Errorf("api call to CreateToken failed: %w", orka.ExtractError(resp.Body))
	}

	fmt.Println("Created token:", *resp.JSON200.Token)

	return nil
}
