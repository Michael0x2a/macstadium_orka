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
        "The name of the VM to delete",
    )
    flag.Parse()

    config, err := utils.ParseAPIConfig(*configPath)
    if err != nil {
        panic(err.Error())
    }

    err = orkaVMPurge(context.Background(), config, *vmName)
    if err != nil {
        panic(err.Error())
    }
}

func orkaVMPurge(ctx context.Context, config utils.APIConfig, vmName string) error {
    auth := orka.Authentication{
        BearerToken: config.BearerToken,
    }

    client, err := orka.NewOrkaClient(config.ServerURL, auth)
    if err != nil {
        return fmt.Errorf("could not construct standard client: %w", err)
    }

    // Despite its name, this method actually supports both administrative and non-administrative
    // requests.
    //
    // This name is garbled due to a quirk with the Postman collection -> OpenAPI v3.0 conversion
    // process. Basically, the Postman collection contains two examples for the /resources/vm/purge endpoint:
    //
    // - Purge user's VMs: https://documenter.getpostman.com/view/6574930/S1ETRGzt?version=latest#7b6b9ae5-aa89-4ac3-9f50-f19becf54529
    // - Purge user's VMs (Admin): https://documenter.getpostman.com/view/6574930/S1ETRGzt?version=latest#bf95ceb8-5903-48a7-b949-80e0175b51e6
    //
    // The converter ends up picking just the latter (and capitalizing that apostrophe'd "s").
    resp, err := client.PurgeUserSVMsAdminWithResponse(ctx, orka.PurgeUserSVMsAdminJSONRequestBody{
        OrkaVmName: &vmName,
    })
    if err != nil {
        return fmt.Errorf("api call to PurgeUserSVMsAdmin failed: %w", err)
    }
    if resp.JSON200 == nil {
        return fmt.Errorf("api call to PurgeUserSVMsAdmin failed: %w", orka.ExtractError(resp.Body))
    }

    fmt.Println("Deleted VM:", *resp.JSON200.Message)

    return nil
}
