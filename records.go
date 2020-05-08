package gonjalla

import "encoding/json"

// ValidTTL is an array containing all the valid TTL values
var ValidTTL = []int{60, 300, 900, 3600, 10800, 21600, 86400}

// Record struct contains data returned by `list-records`
type Record struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Type     string `json:"type"`
	Content  string `json:"content"`
	TTL      int    `json:"ttl"`
	Priority *int   `json:"prio,omitempty"`
}

// ListRecords returns a listing of all records for a given domain
func ListRecords(token string, domain string) ([]Record, error) {
	params := map[string]interface{}{
		"domain": domain,
	}
	data, err := Request(token, "list-records", params)
	if err != nil {
		return nil, err
	}

	type Response struct {
		Records []Record `json:"records"`
	}

	var response Response
	err = json.Unmarshal(data, &response)
	if err != nil {
		return nil, err
	}

	return response.Records, nil
}

// AddRecord adds a given record to a given domain.
// Returns a new Record struct, containing the response from the API if
// successful. This response will have some fields like ID (which can only
// be known after the execution) filled.
func AddRecord(token string, domain string, record Record) (Record, error) {
	marshal, err := json.Marshal(record)
	if err != nil {
		return Record{}, err
	}

	params := map[string]interface{}{
		"domain": domain,
	}
	err = json.Unmarshal(marshal, &params)
	if err != nil {
		return Record{}, err
	}

	data, err := Request(token, "add-record", params)
	if err != nil {
		return Record{}, err
	}

	var response Record
	err = json.Unmarshal(data, &response)
	if err != nil {
		return Record{}, err
	}

	return response, nil
}

// RemoveRecord removes a given record from a given domain.
// If there are no errors it will return `nil`.
func RemoveRecord(token string, domain string, id int) error {
	params := map[string]interface{}{
		"domain": domain,
		"id": id,
	}

	_, err := Request(token, "remove-record", params)
	if err != nil {
		return err
	}

	return nil
}

// EditRecord edits a record for a given domain.
// This function is fairly dumb. It takes in a `Record` struct, and uses all
// its filled fields to send to Njalla.
// So, if you want to only change a given field, get the `Record` object from
// say ListRecords, change the one field you want, and then pass that here.
func EditRecord(token string, domain string, record Record) error {
	marshal, err := json.Marshal(record)
	if err != nil {
		return err
	}

	params := map[string]interface{}{
		"domain": domain,
	}
	err = json.Unmarshal(marshal, &params)
	if err != nil {
		return err
	}

	_, err = Request(token, "edit-record", params)
	if err != nil {
		return err
	}

	return nil
}
