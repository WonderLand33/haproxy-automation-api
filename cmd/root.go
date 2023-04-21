package cmd

import (
	"context"
	"haproxy-automation-api/internal/pkg/route"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.Flags().StringVar(&addr, "addr", "127.0.0.1:23333", "route addr")
}

var (
	addr string
)

var rootCmd = &cobra.Command{
	Use: "haproxy-automation-api",
	Run: func(cmd *cobra.Command, args []string) {

		e := echo.New()
		route.Install(e)

		// Start server
		go func() {
			if err := e.Start(addr); err != nil && err != http.ErrServerClosed {
				log.Fatalln("shutting down the server")
			}
		}()

		// Wait for interrupt signal to gracefully shutdown the server with a timeout of 10 seconds.
		// Use a buffered channel to avoid missing signals as recommended for signal.Notify
		quit := make(chan os.Signal, 1)
		signal.Notify(quit, os.Interrupt)
		<-quit
		log.Println("exiting...")
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		if err := e.Shutdown(ctx); err != nil {
			log.Fatalln(err)
		}
		log.Println("Bye bye")
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
