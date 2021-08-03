import React from "react";
import { useForm } from "react-hook-form";
import { Input, Textarea } from "../../components/Input";
import Layout from "../../components/Layout";

// Title, description, tags, checkpoints
// Title, instructions, links, numbered?
// TODO: Focus more on links?
function CreateMap() {
  const {
    register,
    handleSubmit,
    formState: { errors },
  } = useForm();

  // const [checkpoints, setCheckpoints] = useState([]);

  const onSubmit = () => {};

  return (
    <Layout title="Create | Roadmap">
      <main className="grid grid-cols-12">
        <section className="col-start-2 col-end-12 max-w-sm">
          <h1 className="text-3xl text-gray-800 font-medium tracking-wide mt-8 mb-3">
            Create roadmap
          </h1>

          <form onSubmit={handleSubmit(onSubmit)}>
            <Input
              id="title"
              name="Title"
              error={errors.title}
              register={register("title", {
                required: { value: true, message: "Title is required" },
              })}
            />
            <Textarea
              id="description"
              name="Description"
              error={errors.description}
              register={register("description", {
                required: { value: true, message: "Description is required" },
              })}
            />
          </form>
        </section>
      </main>
    </Layout>
  );
}

// function CreateCheckpoint() {
//   const [title, setTitle] = useState<string>("");
//   const [instructions, setInstructions] = useState<string>("");
//   const [links, setLinks] = useState<string[]>([]);

//   const handleAddLink = () => {};

//   return (
//     <div>
//       <input
//         type="text"
//         value={title}
//         onChange={(e) => setTitle(e.target.value)}
//       />
//       <textarea
//         rows={3}
//         value={instructions}
//         onChange={(e) => setInstructions(e.target.value)}
//       />
//       <div>
//         <button type="button" onClick={handleAddLink}>
//           Add link
//         </button>
//       </div>
//     </div>
//   );
// }

export default CreateMap;
