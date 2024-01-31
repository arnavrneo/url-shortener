import express from "express";
import { register } from "../controllers/authController.js";
// import {checkSchema} from "express-validator";
// import {userSchema} from "../models/User.js";

const router = express.Router();

router.post("/", register);

export default router;