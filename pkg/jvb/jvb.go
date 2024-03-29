package jvb

import (
	"github.com/jitsi-contrib/jitsi-exporter/pkg/generic"
	"github.com/prometheus/client_golang/prometheus"
)

var Component = "jvb"

var Metrics = []generic.Metric{
	generic.NewBasicMetric("average_participant_stress", prometheus.GaugeValue, "TDOD", generic.WithComponent(Component)),
	generic.NewBasicMetric("bit_rate_download", prometheus.GaugeValue, "the current incoming bitrate (RTP) in kilobits per second", generic.WithComponent(Component)),
	generic.NewBasicMetric("bit_rate_upload", prometheus.GaugeValue, "the current outgoing bitrate (RTP) in kilobits per second", generic.WithComponent(Component)),
	generic.NewMetric(
		prometheus.NewDesc(prometheus.BuildFQName(generic.Namespace, Component, "colibri2"), "TODO", nil, nil),
		generic.BoolGetter("colibri2", prometheus.GaugeValue, nil),
	),
	// TODO conference_sizes
	generic.NewBasicMetric("conferences", prometheus.GaugeValue, "the current number of conferences", generic.WithComponent(Component)),
	// TODO conferences_by_audio_senders
	// TODO conferences_by_video_senders
	// TODO current_timestamp ?
	generic.NewMetric(
		prometheus.NewDesc(prometheus.BuildFQName(generic.Namespace, Component, "drain"), "TODO", nil, nil),
		generic.BoolGetter("drain", prometheus.GaugeValue, nil),
	),
	generic.NewBasicMetric("dtls_failed_endpoints", prometheus.GaugeValue, "the total number of endpoints which failed to establish a DTLS connection", generic.WithComponent(Component)),
	generic.NewBasicMetric("endpoints", prometheus.GaugeValue, "the current number of endpoints, including octo endpoints", generic.WithComponent(Component)),
	generic.NewBasicMetric("endpoints_sending_audio", prometheus.GaugeValue, "current number of endpoints sending (non-silence) audio", generic.WithComponent(Component)),
	generic.NewBasicMetric("endpoints_sending_video", prometheus.GaugeValue, "current number of endpoints sending video", generic.WithComponent(Component)),
	generic.NewBasicMetric("endpoints_with_high_outgoing_loss", prometheus.GaugeValue, "TDOD", generic.WithComponent(Component)),
	generic.NewBasicMetric("endpoints_with_spurious_remb", prometheus.GaugeValue, "total number of endpoints which have sent an RTCP REMB packet when REMB was not signaled", generic.WithComponent(Component)),
	generic.NewBasicMetric("endpoints_with_suspended_sources", prometheus.GaugeValue, "TDOD", generic.WithComponent(Component)),
	generic.NewMetric(
		prometheus.NewDesc(prometheus.BuildFQName(generic.Namespace, Component, "graceful_shutdown"), "whether jitsi-videobridge is currently in graceful shutdown mode (hosting existing conferences, but not accepting new ones)", nil, nil),
		generic.BoolGetter("graceful_shutdown", prometheus.GaugeValue, nil),
	),
	generic.NewMetric(
		prometheus.NewDesc(prometheus.BuildFQName(generic.Namespace, Component, "healthy"), "TODO", nil, nil),
		generic.BoolGetter("healthy", prometheus.GaugeValue, nil),
	),
	generic.NewBasicMetric("inactive_conferences", prometheus.GaugeValue, "current number of conferences in which no endpoints are sending audio nor video. Note that this includes conferences which are currently using a peer-to-peer transport", generic.WithComponent(Component)),
	generic.NewBasicMetric("inactive_endpoints", prometheus.GaugeValue, "current number of endpoints in inactive conferences (see inactive_conferences)", generic.WithComponent(Component)),
	generic.NewBasicMetric("incoming_loss", prometheus.GaugeValue, "TDOD", generic.WithComponent(Component)),
	generic.NewBasicMetric("jitter_aggregate", prometheus.GaugeValue, "TDOD", generic.WithComponent(Component)),
	generic.NewBasicMetric("largest_conference", prometheus.GaugeValue, "the size of the current largest conference (counting all endpoints, including octo endpoints which are connected to a different jitsi-videobridge instance)", generic.WithComponent(Component)),
	generic.NewBasicMetric("local_active_endpoints", prometheus.GaugeValue, "the current number of local endpoints (not octo) which are in an active conference. This includes endpoints which are not sending audio or video, but are in an active conference (i.e. they are receive-only)", generic.WithComponent(Component)),
	generic.NewBasicMetric("local_endpoints", prometheus.GaugeValue, "the current number of local (non-octo) endpoints", generic.WithComponent(Component)),
	generic.NewBasicMetric("muc_clients_configured", prometheus.GaugeValue, "TDOD", generic.WithComponent(Component)),
	generic.NewBasicMetric("muc_clients_connected", prometheus.GaugeValue, "TDOD", generic.WithComponent(Component)),
	generic.NewBasicMetric("mucs_configured", prometheus.GaugeValue, "TDOD", generic.WithComponent(Component)),
	generic.NewBasicMetric("mucs_joined", prometheus.GaugeValue, "TDOD", generic.WithComponent(Component)),
	generic.NewBasicMetric("num_eps_no_msg_transport_after_delay", prometheus.GaugeValue, "TDOD", generic.WithComponent(Component)),
	generic.NewBasicMetric("num_eps_oversending", prometheus.GaugeValue, "current number of endpoints to which we are oversending", generic.WithComponent(Component)),
	generic.NewBasicMetric("num_relays_no_msg_transport_after_delay", prometheus.GaugeValue, "TDOD", generic.WithComponent(Component)),
	generic.NewBasicMetric("octo_conferences", prometheus.GaugeValue, "current number of conferences in which octo is enabled", generic.WithComponent(Component)),
	generic.NewBasicMetric("octo_endpoints", prometheus.GaugeValue, "current number of octo endpoints (connected to remove jitsi-videobridge instances)", generic.WithComponent(Component)),
	generic.NewBasicMetric("octo_receive_bitrate", prometheus.GaugeValue, "current incoming bitrate on the octo channel (combined for all conferences) in bits per second", generic.WithComponent(Component)),
	generic.NewBasicMetric("octo_receive_packet_rate", prometheus.GaugeValue, "current incoming packet rate on the octo channel (combined for all conferences) in packets per second", generic.WithComponent(Component)),
	generic.NewBasicMetric("octo_send_bitrate", prometheus.GaugeValue, "current outgoing bitrate on the octo channel (combined for all conferences) in bits per second", generic.WithComponent(Component)),
	generic.NewBasicMetric("octo_send_packet_rate", prometheus.GaugeValue, "current outgoing packet rate on the octo channel (combined for all conferences) in packets per second", generic.WithComponent(Component)),
	generic.NewBasicMetric("outgoing_loss", prometheus.GaugeValue, "TDOD", generic.WithComponent(Component)),
	generic.NewBasicMetric("overall_loss", prometheus.GaugeValue, "TDOD", generic.WithComponent(Component)),
	generic.NewBasicMetric("p2p_conferences", prometheus.GaugeValue, "current number of peer-to-peer conferences. These are conferences of size 2 in which no endpoint is sending audio not video. Presumably the endpoints are using a peer-to-peer transport at this time", generic.WithComponent(Component)),
	generic.NewBasicMetric("packet_rate_download", prometheus.GaugeValue, "current RTP incoming packet rate in packets per second", generic.WithComponent(Component)),
	generic.NewBasicMetric("packet_rate_upload", prometheus.GaugeValue, "current RTP outgoing packet rate in packets per second", generic.WithComponent(Component)),
	generic.NewBasicMetric("participants", prometheus.GaugeValue, "current number of endpoints, including octo endpoints. Deprecated.", generic.WithComponent(Component)),
	generic.NewBasicMetric("preemptive_kfr_sent", prometheus.GaugeValue, "total number of preemptive keyframe requests sent", generic.WithComponent(Component)),
	generic.NewBasicMetric("preemptive_kfr_suppressed", prometheus.GaugeValue, "TDOD", generic.WithComponent(Component)),
	generic.NewBasicMetric("receive_only_endpoints", prometheus.GaugeValue, "current number of endpoints which are not sending audio nor video", generic.WithComponent(Component)),
	generic.NewBasicMetric("rtt_aggregate", prometheus.GaugeValue, "round-trip-time measured via RTCP averaged over all local endpoints with a valid RTT measurement in milliseconds", generic.WithComponent(Component)),
	generic.NewBasicMetric("stress_level", prometheus.GaugeValue, "current stress level on the bridge, with 0 indicating no load and 1 indicating the load is at full capacity (though values >1 are permitted)", generic.WithComponent(Component)),
	generic.NewBasicMetric("threads", prometheus.GaugeValue, "current number of JVM threads", generic.WithComponent(Component)),
	generic.NewBasicMetric("total_aimd_bwe_expirations", prometheus.CounterValue, "TDOD", generic.WithComponent(Component)),
	generic.NewBasicMetric("total_bytes_received", prometheus.CounterValue, "total number of bytes received in RTP", generic.WithComponent(Component)),
	generic.NewBasicMetric("total_bytes_received_octo", prometheus.CounterValue, "total number of bytes received on the octo channel", generic.WithComponent(Component)),
	generic.NewBasicMetric("total_bytes_sent", prometheus.CounterValue, "total number of bytes sent in RTP", generic.WithComponent(Component)),
	generic.NewBasicMetric("total_bytes_sent_octo", prometheus.CounterValue, "total number of bytes sent on the octo channel", generic.WithComponent(Component)),
	generic.NewBasicMetric("total_colibri_web_socket_messages_received", prometheus.CounterValue, "total number of messages received on a Colibri bridge channel messages received on a WebSocket", generic.WithComponent(Component)),
	generic.NewBasicMetric("total_colibri_web_socket_messages_sent", prometheus.CounterValue, "total number of messages sent over a Colibri bridge channel messages sent over a WebSocket", generic.WithComponent(Component)),
	generic.NewBasicMetric("total_conference_seconds", prometheus.CounterValue, "total number of conference-seconds served (only updates once a conference expires)", generic.WithComponent(Component)),
	generic.NewBasicMetric("total_conferences_completed", prometheus.CounterValue, "total number of conferences completed", generic.WithComponent(Component)),
	generic.NewBasicMetric("total_conferences_created", prometheus.CounterValue, "total number of conferences created", generic.WithComponent(Component)),
	generic.NewBasicMetric("total_data_channel_messages_received", prometheus.CounterValue, "total number of Colibri bridge channel messages received on SCTP data channels", generic.WithComponent(Component)),
	generic.NewBasicMetric("total_data_channel_messages_sent", prometheus.CounterValue, "total number of Colibri bridge channel messages sent over SCTP data channels", generic.WithComponent(Component)),
	generic.NewBasicMetric("total_dominant_speaker_changes", prometheus.CounterValue, "total number of times the dominant speaker in a conference changed", generic.WithComponent(Component)),
	generic.NewBasicMetric("total_failed_conferences", prometheus.CounterValue, "total number of conferences in which no endpoints succeeded to establish an ICE connection", generic.WithComponent(Component)),
	generic.NewBasicMetric("total_ice_failed", prometheus.CounterValue, "total number of endpoints which failed to establish an ICE connection", generic.WithComponent(Component)),
	generic.NewBasicMetric("total_ice_succeeded", prometheus.CounterValue, "total number of endpoints which successfully established an ICE connection", generic.WithComponent(Component)),
	generic.NewBasicMetric("total_ice_succeeded_relayed", prometheus.CounterValue, "total number of endpoints which connected through a TURN relay (currently broken)", generic.WithComponent(Component)),
	generic.NewBasicMetric("total_ice_succeeded_tcp", prometheus.CounterValue, "total number of endpoints which connected through via ICE/TCP (currently broken)", generic.WithComponent(Component)),
	generic.NewBasicMetric("total_keyframes_received", prometheus.CounterValue, "TDOD", generic.WithComponent(Component)),
	generic.NewBasicMetric("total_layering_changes_received", prometheus.CounterValue, "TDOD", generic.WithComponent(Component)),
	generic.NewBasicMetric("total_loss_controlled_participant_seconds", prometheus.CounterValue, "TDOD", generic.WithComponent(Component)),
	generic.NewBasicMetric("total_loss_degraded_participant_seconds", prometheus.CounterValue, "TDOD", generic.WithComponent(Component)),
	generic.NewBasicMetric("total_loss_limited_participant_seconds", prometheus.CounterValue, "TDOD", generic.WithComponent(Component)),
	generic.NewBasicMetric("total_packets_received", prometheus.CounterValue, "total number of RTP packets received", generic.WithComponent(Component)),
	generic.NewBasicMetric("total_packets_received_octo", prometheus.CounterValue, "total number packets received on the octo channel", generic.WithComponent(Component)),
	generic.NewBasicMetric("total_packets_sent", prometheus.CounterValue, "total number of RTP packets sent", generic.WithComponent(Component)),
	generic.NewBasicMetric("total_packets_sent_octo", prometheus.CounterValue, "total number packets sent over the octo channel", generic.WithComponent(Component)),
	generic.NewBasicMetric("total_partially_failed_conferences", prometheus.CounterValue, "total number of conferences in which at least one endpoint failed to establish an ICE connection", generic.WithComponent(Component)),
	generic.NewBasicMetric("total_participants", prometheus.CounterValue, "total number of endpoints created", generic.WithComponent(Component)),
	generic.NewBasicMetric("total_relays", prometheus.CounterValue, "TDOD", generic.WithComponent(Component)),
	generic.NewBasicMetric("total_video_stream_milliseconds_received", prometheus.CounterValue, "TDOD", generic.WithComponent(Component)),
	generic.NewBasicMetric("videochannels", prometheus.GaugeValue, "current number of endpoints with a video channel (i.e. which support receiving video). Deprecated.", generic.WithComponent(Component)),
}
