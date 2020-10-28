// +build integration

package rest_test

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/bxcodec/faker/v3"
	"github.com/stretchr/testify/require"

	"github.com/selmison/seed-desafio-cdc/pkg/actor/domain"
	"github.com/selmison/seed-desafio-cdc/pkg/actor/infra"
	"github.com/selmison/seed-desafio-cdc/pkg/api/rest"
	coreDomain "github.com/selmison/seed-desafio-cdc/pkg/core/domain"
	coreInfra "github.com/selmison/seed-desafio-cdc/pkg/core/infra"
	"github.com/selmison/seed-desafio-cdc/testdata"
)

var (
	fakeActors = testdata.FakeActors
)

func Test_RestApi_Post_Actors(t *testing.T) {
	teardownTestCase, _, err := infra.SetupActorTestCase(t, nil)
	if err != nil {
		t.Fatalf("test: failed to setup test case: %v\n", err)
	}
	defer teardownTestCase(t)
	srv := httptest.NewServer(rest.NewServer())
	defer srv.Close()
	fakeURL := fmt.Sprintf("%s/%s", srv.URL, "actors")
	name, err := domain.NewName(faker.Name())
	if err != nil {
		t.Fatalf("test: failed to create a new valid name: %v\n", err)
	}
	email, err := domain.NewEmail(faker.Email())
	if err != nil {
		t.Fatalf("test: failed to create a new valid e-mail: %v\n", err)

	}
	desc, err := domain.NewDesc(faker.Sentence())
	if err != nil {
		t.Fatalf("test: failed to create a new valid description: %v\n", err)
	}
	fakeNewActorDTO := domain.NewActorDTO{
		Name:        name,
		Email:       email,
		Description: desc,
	}
	fakeNewActorJSONWithNameBlank := fmt.Sprintf(
		`{"name": "%s", "e-mail": "%s", "description": "%s"}`,
		"     ",
		fakeNewActorDTO.Email.String(),
		fakeNewActorDTO.Description.String(),
	)
	fakeNewActorJSON := fmt.Sprintf(
		`{"name": "%s", "e-mail": "%s", "description": "%s"}`,
		fakeNewActorDTO.Name.String(),
		fakeNewActorDTO.Email.String(),
		fakeNewActorDTO.Description.String(),
	)
	type request struct {
		url  string
		body io.Reader
	}
	type statusResponse struct {
		status int
		body   string
	}
	tests := []struct {
		name    string
		req     request
		want    statusResponse
		wantErr bool
	}{
		{
			name: "When name field is blank",
			req: request{
				fakeURL,
				strings.NewReader(fakeNewActorJSONWithNameBlank),
			},
			want: statusResponse{
				status: http.StatusBadRequest,
				body:   http.StatusText(http.StatusBadRequest),
			},
			wantErr: false,
		},
		{
			name: "When everything is right",
			req: request{
				fakeURL,
				strings.NewReader(fakeNewActorJSON),
			},
			want: statusResponse{
				status: http.StatusCreated,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := http.Post(tt.req.url, "application/json; charset=UTF-8", tt.req.body)
			if (err != nil) != tt.wantErr {
				require.Error(t, err)
			}
			if got != nil {
				defer got.Body.Close()
				require.Equal(t, tt.want.status, got.StatusCode, "they should be equal")
				switch tt.want.status {
				case http.StatusCreated:
					createResponse := infra.CreateResponse{}
					if err := json.NewDecoder(got.Body).Decode(&createResponse); err != nil {
						t.Fatalf("could not decode response: %v", err)
					}
					require.Equal(t, createResponse.Actor, createResponse.Actor, "they should be equal")
				}
			}
		})
	}
}

