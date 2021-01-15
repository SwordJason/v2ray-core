package congestion

import "github.com/SwordJason/v2ray-core/external/github.com/lucas-clemente/quic-go/internal/protocol"

type connectionStats struct {
	slowstartPacketsLost protocol.PacketNumber
	slowstartBytesLost   protocol.ByteCount
}
