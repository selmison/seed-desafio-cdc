// Code generated by goa v3.2.4, DO NOT EDIT.
//
// catalog HTTP server types
//
// Command:
// $ goa gen github.com/selmison/seed-desafio-cdc/design

package server

import (
	"unicode/utf8"

	catalog "github.com/selmison/seed-desafio-cdc/gen/catalog"
	goa "goa.design/goa/v3/pkg"
)

// CreateActorRequestBody is the type of the "catalog" service "create_actor"
// endpoint HTTP request body.
type CreateActorRequestBody struct {
	Name        *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	Email       *string `form:"email,omitempty" json:"email,omitempty" xml:"email,omitempty"`
	Description *string `form:"description,omitempty" json:"description,omitempty" xml:"description,omitempty"`
}

// CreateBookRequestBody is the type of the "catalog" service "create_book"
// endpoint HTTP request body.
type CreateBookRequestBody struct {
	Title       *string  `form:"title,omitempty" json:"title,omitempty" xml:"title,omitempty"`
	Synopsis    *string  `form:"synopsis,omitempty" json:"synopsis,omitempty" xml:"synopsis,omitempty"`
	Summary     *string  `form:"summary,omitempty" json:"summary,omitempty" xml:"summary,omitempty"`
	Price       *float32 `form:"price,omitempty" json:"price,omitempty" xml:"price,omitempty"`
	Pages       *int     `form:"pages,omitempty" json:"pages,omitempty" xml:"pages,omitempty"`
	Isbn        *string  `form:"isbn,omitempty" json:"isbn,omitempty" xml:"isbn,omitempty"`
	Issue       *string  `form:"issue,omitempty" json:"issue,omitempty" xml:"issue,omitempty"`
	CategoryIds []string `form:"category_ids,omitempty" json:"category_ids,omitempty" xml:"category_ids,omitempty"`
	ActorIds    []string `form:"actor_ids,omitempty" json:"actor_ids,omitempty" xml:"actor_ids,omitempty"`
}

// CreateCategoryRequestBody is the type of the "catalog" service
// "create_category" endpoint HTTP request body.
type CreateCategoryRequestBody struct {
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
}

// CreateCustomerRequestBody is the type of the "catalog" service
// "create_customer" endpoint HTTP request body.
type CreateCustomerRequestBody struct {
	FirstName *string                `form:"first_name,omitempty" json:"first_name,omitempty" xml:"first_name,omitempty"`
	LastName  *string                `form:"last_name,omitempty" json:"last_name,omitempty" xml:"last_name,omitempty"`
	Email     *string                `form:"email,omitempty" json:"email,omitempty" xml:"email,omitempty"`
	Document  *string                `form:"document,omitempty" json:"document,omitempty" xml:"document,omitempty"`
	Address   *AddressDTORequestBody `form:"address,omitempty" json:"address,omitempty" xml:"address,omitempty"`
	Phone     *string                `form:"phone,omitempty" json:"phone,omitempty" xml:"phone,omitempty"`
}

// CreateCountryRequestBody is the type of the "catalog" service
// "create_country" endpoint HTTP request body.
type CreateCountryRequestBody struct {
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
}

// CreateStateRequestBody is the type of the "catalog" service "create_state"
// endpoint HTTP request body.
type CreateStateRequestBody struct {
	Name      *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	CountryID *string `form:"country_id,omitempty" json:"country_id,omitempty" xml:"country_id,omitempty"`
}

// CreateActorResponseBody is the type of the "catalog" service "create_actor"
// endpoint HTTP response body.
type CreateActorResponseBody struct {
	ID          string `form:"id" json:"id" xml:"id"`
	Name        string `form:"name" json:"name" xml:"name"`
	Email       string `form:"email" json:"email" xml:"email"`
	Description string `form:"description" json:"description" xml:"description"`
	CreatedAt   string `form:"created_at" json:"created_at" xml:"created_at"`
}

// ShowActorResponseBody is the type of the "catalog" service "show_actor"
// endpoint HTTP response body.
type ShowActorResponseBody struct {
	ID          string `form:"id" json:"id" xml:"id"`
	Name        string `form:"name" json:"name" xml:"name"`
	Email       string `form:"email" json:"email" xml:"email"`
	Description string `form:"description" json:"description" xml:"description"`
	CreatedAt   string `form:"created_at" json:"created_at" xml:"created_at"`
}

