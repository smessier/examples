// Code generated by goa v3.0.0, DO NOT EDIT.
//
// chatter HTTP server
//
// Command:
// $ goa gen goa.design/examples/streaming/design -o
// $(GOPATH)/src/goa.design/examples/streaming

package server

import (
	"context"
	"io"
	"net/http"
	"sync"
	"time"

	"github.com/gorilla/websocket"
	chatter "goa.design/examples/streaming/gen/chatter"
	goa "goa.design/goa/v3"
	goahttp "goa.design/goa/v3/http"
)

// Server lists the chatter service endpoint HTTP handlers.
type Server struct {
	Mounts    []*MountPoint
	Login     http.Handler
	Echoer    http.Handler
	Listener  http.Handler
	Summary   http.Handler
	Subscribe http.Handler
	History   http.Handler
}

// ErrorNamer is an interface implemented by generated error structs that
// exposes the name of the error as defined in the design.
type ErrorNamer interface {
	ErrorName() string
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

// ConnConfigurer holds the websocket connection configurer functions for the
// streaming endpoints in "chatter" service.
type ConnConfigurer struct {
	EchoerFn    goahttp.ConnConfigureFunc
	ListenerFn  goahttp.ConnConfigureFunc
	SummaryFn   goahttp.ConnConfigureFunc
	SubscribeFn goahttp.ConnConfigureFunc
	HistoryFn   goahttp.ConnConfigureFunc
}

// echoerServerStream implements the chatter.EchoerServerStream interface.
type echoerServerStream struct {
	once sync.Once
	// upgrader is the websocket connection upgrader.
	upgrader goahttp.Upgrader
	// connConfigFn is the websocket connection configurer.
	connConfigFn goahttp.ConnConfigureFunc
	// cancel is the context cancellation function which cancels the request
	// context when invoked.
	cancel context.CancelFunc
	// w is the HTTP response writer used in upgrading the connection.
	w http.ResponseWriter
	// r is the HTTP request.
	r *http.Request
	// conn is the underlying websocket connection.
	conn *websocket.Conn
}

// listenerServerStream implements the chatter.ListenerServerStream interface.
type listenerServerStream struct {
	once sync.Once
	// upgrader is the websocket connection upgrader.
	upgrader goahttp.Upgrader
	// connConfigFn is the websocket connection configurer.
	connConfigFn goahttp.ConnConfigureFunc
	// cancel is the context cancellation function which cancels the request
	// context when invoked.
	cancel context.CancelFunc
	// w is the HTTP response writer used in upgrading the connection.
	w http.ResponseWriter
	// r is the HTTP request.
	r *http.Request
	// conn is the underlying websocket connection.
	conn *websocket.Conn
}

// summaryServerStream implements the chatter.SummaryServerStream interface.
type summaryServerStream struct {
	once sync.Once
	// upgrader is the websocket connection upgrader.
	upgrader goahttp.Upgrader
	// connConfigFn is the websocket connection configurer.
	connConfigFn goahttp.ConnConfigureFunc
	// cancel is the context cancellation function which cancels the request
	// context when invoked.
	cancel context.CancelFunc
	// w is the HTTP response writer used in upgrading the connection.
	w http.ResponseWriter
	// r is the HTTP request.
	r *http.Request
	// conn is the underlying websocket connection.
	conn *websocket.Conn
}

// subscribeServerStream implements the chatter.SubscribeServerStream interface.
type subscribeServerStream struct {
	once sync.Once
	// upgrader is the websocket connection upgrader.
	upgrader goahttp.Upgrader
	// connConfigFn is the websocket connection configurer.
	connConfigFn goahttp.ConnConfigureFunc
	// cancel is the context cancellation function which cancels the request
	// context when invoked.
	cancel context.CancelFunc
	// w is the HTTP response writer used in upgrading the connection.
	w http.ResponseWriter
	// r is the HTTP request.
	r *http.Request
	// conn is the underlying websocket connection.
	conn *websocket.Conn
}

// historyServerStream implements the chatter.HistoryServerStream interface.
type historyServerStream struct {
	once sync.Once
	// upgrader is the websocket connection upgrader.
	upgrader goahttp.Upgrader
	// connConfigFn is the websocket connection configurer.
	connConfigFn goahttp.ConnConfigureFunc
	// cancel is the context cancellation function which cancels the request
	// context when invoked.
	cancel context.CancelFunc
	// w is the HTTP response writer used in upgrading the connection.
	w http.ResponseWriter
	// r is the HTTP request.
	r *http.Request
	// conn is the underlying websocket connection.
	conn *websocket.Conn
	// view is the view to render chatter.ChatSummary result type before sending to
	// the websocket connection.
	view string
}

// New instantiates HTTP handlers for all the chatter service endpoints.
func New(
	e *chatter.Endpoints,
	mux goahttp.Muxer,
	dec func(*http.Request) goahttp.Decoder,
	enc func(context.Context, http.ResponseWriter) goahttp.Encoder,
	eh func(context.Context, http.ResponseWriter, error),
	up goahttp.Upgrader,
	cfn *ConnConfigurer,
) *Server {
	if cfn == nil {
		cfn = &ConnConfigurer{}
	}
	return &Server{
		Mounts: []*MountPoint{
			{"Login", "POST", "/login"},
			{"Echoer", "GET", "/echoer"},
			{"Listener", "GET", "/listener"},
			{"Summary", "GET", "/summary"},
			{"Subscribe", "GET", "/subscribe"},
			{"History", "GET", "/history"},
		},
		Login:     NewLoginHandler(e.Login, mux, dec, enc, eh),
		Echoer:    NewEchoerHandler(e.Echoer, mux, dec, enc, eh, up, cfn.EchoerFn),
		Listener:  NewListenerHandler(e.Listener, mux, dec, enc, eh, up, cfn.ListenerFn),
		Summary:   NewSummaryHandler(e.Summary, mux, dec, enc, eh, up, cfn.SummaryFn),
		Subscribe: NewSubscribeHandler(e.Subscribe, mux, dec, enc, eh, up, cfn.SubscribeFn),
		History:   NewHistoryHandler(e.History, mux, dec, enc, eh, up, cfn.HistoryFn),
	}
}

// NewConnConfigurer initializes the websocket connection configurer function
// with fn for all the streaming endpoints in "chatter" service.
func NewConnConfigurer(fn goahttp.ConnConfigureFunc) *ConnConfigurer {
	return &ConnConfigurer{
		EchoerFn:    fn,
		ListenerFn:  fn,
		SummaryFn:   fn,
		SubscribeFn: fn,
		HistoryFn:   fn,
	}
}

// Service returns the name of the service served.
func (s *Server) Service() string { return "chatter" }

// Use wraps the server handlers with the given middleware.
func (s *Server) Use(m func(http.Handler) http.Handler) {
	s.Login = m(s.Login)
	s.Echoer = m(s.Echoer)
	s.Listener = m(s.Listener)
	s.Summary = m(s.Summary)
	s.Subscribe = m(s.Subscribe)
	s.History = m(s.History)
}

// Mount configures the mux to serve the chatter endpoints.
func Mount(mux goahttp.Muxer, h *Server) {
	MountLoginHandler(mux, h.Login)
	MountEchoerHandler(mux, h.Echoer)
	MountListenerHandler(mux, h.Listener)
	MountSummaryHandler(mux, h.Summary)
	MountSubscribeHandler(mux, h.Subscribe)
	MountHistoryHandler(mux, h.History)
}

// MountLoginHandler configures the mux to serve the "chatter" service "login"
// endpoint.
func MountLoginHandler(mux goahttp.Muxer, h http.Handler) {
	f, ok := h.(http.HandlerFunc)
	if !ok {
		f = func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		}
	}
	mux.Handle("POST", "/login", f)
}

