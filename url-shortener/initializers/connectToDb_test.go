package initializers

import "testing"

func TestConnectToDb(t *testing.T) {
	ConnectToDb("mongodb+srv://arnav:urlshortener@cluster0.8hlrx0u.mongodb.net/?retryWrites=true&w=majority") //TODO: remove the uri
}
