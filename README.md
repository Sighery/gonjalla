# Unofficial Golang library for the Njalla API

[Njalla][] is a privacy-oriented domain name registration service. Recently
they released their [official API][].

This Golang library covers _some_ methods of that API. For the moment, those
are:
* `list-domains`
* `get-domain`
* `list-records`
* `add-record`
* `edit-record`
* `remove-record`
* `find-domains`
* `check-task`
* `register-domain`
* `list-servers`
* `list-server-images`
* `list-server-types`
* `stop-server`
* `restart-server`
* `reset-server`
* `add-server`
* `remove-server`

**TO NOTE**: Even though `record` methods are implemented, I'm fairly certain
they'll fail (silently or not) in some cases. I deal mostly with `TXT`, `MX`,
`A` and `AAAA` DNS records. Some records have different/more variables, and
since I don't use them I decided against implementing them. Chances are the
methods will fail when trying to deal with those types of records (like `SSH`
records).

The code is fairly simple, and most methods are tested by using mocks on
the API request. The mocked returned data is based on the same data the API
returns.

Feel free to send in a PR to add more (or to cover all types of DNS records),
as long as they're tested and documented.

### Usage

Most of the methods are pretty self-explanatory if you read the documentation
and read the function signature. Some of the most "complex" operations, like
creating/updating records can be figured out from the tests. But here's some
examples:

```golang
package main

import (
	"fmt"

	"github.com/Sighery/gonjalla"
)

func main() {
	token := "api-token"
	domain := "your-domain"

	// 1. Listing records in a domain
	records, err1 := gonjalla.ListRecords(token, domain)
	if err1 != nil {
		fmt.Println(err1)
	}

	// It will print an array of gonjalla.Record
	fmt.Println(records)


	// 2. Adding a new record to a domain
	priority := 10
	adding := gonjalla.Record{
		Name: "@",
		Type: "MX",
		Content: "testing.com",
		TTL: 10800,
		Priority: &priority,
	}

	confirmation, err2 := gonjalla.AddRecord(token, domain, adding)
	if err2 != nil {
		fmt.Println(err2)
	}

	// confirmation will be a gonjalla.Record with the response from the
	// server, so it should be pretty similar to your starting
	// gonjalla.Record but this will contain the ID filled in, which is
	// needed for updates
	fmt.Println(confirmation)


	// 3. Updating a record of a given domain
	// Let's assume we got the ID of the record created in step 2
	// and we want to change either some or all fields
	id_we_look_for := confirmation.ID

	// The edit method updates all the fields due to limitations of the
	// API. Get the record from the API if you only want to change some,
	// but not all, fields
	for _, record := range records {
		if record.ID == id_we_look_for {
			record.Content = "edited-value"
			record.TTL = 900

			err3 := gonjalla.EditRecord(token, domain, record)
			if err3 != nil {
				fmt.Println(err3)
			}
		}
	}

	// If you don't care about overwriting previous values
	new_priority := 20
	editing := gonjalla.Record{
		ID: id_we_look_for,
		Name: "@",
		Type: "MX",
		Content: "edited-thing",
		TTL: 300,
		Priority: &new_priority,
	}

	err4 := gonjalla.EditRecord(token, domain, editing)
	if err4 != nil {
		fmt.Println(err4)
	}
}
```

Some actual code making use of this library (mainly dealing with records) can
also be seen at the [Njalla Terraform provider].

[Njalla]: https://njal.la
[official API]: https://njal.la/api/
[Njalla Terraform provider]: https://github.com/Sighery/terraform-provider-njalla