// NewLoginHandler creates a HTTP handler which loads the HTTP request and
// calls the "chatter" service "login" endpoint.
func NewLoginHandler(
	endpoint goa.Endpoint,
	mux goahttp.Muxer,
	dec func(*http.Request) goahttp.Decoder,
	enc func(context.Context, http.ResponseWriter) goahttp.Encoder,
	eh func(context.Context, http.ResponseWriter, error),
) http.Handler {
	var (
		decodeRequest  = DecodeLoginRequest(mux, dec)
		encodeResponse = EncodeLoginResponse(enc)
		encodeError    = EncodeLoginError(enc)
	)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), goahttp.AcceptTypeKey, r.Header.Get("Accept"))
		ctx = context.WithValue(ctx, goa.MethodKey, "login")
		ctx = context.WithValue(ctx, goa.ServiceKey, "chatter")
		payload, err := decodeRequest(r)
		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				eh(ctx, w, err)
			}
			return
		}

		res, err := endpoint(ctx, payload)

		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				eh(ctx, w, err)
			}
			return
		}
		if err := encodeResponse(ctx, w, res); err != nil {
			eh(ctx, w, err)
		}
	})
}

// MountEchoerHandler configures the mux to serve the "chatter" service
// "echoer" endpoint.
func MountEchoerHandler(mux goahttp.Muxer, h http.Handler) {
	f, ok := h.(http.HandlerFunc)
	if !ok {
		f = func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		}
	}
	mux.Handle("GET", "/echoer", f)
}

