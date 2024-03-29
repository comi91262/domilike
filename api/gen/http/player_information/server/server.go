// Code generated by goa v3.11.0, DO NOT EDIT.
//
// player-information HTTP server
//
// Command:
// $ goa gen github.com/comi91262/domilike/design

package server

import (
	"context"
	"net/http"
	"regexp"

	playerinformation "github.com/comi91262/domilike/gen/player_information"
	goahttp "goa.design/goa/v3/http"
	goa "goa.design/goa/v3/pkg"
	"goa.design/plugins/v3/cors"
)

// Server lists the player-information service endpoint HTTP handlers.
type Server struct {
	Mounts           []*MountPoint
	Create           http.Handler
	Delete           http.Handler
	GetCoins         http.Handler
	GetVictoryPoints http.Handler
	GetDecks         http.Handler
	GetDiscards      http.Handler
	GetHands         http.Handler
	GetPlayArea      http.Handler
	CORS             http.Handler
}

// MountPoint holds information about the mounted endpoints.
type MountPoint struct {
	// Method is the name of the service method served by the mounted HTTP handler.
	Method string
	// Verb is the HTTP method used to match requests to the mounted handler.
	Verb string
	// Pattern is the HTTP request path pattern used to match requests to the
	// mounted handler.
	Pattern string
}

// New instantiates HTTP handlers for all the player-information service
// endpoints using the provided encoder and decoder. The handlers are mounted
// on the given mux using the HTTP verb and path defined in the design.
// errhandler is called whenever a response fails to be encoded. formatter is
// used to format errors returned by the service methods prior to encoding.
// Both errhandler and formatter are optional and can be nil.
func New(
	e *playerinformation.Endpoints,
	mux goahttp.Muxer,
	decoder func(*http.Request) goahttp.Decoder,
	encoder func(context.Context, http.ResponseWriter) goahttp.Encoder,
	errhandler func(context.Context, http.ResponseWriter, error),
	formatter func(ctx context.Context, err error) goahttp.Statuser,
) *Server {
	return &Server{
		Mounts: []*MountPoint{
			{"Create", "POST", "/api/player-informations"},
			{"Delete", "DELETE", "/api/player-informations"},
			{"GetCoins", "GET", "/api/coins"},
			{"GetVictoryPoints", "GET", "/api/victory-points"},
			{"GetDecks", "GET", "/api/decks"},
			{"GetDiscards", "GET", "/api/discards"},
			{"GetHands", "GET", "/api/hands"},
			{"GetPlayArea", "GET", "/api/play-areas"},
			{"CORS", "OPTIONS", "/api/player-informations"},
			{"CORS", "OPTIONS", "/api/coins"},
			{"CORS", "OPTIONS", "/api/victory-points"},
			{"CORS", "OPTIONS", "/api/decks"},
			{"CORS", "OPTIONS", "/api/discards"},
			{"CORS", "OPTIONS", "/api/hands"},
			{"CORS", "OPTIONS", "/api/play-areas"},
		},
		Create:           NewCreateHandler(e.Create, mux, decoder, encoder, errhandler, formatter),
		Delete:           NewDeleteHandler(e.Delete, mux, decoder, encoder, errhandler, formatter),
		GetCoins:         NewGetCoinsHandler(e.GetCoins, mux, decoder, encoder, errhandler, formatter),
		GetVictoryPoints: NewGetVictoryPointsHandler(e.GetVictoryPoints, mux, decoder, encoder, errhandler, formatter),
		GetDecks:         NewGetDecksHandler(e.GetDecks, mux, decoder, encoder, errhandler, formatter),
		GetDiscards:      NewGetDiscardsHandler(e.GetDiscards, mux, decoder, encoder, errhandler, formatter),
		GetHands:         NewGetHandsHandler(e.GetHands, mux, decoder, encoder, errhandler, formatter),
		GetPlayArea:      NewGetPlayAreaHandler(e.GetPlayArea, mux, decoder, encoder, errhandler, formatter),
		CORS:             NewCORSHandler(),
	}
}

// Service returns the name of the service served.
func (s *Server) Service() string { return "player-information" }

