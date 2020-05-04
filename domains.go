package gonjalla

import (
	"encoding/json"
	"time"
)

// Domain struct contains data returned by `list-domains` and `get-domains`
type Domain struct {
	Name           string    `json:"name"`
	Status         string    `json:"status"`
	Expiry         time.Time `json:"expiry"`
	Locked         *bool     `json:"locked,omitempty"`
	Mailforwarding *bool     `json:"mailforwarding,omitempty"`
	MaxNameservers *int      `json:"max_nameservers,omitempty"`
}

// ListDomains returns a listing of domains with minimal data
func ListDomains(token string) ([]Domain, error) {
	params := map[string]interface{}{}

	data, err := Request(token, "list-domains", params)
	if err != nil {
		return nil, err
	}

	type Response struct {
		Domains []Domain `json:"domains"`
	}

	var response Response
	err = json.Unmarshal(data, &response)
	if err != nil {
		return nil, err
	}

	return response.Domains, nil
}

// GetDomain returns detailed information for each domain
func GetDomain(token string, domain string) (Domain, error) {
	params := map[string]interface{}{
		"domain": domain,
	}

	data, err := Request(token, "get-domain", params)
	if err != nil {
		return Domain{}, err
	}

	var domainStruct Domain
	err = json.Unmarshal(data, &domainStruct)
	if err != nil {
		return Domain{}, err
	}

	return domainStruct, nil
}
