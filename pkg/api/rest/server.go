package rest

import (
	"fmt"
	"log"
	"net/http"
	"os"

	kitLog "github.com/go-kit/kit/log"

	"github.com/julienschmidt/httprouter"

	"github.com/selmison/seed-desafio-cdc/pkg/actor/infra"
	"github.com/selmison/seed-desafio-cdc/pkg/actor/service"
)

type server struct {
	router *httprouter.Router
	svc    service.Service
}

// InitHttpServer initialize returns a Http Server.
func InitHttpServer(address string) error {
	s := NewServer()
	fmt.Printf("The server is on tap now: http://%s\n", address)
	if err := http.ListenAndServe(address, s); err != nil {
		return err
	}
	return nil
}

// NewServer returns a new Server.
func NewServer() http.Handler {
	r := httprouter.New()
	r.GET("/", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		if _, err := fmt.Fprint(w, "Welcome!\n"); err != nil {
			log.Println(err)
		}
	})
	srv := &server{router: r}
	logger := kitLog.NewLogfmtLogger(os.Stderr)

	actorRepo := infra.NewActorRepository()
	actorSvc := service.NewService(actorRepo, logger)
	infra.NewActorRoutes(srv, actorSvc, logger)

	return srv
}

// ServeHTTP implements Handler.
func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

// AddRoute implements Router.
func (s *server) AddRoute(method string, pattern string, handler http.Handler) {
	s.router.Handler(method, pattern, handler)
}
