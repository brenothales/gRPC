package main

import (
	"database/sql"
	"log"
	"net"

	_ "github.com/mattn/go-sqlite3"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"github.com/brenothales/gRPC/internal/entity"
	"github.com/brenothales/gRPC/internal/pb"
	"github.com/brenothales/gRPC/internal/service"
)

func main() {
	db, err := sql.Open("sqlite3", "./db.sqlite")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	categoryDb := entity.NewCategory(db)
	categoryService := service.NewCategoryService(*categoryDb)

	CourseDb := entity.NewCourse(db)
	courseService := service.NewCourseService(*CourseDb)

	grpcServer := grpc.NewServer()
	pb.RegisterCategoryServiceServer(grpcServer, categoryService)
	pb.RegisterCourseServiceServer(grpcServer, courseService)
	reflection.Register(grpcServer)

	lis, err := net.Listen("tcp", ":50051")
	log.Printf("server listening at %v", lis.Addr())
	if err != nil {
		panic(err)
	}

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
		panic(err)
	}
}