// CreateBookResponseBody is the type of the "catalog" service "create_book"
// endpoint HTTP response body.
type CreateBookResponseBody struct {
	ID          string   `form:"id" json:"id" xml:"id"`
	Title       string   `form:"title" json:"title" xml:"title"`
	Synopsis    string   `form:"synopsis" json:"synopsis" xml:"synopsis"`
	Summary     *string  `form:"summary,omitempty" json:"summary,omitempty" xml:"summary,omitempty"`
	Price       float32  `form:"price" json:"price" xml:"price"`
	Pages       int      `form:"pages" json:"pages" xml:"pages"`
	Isbn        string   `form:"isbn" json:"isbn" xml:"isbn"`
	Issue       string   `form:"issue" json:"issue" xml:"issue"`
	CategoryIds []string `form:"category_ids" json:"category_ids" xml:"category_ids"`
	ActorIds    []string `form:"actor_ids" json:"actor_ids" xml:"actor_ids"`
}

// ListBooksResponseBody is the type of the "catalog" service "list_books"
// endpoint HTTP response body.
type ListBooksResponseBody []*BookDTOResponse

// ShowBookResponseBody is the type of the "catalog" service "show_book"
// endpoint HTTP response body.
type ShowBookResponseBody struct {
	ID          string   `form:"id" json:"id" xml:"id"`
	Title       string   `form:"title" json:"title" xml:"title"`
	Synopsis    string   `form:"synopsis" json:"synopsis" xml:"synopsis"`
	Summary     *string  `form:"summary,omitempty" json:"summary,omitempty" xml:"summary,omitempty"`
	Price       float32  `form:"price" json:"price" xml:"price"`
	Pages       int      `form:"pages" json:"pages" xml:"pages"`
	Isbn        string   `form:"isbn" json:"isbn" xml:"isbn"`
	Issue       string   `form:"issue" json:"issue" xml:"issue"`
	CategoryIds []string `form:"category_ids" json:"category_ids" xml:"category_ids"`
	ActorIds    []string `form:"actor_ids" json:"actor_ids" xml:"actor_ids"`
}

// CreateCategoryResponseBody is the type of the "catalog" service
// "create_category" endpoint HTTP response body.
type CreateCategoryResponseBody struct {
	ID   string `form:"id" json:"id" xml:"id"`
	Name string `form:"name" json:"name" xml:"name"`
}

// ShowCategoryResponseBody is the type of the "catalog" service
// "show_category" endpoint HTTP response body.
type ShowCategoryResponseBody struct {
	ID   string `form:"id" json:"id" xml:"id"`
	Name string `form:"name" json:"name" xml:"name"`
}

// CreateCustomerResponseBody is the type of the "catalog" service
// "create_customer" endpoint HTTP response body.
type CreateCustomerResponseBody struct {
	ID        string                  `form:"id" json:"id" xml:"id"`
	FirstName string                  `form:"first_name" json:"first_name" xml:"first_name"`
	LastName  string                  `form:"last_name" json:"last_name" xml:"last_name"`
	Email     string                  `form:"email" json:"email" xml:"email"`
	Document  string                  `form:"document" json:"document" xml:"document"`
	Address   *AddressDTOResponseBody `form:"address" json:"address" xml:"address"`
	Phone     string                  `form:"phone" json:"phone" xml:"phone"`
}

// CreateCountryResponseBody is the type of the "catalog" service
// "create_country" endpoint HTTP response body.
type CreateCountryResponseBody struct {
	ID       string  `form:"id" json:"id" xml:"id"`
	Name     string  `form:"name" json:"name" xml:"name"`
	StateIds *string `form:"state_ids,omitempty" json:"state_ids,omitempty" xml:"state_ids,omitempty"`
}

// CreateStateResponseBody is the type of the "catalog" service "create_state"
// endpoint HTTP response body.
type CreateStateResponseBody struct {
	ID        string `form:"id" json:"id" xml:"id"`
	Name      string `form:"name" json:"name" xml:"name"`
	CountryID string `form:"country_id" json:"country_id" xml:"country_id"`
}

// CreateActorInvalidFieldsResponseBody is the type of the "catalog" service
// "create_actor" endpoint HTTP response body for the "invalid_fields" error.
type CreateActorInvalidFieldsResponseBody struct {
	// Name is the name of this class of errors.
	Name string `form:"name" json:"name" xml:"name"`
	// ID is a unique identifier for this particular occurrence of the problem.
	ID string `form:"id" json:"id" xml:"id"`
	// Message is a human-readable explanation specific to this occurrence of the
	// problem.
	Message string `form:"message" json:"message" xml:"message"`
	// Is the error temporary?
	Temporary bool `form:"temporary" json:"temporary" xml:"temporary"`
	// Is the error a timeout?
	Timeout bool `form:"timeout" json:"timeout" xml:"timeout"`
	// Is the error a server-side fault?
	Fault bool `form:"fault" json:"fault" xml:"fault"`
}

