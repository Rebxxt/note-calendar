package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"
	"net"
	notes "notes/gen/go"
)

type server struct {
	notes.UnimplementedNotesServer
}

func (s *server) Get(context.Context, *notes.NoteGetRequest) (*notes.NoteGetResponse, error) {
	return &notes.NoteGetResponse{}, nil
}
func (s *server) List(context.Context, *notes.NoteListRequest) (*notes.NoteListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (s *server) Create(context.Context, *notes.NoteCreateRequest) (*notes.NoteCreateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (s *server) Delete(context.Context, *notes.NoteDeleteRequest) (*notes.NoteDeleteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (s *server) Update(context.Context, *notes.NoteUpdateRequest) (*notes.NoteUpdateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}

func main() {
	srv := grpc.NewServer()
	listener, err := net.Listen("tcp", "0.0.0.0:9000")
	if err != nil {
		panic(err)
	}

	reflection.Register(srv)
	notes.RegisterNotesServer(srv, &server{})

	err = srv.Serve(listener)
	if err != nil {
		panic(err)
	}
	fmt.Println("Starting server...")
}
