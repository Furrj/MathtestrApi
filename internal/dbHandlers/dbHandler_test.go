package dbHandlers

import (
	"context"
	"testing"
)

func TestDbHandler(t *testing.T) {
	dbHandler := InitDBHandler()
	defer dbHandler.DB.Close(context.Background())

	t.Run("Ping connection", func(t *testing.T) {
		if err := dbHandler.DB.Ping(context.Background()); err != nil {
			t.Errorf("Error initializing database: %+v\n", err)
		}
	})
	t.Run("Existing CheckIfUsernameExists", func(t *testing.T) {
		got, err := dbHandler.CheckIfUsernameExists("a")

		if err != nil {
			t.Errorf("Error querying for username: %+v\n", err)
		}
		if !got {
			t.Errorf("Username exists but returned false")
		}
	})
	t.Run("Non-existent CheckIfUsernameExists", func(t *testing.T) {
		got, err := dbHandler.CheckIfUsernameExists("")

		if err != nil {
			t.Errorf("Error querying for username: %+v\n", err)
		}
		if got {
			t.Errorf("Username doesn't exist but returned true")
		}
	})
	t.Run("GetUserByUsername", func(t *testing.T) {
		got, err := dbHandler.GetUserByUsername("a")
		want := "a"
		if err != nil {
			t.Errorf("Error when querying for user: %+v\n", err)
		}
		if got.Username != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})
}
