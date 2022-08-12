package cmd

import (
	"net/http"

	"github.com/jitsi-contrib/jitsi-exporter/pkg/generic"
	"github.com/jitsi-contrib/jitsi-exporter/pkg/jicofo"
	"github.com/jitsi-contrib/jitsi-exporter/pkg/jvb"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var metricsPath, metricsAddr, logLevel, jicofoUrl, jvbUrl string
var httpClient = &http.Client{
	Transport: &http.Transport{
		MaxIdleConns: 100,
	},
}

var rootCmd = &cobra.Command{
	Use:       "jitsi-exporter [components]",
	Short:     "A Prometheus exporter for Jitsi. Supports Jicofo and JVB.",
	ValidArgs: []string{"jicofo", "jvb"},
	Args:      cobra.MatchAll(cobra.RangeArgs(1, 2), cobra.OnlyValidArgs),
	Run: func(cmd *cobra.Command, args []string) {
		lvl, _ := log.ParseLevel(logLevel)
		log.SetLevel(lvl)
		log.Info("Starting Jitsi Colibri Exporter")

		collectors := []prometheus.Collector{}
		for _, arg := range args {
			if arg == "jicofo" {
				log.Info("Collect Jicofo metrics")
				collectors = append(collectors, generic.NewCollector(httpClient, jicofoUrl, jicofo.Metrics))
			}
			if arg == "jvb" {
				log.Info("Collect JVB metrics")
				collectors = append(collectors, generic.NewCollector(httpClient, jvbUrl, jvb.Metrics))
			}
		}
		prometheus.MustRegister(collectors...)

		http.Handle(metricsPath, promhttp.Handler())
		log.Fatal(http.ListenAndServe(metricsAddr, nil))
	},
}

func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	rootCmd.PersistentFlags().StringVar(&metricsPath, "path", "/metrics", "Path of the metrics")
	rootCmd.PersistentFlags().StringVar(&metricsAddr, "address", ":9210", "Address to listen")
	rootCmd.PersistentFlags().StringVar(&logLevel, "log-level", "info", "Log level.")
	rootCmd.Flags().StringVar(&jicofoUrl, "jicofo-url", "http://127.0.0.1:8888/stats", "URL of the Jicofo stats endpoint")
	rootCmd.Flags().StringVar(&jvbUrl, "jvb-url", "http://127.0.0.1:8080/colibri/stats", "URL of the JVB stats endpoint")
}
