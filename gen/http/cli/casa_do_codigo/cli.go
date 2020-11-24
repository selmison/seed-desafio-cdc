// Code generated by goa v3.2.4, DO NOT EDIT.
//
// casa_do_codigo HTTP client CLI support package
//
// Command:
// $ goa gen github.com/selmison/seed-desafio-cdc/design

package cli

import (
	"flag"
	"fmt"
	"net/http"
	"os"

	"github.com/go-kit/kit/endpoint"
	catalogc "github.com/selmison/seed-desafio-cdc/gen/http/catalog/client"
	goahttp "goa.design/goa/v3/http"
)

// UsageCommands returns the set of commands and sub-commands using the format
//
//    command (subcommand1|subcommand2|...)
//
func UsageCommands() string {
	return `catalog (create-actor|list-actors|show-actor|create-book|list-books|show-book|create-cart|create-category|list-categories|show-category|create-country|list-countries|show-country|create-coupon|create-customer|create-purchase|create-state|list-states|show-state)
`
}

// UsageExamples produces an example of a valid invocation of the CLI tool.
func UsageExamples() string {
	return os.Args[0] + ` catalog create-actor --body '{
      "description": "ij9",
      "email": "xzavier.dickens@jacobsonernser.info",
      "name": "Minima non consequatur corporis molestias natus."
   }'` + "\n" +
		""
}

