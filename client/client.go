package main

import (
	"context"
	"github.com/golark/datagrabber/dgproto"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"io"
)

func exampleSingleShotGrpcClient() error {

	var opts []grpc.DialOption
	opts = append(opts, grpc.WithInsecure())

	conn, err := grpc.Dial("localhost:8090",opts...)
	if err != nil {
		log.WithFields(log.Fields{"err":err}).Error("cant dial grpc port")
		return err
	}
	defer conn.Close()

	client := dgproto.NewDataServiceClient(conn)

	r, err := client.DataInquiry(context.Background(), &dgproto.DataReq{Identifier:"car"})
	if err != nil {
		log.WithFields(log.Fields{"err":err}).Error("cant get data inquiry client")
		return err
	} else {
		for {
			header, err := r.Recv()
			if err == io.EOF {
				break
			} else {
				log.WithFields(log.Fields{"header:":header}).Info("received new data")
			}
		}
	}
	return nil

}

func main() {
	exampleSingleShotGrpcClient()
}
