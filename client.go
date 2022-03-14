package main

import (
	"context"
	"fmt"
	"io"
	"log"

	pb "github.com/zibloidix/messanger-api-grpc-golang/messangerpb"
	"google.golang.org/grpc"
)

func main() {
	// 1. Делаем соединение для gRPC
	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
	defer cc.Close()

	// 2. Делаем новго клиента
	c := pb.NewMessangerServiceClient(cc)

	// 3. Получаем поток для чтения и записи
	stream, err := c.SendMessage(context.Background())
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
	// 4. Отправим сообщние на сервер
	stream.Send(&pb.SendMessageRequest{
		Message: &pb.Message{
			Uuid: "70182947-cb69-4943-94c7-96236637b7c6",
			User: 500,
			Chat: 300,
			Msg:  "Hello message 1 from client",
		},
	})

	stream.Send(&pb.SendMessageRequest{
		Message: &pb.Message{
			Uuid: "70182947-cb69-4943-94c7-96236637b7c6",
			User: 500,
			Chat: 300,
			Msg:  "Hello message 2 from client",
		},
	})

	stream.Send(&pb.SendMessageRequest{
		Message: &pb.Message{
			Uuid: "70182947-cb69-4943-94c7-96236637b7c6",
			User: 500,
			Chat: 300,
			Msg:  "Hello message 3 from client",
		},
	})

	stream.Send(&pb.SendMessageRequest{
		Message: &pb.Message{
			Uuid: "70182947-cb69-4943-94c7-96236637b7c6",
			User: 500,
			Chat: 300,
			Msg:  "Hello message 4 from client",
		},
	})

	// 5. Начианем чтение сообщений из потока от сервера
	for {
		log.Printf("Client step")

		resp, err := stream.Recv()
		if err == io.EOF {
			// Поток закончен
			stream.CloseSend()
		}
		if err != nil {
			// Ошибка чтения из потока
			log.Fatalf("Error: %v", err)
			break
		}
		uuid := resp.GetMessage().GetUuid()
		user := resp.GetMessage().GetUser()
		chat := resp.GetMessage().GetChat()
		msg := resp.GetMessage().GetMsg()
		logMsg := fmt.Sprintf("Message from SERVER: {UUID: %s, USER: %d, CHAT: %d, MSG: %s}", uuid, user, chat, msg)
		log.Println(logMsg)
	}
}
