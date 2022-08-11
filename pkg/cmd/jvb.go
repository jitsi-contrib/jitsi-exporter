package cmd

import (
	"github.com/jitsi-contrib/jitsi-exporter/pkg/jvb"
	"github.com/spf13/cobra"
)

var jvbCmd = &cobra.Command{
	Use:   "jvb",
	Short: "Export JVB metrics",
	Run: func(cmd *cobra.Command, args []string) {
		prober = jvb.NewProber(httpClient, targetUrl)
	},
}

func init() {
	rootCmd.AddCommand(jvbCmd)
	jvbCmd.Flags().StringVar(&targetUrl, "url", "http://127.0.0.1:8080/colibri/stats", "URL of the JVB stats endpoint")
}
