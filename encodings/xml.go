package encodings

import (
	"context"
	"log"

	xml "goa.design/examples/encodings/gen/xml"
)

// xml service example implementation.
// The example methods log the requests and return zero values.
type xmlsrvc struct {
	logger *log.Logger
}

// NewXML returns the xml service implementation.
func NewXML(logger *log.Logger) xml.Service {
	return &xmlsrvc{logger}
}

// Convert implements convert.
func (s *xmlsrvc) Convert(ctx context.Context, p *xml.ConvertPayload) (res *xml.JSONObject, err error) {
	res = &xml.JSONObject{}
	s.logger.Print("xml.convert")
	s.logger.Print("payload", p)
	res = &xml.JSONObject{
		Name:        p.Name,
		Description: p.Description,
	}
	return
}
