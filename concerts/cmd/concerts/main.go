package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/google/uuid"
	goahttp "goa.design/goa/v3/http"

	// Use gen prefix for generated packages
	genconcerts "concerts/gen/concerts"
	genconcertshttpserver "concerts/gen/http/concerts/server"

    "github.com/vmihailenco/msgpack/v5"
    "strings"
)

// ConcertsService implements the genconcerts.Service interface
type ConcertsService struct {
	concerts []*genconcerts.Concert // Using a simple slice for in-memory storage
}

// List upcoming concerts with optional pagination.
func (m *ConcertsService) List(ctx context.Context, p *genconcerts.ListPayload) ([]*genconcerts.Concert, error) {
	// implement pagination logic
	start := (p.Page - 1) * p.Limit
	end := start + p.Limit
	if end > len(m.concerts) {
		end = len(m.concerts)
	}
	return m.concerts[start:end], nil
}

// Create a new concerts entry.
func (m *ConcertsService) Create(ctx context.Context, p *genconcerts.ConcertPayload) (*genconcerts.Concert, error) {
	newConcert := &genconcerts.Concert{
		ID:     uuid.New().String(), // Using Google’s UUID library for unique identifiers
		Artist: p.Artist,
		Date:   p.Date,
		Venue:  p.Venue,
		Price:  p.Price,
	}
	m.concerts = append(m.concerts, newConcert)
	return newConcert, nil
}

// Get a single concert by ID.
func (m *ConcertsService) Show(ctx context.Context, p *genconcerts.ShowPayload) (*genconcerts.Concert, error) {
	for _, concert := range m.concerts {
		if concert.ID == p.ConcertID {
			return concert, nil
		}
	}
	// Use designed error
	return nil, genconcerts.MakeNotFound(fmt.Errorf("concert not found: %s", p.ConcertID))
}

// Update an existing concert by ID.
func (m *ConcertsService) Update(ctx context.Context, p *genconcerts.UpdatePayload) (*genconcerts.Concert, error) {
	for i, concert := range m.concerts {
		if concert.ID == p.ConcertID {
			if p.Artist != nil {
				concert.Artist = *p.Artist
			}
			if p.Date != nil {
				concert.Date = *p.Date
			}
			if p.Venue != nil {
				concert.Venue = *p.Venue
			}
			if p.Price != nil {
				concert.Price = *p.Price
			}
			m.concerts[i] = concert
			return concert, nil
		}
	}
	return nil, genconcerts.MakeNotFound(fmt.Errorf("concert not found: %s", p.ConcertID))
}

// Remove a concert from the system by ID.
func (m *ConcertsService) Delete(ctx context.Context, p *genconcerts.DeletePayload) error {
	for i, concert := range m.concerts {
		if concert.ID == p.ConcertID {
			m.concerts = append(m.concerts[:i], m.concerts[i+1:]...)
			return nil
		}
	}
	return genconcerts.MakeNotFound(fmt.Errorf("concert not found: %s", p.ConcertID))
}

type (
    // MessagePack encoder implementation
    msgpackEnc struct {
        w http.ResponseWriter
    }

    // MessagePack decoder implementation
    msgpackDec struct {
        r *http.Request
    }
)

// Custom encoder constructor - this creates our MessagePack encoder
func msgpackEncoder(ctx context.Context, w http.ResponseWriter) goahttp.Encoder {
    return &msgpackEnc{w: w}
}

func (e *msgpackEnc) Encode(v any) error {
    e.w.Header().Set("Content-Type", "application/msgpack")
    return msgpack.NewEncoder(e.w).Encode(v)
}

// Custom decoder constructor - this handles incoming MessagePack data
func msgpackDecoder(r *http.Request) goahttp.Decoder {
    return &msgpackDec{r: r}
}

func (d *msgpackDec) Decode(v any) error {
    return msgpack.NewDecoder(d.r.Body).Decode(v)
}


// main instantiates the service and starts the HTTP server.
func main() {
	// Instantiate the service
	svc := &ConcertsService{}

	// Wrap it in the generated endpoints
	endpoints := genconcerts.NewEndpoints(svc)

	// Build an HTTP handler
	mux := goahttp.NewMuxer() // Uses Goa’s built-in HTTP multiplexer
	// handler := genconcertshttpserver.New(endpoints, mux, goahttp.RequestDecoder, goahttp.ResponseEncoder, nil, nil)


    // Smart encoder selection based on what the client wants (Accept header)
    encodeFunc := func(ctx context.Context, w http.ResponseWriter) goahttp.Encoder {
        accept := ctx.Value(goahttp.AcceptTypeKey).(string)
        
        // Parse Accept header which may contain multiple types with q-values
        // For example: "application/json;q=0.9,application/msgpack"
        types := strings.Split(accept, ",")
        for _, t := range types {
            mt := strings.TrimSpace(strings.Split(t, ";")[0])
            switch mt {
            case "application/msgpack":
                return msgpackEncoder(ctx, w)
            case "application/json", "*/*":
                return goahttp.ResponseEncoder(ctx, w)
            }
        }
        
        // When in doubt, JSON is our friend!
        return goahttp.ResponseEncoder(ctx, w)
    }

    // Smart decoder selection based on what the client is sending (Content-Type)
    decodeFunc := func(r *http.Request) goahttp.Decoder {
        if r.Header.Get("Content-Type") == "application/msgpack" {
            return msgpackDecoder(r)
        }
        return goahttp.RequestDecoder(r)
    }

    // Wire up our custom encoder/decoder
    handler := genconcertshttpserver.New(
        endpoints,
        mux,
        decodeFunc,
        encodeFunc,
        nil,
        nil,
    )

	// Logs all mounted routes for debugging
	genconcertshttpserver.Mount(mux, handler)

	// Create a new HTTP server
	port := "8080"
	server := &http.Server{Addr: ":" + port, Handler: mux}

	// Log the supported routes
	for _, mount := range handler.Mounts {
		log.Printf("%q mounted on %s %s", mount.Method, mount.Verb, mount.Pattern)
	}

	// Start the server (this will block the execution)
	log.Printf("Starting concerts service on :%s", port)
	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}

}
