package infra

import (
	"context"

	"github.com/go-kit/kit/endpoint"
	"github.com/julienschmidt/httprouter"

	"github.com/selmison/seed-desafio-cdc/pkg/actor/domain"
	"github.com/selmison/seed-desafio-cdc/pkg/actor/service"
	coreDomain "github.com/selmison/seed-desafio-cdc/pkg/core/domain"
)

type CreateRequest struct {
	domain.NewActorDTO
}

type CreateResponse struct {
	Actor domain.ActorDTO `json:"actor"`
	Err   error           `json:"error,omitempty"`
}

func (c CreateResponse) Error() error { return c.Err }

func MakeCreateEndpoint(svc service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(CreateRequest)
		id, err := coreDomain.GenerateId()
		if err != nil {
			return CreateResponse{domain.ActorDTO{}, err}, nil
		}
		actor := domain.Actor{
			Id:          id,
			Name:        req.Name,
			Email:       req.Email,
			Description: req.Description,
			CreatedAt:   domain.GenerateTime(),
		}
		if err = svc.Create(ctx, actor); err != nil {
			return CreateResponse{domain.ActorDTO{}, err}, nil
		}
		return CreateResponse{domain.ActorDTO{
			Id:          actor.Id,
			Name:        actor.Name,
			Email:       actor.Email,
			Description: actor.Description,
			CreatedAt:   actor.CreatedAt,
		}, nil}, nil
	}
}

func MakeDestroyEndpoint(svc service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		id := httprouter.ParamsFromContext(ctx).ByName("id")
		if err := svc.Destroy(ctx, id); err != nil {
			return err, nil
		}
		return nil, nil
	}
}

type ListResponse struct {
	Actors []domain.ActorDTO `json:"actors"`
	Err    error             `json:"error,omitempty"`
}

func (c ListResponse) Error() error { return c.Err }

func MakeListEndpoint(svc service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		actors, err := svc.List(ctx)
		if err != nil {
			return ListResponse{nil, err}, nil
		}
		actorDTOS := make([]domain.ActorDTO, len(actors))
		for i, actor := range actors {
			actorDTOS[i] = domain.ActorDTO{
				Id:          actor.Id,
				Name:        actor.Name,
				Email:       actor.Email,
				Description: actor.Description,
				CreatedAt:   actor.CreatedAt,
			}
		}
		return ListResponse{actorDTOS, nil}, nil
	}
}

type ShowResponse struct {
	Actor domain.ActorDTO `json:"actor"`
	Err   error           `json:"error,omitempty"`
}

func (c ShowResponse) Error() error { return c.Err }

func MakeShowEndpoint(svc service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		id := httprouter.ParamsFromContext(ctx).ByName("id")
		actor, err := svc.Show(ctx, id)
		if err != nil {
			return ShowResponse{domain.ActorDTO{}, err}, nil
		}
		actorDTO := domain.ActorDTO{
			Id:          actor.Id,
			Name:        actor.Name,
			Email:       actor.Email,
			Description: actor.Description,
			CreatedAt:   actor.CreatedAt,
		}
		return ShowResponse{actorDTO, nil}, nil
	}
}

type UpdateRequest struct {
	domain.UpdateActorDTO
}

func MakeUpdateEndpoint(svc service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		id := httprouter.ParamsFromContext(ctx).ByName("id")
		req := request.(UpdateRequest)
		return svc.Update(ctx, id, req.UpdateActorDTO), nil
	}
}
