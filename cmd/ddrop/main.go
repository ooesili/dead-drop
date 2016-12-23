package main

import (
	"crypto/rand"
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/csrf"
	"github.com/justinas/alice"
	"github.com/ooesili/dead-drop/internal"
	"github.com/ooesili/dead-drop/internal/config"
	"github.com/ooesili/dead-drop/internal/handler"
	"github.com/ooesili/dead-drop/internal/localdrop"
	"github.com/ooesili/dead-drop/internal/view"
)

func main() {
	if err := realMain(); err != nil {
		fmt.Fprintf(os.Stderr, "ddrop: %s\n", err)
		os.Exit(1)
	}
}

func realMain() error {
	config, err := internal.ConfigSchema.NewConfig(os.LookupEnv)
	if err != nil {
		return err
	}

	app := newApp()

	addr := getAddr(config)
	fmt.Printf("listening on %s for dead drop\n", addr)
	return http.ListenAndServe(addr, app)
}

func getAddr(config config.Config) string {
	return config.Get("BIND_ADDR") + ":" + config.Get("PORT")
}

func mustGetRandomBytes(size int) []byte {
	csrfKey := make([]byte, size)
	if _, err := rand.Read(csrfKey); err != nil {
		panic(err)
	}
	return csrfKey
}

func newApp() http.Handler {
	dropper := localdrop.LocalDrop{
		Out:     os.Stdout,
		BaseDir: ".",
	}

	router := handler.NewRouter(
		dropper,
		view.New(),
	)
	router.ServeFiles("/assets/*filepath", internal.AssetFS)

	wrapCSRF := csrf.Protect(
		mustGetRandomBytes(32),
		csrf.Secure(false),
	)

	return alice.New(wrapCSRF).Then(router)
}
