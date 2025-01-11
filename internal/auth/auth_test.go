package auth

import (
	"fmt"
	"net/http"
	"testing"
)

// TestHelloName calls greetings.Hello with a name, checking
// for a valid return value.
func TestGetAPIKey(t *testing.T) {
	noHeaders, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	bearer, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	apikey, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}
	bearer.Header.Add("Authorization", "Bearer this-doesnot-matter")
	apikey.Header.Add("Authorization", "ApiKey my-cool-key")

	cases := []struct {
		request  http.Request
		key      string
		errorMsg string
	}{
		{
			request:  *noHeaders,
			key:      "",
			errorMsg: "no authorization header included",
		},
		{
			request:  *bearer,
			key:      "",
			errorMsg: "malformed authorization header",
		},
		{
			request:  *apikey,
			key:      "my-cool-key",
			errorMsg: "",
		},
	}

	for i, c := range cases {
		t.Run(fmt.Sprintf("Test case %v", i), func(t *testing.T) {
			actual, err := GetAPIKey(c.request.Header)
			if len(c.errorMsg) > 0 {
				if err == nil {
					t.Fatal("Expected an error to be thrown.")
				}

				if err.Error() != c.errorMsg {
					t.Fatalf(
						"Incorrect error message.  Got: %s Expected: %s",
						err.Error(),
						c.errorMsg,
					)
				}
			}

			if actual != c.key {
				t.Fatalf("Incorrect key. Got: %s Expected: %s", actual, c.key)
			}
		})
	}
}
