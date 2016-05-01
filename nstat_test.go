package nstat

import (
	"testing"
)

var snmpFile = []byte(`Ip: Forwarding DefaultTTL InReceives InHdrErrors InAddrErrors ForwDatagrams InUnknownProtos InDiscards InDelivers OutRequests OutDiscards OutNoRoutes ReasmTimeout ReasmReqds ReasmOKs ReasmFails FragOKs FragFails FragCreates
Ip: 2 64 1231 0 0 0 0 0 1227 972 0 0 0 0 0 0 0 0 0
Icmp: InMsgs InErrors InCsumErrors InDestUnreachs InTimeExcds InParmProbs InSrcQuenchs InRedirects InEchos InEchoReps InTimestamps InTimestampReps InAddrMasks InAddrMaskReps OutMsgs OutErrors OutDestUnreachs OutTimeExcds OutParmProbs OutSrcQuenchs OutRedirects OutEchos OutEchoReps OutTimestamps OutTimestampReps OutAddrMasks OutAddrMaskReps
Icmp: 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0
Tcp: RtoAlgorithm RtoMin RtoMax MaxConn ActiveOpens PassiveOpens AttemptFails EstabResets CurrEstab InSegs OutSegs RetransSegs InErrs OutRsts InCsumErrors
Tcp: 1 200 120000 -1 2 4 0 0 1 796 541 0 0 1 0
Udp: InDatagrams NoPorts InErrors OutDatagrams RcvbufErrors SndbufErrors InCsumErrors
Udp: 431 0 0 431 0 0 0
UdpLite: InDatagrams NoPorts InErrors OutDatagrams RcvbufErrors SndbufErrors InCsumErrors
UdpLite: 0 0 0 0 0 0 0
`)

var netstatFile = []byte(`TcpExt: SyncookiesSent SyncookiesRecv SyncookiesFailed EmbryonicRsts PruneCalled RcvPruned OfoPruned OutOfWindowIcmps LockDroppedIcmps ArpFilter TW TWRecycled TWKilled PAWSPassive PAWSActive PAWSEstab DelayedACKs DelayedACKLocked DelayedACKLost ListenOverflows ListenDrops TCPPrequeued TCPDirectCopyFromBacklog TCPDirectCopyFromPrequeue TCPPrequeueDropped TCPHPHits TCPHPHitsToUser TCPPureAcks TCPHPAcks TCPRenoRecovery TCPSackRecovery TCPSACKReneging TCPFACKReorder TCPSACKReorder TCPRenoReorder TCPTSReorder TCPFullUndo TCPPartialUndo TCPDSACKUndo TCPLossUndo TCPLostRetransmit TCPRenoFailures TCPSackFailures TCPLossFailures TCPFastRetrans TCPForwardRetrans TCPSlowStartRetrans TCPTimeouts TCPLossProbes TCPLossProbeRecovery TCPRenoRecoveryFail TCPSackRecoveryFail TCPSchedulerFailed TCPRcvCollapsed TCPDSACKOldSent TCPDSACKOfoSent TCPDSACKRecv TCPDSACKOfoRecv TCPAbortOnData TCPAbortOnClose TCPAbortOnMemory TCPAbortOnTimeout TCPAbortOnLinger TCPAbortFailed TCPMemoryPressures TCPSACKDiscard TCPDSACKIgnoredOld TCPDSACKIgnoredNoUndo TCPSpuriousRTOs TCPMD5NotFound TCPMD5Unexpected TCPSackShifted TCPSackMerged TCPSackShiftFallback TCPBacklogDrop TCPMinTTLDrop TCPDeferAcceptDrop IPReversePathFilter TCPTimeWaitOverflow TCPReqQFullDoCookies TCPReqQFullDrop TCPRetransFail TCPRcvCoalesce TCPOFOQueue TCPOFODrop TCPOFOMerge TCPChallengeACK TCPSYNChallenge TCPFastOpenActive TCPFastOpenPassive TCPFastOpenPassiveFail TCPFastOpenListenOverflow TCPFastOpenCookieReqd TCPSpuriousRtxHostQueues BusyPollRxPackets
TcpExt: 0 0 0 0 0 0 0 0 0 0 1 0 0 0 0 0 5 0 0 0 0 3 0 0 0 221 0 10 517 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 2 0 0 0 0 0 0 0 0 0 0 0 0
IpExt: InNoRoutes InTruncatedPkts InMcastPkts OutMcastPkts InBcastPkts OutBcastPkts InOctets OutOctets InMcastOctets OutMcastOctets InBcastOctets OutBcastOctets InCsumErrors InNoECTPkts InECT1Pkts InECT0Pkts InCEPkts
IpExt: 4 0 0 0 0 0 127293 89075 0 0 0 0 0 1269 0 0 0
`)

