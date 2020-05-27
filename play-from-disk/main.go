package main

import (
	"github.com/pion/webrtc-v3-api-examples/webrtc"
)

func main() {
	// User's can configure codec preference per PeerConnection
	// We will also have a default with everything enabled that we support
	s := &webrtc.SettingEngine{
		Codecs: []webrtc.RtpCodecCapability{
			webrtc.RtpCodecCapabilityVP8(),
		},
	}

	peerConnection, err := webrtc.NewAPI(s).NewPeerConnection(webrtc.Configuration{})
	if err != nil {
		panic(err)
	}

	// track := webrtc.NewMediaStreamTrack("my-video-track", func(s webrtc.RtpSender, codecs []webrtc.RtpCodecCapability) {
	// 	// Codecs will at least have one entry, SetRemoteDescription will fail otherwise
	// 	// User can look at list and decide if they want to send H264/H265/VP8/VP9/AV1

	// 	// SetCodec could fail if negotiation isn't complete, or codec selected isn't actually supported
	// 	if err := s.SetCodec(codecs[0]); err != nil {
	// 		panic(err)
	// 	}

	// 	// Read packets from disk and call s.WriteSample
	// })

	// if _, err = peerConnection.AddTrack(track); err != nil {
	// 	panic(err)
	// }

	// Signaling
}
