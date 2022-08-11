package collector

// import (
// 	"net/http"

// 	"github.com/prometheus/client_golang/prometheus"
// 	"github.com/sirupsen/logrus"
// )

// type Collector struct {
// 	client *http.Client
// 	target string
// 	log    *logrus.Logger

// 	up                                         *prometheus.Desc
// 	threads                                    *prometheus.Desc
// 	used_memory                                *prometheus.Desc
// 	total_memory                               *prometheus.Desc
// 	cpu_usage                                  *prometheus.Desc
// 	bit_rate_download                          *prometheus.Desc
// 	bit_rate_upload                            *prometheus.Desc
// 	packet_rate_download                       *prometheus.Desc
// 	packet_rate_upload                         *prometheus.Desc
// 	loss_rate_download                         *prometheus.Desc
// 	loss_rate_upload                           *prometheus.Desc
// 	rtp_loss                                   *prometheus.Desc
// 	jitter_aggregate                           *prometheus.Desc
// 	rtt_aggregate                              *prometheus.Desc
// 	largest_conference                         *prometheus.Desc
// 	conference_sizes                           *prometheus.Desc
// 	audiochannels                              *prometheus.Desc
// 	videochannels                              *prometheus.Desc
// 	endpoints_sending_audio                    *prometheus.Desc
// 	endpoints_sending_video                    *prometheus.Desc
// 	conferences                                *prometheus.Desc
// 	participants                               *prometheus.Desc
// 	videostreams                               *prometheus.Desc
// 	stress_level                               *prometheus.Desc
// 	total_loss_controlled_participant_seconds  *prometheus.Desc
// 	total_loss_limited_participant_seconds     *prometheus.Desc
// 	total_loss_degraded_participant_seconds    *prometheus.Desc
// 	total_conference_seconds                   *prometheus.Desc
// 	total_conferences_created                  *prometheus.Desc
// 	total_failed_conferences                   *prometheus.Desc
// 	total_partially_failed_conferences         *prometheus.Desc
// 	total_data_channel_messages_received       *prometheus.Desc
// 	total_data_channel_messages_sent           *prometheus.Desc
// 	total_colibri_web_socket_messages_received *prometheus.Desc
// 	total_colibri_web_socket_messages_sent     *prometheus.Desc
// 	octo_send_bitrate                          *prometheus.Desc
// 	octo_receive_bitrate                       *prometheus.Desc
// 	octo_conferences                           *prometheus.Desc
// 	octo_endpoints                             *prometheus.Desc
// 	octo_send_packet_rate                      *prometheus.Desc
// 	octo_receive_packet_rate                   *prometheus.Desc
// }

// type Target struct {
// 	Name     string
// 	Endpoint string
// }

// func NewJVBCollector(client *http.Client, targetURL string) *Collector {
// 	return &Collector{
// 		client: client,
// 		target: targetURL,

