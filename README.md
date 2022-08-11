# Jitsi Prometheus Exporter
Exporter that grabs various metrics from [Jitsi](https://jitsi.org), especially form the video bridges, and publishes them as [Prometheus](https://prometheus.io) metrics.

There is a [documentation](https://github.com/jitsi/jitsi-videobridge/blob/master/doc/statistics.md) of the published statistics by the video bridges.

## Install

Grab the binary from the release. Or user the `ghcr.io/jitsi-contrib/jitsi-exporter` container image.

## Usage

Run this as a sidecar of either Jicofo or Jvb.

```shel
jitsi-exporter jicfo
```

```shel
jitsi-exporter jvb
```

Defaults should work, but some flags are available. For more information add the flag `--help`.
