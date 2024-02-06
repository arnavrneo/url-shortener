import jwt from "jsonwebtoken";
import User from "../models/User.js";

const requireAuth = (req, res, next) => {
    const token = req.cookies.jwt;

    if (token) {
        jwt.verify(token, process.env.SECRET, (err, decodedToken) => {
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
        jwt.verify(token, process.env.SECRET, async (err, decodedToken) => {
            if (err) {
                next();
            } else {
                console.log(decodedToken)
                // let user = await User.findById(decodedToken.email);
                next();
            }
        })
    } else {

    }
}