// NewEchoerHandler creates a HTTP handler which loads the HTTP request and
// calls the "chatter" service "echoer" endpoint.
func NewEchoerHandler(
	endpoint goa.Endpoint,
	mux goahttp.Muxer,
	dec func(*http.Request) goahttp.Decoder,
	enc func(context.Context, http.ResponseWriter) goahttp.Encoder,
	eh func(context.Context, http.ResponseWriter, error),
	up goahttp.Upgrader,
	connConfigFn goahttp.ConnConfigureFunc,
) http.Handler {
	var (
		decodeRequest = DecodeEchoerRequest(mux, dec)
		encodeError   = EncodeEchoerError(enc)
	)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), goahttp.AcceptTypeKey, r.Header.Get("Accept"))
		ctx = context.WithValue(ctx, goa.MethodKey, "echoer")
		ctx = context.WithValue(ctx, goa.ServiceKey, "chatter")
		payload, err := decodeRequest(r)
		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				eh(ctx, w, err)
			}
			return
		}

		var cancel context.CancelFunc
		{
			ctx, cancel = context.WithCancel(ctx)
		}
		v := &chatter.EchoerEndpointInput{
			Stream: &echoerServerStream{
				upgrader:     up,
				connConfigFn: connConfigFn,
				cancel:       cancel,
				w:            w,
				r:            r,
			},
			Payload: payload.(*chatter.EchoerPayload),
		}
		_, err = endpoint(ctx, v)

		if err != nil {
			if _, ok := err.(websocket.HandshakeError); ok {
				return
			}
			if err := encodeError(ctx, w, err); err != nil {
				eh(ctx, w, err)
			}
			return
		}
	})
}

// MountListenerHandler configures the mux to serve the "chatter" service
// "listener" endpoint.
func MountListenerHandler(mux goahttp.Muxer, h http.Handler) {
	f, ok := h.(http.HandlerFunc)
	if !ok {
		f = func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		}
	}
	mux.Handle("GET", "/listener", f)
}

// NewListenerHandler creates a HTTP handler which loads the HTTP request and
// calls the "chatter" service "listener" endpoint.
func NewListenerHandler(
	endpoint goa.Endpoint,
	mux goahttp.Muxer,
	dec func(*http.Request) goahttp.Decoder,
	enc func(context.Context, http.ResponseWriter) goahttp.Encoder,
	eh func(context.Context, http.ResponseWriter, error),
	up goahttp.Upgrader,
	connConfigFn goahttp.ConnConfigureFunc,
) http.Handler {
	var (
		decodeRequest = DecodeListenerRequest(mux, dec)
		encodeError   = EncodeListenerError(enc)
	)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), goahttp.AcceptTypeKey, r.Header.Get("Accept"))
		ctx = context.WithValue(ctx, goa.MethodKey, "listener")
		ctx = context.WithValue(ctx, goa.ServiceKey, "chatter")
		payload, err := decodeRequest(r)
		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				eh(ctx, w, err)
			}
			return
		}

		var cancel context.CancelFunc
		{
			ctx, cancel = context.WithCancel(ctx)
		}
		v := &chatter.ListenerEndpointInput{
			Stream: &listenerServerStream{
				upgrader:     up,
				connConfigFn: connConfigFn,
				cancel:       cancel,
				w:            w,
				r:            r,
			},
			Payload: payload.(*chatter.ListenerPayload),
		}
		_, err = endpoint(ctx, v)

		if err != nil {
			if _, ok := err.(websocket.HandshakeError); ok {
				return
			}
			if err := encodeError(ctx, w, err); err != nil {
				eh(ctx, w, err)
			}
			return
		}
	})
}

