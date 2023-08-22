package source

import (
	"context"

	"github.com/kong11213613/haremicro/client"
	"github.com/kong11213613/haremicro/config/encoder"
	"github.com/kong11213613/haremicro/config/encoder/json"
)

type Options struct {
	// Encoder
	Encoder encoder.Encoder

	// for alternative data
	Context context.Context

	// Client to use for RPC
	Client client.Client
}

type Option func(o *Options)

func NewOptions(opts ...Option) Options {
	options := Options{
		Encoder: json.NewEncoder(),
		Context: context.Background(),
		Client:  client.DefaultClient,
	}

	for _, o := range opts {
		o(&options)
	}

	return options
}

// WithEncoder sets the source encoder
func WithEncoder(e encoder.Encoder) Option {
	return func(o *Options) {
		o.Encoder = e
	}
}

// WithClient sets the source client
func WithClient(c client.Client) Option {
	return func(o *Options) {
		o.Client = c
	}
}