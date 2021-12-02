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
        "The name of the VM to deploy",
    )
    enableVNC := flag.Bool(
        "enable-vnc",
        true,
        "Whether we should enable the VNC console. Defaults to true",
    )
    flag.Parse()

    config, err := utils.ParseAPIConfig(*configPath)
    if err != nil {
        panic(err.Error())
    }

    err = orkaVMDeploy(context.Background(), config, *vmName, *enableVNC)
    if err != nil {
        panic(err.Error())
    }
}

func orkaVMDeploy(ctx context.Context, config utils.APIConfig, vmName string, enableVNC bool) error {
    auth := orka.Authentication{
        BearerToken: config.BearerToken,
    }

    client, err := orka.NewOrkaClient(config.ServerURL, auth)
    if err != nil {
        return fmt.Errorf("could not construct standard client: %w", err)
    }

    resp, err := client.DeployVmConfigurationWithResponse(ctx, orka.DeployVmConfigurationJSONRequestBody{
        OrkaVmName:     &vmName,
        VncConsole:     &enableVNC,
    })
    if err != nil {
        return fmt.Errorf("api call to DeployVmConfiguration failed: %w", err)
    }
    if resp.JSON200 == nil {
        return fmt.Errorf("api call to DeployVmConfiguration failed: %w", orka.ExtractError(resp.Body))
    }

    fmt.Println("Created VM:")
    fmt.Println("    IP:               ", *resp.JSON200.Ip)
    fmt.Println("    SSH Port:         ", *resp.JSON200.SshPort)
    fmt.Println("    VNC Port:         ", resp.JSON200.VncPort)
    fmt.Println("    Screen Share Port:", resp.JSON200.ScreenSharePort)
    fmt.Println("    VM ID:            ", *resp.JSON200.VmId)
    fmt.Println("    VCPU:             ", *resp.JSON200.Vcpu)
    fmt.Println("    RAM:              ", *resp.JSON200.Ram)

    return nil
}