// 		up:                      prometheus.NewDesc(FQName("colibri", "up"), "Whether the Azure ServiceBus scrape was successful", nil, nil),
// 		threads:                 prometheus.NewDesc(FQName("colibri", "threads"), "threads", nil, nil),
// 		used_memory:             prometheus.NewDesc(FQName("colibri", "used_memory"), "used_memory", nil, nil),
// 		total_memory:            prometheus.NewDesc(FQName("colibri", "total_memory"), "total_memory", nil, nil),
// 		cpu_usage:               prometheus.NewDesc(FQName("colibri", "cpu_usage"), "cpu_usage", nil, nil),
// 		bit_rate_download:       prometheus.NewDesc(FQName("colibri", "bit_rate_download"), "bit_rate_download", nil, nil),
// 		bit_rate_upload:         prometheus.NewDesc(FQName("colibri", "bit_rate_upload"), "bit_rate_upload", nil, nil),
// 		packet_rate_download:    prometheus.NewDesc(FQName("colibri", "packet_rate_download"), "packet_rate_download", nil, nil),
// 		packet_rate_upload:      prometheus.NewDesc(FQName("colibri", "packet_rate_upload"), "packet_rate_upload", nil, nil),
// 		loss_rate_download:      prometheus.NewDesc(FQName("colibri", "loss_rate_download"), "loss_rate_download", nil, nil),
// 		loss_rate_upload:        prometheus.NewDesc(FQName("colibri", "loss_rate_upload"), "loss_rate_upload", nil, nil),
// 		rtp_loss:                prometheus.NewDesc(FQName("colibri", "rtp_loss"), "rtp_loss", nil, nil),
// 		jitter_aggregate:        prometheus.NewDesc(FQName("colibri", "jitter_aggregate"), "jitter_aggregate", nil, nil),
// 		rtt_aggregate:           prometheus.NewDesc(FQName("colibri", "rtt_aggregate"), "rtt_aggregate", nil, nil),
// 		largest_conference:      prometheus.NewDesc(FQName("colibri", "largest_conference"), "largest_conference", nil, nil),
// 		audiochannels:           prometheus.NewDesc(FQName("colibri", "audiochannels"), "audiochannels", nil, nil),
// 		videochannels:           prometheus.NewDesc(FQName("colibri", "videochannels"), "videochannels", nil, nil),
// 		endpoints_sending_audio: prometheus.NewDesc(FQName("colibri", "endpoints_sending_audio"), "endpoints_sending_audio", nil, nil),
// 		endpoints_sending_video: prometheus.NewDesc(FQName("colibri", "endpoints_sending_video"), "endpoints_sending_video", nil, nil),
// 		conferences:             prometheus.NewDesc(FQName("colibri", "conferences"), "conferences", nil, nil),
// 		participants:            prometheus.NewDesc(FQName("colibri", "participants"), "participants", nil, nil),
// 		videostreams:            prometheus.NewDesc(FQName("colibri", "videostreams"), "videostreams", nil, nil),
// 		stress_level:            prometheus.NewDesc(FQName("colibri", "stress_level"), "stress_level", nil, nil),
// 		total_loss_controlled_participant_seconds:  prometheus.NewDesc(FQName("colibri", "total_loss_controlled_participant_seconds"), "total_loss_controlled_participant_seconds", nil, nil),
// 		total_loss_limited_participant_seconds:     prometheus.NewDesc(FQName("colibri", "total_loss_limited_participant_seconds"), "total_loss_limited_participant_seconds", nil, nil),
// 		total_loss_degraded_participant_seconds:    prometheus.NewDesc(FQName("colibri", "total_loss_degraded_participant_seconds"), "total_loss_degraded_participant_seconds", nil, nil),
// 		total_conference_seconds:                   prometheus.NewDesc(FQName("colibri", "total_conference_seconds"), "total_conference_seconds", nil, nil),
// 		total_conferences_created:                  prometheus.NewDesc(FQName("colibri", "total_conferences_created"), "total_conferences_created", nil, nil),
// 		total_failed_conferences:                   prometheus.NewDesc(FQName("colibri", "total_failed_conferences"), "total_failed_conferences", nil, nil),
// 		total_partially_failed_conferences:         prometheus.NewDesc(FQName("colibri", "total_partially_failed_conferences"), "total_partially_failed_conferences", nil, nil),
// 		total_data_channel_messages_received:       prometheus.NewDesc(FQName("colibri", "total_data_channel_messages_received"), "total_data_channel_messages_received", nil, nil),
// 		total_data_channel_messages_sent:           prometheus.NewDesc(FQName("colibri", "total_data_channel_messages_sent"), "total_data_channel_messages_sent", nil, nil),
// 		total_colibri_web_socket_messages_received: prometheus.NewDesc(FQName("colibri", "total_colibri_web_socket_messages_received"), "total_colibri_web_socket_messages_received", nil, nil),
// 		total_colibri_web_socket_messages_sent:     prometheus.NewDesc(FQName("colibri", "total_colibri_web_socket_messages_sent"), "total_colibri_web_socket_messages_sent", nil, nil),
// 		conference_sizes:                           prometheus.NewDesc(FQName("colibri", "conference_sizes"), "conference_sizes", nil, nil),
// 		octo_send_bitrate:                          prometheus.NewDesc(FQName("colibri", "octo_send_bitrate"), "octo_send_bitrate", nil, nil),
// 		octo_receive_bitrate:                       prometheus.NewDesc(FQName("colibri", "octo_receive_bitrate"), "octo_receive_bitrate", nil, nil),
// 		octo_conferences:                           prometheus.NewDesc(FQName("colibri", "octo_conferences"), "octo_conferences", nil, nil),
// 		octo_endpoints:                             prometheus.NewDesc(FQName("colibri", "octo_endpoints"), "octo_endpoints", nil, nil),
// 		octo_send_packet_rate:                      prometheus.NewDesc(FQName("colibri", "octo_send_packet_rate"), "octo_send_packet_rate", nil, nil),
// 		octo_receive_packet_rate:                   prometheus.NewDesc(FQName("colibri", "octo_receive_packet_rate"), "octo_receive_packet_rate", nil, nil),
// 	}
// }

