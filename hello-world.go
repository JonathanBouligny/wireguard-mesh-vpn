package main

import (
	"fmt"
	"log"
	"os/exec"

	"golang.zx2c4.com/wireguard/wgctrl"
)

func main() {

	// Create interfaces
	// Assign ip addresses
	// Create routes
	// Name interfaces
	// Authorization
	// Authentication
	// Relays
	// Egress Node
	// Message Network
	// Lighthouse? - what relays to use or what nodes you can holepunch through

	fmt.Print("Hello welcome to wireguard\n")
	createCmd := exec.Command("ip", "link", "add", "dev", "wg0", "type", "wireguard")

	// Run the command
	err := createCmd.Run()
	if err != nil {
		log.Fatalf("Command execution failed: %s", err)
	}

	// Create a new client
	client, err := wgctrl.New()

	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	// List all WireGuard devices
	devices, err := client.Devices()
	if err != nil {
		log.Fatalf("Failed to list devices: %v", err)
	}

	for _, device := range devices {
		fmt.Printf("Device name: %s\n", device.Name)
	}

	// FEELS BAD MAN
	client.Close()

	// FELT
	delCommand := exec.Command("ip", "link", "del", "dev", "wg0")

	err2 := delCommand.Run()
	if err2 != nil {
		log.Fatalf("Command execution failed: %s", err)
	}
}
