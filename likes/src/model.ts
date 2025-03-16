import mongoose, { Document, Schema, Model } from "mongoose";

// Define the Like interface for TypeScript
export interface LikeDocument extends Document {
  userId: number;
  postId: number;
  createdAt: Date;
  updatedAt: Date;
}

// Define the Like schema
const LikeSchema = new Schema<LikeDocument>(
  {
    userId: { type: Number, required: true },
    postId: { type: Number, required: true },
  },
  { timestamps: true } // Adds createdAt and updatedAt fields
);

// Create and export the Like model
const Like: Model<LikeDocument> = mongoose.model<LikeDocument>(
  "Like",
  LikeSchema
);

export default Like;
