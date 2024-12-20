package handlers

import (
	"log"
	"runtime"

	"github.com/MikeDev101/cloudlink-phi/server/pkg/constants"
	"github.com/MikeDev101/cloudlink-phi/server/pkg/signaling/message"
	"github.com/MikeDev101/cloudlink-phi/server/pkg/structs"
)

type MetadataPacket struct {
	OperatingSystem string `json:"os"`
	Architecture    string `json:"architecture"`
	ServerVersion   string `json:"version"`
	GoVersion       string `json:"go_version"`
}

func META(s *structs.Server, client *structs.Client, packet *structs.SignalPacket) {
	err := message.Code(
		client,
		"ACK_META",
		&MetadataPacket{ // Read system information from the OS
			OperatingSystem: runtime.GOOS,
			Architecture:    runtime.GOARCH,
			GoVersion:       runtime.Version(),
			ServerVersion:   constants.Version,
		},
		packet.Listener,
		nil,
	)
	if err != nil {
		log.Printf("Send ACK_META response to META opcode error: %s", err.Error())
	}
}
