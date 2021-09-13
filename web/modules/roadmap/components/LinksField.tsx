import { PrimaryButton } from "components/element";
import { InputField } from "components/form";
import React, { useState } from "react";
import { useFormContext, useWatch } from "react-hook-form";
import Add from "svgs/add.svg";
import Close from "svgs/close.svg";
import Link from "svgs/link.svg";
import { pluralize } from "utils";
import { RoadmapValues } from "../types";

type Props = {
  checkpointIdx: number;
};

export function LinksField({ checkpointIdx }: Props) {
  const {
    register,
    control,
    setValue,
    formState: { errors },
  } = useFormContext<RoadmapValues>();
  console.log("ERRORS", errors);

  // const { fields, append, remove } = useFieldArray({
  //   name: `checkpoints.${checkpointIdx}.links`,
  //   control,
  // });

  const links = useWatch({
    name: `checkpoints.${checkpointIdx}.links`,
    defaultValue: [],
    control,
  });
  console.log("LINKS", links, `checkpoints.${checkpointIdx}.links`);

  const [link, setLink] = useState<string>("");

  const handleKeyDown = (e: React.KeyboardEvent<HTMLInputElement>) => {
    if (e.key === "Enter") {
      e.preventDefault();
      appendLink();
    }
  };

  const appendLink = () => {
    if (link.length > 0) {
      setValue(`checkpoints.${checkpointIdx}.links`, [...links, link]);
      setLink("");
    }
  };

  const removeLink = (idx: number) => {
    setValue(
      `checkpoints.${checkpointIdx}.links`,
      links.filter((_, i) => i !== idx)
    );
  };

  return (
    <>
      <input
        className="hidden"
        {...register(`checkpoints.${checkpointIdx}.links`)}
      />
      <div className="flex">
        <InputField
          id="link"
          label="LINKS"
          placeholder="Add a new link"
          value={link}
          onKeyDown={handleKeyDown}
          onChange={(e) => setLink(e.target.value)}
        >
          <span className="flex justify-center items-center ml-2">
            <PrimaryButton onClick={appendLink} data-testid="add-link">
              <Add className="w-5 h-5 fill-current" />
            </PrimaryButton>
          </span>
        </InputField>
      </div>

      <div className="flex flex-col mb-7">
        <div className="flex items-center mb-1">
          <Link className="w-4 h-4 fill-current text-gray-500 mr-2" />
          {links.length === 0 ? (
            <p className="text-gray-500 italic">No links</p>
          ) : (
            <p className="text-gray-500 italic">
              {links.length} {pluralize("link", links.length)}
            </p>
          )}
        </div>

        {links.map((link, idx) => (
          <div
            className="flex items-center bg-blueGray-100 pr-2 mb-1 rounded"
            key={idx}
          >
            <button type="button" onClick={() => removeLink(idx)}>
              <Close className="w-4 h-4 fill-current text-blueGray-500 mx-1" />
            </button>
            <span className="truncate">{link}</span>
          </div>
        ))}
      </div>
    </>
  );
}
