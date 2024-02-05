import {mongoose} from "mongoose";
import bcrypt from "bcrypt";

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

// fire a func after saving to db
// userSchema.post('save', function (doc, next) {
//
//     next();
// })

// fire a func before saving to db
userSchema.pre('save', async function (next) {
    const salt = await bcrypt.genSalt();
    this.password = await bcrypt.hash(this.password, salt);
    next();
})

// static method to login user
userSchema.statics.login = async function(username, email, password) {
    const user = await this.findOne({ email: email }) // value will be placed inside the user constant
    if (user) {
        const auth = await bcrypt.compare(password, user.password);
        if (auth) {
            return user;
        }
        throw Error('incorrect password');
    }
    throw Error('incorrect email');
}

const User = mongoose.model("url_short_express_backend", userSchema);

export default User;