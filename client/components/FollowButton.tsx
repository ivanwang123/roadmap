import { useMutation } from "@apollo/client";
import React from "react";
import { UserInfoFieldsFragment } from "../graphql/generated/generated";
import { TOGGLE_FOLLOW_ROADMAP_MUTATION } from "../graphql/mutations/toggleFollowRoadmap";
import Add from "../svgs/add.svg";
import Check from "../svgs/check.svg";
import User from "../svgs/user.svg";
import Icon from "./Icon";

type FollowerType = {
  id: number;
};

type Props = {
  me?: UserInfoFieldsFragment;
  followers: FollowerType[];
  roadmapId: number;
};

function FollowButton({ me, followers, roadmapId }: Props) {
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

  let followButton = null;
  if (me) {
    if (followers.some((follower) => follower.id === me.id)) {
      followButton = (
        <button
          type="button"
          className="flex items-center bg-secondary text-sm tracking-wide rounded focus:outline-none"
          onClick={handleToggleFollow}
        >
          <div className="flex items-center bg-blueGray-200 text-gray-400 px-2 py-1 rounded-l">
            <User className="fill-current mr-1" width={14} height={14} />
            <span className="font-medium">{followers.length}</span>
          </div>
          <Check
            className="fill-current text-gray-400 mx-1"
            width={16}
            height={16}
          />
          <div className="text-gray-400 pr-2 py-1 rounded-r">Following</div>
        </button>
      );
    } else {
      followButton = (
        <button
          type="button"
          className="flex items-center bg-blue-200 text-sm tracking-wide rounded focus:outline-none"
          onClick={handleToggleFollow}
        >
          <div className="flex items-center bg-blue-300 text-blue-500 px-2 py-1 rounded-l">
            <User className="fill-current mr-1" width={14} height={14} />
            <span className="font-medium">{followers.length}</span>
          </div>
          <Add
            className="fill-current text-blue-400 mx-1"
            width={16}
            height={16}
          />
          <div className="text-blue-400 pr-2 py-1 rounded-r">Follow</div>
        </button>
      );
    }
  } else {
    followButton = (
      <div className="flex items-center">
        <Icon icon={User} size={12} />
        <span className="text-right font-medium ml-2 mr-1">
          {followers.length}
        </span>
        <span>followers</span>
      </div>
    );
  }

  return followButton;
}

export default FollowButton;