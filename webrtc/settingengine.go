package webrtc

type API struct {
}

func NewAPI(s *SettingEngine) *API {
	return nil
}

type RtpCodecCapability struct {
	Name      string
	ClockRate int
}

type SettingEngine struct {
	Codecs []RtpCodecCapability
}
