import Like from "./model";
import { LikeData } from "./types";

export async function save(data: LikeData) {
  return Like.create(data);
}
