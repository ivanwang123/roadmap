import { Icon } from "components/element";
import { Checkpoints } from "modules/checkpoint";
import React from "react";
import Book from "svgs/book.svg";
import Flag from "svgs/flag-big.svg";
import { RoadmapQuery, UserInfoFieldsFragment } from "types/graphql-generated";
import { pluralize } from "utils";
import { FollowButton } from "./FollowButton";

type Props = {
  roadmap: RoadmapQuery["roadmap"];
  me?: UserInfoFieldsFragment;
};

export function Sidebar({ roadmap, me }: Props) {
  return (
    <section className="w-64 h-full">
      {/* TODO: Remove border? */}
      <div className="sticky top-0 flex flex-col h-full border-r- border-secondary pl-10 pt-10 overflow-auto">
        <div className="flex text-gray-400 tracking-wide mb-2">
          <div className="flex items-center mr-10">
            <FollowButton
              me={me}
              followers={roadmap.followers}
              roadmapId={roadmap.id}
            />
          </div>
        </div>
        <div className="text-gray-400 tracking-wide mb-8">
          <div className="flex items-center mb-2">
            <Icon icon={Flag} size={12} />
            <span className="text-right font-semibold ml-2 mr-1">
              {roadmap.checkpoints.length}
            </span>
            <span>{pluralize("checkpoint", roadmap.checkpoints.length)}</span>
          </div>
          <div className="flex items-center">
            <Icon icon={Book} size={12} />
            {/* TODO: Get number of resources */}
            <span className="text-right font-semibold ml-2 mr-1">12</span>
            <span>{pluralize("resource", 12)}</span>
          </div>
        </div>

        <Checkpoints checkpoints={roadmap.checkpoints} />
      </div>
    </section>
  );
}
