package main

import (
	"github.com/OsipyanG/market/services/gateway-msv/config"
	"github.com/OsipyanG/market/services/gateway-msv/internal/app"
)

func main() {
	cfg := config.MustLoad()
	app.Run(cfg)
}
