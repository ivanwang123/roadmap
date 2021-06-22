import { Roadmap } from "./roadmapTypes";

export type Checkpoint = {
  id: number;
  title: string;
  instructions: string;
  links: string[];
  isCompleted: boolean;
  roadmap: Roadmap;
  createdAt: Date;
  updatedAt: Date;
};