// MountSummaryHandler configures the mux to serve the "chatter" service
// "summary" endpoint.
func MountSummaryHandler(mux goahttp.Muxer, h http.Handler) {
	f, ok := h.(http.HandlerFunc)
	if !ok {
		f = func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		}
	}
	mux.Handle("GET", "/summary", f)
}

// NewSummaryHandler creates a HTTP handler which loads the HTTP request and
// calls the "chatter" service "summary" endpoint.
func NewSummaryHandler(
	endpoint goa.Endpoint,
	mux goahttp.Muxer,
	dec func(*http.Request) goahttp.Decoder,
	enc func(context.Context, http.ResponseWriter) goahttp.Encoder,
	eh func(context.Context, http.ResponseWriter, error),
	up goahttp.Upgrader,
	connConfigFn goahttp.ConnConfigureFunc,
) http.Handler {
	var (
		decodeRequest = DecodeSummaryRequest(mux, dec)
		encodeError   = EncodeSummaryError(enc)
	)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), goahttp.AcceptTypeKey, r.Header.Get("Accept"))
		ctx = context.WithValue(ctx, goa.MethodKey, "summary")
		ctx = context.WithValue(ctx, goa.ServiceKey, "chatter")
		payload, err := decodeRequest(r)
		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				eh(ctx, w, err)
			}
			return
		}

		var cancel context.CancelFunc
		{
			ctx, cancel = context.WithCancel(ctx)
		}
		v := &chatter.SummaryEndpointInput{
			Stream: &summaryServerStream{
				upgrader:     up,
				connConfigFn: connConfigFn,
				cancel:       cancel,
				w:            w,
				r:            r,
			},
			Payload: payload.(*chatter.SummaryPayload),
		}
		_, err = endpoint(ctx, v)

		if err != nil {
			if _, ok := err.(websocket.HandshakeError); ok {
				return
			}
			if err := encodeError(ctx, w, err); err != nil {
				eh(ctx, w, err)
			}
			return
		}
	})
}

// MountSubscribeHandler configures the mux to serve the "chatter" service
// "subscribe" endpoint.
func MountSubscribeHandler(mux goahttp.Muxer, h http.Handler) {
	f, ok := h.(http.HandlerFunc)
	if !ok {
		f = func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		}
	}
	mux.Handle("GET", "/subscribe", f)
}

// NewSubscribeHandler creates a HTTP handler which loads the HTTP request and
// calls the "chatter" service "subscribe" endpoint.
func NewSubscribeHandler(
	endpoint goa.Endpoint,
	mux goahttp.Muxer,
	dec func(*http.Request) goahttp.Decoder,
	enc func(context.Context, http.ResponseWriter) goahttp.Encoder,
	eh func(context.Context, http.ResponseWriter, error),
	up goahttp.Upgrader,
	connConfigFn goahttp.ConnConfigureFunc,
) http.Handler {
	var (
		decodeRequest = DecodeSubscribeRequest(mux, dec)
		encodeError   = EncodeSubscribeError(enc)
	)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), goahttp.AcceptTypeKey, r.Header.Get("Accept"))
		ctx = context.WithValue(ctx, goa.MethodKey, "subscribe")
		ctx = context.WithValue(ctx, goa.ServiceKey, "chatter")
		payload, err := decodeRequest(r)
		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				eh(ctx, w, err)
			}
			return
		}

		var cancel context.CancelFunc
		{
			ctx, cancel = context.WithCancel(ctx)
		}
		v := &chatter.SubscribeEndpointInput{
			Stream: &subscribeServerStream{
				upgrader:     up,
				connConfigFn: connConfigFn,
				cancel:       cancel,
				w:            w,
				r:            r,
			},
			Payload: payload.(*chatter.SubscribePayload),
		}
		_, err = endpoint(ctx, v)

		if err != nil {
			if _, ok := err.(websocket.HandshakeError); ok {
				return
			}
			if err := encodeError(ctx, w, err); err != nil {
				eh(ctx, w, err)
			}
			return
		}
	})
}

