package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"net/url"
	"os"
	"os/signal"
	"sync"
	"syscall"

	db "github.com/tomchan123/go-news-mailer/internal/db"
	"github.com/tomchan123/go-news-mailer/internal/gen/email"
	subscription "github.com/tomchan123/go-news-mailer/internal/gen/subscription"
	"github.com/tomchan123/go-news-mailer/internal/mailer"
	news "github.com/tomchan123/go-news-mailer/internal/news"
)

func main() {
	// Define command line flags, add any other flag required to configure the
	// service.
	var (
		hostF     = flag.String("host", "development", "Server host (valid values: development)")
		domainF   = flag.String("domain", "", "Host domain name (overrides host domain specified in service design)")
		httpPortF = flag.String("http-port", "", "HTTP port (overrides host HTTP port specified in service design)")
		secureF   = flag.Bool("secure", false, "Use secure scheme (https or grpcs)")
		dbgF      = flag.Bool("debug", false, "Log request and response bodies")
	)
	flag.Parse()

	// Setup logger. Replace logger with your own log package of choice.
	var (
		logger *log.Logger
	)
	{
		logger = log.New(os.Stderr, "[news] ", log.Ltime)
	}

	// Setup DB
	var (
		dbx *db.Mongodb
	)
	{
		dbx = &db.Mongodb{}
		d, err := dbx.ConnectDB("mongodb://localhost:27017")
		if err != nil {
			log.Fatal(err)
		}
		defer d()
	}

	// Setup News Server
	var (
		ns *mailer.NewsServer
	)
	{
		ns = mailer.CreateNewsServer()
	}

	// Setup Mail Server
	var (
		ms *mailer.MailServer
	)
	{
		ms = mailer.CreateMailer(dbx, ns)
		go ms.Start()
	}

	// Initialize the services.
	var (
		subscriptionSvc subscription.Service
		emailSvc        email.Service
	)
	{
		subscriptionSvc = news.NewSubscription(logger, dbx)
		emailSvc = news.NewEmail(logger, ms)
	}

	// Wrap the services in endpoints that can be invoked from other services
	// potentially running in different processes.
	var (
		subscriptionEndpoints *subscription.Endpoints
		emailEndpoints        *email.Endpoints
	)
	{
		subscriptionEndpoints = subscription.NewEndpoints(subscriptionSvc)
		emailEndpoints = email.NewEndpoints(emailSvc)
	}

	// Create channel used by both the signal handler and server goroutines
	// to notify the main goroutine when to stop the server.
	errc := make(chan error)

	// Setup interrupt handler. This optional step configures the process so
	// that SIGINT and SIGTERM signals cause the services to stop gracefully.
	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errc <- fmt.Errorf("%s", <-c)
	}()

	var wg sync.WaitGroup
	ctx, cancel := context.WithCancel(context.Background())

	// Start the servers and send errors (if any) to the error channel.
	switch *hostF {
	case "development":
		{
			addr := "http://localhost:8888"
			u, err := url.Parse(addr)
			if err != nil {
				logger.Fatalf("invalid URL %#v: %s\n", addr, err)
			}
			if *secureF {
				u.Scheme = "https"
			}
			if *domainF != "" {
				u.Host = *domainF
			}
			if *httpPortF != "" {
				h, _, err := net.SplitHostPort(u.Host)
				if err != nil {
					logger.Fatalf("invalid URL %#v: %s\n", u.Host, err)
				}
				u.Host = net.JoinHostPort(h, *httpPortF)
			} else if u.Port() == "" {
				u.Host = net.JoinHostPort(u.Host, "80")
			}
			handleHTTPServer(ctx, u, subscriptionEndpoints, emailEndpoints, &wg, errc, logger, *dbgF)
		}

	default:
		logger.Fatalf("invalid host argument: %q (valid hosts: development)\n", *hostF)
	}

	// Wait for signal.
	logger.Printf("exiting (%v)", <-errc)

	// Send cancellation signal to the goroutines.
	cancel()

	wg.Wait()
	logger.Println("exited")
}