// func (c *Collector) Describe(ch chan<- *prometheus.Desc) {
// 	ch <- c.up
// 	ch <- c.threads
// 	ch <- c.used_memory
// 	ch <- c.total_memory
// 	ch <- c.cpu_usage
// 	ch <- c.bit_rate_download
// 	ch <- c.bit_rate_upload
// 	ch <- c.packet_rate_download
// 	ch <- c.packet_rate_upload
// 	ch <- c.loss_rate_download
// 	ch <- c.loss_rate_upload
// 	ch <- c.rtp_loss
// 	ch <- c.jitter_aggregate
// 	ch <- c.rtt_aggregate
// 	ch <- c.largest_conference
// 	ch <- c.conference_sizes
// 	ch <- c.audiochannels
// 	ch <- c.videochannels
// 	ch <- c.endpoints_sending_audio
// 	ch <- c.endpoints_sending_video
// 	ch <- c.conferences
// 	ch <- c.participants
// 	ch <- c.videostreams
// 	ch <- c.stress_level
// 	ch <- c.total_loss_controlled_participant_seconds
// 	ch <- c.total_loss_limited_participant_seconds
// 	ch <- c.total_loss_degraded_participant_seconds
// 	ch <- c.total_conference_seconds
// 	ch <- c.total_conferences_created
// 	ch <- c.total_failed_conferences
// 	ch <- c.total_partially_failed_conferences
// 	ch <- c.total_data_channel_messages_received
// 	ch <- c.total_data_channel_messages_sent
// 	ch <- c.total_colibri_web_socket_messages_received
// 	ch <- c.total_colibri_web_socket_messages_sent
// 	ch <- c.octo_send_bitrate
// 	ch <- c.octo_receive_bitrate
// 	ch <- c.octo_conferences
// 	ch <- c.octo_endpoints
// 	ch <- c.octo_send_packet_rate
// 	ch <- c.octo_receive_packet_rate
// }

// func conferenceSizesHelper(conferenceSizes []uint64) (conferenceSizesHistogram map[float64]uint64, sum uint64) {
// 	var sizes = make(map[float64]uint64)
// 	var values []uint64
// 	for _, v := range conferenceSizes {
// 		values = append(values, v)
// 	}

// 	//calculate sum (makes this metric independent from conferences metric)
// 	sum = 0
// 	for _, v := range values {
// 		sum += v
// 	}

// 	//for the histgram buckets we need to omit the last field b/c the +inf bucket is added automatically
// 	values = values[:len(values)-1]

// 	//the bucket values have to be cumulative
// 	var i int
// 	for i = len(values) - 1; i >= 0; i-- {
// 		var cumulative uint64
// 		var j int
// 		for j = i; j >= 0; j-- {
// 			cumulative += values[j]
// 		}
// 		values[i] = cumulative
// 	}

// 	for i, v := range values {
// 		sizes[float64(i)] = v
// 	}

// 	return sizes, sum
// }

// func (c *Collector) Collect(ch chan<- prometheus.Metric) {
// 	jsonData, err := probeColibri(c.client, c.target)

// 	if err != nil {
// 		c.log.Error(err)
// 		// client call failed, set the up metric value to 0
// 		ch <- prometheus.MustNewConstMetric(c.up, prometheus.GaugeValue, 0)

