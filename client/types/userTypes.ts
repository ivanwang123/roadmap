import { Roadmap } from "./roadmapTypes";

export type User = {
  id: number;
  username: string;
  email: string;
  password: string;
  followingRoadmaps: Roadmap[];
  createdRoadmaps: Roadmap[];
  createdAt: Date;
  updatedAt: Date;
};
