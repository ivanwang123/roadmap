import { gql } from '@apollo/client';
export type Maybe<T> = T | null;
export type Exact<T extends { [key: string]: unknown }> = { [K in keyof T]: T[K] };
export type MakeOptional<T, K extends keyof T> = Omit<T, K> & { [SubKey in K]?: Maybe<T[SubKey]> };
export type MakeMaybe<T, K extends keyof T> = Omit<T, K> & { [SubKey in K]: Maybe<T[SubKey]> };
/** All built-in and custom scalars, mapped to their actual values */
export type Scalars = {
  ID: number;
  String: string;
  Boolean: boolean;
  Int: number;
  Float: number;
  Time: Date;
};


export type Checkpoint = {
  id: Scalars['ID'];
  title: Scalars['String'];
  instructions: Scalars['String'];
  links: Array<Link>;
  status?: Maybe<Status>;
  roadmap: Roadmap;
  createdAt: Scalars['Time'];
  updatedAt: Scalars['Time'];
};

export type CheckpointStatus = {
  userId: Scalars['Int'];
  checkpointId: Scalars['Int'];
  roadmapId: Scalars['Int'];
  status: Status;
};

export type FollowRoadmap = {
  roadmapId: Scalars['Int'];
};

export type GetRoadmap = {
  id: Scalars['Int'];
};

export type GetRoadmaps = {
  cursorId: Scalars['Int'];
  cursorValue: Scalars['String'];
  sort: Sort;
};

export type GetUser = {
  id: Scalars['Int'];
};

export type Link = {
  url: Scalars['String'];
  title: Scalars['String'];
  description: Scalars['String'];
  image: Scalars['String'];
};

export type Login = {
  username?: Maybe<Scalars['String']>;
  email?: Maybe<Scalars['String']>;
  password: Scalars['String'];
};

export type Mutation = {
  login: User;
  logout: Scalars['Boolean'];
  createUser: User;
  createCheckpoint: Checkpoint;
  createRoadmap: Roadmap;
  toggleFollowRoadmap: Roadmap;
  updateCheckpointStatus: Checkpoint;
};


export type MutationLoginArgs = {
  input: Login;
};


export type MutationCreateUserArgs = {
  input: NewUser;
};


export type MutationCreateCheckpointArgs = {
  input: NewCheckpoint;
};


export type MutationCreateRoadmapArgs = {
  input: NewRoadmap;
};


export type MutationToggleFollowRoadmapArgs = {
  input: FollowRoadmap;
};


export type MutationUpdateCheckpointStatusArgs = {
  input: UpdateStatus;
};

export type NewCheckpoint = {
  title: Scalars['String'];
  instructions: Scalars['String'];
  links: Array<Scalars['String']>;
  roadmapId: Scalars['Int'];
};

export type NewRoadmap = {
  title: Scalars['String'];
  description: Scalars['String'];
  creatorId: Scalars['Int'];
};

export type NewUser = {
  username: Scalars['String'];
  email: Scalars['String'];
  password: Scalars['String'];
};

export type Query = {
  user: User;
  roadmap: Roadmap;
  users: Array<User>;
  roadmaps: Array<Roadmap>;
  me?: Maybe<User>;
};


export type QueryUserArgs = {
  input?: Maybe<GetUser>;
};


export type QueryRoadmapArgs = {
  input?: Maybe<GetRoadmap>;
};


export type QueryRoadmapsArgs = {
  input?: Maybe<GetRoadmaps>;
};

export type Roadmap = {
  id: Scalars['ID'];
  title: Scalars['String'];
  description: Scalars['String'];
  creator: User;
  checkpoints: Array<Checkpoint>;
  followers: Array<User>;
  createdAt: Scalars['Time'];
  updatedAt: Scalars['Time'];
};

export type RoadmapFollower = {
  user: User;
  roadmap: Roadmap;
};

export enum Sort {
  Newest = 'NEWEST',
  Oldest = 'OLDEST',
  MostFollowers = 'MOST_FOLLOWERS',
  MostCheckpoints = 'MOST_CHECKPOINTS',
  LeastCheckpoints = 'LEAST_CHECKPOINTS'
}

export enum Status {
  Complete = 'COMPLETE',
  Incomplete = 'INCOMPLETE',
  Skip = 'SKIP'
}


export type UpdateStatus = {
  checkpointId: Scalars['Int'];
  status: Status;
};

export type User = {
  id: Scalars['ID'];
  username: Scalars['String'];
  email: Scalars['String'];
  password: Scalars['String'];
  followingRoadmaps: Array<Roadmap>;
  createdRoadmaps: Array<Roadmap>;
  createdAt: Scalars['Time'];
  updatedAt: Scalars['Time'];
};

export type LinkFieldsFragment = { url: string, title: string, description: string, image: string };

export type RoadmapInfoFieldsFragment = { id: number, title: string, description: string, createdAt: Date, updatedAt: Date, creator: { id: number, username: string }, followers: Array<{ id: number }>, checkpoints: Array<{ id: number }> };

export type UserInfoFieldsFragment = { id: number, username: string, email: string, createdAt: Date, updatedAt: Date };

export type LoginMutationVariables = Exact<{
  email: Scalars['String'];
  password: Scalars['String'];
}>;


