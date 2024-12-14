import * as store from "./store";

export function createLike(data: LikeData) {
    return store.save(data)
}