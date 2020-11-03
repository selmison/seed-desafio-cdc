// Code generated by goa v3.2.4, DO NOT EDIT.
//
// books HTTP client CLI support package
//
// Command:
// $ goa gen github.com/selmison/seed-desafio-cdc/design

package client

import (
	"encoding/json"
	"fmt"
	"unicode/utf8"

	books "github.com/selmison/seed-desafio-cdc/gen/books"
	goa "goa.design/goa/v3/pkg"
)

// BuildCreateBookPayload builds the payload for the books create_book endpoint
// from CLI flags.
func BuildCreateBookPayload(booksCreateBookBody string) (*books.CreateBookDTO, error) {
	var err error
	var body CreateBookRequestBody
	{
		err = json.Unmarshal([]byte(booksCreateBookBody), &body)
		if err != nil {
			return nil, fmt.Errorf("invalid JSON for body, \nerror: %s, \nexample of valid JSON:\n%s", err, "'{\n      \"actor_id\": \"Exercitationem atque velit.\",\n      \"category_id\": \"Soluta eos ipsa ad.\",\n      \"isbn\": \"Earum qui est nam quos aperiam quidem.\",\n      \"issue\": \"Quisquam in cum numquam.\",\n      \"pages\": 1974383956477392325,\n      \"price\": 20.982689,\n      \"summary\": \"Recusandae sit minus.\",\n      \"synopsis\": \"1km\",\n      \"title\": \"Debitis eos et culpa.\"\n   }'")
		}
		if utf8.RuneCountInString(body.Synopsis) > 500 {
			err = goa.MergeErrors(err, goa.InvalidLengthError("body.synopsis", body.Synopsis, utf8.RuneCountInString(body.Synopsis), 500, false))
		}
		if body.Price < 20 {
			err = goa.MergeErrors(err, goa.InvalidRangeError("body.price", body.Price, 20, true))
		}
		if body.Pages < 100 {
			err = goa.MergeErrors(err, goa.InvalidRangeError("body.pages", body.Pages, 100, true))
		}
		if err != nil {
			return nil, err
		}
	}
	v := &books.CreateBookDTO{
		Title:      body.Title,
		Synopsis:   body.Synopsis,
		Summary:    body.Summary,
		Price:      body.Price,
		Pages:      body.Pages,
		Isbn:       body.Isbn,
		Issue:      body.Issue,
		CategoryID: body.CategoryID,
		ActorID:    body.ActorID,
	}

	return v, nil
}
