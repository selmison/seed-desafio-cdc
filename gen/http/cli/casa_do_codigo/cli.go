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
	actorsc "github.com/selmison/seed-desafio-cdc/gen/http/actors/client"
	categoriesc "github.com/selmison/seed-desafio-cdc/gen/http/categories/client"
	goahttp "goa.design/goa/v3/http"
)

// UsageCommands returns the set of commands and sub-commands using the format
//
//    command (subcommand1|subcommand2|...)
//
func UsageCommands() string {
	return `actors create-actor
categories create-category
`
}

// UsageExamples produces an example of a valid invocation of the CLI tool.
func UsageExamples() string {
	return os.Args[0] + ` actors create-actor --body '{
      "description": "h7r",
      "e-mail": "Dicta rerum nesciunt perspiciatis.",
      "name": "Sint atque."
   }'` + "\n" +
		os.Args[0] + ` categories create-category --body '{
      "name": "Sapiente facilis mollitia aut."
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
		actorsFlags = flag.NewFlagSet("actors", flag.ContinueOnError)

		actorsCreateActorFlags    = flag.NewFlagSet("create-actor", flag.ExitOnError)
		actorsCreateActorBodyFlag = actorsCreateActorFlags.String("body", "REQUIRED", "")

		categoriesFlags = flag.NewFlagSet("categories", flag.ContinueOnError)

		categoriesCreateCategoryFlags    = flag.NewFlagSet("create-category", flag.ExitOnError)
		categoriesCreateCategoryBodyFlag = categoriesCreateCategoryFlags.String("body", "REQUIRED", "")
	)
	actorsFlags.Usage = actorsUsage
	actorsCreateActorFlags.Usage = actorsCreateActorUsage

	categoriesFlags.Usage = categoriesUsage
	categoriesCreateCategoryFlags.Usage = categoriesCreateCategoryUsage

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
		case "actors":
			svcf = actorsFlags
		case "categories":
			svcf = categoriesFlags
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
		case "actors":
			switch epn {
			case "create-actor":
				epf = actorsCreateActorFlags

			}

		case "categories":
			switch epn {
			case "create-category":
				epf = categoriesCreateCategoryFlags

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
		case "actors":
			c := actorsc.NewClient(scheme, host, doer, enc, dec, restore)
			switch epn {
			case "create-actor":
				endpoint = c.CreateActor()
				data, err = actorsc.BuildCreateActorPayload(*actorsCreateActorBodyFlag)
			}
		case "categories":
			c := categoriesc.NewClient(scheme, host, doer, enc, dec, restore)
			switch epn {
			case "create-category":
				endpoint = c.CreateCategory()
				data, err = categoriesc.BuildCreateCategoryPayload(*categoriesCreateCategoryBodyFlag)
			}
		}
	}
	if err != nil {
		return nil, nil, err
	}

	return endpoint, data, nil
}

// actorsUsage displays the usage of the actors command and its subcommands.
func actorsUsage() {
	fmt.Fprintf(os.Stderr, `The actors service performs operations on actors
Usage:
    %s [globalflags] actors COMMAND [flags]

COMMAND:
    create-actor: CreateActor implements create_actor.

Additional help:
    %s actors COMMAND --help
`, os.Args[0], os.Args[0])
}
func actorsCreateActorUsage() {
	fmt.Fprintf(os.Stderr, `%s [flags] actors create-actor -body JSON

CreateActor implements create_actor.
    -body JSON: 

Example:
    `+os.Args[0]+` actors create-actor --body '{
      "description": "h7r",
      "e-mail": "Dicta rerum nesciunt perspiciatis.",
      "name": "Sint atque."
   }'
`, os.Args[0])
}

// categoriesUsage displays the usage of the categories command and its
// subcommands.
func categoriesUsage() {
	fmt.Fprintf(os.Stderr, `The categories service performs operations on categories
Usage:
    %s [globalflags] categories COMMAND [flags]

COMMAND:
    create-category: CreateCategory implements create_category.

Additional help:
    %s categories COMMAND --help
`, os.Args[0], os.Args[0])
}
func categoriesCreateCategoryUsage() {
	fmt.Fprintf(os.Stderr, `%s [flags] categories create-category -body JSON

CreateCategory implements create_category.
    -body JSON: 

Example:
    `+os.Args[0]+` categories create-category --body '{
      "name": "Sapiente facilis mollitia aut."
   }'
`, os.Args[0])
}
