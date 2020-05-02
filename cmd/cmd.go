package cmd

import (
	"github.com/golark/datagrabber/dgproto"
	log "github.com/sirupsen/logrus"
	gRPC "google.golang.org/grpc"
	"net"
	"strconv"
)

const (
	GrpcPort = ":8090"
)

type DataReq struct {

}

func (d DataReq) DataInquiry(req *dgproto.SearchReq, stream dgproto.DataService_DataInquiryServer) error {

	log.WithFields(log.Fields{"req:":req.String()}).Info("DataInquiry")

	for i:=0;i<10;i++ {
		resp := dgproto.DataHeaderResp{ColHeader:strconv.Itoa(i), RowHeader:strconv.Itoa(i+10)}
		stream.Send(&resp)
	}

	return nil
}


func ServeGrpc() {

	l, err := net.Listen("tcp", GrpcPort)

	if err!=nil {
		log.WithFields(log.Fields{"err":err}).Error("cant listen")
		return
	}

	s := gRPC.NewServer()

	dgproto.RegisterDataServiceServer(s, DataReq{})
	if err = s.Serve(l); err != nil {
		log.WithFields(log.Fields{"err":err}).Error("error serving")
	}

}
