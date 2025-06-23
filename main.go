package main

import (
    "log"
    "net"

    "google.golang.org/grpc"
    "github.com/nikhil1010918/grpc-report-service/proto"
    "github.com/nikhil1010918/grpc-report-service/server"
    "github.com/nikhil1010918/grpc-report-service/cron"
)

func main() {
    lis, err := net.Listen("tcp", ":50051")
    if err != nil {
        log.Fatalf("Failed to listen: %v", err)
    }

    grpcServer := grpc.NewServer()
    reportServer := server.NewReportServer()

    proto.RegisterReportServiceServer(grpcServer, reportServer)
    cron.StartCronJob(reportServer)

    log.Println("gRPC server running on port 50051")
    if err := grpcServer.Serve(lis); err != nil {
        log.Fatalf("Failed to serve: %v", err)
    }
}