// Use wraps the server handlers with the given middleware.
func (s *Server) Use(m func(http.Handler) http.Handler) {
	s.Create = m(s.Create)
	s.Delete = m(s.Delete)
	s.GetCoins = m(s.GetCoins)
	s.GetVictoryPoints = m(s.GetVictoryPoints)
	s.GetDecks = m(s.GetDecks)
	s.GetDiscards = m(s.GetDiscards)
	s.GetHands = m(s.GetHands)
	s.GetPlayArea = m(s.GetPlayArea)
	s.CORS = m(s.CORS)
}

// MethodNames returns the methods served.
func (s *Server) MethodNames() []string { return playerinformation.MethodNames[:] }

// Mount configures the mux to serve the player-information endpoints.
func Mount(mux goahttp.Muxer, h *Server) {
	MountCreateHandler(mux, h.Create)
	MountDeleteHandler(mux, h.Delete)
	MountGetCoinsHandler(mux, h.GetCoins)
	MountGetVictoryPointsHandler(mux, h.GetVictoryPoints)
	MountGetDecksHandler(mux, h.GetDecks)
	MountGetDiscardsHandler(mux, h.GetDiscards)
	MountGetHandsHandler(mux, h.GetHands)
	MountGetPlayAreaHandler(mux, h.GetPlayArea)
	MountCORSHandler(mux, h.CORS)
}

// Mount configures the mux to serve the player-information endpoints.
func (s *Server) Mount(mux goahttp.Muxer) {
	Mount(mux, s)
}

// MountCreateHandler configures the mux to serve the "player-information"
// service "create" endpoint.
func MountCreateHandler(mux goahttp.Muxer, h http.Handler) {
	f, ok := HandlePlayerInformationOrigin(h).(http.HandlerFunc)
	if !ok {
		f = func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		}
	}
	mux.Handle("POST", "/api/player-informations", f)
}

// NewCreateHandler creates a HTTP handler which loads the HTTP request and
// calls the "player-information" service "create" endpoint.
func NewCreateHandler(
	endpoint goa.Endpoint,
	mux goahttp.Muxer,
	decoder func(*http.Request) goahttp.Decoder,
	encoder func(context.Context, http.ResponseWriter) goahttp.Encoder,
	errhandler func(context.Context, http.ResponseWriter, error),
	formatter func(ctx context.Context, err error) goahttp.Statuser,
) http.Handler {
	var (
		decodeRequest  = DecodeCreateRequest(mux, decoder)
		encodeResponse = EncodeCreateResponse(encoder)
		encodeError    = goahttp.ErrorEncoder(encoder, formatter)
	)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), goahttp.AcceptTypeKey, r.Header.Get("Accept"))
		ctx = context.WithValue(ctx, goa.MethodKey, "create")
		ctx = context.WithValue(ctx, goa.ServiceKey, "player-information")
		payload, err := decodeRequest(r)
		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				errhandler(ctx, w, err)
			}
			return
		}
		res, err := endpoint(ctx, payload)
		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				errhandler(ctx, w, err)
			}
			return
		}
		if err := encodeResponse(ctx, w, res); err != nil {
			errhandler(ctx, w, err)
		}
	})
}

// MountDeleteHandler configures the mux to serve the "player-information"
// service "delete" endpoint.
func MountDeleteHandler(mux goahttp.Muxer, h http.Handler) {
	f, ok := HandlePlayerInformationOrigin(h).(http.HandlerFunc)
	if !ok {
		f = func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		}
	}
	mux.Handle("DELETE", "/api/player-informations", f)
}

// NewDeleteHandler creates a HTTP handler which loads the HTTP request and
// calls the "player-information" service "delete" endpoint.
func NewDeleteHandler(
	endpoint goa.Endpoint,
	mux goahttp.Muxer,
	decoder func(*http.Request) goahttp.Decoder,
	encoder func(context.Context, http.ResponseWriter) goahttp.Encoder,
	errhandler func(context.Context, http.ResponseWriter, error),
	formatter func(ctx context.Context, err error) goahttp.Statuser,
) http.Handler {
	var (
		decodeRequest  = DecodeDeleteRequest(mux, decoder)
		encodeResponse = EncodeDeleteResponse(encoder)
		encodeError    = goahttp.ErrorEncoder(encoder, formatter)
	)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), goahttp.AcceptTypeKey, r.Header.Get("Accept"))
		ctx = context.WithValue(ctx, goa.MethodKey, "delete")
		ctx = context.WithValue(ctx, goa.ServiceKey, "player-information")
		payload, err := decodeRequest(r)
		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				errhandler(ctx, w, err)
			}
			return
		}
		res, err := endpoint(ctx, payload)
		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				errhandler(ctx, w, err)
			}
			return
		}
		if err := encodeResponse(ctx, w, res); err != nil {
			errhandler(ctx, w, err)
		}
	})
}

