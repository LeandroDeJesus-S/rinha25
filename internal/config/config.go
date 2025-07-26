package config

import (
	"os"
	"runtime"
	"strconv"
)

const (
	DefaultAppPort = 9000
	DefaultQueueSize = 5000
	DefaultMaxRetries = 3
)

type configOption func(*config)

type config struct {
	AppPort int
	RedisAddr string
	
	PaymentProcessorDefaultUrl string
	PaymentProcessorFallbackUrl string

	Workers int
	QueueSize int
	MaxRetries int
}

func New(opt ...configOption) *config {
	cfg := defaultConfig()

	for _, o := range opt {
		o(cfg)
	}

	return cfg
}

func defaultConfig() *config {
	appPort, err := strconv.Atoi(os.Getenv("APP_PORT"))
	if err != nil {
		appPort = DefaultAppPort
	}
	return &config{
		AppPort: appPort,
		RedisAddr: os.Getenv("REDIS_ADDR"),

		Workers: runtime.NumCPU(),
		QueueSize: DefaultQueueSize,
		MaxRetries: DefaultMaxRetries,

		PaymentProcessorDefaultUrl: os.Getenv("PAYMENT_PROCESSOR_URL_DEFAULT"),
		PaymentProcessorFallbackUrl: os.Getenv("PAYMENT_PROCESSOR_URL_FALLBACK"),
	}
}

func WithAppPort(p int) configOption {
	return func(cfg *config) {
		cfg.AppPort = p
	}
}

func WithRedisAddr(a string) configOption {
	return func(cfg *config) {
		cfg.RedisAddr = a
	}
}

func WithPaymentProcessorDefaultUrl(u string) configOption {
	return func(c *config) {
		c.PaymentProcessorDefaultUrl = u
	}
}

func WithPaymentProcessorFallbackUrl(u string) configOption {
	return func(c *config) {
		c.PaymentProcessorFallbackUrl = u
	}
}

func WithWorkers(n int) configOption {
	return func(cfg *config) {
		cfg.Workers = n
	}
}

func WithQueueSize(s int) configOption {
	return func(cfg *config) {
		cfg.QueueSize = s
	}
}

func WithMaxRetries(n int) configOption {
	return func(cfg *config) {
		cfg.MaxRetries = n
	}
}