// ShowActorNotFoundResponseBody is the type of the "catalog" service
// "show_actor" endpoint HTTP response body for the "not_found" error.
type ShowActorNotFoundResponseBody struct {
	// Name is the name of this class of errors.
	Name string `form:"name" json:"name" xml:"name"`
	// ID is a unique identifier for this particular occurrence of the problem.
	ID string `form:"id" json:"id" xml:"id"`
	// Message is a human-readable explanation specific to this occurrence of the
	// problem.
	Message string `form:"message" json:"message" xml:"message"`
	// Is the error temporary?
	Temporary bool `form:"temporary" json:"temporary" xml:"temporary"`
	// Is the error a timeout?
	Timeout bool `form:"timeout" json:"timeout" xml:"timeout"`
	// Is the error a server-side fault?
	Fault bool `form:"fault" json:"fault" xml:"fault"`
}

// CreateBookInvalidFieldsResponseBody is the type of the "catalog" service
// "create_book" endpoint HTTP response body for the "invalid_fields" error.
type CreateBookInvalidFieldsResponseBody struct {
	// Name is the name of this class of errors.
	Name string `form:"name" json:"name" xml:"name"`
	// ID is a unique identifier for this particular occurrence of the problem.
	ID string `form:"id" json:"id" xml:"id"`
	// Message is a human-readable explanation specific to this occurrence of the
	// problem.
	Message string `form:"message" json:"message" xml:"message"`
	// Is the error temporary?
	Temporary bool `form:"temporary" json:"temporary" xml:"temporary"`
	// Is the error a timeout?
	Timeout bool `form:"timeout" json:"timeout" xml:"timeout"`
	// Is the error a server-side fault?
	Fault bool `form:"fault" json:"fault" xml:"fault"`
}

// ShowBookNotFoundResponseBody is the type of the "catalog" service
// "show_book" endpoint HTTP response body for the "not_found" error.
type ShowBookNotFoundResponseBody struct {
	// Name is the name of this class of errors.
	Name string `form:"name" json:"name" xml:"name"`
	// ID is a unique identifier for this particular occurrence of the problem.
	ID string `form:"id" json:"id" xml:"id"`
	// Message is a human-readable explanation specific to this occurrence of the
	// problem.
	Message string `form:"message" json:"message" xml:"message"`
	// Is the error temporary?
	Temporary bool `form:"temporary" json:"temporary" xml:"temporary"`
	// Is the error a timeout?
	Timeout bool `form:"timeout" json:"timeout" xml:"timeout"`
	// Is the error a server-side fault?
	Fault bool `form:"fault" json:"fault" xml:"fault"`
}

// CreateCategoryInvalidFieldsResponseBody is the type of the "catalog" service
// "create_category" endpoint HTTP response body for the "invalid_fields" error.
type CreateCategoryInvalidFieldsResponseBody struct {
	// Name is the name of this class of errors.
	Name string `form:"name" json:"name" xml:"name"`
	// ID is a unique identifier for this particular occurrence of the problem.
	ID string `form:"id" json:"id" xml:"id"`
	// Message is a human-readable explanation specific to this occurrence of the
	// problem.
	Message string `form:"message" json:"message" xml:"message"`
	// Is the error temporary?
	Temporary bool `form:"temporary" json:"temporary" xml:"temporary"`
	// Is the error a timeout?
	Timeout bool `form:"timeout" json:"timeout" xml:"timeout"`
	// Is the error a server-side fault?
	Fault bool `form:"fault" json:"fault" xml:"fault"`
}

// ShowCategoryNotFoundResponseBody is the type of the "catalog" service
// "show_category" endpoint HTTP response body for the "not_found" error.
type ShowCategoryNotFoundResponseBody struct {
	// Name is the name of this class of errors.
	Name string `form:"name" json:"name" xml:"name"`
	// ID is a unique identifier for this particular occurrence of the problem.
	ID string `form:"id" json:"id" xml:"id"`
	// Message is a human-readable explanation specific to this occurrence of the
	// problem.
	Message string `form:"message" json:"message" xml:"message"`
	// Is the error temporary?
	Temporary bool `form:"temporary" json:"temporary" xml:"temporary"`
	// Is the error a timeout?
	Timeout bool `form:"timeout" json:"timeout" xml:"timeout"`
	// Is the error a server-side fault?
	Fault bool `form:"fault" json:"fault" xml:"fault"`
}