// MountGetCoinsHandler configures the mux to serve the "player-information"
// service "get_coins" endpoint.
func MountGetCoinsHandler(mux goahttp.Muxer, h http.Handler) {
	f, ok := HandlePlayerInformationOrigin(h).(http.HandlerFunc)
	if !ok {
		f = func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		}
	}
	mux.Handle("GET", "/api/coins", f)
}

// NewGetCoinsHandler creates a HTTP handler which loads the HTTP request and
// calls the "player-information" service "get_coins" endpoint.
func NewGetCoinsHandler(
	endpoint goa.Endpoint,
	mux goahttp.Muxer,
	decoder func(*http.Request) goahttp.Decoder,
	encoder func(context.Context, http.ResponseWriter) goahttp.Encoder,
	errhandler func(context.Context, http.ResponseWriter, error),
	formatter func(ctx context.Context, err error) goahttp.Statuser,
) http.Handler {
	var (
		decodeRequest  = DecodeGetCoinsRequest(mux, decoder)
		encodeResponse = EncodeGetCoinsResponse(encoder)
		encodeError    = goahttp.ErrorEncoder(encoder, formatter)
	)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), goahttp.AcceptTypeKey, r.Header.Get("Accept"))
		ctx = context.WithValue(ctx, goa.MethodKey, "get_coins")
		ctx = context.WithValue(ctx, goa.ServiceKey, "player-information")
		payload, err := decodeRequest(r)
		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				errhandler(ctx, w, err)
			}
			return
		}
		res, err := endpoint(ctx, payload)
		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				errhandler(ctx, w, err)
			}
			return
		}
		if err := encodeResponse(ctx, w, res); err != nil {
			errhandler(ctx, w, err)
		}
	})
}

// MountGetVictoryPointsHandler configures the mux to serve the
// "player-information" service "get_victory_points" endpoint.
func MountGetVictoryPointsHandler(mux goahttp.Muxer, h http.Handler) {
	f, ok := HandlePlayerInformationOrigin(h).(http.HandlerFunc)
	if !ok {
		f = func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		}
	}
	mux.Handle("GET", "/api/victory-points", f)
}

// NewGetVictoryPointsHandler creates a HTTP handler which loads the HTTP
// request and calls the "player-information" service "get_victory_points"
// endpoint.
func NewGetVictoryPointsHandler(
	endpoint goa.Endpoint,
	mux goahttp.Muxer,
	decoder func(*http.Request) goahttp.Decoder,
	encoder func(context.Context, http.ResponseWriter) goahttp.Encoder,
	errhandler func(context.Context, http.ResponseWriter, error),
	formatter func(ctx context.Context, err error) goahttp.Statuser,
) http.Handler {
	var (
		decodeRequest  = DecodeGetVictoryPointsRequest(mux, decoder)
		encodeResponse = EncodeGetVictoryPointsResponse(encoder)
		encodeError    = goahttp.ErrorEncoder(encoder, formatter)
	)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), goahttp.AcceptTypeKey, r.Header.Get("Accept"))
		ctx = context.WithValue(ctx, goa.MethodKey, "get_victory_points")
		ctx = context.WithValue(ctx, goa.ServiceKey, "player-information")
		payload, err := decodeRequest(r)
		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				errhandler(ctx, w, err)
			}
			return
		}
		res, err := endpoint(ctx, payload)
		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				errhandler(ctx, w, err)
			}
			return
		}
		if err := encodeResponse(ctx, w, res); err != nil {
			errhandler(ctx, w, err)
		}
	})
}

// MountGetDecksHandler configures the mux to serve the "player-information"
// service "get_decks" endpoint.
func MountGetDecksHandler(mux goahttp.Muxer, h http.Handler) {
	f, ok := HandlePlayerInformationOrigin(h).(http.HandlerFunc)
	if !ok {
		f = func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		}
	}
	mux.Handle("GET", "/api/decks", f)
}

