package collector

// import (
// 	"encoding/json"
// 	"io"
// 	"net/http"
// )

// type ColibriMetrics struct {
// 	Threads                               float64  `json:"threads"`
// 	UsedMemory                            float64  `json:"used_memory"`
// 	TotalMemory                           float64  `json:"total_memory"`
// 	CPUUsage                              float64  `json:"cpu_usage"`
// 	BitRateDownload                       float64  `json:"bit_rate_download"`
// 	BitRateUpload                         float64  `json:"bit_rate_upload"`
// 	PacketRateDownload                    float64  `json:"packet_rate_download"`
// 	PacketRateUpload                      float64  `json:"packet_rate_upload"`
// 	LossRateDownload                      float64  `json:"loss_rate_download"`
// 	LossRateUpload                        float64  `json:"loss_rate_upload"`
// 	RTPLoss                               float64  `json:"rtp_loss"`
// 	JitterAggregate                       float64  `json:"jitter_aggregate"`
// 	RTTAggregate                          float64  `json:"rtt_aggregate"`
// 	LargestConference                     float64  `json:"largest_conference"`
// 	ConferenceSizes                       []uint64 `json:"conference_sizes"`
// 	Audiochannels                         float64  `json:"audiochannels"`
// 	Videochannels                         float64  `json:"videochannels"`
// 	EndpointsSendingAudio                 float64  `json:"endpoints_sending_audio"`
// 	EndpointsSendingVideo                 float64  `json:"endpoints_sending_video"`
// 	Conferences                           float64  `json:"conferences"`
// 	Participants                          float64  `json:"participants"`
// 	Videostreams                          float64  `json:"videostreams"`
// 	StressLevel                           float64  `json:"stress_level"`
// 	TotalLossControlledParticipantSeconds float64  `json:"total_loss_controlled_participant_seconds"`
// 	TotalLossLimitedParticipantSeconds    float64  `json:"total_loss_limited_participant_seconds"`
// 	TotalLossDegradedParticipantSeconds   float64  `json:"total_loss_degraded_participant_seconds"`
// 	TotalConferenceSeconds                float64  `json:"total_conference_seconds"`
// 	TotalConferencesCreated               float64  `json:"total_conferences_created"`
// 	TotalFailedConferences                float64  `json:"total_failed_conferences"`
// 	TotalPartiallyFailedConferences       float64  `json:"total_partially_failed_conferences"`
// 	TotalDataChannelMessagesReceived      float64  `json:"total_data_channel_messages_received"`
// 	TotalDataChannelMessagesSent          float64  `json:"total_data_channel_messages_sent"`
// 	TotalColibriWebsocketMessagesReceived float64  `json:"total_colibri_web_socket_messages_received"`
// 	TotalColibriWebsocketMessagesSent     float64  `json:"total_colibri_web_socket_messages_sent"`
// 	OctoSendBitRate                       float64  `json:"octo_send_bitrate"`
// 	OctoReceiveBitRate                    float64  `json:"octo_receive_bitrate"`
// 	OctoConferences                       float64  `json:"octo_conferences"`
// 	OctoEndpoints                         float64  `json:"octo_endpoints"`
// 	OctoSendPacketRate                    float64  `json:"octo_send_packet_rate"`
// 	OctoReceivePacketRate                 float64  `json:"octo_receive_packet_rate"`
// }

// func probeColibri(client *http.Client, target string) (ColibriMetrics, error) {
// 	var jsonData ColibriMetrics

// 	resp, err := client.Get(target)
// 	if err != nil {
// 		return jsonData, err
// 	}
// 	defer resp.Body.Close()

// 	bytes, err := io.ReadAll(resp.Body)
// 	if err != nil {
// 		return jsonData, err
// 	}

// 	err = json.Unmarshal([]byte(bytes), &jsonData)
// 	if err != nil {
// 		return jsonData, err
// 	}

// 	return jsonData, nil
// }

// // inactive_endpoints
// // inactive_conferences
// // total_ice_succeeded_relayed
// // total_ice_succeeded
// // octo_version
// // local_active_endpoints
// // muc_clients_connected
// // total_packets_received
// // p2p_conferences
// // total_dominant_speaker_changes
// // receive_only_endpoints
// // version
// // total_bytes_sent_octo
// // num_eps_oversending
// // total_conferences_completed
// // num_eps_no_msg_transport_after_delay
// // region
// // muc_clients_configured
// // outgoing_loss
// // overall_loss
// // total_packets_sent_octo
// // conferences_by_video_senders
// // endpoints_with_high_outgoing_loss
// // total_ice_succeeded_tcp
// // total_packets_dropped_octo
// // average_participant_stress
// // total_packets_sent
// // incoming_loss
// // total_bytes_received_octo
// // conferences_by_audio_senders
// // total_ice_failed
// // total_packets_received_octo
// // graceful_shutdown
// // total_bytes_received
// // dtls_failed_endpoints
// // total_bytes_sent
// // mucs_confi
// // mucs_joined
// // relay_id