// CreateCustomerInvalidFieldsResponseBody is the type of the "catalog" service
// "create_customer" endpoint HTTP response body for the "invalid_fields" error.
type CreateCustomerInvalidFieldsResponseBody struct {
	// Name is the name of this class of errors.
	Name string `form:"name" json:"name" xml:"name"`
	// ID is a unique identifier for this particular occurrence of the problem.
	ID string `form:"id" json:"id" xml:"id"`
	// Message is a human-readable explanation specific to this occurrence of the
	// problem.
	Message string `form:"message" json:"message" xml:"message"`
	// Is the error temporary?
	Temporary bool `form:"temporary" json:"temporary" xml:"temporary"`
	// Is the error a timeout?
	Timeout bool `form:"timeout" json:"timeout" xml:"timeout"`
	// Is the error a server-side fault?
	Fault bool `form:"fault" json:"fault" xml:"fault"`
}

// CreateCountryInvalidFieldsResponseBody is the type of the "catalog" service
// "create_country" endpoint HTTP response body for the "invalid_fields" error.
type CreateCountryInvalidFieldsResponseBody struct {
	// Name is the name of this class of errors.
	Name string `form:"name" json:"name" xml:"name"`
	// ID is a unique identifier for this particular occurrence of the problem.
	ID string `form:"id" json:"id" xml:"id"`
	// Message is a human-readable explanation specific to this occurrence of the
	// problem.
	Message string `form:"message" json:"message" xml:"message"`
	// Is the error temporary?
	Temporary bool `form:"temporary" json:"temporary" xml:"temporary"`
	// Is the error a timeout?
	Timeout bool `form:"timeout" json:"timeout" xml:"timeout"`
	// Is the error a server-side fault?
	Fault bool `form:"fault" json:"fault" xml:"fault"`
}

// CreateStateInvalidFieldsResponseBody is the type of the "catalog" service
// "create_state" endpoint HTTP response body for the "invalid_fields" error.
type CreateStateInvalidFieldsResponseBody struct {
	// Name is the name of this class of errors.
	Name string `form:"name" json:"name" xml:"name"`
	// ID is a unique identifier for this particular occurrence of the problem.
	ID string `form:"id" json:"id" xml:"id"`
	// Message is a human-readable explanation specific to this occurrence of the
	// problem.
	Message string `form:"message" json:"message" xml:"message"`
	// Is the error temporary?
	Temporary bool `form:"temporary" json:"temporary" xml:"temporary"`
	// Is the error a timeout?
	Timeout bool `form:"timeout" json:"timeout" xml:"timeout"`
	// Is the error a server-side fault?
	Fault bool `form:"fault" json:"fault" xml:"fault"`
}

// BookDTOResponse is used to define fields on response body types.
type BookDTOResponse struct {
	ID          string   `form:"id" json:"id" xml:"id"`
	Title       string   `form:"title" json:"title" xml:"title"`
	Synopsis    string   `form:"synopsis" json:"synopsis" xml:"synopsis"`
	Summary     *string  `form:"summary,omitempty" json:"summary,omitempty" xml:"summary,omitempty"`
	Price       float32  `form:"price" json:"price" xml:"price"`
	Pages       int      `form:"pages" json:"pages" xml:"pages"`
	Isbn        string   `form:"isbn" json:"isbn" xml:"isbn"`
	Issue       string   `form:"issue" json:"issue" xml:"issue"`
	CategoryIds []string `form:"category_ids" json:"category_ids" xml:"category_ids"`
	ActorIds    []string `form:"actor_ids" json:"actor_ids" xml:"actor_ids"`
}

// AddressDTOResponseBody is used to define fields on response body types.
type AddressDTOResponseBody struct {
	Address    string `form:"address" json:"address" xml:"address"`
	Complement string `form:"complement" json:"complement" xml:"complement"`
	City       string `form:"city" json:"city" xml:"city"`
	CountryID  string `form:"country_id" json:"country_id" xml:"country_id"`
	StateID    string `form:"state_id" json:"state_id" xml:"state_id"`
	Cep        string `form:"cep" json:"cep" xml:"cep"`
}