// ParseEndpoint returns the endpoint and payload as specified on the command
// line.
func ParseEndpoint(
	scheme, host string,
	doer goahttp.Doer,
	enc func(*http.Request) goahttp.Encoder,
	dec func(*http.Response) goahttp.Decoder,
	restore bool,
) (endpoint.Endpoint, interface{}, error) {
	var (
		catalogFlags = flag.NewFlagSet("catalog", flag.ContinueOnError)

		catalogCreateActorFlags    = flag.NewFlagSet("create-actor", flag.ExitOnError)
		catalogCreateActorBodyFlag = catalogCreateActorFlags.String("body", "REQUIRED", "")

		catalogListActorsFlags = flag.NewFlagSet("list-actors", flag.ExitOnError)

		catalogShowActorFlags  = flag.NewFlagSet("show-actor", flag.ExitOnError)
		catalogShowActorIDFlag = catalogShowActorFlags.String("id", "REQUIRED", "ID")

		catalogCreateBookFlags    = flag.NewFlagSet("create-book", flag.ExitOnError)
		catalogCreateBookBodyFlag = catalogCreateBookFlags.String("body", "REQUIRED", "")

		catalogListBooksFlags = flag.NewFlagSet("list-books", flag.ExitOnError)

		catalogShowBookFlags  = flag.NewFlagSet("show-book", flag.ExitOnError)
		catalogShowBookIDFlag = catalogShowBookFlags.String("id", "REQUIRED", "ID")

		catalogCreateCartFlags    = flag.NewFlagSet("create-cart", flag.ExitOnError)
		catalogCreateCartBodyFlag = catalogCreateCartFlags.String("body", "REQUIRED", "")

		catalogCreateCategoryFlags    = flag.NewFlagSet("create-category", flag.ExitOnError)
		catalogCreateCategoryBodyFlag = catalogCreateCategoryFlags.String("body", "REQUIRED", "")

		catalogListCategoriesFlags = flag.NewFlagSet("list-categories", flag.ExitOnError)

		catalogShowCategoryFlags  = flag.NewFlagSet("show-category", flag.ExitOnError)
		catalogShowCategoryIDFlag = catalogShowCategoryFlags.String("id", "REQUIRED", "ID")

		catalogCreateCountryFlags    = flag.NewFlagSet("create-country", flag.ExitOnError)
		catalogCreateCountryBodyFlag = catalogCreateCountryFlags.String("body", "REQUIRED", "")

		catalogListCountriesFlags = flag.NewFlagSet("list-countries", flag.ExitOnError)

		catalogShowCountryFlags  = flag.NewFlagSet("show-country", flag.ExitOnError)
		catalogShowCountryIDFlag = catalogShowCountryFlags.String("id", "REQUIRED", "ID")

		catalogCreateCouponFlags    = flag.NewFlagSet("create-coupon", flag.ExitOnError)
		catalogCreateCouponBodyFlag = catalogCreateCouponFlags.String("body", "REQUIRED", "")

		catalogCreateCustomerFlags    = flag.NewFlagSet("create-customer", flag.ExitOnError)
		catalogCreateCustomerBodyFlag = catalogCreateCustomerFlags.String("body", "REQUIRED", "")

		catalogCreatePurchaseFlags    = flag.NewFlagSet("create-purchase", flag.ExitOnError)
		catalogCreatePurchaseBodyFlag = catalogCreatePurchaseFlags.String("body", "REQUIRED", "")

		catalogCreateStateFlags    = flag.NewFlagSet("create-state", flag.ExitOnError)
		catalogCreateStateBodyFlag = catalogCreateStateFlags.String("body", "REQUIRED", "")

		catalogListStatesFlags = flag.NewFlagSet("list-states", flag.ExitOnError)

		catalogShowStateFlags  = flag.NewFlagSet("show-state", flag.ExitOnError)
		catalogShowStateIDFlag = catalogShowStateFlags.String("id", "REQUIRED", "ID")
	)
	catalogFlags.Usage = catalogUsage
	catalogCreateActorFlags.Usage = catalogCreateActorUsage
	catalogListActorsFlags.Usage = catalogListActorsUsage
	catalogShowActorFlags.Usage = catalogShowActorUsage
	catalogCreateBookFlags.Usage = catalogCreateBookUsage
	catalogListBooksFlags.Usage = catalogListBooksUsage
	catalogShowBookFlags.Usage = catalogShowBookUsage
	catalogCreateCartFlags.Usage = catalogCreateCartUsage
	catalogCreateCategoryFlags.Usage = catalogCreateCategoryUsage
	catalogListCategoriesFlags.Usage = catalogListCategoriesUsage
	catalogShowCategoryFlags.Usage = catalogShowCategoryUsage
	catalogCreateCountryFlags.Usage = catalogCreateCountryUsage
	catalogListCountriesFlags.Usage = catalogListCountriesUsage
	catalogShowCountryFlags.Usage = catalogShowCountryUsage
	catalogCreateCouponFlags.Usage = catalogCreateCouponUsage
	catalogCreateCustomerFlags.Usage = catalogCreateCustomerUsage
	catalogCreatePurchaseFlags.Usage = catalogCreatePurchaseUsage
	catalogCreateStateFlags.Usage = catalogCreateStateUsage
	catalogListStatesFlags.Usage = catalogListStatesUsage
	catalogShowStateFlags.Usage = catalogShowStateUsage

	if err := flag.CommandLine.Parse(os.Args[1:]); err != nil {
		return nil, nil, err
	}

	if flag.NArg() < 2 { // two non flag args are required: SERVICE and ENDPOINT (aka COMMAND)
		return nil, nil, fmt.Errorf("not enough arguments")
	}

	var (
		svcn string
		svcf *flag.FlagSet
	)
	{
		svcn = flag.Arg(0)
		switch svcn {
		case "catalog":
			svcf = catalogFlags
		default:
			return nil, nil, fmt.Errorf("unknown service %q", svcn)
		}
	}
	if err := svcf.Parse(flag.Args()[1:]); err != nil {
		return nil, nil, err
	}

	var (
		epn string
		epf *flag.FlagSet
	)
	{
		epn = svcf.Arg(0)
		switch svcn {
		case "catalog":
			switch epn {
			case "create-actor":
				epf = catalogCreateActorFlags

			case "list-actors":
				epf = catalogListActorsFlags

			case "show-actor":
				epf = catalogShowActorFlags

			case "create-book":
				epf = catalogCreateBookFlags

			case "list-books":
				epf = catalogListBooksFlags

			case "show-book":
				epf = catalogShowBookFlags

			case "create-cart":
				epf = catalogCreateCartFlags

			case "create-category":
				epf = catalogCreateCategoryFlags

			case "list-categories":
				epf = catalogListCategoriesFlags

			case "show-category":
				epf = catalogShowCategoryFlags

			case "create-country":
				epf = catalogCreateCountryFlags

			case "list-countries":
				epf = catalogListCountriesFlags

			case "show-country":
				epf = catalogShowCountryFlags

			case "create-coupon":
				epf = catalogCreateCouponFlags

			case "create-customer":
				epf = catalogCreateCustomerFlags

			case "create-purchase":
				epf = catalogCreatePurchaseFlags

			case "create-state":
				epf = catalogCreateStateFlags

			case "list-states":
				epf = catalogListStatesFlags

			case "show-state":
				epf = catalogShowStateFlags

			}

		}
	}
	if epf == nil {
		return nil, nil, fmt.Errorf("unknown %q endpoint %q", svcn, epn)
	}

	// Parse endpoint flags if any
	if svcf.NArg() > 1 {
		if err := epf.Parse(svcf.Args()[1:]); err != nil {
			return nil, nil, err
		}
	}

	var (
		data     interface{}
		endpoint endpoint.Endpoint
		err      error
	)
	{
		switch svcn {
		case "catalog":
			c := catalogc.NewClient(scheme, host, doer, enc, dec, restore)
			switch epn {
			case "create-actor":
				endpoint = c.CreateActor()
				data, err = catalogc.BuildCreateActorPayload(*catalogCreateActorBodyFlag)
			case "list-actors":
				endpoint = c.ListActors()
				data = nil
			case "show-actor":
				endpoint = c.ShowActor()
				data, err = catalogc.BuildShowActorPayload(*catalogShowActorIDFlag)
			case "create-book":
				endpoint = c.CreateBook()
				data, err = catalogc.BuildCreateBookPayload(*catalogCreateBookBodyFlag)
			case "list-books":
				endpoint = c.ListBooks()
				data = nil
			case "show-book":
				endpoint = c.ShowBook()
				data, err = catalogc.BuildShowBookPayload(*catalogShowBookIDFlag)
			case "create-cart":
				endpoint = c.CreateCart()
				data, err = catalogc.BuildCreateCartPayload(*catalogCreateCartBodyFlag)
			case "create-category":
				endpoint = c.CreateCategory()
				data, err = catalogc.BuildCreateCategoryPayload(*catalogCreateCategoryBodyFlag)
			case "list-categories":
				endpoint = c.ListCategories()
				data = nil
			case "show-category":
				endpoint = c.ShowCategory()
				data, err = catalogc.BuildShowCategoryPayload(*catalogShowCategoryIDFlag)
			case "create-country":
				endpoint = c.CreateCountry()
				data, err = catalogc.BuildCreateCountryPayload(*catalogCreateCountryBodyFlag)
			case "list-countries":
				endpoint = c.ListCountries()
				data = nil
			case "show-country":
				endpoint = c.ShowCountry()
				data, err = catalogc.BuildShowCountryPayload(*catalogShowCountryIDFlag)
			case "create-coupon":
				endpoint = c.CreateCoupon()
				data, err = catalogc.BuildCreateCouponPayload(*catalogCreateCouponBodyFlag)
			case "create-customer":
				endpoint = c.CreateCustomer()
				data, err = catalogc.BuildCreateCustomerPayload(*catalogCreateCustomerBodyFlag)
			case "create-purchase":
				endpoint = c.CreatePurchase()
				data, err = catalogc.BuildCreatePurchasePayload(*catalogCreatePurchaseBodyFlag)
			case "create-state":
				endpoint = c.CreateState()
				data, err = catalogc.BuildCreateStatePayload(*catalogCreateStateBodyFlag)
			case "list-states":
				endpoint = c.ListStates()
				data = nil
			case "show-state":
				endpoint = c.ShowState()
				data, err = catalogc.BuildShowStatePayload(*catalogShowStateIDFlag)
			}
		}
	}
	if err != nil {
		return nil, nil, err
	}

	return endpoint, data, nil
}

