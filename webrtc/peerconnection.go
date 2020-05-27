package webrtc

type RtpCodecType int

const (
	RtpCodecTypeAudio RtpCodecType = iota + 1
	RtpCodecTypeVideo
)

type PeerConnection struct{}
type Configuration struct{}
type RtpTransceiver struct{}
type RtpReceiver struct{}
type RtpSender struct{}
type Packet struct{}

func (api *API) NewPeerConnection(c Configuration) (*PeerConnection, error) {
	return nil, nil
}

// PeerConnection APIs
func (pc *PeerConnection) AddTransceiverFromKind(r RtpCodecType) (*RtpTransceiver, error) {
	return nil, nil
}

func (pc *PeerConnection) OnTrack(cb func(*MediaStreamTrack, *RtpReceiver, []*MediaStream)) {}

// RtpTransceiver APIs
func (t *RtpTransceiver) setCodecPreferences([]RtpCodecCapability) error { return nil }

// RtpReceiver APIs
func (r *RtpReceiver) ReadRTP() ([]byte, error) {
	return nil, nil
}