// AddressDTORequestBody is used to define fields on request body types.
type AddressDTORequestBody struct {
	Address    *string `form:"address,omitempty" json:"address,omitempty" xml:"address,omitempty"`
	Complement *string `form:"complement,omitempty" json:"complement,omitempty" xml:"complement,omitempty"`
	City       *string `form:"city,omitempty" json:"city,omitempty" xml:"city,omitempty"`
	CountryID  *string `form:"country_id,omitempty" json:"country_id,omitempty" xml:"country_id,omitempty"`
	StateID    *string `form:"state_id,omitempty" json:"state_id,omitempty" xml:"state_id,omitempty"`
	Cep        *string `form:"cep,omitempty" json:"cep,omitempty" xml:"cep,omitempty"`
}

// NewCreateActorResponseBody builds the HTTP response body from the result of
// the "create_actor" endpoint of the "catalog" service.
func NewCreateActorResponseBody(res *catalog.ActorDTO) *CreateActorResponseBody {
	body := &CreateActorResponseBody{
		ID:          res.ID,
		Name:        res.Name,
		Email:       res.Email,
		Description: res.Description,
		CreatedAt:   res.CreatedAt,
	}
	return body
}

// NewShowActorResponseBody builds the HTTP response body from the result of
// the "show_actor" endpoint of the "catalog" service.
func NewShowActorResponseBody(res *catalog.ActorDTO) *ShowActorResponseBody {
	body := &ShowActorResponseBody{
		ID:          res.ID,
		Name:        res.Name,
		Email:       res.Email,
		Description: res.Description,
		CreatedAt:   res.CreatedAt,
	}
	return body
}

// NewCreateBookResponseBody builds the HTTP response body from the result of
// the "create_book" endpoint of the "catalog" service.
func NewCreateBookResponseBody(res *catalog.BookDTO) *CreateBookResponseBody {
	body := &CreateBookResponseBody{
		ID:       res.ID,
		Title:    res.Title,
		Synopsis: res.Synopsis,
		Summary:  res.Summary,
		Price:    res.Price,
		Pages:    res.Pages,
		Isbn:     res.Isbn,
		Issue:    res.Issue,
	}
	if res.CategoryIds != nil {
		body.CategoryIds = make([]string, len(res.CategoryIds))
		for i, val := range res.CategoryIds {
			body.CategoryIds[i] = val
		}
	}
	if res.ActorIds != nil {
		body.ActorIds = make([]string, len(res.ActorIds))
		for i, val := range res.ActorIds {
			body.ActorIds[i] = val
		}
	}
	return body
}

// NewListBooksResponseBody builds the HTTP response body from the result of
// the "list_books" endpoint of the "catalog" service.
func NewListBooksResponseBody(res []*catalog.BookDTO) ListBooksResponseBody {
	body := make([]*BookDTOResponse, len(res))
	for i, val := range res {
		body[i] = marshalCatalogBookDTOToBookDTOResponse(val)
	}
	return body
}

// NewShowBookResponseBody builds the HTTP response body from the result of the
// "show_book" endpoint of the "catalog" service.
func NewShowBookResponseBody(res *catalog.BookDTO) *ShowBookResponseBody {
	body := &ShowBookResponseBody{
		ID:       res.ID,
		Title:    res.Title,
		Synopsis: res.Synopsis,
		Summary:  res.Summary,
		Price:    res.Price,
		Pages:    res.Pages,
		Isbn:     res.Isbn,
		Issue:    res.Issue,
	}
	if res.CategoryIds != nil {
		body.CategoryIds = make([]string, len(res.CategoryIds))
		for i, val := range res.CategoryIds {
			body.CategoryIds[i] = val
		}
	}
	if res.ActorIds != nil {
		body.ActorIds = make([]string, len(res.ActorIds))
		for i, val := range res.ActorIds {
			body.ActorIds[i] = val
		}
	}
	return body
}

// NewCreateCategoryResponseBody builds the HTTP response body from the result
// of the "create_category" endpoint of the "catalog" service.
func NewCreateCategoryResponseBody(res *catalog.CategoryDTO) *CreateCategoryResponseBody {
	body := &CreateCategoryResponseBody{
		ID:   res.ID,
		Name: res.Name,
	}
	return body
}

// NewShowCategoryResponseBody builds the HTTP response body from the result of
// the "show_category" endpoint of the "catalog" service.
func NewShowCategoryResponseBody(res *catalog.CategoryDTO) *ShowCategoryResponseBody {
	body := &ShowCategoryResponseBody{
		ID:   res.ID,
		Name: res.Name,
	}
	return body
}

