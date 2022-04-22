package main

import (
	"context"
	"flag"
	"fmt"
	pb "grpc-stat-client/pb"
	"io"
	"log"
	"time"

	"google.golang.org/grpc"
)

var (
	serverAddr = flag.String("server_addr", ":10000", "The server address in the format of host:port")
)

func main() {
	flag.Parse()
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithInsecure())
	opts = append(opts, grpc.WithBlock())

	conn, err := grpc.Dial(":10000", opts...)
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}
	defer conn.Close()
	client := pb.NewStatServiceClient(conn)

	/***** Add *****/
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	item := pb.Stat{
		Platform: 0,
		Account:  "daochun.zhao@gmail.com",
		Path:     "/account/login",
		Action:   "click",
		Time:     time.Now().Unix(),
	}
	result, err := client.Add(ctx, &item)
	if err != nil {
		log.Fatalf("%v.Add, error: %v", client, err)
	}

	fmt.Println("Add:")
	fmt.Printf("%+v\n", result)
	fmt.Println("")
	/***** Add *****/

	/***** BatchAdd *****/
	ctx, cancel = context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	statList := &pb.StatList{
		StatArray: []*pb.Stat{
			&pb.Stat{Platform: 0, Account: "daochun.zhao@gmail.com", Path: "/account/login", Action: "click", Time: 123456789},
			&pb.Stat{Platform: 1, Account: "daochun.zhao@iherb.com", Path: "/account/dashboard", Action: "click", Time: 223456789},
			&pb.Stat{Platform: 2, Account: "daochun.zhao@mail.com", Path: "/account/money", Action: "click", Time: 323456789},
			&pb.Stat{Platform: 3, Account: "daochun.zhao@126.com", Path: "/account/title", Action: "click", Time: 423456789},
		},
	}
	stream, err := client.BatchAdd(ctx)

	wait := make(chan struct{})
	fmt.Println("Batch Add:")

	// receive
	go func() {
		for {
			stat, err := stream.Recv()
			if err == io.EOF {
				close(wait)
				return
			}
			if err != nil {
				log.Fatalf("failed to receive a not: %v", err)
			}
			if stat != nil {
				fmt.Printf("%+v\n", stat)
			}
		}
	}()

	// send
	for _, stat := range statList.StatArray {
		if err := stream.Send(stat); err != nil {
			log.Fatalf("failed to receive a note: %v", err)
		}
	}
	stream.CloseSend()

	// wait receiving data
	<-wait
	fmt.Println("")
	/***** BatchAdd *****/

	/***** BatchAddWithSummary *****/
	ctx, cancel = context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	statList2 := &pb.StatList{
		StatArray: []*pb.Stat{
			&pb.Stat{Platform: 0, Account: "daochun.zhao@gmail.com", Path: "/account/login", Action: "click", Time: 123456789},
			&pb.Stat{Platform: 1, Account: "daochun.zhao@iherb.com", Path: "/account/dashboard", Action: "click", Time: 223456789},
			&pb.Stat{Platform: 2, Account: "daochun.zhao@mail.com", Path: "/account/money", Action: "click", Time: 323456789},
			&pb.Stat{Platform: 3, Account: "daochun.zhao@126.com", Path: "/account/title", Action: "click", Time: 423456789},
		},
	}
	stream2, err := client.BatchAddWithSummary(ctx)

	// send
	for _, stat := range statList2.StatArray {
		if err := stream2.Send(stat); err != nil {
			log.Fatalf("failed to receive a note: %v", err)
		}
	}
	// receive
	summary, _ := stream2.CloseAndRecv()

	fmt.Println("Batch Add with Summary:")
	fmt.Printf("%+v", summary)
	fmt.Println("")
	fmt.Println("")
	/***** BatchAddWithSummary *****/

	/***** Search *****/
	ctx, cancel = context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	searchParameter := pb.SearchParameters{
		Keyword:  "daochun",
		Page:     1,
		PageSize: 10,
	}
	list, err := client.Search(ctx, &searchParameter)

	if err != nil {
		log.Fatalf("%v.List, error: %v", client, err)
	}
	fmt.Println("Search:")
	fmt.Printf("%+v\n", list)
	fmt.Println("")
	/***** Search *****/

	/***** List Search Data *****/
	ctx, cancel = context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	searchParameter2 := pb.SearchParameters{
		Keyword:  "daochun",
		Page:     1,
		PageSize: 10,
	}
	stream3, err := client.ListSearchData(ctx, &searchParameter2)

	if err != nil {
		log.Fatalf("%v.List, error: %v", client, err)
	}

	wait3 := make(chan struct{})
	fmt.Println("List Search Data:")

	// receive
	go func() {
		for {
			stat, err := stream3.Recv()
			if err == io.EOF {
				close(wait3)
				return
			}
			if err != nil {
				log.Fatalf("failed to receive a not: %v", err)
			}
			if stat != nil {
				fmt.Printf("%+v\n", stat)
			}
		}
	}()

	// waiting for receiving all
	<-wait3
	/***** List Search Data *****/
}
