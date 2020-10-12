package main

import (
	"flag"
	"net"
	"os"
	"time"

	"github.com/isgo-golgo13/udp_proto_svc/proto_svc"

	"github.com/golang/protobuf/proto"
	log "github.com/sirupsen/logrus"
)

var (
	server     = flag.String("s", "127.0.0.1", "ip")
	serverPort = flag.String("p", "7000", "host: ip:port")
)

func init() {
	log.SetOutput(os.Stdout)
	log.SetLevel(log.InfoLevel)
}

func main() {
	log.WithFields(log.Fields{
		"time": time.Now().String(),
	}).Info("Client Start Time")

	flag.Parse()
	StartClient(*server, *serverPort)
}

func StartClient(server, serverPort string) {
	remoteAddr, err := net.ResolveUDPAddr("udp", server+":"+serverPort)
	proto_svc.CheckError(err)

	localAddr, err := net.ResolveUDPAddr("udp", server+":0")
	proto_svc.CheckError(err)

	conn, err := net.DialUDP("udp", localAddr, remoteAddr)
	proto_svc.CheckError(err)

	defer conn.Close()

	i := 1
	for {
		packet := proto_svc.CreatePacket(int32(i), "Data Payload")
		now := time.Now().Unix()
		packet.PayloadSentTime = now
		data, err := proto.Marshal(packet)
		if err != nil {
			log.Fatal("marshalling error: ", err)
		}
		buf := []byte(data)
		_, err = conn.Write(buf)
		if err != nil {
			log.Println(err)
		}

		i++
		time.Sleep(time.Second * 1)
	}

}
