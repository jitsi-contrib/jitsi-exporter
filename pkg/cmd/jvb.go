package cmd

import (
	"github.com/jitsi-contrib/jitsi-exporter/pkg/jvb"
	"github.com/spf13/cobra"
)

var jvbUrl string

var jvbCmd = &cobra.Command{
	Use:   "jvb",
	Short: "Export JVB metrics",
	Run: func(cmd *cobra.Command, args []string) {
		prober = jvb.NewProber(httpClient, jvbUrl)
	},
}

func init() {
	rootCmd.AddCommand(jvbCmd)
	jvbCmd.Flags().StringVar(&jvbUrl, "url", "http://127.0.0.1:8080/colibri/stats", "URL of the JVB stats endpoint")
}
