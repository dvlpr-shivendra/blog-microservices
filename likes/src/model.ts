// Import mongoose
import mongoose from "mongoose";

// Define the Task schema
const LikeSchema = new mongoose.Schema(
    {
        userId: {
            type: Number,
            required: true,
        },
        postId: {
            type: Number,
            required: true,
        },
    },
    {
        timestamps: true,  // Add createdAt and updatedAt fields
    }
);

// Create and export the Like model
const Like = mongoose.model("Like", LikeSchema);

export default Like;