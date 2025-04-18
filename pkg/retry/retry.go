package retry

import (
	"errors"
	"fmt"
	"time"
)

// Retry function that retries the provided function up to maxRetries times with a delay between attempts.
func Retry(attempts int, sleep time.Duration, fn func() error) error {
	for i := 0; i < attempts; i++ {
		err := fn()
		if err == nil {
			return nil
		}
		fmt.Printf("Attempt %d failed; retrying in %v...\n", i+1, sleep)
		time.Sleep(sleep)
	}
	return errors.New("all retry attempts failed")
}
