package jvb

import (
	"io"
	"net/http"

	"github.com/jitsi-contrib/jitsi-exporter/pkg/generic"
	"github.com/prometheus/client_golang/prometheus"
	log "github.com/sirupsen/logrus"
)

type prober struct {
	client *http.Client
	target string
}

func NewProber(client *http.Client, target string) generic.Prober {
	return &prober{
		client: client,
		target: target,
	}
}

func (p *prober) Name() string {
	return "jvb"
}

func (p *prober) Labels() []string {
	return []string{
		"version",
		"region",
		// TODO relay_id ?
	}
}

func (c *prober) Probe() (string, bool) {
	resp, err := c.client.Get(c.target)
	if err != nil {
		log.Error(err)
		return "", false
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Error(err)
		return "", false
	}

	return string(data), true
}

func (p *prober) MetricGetters() []generic.MetricGetter {
	return []generic.MetricGetter{
		generic.NewMetricGetter(prometheus.GaugeValue, "average_participant_stress", "TDOD"),
		generic.NewMetricGetter(prometheus.GaugeValue, "bit_rate_download", "the current incoming bitrate (RTP) in kilobits per second"),
		generic.NewMetricGetter(prometheus.GaugeValue, "bit_rate_upload", "the current outgoing bitrate (RTP) in kilobits per second"),
		generic.NewBoolMetricGetter(prometheus.GaugeValue, "colibri2", "TDOD"),
		// TODO conference_sizes
		generic.NewMetricGetter(prometheus.GaugeValue, "conferences", "the current number of conferences"),
		// TODO conferences_by_audio_senders
		// TODO conferences_by_video_senders
		// TODO current_timestamp ?
		generic.NewBoolMetricGetter(prometheus.GaugeValue, "drain", "TDOD"),
		generic.NewMetricGetter(prometheus.GaugeValue, "dtls_failed_endpoints", "the total number of endpoints which failed to establish a DTLS connection"),
		generic.NewMetricGetter(prometheus.GaugeValue, "endpoints", "the current number of endpoints, including octo endpoints"),
		generic.NewMetricGetter(prometheus.GaugeValue, "endpoints_sending_audio", "current number of endpoints sending (non-silence) audio"),
		generic.NewMetricGetter(prometheus.GaugeValue, "endpoints_sending_video", "current number of endpoints sending video"),
		generic.NewMetricGetter(prometheus.GaugeValue, "endpoints_with_high_outgoing_loss", "TDOD"),
		generic.NewMetricGetter(prometheus.GaugeValue, "endpoints_with_spurious_remb", "total number of endpoints which have sent an RTCP REMB packet when REMB was not signaled"),
		generic.NewMetricGetter(prometheus.GaugeValue, "endpoints_with_suspended_sources", "TDOD"),
		generic.NewBoolMetricGetter(prometheus.GaugeValue, "graceful_shutdown", "whether jitsi-videobridge is currently in graceful shutdown mode (hosting existing conferences, but not accepting new ones)"),
		generic.NewBoolMetricGetter(prometheus.GaugeValue, "healthy", "TDOD"),
		generic.NewMetricGetter(prometheus.GaugeValue, "inactive_conferences", "current number of conferences in which no endpoints are sending audio nor video. Note that this includes conferences which are currently using a peer-to-peer transport"),
		generic.NewMetricGetter(prometheus.GaugeValue, "inactive_endpoints", "current number of endpoints in inactive conferences (see inactive_conferences)"),
		generic.NewMetricGetter(prometheus.GaugeValue, "incoming_loss", "TDOD"),
		generic.NewMetricGetter(prometheus.GaugeValue, "jitter_aggregate", "TDOD"),
		generic.NewMetricGetter(prometheus.GaugeValue, "largest_conference", "the size of the current largest conference (counting all endpoints, including octo endpoints which are connected to a different jitsi-videobridge instance)"),
		generic.NewMetricGetter(prometheus.GaugeValue, "local_active_endpoints", "the current number of local endpoints (not octo) which are in an active conference. This includes endpoints which are not sending audio or video, but are in an active conference (i.e. they are receive-only)"),
		generic.NewMetricGetter(prometheus.GaugeValue, "local_endpoints", "the current number of local (non-octo) endpoints"),
		generic.NewMetricGetter(prometheus.GaugeValue, "muc_clients_configured", "TDOD"),
		generic.NewMetricGetter(prometheus.GaugeValue, "muc_clients_connected", "TDOD"),
		generic.NewMetricGetter(prometheus.GaugeValue, "mucs_configured", "TDOD"),
		generic.NewMetricGetter(prometheus.GaugeValue, "mucs_joined", "TDOD"),
		generic.NewMetricGetter(prometheus.GaugeValue, "num_eps_no_msg_transport_after_delay", "TDOD"),
		generic.NewMetricGetter(prometheus.GaugeValue, "num_eps_oversending", "current number of endpoints to which we are oversending"),
		generic.NewMetricGetter(prometheus.GaugeValue, "num_relays_no_msg_transport_after_delay", "TDOD"),
		generic.NewMetricGetter(prometheus.GaugeValue, "octo_conferences", "current number of conferences in which octo is enabled"),
		generic.NewMetricGetter(prometheus.GaugeValue, "octo_endpoints", "current number of octo endpoints (connected to remove jitsi-videobridge instances)"),
		generic.NewMetricGetter(prometheus.GaugeValue, "octo_receive_bitrate", "current incoming bitrate on the octo channel (combined for all conferences) in bits per second"),
		generic.NewMetricGetter(prometheus.GaugeValue, "octo_receive_packet_rate", "current incoming packet rate on the octo channel (combined for all conferences) in packets per second"),
		generic.NewMetricGetter(prometheus.GaugeValue, "octo_send_bitrate", "current outgoing bitrate on the octo channel (combined for all conferences) in bits per second"),
		generic.NewMetricGetter(prometheus.GaugeValue, "octo_send_packet_rate", "current outgoing packet rate on the octo channel (combined for all conferences) in packets per second"),
		generic.NewMetricGetter(prometheus.GaugeValue, "outgoing_loss", "TDOD"),
		generic.NewMetricGetter(prometheus.GaugeValue, "overall_loss", "TDOD"),
		generic.NewMetricGetter(prometheus.GaugeValue, "p2p_conferences", "current number of peer-to-peer conferences. These are conferences of size 2 in which no endpoint is sending audio not video. Presumably the endpoints are using a peer-to-peer transport at this time"),
		generic.NewMetricGetter(prometheus.GaugeValue, "packet_rate_download", "current RTP incoming packet rate in packets per second"),
		generic.NewMetricGetter(prometheus.GaugeValue, "packet_rate_upload", "current RTP outgoing packet rate in packets per second"),
		generic.NewMetricGetter(prometheus.GaugeValue, "participants", "current number of endpoints, including octo endpoints. Deprecated."),
		generic.NewMetricGetter(prometheus.GaugeValue, "preemptive_kfr_sent", "total number of preemptive keyframe requests sent"),
		generic.NewMetricGetter(prometheus.GaugeValue, "preemptive_kfr_suppressed", "TDOD"),
		generic.NewMetricGetter(prometheus.GaugeValue, "receive_only_endpoints", "current number of endpoints which are not sending audio nor video"),
		generic.NewMetricGetter(prometheus.GaugeValue, "rtt_aggregate", "round-trip-time measured via RTCP averaged over all local endpoints with a valid RTT measurement in milliseconds"),
		generic.NewMetricGetter(prometheus.GaugeValue, "stress_level", "current stress level on the bridge, with 0 indicating no load and 1 indicating the load is at full capacity (though values >1 are permitted)"),
		generic.NewMetricGetter(prometheus.GaugeValue, "threads", "current number of JVM threads"),
		generic.NewMetricGetter(prometheus.CounterValue, "total_aimd_bwe_expirations", "TDOD"),
		generic.NewMetricGetter(prometheus.CounterValue, "total_bytes_received", "total number of bytes received in RTP"),
		generic.NewMetricGetter(prometheus.CounterValue, "total_bytes_received_octo", "total number of bytes received on the octo channel"),
		generic.NewMetricGetter(prometheus.CounterValue, "total_bytes_sent", "total number of bytes sent in RTP"),
		generic.NewMetricGetter(prometheus.CounterValue, "total_bytes_sent_octo", "total number of bytes sent on the octo channel"),
		generic.NewMetricGetter(prometheus.CounterValue, "total_colibri_web_socket_messages_received", "total number of messages received on a Colibri bridge channel messages received on a WebSocket"),
		generic.NewMetricGetter(prometheus.CounterValue, "total_colibri_web_socket_messages_sent", "total number of messages sent over a Colibri bridge channel messages sent over a WebSocket"),
		generic.NewMetricGetter(prometheus.CounterValue, "total_conference_seconds", "total number of conference-seconds served (only updates once a conference expires)"),
		generic.NewMetricGetter(prometheus.CounterValue, "total_conferences_completed", "total number of conferences completed"),
		generic.NewMetricGetter(prometheus.CounterValue, "total_conferences_created", "total number of conferences created"),
		generic.NewMetricGetter(prometheus.CounterValue, "total_data_channel_messages_received", "total number of Colibri bridge channel messages received on SCTP data channels"),
		generic.NewMetricGetter(prometheus.CounterValue, "total_data_channel_messages_sent", "total number of Colibri bridge channel messages sent over SCTP data channels"),
		generic.NewMetricGetter(prometheus.CounterValue, "total_dominant_speaker_changes", "total number of times the dominant speaker in a conference changed"),
		generic.NewMetricGetter(prometheus.CounterValue, "total_failed_conferences", "total number of conferences in which no endpoints succeeded to establish an ICE connection"),
		generic.NewMetricGetter(prometheus.CounterValue, "total_ice_failed", "total number of endpoints which failed to establish an ICE connection"),
		generic.NewMetricGetter(prometheus.CounterValue, "total_ice_succeeded", "total number of endpoints which successfully established an ICE connection"),
		generic.NewMetricGetter(prometheus.CounterValue, "total_ice_succeeded_relayed", "total number of endpoints which connected through a TURN relay (currently broken)"),
		generic.NewMetricGetter(prometheus.CounterValue, "total_ice_succeeded_tcp", "total number of endpoints which connected through via ICE/TCP (currently broken)"),
		generic.NewMetricGetter(prometheus.CounterValue, "total_keyframes_received", "TDOD"),
		generic.NewMetricGetter(prometheus.CounterValue, "total_layering_changes_received", "TDOD"),
		generic.NewMetricGetter(prometheus.CounterValue, "total_loss_controlled_participant_seconds", "TDOD"),
		generic.NewMetricGetter(prometheus.CounterValue, "total_loss_degraded_participant_seconds", "TDOD"),
		generic.NewMetricGetter(prometheus.CounterValue, "total_loss_limited_participant_seconds", "TDOD"),
		generic.NewMetricGetter(prometheus.CounterValue, "total_packets_received", "total number of RTP packets received"),
		generic.NewMetricGetter(prometheus.CounterValue, "total_packets_received_octo", "total number packets received on the octo channel"),
		generic.NewMetricGetter(prometheus.CounterValue, "total_packets_sent", "total number of RTP packets sent"),
		generic.NewMetricGetter(prometheus.CounterValue, "total_packets_sent_octo", "total number packets sent over the octo channel"),
		generic.NewMetricGetter(prometheus.CounterValue, "total_partially_failed_conferences", "total number of conferences in which at least one endpoint failed to establish an ICE connection"),
		generic.NewMetricGetter(prometheus.CounterValue, "total_participants", "total number of endpoints created"),
		generic.NewMetricGetter(prometheus.CounterValue, "total_relays", "TDOD"),
		generic.NewMetricGetter(prometheus.CounterValue, "total_video_stream_milliseconds_received", "TDOD"),
		generic.NewMetricGetter(prometheus.GaugeValue, "videochannels", "current number of endpoints with a video channel (i.e. which support receiving video). Deprecated."),
	}
}
