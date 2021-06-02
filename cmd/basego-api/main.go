package main

import (
	"context"
	"flag"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	appV1 "github.com/jonylim/basego/internal/app/basego-api/v1"
	"github.com/jonylim/basego/internal/pkg/common/api"
	"github.com/jonylim/basego/internal/pkg/common/asset"
	"github.com/jonylim/basego/internal/pkg/common/constant/envvar"
	"github.com/jonylim/basego/internal/pkg/common/data/db"
	"github.com/jonylim/basego/internal/pkg/common/data/redis"
	"github.com/jonylim/basego/internal/pkg/common/logger"
	"github.com/jonylim/basego/internal/pkg/common/send/email"
	"github.com/jonylim/basego/internal/pkg/common/storage"

	"github.com/joho/godotenv"
	"github.com/julienschmidt/httprouter"
)

// AppName defines the app's name.
const AppName = "basego-api"

// ShutdownTimeout defines the duration which the server will wait for pending requests to complete before shutdown.
const ShutdownTimeout = 10 * time.Second

func main() {
	// Hide log timestamp.
	logger.HideTimestamp()
	logger.Println("main", "Starting app...")

	// Define argument flags.
	envFile := flag.String("env-file", "", "File containing environment variables to load")
	srvPort := flag.String("port", "", "HTTP server port")

	// Parse app argument flags.
	flag.Parse()

	// Load the environment variables.
	if *envFile != "" {
		if err := godotenv.Load(*envFile); err != nil {
			logger.Println("main", fmt.Sprintf("ERROR: %v", err))
			osExit1()
		} else {
			logger.Println("main", "Environment variables loaded from "+*envFile)
		}
	}
	if err := godotenv.Load(); err != nil {
		logger.Println("main", fmt.Sprintf("WARN: %v", err))
	} else {
		logger.Println("main", "Environment variables loaded from default")
	}

	// Init logger.
	logger.Init()

	// Init databases.
	db.Init()
	defer db.CloseAll()

	// Init Redis.
	redis.Init()
	defer redis.Close()

	// Initialize asset configurations.
	asset.Init()

	// Init storage configurations.
	storage.Init()

	// Init email sender.
	email.Init()

	// Create the server
	srv := newServer(*srvPort)

	stopped := make(chan bool)
	go func() {
		chsig := make(chan os.Signal, 1)
		signal.Notify(chsig, syscall.SIGINT, syscall.SIGTERM)

		// Wait for an interrupt or terminate signal.
		sig := <-chsig
		logger.Println("main", fmt.Sprintf("Signal received: %v", sig))

		// Shut the HTTP server down.
		shutdown := make(chan bool)
		go func() {
			logger.Println("main", "Shutting HTTP server down...")
			ctx, cancel := context.WithTimeout(context.Background(), ShutdownTimeout)
			defer cancel()
			if err := srv.Shutdown(ctx); err != nil {
				// Error from closing listeners, or context timeout.
				logger.Println("main", fmt.Sprintf("ERROR: HTTP server shutdown: %v", logger.FromError(err)))
			}
			close(shutdown)
		}()

		// Wait for shutdown to complete.
		<-shutdown
		close(stopped)
	}()
	go notifyAppStartedEvent(AppName)

	// Start the server.
	go startServer(srv)

	// Wait until stopped.
	<-stopped
	notifyAppStoppedEvent(AppName)
	logger.Println("main", "App stopped")
}

func newServer(port string) *http.Server {
	// Validate the server port.
	if port == "" {
		port = os.Getenv(envvar.ServerPort)
		if port == "" {
			logger.Println("main", "ERROR: Server port is undefined")
			osExit1()
		}
	}

	// Create the router.
	router := httprouter.New()
	router.GET("/", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		w.WriteHeader(404)
	})
	router.GET("/test/show_request_info", handleTestShowRequestInfo)

	// Route APIs.
	appV1.RouteAPIs(router)

	// Route apiDoc.
	router.ServeFiles("/apidoc/*filepath", http.Dir("apidoc-basego-api"))

	// Create new server.
	return &http.Server{Addr: ":" + port, Handler: router}
}

func startServer(srv *http.Server) {
	logger.Println("main", fmt.Sprintf("Serving HTTP server on %s...", srv.Addr))
	if err := srv.ListenAndServe(); err != http.ErrServerClosed {
		// Error starting or closing listener.
		logger.Println("main", fmt.Sprintf("ERROR: HTTP server ListenAndServe: %v", logger.FromError(err)))

		// Wait for 1 second before exit.
		time.Sleep(1 * time.Second)
		osExit1()
	}
}

func handleTestShowRequestInfo(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var sb strings.Builder
	sb.WriteString("RemoteAddr: " + r.RemoteAddr +
		"\nClient IP Address: " + api.GetClientIPAddress(r) +
		"\n")
	for key, values := range r.Header {
		sb.WriteString("\n" + key + ": \"" + strings.Join(values, "\", \"") + "\"")
	}
	w.Write([]byte(sb.String()))
}

/*
func handleDelay(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	c := "reqid"
	start := time.Now()
	logger.Debug("handleDelay", fmt.Sprintf("[%v] Request received at %v", c, start))

	time.Sleep(8 * time.Second)
	// var data []byte
	// for i := 0; i <= 5000000; i++ {
	// 	data = []byte(fmt.Sprintf("%v", time.Now()))
	// }
	// w.Write(data)

	logger.Debug("handleDelay", fmt.Sprintf("[%v] Request completed at %v", c, time.Now()))
}
*/

func osExit1() {
	// Wait for 1 second before exit.
	notifyAppStartFailedEvent(AppName)
	time.Sleep(1 * time.Second)
	os.Exit(1)
}

func notifyAppStartedEvent(appName string) {
	//
}

func notifyAppStoppedEvent(appName string) {
	//
}

func notifyAppStartFailedEvent(appName string) {
	//
}
