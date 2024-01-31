import express from "express";
import {shortRedirect} from "../controllers/authController.js";

const router = express.Router();

router.use("/:id", shortRedirect);

export default router;