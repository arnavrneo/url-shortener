// import { MongoClient } from "mongodb";
//
// const connectionString = process.env.MONGODB_URI || "";
// const client = new MongoClient(connectionString);
//
// let conn;
// try {
//     conn = await client.connect();
// } catch (e) {
//     console.log(e);
// }
//
// let db = conn.db(process.env.DATABASE_NAME);
// export default db;

import mongoose from "mongoose";

const connectionString = process.env.MONGODB_URI || "";
mongoose.Promise = global.Promise;
try {
    mongoose.connect("mongodb+srv://arnav:urlshortener@cluster0.8hlrx0u.mongodb.net/?retryWrites=true&w=majority");
    console.log("connected successfully");
} catch (e) {
    console.log(e);
}

export default mongoose;