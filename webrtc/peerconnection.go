package webrtc

type RTPCodecType int

const (
	RTPCodecTypeAudio RTPCodecType = iota + 1
	RTPCodecTypeVideo
)

type PeerConnection struct{}
type Configuration struct{}
type RTPTransceiver struct{}
type RTPReceiver struct{}
type Packet struct{}

func (api *API) NewPeerConnection(c Configuration) (*PeerConnection, error) {
	return nil, nil
}

func (pc *PeerConnection) AddTransceiverFromKind(r RTPCodecType) (*RTPTransceiver, error) {
	return nil, nil
}

func (pc *PeerConnection) OnTrack(cb func(*MediaStreamTrack, *RTPReceiver)) {}

func (r *RTPReceiver) ReadRTP() ([]byte, error) {
	return nil, nil
}