// catalogUsage displays the usage of the catalog command and its subcommands.
func catalogUsage() {
	fmt.Fprintf(os.Stderr, `The catalog service performs operations on catalog
Usage:
    %s [globalflags] catalog COMMAND [flags]

COMMAND:
    create-actor: CreateActor implements create_actor.
    list-actors: ListActors implements list_actors.
    show-actor: ShowActor implements show_actor.
    create-book: CreateBook implements create_book.
    list-books: ListBooks implements list_books.
    show-book: ShowBook implements show_book.
    create-cart: CreateCart implements create_cart.
    create-category: CreateCategory implements create_category.
    list-categories: ListCategories implements list_categories.
    show-category: ShowCategory implements show_category.
    create-country: CreateCountry implements create_country.
    list-countries: ListCountries implements list_countries.
    show-country: ShowCountry implements show_country.
    create-coupon: CreateCoupon implements create_coupon.
    create-customer: CreateCustomer implements create_customer.
    create-purchase: CreatePurchase implements create_purchase.
    create-state: CreateState implements create_state.
    list-states: ListStates implements list_states.
    show-state: ShowState implements show_state.

Additional help:
    %s catalog COMMAND --help
`, os.Args[0], os.Args[0])
}
func catalogCreateActorUsage() {
	fmt.Fprintf(os.Stderr, `%s [flags] catalog create-actor -body JSON

CreateActor implements create_actor.
    -body JSON: 

Example:
    `+os.Args[0]+` catalog create-actor --body '{
      "description": "ij9",
      "email": "xzavier.dickens@jacobsonernser.info",
      "name": "Minima non consequatur corporis molestias natus."
   }'
`, os.Args[0])
}

