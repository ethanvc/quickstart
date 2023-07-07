/*
 *
 * Copyright 2015 gRPC authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

// Package main implements a client for Greeter service.
package main

import (
	"context"
	"flag"
	"fmt"
	pb "github.com/ethanvc/quickstart/grpc/sendrawbytes/helloworldx"
	"google.golang.org/protobuf/proto"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	defaultName = "world"
)

var (
	addr = flag.String("addr", "localhost:50051", "the address to connect to")
	name = flag.String("name", defaultName, "Name to greet")
)

func main() {
	flag.Parse()
	// Set up a connection to the server.
	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewGreeterClient(conn)

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), 10000*time.Second)
	defer cancel()
	SendBytes(conn)
	r, err := c.SayHello(ctx, &pb.HelloRequest{Name: *name})
	r, err = c.SayHello(ctx, &pb.HelloRequest{Name: *name})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", r.GetMessage())
}

func SendBytes(conn *grpc.ClientConn) {
	opts := []grpc.CallOption{grpc.ForceCodec(internalCodec{})}
	req := &pb.HelloRequest{Name: "nihao"}
	reqBytes, _ := proto.Marshal(req)
	var respBytes []byte
	err := conn.Invoke(context.Background(), pb.Greeter_SayHello_FullMethodName, reqBytes, &respBytes, opts...)
	var resp pb.HelloRequest
	err = proto.Unmarshal(respBytes, &resp)
	fmt.Println(err)
}

type internalCodec struct {
}

func (internalCodec) Marshal(v any) ([]byte, error) {
	return v.([]byte), nil
}

func (internalCodec) Unmarshal(data []byte, v interface{}) error {
	realV := v.(*[]byte)
	*realV = data
	return nil
}

func (internalCodec) Name() string {
	return "xx"
}
