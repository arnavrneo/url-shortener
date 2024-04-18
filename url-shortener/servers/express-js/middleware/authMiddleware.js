import jwt from "jsonwebtoken";
import User from "../models/User.js";

const SECRET = process.env.SECRET;

const requireAuth = (req, res, next) => {
    const token = req.cookies.jwt;

    if (token) {
        jwt.verify(token, SECRET, (err, decodedToken) => {
            if (err) {
                res.status(404).json({"error": "unauthorized"})
            } else {
                next();
            }
        });
    } else {
        res.status(404).json({"error": "cookie not found"})
    }
}

export default requireAuth;

export const checkUser = async (req, res, next) => {
    const token = req.cookies.jwt;

    if (token) {
        jwt.verify(token, SECRET, async (err, decodedToken) => {
            if (err) {
                res.locals.user = null;
                next();
            } else {
                res.locals.user = await User.findOne({email: decodedToken.email});
                next();
            }
        })
    } else {
        res.locals.user = null;
        next();
    }
}