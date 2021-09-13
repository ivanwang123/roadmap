import { PrimaryButton } from "components/element";
import { InputField } from "components/form";
import React, { useState } from "react";
import { useFormContext, useWatch } from "react-hook-form";
import Add from "svgs/add.svg";
import Close from "svgs/close.svg";
import Tag from "svgs/tag.svg";
import { RoadmapValues } from "../types";

export function TagsField() {
  const {
    control,
    setValue,
    formState: { errors },
  } = useFormContext<RoadmapValues>();
  console.log("ERRORS", errors);

  const tags = useWatch({
    name: "tags",
    defaultValue: [],
    control,
  });

  const [tag, setTag] = useState<string>("");

  const handleKeyDown = (e: React.KeyboardEvent<HTMLInputElement>) => {
    if (e.key === "Enter") {
      e.preventDefault();
      appendTag();
    }
  };

  const appendTag = () => {
    if (tag.length > 0) {
      setValue("tags", [...tags, tag]);
      setTag("");
    }
  };

  const removeTag = (idx: number) => {
    setValue(
      "tags",
      tags.filter((_, i) => i !== idx)
    );
  };

  return (
    <>
      <div className="flex">
        <InputField
          id="tag"
          label="TAGS"
          placeholder="Add a new tag"
          value={tag}
          onKeyDown={handleKeyDown}
          onChange={(e) => setTag(e.target.value)}
        >
          <span className="flex justify-center items-center ml-2">
            <PrimaryButton onClick={appendTag} data-testid="add-tag">
              <Add className="w-5 h-5 fill-current" />
            </PrimaryButton>
          </span>
        </InputField>
      </div>

      <div className="flex flex-wrap items-center mb-7">
        <Tag className="w-4 h-4 fill-current text-gray-500 mr-2 mb-1" />
        {tags.length === 0 && <p className="text-gray-500 italic">No tags</p>}

        {tags.map((tag, idx) => {
          return (
            <div
              className="flex items-center max-w-full bg-blueGray-100 pr-2 mr-2 mb-1 rounded"
              key={idx}
            >
              <button type="button" onClick={() => removeTag(idx)}>
                <Close className="w-4 h-4 fill-current text-blueGray-500 mx-1" />
              </button>
              <span className="truncate">{tag}</span>
            </div>
          );
        })}
      </div>
    </>
  );
}
