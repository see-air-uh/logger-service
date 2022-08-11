package main

import (
	"context"

	"github.com/see-air-uh/logger-service/data"
	"github.com/see-air-uh/logger-service/logs"
)

type LogServer struct {
	// this is going to be required for pretty much every service with GRPC
	logs.UnimplementedLogServiceServer
	Models data.Models
}

func (l *LogServer) WriteLog(ctx context.Context, req *logs.LogRequest) (*logs.LogResponse, error) {

	input := req.GetLogEntry()

	// write the log
	logEntry := data.LogEntry{
		Name: input.Name,
		Data: input.Data,
	}

	err := l.Models.LogEntry.Insert(logEntry)
	if err != nil {
		res := &logs.LogResponse{Result: "failed"}
		return res, err
	}
	res := &logs.LogResponse{Result: "logged"}
	return res, nil

}