var snmp6File = []byte(`Ip6InReceives                   	0
Ip6InHdrErrors                  	0
Ip6InTooBigErrors               	0
Ip6InNoRoutes                   	0
Ip6InAddrErrors                 	0
Ip6InUnknownProtos              	0
Ip6InTruncatedPkts              	0
Ip6InDiscards                   	0
Ip6InDelivers                   	0
Ip6OutForwDatagrams             	0
Ip6OutRequests                  	14
Ip6OutDiscards                  	0
Ip6OutNoRoutes                  	5
Ip6ReasmTimeout                 	0
Ip6ReasmReqds                   	0
Ip6ReasmOKs                     	0
Ip6ReasmFails                   	0
Ip6FragOKs                      	0
Ip6FragFails                    	0
Ip6FragCreates                  	0
Ip6InMcastPkts                  	0
Ip6OutMcastPkts                 	22
Ip6InOctets                     	0
Ip6OutOctets                    	1068
Ip6InMcastOctets                	0
Ip6OutMcastOctets               	1776
Ip6InBcastOctets                	0
Ip6OutBcastOctets               	0
Ip6InNoECTPkts                  	0
Ip6InECT1Pkts                   	0
Ip6InECT0Pkts                   	0
Ip6InCEPkts                     	0
Icmp6InMsgs                     	0
Icmp6InErrors                   	0
Icmp6OutMsgs                    	14
Icmp6OutErrors                  	0
Icmp6InCsumErrors               	0
Icmp6InDestUnreachs             	0
Icmp6InPktTooBigs               	0
Icmp6InTimeExcds                	0
Icmp6InParmProblems             	0
Icmp6InEchos                    	0
Icmp6InEchoReplies              	0
Icmp6InGroupMembQueries         	0
Icmp6InGroupMembResponses       	0
Icmp6InGroupMembReductions      	0
Icmp6InRouterSolicits           	0
Icmp6InRouterAdvertisements     	0
Icmp6InNeighborSolicits         	0
Icmp6InNeighborAdvertisements   	0
Icmp6InRedirects                	0
Icmp6InMLDv2Reports             	0
Icmp6OutDestUnreachs            	0
Icmp6OutPktTooBigs              	0
Icmp6OutTimeExcds               	0
Icmp6OutParmProblems            	0
Icmp6OutEchos                   	0
Icmp6OutEchoReplies             	0
Icmp6OutGroupMembQueries        	0
Icmp6OutGroupMembResponses      	0
Icmp6OutGroupMembReductions     	0
Icmp6OutRouterSolicits          	3
Icmp6OutRouterAdvertisements    	0
Icmp6OutNeighborSolicits        	3
Icmp6OutNeighborAdvertisements  	0
Icmp6OutRedirects               	0
Icmp6OutMLDv2Reports            	8
Icmp6OutType133                 	3
Icmp6OutType135                 	3
Icmp6OutType143                 	8
Udp6InDatagrams                 	0
Udp6NoPorts                     	0
Udp6InErrors                    	0
Udp6OutDatagrams                	0
Udp6RcvbufErrors                	0
Udp6SndbufErrors                	0
Udp6InCsumErrors                	0
UdpLite6InDatagrams             	0
UdpLite6NoPorts                 	0
UdpLite6InErrors                	0
UdpLite6OutDatagrams            	0
UdpLite6RcvbufErrors            	0
UdpLite6SndbufErrors            	0
UdpLite6InCsumErrors            	0
`)

var result map[string]int64

func BenchmarkNstatGet(b *testing.B) {
	counters := map[string]int64{}
	c := &Counters{}
	for n := 0; n < b.N; n++ {
		counters = c.Get()
	}
	result = counters
}

func BenchmarkNstatGetDumpZeros(b *testing.B) {
	counters := map[string]int64{}
	c := &Counters{
		DumpZeros: true,
	}
	for n := 0; n < b.N; n++ {
		counters = c.Get()
	}
	result = counters
}

func BenchmarkParseSNMPFile(b *testing.B) {
	counters := map[string]int64{}
	c := &Counters{}
	c.Get()
	for n := 0; n < b.N; n++ {
		c.parseUglyFile(snmpFile)
	}
	result = counters
}

func BenchmarkParseNetstatFile(b *testing.B) {
	counters := map[string]int64{}
	c := &Counters{}
	c.Get()
	for n := 0; n < b.N; n++ {
		c.parseUglyFile(netstatFile)
	}
	result = counters
}

func BenchmarkParseSNMP6File(b *testing.B) {
	counters := map[string]int64{}
	c := &Counters{}
	c.Get()
	for n := 0; n < b.N; n++ {
		c.parseNiceFile(snmp6File)
	}
	result = counters
}
