package main

import (
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

	track := webrtc.NewMediaStreamTrack("my-video-track", func(s webrtc.RtpSender, codecs []webrtc.RtpCodecCapability) {
		// Codecs will at least have one entry, SetRemoteDescription will fail otherwise
		// User can look at list and decide if they want to send H264/H265/VP8/VP9/AV1

		// SetCodec could fail if negotiation isn't complete, or codec selected isn't actually supported
		if err := s.SetCodec(codecs[0]); err != nil {
			panic(err)
		}

		// Read packets from disk and call s.WriteSample
	})

	if _, err = peerConnection.AddTrack(track); err != nil {
		panic(err)
	}

	// Signaling
}
