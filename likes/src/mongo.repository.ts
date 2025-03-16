import { Model } from "mongoose";
import { LikeDocument } from "./model";
import { Repository } from "./repository.interface";
import { LikeData } from "./types";

export class MongoRepository implements Repository {
  constructor(private model: Model<LikeDocument>) {}

  async create(data: LikeData): Promise<LikeDocument> {
    return this.model.create(data);
  }
}
