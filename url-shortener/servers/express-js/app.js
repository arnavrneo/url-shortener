import "./loadEnv.js";

const express = require('express')

const app = express()

// static files and middleware
// app.use(express.static('./static-path'))

// MONGODB connection
const client = new Mono

app.get('/', (req,res) => { // callback func
    res.status(200).send('HomePage');
})


app.get("/short/:id/bla/:poo", (req,res) => {
    const {id, poo} = req.params;
    res.status(200).send(`Fetched key: ${id} ${poo}`)
})

app.post("/api/login", (req, res) => {

})

// all should come at last
app.all("*", (req,res) => {
    res.status(404).send("<h1>Resource not found.</h1>")
})


// POST: /api/login
// POST: /api/logout
// POST: /api/register
// POST: /api/shorten (require auth)
// GET: /short/:id (redis)
// GET: /api/user (require auth)

const PORT = process.env.PORT
app.listen(PORT, () => {
    console.log(`Server is running on port ${PORT}`);
})
