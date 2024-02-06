import jwt from "jsonwebtoken";

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