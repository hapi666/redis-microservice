package main

import (
	"context"
	"log"
	"net"
	"redis-microservice/config"
	proto "redis-microservice/pb"
	"redis-microservice/redis"

	"github.com/golang/protobuf/ptypes/empty"

	"google.golang.org/grpc"
)

var (
	filePath = "conf"
	typeName = "yaml"
	pool     *redis.Pool
)

type DBServer struct{}

func (dbs *DBServer) PushQueue(ctx context.Context, in *proto.PushQueueRequest) (*empty.Empty, error) {
	return &empty.Empty{}, pool.PushQueue(in.Urls)
}

func (dbs *DBServer) PopQueue(ctx context.Context, in *empty.Empty) (*proto.PopQueueReply, error) {
	url, err := pool.PopQueue()
	return &proto.PopQueueReply{Url: url}, err
}

func (dbs *DBServer) RangeQueue(ctx context.Context, in *empty.Empty) (*proto.RangeQueueReply, error) {
	urls, err := pool.RangeQueue()
	return &proto.RangeQueueReply{Urls: urls}, err
}

func (dbs *DBServer) SisMember(ctx context.Context, in *proto.SisMemberRequest) (*proto.SisMemberReply, error) {
	isExist, err := pool.SisMember(in.Url)
	return &proto.SisMemberReply{IsExist: isExist}, err
}

func (dbs *DBServer) SadD(ctx context.Context, in *proto.SadDRequest) (*empty.Empty, error) {
	return &empty.Empty{}, pool.SadD(in.CrawledURL)
}

func main() {
	cfg, err := config.NewConfig(filePath, typeName)
	if err != nil {
		log.Fatalf("Failed to new a cfg, err: %s", err)
	}
	lis, err := net.Listen("tcp", ":"+cfg.GetString("network.port"))
	if err != nil {
		panic(err)
	}
	ser := grpc.NewServer()
	pool, err = redis.NewRedisPool(":"+cfg.GetString("redis.port"), "")
	if err != nil {
		log.Fatal(err)
	}
	proto.RegisterDBServiceServer(ser, &DBServer{})
	err = ser.Serve(lis)
	if err != nil {
		log.Fatal(err)
	}
}
