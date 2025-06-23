package server

import (
    "context"
    "fmt"
    "log"
    "sync"
    "time"

    "github.com/nikhil1010918/grpc-report-service/proto"
)

type ReportServer struct {
    proto.UnimplementedReportServiceServer
    reports map[string]string
    mu      sync.Mutex
}

func NewReportServer() *ReportServer {
    return &ReportServer{
        reports: make(map[string]string),
    }
}

func (s *ReportServer) GenerateReport(ctx context.Context, req *proto.GenerateReportRequest) (*proto.GenerateReportResponse, error) {
    s.mu.Lock()
    defer s.mu.Unlock()

    reportID := fmt.Sprintf("report-%s-%d", req.UserId, time.Now().UnixNano())
    report := fmt.Sprintf("Report for user %s generated at %s", req.UserId, time.Now().Format(time.RFC3339))

    s.reports[reportID] = report

    log.Printf("[Generated] %s -> %s", reportID, report)

    return &proto.GenerateReportResponse{ReportId: reportID}, nil
}

func (s *ReportServer) HealthCheck(ctx context.Context, req *proto.HealthCheckRequest) (*proto.HealthCheckResponse, error) {
    return &proto.HealthCheckResponse{Status: "SERVING"}, nil
}
