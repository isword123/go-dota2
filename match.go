package dota2

import (
	"fmt"
	"log"

	"github.com/Philipp15b/go-steam/dota/protocol/protobuf"
	"github.com/Philipp15b/go-steam/protocol/gamecoordinator"
)

// Sends a request to the Dota 2 GC requesting details for the given matchid.
func (c *Client) MatchDetails(matchID uint64) (*protobuf.CMsgGCMatchDetailsResponse, error) {
	if !c.gcReady {
		return nil, fmt.Errorf("GC not ready")
	}

	log.Printf("Requesting match details for match ID: %d", matchID)

	msgToGC := gamecoordinator.NewGCMsgProtobuf(
		AppId,
		uint32(protobuf.EDOTAGCMsg_k_EMsgGCMatchDetailsRequest),
		&protobuf.CMsgGCMatchDetailsRequest{
			MatchId: &matchID,
		})

	response := new(protobuf.CMsgGCMatchDetailsResponse)
	packet, err := c.runJob(msgToGC)
	if err != nil {
		return nil, err
	}
	packet.ReadProtoMsg(response) // Interpret GCPacket and populate `response` with data
	return response, nil
}
