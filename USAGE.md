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
	res, err := s.APISecurityPolicies.DeleteAPISecurityPolicyPolicyID(ctx, operations.DeleteAPISecurityPolicyPolicyIDRequest{
		PolicyID: "04ae1a0e-dcb7-4d2b-b7a6-f7ca105f8c92",
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