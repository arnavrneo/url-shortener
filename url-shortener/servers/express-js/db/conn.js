// import mongoose, {model, Schema} from "mongoose";
//
// const connectionString = process.env.MONGODB_URI || "";
//
// async function connectToDatabase() {
//     mongoose.Promise = global.Promise;
//     let db;
//     try {
//         db = await mongoose.connect(connectionString);
//         console.log("connected successfully");
//     } catch (e) {
//         console.log(e);
//     }
//     return db;
// }
//
//
// export default connectToDatabase;