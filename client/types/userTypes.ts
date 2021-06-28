import { RoadmapType } from "./roadmapTypes";

export type UserType = {
  id: number;
  username: string;
  email: string;
  password: string;
  followingRoadmaps: RoadmapType[];
  createdRoadmaps: RoadmapType[];
  createdAt: Date;
  updatedAt: Date;
};

export type CreatorType = {
  id: number;
  username: string;
};

export type FollowerType = {
  id: number;
};
