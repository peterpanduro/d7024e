package helpers

import (
	"testing"
)

// TestNewHTTPError tests the creation of a new HTTPError
func TestNewHTTPError(t *testing.T) {
	// Define test cases
	tests := []struct {
		code     int
		message  string
		expected *HTTPError
	}{
		{404, "Not Found", &HTTPError{404, "Not Found"}},
		{500, "Internal Server Error", &HTTPError{500, "Internal Server Error"}},
	}

	for _, test := range tests {
		// Create a new HTTPError using NewHTTPError
		err := NewHTTPError(test.code, test.message)

		// Check if the returned error matches the expected result
		if err.Code != test.expected.Code || err.Message != test.expected.Message {
			t.Errorf("Expected HTTPError to have Code %d and Message '%s', but got Code %d and Message '%s'",
				test.expected.Code, test.expected.Message, err.Code, err.Message)
		}
	}
}

// TestErrorMethod tests the Error method of the HTTPError struct
func TestErrorMethod(t *testing.T) {
	// Define a test case
	err := &HTTPError{Code: 400, Message: "Bad Request"}

	// Call the Error() method
	result := err.Error()

	// Verify the result matches the message
	if result != "Bad Request" {
		t.Errorf("Expected error message 'Bad Request', but got '%s'", result)
	}
}
