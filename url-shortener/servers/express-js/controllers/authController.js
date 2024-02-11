import User from "../models/User.js";
import jwt from "jsonwebtoken";

const handleErrors = (err) => {
    console.log(err.message, err.code);
    let errors = { username: '', email: '', password: ''};

    // incorrect email
    if (err.message === 'incorrect email') {
        errors.email = 'the email is not registered';
    }

    // incorrect password
    if (err.message === 'incorrect password') {
        errors.password = 'the password is incorrect';
    }

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
        const token = createToken(user.email)

        // using lib
        res.cookie('jwt', token, {
            maxAge: maxAge * 1000, // in milliseconds
            //secure: true, // only over https
            httpOnly: true // cant access it from frontend
        });
        res.status(201).json({'message': 'register successful'});

    } catch (e) {
        const errors = handleErrors(e);
        res.status(400).json({ errors });
    }
}

export async function login(req, res) {
    const { username, email, password } = req.body;

    try {
        const user = await User.login(username, email, password);
        const token = createToken(user.email)

        // using lib
        res.cookie('jwt', token, {
            maxAge: maxAge * 1000, // in milliseconds
            //secure: true, // only over https
            httpOnly: true // cant access it from frontend
        });
        res.status(200).json({"message": "login request successful",});
    } catch (err) {
        const errors = handleErrors(err)
        res.status(400).json({ errors });
    }

    // way to set cookie
    //res.setHeader('Set-Cookie', 'newUser=true');
    //res.status(201).json({ 'message': 'login successful'})
}

export function logout(req, res) {
    res.cookie('jwt', '', { maxAge: 1 });
    res.status(200).send({"message": "successfully logged out",})
}

// func GenerateShortKey() string {
//     const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
//     const keyLength = 6
//
//     shortKey := make([]byte, keyLength)
//     for i := range shortKey {
//         shortKey[i] = charset[rand.Intn(len(charset))]
//     }
//     return string(shortKey)
// }

function generateShortkey() {
    const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789";
    const keyLength = 6;

    const shortKey = []
    for (let i=0; i<keyLength; i++) {
        shortKey.push(charset[Math.floor(Math.random() * charset.length)]);
    }

    return shortKey
}

export function shorten(req, res) {
    const shortKey = generateShortkey();
    res.status(200).send("<h1>Reached the shorten page</h1>")
}

export function shortRedirect(req, res) {
    const param = req.params.id;
    res.status(200).send(`<h1>Redirect from here. Got id: ${param}</h1>`)
}

export async function getUser(req, res) {
    // console.log("EMAIL: ", res.locals.user);
    res.status(200).json({ username: res.locals.user.username, email: res.locals.user.email });
}

const maxAge = 60 * 60 // in seconds
const createToken = (email) => {
    return jwt.sign(
        { email },
        process.env.SECRET,
        { expiresIn: maxAge })

}
