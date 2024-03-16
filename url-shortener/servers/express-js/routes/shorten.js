import express from "express";
import {shorten} from "../controllers/authController.js";
import multer from "multer";

const router = express.Router();
const upload = multer();

router.post("/", upload.none(), shorten);

export default router;