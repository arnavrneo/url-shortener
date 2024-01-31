export function register(req, res) {
    const {username, email, password } = req.body;
    console.log(username, email, password);
    res.status(200).send("new signup")
}

export function login(req, res) {
    const { username, email, password } = req.body;
    res.status(200).send("new login reached")
}

export function logout(req, res) {
    res.status(200).send("logout successful.")
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

