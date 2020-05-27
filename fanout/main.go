package main

import (
	"github.com/pion/webrtc-v3-api-examples/webrtc"
)

func main() {
	s := &webrtc.SettingEngine{
		Codecs: []webrtc.RtpCodecCapability{
			webrtc.RtpCodecCapabilityVP8(),
		},
	}

	peerConnection, err := webrtc.NewAPI(s).NewPeerConnection(webrtc.Configuration{})
	if err != nil {
		panic(err)
	}

	track := webrtc.NewMediaStreamTrack("fan-out-track", func(s webrtc.RtpSender, codecs []webrtc.RtpCodecCapability) {
		// Codecs will at least have one entry, SetRemoteDescription will fail otherwise
		// User can look at list and decide if they want to send H264/H265/VP8/VP9/AV1

		// SetCodec could fail if negotiation isn't complete, or codec selected isn't actually supported
		if err := s.SetCodec(codecs[0]); err != nil {
			panic(err)
		}
	})

	if _, err = peerConnection.AddTrack(track); err != nil {
		panic(err)
	}

	// Signaling
}