// NewGetDecksHandler creates a HTTP handler which loads the HTTP request and
// calls the "player-information" service "get_decks" endpoint.
func NewGetDecksHandler(
	endpoint goa.Endpoint,
	mux goahttp.Muxer,
	decoder func(*http.Request) goahttp.Decoder,
	encoder func(context.Context, http.ResponseWriter) goahttp.Encoder,
	errhandler func(context.Context, http.ResponseWriter, error),
	formatter func(ctx context.Context, err error) goahttp.Statuser,
) http.Handler {
	var (
		decodeRequest  = DecodeGetDecksRequest(mux, decoder)
		encodeResponse = EncodeGetDecksResponse(encoder)
		encodeError    = goahttp.ErrorEncoder(encoder, formatter)
	)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), goahttp.AcceptTypeKey, r.Header.Get("Accept"))
		ctx = context.WithValue(ctx, goa.MethodKey, "get_decks")
		ctx = context.WithValue(ctx, goa.ServiceKey, "player-information")
		payload, err := decodeRequest(r)
		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				errhandler(ctx, w, err)
			}
			return
		}
		res, err := endpoint(ctx, payload)
		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				errhandler(ctx, w, err)
			}
			return
		}
		if err := encodeResponse(ctx, w, res); err != nil {
			errhandler(ctx, w, err)
		}
	})
}

// MountGetDiscardsHandler configures the mux to serve the "player-information"
// service "get_discards" endpoint.
func MountGetDiscardsHandler(mux goahttp.Muxer, h http.Handler) {
	f, ok := HandlePlayerInformationOrigin(h).(http.HandlerFunc)
	if !ok {
		f = func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		}
	}
	mux.Handle("GET", "/api/discards", f)
}

// NewGetDiscardsHandler creates a HTTP handler which loads the HTTP request
// and calls the "player-information" service "get_discards" endpoint.
func NewGetDiscardsHandler(
	endpoint goa.Endpoint,
	mux goahttp.Muxer,
	decoder func(*http.Request) goahttp.Decoder,
	encoder func(context.Context, http.ResponseWriter) goahttp.Encoder,
	errhandler func(context.Context, http.ResponseWriter, error),
	formatter func(ctx context.Context, err error) goahttp.Statuser,
) http.Handler {
	var (
		decodeRequest  = DecodeGetDiscardsRequest(mux, decoder)
		encodeResponse = EncodeGetDiscardsResponse(encoder)
		encodeError    = goahttp.ErrorEncoder(encoder, formatter)
	)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), goahttp.AcceptTypeKey, r.Header.Get("Accept"))
		ctx = context.WithValue(ctx, goa.MethodKey, "get_discards")
		ctx = context.WithValue(ctx, goa.ServiceKey, "player-information")
		payload, err := decodeRequest(r)
		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				errhandler(ctx, w, err)
			}
			return
		}
		res, err := endpoint(ctx, payload)
		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				errhandler(ctx, w, err)
			}
			return
		}
		if err := encodeResponse(ctx, w, res); err != nil {
			errhandler(ctx, w, err)
		}
	})
}

// MountGetHandsHandler configures the mux to serve the "player-information"
// service "get_hands" endpoint.
func MountGetHandsHandler(mux goahttp.Muxer, h http.Handler) {
	f, ok := HandlePlayerInformationOrigin(h).(http.HandlerFunc)
	if !ok {
		f = func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		}
	}
	mux.Handle("GET", "/api/hands", f)
}

// NewGetHandsHandler creates a HTTP handler which loads the HTTP request and
// calls the "player-information" service "get_hands" endpoint.
func NewGetHandsHandler(
	endpoint goa.Endpoint,
	mux goahttp.Muxer,
	decoder func(*http.Request) goahttp.Decoder,
	encoder func(context.Context, http.ResponseWriter) goahttp.Encoder,
	errhandler func(context.Context, http.ResponseWriter, error),
	formatter func(ctx context.Context, err error) goahttp.Statuser,
) http.Handler {
	var (
		decodeRequest  = DecodeGetHandsRequest(mux, decoder)
		encodeResponse = EncodeGetHandsResponse(encoder)
		encodeError    = goahttp.ErrorEncoder(encoder, formatter)
	)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), goahttp.AcceptTypeKey, r.Header.Get("Accept"))
		ctx = context.WithValue(ctx, goa.MethodKey, "get_hands")
		ctx = context.WithValue(ctx, goa.ServiceKey, "player-information")
		payload, err := decodeRequest(r)
		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				errhandler(ctx, w, err)
			}
			return
		}
		res, err := endpoint(ctx, payload)
		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				errhandler(ctx, w, err)
			}
			return
		}
		if err := encodeResponse(ctx, w, res); err != nil {
			errhandler(ctx, w, err)
		}
	})
}

