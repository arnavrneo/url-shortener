package initializers

import "testing"

func TestConnectToDb(t *testing.T) {
	err := ConnectToDb("mongodb+srv://arnav:urlshortener@cluster0.8hlrx0u.mongodb.net/?retryWrites=true&w=majority") //TODO: remove the uri
	if err != nil {
		t.Error("database connection not established")
	}
}
