import express from "express";
import {shorten} from "../controllers/authController.js";

const router = express.Router();

router.get("/", shorten);

export default router;