package main

import (
	"context"
	"flag"
	"github.com/iamgadfly/go-echo-api/pkg/api/user_v1/proto"
	"google.golang.org/grpc"
	"log"
	"strconv"
)

func main() {
	flag.Parse()
	conn, err := grpc.Dial(":8888", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	c := proto.NewUserClient(conn)
	ctx := context.Background()

	if flag.NArg() == 1 {
		userId, err := strconv.Atoi(flag.Arg(0))
		if err != nil {
			log.Fatal(err)
		}
		res, err := c.Get(ctx, &proto.UserIdRequest{UserId: uint64(userId)})
		if err != nil {
			log.Fatal(err)
		}
		log.Println(res.GetData())
	} else {
		res, err := c.GetUsers(ctx, &proto.EmptyParams{})
		if err != nil {
			log.Fatal(err)
		}
		log.Println(res.GetData())
	}
}
