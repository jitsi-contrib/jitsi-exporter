package cmd

import (
	"github.com/jitsi-contrib/jitsi-exporter/pkg/jicofo"
	"github.com/spf13/cobra"
)

var jicofoUrl string

var jicofoCmd = &cobra.Command{
	Use:   "jicofo",
	Short: "Export Jicofo metrics",
	Run: func(cmd *cobra.Command, args []string) {
		prober = jicofo.NewProber(httpClient, jicofoUrl)
	},
}

func init() {
	rootCmd.AddCommand(jicofoCmd)
	jicofoCmd.Flags().StringVar(&jicofoUrl, "url", "http://127.0.0.1:8888/stats", "URL of the Jicofo stats endpoint")
}
