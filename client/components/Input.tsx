import React from "react";

type Props = {
  id: string;
  name: string;
  register: any;
  type?: string;
  error?: any;
};

function Input({ id, name, register, type = "text", error }: Props) {
  return (
    <div className="flex flex-col my-8">
      <label htmlFor={id} className="text-gray-800 font-medium tracking-wide">
        {name}
      </label>
      <input
        type={type}
        id={id}
        className={`border-b-2 pt-1 focus:border-gray-800 focus:outline-none ${
          error && "border-red-500"
        }`}
        {...register}
      />
      <span className="h-6 text-red-500 mt-1">{error?.message}</span>
    </div>
  );
}

function Textarea({ id, name, register, error }: Props) {
  return (
    <div className="flex flex-col my-8">
      <label htmlFor={id} className="text-gray-800 font-medium tracking-wide">
        {name}
      </label>
      <textarea
        id={id}
        rows={4}
        className={`border-b-2 pt-1 resize-none focus:border-gray-800 focus:outline-none ${
          error && "border-red-500"
        }`}
        {...register}
      />
      <span className="h-6 text-red-500 mt-1">{error?.message}</span>
    </div>
  );
}

export { Input, Textarea };
