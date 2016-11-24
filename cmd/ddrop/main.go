package main

import (
	"crypto/rand"
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/csrf"
	"github.com/julienschmidt/httprouter"
	"github.com/ooesili/dead-drop/internal"
)

func main() {
	if err := realMain(); err != nil {
		fmt.Fprintf(os.Stderr, "ddrop: %s\n", err)
		os.Exit(1)
	}
}

func realMain() error {
	dropHandler := internal.DropHandler{}

	router := httprouter.New()
	router.GET("/", dropHandler.DropForm)
	router.POST("/drop", dropHandler.Drop)
	router.GET("/complete", dropHandler.Complete)
	router.ServeFiles("/assets/*filepath", internal.AssetFS)

	wrapCSRF := csrf.Protect(
		mustGetRandomBytes(32),
		csrf.Secure(false),
	)

	addr := getAddr()
	fmt.Printf("listening on %s for dead drop\n", addr)
	return http.ListenAndServe(addr, wrapCSRF(router))
}

func getAddr() string {
	host := "0.0.0.0"
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	return fmt.Sprintf("%s:%s", host, port)
}

func mustGetRandomBytes(size int) []byte {
	csrfKey := make([]byte, size)
	if _, err := rand.Read(csrfKey); err != nil {
		panic(err)
	}
	return csrfKey
}
