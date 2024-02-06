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
import requireAuth from "./middleware/authMiddleware.js";


const PORT = process.env.PORT;
const app = express();

app.use(express.json());
app.use(cookieParser());

mongoose.connect(process.env.MONGODB_URI)
    .then((result) => app.listen(PORT))
    .catch((err) => console.log(err));

// load the routes
app.use("/api/signup", register);
app.use("/api/login", login);
app.use("/api/logout", logout);
app.use("/api/shorten", requireAuth , shorten);
app.use("/api/short", short)
app.use("/api/user", requireAuth, user);

// global error handling
app.use((err, _req, res, next) => {
    res.status(500).send("Error occurred");
})

// app.listen(PORT, () => {
//     console.log(`Server listening on port: ${PORT}`);
// })
