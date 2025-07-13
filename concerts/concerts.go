package concertsapi

import (
	concerts "concerts/gen/concerts"
	"context"

	"goa.design/clue/log"
)

// concerts service example implementation.
// The example methods log the requests and return zero values.
type concertssrvc struct{}

// NewConcerts returns the concerts service implementation.
func NewConcerts() concerts.Service {
	return &concertssrvc{}
}

// List concerts with optional pagination. Returns an array of concerts sorted
// by date.
func (s *concertssrvc) List(ctx context.Context, p *concerts.ListPayload) (res []*concerts.Concert, err error) {
	log.Printf(ctx, "concerts.list")
	return
}

// Create a new concert entry. All fields are required to ensure complete
// concert information.
func (s *concertssrvc) Create(ctx context.Context, p *concerts.ConcertPayload) (res *concerts.Concert, err error) {
	res = &concerts.Concert{}
	log.Printf(ctx, "concerts.create")
	return
}

// Get a single concert by its unique ID.
func (s *concertssrvc) Show(ctx context.Context, p *concerts.ShowPayload) (res *concerts.Concert, err error) {
	res = &concerts.Concert{}
	log.Printf(ctx, "concerts.show")
	return
}

// Update an existing concert by ID. Only provided fields will be updated.
func (s *concertssrvc) Update(ctx context.Context, p *concerts.UpdatePayload) (res *concerts.Concert, err error) {
	res = &concerts.Concert{}
	log.Printf(ctx, "concerts.update")
	return
}

// Remove a concert from the system by ID. This operation cannot be undone.
func (s *concertssrvc) Delete(ctx context.Context, p *concerts.DeletePayload) (err error) {
	log.Printf(ctx, "concerts.delete")
	return
}
