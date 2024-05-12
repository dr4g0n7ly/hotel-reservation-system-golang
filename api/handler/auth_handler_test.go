package api

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/dr4g0n7ly/hotel-management-system-golang/db"
	"github.com/dr4g0n7ly/hotel-management-system-golang/db/fixtures"
	"github.com/dr4g0n7ly/hotel-management-system-golang/types"
	"github.com/gofiber/fiber/v2"
)

func InsertTestUser(t *testing.T, userStore db.UserStore) *types.User {
	user, err := types.NewUserFromParams(types.CreateUserParams{
		Email:     "james@foo.com",
		FirstName: "james",
		LastName:  "foo",
		Password:  "securepassword",
	})
	if err != nil {
		t.Fatal(err)
	}
	_, err = userStore.InsertUser(context.Background(), user)
	if err != nil {
		t.Fatal(err)
	}
	return user
}

func TestAuthenticateSuccess(t *testing.T) {
	tdb := setup(t)
	defer tdb.teardown(t)

	insertedUser := fixtures.AddUser(tdb.Store, "james", "foo", false)

	app := fiber.New()
	authHandler := NewAuthHandler(*tdb.Store)
	app.Post("/auth", authHandler.HandleAuthenticate)

	params := AuthParams{
		Email:    "james@foo.com",
		Password: "james_foo",
	}
	b, _ := json.Marshal(params)

	req := httptest.NewRequest("POST", "/auth", bytes.NewReader(b))
	req.Header.Add("Content-Type", "application/json")
	res, err := app.Test(req)
	if err != nil {
		t.Fatal(err)
	}
	if res.StatusCode != http.StatusOK {
		t.Fatalf("expected http status of 200 but got %d", res.StatusCode)
	}

	var authResp AuthResponse
	if err := json.NewDecoder(res.Body).Decode(&authResp); err != nil {
		t.Error(err)
	}
	if authResp.Token == "" {
		t.Fatalf("expected the JWT token to be present in the authResponse")
	}
	insertedUser.EncryptedPassword = ""
	if !reflect.DeepEqual(insertedUser, authResp.User) {
		fmt.Printf("\nINSERTED USER: %+v", insertedUser)
		fmt.Printf("\nAUTHENTI USER: %+v", authResp.User)
		t.Fatal("expected user to be insertedUser")
	}
}

func TestAuthenticateIncorrectPassword(t *testing.T) {
	tdb := setup(t)
	defer tdb.teardown(t)

	insertedUser := InsertTestUser(t, tdb.User)
	_ = insertedUser

	app := fiber.New()
	authHandler := NewAuthHandler(*tdb.Store)
	app.Post("/auth", authHandler.HandleAuthenticate)

	params := AuthParams{
		Email:    "james@foo.com",
		Password: "wrongsecurepassword",
	}
	b, _ := json.Marshal(params)

	req := httptest.NewRequest("POST", "/auth", bytes.NewReader(b))
	req.Header.Add("Content-Type", "application/json")
	res, err := app.Test(req)
	if err != nil {
		t.Fatal(err)
	}
	if res.StatusCode != http.StatusUnauthorized {
		t.Fatalf("expected http status of 401 but got %d", res.StatusCode)
	}

	var genResp GenericResponse
	if err := json.NewDecoder(res.Body).Decode(&genResp); err != nil {
		t.Fatal(err)
	}
	if genResp.Type != "error" {
		t.Fatalf("expected gen response type to be error but got %s", genResp.Type)
	}
	if genResp.Msg != "Invalid credentials" {
		t.Fatalf("expected gen response msg to be <Invalid credentials> but got %s", genResp.Msg)
	}
}
