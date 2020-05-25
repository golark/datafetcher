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

// DataInquiry
// serve incoming data requests by sending a stream of points
func (d DataReq) DataInquiry(req *dgproto.DataReq, stream dgproto.DataService_DataInquiryServer) error {

	log.WithFields(log.Fields{"request identifier":req.Identifier}).Info("received DataInquiry")

	// step 1 - request data
	points, err := symphoniser.DataInquiry(req.Identifier)
	if err != nil {
		return err
	}

	// step 2 - stream
	for i := range points {
		stream.Send(&points[i])
	}

	return nil
}

// HeaderInquiry
// send back a stream of headers
func (d DataReq) HeaderInquiry(req *dgproto.HeaderReq, stream dgproto.DataService_HeaderInquiryServer) error {

	log.WithFields(log.Fields{"request identifier":req.Identifier}).Info("received HeaderInquiry")

	// step 1 - request headers
	rowHeaders, colHeaders := symphoniser.GetDataHeaders("covid") // @TODO add request identifier

	// step 2 - send the stream
	for _, row := range rowHeaders {
		for _, col := range colHeaders {
			resp := dgproto.HeaderResp{ColHeader: col, RowHeader: row}
			stream.Send(&resp)
		}
	}

	return nil
}

// ServeGrpc
// entry point cmd grpc server
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

