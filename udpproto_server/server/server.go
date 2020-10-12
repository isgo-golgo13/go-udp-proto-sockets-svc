package main

import (
	"flag"
	"net"
	"os"
	"time"

	"github.com/isgo-golgo13/udp_proto_svc/proto_data"
	"github.com/isgo-golgo13/udp_proto_svc/proto_svc"

	"github.com/golang/protobuf/proto"
	log "github.com/sirupsen/logrus"
)

var (
	port = flag.String("p", "7000", "host: ip:port")
)

func init() {
	log.SetOutput(os.Stdout)
	log.SetLevel(log.InfoLevel)
}

func main() {
	log.WithFields(log.Fields{
		"time": time.Now().String(),
	}).Info("Server Start Time")

	flag.Parse()
	StartServer()

}

func StartServer() {
	serverAddr, err := net.ResolveUDPAddr("udp", ":"+*port)
	proto_svc.CheckError(err)

	serverConn, err := net.ListenUDP("udp", serverAddr)
	proto_svc.CheckError(err)
	defer serverConn.Close()

	buf := make([]byte, 1024)

	log.Println("Listening on port " + *port)
	for {
		n, addr, err := serverConn.ReadFromUDP(buf)
		packet := &proto_data.Packet{}
		err = proto.Unmarshal(buf[0:n], packet)
		log.Printf("Received %d sent at %s from %s as payload data [%s]", packet.Id, time.Unix(packet.PayloadSentTime, 0), addr, packet.Payload)

		if err != nil {
			log.Fatal("Error: ", err)
		}
	}
}
