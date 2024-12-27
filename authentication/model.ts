import mongoose from "mongoose";
import bcrypt from "bcrypt";

const UserSchema = new mongoose.Schema(
    {
        name: {
            type: String,
            required: true,
        },
        email: {
            type: String,
            required: true,
        },
        username: {
            type: String,
            required: true,
        },
        avatar: {
            type: String,
            default: "",
        },
        password: {
            type: String,
            required: true,
        }
    },
    {
        timestamps: true,
    }
);

UserSchema.pre("save", async function (next) {
    const saltRounds = 10;
    this.password = await bcrypt.hash(this.password, saltRounds);
    next();
});


// Create and export the User model
const User = mongoose.model("User", UserSchema);

export default User;