// MountGetPlayAreaHandler configures the mux to serve the "player-information"
// service "get_play_area" endpoint.
func MountGetPlayAreaHandler(mux goahttp.Muxer, h http.Handler) {
	f, ok := HandlePlayerInformationOrigin(h).(http.HandlerFunc)
	if !ok {
		f = func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		}
	}
	mux.Handle("GET", "/api/play-areas", f)
}

// NewGetPlayAreaHandler creates a HTTP handler which loads the HTTP request
// and calls the "player-information" service "get_play_area" endpoint.
func NewGetPlayAreaHandler(
	endpoint goa.Endpoint,
	mux goahttp.Muxer,
	decoder func(*http.Request) goahttp.Decoder,
	encoder func(context.Context, http.ResponseWriter) goahttp.Encoder,
	errhandler func(context.Context, http.ResponseWriter, error),
	formatter func(ctx context.Context, err error) goahttp.Statuser,
) http.Handler {
	var (
		decodeRequest  = DecodeGetPlayAreaRequest(mux, decoder)
		encodeResponse = EncodeGetPlayAreaResponse(encoder)
		encodeError    = goahttp.ErrorEncoder(encoder, formatter)
	)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), goahttp.AcceptTypeKey, r.Header.Get("Accept"))
		ctx = context.WithValue(ctx, goa.MethodKey, "get_play_area")
		ctx = context.WithValue(ctx, goa.ServiceKey, "player-information")
		payload, err := decodeRequest(r)
		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				errhandler(ctx, w, err)
			}
			return
		}
		res, err := endpoint(ctx, payload)
		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				errhandler(ctx, w, err)
			}
			return
		}
		if err := encodeResponse(ctx, w, res); err != nil {
			errhandler(ctx, w, err)
		}
	})
}

// MountCORSHandler configures the mux to serve the CORS endpoints for the
// service player-information.
func MountCORSHandler(mux goahttp.Muxer, h http.Handler) {
	h = HandlePlayerInformationOrigin(h)
	mux.Handle("OPTIONS", "/api/player-informations", h.ServeHTTP)
	mux.Handle("OPTIONS", "/api/coins", h.ServeHTTP)
	mux.Handle("OPTIONS", "/api/victory-points", h.ServeHTTP)
	mux.Handle("OPTIONS", "/api/decks", h.ServeHTTP)
	mux.Handle("OPTIONS", "/api/discards", h.ServeHTTP)
	mux.Handle("OPTIONS", "/api/hands", h.ServeHTTP)
	mux.Handle("OPTIONS", "/api/play-areas", h.ServeHTTP)
}

// NewCORSHandler creates a HTTP handler which returns a simple 200 response.
func NewCORSHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
	})
}

// HandlePlayerInformationOrigin applies the CORS response headers
// corresponding to the origin for the service player-information.
func HandlePlayerInformationOrigin(h http.Handler) http.Handler {
	spec0 := regexp.MustCompile(".*localhost.*")
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		origin := r.Header.Get("Origin")
		if origin == "" {
			// Not a CORS request
			h.ServeHTTP(w, r)
			return
		}
		if cors.MatchOriginRegexp(origin, spec0) {
			w.Header().Set("Access-Control-Allow-Origin", origin)
			w.Header().Set("Vary", "Origin")
			w.Header().Set("Access-Control-Expose-Headers", "X-Time, X-Api-Version")
			w.Header().Set("Access-Control-Max-Age", "100")
			w.Header().Set("Access-Control-Allow-Credentials", "true")
			if acrm := r.Header.Get("Access-Control-Request-Method"); acrm != "" {
				// We are handling a preflight request
				w.Header().Set("Access-Control-Allow-Methods", "GET, POST")
				w.Header().Set("Access-Control-Allow-Headers", "X-Requested-With")
			}
			h.ServeHTTP(w, r)
			return
		}
		h.ServeHTTP(w, r)
		return
	})
}
