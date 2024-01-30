import express from 'express';
import './loadEnv.js';
import email from './routes/email.js';

const PORT = process.env.PORT;
const app = express();

app.use(express.json());

// load the routes
app.use("/email", email);

// global error handling
app.use((err, _req, res, next) => {
    res.status(500).send("Error occured");
})

app.listen(PORT, () => {
    console.log(`Server listening on port: ${PORT}`);
})
