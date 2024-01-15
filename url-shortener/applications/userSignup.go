package applications

func SignUp() {
	// Get the email/pass off req body
	var body struct {
		Name     string
		Email    string
		Password string
	}

	if body.Name != "" {
		return
	}
	// Hash the pass

	// Create the user

	// Respond
}
