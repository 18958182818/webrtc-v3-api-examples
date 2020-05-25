This repository demonstrates all the use cases we want to support with the Pion WebRTC `/v3` Media API.

The root of the repository contains the stubs that demonstrate the new API, and each sub-directory demonstrates
actual usage of the API. This API is complete when we can demonstrate working usage of the following.

#### [Play Files from Disk](play-from-disk)
Example of playing file from disk.
* User declares codecs they are reading from disk, defined via [RTCRtpCodecCapability](https://draft.ortc.org/#dom-rtcrtpcodeccapability)
* On each negotiated PeerConnection they are notified if the receiver can accept what they are offering
* Read packets from disk and send until completed

#### [Save files to Disk](save-to-disk)
Example of saving user input to disk,
* User declares codecs they are saving to disk, defined via [RTCRtpCodecCapability](https://draft.ortc.org/#dom-rtcrtpcodeccapability)
* On each negotiated PeerConnection they are notified if the receiver can send what they are offering
* In OnTrack user Read packets from webrtc.RTPReceiver and save to disk

#### [Fan out WebRTC Input](fanout)
Example of simple SFU
* User declares codec they wish to fan out, defined via [RTCRtpCodecCapability](https://draft.ortc.org/#dom-rtcrtpcodeccapability)
* Uploader connects and we assert that our chosen codec is supported, if so we start reading from RTPReceiver
* Downloader connects and we assert our chosen codec is supported, if so we save our new RTPSender to a slice
* RTPReceiver on every read sends to every RTPSender

We can also implement some basic receiver feedback. These need to be supported, but are not mandatory
* Determine the lowest support bitrate across all receivers, and forward it back to the sender
* Forward NACK/PLI messages from each receiver back to the sender
* Simulcast (see below)

#### [Error Resilience Send](error-resilience-send)
A user sending video should be able to receive NACKs and respond to them

* User declares the codecs they are willing to send, defined via [RTCRtpCodecCapability](https://draft.ortc.org/#dom-rtcrtpcodeccapability). In `rtcpFeedback` they can declare they support NACK
* When generating SSRC we will check if they wish to support `NACK`, if they do we create another SSRC for that as well
* User can then call [GetSynchronizationSource](https://draft.ortc.org/#dom-rtcrtpreceiver-getsynchronizationsources), this will return a list of SSRCes known by this RTPSender.
* User can call WriteRTP with a packet they crafted themselves, they can determine the proper SSRC from the slice of SSRCes returned in `GetSynchronizationSources`
* User reads RTCP packets and when they recieve a NACK they can respond

#### [Error Resilience Receive](error-resilience-receive)
A user receiving video should be able to send NACKs and receive retransmissions

* User declares the codecs they are willing to receive, defined via [RTCRtpCodecCapability](https://draft.ortc.org/#dom-rtcrtpcodeccapability)
* In OnTrack user Read packets from webrtc.RTPReceiver and save to disk
* If user wants a certain packet to be re-sent they can call WriteRTCP on the PeerConnection with the Seqnum/SSRC of their choice

#### [Congestion Control Send](congestion-control-send)
A user sending video should be able to receive REMB/Receiver Reports and adjust bitrate

* User declares the codecs they are willing to send, in `rtcpFeedback` they can declare they support the congestion control algorithm of their choice
* Implementation details differ between Sender/Receiver side congestion control but users will Send/Receive RTP and RTCP packets like resilience methods

#### [Congestion Control Receive](congestion-control-receive)
A user receiving video should be able to send REMB/Receiver Reports

* User declares the codecs they are willing to send, defined via [RTCRtpCodecCapability](https://draft.ortc.org/#dom-rtcrtpcodeccapability). In `rtcpFeedback` they can declare they support the congestion control algorithm of their choice
* Implementation details differ between Sender/Receiver side congestion control but users will Send/Receive RTP and RTCP packets like resilience methods

#### [Simulcast Send](simulcast-send)
A user should be able to send multiple SSRCes for a single Track

* A user creates a track as the have done before
* When calling AddTransceiver user can now pass `SendEncodings`. This is how users in the browser send Simulcast [webrtc-pc#example-11](https://www.w3.org/TR/webrtc/#example-11)
```
pc.AddTransceiverFromTrack(track, {
  Direction: webrtc.RTPTransceiverDirectionSendrecv,
  SendEncodings: [
    {Rid: 'q'}
    {Rid: 'h'},
    {Rid: 'f'},
  ]
})
```

* User can then call [GetSynchronizationSource](https://draft.ortc.org/#dom-rtcrtpreceiver-getsynchronizationsources), this will return a list of SSRCes known by this RTPSender.
* User will be unable to use `WriteSample`, but will have to packetize and send themselves

#### [Simulcast Receive](simulcast-send)
A user should be able to receive SSRCes for a single Track

* OnTrack will be fired for the incoming simulcast track
* User can then call [GetSynchronizationSource](https://draft.ortc.org/#dom-rtcrtpreceiver-getsynchronizationsources), this will return a list of SSRCes known by this RTPSender.
* This will contain the details needed to get the details around each unique SSRC (and what rid they represent)

#### [Portable getUserMedia](portable-getusermedia)
Users should be able to call getUserMedia and have it work in both their Go and WASM code.

Everything should be behind platform flags in the `mediadevices` repo so user doesn't need to write platform specific
code.

------

#### Write/WriteRTP needs to allow users to send Simulcast and RTX, but we also want to autofill SSRC/PayloadType for the easy case
* Can have `RTPSenderWriteParamaters`. Users can specify `RID`/`isRetransmissions`
    - This might become ad-hoc and messy
* Can check the SSRC and conditionally rewrite.
    - Smaller API surface, but possibly more rough edges?