func Test_RestApi_Get_Actors(t *testing.T) {
	teardownTestCase, _, err := infra.SetupActorTestCase(t, fakeActors)
	if err != nil {
		t.Fatalf("test: failed to setup test case: %v\n", err)
	}
	defer teardownTestCase(t)
	srv := httptest.NewServer(rest.NewServer())
	defer srv.Close()
	fakeURL := fmt.Sprintf("%s/%s", srv.URL, "actors")
	type response struct {
		status int
		body   []domain.ActorDTO
	}
	tests := []struct {
		name    string
		url     string
		want    response
		wantErr bool
	}{
		{
			name: "When everything is right",
			url:  fakeURL,
			want: response{
				status: http.StatusOK,
				body:   testdata.FakeActorDTOs,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := http.Get(tt.url)
			if (err != nil) != tt.wantErr {
				require.Error(t, err)
			}
			if got != nil {
				defer got.Body.Close()
				require.Equal(t, tt.want.status, got.StatusCode, "they should be equal")
				switch tt.want.status {
				case http.StatusOK:
					listResponse := infra.ListResponse{}
					if err := json.NewDecoder(got.Body).Decode(&listResponse); err != nil {
						t.Fatalf("could not decode response: %v", err)
					}
					require.ElementsMatch(
						t,
						listResponse.Actors,
						tt.want.body,
						"they should be equal",
					)
				}
			}
		})
	}
}

func Test_RestApi_Get_Actor(t *testing.T) {
	teardownTestCase, _, err := infra.SetupActorTestCase(t, fakeActors)
	if err != nil {
		t.Fatalf("test: failed to setup test case: %v\n", err)
	}
	defer teardownTestCase(t)
	srv := httptest.NewServer(rest.NewServer())
	defer srv.Close()
	fakeExistentActor := fakeActors[0]
	fakeExistentId := fakeExistentActor.Id
	fakeDoesNonExistentId := "fakeDoesNonExistentId"
	fakeExistentActorDTO := domain.ActorDTO{
		Id:          fakeExistentId,
		Name:        fakeExistentActor.Name,
		Email:       fakeExistentActor.Email,
		Description: fakeExistentActor.Description,
		CreatedAt:   fakeExistentActor.CreatAt(),
	}
	fakeURL := func(id string) string {
		return fmt.Sprintf("%s/%s/%s", srv.URL, "actors", id)
	}
	type statusResponse struct {
		status int
		body   []byte
	}
	tests := []struct {
		name    string
		url     string
		want    statusResponse
		wantErr bool
	}{
		{
			name: "When id doesn't exist",
			url:  fakeURL(fakeDoesNonExistentId),
			want: statusResponse{
				status: http.StatusNotFound,
				body:   []byte(http.StatusText(http.StatusNotFound)),
			},
		},
		{
			name: "When id exists",
			url:  fakeURL(fakeExistentId.String()),
			want: statusResponse{
				status: http.StatusOK,
				body:   coreInfra.ToJSON(fakeExistentActorDTO),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := http.Get(tt.url)
			if (err != nil) != tt.wantErr {
				require.Error(t, err)
			}
			if got != nil {
				defer got.Body.Close()
				require.Equal(t, tt.want.status, got.StatusCode, "they should be equal")
				switch tt.want.status {
				case http.StatusOK:
					expectedBody := domain.ActorDTO{}
					if err := json.Unmarshal(tt.want.body, &expectedBody); err != nil {
						t.Fatalf("could not decode expectedBody: %v", err)
					}
					showResponse := infra.ShowResponse{}
					if err := json.NewDecoder(got.Body).Decode(&showResponse); err != nil {
						t.Fatalf("could not decode response: %v", err)
					}
					require.Equal(t, expectedBody, showResponse.Actor, "they should be equal")
				}
			}
		})
	}
}

func Test_RestApi_Delete_Actor(t *testing.T) {
	teardownTestCase, _, err := infra.SetupActorTestCase(t, fakeActors)
	if err != nil {
		t.Fatalf("test: failed to setup test case: %v\n", err)
	}
	defer teardownTestCase(t)
	srv := httptest.NewServer(rest.NewServer())
	defer srv.Close()
	fakeExistentId := fakeActors[0].Id
	fakeNonExistentId, err := coreDomain.NewId("fakeNonExistentId")
	if err != nil {
		t.Fatalf("test: failed to create a new valid id: %v\n", err)
	}
	fakeURL := func(id string) string {
		return fmt.Sprintf("%s/%s/%s", srv.URL, "actors", id)
	}
	type statusResponse struct {
		status int
		body   string
	}
	tests := []struct {
		name    string
		url     string
		want    statusResponse
		wantErr bool
	}{
		{
			name: "When id doesn't exist",
			url:  fakeURL(fakeNonExistentId.String()),
			want: statusResponse{
				status: http.StatusNotFound,
				body:   http.StatusText(http.StatusNotFound),
			},
			wantErr: false,
		},
		{
			name: "When id exists",
			url:  fakeURL(fakeExistentId.String()),
			want: statusResponse{
				status: http.StatusOK,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client := &http.Client{}
			req, err := http.NewRequest(http.MethodDelete, tt.url, nil)
			got, err := client.Do(req)
			if (err != nil) != tt.wantErr {
				require.Error(t, err)
			}
			if got != nil {
				defer got.Body.Close()
				require.Equal(t, tt.want.status, got.StatusCode, "they should be equal")
				switch tt.want.status {
				case http.StatusOK:
					bs, err := ioutil.ReadAll(got.Body)
					if err != nil {
						t.Fatalf("read body: %v", err)
					}
					actualBody := strings.TrimSpace(string(bs))
					require.Equal(t, tt.want.body, actualBody, "they should be equal")
				}
			}
		})
	}
}

