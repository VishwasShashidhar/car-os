package e2e

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/VishwasShashidhar/car-os/src/controllers"
	"github.com/go-playground/assert/v2"
	"github.com/go-resty/resty/v2"
)

func TestHomeGreetUser(t *testing.T) {

	client := resty.New()

	resp, _ := client.R().Get("http://localhost:8080/home")

	assert.Equal(t, 200, resp.StatusCode())

	homeResponse := controllers.Home{}

	err := json.Unmarshal(resp.Body(), &homeResponse)
	if err != nil {
		fmt.Println(err)
		return
	}

	assert.Equal(t, "Hello, user!", homeResponse.Message)
}
