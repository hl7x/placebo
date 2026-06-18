package random

import (
	"fmt"
	"math/rand"
)

type Provider struct {
	ID            string
	FirstName     string
	MiddleName    string
	MiddleInitial string
	LastName      string
}

func NewProvider() *Provider {
	p := &Provider{}

	p.ProviderName().ProviderID()

	return p
}

func (p *Provider) ProviderName() *Provider {

	name := NewName()
	p.FirstName = name.FirstName
	p.MiddleName = name.MiddleName
	p.LastName = name.LastName

	return p
}
func (p *Provider) ProviderID() *Provider {

	randomId := rand.Intn(1000000000)

	p.ID = fmt.Sprint(randomId)

	return p
}
