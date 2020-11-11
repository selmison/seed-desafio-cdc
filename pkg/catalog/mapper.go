package catalog

import (
	"strings"

	"github.com/google/uuid"

	catalogGen "github.com/selmison/seed-desafio-cdc/gen/catalog"
)

func mapAddressDTOToAddress(dto catalogGen.AddressDTO, customerID string) Address {
	return Address{
		Address:    dto.Address,
		Complement: dto.Complement,
		City:       dto.City,
		StateID:    dto.StateID,
		State:      State{ID: dto.StateID},
		Cep:        dto.Cep,
	}
}

func mapCreateBookDTOToBook(dto catalogGen.CreateBookDTO) Book {
	var summary string
	if dto.Summary != nil {
		summary = strings.TrimSpace(*dto.Summary)
	}
	actors := mapIDsToActors(dto.ActorIds)
	categories := mapIDsToCategories(dto.CategoryIds)
	book := Book{
		ID:         uuid.New().String(),
		Title:      strings.TrimSpace(dto.Title),
		Synopsis:   strings.TrimSpace(dto.Synopsis),
		Summary:    summary,
		Price:      dto.Price,
		Pages:      dto.Pages,
		Isbn:       strings.TrimSpace(dto.Isbn),
		Issue:      strings.TrimSpace(dto.Issue),
		Categories: categories,
		Actors:     actors,
	}
	return book
}

func mapCreateCartDTOToCart(dto catalogGen.CreateCartDTO) Cart {
	id := uuid.New().String()
	return Cart{
		ID:    id,
		Items: mapItemsDTOToItems(dto.Items, id),
	}
}

func mapCreateCustomerDTOToCustomer(dto catalogGen.CreateCustomerDTO) Customer {
	id := uuid.New().String()
	customer := Customer{
		ID:        id,
		FirstName: dto.FirstName,
		LastName:  dto.LastName,
		Email:     dto.Email,
		Document:  dto.Document,
		Address:   mapAddressDTOToAddress(*dto.Address, id),
		Phone:     dto.Phone,
	}
	return customer
}

func mapItemsDTOToItems(dtos []*catalogGen.ItemDTO, cartID string) []*Item {
	items := make([]*Item, len(dtos))
	for i, dto := range dtos {
		items[i] = &Item{
			BookID: dto.BookID,
			Amount: dto.Amount,
			CartID: cartID,
		}
	}
	return items
}

func (s *service) mapBookToBookDTO(book Book) (*catalogGen.BookDTO, error) {
	var actors []*Actor
	if err := s.repo.Model(&book).Association("Actors").Find(&actors); err != nil {
		return nil, err
	}
	var categories []*Category
	if err := s.repo.Model(&book).Association("Categories").Find(&categories); err != nil {
		return nil, err
	}
	actorIDs := mapActorsToIDs(actors)
	categoryIDs := mapCategoriesToIDs(categories)
	return &catalogGen.BookDTO{
		ID:          book.ID,
		Title:       book.Title,
		Synopsis:    book.Synopsis,
		Summary:     &book.Summary,
		Price:       book.Price,
		Pages:       book.Pages,
		Isbn:        book.Isbn,
		Issue:       book.Issue,
		CategoryIds: categoryIDs,
		ActorIds:    actorIDs,
	}, nil
}

func mapCustomerToCustomerDTO(customer Customer) *catalogGen.CustomerDTO {
	return &catalogGen.CustomerDTO{
		ID:        customer.ID,
		FirstName: customer.FirstName,
		LastName:  customer.LastName,
		Email:     customer.Email,
		Document:  customer.Document,
		Address: &catalogGen.AddressDTO{
			Address:    customer.Address.Address,
			Complement: customer.Address.Complement,
			City:       customer.Address.City,
			StateID:    customer.Address.StateID,
			Cep:        customer.Address.Cep,
		},
		Phone: customer.Phone,
	}
}

func mapCartToCartDTO(cart Cart) *catalogGen.CartDTO {
	itemDTOs := make([]*catalogGen.ItemDTO, len(cart.Items))
	for i, item := range cart.Items {
		itemDTOs[i] = &catalogGen.ItemDTO{
			BookID: item.BookID,
			Amount: item.Amount,
		}
	}
	return &catalogGen.CartDTO{
		ID:         cart.ID,
		Total:      cart.Total(),
		Items:      itemDTOs,
		CustomerID: "",
		CouponID:   nil,
	}
}

func mapIDsToActors(ids []string) []*Actor {
	actors := make([]*Actor, len(ids))
	for i, id := range ids {
		actors[i] = &Actor{ID: id}
	}
	return actors
}

func mapIDsToCarts(ids []string) []*Cart {
	carts := make([]*Cart, len(ids))
	for i, id := range ids {
		carts[i] = &Cart{ID: id}
	}
	return carts
}

func mapIDsToCategories(ids []string) []*Category {
	categories := make([]*Category, len(ids))
	for i, id := range ids {
		categories[i] = &Category{ID: id}
	}
	return categories
}

func mapActorsToIDs(actors []*Actor) []string {
	ids := make([]string, len(actors))
	for i, actor := range actors {
		ids[i] = actor.ID
	}
	return ids
}

func mapCartsToIDs(carts []*Cart) []string {
	ids := make([]string, len(carts))
	for i, cart := range carts {
		ids[i] = cart.ID
	}
	return ids
}

func mapCategoriesToIDs(categories []*Category) []string {
	ids := make([]string, len(categories))
	for i, category := range categories {
		ids[i] = category.ID
	}
	return ids
}
