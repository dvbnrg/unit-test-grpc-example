package main

import (
	"context"
	"reflect"
	"testing"
	pb "unit-test-grpc-example/proto"
)

func Test_server_SayHello(t *testing.T) {
	type fields struct {
		UnimplementedGreeterServer pb.UnimplementedGreeterServer
	}
	type args struct {
		ctx context.Context
		in  *pb.HelloRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *pb.HelloReply
		wantErr bool
	}{
		// TODO: Add test cases.
		{"Say Hello Test", fields{}, args{context.Background(), &pb.HelloRequest{Name: "world"}}, &pb.HelloReply{Message: "Hello world"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &server{
				UnimplementedGreeterServer: tt.fields.UnimplementedGreeterServer,
			}
			got, err := s.SayHello(tt.args.ctx, tt.args.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("server.SayHello() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("server.SayHello() = %v, want %v", got, tt.want)
			}
		})
	}
}
