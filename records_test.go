package gonjalla

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/Sighery/gonjalla/mocks"
)

func TestListRecordsExpected(t *testing.T) {
	token := "test-token"
	domain := "testing.com"
	Client = &mocks.MockClient{}

	testData := `{
		"jsonrpc": "2.0",
		"result": {
			"records": [
				{
					"id": 1337,
					"name": "_acme-challenge",
					"type": "TXT",
					"content": "long-string",
					"ttl": 10800
				},
				{
					"id": 1338,
					"name": "@",
					"type": "A",
					"content": "1.2.3.4",
					"ttl": 3600
				},
				{
					"id": 1339,
					"name": "@",
					"type": "AAAA",
					"content": "2001:0DB8:0000:0000:0000:8A2E:0370:7334",
					"ttl": 900
				},
				{
					"id": 1340,
					"name": "@",
					"type": "MX",
					"content": "mail.protonmail.ch",
					"ttl": 300,
					"prio": 10
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

	records, err := ListRecords(token, domain)
	if err != nil {
		t.Error(err)
	}

	priority := 10

	expected := []Record{
		{
			ID:      1337,
			Name:    "_acme-challenge",
			Type:    "TXT",
			Content: "long-string",
			TTL:     10800,
		},
		{
			ID:      1338,
			Name:    "@",
			Type:    "A",
			Content: "1.2.3.4",
			TTL:     3600,
		},
		{
			ID:      1339,
			Name:    "@",
			Type:    "AAAA",
			Content: "2001:0DB8:0000:0000:0000:8A2E:0370:7334",
			TTL:     900,
		},
		{
			ID:       1340,
			Name:     "@",
			Type:     "MX",
			Content:  "mail.protonmail.ch",
			TTL:      300,
			Priority: &priority,
		},
	}

	assert.Equal(t, records, expected)
}

func TestListRecordsError(t *testing.T) {
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

	records, err := ListRecords(token, domain)
	if records != nil {
		t.Error("records isn't nil")
	}

	assert.Error(t, err)
}

func TestAddRecordExpected(t *testing.T) {
	token := "test-token"
	domain := "testing.com"
	Client = &mocks.MockClient{}

	testData := `{
		"jsonrpc": "2.0",
		"result": {
			"id": 1337,
			"name": "@",
			"type": "MX",
			"content": "testing.com",
			"ttl": 10800,
			"prio": 10
		}
	}`
	r := ioutil.NopCloser(bytes.NewReader([]byte(testData)))

	mocks.GetDoFunc = func(*http.Request) (*http.Response, error) {
		return &http.Response{
			StatusCode: 200,
			Body:       r,
		}, nil
	}

	priority := 10

	adding := Record{
		Name:     "@",
		Type:     "MX",
		Content:  "testing.com",
		TTL:      10800,
		Priority: &priority,
	}

	record, err := AddRecord(token, domain, adding)
	if err != nil {
		t.Error(err)
	}

	expected := Record{
		ID:       1337,
		Name:     "@",
		Type:     "MX",
		Content:  "testing.com",
		TTL:      10800,
		Priority: &priority,
	}

	assert.Equal(t, record, expected)
}

func TestAddRecordError(t *testing.T) {
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

	priority := 10
	adding := Record{
		Name:     "@",
		Type:     "MX",
		Content:  "testing.com",
		TTL:      10800,
		Priority: &priority,
	}

	_, err := AddRecord(token, domain, adding)
	assert.Error(t, err)
}

func TestRemoveRecordExpected(t *testing.T) {
	token := "test-token"
	domain := "testing.com"
	id := 1337
	Client = &mocks.MockClient{}

	testData := `{
		"jsonrpc": "2.0",
		"result": {}
	}`
	r := ioutil.NopCloser(bytes.NewReader([]byte(testData)))

	mocks.GetDoFunc = func(*http.Request) (*http.Response, error) {
		return &http.Response{
			StatusCode: 200,
			Body:       r,
		}, nil
	}

	err := RemoveRecord(token, domain, id)
	assert.Nil(t, err)
}

func TestRemoveRecordError(t *testing.T) {
	token := "test-token"
	domain := "testing.com"
	id := 1337
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

	err := RemoveRecord(token, domain, id)
	assert.Error(t, err)
}

func TestEditRecordExpected(t *testing.T) {
	token := "test-token"
	domain := "testing.com"
	Client = &mocks.MockClient{}

	testData := `{
		"jsonrpc": "2.0",
		"result": {
			"id": 1337,
			"name": "@",
			"type": "MX",
			"content": "testing.com",
			"ttl": 10800,
			"prio": 10
		}
	}`
	r := ioutil.NopCloser(bytes.NewReader([]byte(testData)))

	mocks.GetDoFunc = func(*http.Request) (*http.Response, error) {
		return &http.Response{
			StatusCode: 200,
			Body:       r,
		}, nil
	}

	priority := 10
	editing := Record{
		ID: 1337,
		Name: "@",
		Type: "MX",
		Content: "testing.com",
		TTL: 10800,
		Priority: &priority,
	}

	err := EditRecord(token, domain, editing)
	assert.Nil(t, err)
}

func TestEditRecordError(t *testing.T) {
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

	priority := 10
	editing := Record{
		ID: 1337,
		Name: "@",
		Type: "MX",
		Content: "testing.com",
		TTL: 10800,
		Priority: &priority,
	}

	err := EditRecord(token, domain, editing)
	assert.Error(t, err)
}
