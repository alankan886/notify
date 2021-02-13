package mailgun

import (
	"github.com/mailgun/mailgun-go/v4"
)

// Option describes a functional parameter for the Mailgun constructor.
type Option func(*Mailgun)

// WithEurope sets the API Mailgun base url to Europe region.
func WithEurope() Option {
	return func(m *Mailgun) {
		m.client.SetAPIBase(mailgun.APIBaseEU)
	}
}