// NewCreateCustomerResponseBody builds the HTTP response body from the result
// of the "create_customer" endpoint of the "catalog" service.
func NewCreateCustomerResponseBody(res *catalog.CustomerDTO) *CreateCustomerResponseBody {
	body := &CreateCustomerResponseBody{
		ID:        res.ID,
		FirstName: res.FirstName,
		LastName:  res.LastName,
		Email:     res.Email,
		Document:  res.Document,
		Phone:     res.Phone,
	}
	if res.Address != nil {
		body.Address = marshalCatalogAddressDTOToAddressDTOResponseBody(res.Address)
	}
	return body
}

// NewCreateCountryResponseBody builds the HTTP response body from the result
// of the "create_country" endpoint of the "catalog" service.
func NewCreateCountryResponseBody(res *catalog.CountryDTO) *CreateCountryResponseBody {
	body := &CreateCountryResponseBody{
		ID:       res.ID,
		Name:     res.Name,
		StateIds: res.StateIds,
	}
	return body
}

// NewCreateStateResponseBody builds the HTTP response body from the result of
// the "create_state" endpoint of the "catalog" service.
func NewCreateStateResponseBody(res *catalog.StateDTO) *CreateStateResponseBody {
	body := &CreateStateResponseBody{
		ID:        res.ID,
		Name:      res.Name,
		CountryID: res.CountryID,
	}
	return body
}

// NewCreateActorInvalidFieldsResponseBody builds the HTTP response body from
// the result of the "create_actor" endpoint of the "catalog" service.
func NewCreateActorInvalidFieldsResponseBody(res *goa.ServiceError) *CreateActorInvalidFieldsResponseBody {
	body := &CreateActorInvalidFieldsResponseBody{
		Name:      res.Name,
		ID:        res.ID,
		Message:   res.Message,
		Temporary: res.Temporary,
		Timeout:   res.Timeout,
		Fault:     res.Fault,
	}
	return body
}

// NewShowActorNotFoundResponseBody builds the HTTP response body from the
// result of the "show_actor" endpoint of the "catalog" service.
func NewShowActorNotFoundResponseBody(res *goa.ServiceError) *ShowActorNotFoundResponseBody {
	body := &ShowActorNotFoundResponseBody{
		Name:      res.Name,
		ID:        res.ID,
		Message:   res.Message,
		Temporary: res.Temporary,
		Timeout:   res.Timeout,
		Fault:     res.Fault,
	}
	return body
}

// NewCreateBookInvalidFieldsResponseBody builds the HTTP response body from
// the result of the "create_book" endpoint of the "catalog" service.
func NewCreateBookInvalidFieldsResponseBody(res *goa.ServiceError) *CreateBookInvalidFieldsResponseBody {
	body := &CreateBookInvalidFieldsResponseBody{
		Name:      res.Name,
		ID:        res.ID,
		Message:   res.Message,
		Temporary: res.Temporary,
		Timeout:   res.Timeout,
		Fault:     res.Fault,
	}
	return body
}

// NewShowBookNotFoundResponseBody builds the HTTP response body from the
// result of the "show_book" endpoint of the "catalog" service.
func NewShowBookNotFoundResponseBody(res *goa.ServiceError) *ShowBookNotFoundResponseBody {
	body := &ShowBookNotFoundResponseBody{
		Name:      res.Name,
		ID:        res.ID,
		Message:   res.Message,
		Temporary: res.Temporary,
		Timeout:   res.Timeout,
		Fault:     res.Fault,
	}
	return body
}

// NewCreateCategoryInvalidFieldsResponseBody builds the HTTP response body
// from the result of the "create_category" endpoint of the "catalog" service.
func NewCreateCategoryInvalidFieldsResponseBody(res *goa.ServiceError) *CreateCategoryInvalidFieldsResponseBody {
	body := &CreateCategoryInvalidFieldsResponseBody{
		Name:      res.Name,
		ID:        res.ID,
		Message:   res.Message,
		Temporary: res.Temporary,
		Timeout:   res.Timeout,
		Fault:     res.Fault,
	}
	return body
}

// NewShowCategoryNotFoundResponseBody builds the HTTP response body from the
// result of the "show_category" endpoint of the "catalog" service.
func NewShowCategoryNotFoundResponseBody(res *goa.ServiceError) *ShowCategoryNotFoundResponseBody {
	body := &ShowCategoryNotFoundResponseBody{
		Name:      res.Name,
		ID:        res.ID,
		Message:   res.Message,
		Temporary: res.Temporary,
		Timeout:   res.Timeout,
		Fault:     res.Fault,
	}
	return body
}

