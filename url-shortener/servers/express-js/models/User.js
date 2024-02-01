import {mongoose} from "mongoose";

const userSchema = new mongoose.Schema({
    username: {
        type: String,
        required: [true, "Please enter a username"]
    },
    email: {
        type: String,
        required: [true, "Please enter an email"],
        unique: true,
        lowercase: true
    },
    password: {
        type: String,
        required: [true, "Please enter a password"],
    }
})

const User = mongoose.model("url_short_express_backend", userSchema);

export default User;