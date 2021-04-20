package function

import (
	"fmt"
)

// Handle a serverless request
func Handle(req []byte) (interface{}, error) {

	fmt.Printf("Hello, Go Interface. You said: %s", string(req))

	return nil, nil
}