// 	} else {
// 		// client call succeeded, set the up metric value to 1
// 		ch <- prometheus.MustNewConstMetric(c.up, prometheus.GaugeValue, 1)
// 		ch <- prometheus.MustNewConstMetric(c.threads, prometheus.GaugeValue, jsonData.Threads)
// 		ch <- prometheus.MustNewConstMetric(c.used_memory, prometheus.GaugeValue, jsonData.UsedMemory)
// 		ch <- prometheus.MustNewConstMetric(c.total_memory, prometheus.GaugeValue, jsonData.TotalMemory)
// 		ch <- prometheus.MustNewConstMetric(c.cpu_usage, prometheus.GaugeValue, jsonData.CPUUsage)
// 		ch <- prometheus.MustNewConstMetric(c.bit_rate_download, prometheus.GaugeValue, jsonData.BitRateDownload)
// 		ch <- prometheus.MustNewConstMetric(c.bit_rate_upload, prometheus.GaugeValue, jsonData.BitRateUpload)
// 		ch <- prometheus.MustNewConstMetric(c.packet_rate_download, prometheus.GaugeValue, jsonData.PacketRateDownload)
// 		ch <- prometheus.MustNewConstMetric(c.packet_rate_upload, prometheus.GaugeValue, jsonData.PacketRateUpload)
// 		ch <- prometheus.MustNewConstMetric(c.loss_rate_download, prometheus.GaugeValue, jsonData.LossRateUpload)
// 		ch <- prometheus.MustNewConstMetric(c.loss_rate_upload, prometheus.GaugeValue, jsonData.LossRateDownload)
// 		ch <- prometheus.MustNewConstMetric(c.rtp_loss, prometheus.GaugeValue, jsonData.RTPLoss)
// 		ch <- prometheus.MustNewConstMetric(c.jitter_aggregate, prometheus.GaugeValue, jsonData.JitterAggregate)
// 		ch <- prometheus.MustNewConstMetric(c.rtt_aggregate, prometheus.GaugeValue, jsonData.RTTAggregate)
// 		ch <- prometheus.MustNewConstMetric(c.largest_conference, prometheus.GaugeValue, jsonData.LargestConference)
// 		ch <- prometheus.MustNewConstMetric(c.audiochannels, prometheus.GaugeValue, jsonData.Audiochannels)
// 		ch <- prometheus.MustNewConstMetric(c.videochannels, prometheus.GaugeValue, jsonData.Videochannels)
// 		ch <- prometheus.MustNewConstMetric(c.endpoints_sending_audio, prometheus.GaugeValue, jsonData.EndpointsSendingAudio)
// 		ch <- prometheus.MustNewConstMetric(c.endpoints_sending_video, prometheus.GaugeValue, jsonData.EndpointsSendingVideo)
// 		ch <- prometheus.MustNewConstMetric(c.conferences, prometheus.GaugeValue, jsonData.Conferences)
// 		ch <- prometheus.MustNewConstMetric(c.participants, prometheus.GaugeValue, jsonData.Participants)
// 		ch <- prometheus.MustNewConstMetric(c.videostreams, prometheus.GaugeValue, jsonData.Videostreams)
// 		ch <- prometheus.MustNewConstMetric(c.stress_level, prometheus.GaugeValue, jsonData.StressLevel)
// 		ch <- prometheus.MustNewConstMetric(c.total_loss_controlled_participant_seconds, prometheus.CounterValue, jsonData.TotalLossControlledParticipantSeconds)
// 		ch <- prometheus.MustNewConstMetric(c.total_loss_limited_participant_seconds, prometheus.CounterValue, jsonData.TotalLossLimitedParticipantSeconds)
// 		ch <- prometheus.MustNewConstMetric(c.total_loss_degraded_participant_seconds, prometheus.CounterValue, jsonData.TotalLossDegradedParticipantSeconds)
// 		ch <- prometheus.MustNewConstMetric(c.total_conference_seconds, prometheus.CounterValue, jsonData.TotalConferenceSeconds)
// 		ch <- prometheus.MustNewConstMetric(c.total_conferences_created, prometheus.CounterValue, jsonData.TotalConferencesCreated)
// 		ch <- prometheus.MustNewConstMetric(c.total_failed_conferences, prometheus.CounterValue, jsonData.TotalFailedConferences)
// 		ch <- prometheus.MustNewConstMetric(c.total_partially_failed_conferences, prometheus.CounterValue, jsonData.TotalPartiallyFailedConferences)
// 		ch <- prometheus.MustNewConstMetric(c.total_data_channel_messages_received, prometheus.CounterValue, jsonData.TotalDataChannelMessagesReceived)
// 		ch <- prometheus.MustNewConstMetric(c.total_data_channel_messages_sent, prometheus.CounterValue, jsonData.TotalDataChannelMessagesSent)
// 		ch <- prometheus.MustNewConstMetric(c.total_colibri_web_socket_messages_received, prometheus.CounterValue, jsonData.TotalColibriWebsocketMessagesReceived)
// 		ch <- prometheus.MustNewConstMetric(c.total_colibri_web_socket_messages_sent, prometheus.CounterValue, jsonData.TotalColibriWebsocketMessagesSent)
// 		ch <- prometheus.MustNewConstMetric(c.octo_send_bitrate, prometheus.GaugeValue, jsonData.OctoSendBitRate)
// 		ch <- prometheus.MustNewConstMetric(c.octo_receive_bitrate, prometheus.GaugeValue, jsonData.OctoReceiveBitRate)
// 		ch <- prometheus.MustNewConstMetric(c.octo_conferences, prometheus.GaugeValue, jsonData.OctoConferences)
// 		ch <- prometheus.MustNewConstMetric(c.octo_endpoints, prometheus.GaugeValue, jsonData.OctoEndpoints)
// 		ch <- prometheus.MustNewConstMetric(c.octo_send_packet_rate, prometheus.GaugeValue, jsonData.OctoSendPacketRate)
// 		ch <- prometheus.MustNewConstMetric(c.octo_receive_packet_rate, prometheus.GaugeValue, jsonData.OctoReceivePacketRate)

// 		var combinedConferenceSizes = make(map[float64]uint64)
// 		var combinedSum uint64
// 		conSizes, sum := conferenceSizesHelper(jsonData.ConferenceSizes)
// 		for bucket, numConferences := range conSizes {
// 			combinedConferenceSizes[bucket] += numConferences
// 		}
// 		combinedSum += sum

// 		ch <- prometheus.MustNewConstHistogram(c.conference_sizes, combinedSum, float64(combinedSum), combinedConferenceSizes)

// 	}
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