func catalogListActorsUsage() {
	fmt.Fprintf(os.Stderr, `%s [flags] catalog list-actors

ListActors implements list_actors.

Example:
    `+os.Args[0]+` catalog list-actors
`, os.Args[0])
}

func catalogShowActorUsage() {
	fmt.Fprintf(os.Stderr, `%s [flags] catalog show-actor -id STRING

ShowActor implements show_actor.
    -id STRING: ID

Example:
    `+os.Args[0]+` catalog show-actor --id "Provident inventore sit sed assumenda accusantium."
`, os.Args[0])
}

func catalogCreateBookUsage() {
	fmt.Fprintf(os.Stderr, `%s [flags] catalog create-book -body JSON

CreateBook implements create_book.
    -body JSON: 

Example:
    `+os.Args[0]+` catalog create-book --body '{
      "actor_ids": [
         "Fugit ut praesentium.",
         "Hic est soluta.",
         "Et accusantium ipsa commodi corporis excepturi."
      ],
      "category_ids": [
         "Quae perferendis adipisci excepturi.",
         "Qui repellat dolor delectus.",
         "Distinctio iusto corporis eum amet.",
         "Tempora alias quis."
      ],
      "isbn": "Quia distinctio.",
      "issue": "Quasi harum nam.",
      "pages": 6926372662986759706,
      "price": 20.820534,
      "summary": "Ullam facere minima cupiditate consectetur.",
      "synopsis": "nbr",
      "title": "Officiis accusamus quibusdam aut eum."
   }'
`, os.Args[0])
}

func catalogListBooksUsage() {
	fmt.Fprintf(os.Stderr, `%s [flags] catalog list-books

ListBooks implements list_books.

Example:
    `+os.Args[0]+` catalog list-books
`, os.Args[0])
}

func catalogShowBookUsage() {
	fmt.Fprintf(os.Stderr, `%s [flags] catalog show-book -id STRING

ShowBook implements show_book.
    -id STRING: ID

Example:
    `+os.Args[0]+` catalog show-book --id "Rem quasi voluptatum ut."
`, os.Args[0])
}

func catalogCreateCartUsage() {
	fmt.Fprintf(os.Stderr, `%s [flags] catalog create-cart -body JSON

CreateCart implements create_cart.
    -body JSON: 

Example:
    `+os.Args[0]+` catalog create-cart --body '{
      "coupon_id": "Amet consequatur.",
      "customer_id": "Sed maiores qui incidunt rem rem quia.",
      "items": [
         {
            "amount": 295474086,
            "book_id": "Eos accusamus."
         },
         {
            "amount": 295474086,
            "book_id": "Eos accusamus."
         }
      ],
      "total": 0.9235589
   }'
`, os.Args[0])
}

func catalogCreateCategoryUsage() {
	fmt.Fprintf(os.Stderr, `%s [flags] catalog create-category -body JSON

CreateCategory implements create_category.
    -body JSON: 

Example:
    `+os.Args[0]+` catalog create-category --body '{
      "name": "Itaque nesciunt assumenda aperiam excepturi harum enim."
   }'
`, os.Args[0])
}

func catalogListCategoriesUsage() {
	fmt.Fprintf(os.Stderr, `%s [flags] catalog list-categories

ListCategories implements list_categories.

Example:
    `+os.Args[0]+` catalog list-categories
`, os.Args[0])
}

func catalogShowCategoryUsage() {
	fmt.Fprintf(os.Stderr, `%s [flags] catalog show-category -id STRING

ShowCategory implements show_category.
    -id STRING: ID

Example:
    `+os.Args[0]+` catalog show-category --id "Placeat aut hic iure voluptatem."
`, os.Args[0])
}

func catalogCreateCountryUsage() {
	fmt.Fprintf(os.Stderr, `%s [flags] catalog create-country -body JSON

CreateCountry implements create_country.
    -body JSON: 

Example:
    `+os.Args[0]+` catalog create-country --body '{
      "name": "Doloribus rerum corporis nostrum veniam."
   }'
`, os.Args[0])
}

