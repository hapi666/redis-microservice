package main

import (
	"context"
	"net"
	proto "redis-microservice/pb"
	"redis-microservice/redis"

	"github.com/golang/protobuf/ptypes/empty"

	"google.golang.org/grpc"
)

var db redis.DB

type DBServer struct{}

func (dbs *DBServer) PushQueue(ctx context.Context, in *proto.PushQueueRequest) (*empty.Empty, error) {
	return nil, db.PushQueue(in.Urls)
}

func (dbs *DBServer) PopQueue(ctx context.Context, in *empty.Empty) (*proto.PopQueueReply, error) {
	url, err := db.PopQueue()
	return &proto.PopQueueReply{Url: url}, err
}

func (dbs *DBServer) RangeQueue(ctx context.Context, in *empty.Empty) (*proto.RangeQueueReply, error) {
	urls, err := db.RangeQueue()
	return &proto.RangeQueueReply{Urls: urls}, err
}

func (dbs *DBServer) SisMember(ctx context.Context, in *proto.SisMemberRequest) (*proto.SisMemberReply, error) {
	isExist, err := db.SisMember(in.Url)
	return &proto.SisMemberReply{IsExist: isExist}, err
}

func (dbs *DBServer) SadD(ctx context.Context, in *proto.SadDRequest) (*empty.Empty, error) {
	return nil, db.SadD(in.CrawledURL)
}

func main() {
	lis, err := net.Listen("", "")
	if err != nil {
		panic(err)
	}
	ser := grpc.NewServer()
	DBPool, err := redis.NewRedisPool("", "")
	if err != nil {
		panic(err)
	}
	db, err = DBPool.GetRedisConn()
	if err != nil {
		panic(err)
	}
	proto.RegisterDBServiceServer(ser, &DBServer{})
	err = ser.Serve(lis)
	if err != nil {
		panic(err)
	}
}
