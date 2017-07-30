//go:generate bitfanDoc
// HTTPPoller allows you to call an HTTP Endpoint, decode the output of it into an event
package httppoller

import (
	"encoding/json"

	"github.com/parnurzeal/gorequest"
	"github.com/vjeantet/bitfan/codecs"
	"github.com/vjeantet/bitfan/processors"
)

func New() processors.Processor {
	return &processor{opt: &options{}}
}

type options struct {
	// Add a field to an event
	Add_field map[string]interface{}

	// Add any number of arbitrary tags to your event.
	// This can help with processing later.
	Tags []string

	// Add a type field to all events handled by this input
	Type string

	// The codec used for input data. Input codecs are a convenient method for decoding
	// your data before it enters the input, without needing a separate filter in your bitfan pipeline
	// @Default "plain"
	// @Type codec
	Codec codecs.Codec `mapstructure:"codec"`

	// Use CRON or BITFAN notation
	// @ExampleLS interval => "every_10s"
	Interval string `mapstructure:"interval"`

	// Http Method
	// @Default : "GET"
	Method string `mapstructure:"method"`

	// URL
	// @ExampleLS : url=> "http://google.fr"
	Url string `mapstructure:"url" validate:"required"`
}

type processor struct {
	processors.Base

	opt     *options
	request *gorequest.SuperAgent
}

func (p *processor) Configure(ctx processors.ProcessorContext, conf map[string]interface{}) error {
	return p.ConfigureAndValidate(ctx, conf, p.opt)
}

func (p *processor) Start(e processors.IPacket) error {
	p.request = gorequest.New()
	return nil
}

func (p *processor) Tick(e processors.IPacket) error {
	return p.Receive(e)
}

func (p *processor) Receive(e processors.IPacket) error {
	var (
		errs []error
		resp gorequest.Response
		body string
	)

	switch p.opt.Method {
	case "GET":
		resp, body, errs = p.request.Get(p.opt.Url).End()
	default:
		p.Logger.Warnf("Method %s not implemented", p.opt.Method)
		return nil
	}

	if errs != nil {
		p.Logger.Warnf("while http requesting %s : %#v", p.opt.Url, errs)
		return nil
	}
	if resp.StatusCode >= 400 {
		p.Logger.Warnf("http response code %s : %d (%s)", p.opt.Url, resp.StatusCode, resp.Status)
		return nil
	}
	e = p.NewPacket(p.opt.Url, map[string]interface{}{})
	e.SetMessage(p.opt.Url)
	json.Unmarshal([]byte(body), e.Fields())

	p.Send(e, 0)

	return nil
}
