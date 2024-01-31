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


const PORT = process.env.PORT;
const app = express();

app.use(express.json());

mongoose.connect(process.env.MONGODB_URI)
    .then((result) => app.listen(process.env.PORT))
    .catch((err) => console.log(err));

// load the routes
// app.use("/email", email);
// TODO: convert the api to /api/*
app.use("/signup", register);
app.use("/login", login);
app.use("/logout", logout);
app.use("/shorten", shorten);
app.use("/short", short)
app.use("/user", user);

// global error handling
app.use((err, _req, res, next) => {
    res.status(500).send("Error occured");
})

// app.listen(PORT, () => {
//     console.log(`Server listening on port: ${PORT}`);
// })
