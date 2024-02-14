import express from 'express';
import './loadEnv.js';
// import email from './routes/email.js';
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
const app = express();

const allowedOrigins = ["http://localhost:3000", "http://localhost:3000/signup", "http://localhost:3000/main"]

app.use(express.json());
app.use(cors({allowedOrigins, credentials: true, origin: true}));
app.use(cookieParser());

mongoose.connect(process.env.MONGODB_URI)
    .then((result) => app.listen(PORT))
    .catch((err) => console.log(err));

// load the routes
app.get("*", checkUser);
app.use("/api/register", register);
app.use("/api/login", login);
app.use("/api/logout", logout);
app.use("/api/shorten", requireAuth , shorten);
app.use("/api/short", short)
app.use("/api/user", requireAuth, user);


// app.listen(PORT, () => {
//     console.log(`Server listening on port: ${PORT}`);
// })
