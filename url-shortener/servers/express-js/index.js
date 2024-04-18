import express from 'express';
import './loadEnv.js';
import register from "./routes/register.js";
import login from "./routes/login.js";
import logout from "./routes/logout.js";
import shorten from "./routes/shorten.js";
import user from "./routes/user.js";
import short from "./routes/short.js";
import {mongoose} from "mongoose";
import cookieParser from "cookie-parser";
import requireAuth, {checkUser} from "./middleware/authMiddleware.js";
import cors from 'cors';

const PORT = process.env.PORT;
const ORIGIN = process.env.ORIGIN;
const MONGODB_URI = process.env.MONGODB_URI;

const app = express();

const allowedOrigins = [ORIGIN]
//const allowedOrigins = ["*"]

app.use(express.json());
app.use(cors({origin: allowedOrigins, credentials: true}));
// app.use(function(req, res, next) {
//     res.setHeader('Access-Control-Allow-Origin', '*');
//     res.setHeader('Access-Control-Allow-Methods', 'GET, POST, PUT, DELETE');
//     res.setHeader('Access-Control-Allow-Headers', 'Content-Type');
//     res.setHeader('Access-Control-Allow-Credentials', true);
//     next();
// });
app.use(cookieParser());

mongoose.connect(MONGODB_URI)
    .then((result) => app.listen(PORT))
    .catch((err) => console.log(err));

// load the routes
app.get("*", checkUser);
app.get("/", (req, res) => {
    res.status(200).json({"message": "the server is running."})
})
app.use("/api/register", register);
app.use("/api/login", login);
app.use("/api/logout", logout);
app.use("/api/shorten", requireAuth , shorten);
app.use("/api/short", short)
app.use("/api/user", requireAuth, user);


// app.listen(PORT, () => {
//     console.log(`Server listening on port: ${PORT}`);
// })
