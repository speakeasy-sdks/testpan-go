<!-- Start SDK Example Usage -->
```go
package main

import (
	"context"
	testpango "github.com/speakeasy-sdks/testpan-go"
	"github.com/speakeasy-sdks/testpan-go/pkg/models/operations"
	"github.com/speakeasy-sdks/testpan-go/pkg/models/shared"
	"log"
)

func main() {
	s := testpango.New(
		testpango.WithSecurity(shared.Security{
			Password: "",
			Username: "",
		}),
	)

	ctx := context.Background()
	res, err := s.Users.DeleteUsersUserID(ctx, operations.DeleteUsersUserIDRequest{
		UserID: "2d4aef6d-76db-4c57-a2a3-fe8efd3db6e2",
	})
	if err != nil {
		log.Fatal(err)
	}

	if res.StatusCode == http.StatusOK {
		// handle response
	}
}

```
<!-- End SDK Example Usage -->