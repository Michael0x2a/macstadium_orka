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
		"api-config-path",
		"",
		"Path to JSON config describing how to interact with the Macstadium Orka API",
	)
	flag.Parse()

	config, err := utils.ParseAPIConfig(*configPath)
	if err != nil {
		panic(err.Error())
	}

	err = orkaImageList(context.Background(), config)
	if err != nil {
		panic(err.Error())
	}
}

func orkaImageList(ctx context.Context, config utils.APIConfig) error {
	auth := orka.Authentication{
		BearerToken: config.BearerToken,
	}

	client, err := orka.NewOrkaClient(config.ServerURL, auth)
	if err != nil {
		return fmt.Errorf("could not construct standard client: %w", err)
	}

	resp, err := client.ListImagesWithResponse(ctx)
	if err != nil {
		return fmt.Errorf("api call to ListImages failed: %w", err)
	}
	if resp.JSON200 == nil {
		return fmt.Errorf("api call to ListImages failed: %w", orka.ExtractError(resp.Body))
	}

	for _, image := range *resp.JSON200.ImageAttributes {
		fmt.Println(*image.Image)
		fmt.Println("    Size:      ", *image.ImageSize)
		fmt.Println("    Modified:  ", *image.Modified)
		fmt.Println("    Date added:", *image.DateAdded)
		fmt.Println("    Owner:     ", *image.Owner)
		fmt.Println("")
	}

	return nil
}
