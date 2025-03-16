import Like from "./model";
import { Repository } from "./repository.interface";
import { LikeData } from "./types";

export class MongoRepository implements Repository {
  constructor() {}

  async create(data: LikeData) {
    return Like.create(data);
  }
}
