package main

import (
    "context"
    "flag"
    "fmt"
    "github.com/Michael0x2a/macstadium_orka/examples/utils"
    "github.com/Michael0x2a/macstadium_orka/pkg/orka"
    "strings"
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

    err = orkaVMList(context.Background(), config)
    if err != nil {
        panic(err.Error())
    }
}

func orkaVMList(ctx context.Context, config utils.APIConfig) error {
    auth := orka.Authentication{
        BearerToken: config.BearerToken,
    }

    client, err := orka.NewOrkaClient(config.ServerURL, auth)
    if err != nil {
        return fmt.Errorf("could not construct standard client: %w", err)
    }

    vmResp, err := client.ListYourVMsWithResponse(ctx)
    if err != nil {
        return fmt.Errorf("api call to ListYourVMs failed: %w", err)
    }
    if vmResp.JSON200 == nil {
        return fmt.Errorf("api call to ListYourVMs failed: %w", orka.ExtractError(vmResp.Body))
    }

    fmt.Println("Listing VMs:")
    fmt.Println()
    for _, vm := range *vmResp.JSON200.VirtualMachineResources {
        fmt.Println(strPtr(vm.VirtualMachineName))
        fmt.Println("    Image:     ", strPtr(vm.Image))
        fmt.Println("    Base image:", strPtr(vm.BaseImage))
        fmt.Println("    Cpu:       ", floatPtr(vm.Cpu))
        fmt.Println("    Vcpu:      ", floatPtr(vm.Vcpu))
        fmt.Println("    Status:    ", strPtr(vm.VmDeploymentStatus))
        fmt.Println()
    }

    configResp, err := client.ListYourVmConfigurationsWithBodyWithResponse(ctx, "application/json", strings.NewReader(""))
    if err != nil {
        return fmt.Errorf("api call to ListYourVMConfigurations failed: %w", err)
    }
    if configResp.JSON200 == nil {
        return fmt.Errorf("api call to ListYourVMConfigurations failed: %w", orka.ExtractError(configResp.Body))
    }

    fmt.Println()
    fmt.Println("----")
    fmt.Println()
    fmt.Println("Listing VM configurations:")
    fmt.Println()
    for _, vm := range *configResp.JSON200.Configs {
        fmt.Println(strPtr(vm.OrkaVmName))
        fmt.Println("    Base image:", strPtr(vm.OrkaBaseImage))
        fmt.Println("    Iso image: ", strPtr(vm.IsoImage))
        fmt.Println("    Cpu:       ", floatPtr(vm.OrkaCpuCore))
        fmt.Println("    Vcpu:      ", floatPtr(vm.VcpuCount))
        fmt.Println()
    }

    return nil
}

func strPtr(s *string) string {
    if s == nil {
        return "nil"
    } else {
        return *s
    }
}

func floatPtr(f *float32) string {
    if f == nil {
        return "nil"
    } else {
        return fmt.Sprintf("%v", *f)
    }
}