// MountHistoryHandler configures the mux to serve the "chatter" service
// "history" endpoint.
func MountHistoryHandler(mux goahttp.Muxer, h http.Handler) {
	f, ok := h.(http.HandlerFunc)
	if !ok {
		f = func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		}
	}
	mux.Handle("GET", "/history", f)
}

// NewHistoryHandler creates a HTTP handler which loads the HTTP request and
// calls the "chatter" service "history" endpoint.
func NewHistoryHandler(
	endpoint goa.Endpoint,
	mux goahttp.Muxer,
	dec func(*http.Request) goahttp.Decoder,
	enc func(context.Context, http.ResponseWriter) goahttp.Encoder,
	eh func(context.Context, http.ResponseWriter, error),
	up goahttp.Upgrader,
	connConfigFn goahttp.ConnConfigureFunc,
) http.Handler {
	var (
		decodeRequest = DecodeHistoryRequest(mux, dec)
		encodeError   = EncodeHistoryError(enc)
	)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), goahttp.AcceptTypeKey, r.Header.Get("Accept"))
		ctx = context.WithValue(ctx, goa.MethodKey, "history")
		ctx = context.WithValue(ctx, goa.ServiceKey, "chatter")
		payload, err := decodeRequest(r)
		if err != nil {
			if err := encodeError(ctx, w, err); err != nil {
				eh(ctx, w, err)
			}
			return
		}

		var cancel context.CancelFunc
		{
			ctx, cancel = context.WithCancel(ctx)
		}
		v := &chatter.HistoryEndpointInput{
			Stream: &historyServerStream{
				upgrader:     up,
				connConfigFn: connConfigFn,
				cancel:       cancel,
				w:            w,
				r:            r,
			},
			Payload: payload.(*chatter.HistoryPayload),
		}
		_, err = endpoint(ctx, v)

		if err != nil {
			if _, ok := err.(websocket.HandshakeError); ok {
				return
			}
			if err := encodeError(ctx, w, err); err != nil {
				eh(ctx, w, err)
			}
			return
		}
	})
}

// Send streams instances of "string" to the "echoer" endpoint websocket
// connection.
func (s *echoerServerStream) Send(v string) error {
	var err error
	// Upgrade the HTTP connection to a websocket connection only once. Connection
	// upgrade is done here so that authorization logic in the endpoint is executed
	// before calling the actual service method which may call Send().
	s.once.Do(func() {
		var conn *websocket.Conn
		conn, err = s.upgrader.Upgrade(s.w, s.r, nil)
		if err != nil {
			return
		}
		if s.connConfigFn != nil {
			conn = s.connConfigFn(conn, s.cancel)
		}
		s.conn = conn
	})
	if err != nil {
		return err
	}
	res := v
	return s.conn.WriteJSON(res)
}

// Recv reads instances of "string" from the "echoer" endpoint websocket
// connection.
func (s *echoerServerStream) Recv() (string, error) {
	var (
		rv  string
		msg *string
		err error
	)
	// Upgrade the HTTP connection to a websocket connection only once. Connection
	// upgrade is done here so that authorization logic in the endpoint is executed
	// before calling the actual service method which may call Recv().
	s.once.Do(func() {
		var conn *websocket.Conn
		conn, err = s.upgrader.Upgrade(s.w, s.r, nil)
		if err != nil {
			return
		}
		if s.connConfigFn != nil {
			conn = s.connConfigFn(conn, s.cancel)
		}
		s.conn = conn
	})
	if err != nil {
		return rv, err
	}
	if err = s.conn.ReadJSON(&msg); err != nil {
		return rv, err
	}
	if msg == nil {
		return rv, io.EOF
	}
	body := *msg
	return body, nil
}