// NewCreateCustomerInvalidFieldsResponseBody builds the HTTP response body
// from the result of the "create_customer" endpoint of the "catalog" service.
func NewCreateCustomerInvalidFieldsResponseBody(res *goa.ServiceError) *CreateCustomerInvalidFieldsResponseBody {
	body := &CreateCustomerInvalidFieldsResponseBody{
		Name:      res.Name,
		ID:        res.ID,
		Message:   res.Message,
		Temporary: res.Temporary,
		Timeout:   res.Timeout,
		Fault:     res.Fault,
	}
	return body
}

// NewCreateCountryInvalidFieldsResponseBody builds the HTTP response body from
// the result of the "create_country" endpoint of the "catalog" service.
func NewCreateCountryInvalidFieldsResponseBody(res *goa.ServiceError) *CreateCountryInvalidFieldsResponseBody {
	body := &CreateCountryInvalidFieldsResponseBody{
		Name:      res.Name,
		ID:        res.ID,
		Message:   res.Message,
		Temporary: res.Temporary,
		Timeout:   res.Timeout,
		Fault:     res.Fault,
	}
	return body
}

// NewCreateStateInvalidFieldsResponseBody builds the HTTP response body from
// the result of the "create_state" endpoint of the "catalog" service.
func NewCreateStateInvalidFieldsResponseBody(res *goa.ServiceError) *CreateStateInvalidFieldsResponseBody {
	body := &CreateStateInvalidFieldsResponseBody{
		Name:      res.Name,
		ID:        res.ID,
		Message:   res.Message,
		Temporary: res.Temporary,
		Timeout:   res.Timeout,
		Fault:     res.Fault,
	}
	return body
}

// NewCreateActorDTO builds a catalog service create_actor endpoint payload.
func NewCreateActorDTO(body *CreateActorRequestBody) *catalog.CreateActorDTO {
	v := &catalog.CreateActorDTO{
		Name:        *body.Name,
		Email:       *body.Email,
		Description: *body.Description,
	}

	return v
}

// NewShowActorShowByIDDTO builds a catalog service show_actor endpoint payload.
func NewShowActorShowByIDDTO(id string) *catalog.ShowByIDDTO {
	v := &catalog.ShowByIDDTO{}
	v.ID = id

	return v
}

// NewCreateBookDTO builds a catalog service create_book endpoint payload.
func NewCreateBookDTO(body *CreateBookRequestBody) *catalog.CreateBookDTO {
	v := &catalog.CreateBookDTO{
		Title:    *body.Title,
		Synopsis: *body.Synopsis,
		Summary:  body.Summary,
		Price:    *body.Price,
		Pages:    *body.Pages,
		Isbn:     *body.Isbn,
		Issue:    *body.Issue,
	}
	v.CategoryIds = make([]string, len(body.CategoryIds))
	for i, val := range body.CategoryIds {
		v.CategoryIds[i] = val
	}
	v.ActorIds = make([]string, len(body.ActorIds))
	for i, val := range body.ActorIds {
		v.ActorIds[i] = val
	}

	return v
}

// NewShowBookShowByIDDTO builds a catalog service show_book endpoint payload.
func NewShowBookShowByIDDTO(id string) *catalog.ShowByIDDTO {
	v := &catalog.ShowByIDDTO{}
	v.ID = id

	return v
}

// NewCreateCategoryDTO builds a catalog service create_category endpoint
// payload.
func NewCreateCategoryDTO(body *CreateCategoryRequestBody) *catalog.CreateCategoryDTO {
	v := &catalog.CreateCategoryDTO{
		Name: *body.Name,
	}

	return v
}

// NewShowCategoryShowByIDDTO builds a catalog service show_category endpoint
// payload.
func NewShowCategoryShowByIDDTO(id string) *catalog.ShowByIDDTO {
	v := &catalog.ShowByIDDTO{}
	v.ID = id

	return v
}

// NewCreateCustomerDTO builds a catalog service create_customer endpoint
// payload.
func NewCreateCustomerDTO(body *CreateCustomerRequestBody) *catalog.CreateCustomerDTO {
	v := &catalog.CreateCustomerDTO{
		FirstName: *body.FirstName,
		LastName:  *body.LastName,
		Email:     *body.Email,
		Document:  *body.Document,
		Phone:     *body.Phone,
	}
	v.Address = unmarshalAddressDTORequestBodyToCatalogAddressDTO(body.Address)

	return v
}

