# EmailRep

This is a go library for interacting with the [EmailRep](https://emailrep.io) service.


[![Go Report Card](https://goreportcard.com/badge/github.com/kaiiyer/emailrep)](https://goreportcard.com/report/github.com/kaiiyer/emailrep)
[![Documentation](https://godoc.org/github.com/kaiiyer/emailrep?status.svg)](http://godoc.org/github.com/kaiiyer/emailrep)

**Rep** function will return the reputation of the email address passed as the argument 

### Install

```bash
go get github.com/kaiiyer/emailrep
```

### Import

```go
import (
	"github.com/kaiiyer/emailrep"
)
```

## Example Usage

```go
package main

import (
	"github.com/kaiiyer/emailrep"
)

func main() {
	var r string
	r = emailrep.Rep("bill@microsoft.com")
	print(r)
}
```

With this example, the response would be:

```json
  {
  "email": "bill@microsoft.com",
  "reputation": "high",
  "suspicious": false,
  "references": 79,
  "details": {
    "blacklisted": false,
    "malicious_activity": false,
    "malicious_activity_recent": false,
    "credentials_leaked": true,
    "credentials_leaked_recent": false,
    "data_breach": true,
    "first_seen": "07/01/2008",
    "last_seen": "05/24/2019",
    "domain_exists": true,
    "domain_reputation": "high",
    "new_domain": false,
    "days_since_domain_creation": 10341,
    "suspicious_tld": false,
    "spam": false,
    "free_provider": false,
    "disposable": false,
    "deliverable": true,
    "accept_all": true,
    "valid_mx": true,
    "spoofable": false,
    "spf_strict": true,
    "dmarc_enforced": true,
    "profiles": [
      "myspace",
      "spotify",
      "twitter",
      "pinterest",
      "flickr",
      "linkedin",
      "vimeo",
      "angellist"
    ]
  }
}
```

**emailrep** also accepts arguments of type `error` that implements the built in `error` interface and the response:

```json
{
  "message": "Sorry we got hit by an error!"
}
```

Full API docs can be found [here](https://docs.emailrep.io).
