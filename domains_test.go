package gonjalla

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/Sighery/gonjalla/mocks"
)

func TestListDomainsExpected(t *testing.T) {
	token := "test-token"
	Client = &mocks.MockClient{}

	testData := `{
		"jsonrpc": "2.0",
		"result": {
			"domains": [
				{
					"name": "testing1.com",
					"status": "active",
					"expiry": "2021-02-20T19:38:48Z"
				},
				{
					"name": "testing2.com",
					"status": "inactive",
					"expiry": "2021-02-20T19:38:48Z"
				}
			]
		}
	}`
	r := ioutil.NopCloser(bytes.NewReader([]byte(testData)))

	mocks.GetDoFunc = func(*http.Request) (*http.Response, error) {
		return &http.Response{
			StatusCode: 200,
			Body:       r,
		}, nil
	}

	domains, err := ListDomains(token)
	if err != nil {
		t.Error(err)
	}

	expectedTime, _ := time.Parse(time.RFC3339, "2021-02-20T19:38:48Z")

	expected := []Domain{
		{
			Name:   "testing1.com",
			Status: "active",
			Expiry: expectedTime,
		},
		{
			Name:   "testing2.com",
			Status: "inactive",
			Expiry: expectedTime,
		},
	}

	assert.Equal(t, domains, expected)
}

func TestListDomainsError(t *testing.T) {
	token := "test-token"
	Client = &mocks.MockClient{}

	testData := `{
		"jsonrpc": "2.0",
		"error": {
			"code": 0,
			"message": "Testing error"
		}
	}`
	r := ioutil.NopCloser(bytes.NewReader([]byte(testData)))

	mocks.GetDoFunc = func(*http.Request) (*http.Response, error) {
		return &http.Response{
			StatusCode: 200,
			Body:       r,
		}, nil
	}

	domains, err := ListDomains(token)
	assert.Nil(t, domains)
	assert.Error(t, err)
}

func TestGetDomainExpected(t *testing.T) {
	token := "test-token"
	domain := "testing.com"
	Client = &mocks.MockClient{}

	testData := `{
		"jsonrpc": "2.0",
		"result": {
			"name": "testing.com",
			"status": "active",
			"expiry": "2021-02-20T19:38:48Z",
			"locked": true,
			"mailforwarding": false,
			"max_nameservers": 10
		}
	}`
	r := ioutil.NopCloser(bytes.NewReader([]byte(testData)))

	mocks.GetDoFunc = func(*http.Request) (*http.Response, error) {
		return &http.Response{
			StatusCode: 200,
			Body:       r,
		}, nil
	}

	result, err := GetDomain(token, domain)
	if err != nil {
		t.Error(err)
	}

	expectedTime, _ := time.Parse(time.RFC3339, "2021-02-20T19:38:48Z")
	locked := true
	mailforwarding := false
	maxNameservers := 10

	expected := Domain{
		Name:           domain,
		Status:         "active",
		Expiry:         expectedTime,
		Locked:         &locked,
		Mailforwarding: &mailforwarding,
		MaxNameservers: &maxNameservers,
	}

	assert.Equal(t, result, expected)
}

func TestGetDomainError(t *testing.T) {
	token := "test-token"
	domain := "testing.com"
	Client = &mocks.MockClient{}

	testData := `{
		"jsonrpc": "2.0",
		"error": {
			"code": 0,
			"message": "Testing error"
		}
	}`
	r := ioutil.NopCloser(bytes.NewReader([]byte(testData)))

	mocks.GetDoFunc = func(*http.Request) (*http.Response, error) {
		return &http.Response{
			StatusCode: 200,
			Body:       r,
		}, nil
	}

	_, err := GetDomain(token, domain)
	assert.Error(t, err)
}

func TestFindDomainsExpected(t *testing.T) {
	token := "test-token"
	query := "testing"
	Client = &mocks.MockClient{}

	testData := `{
		"result": {
			"jsonrpc": "2.0",
			"domains": [
				{
					"name": "testing.com",
					"status": "taken",
					"price": 45
				},
				{
					"name": "testing.net",
					"status": "available",
					"price": 30
				},
				{
					"name": "testing.rocks",
					"status": "in progress",
					"price": 15
				},
				{
					"name": "testing.express",
					"status": "failed",
					"price": 75
				}
			]
		}
	}`
	r := ioutil.NopCloser(bytes.NewReader([]byte(testData)))

	mocks.GetDoFunc = func(*http.Request) (*http.Response, error) {
		return &http.Response{
			StatusCode: 200,
			Body:       r,
		}, nil
	}

	domains, err := FindDomains(token, query)
	if err != nil {
		t.Error(err)
	}

	expected := []MarketDomain{
		{
			Name:   "testing.com",
			Status: "taken",
			Price:  45,
		},
		{
			Name:   "testing.net",
			Status: "available",
			Price:  30,
		},
		{
			Name:   "testing.rocks",
			Status: "in progress",
			Price:  15,
		},
		{
			Name:   "testing.express",
			Status: "failed",
			Price:  75,
		},
	}

	assert.Equal(t, domains, expected)
}

func TestFindDomainsError(t *testing.T) {
	token := "test-token"
	query := "testing"
	Client = &mocks.MockClient{}

	testData := `{
		"jsonrpc": "2.0",
		"error": {
			"code": 0,
			"message": "Testing error"
		}
	}`
	r := ioutil.NopCloser(bytes.NewReader([]byte(testData)))

	mocks.GetDoFunc = func(*http.Request) (*http.Response, error) {
		return &http.Response{
			StatusCode: 200,
			Body:       r,
		}, nil
	}

	domains, err := FindDomains(token, query)
	assert.Nil(t, domains)
	assert.Error(t, err)
}
