import { Form, InputField, SubmitButton, TextareaField } from "components/form";
import { Layout } from "components/layout";
import { Notification } from "modules/notification";
import { CheckpointsField, RoadmapValues, TagsField } from "modules/roadmap";
import React from "react";
import { z } from "zod";

const schema = z.object({
  title: z
    .string()
    .min(1, "Title is required")
    .max(256, "Title must be less than 256 characters long"),
  description: z.string().min(1, "Description is required"),
  tags: z.array(z.string().min(1, "Tag can not be empty")).default([]),
  checkpoints: z.array(
    z.object({
      title: z
        .string()
        .min(1, "Title is required")
        .max(256, "Title must be between 1-256 characters long"),
      instructions: z.string(),
      links: z.array(z.string().min(1, "Link can not be empty")),
    })
  ),
});

function CreateMap() {
  const onSubmit = (data: RoadmapValues) => {
    console.log("DATA", data);
  };

  return (
    <Layout title="Create | Roadmap">
      <main className="flex justify-center w-full">
        <section className="w-full max-w-sm pb-12">
          <h1 className="text-3xl text-black font-semibold mt-8 mb-3">
            Create roadmap
          </h1>

          <Form<RoadmapValues, typeof schema>
            onSubmit={onSubmit}
            schema={schema}
          >
            <>
              <InputField label="TITLE" id="title" />
              <TextareaField label="DESCRIPTION" id="description" />
              <TagsField />
              <CheckpointsField />
              <SubmitButton label="Create" loading={false} />
            </>
          </Form>
          <Notification type="text" style="mt-1" showOnly="error" />
        </section>
      </main>
    </Layout>
  );
}

export default CreateMap;
