import { Checkpoint } from "./checkpointTypes";
import { User } from "./userTypes";

export type Roadmap = {
  id: number;
  title: string;
  description: string;
  creator: User;
  checkpoints: Checkpoint;
  followers: User[];
  createdAt: Date;
  updatedAt: Date;
};
