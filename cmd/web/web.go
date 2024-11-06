package main

import (
	"context"
	"log/slog"
	"net"
	"net/http"
	"os"
	"strconv"

	"time"

	"github.com/datasektionen/logout/pkg/config"
	"github.com/datasektionen/logout/pkg/database"
	"github.com/datasektionen/logout/service"
	"github.com/datasektionen/logout/services/admin"
	"github.com/datasektionen/logout/services/dev"
	"github.com/datasektionen/logout/services/legacyapi"
	"github.com/datasektionen/logout/services/oidcprovider"
	"github.com/datasektionen/logout/services/oidcrp"
	"github.com/datasektionen/logout/services/passkey"
	"github.com/datasektionen/logout/services/static"
	"github.com/datasektionen/logout/services/user"
)

func main() {
	initCtx, cancel := context.WithTimeout(context.Background(), time.Minute)
	db, err := database.ConnectAndMigrate(initCtx)
	if err != nil {
		panic(err)
	}

	s := must(service.NewService(initCtx, db))
	if err := oidcprovider.MountRoutes(s); err != nil {
		panic(err)
	}
	cancel()

	user.MountRoutes(s)
	passkey.MountRoutes(s)
	oidcrp.MountRoutes(s)
	legacyapi.MountRoutes(s)
	dev.MountRoutes(s)
	admin.MountRoutes(s)

	static.Mount()

	port := strconv.Itoa(config.Config.Port)
	l, err := net.Listen("tcp", ":"+port)
	if err != nil {
		slog.Error("Could not start listening for connections", "port", port, "error", err)
		os.Exit(1)
	}
	slog.Info("Server started", "address", "http://localhost:"+port)
	slog.Error("Failed serving http server", "error", http.Serve(l, nil))
	os.Exit(1)
}

func must[T any](t T, err error) T {
	if err != nil {
		panic(err)
	}
	return t
}
