import { PrimaryButton, SecondaryButton } from "components/element";
import { InputField, TextareaField } from "components/form";
import React from "react";
import { useFieldArray, useFormContext } from "react-hook-form";
import Add from "svgs/add.svg";
import { RoadmapValues } from "../types";
import { LinksField } from "./LinksField";

export function CheckpointsField() {
  const { control } = useFormContext<RoadmapValues>();

  const { fields, append, remove } = useFieldArray({
    name: "checkpoints",
    control,
  });

  return (
    <>
      <p className="text-gray-500 text-sm font-medium tracking-wide mb-1">
        CHECKPOINTS
      </p>
      <PrimaryButton
        label="Checkpoint"
        onClick={() =>
          append({
            title: "",
            instructions: "",
            links: [],
          })
        }
      >
        <Add className="w-5 h-5 fill-current mr-2" />
      </PrimaryButton>
      {fields.map((field, idx) => (
        <div className="flex flex-col" key={field.id}>
          <InputField
            id={`checkpoints.${idx}.title`}
            label="TITLE"
            defaultValue={field.title}
          />
          <TextareaField
            id={`checkpoints.${idx}.instructions`}
            label="INSTRUCTIONS"
            defaultValue={field.instructions}
          />
          <LinksField checkpointIdx={idx} />
          <SecondaryButton label="Delete" onClick={() => remove(idx)} />
        </div>
      ))}
    </>
  );
}
