package weheat

import "fmt"

// APIError represents a non-2xx response from the Weheat API.
type APIError struct {
	StatusCode int
	Body       []byte
}

func (e *APIError) Error() string {
	if len(e.Body) == 0 {
		return fmt.Sprintf("weheat: api error %d", e.StatusCode)
	}
	return fmt.Sprintf("weheat: api error %d: %s", e.StatusCode, string(e.Body))
}
