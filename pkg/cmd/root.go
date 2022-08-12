package cmd

import (
	"net/http"

	"github.com/jitsi-contrib/jitsi-exporter/pkg/generic"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var metricsPath, metricsAddr, logLevel string
var prober generic.Prober
var httpClient = &http.Client{
	Transport: &http.Transport{
		MaxIdleConns: 100,
	},
}

var rootCmd = &cobra.Command{
	Use:   "jitsi-exporter",
	Short: "A Prometheus exporter for Jitsi. Supports Jicofo and JVB.",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		lvl, _ := log.ParseLevel(logLevel)
		log.SetLevel(lvl)

		log.Info("Starting Jitsi Colibri Exporter")
	},
	PersistentPostRun: func(cmd *cobra.Command, args []string) {
		collector := generic.NewCollector(prober)
		prometheus.MustRegister(collector)

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
}