// Close closes the "echoer" endpoint websocket connection.
func (s *echoerServerStream) Close() error {
	var err error
	// Upgrade the HTTP connection to a websocket connection only once. Connection
	// upgrade is done here so that authorization logic in the endpoint is executed
	// before calling the actual service method which may call Close().
	s.once.Do(func() {
		var conn *websocket.Conn
		conn, err = s.upgrader.Upgrade(s.w, s.r, nil)
		if err != nil {
			return
		}
		if s.connConfigFn != nil {
			conn = s.connConfigFn(conn, s.cancel)
		}
		s.conn = conn
	})
	if err != nil {
		return err
	}
	if err = s.conn.WriteControl(
		websocket.CloseMessage,
		websocket.FormatCloseMessage(websocket.CloseNormalClosure, "server closing connection"),
		time.Now().Add(time.Second),
	); err != nil {
		return err
	}
	return s.conn.Close()
}

// Recv reads instances of "string" from the "listener" endpoint websocket
// connection.
func (s *listenerServerStream) Recv() (string, error) {
	var (
		rv  string
		msg *string
		err error
	)
	// Upgrade the HTTP connection to a websocket connection only once. Connection
	// upgrade is done here so that authorization logic in the endpoint is executed
	// before calling the actual service method which may call Recv().
	s.once.Do(func() {
		var conn *websocket.Conn
		conn, err = s.upgrader.Upgrade(s.w, s.r, nil)
		if err != nil {
			return
		}
		if s.connConfigFn != nil {
			conn = s.connConfigFn(conn, s.cancel)
		}
		s.conn = conn
	})
	if err != nil {
		return rv, err
	}
	if err = s.conn.ReadJSON(&msg); err != nil {
		return rv, err
	}
	if msg == nil {
		return rv, io.EOF
	}
	body := *msg
	return body, nil
}

// Close closes the "listener" endpoint websocket connection.
func (s *listenerServerStream) Close() error {
	var err error
	// Upgrade the HTTP connection to a websocket connection only once. Connection
	// upgrade is done here so that authorization logic in the endpoint is executed
	// before calling the actual service method which may call Close().
	s.once.Do(func() {
		var conn *websocket.Conn
		conn, err = s.upgrader.Upgrade(s.w, s.r, nil)
		if err != nil {
			return
		}
		if s.connConfigFn != nil {
			conn = s.connConfigFn(conn, s.cancel)
		}
		s.conn = conn
	})
	if err != nil {
		return err
	}
	if err = s.conn.WriteControl(
		websocket.CloseMessage,
		websocket.FormatCloseMessage(websocket.CloseNormalClosure, "server closing connection"),
		time.Now().Add(time.Second),
	); err != nil {
		return err
	}
	return s.conn.Close()
}

// SendAndClose streams instances of "chatter.ChatSummaryCollection" to the
// "summary" endpoint websocket connection and closes the connection.
func (s *summaryServerStream) SendAndClose(v chatter.ChatSummaryCollection) error {
	defer s.conn.Close()
	res := chatter.NewViewedChatSummaryCollection(v, "default")
	body := NewChatSummaryResponseCollection(res.Projected)
	return s.conn.WriteJSON(body)
}

// Recv reads instances of "string" from the "summary" endpoint websocket
// connection.
func (s *summaryServerStream) Recv() (string, error) {
	var (
		rv  string
		msg *string
		err error
	)
	// Upgrade the HTTP connection to a websocket connection only once. Connection
	// upgrade is done here so that authorization logic in the endpoint is executed
	// before calling the actual service method which may call Recv().
	s.once.Do(func() {
		var conn *websocket.Conn
		conn, err = s.upgrader.Upgrade(s.w, s.r, nil)
		if err != nil {
			return
		}
		if s.connConfigFn != nil {
			conn = s.connConfigFn(conn, s.cancel)
		}
		s.conn = conn
	})
	if err != nil {
		return rv, err
	}
	if err = s.conn.ReadJSON(&msg); err != nil {
		return rv, err
	}
	if msg == nil {
		return rv, io.EOF
	}
	body := *msg
	return body, nil
}

