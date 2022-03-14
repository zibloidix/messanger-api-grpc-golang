package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"time"

	pb "github.com/zibloidix/messanger-api-grpc-golang/messangerpb"
	"google.golang.org/grpc"
)

type server struct{}

func (*server) SendMessage(stream pb.MessangerService_SendMessageServer) error {
	var uuid string
	var user int32
	var chat int32
	var msg string

	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return stream.Send(&pb.SendMessageResponse{
				Message: &pb.Message{
					Uuid: "efdae4c2-7fd7-42e0-a7fd-a0d7331bc3cc",
					User: 100,
					Chat: 999,
					Msg:  "User chat is closed",
				},
			})
		}
		if err != nil {
			log.Fatalf("Reading stream error: %v", err)
		}
		uuid = req.GetMessage().GetUuid()
		user = req.GetMessage().GetUser()
		chat = req.GetMessage().GetChat()
		msg = req.GetMessage().GetMsg()

		stream.Send(&pb.SendMessageResponse{
			Message: &pb.Message{
				Uuid: "94146adc-6411-4be9-97d3-ecc2116a2ebf",
				User: 100,
				Chat: 999,
				Msg:  fmt.Sprintf("Original msg: {UUID:%s, USER: %d, CHAT: %d, MSG: %s}.", uuid, user, chat, msg),
			},
		})

		stream.Send(&pb.SendMessageResponse{
			Message: &pb.Message{
				Uuid: "03172e03-8fa1-4625-8ecf-b6997885d1f1",
				User: 100,
				Chat: 000,
				Msg:  fmt.Sprintf("Message from server (%v)!", time.Millisecond),
			},
		})
	}
	return nil
}

func main() {
	log.Println("Server start")
	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Fail to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterMessangerServiceServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Fail to listen: %v", err)
	}
}
