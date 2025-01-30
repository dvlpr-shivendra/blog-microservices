import Like from "./model";

export async function save(data: LikeData) {
  return Like.create(data);
}
