import Like from "./model";

export async function save(data: LikeData) {
    const like = new Like(data);

    await like.save();
}