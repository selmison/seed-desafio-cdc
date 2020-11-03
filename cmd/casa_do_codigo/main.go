package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/url"
	"os"
	"os/signal"
	"strings"
	"sync"
	"syscall"

	kitLog "github.com/go-kit/kit/log"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	actorsGen "github.com/selmison/seed-desafio-cdc/gen/actors"
	booksGen "github.com/selmison/seed-desafio-cdc/gen/books"
	categoriesGen "github.com/selmison/seed-desafio-cdc/gen/categories"
	"github.com/selmison/seed-desafio-cdc/pkg/actors"
	"github.com/selmison/seed-desafio-cdc/pkg/books"
	"github.com/selmison/seed-desafio-cdc/pkg/categories"
	coreDomain "github.com/selmison/seed-desafio-cdc/pkg/core/domain"
)

func main() {
	var (
		hostF     = flag.String("host", "development", "Server host (valid values: development)")
		domainF   = flag.String("domain", "", "Host domain name (overrides host domain specified in service design)")
		httpPortF = flag.String("http-port", "", "HTTP port (overrides host HTTP port specified in service design)")
		secureF   = flag.Bool("secure", false, "Use secure scheme (https or grpcs)")
		dbgF      = flag.Bool("debug", false, "Log request and response bodies")
		dsnF      = flag.String("dsn", "file::memory:?cache=shared", "data source name")
	)
	flag.Parse()

	var (
		logger kitLog.Logger
	)
	{
		logger = kitLog.NewLogfmtLogger(os.Stderr)
		logger = kitLog.With(logger, "ts", kitLog.DefaultTimestampUTC)
		logger = kitLog.With(logger, "caller", kitLog.DefaultCaller)
	}

	var (
		repo          *gorm.DB
		actorsSvc     actorsGen.Service
		categoriesSvc categoriesGen.Service
		booksSvc      booksGen.Service
	)
	{
		var err error
		repo, err = gorm.Open(sqlite.Open(*dsnF), &gorm.Config{
			SkipDefaultTransaction: true,
		})
		if err != nil {
			panic("failed to connect database")
		}
		if err := repo.AutoMigrate(
			&actors.Actor{},
			&categories.Category{},
			&books.Book{},
		); err != nil {
			log.Fatalf("db init: %v", err)
		}
		actorsSvc = actors.NewService(repo, logger)
		categoriesSvc = categories.NewService(repo, logger)
		booksSvc = books.NewService(repo, logger)
	}

	// Wrap the services in endpoints that can be invoked from other services
	// potentially running in different processes.
	var (
		actorsEndpoints     *actorsGen.Endpoints
		categoriesEndpoints *categoriesGen.Endpoints
		booksEndpoints      *booksGen.Endpoints
	)
	{
		actorsEndpoints = actorsGen.NewEndpoints(actorsSvc)
		actorsEndpoints.Use(ValidationCreateActorMiddleware(repo))

		categoriesEndpoints = categoriesGen.NewEndpoints(categoriesSvc)
		categoriesEndpoints.Use(ValidationCreateCategoryMiddleware(repo))

		booksEndpoints = booksGen.NewEndpoints(booksSvc)
		booksEndpoints.Use(ValidationCreateBookMiddleware(repo))
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
			addr := "http://localhost:3333"
			u, err := url.Parse(addr)
			if err != nil {
				if _, err := fmt.Fprintf(os.Stderr, "invalid URL %#v: %s\n", addr, err); err != nil {
					log.Fatalf(coreDomain.FormatToErrFprintf, err)
				}
				os.Exit(1)
			}
			if *secureF {
				u.Scheme = "https"
			}
			if *domainF != "" {
				u.Host = *domainF
			}
			if *httpPortF != "" {
				h := strings.Split(u.Host, ":")[0]
				u.Host = h + ":" + *httpPortF
			} else if u.Port() == "" {
				u.Host += ":80"
			}
			handleHTTPServer(ctx, u, actorsEndpoints, categoriesEndpoints, booksEndpoints, &wg, errc, logger, *dbgF)
		}

	default:
		if _, err := fmt.Fprintf(os.Stderr, "invalid host argument: %q (valid hosts: development)\n", *hostF); err != nil {
			log.Fatalf(coreDomain.FormatToErrFprintf, err)
		}
	}

	// Wait for signal.
	if err := logger.Log("info", fmt.Sprintf("exiting (%v)", <-errc)); err != nil {
		log.Fatalf(coreDomain.FormatToErrKitLog, err)
	}

	// Send cancellation signal to the goroutines.
	cancel()

	wg.Wait()
	if err := logger.Log("info", fmt.Sprintf("exited")); err != nil {
		log.Fatalf(coreDomain.FormatToErrKitLog, err)
	}
}
