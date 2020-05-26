package webrtc

type MediaStreamTrack struct{}
type MediaStream struct{}

func NewMediaStream(mediaStreamLabel string) *MediaStream { return nil }
func (m *MediaStream) AddTrack(t *MediaStreamTrack)       {}

func NewMediaStreamTrack(mediaStreamTrackLabel string, onNegotiated func(RtpSender, []RtpCodecCapability)) *MediaStreamTrack {
	return nil
}
