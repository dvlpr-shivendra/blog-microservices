import * as store from "./store";

export function createUser(data: UserPayload) {
    return store.save(data)
}