// Send streams instances of "chatter.Event" to the "subscribe" endpoint
// websocket connection.
func (s *subscribeServerStream) Send(v *chatter.Event) error {
	var err error
	// Upgrade the HTTP connection to a websocket connection only once. Connection
	// upgrade is done here so that authorization logic in the endpoint is executed
	// before calling the actual service method which may call Send().
	s.once.Do(func() {
		var conn *websocket.Conn
		conn, err = s.upgrader.Upgrade(s.w, s.r, nil)
		if err != nil {
			return
		}
		if s.connConfigFn != nil {
			conn = s.connConfigFn(conn, s.cancel)
		}
		s.conn = conn
	})
	if err != nil {
		return err
	}
	res := v
	body := NewSubscribeResponseBody(res)
	return s.conn.WriteJSON(body)
}

// Close closes the "subscribe" endpoint websocket connection.
func (s *subscribeServerStream) Close() error {
	var err error
	// Upgrade the HTTP connection to a websocket connection only once. Connection
	// upgrade is done here so that authorization logic in the endpoint is executed
	// before calling the actual service method which may call Close().
	s.once.Do(func() {
		var conn *websocket.Conn
		conn, err = s.upgrader.Upgrade(s.w, s.r, nil)
		if err != nil {
			return
		}
		if s.connConfigFn != nil {
			conn = s.connConfigFn(conn, s.cancel)
		}
		s.conn = conn
	})
	if err != nil {
		return err
	}
	if err = s.conn.WriteControl(
		websocket.CloseMessage,
		websocket.FormatCloseMessage(websocket.CloseNormalClosure, "server closing connection"),
		time.Now().Add(time.Second),
	); err != nil {
		return err
	}
	return s.conn.Close()
}

// Send streams instances of "chatter.ChatSummary" to the "history" endpoint
// websocket connection.
func (s *historyServerStream) Send(v *chatter.ChatSummary) error {
	var err error
	// Upgrade the HTTP connection to a websocket connection only once. Connection
	// upgrade is done here so that authorization logic in the endpoint is executed
	// before calling the actual service method which may call Send().
	s.once.Do(func() {
		respHdr := make(http.Header)
		respHdr.Add("goa-view", s.view)
		var conn *websocket.Conn
		conn, err = s.upgrader.Upgrade(s.w, s.r, respHdr)
		if err != nil {
			return
		}
		if s.connConfigFn != nil {
			conn = s.connConfigFn(conn, s.cancel)
		}
		s.conn = conn
	})
	if err != nil {
		return err
	}
	res := chatter.NewViewedChatSummary(v, s.view)
	var body interface{}
	switch s.view {
	case "tiny":
		body = NewHistoryResponseBodyTiny(res.Projected)
	case "default", "":
		body = NewHistoryResponseBody(res.Projected)
	}
	return s.conn.WriteJSON(body)
}

// Close closes the "history" endpoint websocket connection.
func (s *historyServerStream) Close() error {
	var err error
	// Upgrade the HTTP connection to a websocket connection only once. Connection
	// upgrade is done here so that authorization logic in the endpoint is executed
	// before calling the actual service method which may call Close().
	s.once.Do(func() {
		var conn *websocket.Conn
		conn, err = s.upgrader.Upgrade(s.w, s.r, nil)
		if err != nil {
			return
		}
		if s.connConfigFn != nil {
			conn = s.connConfigFn(conn, s.cancel)
		}
		s.conn = conn
	})
	if err != nil {
		return err
	}
	if err = s.conn.WriteControl(
		websocket.CloseMessage,
		websocket.FormatCloseMessage(websocket.CloseNormalClosure, "server closing connection"),
		time.Now().Add(time.Second),
	); err != nil {
		return err
	}
	return s.conn.Close()
}

// SetView sets the view to render the chatter.ChatSummary type before sending
// to the "history" endpoint websocket connection.
func (s *historyServerStream) SetView(view string) {
	s.view = view
}
