import { useMutation } from "@apollo/client";
import { Icon } from "components/element";
import React from "react";
import Add from "svgs/add.svg";
import Check from "svgs/check.svg";
import User from "svgs/user.svg";
import { UserInfoFieldsFragment } from "types/graphql-generated";
import { pluralize } from "utils";
import { TOGGLE_FOLLOW_ROADMAP_MUTATION } from "../api";

type FollowerType = {
  id: number;
};

type Props = {
  me?: UserInfoFieldsFragment;
  followers: FollowerType[];
  roadmapId: number;
};

export function FollowButton({ me, followers, roadmapId }: Props) {
  const [toggleFollow] = useMutation(TOGGLE_FOLLOW_ROADMAP_MUTATION);

  const handleToggleFollow = () => {
    if (me) {
      toggleFollow({
        variables: {
          roadmapId: roadmapId,
        },
      });
    }
  };

  if (me) {
    if (followers.some((follower) => follower.id === me.id)) {
      return (
        <UnFollow
          followers={followers}
          handleToggleFollow={handleToggleFollow}
        />
      );
    } else {
      return (
        <Follow followers={followers} handleToggleFollow={handleToggleFollow} />
      );
    }
  } else {
    return <Followers followers={followers} />;
  }
}

type FollowersProp = {
  followers: FollowerType[];
};

function Followers({ followers }: FollowersProp) {
  return (
    <div className="flex items-center">
      <Icon icon={User} size={12} />
      <span className="text-right font-semibold ml-2 mr-1">
        {followers.length}
      </span>
      <span>{pluralize("follower", followers.length)}</span>
    </div>
  );
}

type ToggleFollowProps = {
  handleToggleFollow: () => void;
} & FollowersProp;

function Follow({ handleToggleFollow, followers }: ToggleFollowProps) {
  return (
    <button
      type="button"
      className="flex items-center bg-blue-200 text-sm tracking-wide rounded focus:outline-none"
      onClick={handleToggleFollow}
    >
      <div className="flex items-center bg-blue-300 text-blue-500 px-2 py-1 rounded-l">
        <User className="fill-current mr-1" width={14} height={14} />
        <span className="font-semibold">{followers.length}</span>
      </div>
      <Add
        className="fill-current text-blue-400 ml-2 mr-1"
        width={16}
        height={16}
      />
      <div className="text-blue-400 pr-2 py-1 rounded-r">Follow</div>
    </button>
  );
}

function UnFollow({ handleToggleFollow, followers }: ToggleFollowProps) {
  return (
    <button
      type="button"
      className="flex items-center bg-secondary text-sm tracking-wide rounded-sm focus:outline-none"
      onClick={handleToggleFollow}
    >
      <div className="flex items-center bg-gray-200 text-gray-400 px-2 py-1 rounded-l">
        <User className="fill-current mr-1" width={14} height={14} />
        <span className="font-semibold">{followers.length}</span>
      </div>
      <Check
        className="fill-current text-gray-400 ml-2 mr-1"
        width={16}
        height={16}
      />
      <div className="text-gray-400 pr-2 py-1 rounded-r">Following</div>
    </button>
  );
}
