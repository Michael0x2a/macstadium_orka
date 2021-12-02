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
	vmName := flag.String(
		"vm-name",
		"",
		"The name of the VM to create",
	)
	baseImage := flag.String(
		"base-image",
		"90GCatalinaSSH.img",
		"Which base image to use. Defaults to '90GCatalinaSSH.img'",
	)
	cpuCore := flag.Int(
		"cpu-core",
		6,
		"How many CPUs to assign to your VM. Defaults to 6",
	)
	vcpuCount := flag.Int(
		"vcpu-count",
		6,
		"Must be half or exactly the same as cpu-core. Defaults to 6",
	)
	flag.Parse()

	config, err := utils.ParseAPIConfig(*configPath)
	if err != nil {
		panic(err.Error())
	}

	err = orkaVMCreate(context.Background(), config, *vmName, *baseImage, float32(*cpuCore), float32(*vcpuCount))
	if err != nil {
		panic(err.Error())
	}
}

func orkaVMCreate(
	ctx context.Context,
	config utils.APIConfig,
	vmName string,
	baseImage string,
	cpuCore float32,
	vcpuCount float32,
) error {
	auth := orka.Authentication{
		BearerToken: config.BearerToken,
	}

	client, err := orka.NewOrkaClient(config.ServerURL, auth)
	if err != nil {
		return fmt.Errorf("could not construct standard client: %w", err)
	}

	// The terminology is a bit confused here -- the `orka vm create` command actually
	// creates a _VM configuration_ under the hood, instead of a VM. See:
	// https://orkadocs.macstadium.com/docs/orka-glossary#vm-configuration
	//
	// The 'orka vm deploy' command is responsible for actually instantiating this config
	// and creating a live VM.
	resp, err := client.CreateVmConfigurationWithResponse(ctx, orka.CreateVmConfigurationJSONRequestBody{
		OrkaVmName:    &vmName,
		OrkaImage:     &vmName,
		OrkaBaseImage: &baseImage,
		OrkaCpuCore:   &cpuCore,
		VcpuCount:     &vcpuCount,
		AttachedDisk:  nil,
		IsoImage:      nil,
		SystemSerial:  nil,
		VncConsole:    nil,
	})
	if err != nil {
		return fmt.Errorf("api call to CreateVmConfiguration failed: %w", err)
	}
	if resp.JSON201 == nil {
		return fmt.Errorf("api call to CreateVmConfiguration failed: %w", orka.ExtractError(resp.Body))
	}

	fmt.Println("Created VM configuration:", *resp.JSON201.Message)

	return nil
}