// NewCreateCountryDTO builds a catalog service create_country endpoint payload.
func NewCreateCountryDTO(body *CreateCountryRequestBody) *catalog.CreateCountryDTO {
	v := &catalog.CreateCountryDTO{
		Name: *body.Name,
	}

	return v
}

// NewCreateStateDTO builds a catalog service create_state endpoint payload.
func NewCreateStateDTO(body *CreateStateRequestBody) *catalog.CreateStateDTO {
	v := &catalog.CreateStateDTO{
		Name:      *body.Name,
		CountryID: *body.CountryID,
	}

	return v
}

// ValidateCreateActorRequestBody runs the validations defined on
// create_actor_request_body
func ValidateCreateActorRequestBody(body *CreateActorRequestBody) (err error) {
	if body.Name == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("name", "body"))
	}
	if body.Email == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("email", "body"))
	}
	if body.Description == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("description", "body"))
	}
	if body.Email != nil {
		err = goa.MergeErrors(err, goa.ValidateFormat("body.email", *body.Email, goa.FormatEmail))
	}
	if body.Description != nil {
		if utf8.RuneCountInString(*body.Description) > 400 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("body.description", *body.Description, utf8.RuneCountInString(*body.Description), 400, false))
		}
	}
	return
}

// ValidateCreateBookRequestBody runs the validations defined on
// create_book_request_body
func ValidateCreateBookRequestBody(body *CreateBookRequestBody) (err error) {
	if body.Title == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("title", "body"))
	}
	if body.Synopsis == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("synopsis", "body"))
	}
	if body.Price == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("price", "body"))
	}
	if body.Pages == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("pages", "body"))
	}
	if body.Isbn == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("isbn", "body"))
	}
	if body.Issue == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("issue", "body"))
	}
	if body.CategoryIds == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("category_ids", "body"))
	}
	if body.ActorIds == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("actor_ids", "body"))
	}
	if body.Synopsis != nil {
		if utf8.RuneCountInString(*body.Synopsis) > 500 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("body.synopsis", *body.Synopsis, utf8.RuneCountInString(*body.Synopsis), 500, false))
		}
	}
	if body.Price != nil {
		if *body.Price < 20 {
			err = goa.MergeErrors(err, goa.InvalidRangeError("body.price", *body.Price, 20, true))
		}
	}
	if body.Pages != nil {
		if *body.Pages < 100 {
			err = goa.MergeErrors(err, goa.InvalidRangeError("body.pages", *body.Pages, 100, true))
		}
	}
	return
}

// ValidateCreateCategoryRequestBody runs the validations defined on
// create_category_request_body
func ValidateCreateCategoryRequestBody(body *CreateCategoryRequestBody) (err error) {
	if body.Name == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("name", "body"))
	}
	return
}

// ValidateCreateCustomerRequestBody runs the validations defined on
// create_customer_request_body
func ValidateCreateCustomerRequestBody(body *CreateCustomerRequestBody) (err error) {
	if body.FirstName == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("first_name", "body"))
	}
	if body.LastName == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("last_name", "body"))
	}
	if body.Email == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("email", "body"))
	}
	if body.Document == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("document", "body"))
	}
	if body.Address == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("address", "body"))
	}
	if body.Phone == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("phone", "body"))
	}
	if body.Email != nil {
		err = goa.MergeErrors(err, goa.ValidateFormat("body.email", *body.Email, goa.FormatEmail))
	}
	if body.Address != nil {
		if err2 := ValidateAddressDTORequestBody(body.Address); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	return
}

// ValidateCreateCountryRequestBody runs the validations defined on
// create_country_request_body
func ValidateCreateCountryRequestBody(body *CreateCountryRequestBody) (err error) {
	if body.Name == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("name", "body"))
	}
	return
}

// ValidateCreateStateRequestBody runs the validations defined on
// create_state_request_body
func ValidateCreateStateRequestBody(body *CreateStateRequestBody) (err error) {
	if body.Name == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("name", "body"))
	}
	if body.CountryID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("country_id", "body"))
	}
	return
}

// ValidateAddressDTORequestBody runs the validations defined on
// AddressDTORequestBody
func ValidateAddressDTORequestBody(body *AddressDTORequestBody) (err error) {
	if body.Address == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("address", "body"))
	}
	if body.Complement == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("complement", "body"))
	}
	if body.City == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("city", "body"))
	}
	if body.CountryID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("country_id", "body"))
	}
	if body.StateID == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("state_id", "body"))
	}
	if body.Cep == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("cep", "body"))
	}
	return
}