func catalogListCountriesUsage() {
	fmt.Fprintf(os.Stderr, `%s [flags] catalog list-countries

ListCountries implements list_countries.

Example:
    `+os.Args[0]+` catalog list-countries
`, os.Args[0])
}

func catalogShowCountryUsage() {
	fmt.Fprintf(os.Stderr, `%s [flags] catalog show-country -id STRING

ShowCountry implements show_country.
    -id STRING: ID

Example:
    `+os.Args[0]+` catalog show-country --id "Recusandae non doloremque molestiae."
`, os.Args[0])
}

func catalogCreateCouponUsage() {
	fmt.Fprintf(os.Stderr, `%s [flags] catalog create-coupon -body JSON

CreateCoupon implements create_coupon.
    -body JSON: 

Example:
    `+os.Args[0]+` catalog create-coupon --body '{
      "code": "Et rem distinctio et dolore enim.",
      "discount": 0.45933485,
      "validity": "Nobis architecto voluptatem adipisci quas aut."
   }'
`, os.Args[0])
}

func catalogCreateCustomerUsage() {
	fmt.Fprintf(os.Stderr, `%s [flags] catalog create-customer -body JSON

CreateCustomer implements create_customer.
    -body JSON: 

Example:
    `+os.Args[0]+` catalog create-customer --body '{
      "address": {
         "address": "Voluptates quia non.",
         "cep": "Itaque ipsa et quae dolorum.",
         "city": "Asperiores et inventore asperiores doloribus commodi.",
         "complement": "Totam dolor.",
         "state_id": "Ad veniam natus eaque recusandae."
      },
      "cart_ids": [
         "Omnis facere veniam porro.",
         "Animi sed eos iure.",
         "Quaerat corporis harum praesentium consequatur dolor quam."
      ],
      "document": "Minus enim possimus eligendi est recusandae sed.",
      "email": "bessie_kub@spinka.info",
      "first_name": "Dicta ut.",
      "last_name": "Et eius dolore quia deserunt expedita.",
      "phone": "Ut ex ex."
   }'
`, os.Args[0])
}

func catalogCreatePurchaseUsage() {
	fmt.Fprintf(os.Stderr, `%s [flags] catalog create-purchase -body JSON

CreatePurchase implements create_purchase.
    -body JSON: 

Example:
    `+os.Args[0]+` catalog create-purchase --body '{
      "cart": {
         "coupon_id": "Minima ullam occaecati qui.",
         "customer_id": "Qui aut maxime asperiores et adipisci.",
         "items": [
            {
               "amount": 295474086,
               "book_id": "Eos accusamus."
            }
         ],
         "total": 0.5975442
      },
      "customer": {
         "address": {
            "address": "Voluptates quia non.",
            "cep": "Itaque ipsa et quae dolorum.",
            "city": "Asperiores et inventore asperiores doloribus commodi.",
            "complement": "Totam dolor.",
            "state_id": "Ad veniam natus eaque recusandae."
         },
         "cart_ids": [
            "A sed est est numquam est distinctio.",
            "Necessitatibus rerum modi earum non vel eos.",
            "Optio error et qui eaque.",
            "Sequi quos qui id consequatur tempora."
         ],
         "document": "Pariatur sequi eveniet.",
         "email": "janick.schmeler@von.biz",
         "first_name": "Aliquam vitae sit odit porro.",
         "last_name": "Quas et non non.",
         "phone": "Et aperiam expedita voluptatibus reprehenderit error facere."
      }
   }'
`, os.Args[0])
}

func catalogCreateStateUsage() {
	fmt.Fprintf(os.Stderr, `%s [flags] catalog create-state -body JSON

CreateState implements create_state.
    -body JSON: 

Example:
    `+os.Args[0]+` catalog create-state --body '{
      "country_id": "Aliquid quaerat quaerat veritatis.",
      "name": "Molestiae qui ducimus qui quae nostrum dicta."
   }'
`, os.Args[0])
}

func catalogListStatesUsage() {
	fmt.Fprintf(os.Stderr, `%s [flags] catalog list-states

ListStates implements list_states.

Example:
    `+os.Args[0]+` catalog list-states
`, os.Args[0])
}

func catalogShowStateUsage() {
	fmt.Fprintf(os.Stderr, `%s [flags] catalog show-state -id STRING

ShowState implements show_state.
    -id STRING: ID

Example:
    `+os.Args[0]+` catalog show-state --id "Ipsum vel quibusdam."
`, os.Args[0])
}
