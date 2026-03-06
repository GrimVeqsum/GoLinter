package logrulesdisabled

import (
	"log"
	"log/slog"
)

var zap = struct {
	Info func(string, ...any)
}{
	Info: func(string, ...any) {},
}

func checkLogRules() {
	log.Println("Starting server")
	log.Println("server запуск")
	log.Println("connection failed!!!")
	log.Println("token leaked")

	log.Println("server started")
	slog.Info("service started")
	zap.Info("request completed")
}