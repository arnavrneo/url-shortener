import express from "express";
import {logout} from "../controllers/authController.js";

const router = express.Router();

router.post("/", logout);

export default router;