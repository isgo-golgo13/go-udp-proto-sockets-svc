package proto_svc

import (
	"log"

	"github.com/isgo-golgo13/udp_proto_svc/proto_data"
)

func CreatePacket(id int32, data string) *proto_data.Packet {
	packet := proto_data.Packet{
		Id:      id,
		Payload: data,
	}
	return &packet
}

func CheckError(err error) {
	if err != nil {
		log.Fatal("Error: ", err)
	}
}
