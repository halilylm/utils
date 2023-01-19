package broker

import (
	"context"
	"crypto/tls"
	"github.com/halilylm/utils/logger"
)

type Options struct {
	Brokers   []string
	Secure    bool
	Logger    logger.Logger
	TLSConfig *tls.Config
	Context   context.Context
}

type PublishOptions struct {
	Context context.Context
}

type SubscribeOptions struct {
	AutoAck bool
	Topic   string
	Context context.Context
}

type Option func(options *Options)
type PublishOption func(options *PublishOptions)
type SubscribeOption func(options *SubscribeOptions)

func AutoAck(auto bool) SubscribeOption {
	return func(s *SubscribeOptions) {
		s.AutoAck = auto
	}
}

func Topic(topic string) SubscribeOption {
	return func(s *SubscribeOptions) {
		s.Topic = topic
	}
}

func SubscribeContext(ctx context.Context) SubscribeOption {
	return func(s *SubscribeOptions) {
		s.Context = ctx
	}
}

func PublishContext(ctx context.Context) PublishOption {
	return func(p *PublishOptions) {
		p.Context = ctx
	}
}

func Brokers(brokers ...string) Option {
	return func(o *Options) {
		o.Brokers = brokers
	}
}

func Secure(s bool) Option {
	return func(o *Options) {
		o.Secure = s
	}
}

func Logger(l logger.Logger) Option {
	return func(o *Options) {
		o.Logger = l
	}
}

func TLSConfig(t *tls.Config) Option {
	return func(o *Options) {
		o.TLSConfig = t
	}
}

func Context(c context.Context) Option {
	return func(o *Options) {
		o.Context = c
	}
}
