import mongoose, {model, Schema} from "mongoose";
import User from "../routes/user.js";

const connectionString = process.env.MONGODB_URI || "";
mongoose.Promise = global.Promise;
try {
    mongoose.connect(process.env.MONGODB_URI);
    console.log("connected successfully");

    const schema = new Schema({
        username: String,
        email: String,
        password: String,
    })
    const user = model('User', schema);

    const userDetail = new user({
        username: "a",
        email: "b",
        password: "c",
    })

    await userDetail.save();

    const find = await user.findOne({});
    console.log(find);
} catch (e) {
    console.log(e);
}

export default mongoose;