import { CheckpointType } from "./checkpointTypes";
import { CreatorType, FollowerType } from "./userTypes";

export type RoadmapType = {
  id: number;
  title: string;
  description: string;
  creator: CreatorType;
  checkpoints: CheckpointType[];
  followers: FollowerType[];
  createdAt: Date;
  updatedAt: Date;
};
