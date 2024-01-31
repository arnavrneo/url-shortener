export function register(req, res) {
    res.status(200).send("<h1>Reached the signup page</h1>")
}

export function login(req, res) {
    res.status(200).send("<h1>Reached the login page</h1>")
}

export function logout(req, res) {
    res.status(200).send("<h1>Reached the logout page</h1>")
}

export function shorten(req, res) {
    res.status(200).send("<h1>Reached the shorten page</h1>")
}

export function shortRedirect(req, res) {
    const param = req.params.id;
    res.status(200).send(`<h1>Redirect from here. Got id: ${param}</h1>`)
}

export function getUser(req, res) {
    res.status(200).send("<h1>Fetch the user data from here</h1>")
}