func Test_RestApi_Update_Actor(t *testing.T) {
	teardownTestCase, _, err := infra.SetupActorTestCase(t, fakeActors)
	if err != nil {
		t.Fatalf("test: failed to setup test case: %v\n", err)
	}
	defer teardownTestCase(t)
	srv := httptest.NewServer(rest.NewServer())
	defer srv.Close()
	fakeExistentId := fakeActors[0].Id
	fakeNonExistentId, err := coreDomain.NewId("fakeNonExistentId")
	if err != nil {
		t.Fatalf("test: failed to create a new valid id: %v\n", err)
	}
	fakeURL := func(id string) string {
		return fmt.Sprintf("%s/%s/%s", srv.URL, "actors", id)
	}
	name, err := domain.NewName(faker.Name())
	if err != nil {
		t.Fatalf("test: failed to create a new valid name: %v\n", err)
	}
	email, err := domain.NewEmail(faker.Email())
	if err != nil {
		t.Fatalf("test: failed to create a new valid e-mail: %v\n", err)

	}
	desc, err := domain.NewDesc(faker.Sentence())
	if err != nil {
		t.Fatalf("test: failed to create a new valid description: %v\n", err)
	}
	fakeUpdateActorDTO := domain.UpdateActorDTO{
		Name:        &name,
		Email:       &email,
		Description: &desc,
	}
	fakeUpdateActorJSON := fmt.Sprintf(
		`{"name": "%s", "e-mail": "%s", "description": "%s"}`,
		fakeUpdateActorDTO.Name.String(),
		fakeUpdateActorDTO.Email.String(),
		fakeUpdateActorDTO.Description.String(),
	)
	type request struct {
		url         string
		contentType string
		body        io.Reader
	}
	type statusResponse struct {
		status int
		body   string
	}
	tests := []struct {
		name    string
		req     request
		want    statusResponse
		wantErr bool
	}{
		{
			name: "When id doesn't exist",
			req: request{
				url:         fakeURL(fakeNonExistentId.String()),
				contentType: "application/json; charset=UTF-8",
				body:        strings.NewReader(fakeUpdateActorJSON),
			},
			want: statusResponse{
				status: http.StatusNotFound,
				body:   http.StatusText(http.StatusNotFound),
			},
			wantErr: false,
		},
		{
			name: "When id exists",
			req: request{
				url:         fakeURL(fakeExistentId.String()),
				contentType: "application/json; charset=UTF-8",
				body:        strings.NewReader(fakeUpdateActorJSON),
			},
			want: statusResponse{
				status: http.StatusOK,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client := &http.Client{}
			req, err := http.NewRequest(http.MethodPut, tt.req.url, tt.req.body)
			got, err := client.Do(req)
			if (err != nil) != tt.wantErr {
				require.Error(t, err)
			}
			if got != nil {
				defer got.Body.Close()
				require.Equal(t, tt.want.status, got.StatusCode, "they should be equal")
				switch tt.want.status {
				case http.StatusOK:
					bs, err := ioutil.ReadAll(got.Body)
					if err != nil {
						t.Fatalf("read body: %v", err)
					}
					actualBody := strings.TrimSpace(string(bs))
					require.Equal(t, tt.want.body, actualBody, "they should be equal")
				}
			}
		})
	}
}
