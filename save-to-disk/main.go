package main

import (
	"fmt"

	"github.com/pion/webrtc-v3-api-examples/webrtc"
)

func main() {
	s := &webrtc.SettingEngine{
		// We can export this so users don't need to define VP8/VP9/H264/Opus/PCM Externally
		Codecs: []webrtc.RtpCodecCapability{
			{
				Name:      "VP8",
				ClockRate: 90000,
			},
		},
	}

	peerConnection, err := webrtc.NewAPI(s).NewPeerConnection(webrtc.Configuration{})
	if err != nil {
		panic(err)
	}

	if _, err = peerConnection.AddTransceiverFromKind(webrtc.RtpCodecTypeVideo); err != nil {
		panic(err)
	}

	peerConnection.OnTrack(func(t *webrtc.MediaStreamTrack, r *webrtc.RtpReceiver, s []*webrtc.MediaStream) {
		for {
			rtp, err := r.ReadRTP()
			if err != nil {
				panic(err)
			}

			// Save RTP packets to disk
			fmt.Println(rtp)
		}
	})

	// Signaling
}
