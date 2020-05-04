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

**TO NOTE**: Even though `record` methods are implemented, I'm fairly certain
they'll fail (silently or not) in some cases. I deal mostly with `TXT`, `MX`,
`A` and `AAAA` DNS records. Some records have different/more variables, and
since I don't use them I decided against implementing them. Chances are the
methods will fail when trying to deal with those types of records (like `SSH`
records).

The code is fairly simple, and all the methods are tested by using mocks on
the API request. The mocked returned data is based on the same data the API
returns.

These methods cover my needs, but feel free to send in a PR to add more (or to
cover all types of DNS records), as long as they're all tested and documented.

### Usage

```golang
package main

import (
	"fmt"

	"github.com/Sighery/gonjalla"
)

func main() {
	token := "api-token"
	domain := "your-domain"

	records, err := ListRecords(token, domain)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(records)
}
```

[Njalla]: https://njal.la
[official API]: https://njal.la/api/
