import * as store from "./store";

import jwt from "jsonwebtoken";

function generateToken(payload: object) {
    return jwt.sign(payload, process.env.JWT_SECRET as string, {
        expiresIn: "24h",
    });
}

export async function createUser(data: UserPayload) {
    const user = await store.save(data);
    return generateToken({ id: user._id, name: user.name });
}