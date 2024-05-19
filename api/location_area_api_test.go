package api

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetLocationAreas(t *testing.T) {
	// Happy path test case
	t.Run("GetLocationAreas with valid requestUrl", func(t *testing.T) {
		// Create a mock server
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// Send a mock JSON response
			mockResponse := `{
				"count": 1054,
				"next": "next_url",
				"previous": "previous_url",
				"results": [
				  {
					"name": "mt-coronet-1f-route-207",
					"url": "mock_url_1"
				  }
				]
			  }`
			fmt.Fprintln(w, mockResponse)
		}))
		defer mockServer.Close()

		// Create a new client with the mock server URL
		client := &Client{httpClient: mockServer.Client()}
		requestUrl := mockServer.URL

		// Call the GetLocationAreas method
		res, err := client.GetLocationAreas(&requestUrl)

		// Validate the response and error
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}
		next_url := "next_url"
		previous_url := "previous_url"

		expected := LocationAreaResponse{
			Count:    1054,
			Next:     &next_url,
			Previous: &previous_url,
			Results: []LocationArea{
				{
					Name: "mt-coronet-1f-route-207",
					Url:  "mock_url_1",
				},
			},
		}

		if res.isEqual(expected) {
			t.Errorf("Expected %v, got %v", expected, res)
		}
	})

	// Edge case: requestUrl is nil
	// t.Run("request with nil requestUrl", func(t *testing.T) {
	// 	// Create a new client
	// 	client := &Client{httpClient: &http.Client{}}

	// 	// Call the GetLocationAreas method with nil requestUrl
	// 	_, err := client.GetLocationAreas(nil)

	// 	// Validate the error
	// 	if err == nil {
	// 		t.Error("Expected error, got nil")
	// 	}
	// })

	// Edge case: server response body is empty
	t.Run("GetLocationAreas with empty server response body", func(t *testing.T) {
		// Create a mock server
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// Send an empty response
		}))
		defer mockServer.Close()

		// Create a new client with the mock server URL
		client := &Client{httpClient: mockServer.Client()}
		requestUrl := mockServer.URL

		// Call the GetLocationAreas method
		_, err := client.GetLocationAreas(&requestUrl)

		// Validate the error
		if err == nil {
			t.Error("Expected error, got nil")
		}
	})

	// Edge case: invalid JSON response from the server
	t.Run("GetLocationAreas with invalid JSON response", func(t *testing.T) {
		// Create a mock server
		mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// Send an invalid JSON response
			fmt.Fprintln(w, "invalid JSON")
		}))
		defer mockServer.Close()

		// Create a new client with the mock server URL
		client := &Client{httpClient: mockServer.Client()}
		requestUrl := mockServer.URL

		// Call the GetLocationAreas method
		_, err := client.GetLocationAreas(&requestUrl)

		// Validate the error
		if err == nil {
			t.Error("Expected error, got nil")
		}
	})
}

func Test_requestWithUrl(t *testing.T) {
	type args struct {
		requestUrl *string
	}
	tests := []struct {
		name    string
		args    args
		want    *http.Request
		wantErr bool
	}{
		{
			name: "Happy path, URL provided",
			args: args{
				requestUrl: getStringPointer("https://example.com"),
			},
			want:    getHTTPRequest("GET", "https://example.com"),
			wantErr: false,
		},
		{
			name: "URL is nil",
			args: args{
				requestUrl: nil,
			},
			want:    getHTTPRequest("GET", "https://pokeapi.co/api/v2/location-area"),
			wantErr: false,
		},
		{
			name: "Invalid URL",
			args: args{
				requestUrl: getStringPointer("invalid-url"),
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := createRequestWithUrl(tt.args.requestUrl)
			if (err != nil) && !tt.wantErr {
				t.Errorf("requestWithUrl() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

		})
	}
}

func getStringPointer(s string) *string {
	return &s
}

func getHTTPRequest(method, url string) *http.Request {
	req, _ := http.NewRequest(method, url, nil)
	return req
}
