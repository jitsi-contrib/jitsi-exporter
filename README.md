# Jitsi Prometheus Exporter
Exporter that grabs various metrics from [Jitsi](https://jitsi.org), especially form the video bridges, and publishes them as [Prometheus](https://prometheus.io) metrics.

There is a [documentation](https://github.com/jitsi/jitsi-videobridge/blob/master/doc/statistics.md) of the published statistics by the video bridges.

## Install

Grab the binary from the release. Or user the `ghcr.io/jitsi-contrib/jitsi-exporter` container image.

## Usage

Run this as a sidecar of either Jicofo or Jvb (or both if single server).

```shel
jitsi-exporter jicofo jvb
```

Defaults should work, but some flags are available. For more information add the flag `--help`.

## Metrics

### Jicofo

```
# HELP jitsi_bridge_failures_participants_moved TODO
# TYPE jitsi_bridge_failures_participants_moved gauge
jitsi_bridge_failures_participants_moved 0
# HELP jitsi_bridge_selector_bridge_count TODO
# TYPE jitsi_bridge_selector_bridge_count gauge
jitsi_bridge_selector_bridge_count 1
# HELP jitsi_bridge_selector_in_shutdown_bridge_count TODO
# TYPE jitsi_bridge_selector_in_shutdown_bridge_count gauge
jitsi_bridge_selector_in_shutdown_bridge_count 0
# HELP jitsi_bridge_selector_lost_bridges TODO
# TYPE jitsi_bridge_selector_lost_bridges gauge
jitsi_bridge_selector_lost_bridges 0
# HELP jitsi_bridge_selector_operational_bridge_count TODO
# TYPE jitsi_bridge_selector_operational_bridge_count gauge
jitsi_bridge_selector_operational_bridge_count 1
# HELP jitsi_bridge_selector_total_least_loaded TODO
# TYPE jitsi_bridge_selector_total_least_loaded counter
jitsi_bridge_selector_total_least_loaded 0
# HELP jitsi_bridge_selector_total_least_loaded_in_conference TODO
# TYPE jitsi_bridge_selector_total_least_loaded_in_conference counter
jitsi_bridge_selector_total_least_loaded_in_conference 0
# HELP jitsi_bridge_selector_total_least_loaded_in_region TODO
# TYPE jitsi_bridge_selector_total_least_loaded_in_region counter
jitsi_bridge_selector_total_least_loaded_in_region 0
# HELP jitsi_bridge_selector_total_least_loaded_in_region_group TODO
# TYPE jitsi_bridge_selector_total_least_loaded_in_region_group counter
jitsi_bridge_selector_total_least_loaded_in_region_group 0
# HELP jitsi_bridge_selector_total_least_loaded_in_region_group_in_conference TODO
# TYPE jitsi_bridge_selector_total_least_loaded_in_region_group_in_conference counter
jitsi_bridge_selector_total_least_loaded_in_region_group_in_conference 0
# HELP jitsi_bridge_selector_total_least_loaded_in_region_in_conference TODO
# TYPE jitsi_bridge_selector_total_least_loaded_in_region_in_conference counter
jitsi_bridge_selector_total_least_loaded_in_region_in_conference 0
# HELP jitsi_bridge_selector_total_not_loaded_in_region TODO
# TYPE jitsi_bridge_selector_total_not_loaded_in_region counter
jitsi_bridge_selector_total_not_loaded_in_region 0
# HELP jitsi_bridge_selector_total_not_loaded_in_region_group TODO
# TYPE jitsi_bridge_selector_total_not_loaded_in_region_group counter
jitsi_bridge_selector_total_not_loaded_in_region_group 0
# HELP jitsi_bridge_selector_total_not_loaded_in_region_group_in_conference TODO
# TYPE jitsi_bridge_selector_total_not_loaded_in_region_group_in_conference counter
jitsi_bridge_selector_total_not_loaded_in_region_group_in_conference 0
# HELP jitsi_bridge_selector_total_not_loaded_in_region_in_conference TODO
# TYPE jitsi_bridge_selector_total_not_loaded_in_region_in_conference counter
jitsi_bridge_selector_total_not_loaded_in_region_in_conference 0
# HELP jitsi_bridge_selector_total_split_due_to_load TODO
# TYPE jitsi_bridge_selector_total_split_due_to_load counter
jitsi_bridge_selector_total_split_due_to_load 0
# HELP jitsi_bridge_selector_total_split_due_to_region TODO
# TYPE jitsi_bridge_selector_total_split_due_to_region counter
jitsi_bridge_selector_total_split_due_to_region 0
# HELP jitsi_conference_sizes TODO
# TYPE jitsi_conference_sizes summary
jitsi_conference_sizes{quantile="0.99"} -1
jitsi_conference_sizes{quantile="0.999"} -1
jitsi_conference_sizes_sum 0
jitsi_conference_sizes_count 0
# HELP jitsi_conference_sizes_average TODO
# TYPE jitsi_conference_sizes_average gauge
jitsi_conference_sizes_average 0
# HELP jitsi_conference_sizes_max TODO
# TYPE jitsi_conference_sizes_max gauge
jitsi_conference_sizes_max 0
# HELP jitsi_conference_sizes_min TODO
# TYPE jitsi_conference_sizes_min gauge
jitsi_conference_sizes_min 0
# HELP jitsi_conference_sizes_quantile TODO
# TYPE jitsi_conference_sizes_quantile histogram
jitsi_conference_sizes_quantile_bucket{le="1"} 0
jitsi_conference_sizes_quantile_bucket{le="2"} 0
jitsi_conference_sizes_quantile_bucket{le="3"} 0
jitsi_conference_sizes_quantile_bucket{le="5"} 0
jitsi_conference_sizes_quantile_bucket{le="10"} 0
jitsi_conference_sizes_quantile_bucket{le="20"} 0
jitsi_conference_sizes_quantile_bucket{le="50"} 0
jitsi_conference_sizes_quantile_bucket{le="100"} 0
jitsi_conference_sizes_quantile_bucket{le="200"} 0
jitsi_conference_sizes_quantile_bucket{le="300"} 0
jitsi_conference_sizes_quantile_bucket{le="400"} 0
jitsi_conference_sizes_quantile_bucket{le="500"} 0
jitsi_conference_sizes_quantile_bucket{le="+Inf"} 0
jitsi_conference_sizes_quantile_sum 0
jitsi_conference_sizes_quantile_count 0
# HELP jitsi_conferences TODO
# TYPE jitsi_conferences gauge
jitsi_conferences 0
# HELP jitsi_endpoint_pairs TODO
# TYPE jitsi_endpoint_pairs gauge
jitsi_endpoint_pairs 0
# HELP jitsi_jibri_detector_available TODO
# TYPE jitsi_jibri_detector_available gauge
jitsi_jibri_detector_available 1
# HELP jitsi_jibri_detector_count TODO
# TYPE jitsi_jibri_detector_count gauge
jitsi_jibri_detector_count 1
# HELP jitsi_jibri_live_streaming_active TODO
# TYPE jitsi_jibri_live_streaming_active gauge
jitsi_jibri_live_streaming_active 0
# HELP jitsi_jibri_live_streaming_pending TODO
# TYPE jitsi_jibri_live_streaming_pending gauge
jitsi_jibri_live_streaming_pending 0
# HELP jitsi_jibri_recording_active TODO
# TYPE jitsi_jibri_recording_active gauge
jitsi_jibri_recording_active 0
# HELP jitsi_jibri_recording_pending TODO
# TYPE jitsi_jibri_recording_pending gauge
jitsi_jibri_recording_pending 0
# HELP jitsi_jibri_sip_call_active TODO
# TYPE jitsi_jibri_sip_call_active gauge
jitsi_jibri_sip_call_active 0
# HELP jitsi_jibri_sip_call_pending TODO
# TYPE jitsi_jibri_sip_call_pending gauge
jitsi_jibri_sip_call_pending 0
# HELP jitsi_jibri_total_live_streaming_failures TODO
# TYPE jitsi_jibri_total_live_streaming_failures counter
jitsi_jibri_total_live_streaming_failures 0
# HELP jitsi_jibri_total_recording_failures TODO
# TYPE jitsi_jibri_total_recording_failures counter
jitsi_jibri_total_recording_failures 0
# HELP jitsi_jibri_total_sip_call_failures TODO
# TYPE jitsi_jibri_total_sip_call_failures counter
jitsi_jibri_total_sip_call_failures 0
# HELP jitsi_jicofo_healthy TODO
# TYPE jitsi_jicofo_healthy gauge
jitsi_jicofo_healthy 1
# HELP jitsi_jicofo_threads TODO
# TYPE jitsi_jicofo_threads gauge
jitsi_jicofo_threads 29
# HELP jitsi_largest_conference TODO
# TYPE jitsi_largest_conference gauge
jitsi_largest_conference 0
# HELP jitsi_participant_notifications_ice_failed TODO
# TYPE jitsi_participant_notifications_ice_failed gauge
jitsi_participant_notifications_ice_failed 0
# HELP jitsi_participant_notifications_request_restart TODO
# TYPE jitsi_participant_notifications_request_restart gauge
jitsi_participant_notifications_request_restart 0
# HELP jitsi_participants TODO
# TYPE jitsi_participants gauge
jitsi_participants 0
# HELP jitsi_queues_jibri_iq_queue_added_packets TODO
# TYPE jitsi_queues_jibri_iq_queue_added_packets gauge
jitsi_queues_jibri_iq_queue_added_packets 0
# HELP jitsi_queues_jibri_iq_queue_dropped_packets TODO
# TYPE jitsi_queues_jibri_iq_queue_dropped_packets gauge
jitsi_queues_jibri_iq_queue_dropped_packets 0
# HELP jitsi_queues_jibri_iq_queue_queue_size_at_remove TODO
# TYPE jitsi_queues_jibri_iq_queue_queue_size_at_remove summary
jitsi_queues_jibri_iq_queue_queue_size_at_remove{quantile="0.99"} -1
jitsi_queues_jibri_iq_queue_queue_size_at_remove{quantile="0.999"} -1
jitsi_queues_jibri_iq_queue_queue_size_at_remove_sum 0
jitsi_queues_jibri_iq_queue_queue_size_at_remove_count 0
# HELP jitsi_queues_jibri_iq_queue_queue_size_at_remove_average_queue_size_at_remove TODO
# TYPE jitsi_queues_jibri_iq_queue_queue_size_at_remove_average_queue_size_at_remove gauge
jitsi_queues_jibri_iq_queue_queue_size_at_remove_average_queue_size_at_remove 0
# HELP jitsi_queues_jibri_iq_queue_queue_size_at_remove_discarded TODO
# TYPE jitsi_queues_jibri_iq_queue_queue_size_at_remove_discarded gauge
jitsi_queues_jibri_iq_queue_queue_size_at_remove_discarded 0
# HELP jitsi_queues_jibri_iq_queue_queue_size_at_remove_max_queue_size_at_remove TODO
# TYPE jitsi_queues_jibri_iq_queue_queue_size_at_remove_max_queue_size_at_remove gauge
jitsi_queues_jibri_iq_queue_queue_size_at_remove_max_queue_size_at_remove 0
# HELP jitsi_queues_jibri_iq_queue_queue_size_at_remove_min_queue_size_at_remove TODO
# TYPE jitsi_queues_jibri_iq_queue_queue_size_at_remove_min_queue_size_at_remove gauge
jitsi_queues_jibri_iq_queue_queue_size_at_remove_min_queue_size_at_remove 0
# HELP jitsi_queues_jibri_iq_queue_queue_size_at_remove_quantile TODO
# TYPE jitsi_queues_jibri_iq_queue_queue_size_at_remove_quantile histogram
jitsi_queues_jibri_iq_queue_queue_size_at_remove_quantile_bucket{le="1"} 0
jitsi_queues_jibri_iq_queue_queue_size_at_remove_quantile_bucket{le="4"} 0
jitsi_queues_jibri_iq_queue_queue_size_at_remove_quantile_bucket{le="16"} 0
jitsi_queues_jibri_iq_queue_queue_size_at_remove_quantile_bucket{le="64"} 0
jitsi_queues_jibri_iq_queue_queue_size_at_remove_quantile_bucket{le="256"} 0
jitsi_queues_jibri_iq_queue_queue_size_at_remove_quantile_bucket{le="1024"} 0
jitsi_queues_jibri_iq_queue_queue_size_at_remove_quantile_bucket{le="4096"} 0
jitsi_queues_jibri_iq_queue_queue_size_at_remove_quantile_bucket{le="16384"} 0
jitsi_queues_jibri_iq_queue_queue_size_at_remove_quantile_bucket{le="+Inf"} 0
jitsi_queues_jibri_iq_queue_queue_size_at_remove_quantile_sum 0
jitsi_queues_jibri_iq_queue_queue_size_at_remove_quantile_count 0
# HELP jitsi_queues_jibri_iq_queue_removed_packets TODO
# TYPE jitsi_queues_jibri_iq_queue_removed_packets gauge
jitsi_queues_jibri_iq_queue_removed_packets 0
# HELP jitsi_slow_health_check TODO
# TYPE jitsi_slow_health_check gauge
jitsi_slow_health_check 0
# HELP jitsi_total_conferences_created TODO
# TYPE jitsi_total_conferences_created counter
jitsi_total_conferences_created 0
# HELP jitsi_total_participants TODO
# TYPE jitsi_total_participants counter
jitsi_total_participants 0
```

## Jvb

```
# HELP jitsi_jvb_average_participant_stress TDOD
# TYPE jitsi_jvb_average_participant_stress gauge
jitsi_jvb_average_participant_stress 0.01
# HELP jitsi_jvb_bit_rate_download the current incoming bitrate (RTP) in kilobits per second
# TYPE jitsi_jvb_bit_rate_download gauge
jitsi_jvb_bit_rate_download 0
# HELP jitsi_jvb_bit_rate_upload the current outgoing bitrate (RTP) in kilobits per second
# TYPE jitsi_jvb_bit_rate_upload gauge
jitsi_jvb_bit_rate_upload 0
# HELP jitsi_jvb_colibri2 TODO
# TYPE jitsi_jvb_colibri2 gauge
jitsi_jvb_colibri2 1
# HELP jitsi_jvb_conferences the current number of conferences
# TYPE jitsi_jvb_conferences gauge
jitsi_jvb_conferences 0
# HELP jitsi_jvb_drain TODO
# TYPE jitsi_jvb_drain gauge
jitsi_jvb_drain 0
# HELP jitsi_jvb_dtls_failed_endpoints the total number of endpoints which failed to establish a DTLS connection
# TYPE jitsi_jvb_dtls_failed_endpoints gauge
jitsi_jvb_dtls_failed_endpoints 0
# HELP jitsi_jvb_endpoints the current number of endpoints, including octo endpoints
# TYPE jitsi_jvb_endpoints gauge
jitsi_jvb_endpoints 0
# HELP jitsi_jvb_endpoints_sending_audio current number of endpoints sending (non-silence) audio
# TYPE jitsi_jvb_endpoints_sending_audio gauge
jitsi_jvb_endpoints_sending_audio 0
# HELP jitsi_jvb_endpoints_sending_video current number of endpoints sending video
# TYPE jitsi_jvb_endpoints_sending_video gauge
jitsi_jvb_endpoints_sending_video 0
# HELP jitsi_jvb_endpoints_with_high_outgoing_loss TDOD
# TYPE jitsi_jvb_endpoints_with_high_outgoing_loss gauge
jitsi_jvb_endpoints_with_high_outgoing_loss 0
# HELP jitsi_jvb_endpoints_with_spurious_remb total number of endpoints which have sent an RTCP REMB packet when REMB was not signaled
# TYPE jitsi_jvb_endpoints_with_spurious_remb gauge
jitsi_jvb_endpoints_with_spurious_remb 0
# HELP jitsi_jvb_endpoints_with_suspended_sources TDOD
# TYPE jitsi_jvb_endpoints_with_suspended_sources gauge
jitsi_jvb_endpoints_with_suspended_sources 0
# HELP jitsi_jvb_graceful_shutdown whether jitsi-videobridge is currently in graceful shutdown mode (hosting existing conferences, but not accepting new ones)
# TYPE jitsi_jvb_graceful_shutdown gauge
jitsi_jvb_graceful_shutdown 0
# HELP jitsi_jvb_healthy TODO
# TYPE jitsi_jvb_healthy gauge
jitsi_jvb_healthy 1
# HELP jitsi_jvb_inactive_conferences current number of conferences in which no endpoints are sending audio nor video. Note that this includes conferences which are currently using a peer-to-peer transport
# TYPE jitsi_jvb_inactive_conferences gauge
jitsi_jvb_inactive_conferences 0
# HELP jitsi_jvb_inactive_endpoints current number of endpoints in inactive conferences (see inactive_conferences)
# TYPE jitsi_jvb_inactive_endpoints gauge
jitsi_jvb_inactive_endpoints 0
# HELP jitsi_jvb_incoming_loss TDOD
# TYPE jitsi_jvb_incoming_loss gauge
jitsi_jvb_incoming_loss 0
# HELP jitsi_jvb_jitter_aggregate TDOD
# TYPE jitsi_jvb_jitter_aggregate gauge
jitsi_jvb_jitter_aggregate 0
# HELP jitsi_jvb_largest_conference the size of the current largest conference (counting all endpoints, including octo endpoints which are connected to a different jitsi-videobridge instance)
# TYPE jitsi_jvb_largest_conference gauge
jitsi_jvb_largest_conference 0
# HELP jitsi_jvb_local_active_endpoints the current number of local endpoints (not octo) which are in an active conference. This includes endpoints which are not sending audio or video, but are in an active conference (i.e. they are receive-only)
# TYPE jitsi_jvb_local_active_endpoints gauge
jitsi_jvb_local_active_endpoints 0
# HELP jitsi_jvb_local_endpoints the current number of local (non-octo) endpoints
# TYPE jitsi_jvb_local_endpoints gauge
jitsi_jvb_local_endpoints 0
# HELP jitsi_jvb_muc_clients_configured TDOD
# TYPE jitsi_jvb_muc_clients_configured gauge
jitsi_jvb_muc_clients_configured 1
# HELP jitsi_jvb_muc_clients_connected TDOD
# TYPE jitsi_jvb_muc_clients_connected gauge
jitsi_jvb_muc_clients_connected 1
# HELP jitsi_jvb_mucs_configured TDOD
# TYPE jitsi_jvb_mucs_configured gauge
jitsi_jvb_mucs_configured 1
# HELP jitsi_jvb_mucs_joined TDOD
# TYPE jitsi_jvb_mucs_joined gauge
jitsi_jvb_mucs_joined 1
# HELP jitsi_jvb_num_eps_no_msg_transport_after_delay TDOD
# TYPE jitsi_jvb_num_eps_no_msg_transport_after_delay gauge
jitsi_jvb_num_eps_no_msg_transport_after_delay 0
# HELP jitsi_jvb_num_eps_oversending current number of endpoints to which we are oversending
# TYPE jitsi_jvb_num_eps_oversending gauge
jitsi_jvb_num_eps_oversending 0
# HELP jitsi_jvb_num_relays_no_msg_transport_after_delay TDOD
# TYPE jitsi_jvb_num_relays_no_msg_transport_after_delay gauge
jitsi_jvb_num_relays_no_msg_transport_after_delay 0
# HELP jitsi_jvb_octo_conferences current number of conferences in which octo is enabled
# TYPE jitsi_jvb_octo_conferences gauge
jitsi_jvb_octo_conferences 0
# HELP jitsi_jvb_octo_endpoints current number of octo endpoints (connected to remove jitsi-videobridge instances)
# TYPE jitsi_jvb_octo_endpoints gauge
jitsi_jvb_octo_endpoints 0
# HELP jitsi_jvb_octo_receive_bitrate current incoming bitrate on the octo channel (combined for all conferences) in bits per second
# TYPE jitsi_jvb_octo_receive_bitrate gauge
jitsi_jvb_octo_receive_bitrate 0
# HELP jitsi_jvb_octo_receive_packet_rate current incoming packet rate on the octo channel (combined for all conferences) in packets per second
# TYPE jitsi_jvb_octo_receive_packet_rate gauge
jitsi_jvb_octo_receive_packet_rate 0
# HELP jitsi_jvb_octo_send_bitrate current outgoing bitrate on the octo channel (combined for all conferences) in bits per second
# TYPE jitsi_jvb_octo_send_bitrate gauge
jitsi_jvb_octo_send_bitrate 0
# HELP jitsi_jvb_octo_send_packet_rate current outgoing packet rate on the octo channel (combined for all conferences) in packets per second
# TYPE jitsi_jvb_octo_send_packet_rate gauge
jitsi_jvb_octo_send_packet_rate 0
# HELP jitsi_jvb_outgoing_loss TDOD
# TYPE jitsi_jvb_outgoing_loss gauge
jitsi_jvb_outgoing_loss 0
# HELP jitsi_jvb_overall_loss TDOD
# TYPE jitsi_jvb_overall_loss gauge
jitsi_jvb_overall_loss 0
# HELP jitsi_jvb_p2p_conferences current number of peer-to-peer conferences. These are conferences of size 2 in which no endpoint is sending audio not video. Presumably the endpoints are using a peer-to-peer transport at this time
# TYPE jitsi_jvb_p2p_conferences gauge
jitsi_jvb_p2p_conferences 0
# HELP jitsi_jvb_packet_rate_download current RTP incoming packet rate in packets per second
# TYPE jitsi_jvb_packet_rate_download gauge
jitsi_jvb_packet_rate_download 0
# HELP jitsi_jvb_packet_rate_upload current RTP outgoing packet rate in packets per second
# TYPE jitsi_jvb_packet_rate_upload gauge
jitsi_jvb_packet_rate_upload 0
# HELP jitsi_jvb_participants current number of endpoints, including octo endpoints. Deprecated.
# TYPE jitsi_jvb_participants gauge
jitsi_jvb_participants 0
# HELP jitsi_jvb_preemptive_kfr_sent total number of preemptive keyframe requests sent
# TYPE jitsi_jvb_preemptive_kfr_sent gauge
jitsi_jvb_preemptive_kfr_sent 0
# HELP jitsi_jvb_preemptive_kfr_suppressed TDOD
# TYPE jitsi_jvb_preemptive_kfr_suppressed gauge
jitsi_jvb_preemptive_kfr_suppressed 0
# HELP jitsi_jvb_receive_only_endpoints current number of endpoints which are not sending audio nor video
# TYPE jitsi_jvb_receive_only_endpoints gauge
jitsi_jvb_receive_only_endpoints 0
# HELP jitsi_jvb_rtt_aggregate round-trip-time measured via RTCP averaged over all local endpoints with a valid RTT measurement in milliseconds
# TYPE jitsi_jvb_rtt_aggregate gauge
jitsi_jvb_rtt_aggregate 0
# HELP jitsi_jvb_stress_level current stress level on the bridge, with 0 indicating no load and 1 indicating the load is at full capacity (though values >1 are permitted)
# TYPE jitsi_jvb_stress_level gauge
jitsi_jvb_stress_level 0
# HELP jitsi_jvb_threads current number of JVM threads
# TYPE jitsi_jvb_threads gauge
jitsi_jvb_threads 36
# HELP jitsi_jvb_total_aimd_bwe_expirations TDOD
# TYPE jitsi_jvb_total_aimd_bwe_expirations counter
jitsi_jvb_total_aimd_bwe_expirations 0
# HELP jitsi_jvb_total_bytes_received total number of bytes received in RTP
# TYPE jitsi_jvb_total_bytes_received counter
jitsi_jvb_total_bytes_received 0
# HELP jitsi_jvb_total_bytes_received_octo total number of bytes received on the octo channel
# TYPE jitsi_jvb_total_bytes_received_octo counter
jitsi_jvb_total_bytes_received_octo 0
# HELP jitsi_jvb_total_bytes_sent total number of bytes sent in RTP
# TYPE jitsi_jvb_total_bytes_sent counter
jitsi_jvb_total_bytes_sent 0
# HELP jitsi_jvb_total_bytes_sent_octo total number of bytes sent on the octo channel
# TYPE jitsi_jvb_total_bytes_sent_octo counter
jitsi_jvb_total_bytes_sent_octo 0
# HELP jitsi_jvb_total_colibri_web_socket_messages_received total number of messages received on a Colibri bridge channel messages received on a WebSocket
# TYPE jitsi_jvb_total_colibri_web_socket_messages_received counter
jitsi_jvb_total_colibri_web_socket_messages_received 0
# HELP jitsi_jvb_total_colibri_web_socket_messages_sent total number of messages sent over a Colibri bridge channel messages sent over a WebSocket
# TYPE jitsi_jvb_total_colibri_web_socket_messages_sent counter
jitsi_jvb_total_colibri_web_socket_messages_sent 0
# HELP jitsi_jvb_total_conference_seconds total number of conference-seconds served (only updates once a conference expires)
# TYPE jitsi_jvb_total_conference_seconds counter
jitsi_jvb_total_conference_seconds 0
# HELP jitsi_jvb_total_conferences_completed total number of conferences completed
# TYPE jitsi_jvb_total_conferences_completed counter
jitsi_jvb_total_conferences_completed 0
# HELP jitsi_jvb_total_conferences_created total number of conferences created
# TYPE jitsi_jvb_total_conferences_created counter
jitsi_jvb_total_conferences_created 0
# HELP jitsi_jvb_total_data_channel_messages_received total number of Colibri bridge channel messages received on SCTP data channels
# TYPE jitsi_jvb_total_data_channel_messages_received counter
jitsi_jvb_total_data_channel_messages_received 0
# HELP jitsi_jvb_total_data_channel_messages_sent total number of Colibri bridge channel messages sent over SCTP data channels
# TYPE jitsi_jvb_total_data_channel_messages_sent counter
jitsi_jvb_total_data_channel_messages_sent 0
# HELP jitsi_jvb_total_dominant_speaker_changes total number of times the dominant speaker in a conference changed
# TYPE jitsi_jvb_total_dominant_speaker_changes counter
jitsi_jvb_total_dominant_speaker_changes 0
# HELP jitsi_jvb_total_failed_conferences total number of conferences in which no endpoints succeeded to establish an ICE connection
# TYPE jitsi_jvb_total_failed_conferences counter
jitsi_jvb_total_failed_conferences 0
# HELP jitsi_jvb_total_ice_failed total number of endpoints which failed to establish an ICE connection
# TYPE jitsi_jvb_total_ice_failed counter
jitsi_jvb_total_ice_failed 0
# HELP jitsi_jvb_total_ice_succeeded total number of endpoints which successfully established an ICE connection
# TYPE jitsi_jvb_total_ice_succeeded counter
jitsi_jvb_total_ice_succeeded 0
# HELP jitsi_jvb_total_ice_succeeded_relayed total number of endpoints which connected through a TURN relay (currently broken)
# TYPE jitsi_jvb_total_ice_succeeded_relayed counter
jitsi_jvb_total_ice_succeeded_relayed 0
# HELP jitsi_jvb_total_ice_succeeded_tcp total number of endpoints which connected through via ICE/TCP (currently broken)
# TYPE jitsi_jvb_total_ice_succeeded_tcp counter
jitsi_jvb_total_ice_succeeded_tcp 0
# HELP jitsi_jvb_total_keyframes_received TDOD
# TYPE jitsi_jvb_total_keyframes_received counter
jitsi_jvb_total_keyframes_received 0
# HELP jitsi_jvb_total_layering_changes_received TDOD
# TYPE jitsi_jvb_total_layering_changes_received counter
jitsi_jvb_total_layering_changes_received 0
# HELP jitsi_jvb_total_loss_controlled_participant_seconds TDOD
# TYPE jitsi_jvb_total_loss_controlled_participant_seconds counter
jitsi_jvb_total_loss_controlled_participant_seconds 0
# HELP jitsi_jvb_total_loss_degraded_participant_seconds TDOD
# TYPE jitsi_jvb_total_loss_degraded_participant_seconds counter
jitsi_jvb_total_loss_degraded_participant_seconds 0
# HELP jitsi_jvb_total_loss_limited_participant_seconds TDOD
# TYPE jitsi_jvb_total_loss_limited_participant_seconds counter
jitsi_jvb_total_loss_limited_participant_seconds 0
# HELP jitsi_jvb_total_packets_received total number of RTP packets received
# TYPE jitsi_jvb_total_packets_received counter
jitsi_jvb_total_packets_received 0
# HELP jitsi_jvb_total_packets_received_octo total number packets received on the octo channel
# TYPE jitsi_jvb_total_packets_received_octo counter
jitsi_jvb_total_packets_received_octo 0
# HELP jitsi_jvb_total_packets_sent total number of RTP packets sent
# TYPE jitsi_jvb_total_packets_sent counter
jitsi_jvb_total_packets_sent 0
# HELP jitsi_jvb_total_packets_sent_octo total number packets sent over the octo channel
# TYPE jitsi_jvb_total_packets_sent_octo counter
jitsi_jvb_total_packets_sent_octo 0
# HELP jitsi_jvb_total_partially_failed_conferences total number of conferences in which at least one endpoint failed to establish an ICE connection
# TYPE jitsi_jvb_total_partially_failed_conferences counter
jitsi_jvb_total_partially_failed_conferences 0
# HELP jitsi_jvb_total_participants total number of endpoints created
# TYPE jitsi_jvb_total_participants counter
jitsi_jvb_total_participants 0
# HELP jitsi_jvb_total_relays TDOD
# TYPE jitsi_jvb_total_relays counter
jitsi_jvb_total_relays 0
# HELP jitsi_jvb_total_video_stream_milliseconds_received TDOD
# TYPE jitsi_jvb_total_video_stream_milliseconds_received counter
jitsi_jvb_total_video_stream_milliseconds_received 0
# HELP jitsi_jvb_videochannels current number of endpoints with a video channel (i.e. which support receiving video). Deprecated.
# TYPE jitsi_jvb_videochannels gauge
jitsi_jvb_videochannels 0
```