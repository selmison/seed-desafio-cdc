package testdata

import (
	"log"

	"github.com/Pallinder/go-randomdata"
	"github.com/bluele/factory-go/factory"
	"github.com/bxcodec/faker/v3"

	actorDomain "github.com/selmison/seed-desafio-cdc/pkg/actor/domain"
	"github.com/selmison/seed-desafio-cdc/pkg/core/domain"
)

const FakeActorsLength = 10

var (
	newActorFactory = factory.NewFactory(
		actorDomain.NewActorDTO{},
	).Attr("Name", func(args factory.Args) (interface{}, error) {
		return actorDomain.NewName(randomdata.FullName(randomdata.RandomGender))
	}).Attr("Email", func(args factory.Args) (interface{}, error) {
		return actorDomain.NewEmail(randomdata.Email())
	}).Attr("Description", func(args factory.Args) (interface{}, error) {
		return actorDomain.NewDesc(randomdata.Paragraph())
	})

	FakeActors    []actorDomain.Actor
	FakeActorDTOs []actorDomain.ActorDTO
	FakeNewActors []actorDomain.NewActorDTO
)

func init() {
	var err error
	FakeActors, FakeNewActors, FakeActorDTOs, err = generateFakeActors(FakeActorsLength)
	if err != nil {
		log.Fatalf("init: failed to generateFakeActors: %v\n", err)
	}

}

func generateFakeActors(length int) (
	[]actorDomain.Actor,
	[]actorDomain.NewActorDTO,
	[]actorDomain.ActorDTO,
	error,
) {
	fakeNewActors := make([]actorDomain.NewActorDTO, length)
	fakeActors := make([]actorDomain.Actor, length)
	var fakeActorDTOs []actorDomain.ActorDTO
	for i := 0; i < length; i++ {
		fakeNewActors[i] = newActorFactory.MustCreate().(actorDomain.NewActorDTO)
		id, err := domain.GenerateId()
		if err != nil {
			return nil, nil, nil, err
		}
		name, err := actorDomain.NewName(faker.Name())
		if err != nil {
			return nil, nil, nil, err
		}
		email, err := actorDomain.NewEmail(faker.Email())
		if err != nil {
			return nil, nil, nil, err
		}
		desc, err := actorDomain.NewDesc(faker.Sentence())
		if err != nil {
			return nil, nil, nil, err
		}
		fakeActors[i] = actorDomain.Actor{
			Id:          id,
			Name:        name,
			Email:       email,
			Description: desc,
		}
		fakeActorDTOs = append(fakeActorDTOs, actorDomain.ActorDTO{
			Id:          fakeActors[i].Id,
			Name:        fakeActors[i].Name,
			Email:       fakeActors[i].Email,
			Description: fakeActors[i].Description,
			CreatedAt:   fakeActors[i].CreatedAt,
		})
	}
	return fakeActors, fakeNewActors, fakeActorDTOs, nil
}
