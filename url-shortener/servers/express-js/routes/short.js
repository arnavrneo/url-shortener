import express from "express";
import {shortRedirect} from "../controllers/authController.js";

const router = express.Router();

router.get("/:id", shortRedirect);

export default router;