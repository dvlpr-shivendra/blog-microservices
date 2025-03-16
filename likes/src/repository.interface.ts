import { LikeData } from "./types";

export interface Repository {
  create(data: LikeData): Promise<any>;
}
