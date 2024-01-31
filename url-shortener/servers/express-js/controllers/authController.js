import User from "../models/User.js";

export async function register(req, res) {
    const {username, email, password } = req.body;
    console.log(username, email, password);

    try {
        const user = await User.create({ username, email, password }); // async func call
        res.status(201).json(user);
    } catch (e) {
        console.log(e);
        res.status(400).send("error, user cannot be created.")
    }
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

