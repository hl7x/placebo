package random

import "placebo/internal/tools"

type Provider struct {
	ID        string
	FirstName string
	LastName  string
}

func ProviderName() *Provider {

	p := &Provider{}

	first := FIRSTNAME
	last := LASTNAME

	number := tools.RandomSelector(first)
	number2 := tools.RandomSelector(last)

	p.FirstName = first[number]
	p.LastName = last[number2]

	return p
}

func ProviderID() *Provider {

}
