export type CheckpointType = {
  id: number;
  title: string;
  instructions: string;
  links: LinkType[];
  status?: StatusType;
  createdAt: Date;
  updatedAt: Date;
};

export type StatusType = "complete" | "incomplete" | "skip" | "current";

export type LinkType = {
  url: string;
  title: string;
  description: string;
  image: string;
};
