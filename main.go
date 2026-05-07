package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func main() {
	var serverURL string
	var token string

	root := &cobra.Command{
		Use:   "ha",
		Short: "HostAfrica VPS CLI",
		Long: `Command-line interface for the HostAfrica VPS API.

Connects to a running ha-proxy-api server. Configure the target with:
  --server   (or HA_SERVER env var)   e.g. http://localhost:9100
  --token    (or HA_TOKEN env var)    Bearer token for authentication`,
	}

	root.PersistentFlags().StringVar(&serverURL, "server", envOr("HA_SERVER", "https://api.hostafrica.com"), "API server base URL")
	root.PersistentFlags().StringVar(&token, "token", envOr("HA_TOKEN", ""), "Bearer token for authentication")

	client := NewClient(serverURL, token)

	root.PersistentPreRunE = func(cmd *cobra.Command, args []string) error {
		if serverURL == "" {
			return fmt.Errorf("server URL is required (--server or HA_SERVER)")
		}
		client.BaseURL = serverURL
		client.Token = token
		return nil
	}

	RegisterGeneratedCommands(root, client)

	if err := root.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func envOr(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}
