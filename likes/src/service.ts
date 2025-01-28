import * as store from "./store";

export class LikeService {
  createLike(data: LikeData) {
    return store.save(data);
  }
}
