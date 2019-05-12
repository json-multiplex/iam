package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"

	"github.com/golang/protobuf/ptypes"
	"github.com/golang/protobuf/ptypes/timestamp"
	"github.com/jmoiron/sqlx"
	pb "github.com/json-multiplex/iam/generated/v0"
	"github.com/json-multiplex/iam/internal/service"
	"github.com/json-multiplex/iam/internal/store"
	_ "github.com/lib/pq"
	"github.com/namsral/flag"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	fs := flag.NewFlagSetWithEnvPrefix(os.Args[0], "IAM_SERVICE", 0)
	db := fs.String("db", "", "database url")
	fs.Parse(os.Args[1:])

	dbConn, err := sqlx.Open("postgres", *db)
	if err != nil {
		log.Fatalf("could not open db connection: %v", err)
	}

	if err := dbConn.Ping(); err != nil {
		log.Fatalf("failed to ping db: %v", err)
	}

	srv := server{
		Service: &service.StoreService{
			Store: &store.DBStore{DB: dbConn},
		},
	}

	grpcServer := grpc.NewServer()
	pb.RegisterIAMServer(grpcServer, &srv)
	reflection.Register(grpcServer)

	l, err := net.Listen("tcp", "localhost:3000")
	if err != nil {
		log.Fatalf("could not listen for tcp: %v", err)
	}

	if err := grpcServer.Serve(l); err != nil {
		log.Fatalf("error serving: %v", err)
	}
}

type server struct {
	Service service.Service
}

func (s *server) CreateAccount(ctx context.Context, in *pb.CreateAccountRequest) (*pb.Account, error) {
	account, err := s.Service.CreateAccount(ctx, service.CreateAccountRequest{
		RootPassword: in.Account.RootPassword,
	})

	if err != nil {
		return nil, err
	}

	createTime, err := ptypes.TimestampProto(account.CreateTime)
	if err != nil {
		return nil, err
	}

	updateTime, err := ptypes.TimestampProto(account.UpdateTime)
	if err != nil {
		return nil, err
	}

	var deleteTime *timestamp.Timestamp
	if !account.DeleteTime.IsZero() {
		var err error
		deleteTime, err = ptypes.TimestampProto(account.DeleteTime)
		if err != nil {
			return nil, err
		}
	}

	return &pb.Account{
		Name:       fmt.Sprintf("accounts/%s", account.ID),
		CreateTime: createTime,
		UpdateTime: updateTime,
		DeleteTime: deleteTime,
	}, nil
}
