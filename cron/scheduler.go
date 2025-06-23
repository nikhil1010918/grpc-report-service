package cron

import (
	"log"

	"github.com/nikhil1010918/grpc-report-service/proto"
	"github.com/robfig/cron/v3"
)

func StartCronJob(server proto.ReportServiceServer) {
	c := cron.New()
	c.AddFunc("@every 10s", func() {
		for _, user := range []string{"42", "100", "7"} {
			_, err := server.GenerateReport(nil, &proto.GenerateReportRequest{UserId: user})
			if err != nil {
				log.Printf("[Cron Error] Failed to generate report for user %s: %v", user, err)
			}
		}
	})
	c.Start()
}
