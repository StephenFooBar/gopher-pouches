package test

import (
	"context"
	"io"
	"log"
	"net/http"
	"sync"

	"github.com/StephenFooBar/gopher-pouches/config"
)

func startHttpServer(wg *sync.WaitGroup, port string) *http.Server {
	srv := &http.Server{Addr: "localhost:" + port}
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "hello world\n")
	})

	feeds := http.FileServer(http.Dir(config.Root + "/test/feeds"))
	http.Handle("/feeds/", http.StripPrefix("/feeds/", feeds))

	go func() {
		defer wg.Done()

		if err := srv.ListenAndServe(); err != http.ErrServerClosed {
			log.Fatalf("ListenAndServe(): %v", err)
		}
	}()

	return srv
}

func RunMockHttpServer(port string) (*http.Server, *sync.WaitGroup) {
	log.Printf("Starting Mock HTTP Server.")
	httpServerExitDone := &sync.WaitGroup{}
	httpServerExitDone.Add(1)
	srv := startHttpServer(httpServerExitDone, port)
	log.Printf("Running Mock HTTP Server. Listening on Port:%s", port)
	return srv, httpServerExitDone
}

func StopMockHttpServer(srv *http.Server, httpServerExitDone *sync.WaitGroup) {
	log.Printf("Stopping Mock HTTP Server.")
	if err := srv.Shutdown(context.TODO()); err != nil {
		panic(err)
	}
	httpServerExitDone.Wait()
	log.Printf("Mock HTTP Server Stopped.")
}