export type LoginMutation = { login: UserInfoFieldsFragment };

export type LogoutMutationVariables = Exact<{ [key: string]: never; }>;


export type LogoutMutation = { logout: boolean };

export type RegisterMutationVariables = Exact<{
  email: Scalars['String'];
  username: Scalars['String'];
  password: Scalars['String'];
}>;


export type RegisterMutation = { createUser: UserInfoFieldsFragment };

export type ToggleFollowRoadmapMutationVariables = Exact<{
  roadmapId: Scalars['Int'];
}>;


export type ToggleFollowRoadmapMutation = { toggleFollowRoadmap: { id: number, followers: Array<{ id: number }> } };

export type UpdateCheckpointStatusMutationVariables = Exact<{
  checkpointId: Scalars['Int'];
  status: Status;
}>;


export type UpdateCheckpointStatusMutation = { updateCheckpointStatus: { id: number, status?: Maybe<Status> } };

export type MeQueryVariables = Exact<{ [key: string]: never; }>;


export type MeQuery = { me?: Maybe<UserInfoFieldsFragment> };

export type RoadmapQueryVariables = Exact<{
  id: Scalars['Int'];
}>;


export type RoadmapQuery = { roadmap: (
    { checkpoints: Array<{ id: number, title: string, instructions: string, status?: Maybe<Status>, links: Array<LinkFieldsFragment> }> }
    & RoadmapInfoFieldsFragment
  ) };

export type RoadmapsQueryVariables = Exact<{
  cursorId: Scalars['Int'];
  cursorValue: Scalars['String'];
  sort: Sort;
}>;


export type RoadmapsQuery = { roadmaps: Array<RoadmapInfoFieldsFragment> };

export type UserQueryVariables = Exact<{
  id: Scalars['Int'];
}>;


export type UserQuery = { user: (
    { followingRoadmaps: Array<RoadmapInfoFieldsFragment>, createdRoadmaps: Array<RoadmapInfoFieldsFragment> }
    & UserInfoFieldsFragment
  ) };

export const LinkFieldsFragmentDoc = gql`
    fragment LinkFields on Link {
  url
  title
  description
  image
}
    `;
export const RoadmapInfoFieldsFragmentDoc = gql`
    fragment RoadmapInfoFields on Roadmap {
  id
  title
  description
  createdAt
  updatedAt
  creator {
    id
    username
  }
  followers {
    id
  }
  checkpoints {
    id
  }
}
    `;
export const UserInfoFieldsFragmentDoc = gql`
    fragment UserInfoFields on User {
  id
  username
  email
  createdAt
  updatedAt
}
    `;
export const LoginDocument = gql`
    mutation Login($email: String!, $password: String!) {
  login(input: {email: $email, password: $password}) {
    ...UserInfoFields
  }
}
    ${UserInfoFieldsFragmentDoc}`;
export const LogoutDocument = gql`
    mutation Logout {
  logout
}
    `;
export const RegisterDocument = gql`
    mutation Register($email: String!, $username: String!, $password: String!) {
  createUser(input: {email: $email, username: $username, password: $password}) {
    ...UserInfoFields
  }
}
    ${UserInfoFieldsFragmentDoc}`;
export const ToggleFollowRoadmapDocument = gql`
    mutation ToggleFollowRoadmap($roadmapId: Int!) {
  toggleFollowRoadmap(input: {roadmapId: $roadmapId}) {
    id
    followers {
      id
    }
  }
}
    `;
export const UpdateCheckpointStatusDocument = gql`
    mutation UpdateCheckpointStatus($checkpointId: Int!, $status: Status!) {
  updateCheckpointStatus(input: {checkpointId: $checkpointId, status: $status}) {
    id
    status
  }
}
    `;
export const MeDocument = gql`
    query Me {
  me {
    ...UserInfoFields
  }
}
    ${UserInfoFieldsFragmentDoc}`;
export const RoadmapDocument = gql`
    query Roadmap($id: Int!) {
  roadmap(input: {id: $id}) {
    ...RoadmapInfoFields
    checkpoints {
      id
      title
      instructions
      status
      links {
        ...LinkFields
      }
    }
  }
}
    ${RoadmapInfoFieldsFragmentDoc}
${LinkFieldsFragmentDoc}`;
export const RoadmapsDocument = gql`
    query Roadmaps($cursorId: Int!, $cursorValue: String!, $sort: Sort!) {
  roadmaps(input: {cursorId: $cursorId, cursorValue: $cursorValue, sort: $sort}) {
    ...RoadmapInfoFields
  }
}
    ${RoadmapInfoFieldsFragmentDoc}`;
export const UserDocument = gql`
    query User($id: Int!) {
  user(input: {id: $id}) {
    ...UserInfoFields
    followingRoadmaps {
      ...RoadmapInfoFields
    }
    createdRoadmaps {
      ...RoadmapInfoFields
    }
  }
}
    ${UserInfoFieldsFragmentDoc}
${RoadmapInfoFieldsFragmentDoc}`;

      export interface PossibleTypesResultData {
        possibleTypes: {
          [key: string]: string[]
        }
      }
      const result: PossibleTypesResultData = {
  "possibleTypes": {}
};
      export default result;
    