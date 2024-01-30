import { MongoClient } from "mongodb";

const connectionString = process.env.MONGODB_URI || "";
const client = new MongoClient(connectionString);

let conn;
try {
    conn = await client.connect();
} catch (e) {
    console.log(e);
}

let db = conn.db(process.env.DATABASE_NAME);
export default db;