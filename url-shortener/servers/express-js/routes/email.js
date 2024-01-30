import express from "express";
import db from "../db/conn.js";

const router = express.Router();

router.get("/", async (req, res) => {
    let coll = await db.collection(process.env.DATABASE_COLLECTION);
    let results = await coll.findOne({ email: "a@a.com" });
    console.log(results)
    res.send(results).status(200);
})

export default router;