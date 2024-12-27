import User from "./model";

export async function save(data: UserPayload) {
    const user = new User(data);

    await user.save();
}