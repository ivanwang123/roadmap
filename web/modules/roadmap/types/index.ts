export type RoadmapValues = {
  title: string;
  description: string;
  tags: string[];
  checkpoints: CheckpointType[];
};

export type CheckpointType = {
  title: string;
  instructions: string;
  links: string[];
};
