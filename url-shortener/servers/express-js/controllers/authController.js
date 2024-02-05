import User from "../models/User.js";
import jwt from "jsonwebtoken";

const handleErrors = (err) => {
    console.log(err.message, err.code);
    let errors = { username: '', email: '', password: ''};

    // duplicate error code
    if (err.code === 11000) {
        errors.email = 'that email is already registered';
        return errors;
    }

    // validation errors
    if (err.message.includes('url_short_express_backend validation failed:')) {
        Object.values(err.errors).forEach(({properties}) => {
            errors[properties.path] = properties.message;
        });
    }

    return errors;
}

export async function register(req, res) {
    const {username, email, password } = req.body;

    try {
        const user = await User.create({ username, email, password }); // async func call
        res.status(201).json(user);
    } catch (e) {
        const errors = handleErrors(e);
        res.status(400).json({ errors });
    }
}

export function login(req, res) {
    const { username, email, password } = req.body;

    const token = createToken(email)

    // using lib
    res.cookie('jwt', token, {
        maxAge: maxAge * 1000, // in milliseconds
        //secure: true, // only over https
        httpOnly: true // cant access it from frontend
    });

    // way to set cookie
    //res.setHeader('Set-Cookie', 'newUser=true');
    res.status(201).json({ 'message': 'login successful'})
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

const maxAge = 60 * 60 // in seconds
const createToken = (email) => {
    return jwt.sign(
        { email },
        process.env.SECRET,
        { expiresIn: maxAge })

}
