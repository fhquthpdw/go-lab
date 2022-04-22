package main

import (
	"context"
	"flag"
	"fmt"
	pb "grpc-stat-server/pb"
	"io"
	"log"
	"net"
	"time"

	"google.golang.org/grpc/reflection"

	"google.golang.org/grpc"
)

func main() {
	flag.Parse()
	s := fmt.Sprintf(":10000")
	lis, err := net.Listen("tcp", s)

	if err != nil {
		log.Fatalf("Listen server error: %v", err)
	}

	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
	pb.RegisterStatServiceServer(grpcServer, newServer())
	reflection.Register(grpcServer)
	log.Println("Server on: ", s)
	if err := grpcServer.Serve(lis); err == nil {
	}
	log.Fatalf("Start server error: %v", err)
}
func newServer() *statServiceServer {
	return &statServiceServer{}
}

type statServiceServer struct {
	//
}

func (s *statServiceServer) Add(ctx context.Context, stat *pb.Stat) (*pb.AddReturn, error) {
	fmt.Println("Add:")
	fmt.Println("")

	result := pb.AddReturn{
		Result: true,
	}
	return &result, nil
}
func (s *statServiceServer) BatchAdd(stream pb.StatService_BatchAddServer) error {
	fmt.Println("Batch Add:")
	fmt.Println("")

	for {
		stat, err := stream.Recv()
		if err == io.EOF {
			break
		}
		stat.Action += "-received"
		stream.Send(stat)
	}
	return nil
}
func (s *statServiceServer) BatchAddWithSummary(stream pb.StatService_BatchAddWithSummaryServer) error {
	fmt.Println("Batch Add with Summary:")
	fmt.Println("")

	var total, success, failed int32
	failed = 1
	for {
		stat, err := stream.Recv()
		if err == io.EOF {
			summary := &pb.BatchAddSummary{
				Total:   total,
				Success: success,
				Failed:  failed,
			}
			return stream.SendAndClose(summary)
		}
		if stat != nil {
			fmt.Printf("%+v", stat)
			total++
			success++
		}
	}
}
func (s *statServiceServer) Search(ctx context.Context, parameters *pb.SearchParameters) (*pb.SearchResult, error) {
	fmt.Println("Search:")
	fmt.Println(parameters)
	fmt.Println("")

	searchResult := &pb.SearchResult{
		Total: 1989,
		Data: &pb.StatList{
			StatArray: []*pb.Stat{
				&pb.Stat{Platform: 0, Account: "daochun.zhao@gmail.com", Path: "/account/login", Action: "click", Time: time.Now().Unix()},
				&pb.Stat{Platform: 1, Account: "daochun.zhao@iherb.com", Path: "/account/dashboard", Action: "click", Time: time.Now().Unix()},
				&pb.Stat{Platform: 2, Account: "daochun.zhao@mail.com", Path: "/account/money", Action: "click", Time: time.Now().Unix()},
				&pb.Stat{Platform: 3, Account: "daochun.zhao@126.com", Path: "/account/title", Action: "click", Time: time.Now().Unix()},
			},
		},
	}

	return searchResult, nil
}
func (s *statServiceServer) ListSearchData(parameters *pb.SearchParameters, stream pb.StatService_ListSearchDataServer) error {
	fmt.Println("List Search Data:")
	fmt.Println(parameters)
	fmt.Println("")

	statSlice := []*pb.Stat{
		&pb.Stat{Platform: 0, Account: "daochun.zhao@gmail.com", Path: "/account/login", Action: "click", Time: time.Now().Unix()},
		&pb.Stat{Platform: 1, Account: "daochun.zhao@iherb.com", Path: "/account/dashboard", Action: "click", Time: time.Now().Unix()},
		&pb.Stat{Platform: 2, Account: "daochun.zhao@mail.com", Path: "/account/money", Action: "click", Time: time.Now().Unix()},
		&pb.Stat{Platform: 3, Account: "daochun.zhao@126.com", Path: "/account/title", Action: "click", Time: time.Now().Unix()},
	}

	for _, stat := range statSlice {
		stream.Send(stat)
	}

	return nil
}
