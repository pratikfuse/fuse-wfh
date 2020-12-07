package auth

import (
	"log"
	"net/http"
	"strconv"
	"sync"
)

type Server struct {
}

var waitGroup sync.WaitGroup

func handleAuthResponse(shutdownSignal chan bool, errorChannel chan error) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		code := req.URL.Query()["code"]
		waitGroup.Add(1)
		credentials := GetAccessToken(code[0], &waitGroup)
		waitGroup.Wait()
		err := SaveCredentialsToFile(credentials)
		shutdownSignal <- true
		errorChannel <- err
	}
}

func (s *Server) RunAuthServer(port int, errorChan chan error, shutdownSignal chan bool) {
	log.Printf("opening auth server on port %d \n", port)
	srv := http.Server{
		Addr: ":" + strconv.Itoa(port),
	}
	http.HandleFunc("/oauth", handleAuthResponse(shutdownSignal, errorChan))
	go func() {
		if err := srv.ListenAndServe(); err != nil {
			errorChan <- err
		}
	}()
	shutdown := <-shutdownSignal

	if shutdown {
		log.Println("Shutting down auth server")
		_ = srv.Close()
	}
}
