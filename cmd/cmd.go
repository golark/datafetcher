package cmd

import (
	"github.com/golark/datagrabber/dgproto"
	"github.com/golark/datagrabber/symphoniser"
	log "github.com/sirupsen/logrus"
	gRPC "google.golang.org/grpc"
	"net"
)

const (
	GrpcPort = ":8090"
)

type DataReq struct {}

func (d DataReq) DataInquiry(req *dgproto.SearchReq, stream dgproto.DataService_DataInquiryServer) error {

	log.WithFields(log.Fields{"request identifier":req.Identifier}).Info("new data inquiry")

	// step 1 - request data inquiry
	rowHeaders, colHeaders := symphoniser.GetDataHeaders("covid") // @TODO add request identifier

	// step 2 - send the stream
	for _, row := range rowHeaders {
		for _, col := range colHeaders {
			resp := dgproto.DataHeaderResp{ColHeader: col, RowHeader: row}
			stream.Send(&resp)
		}
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

