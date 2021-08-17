import React from "react";
import Loader from "svgs/loader.svg";

type Props = {
  label: string;
  loading: boolean;
};

export function SubmitButton({ label, loading }: Props) {
  return (
    <button
      type="submit"
      className="icon-btn-grid items-center bg-green-500 text-white font-medium tracking-wide py-2 rounded disabled:opacity-70"
      disabled={loading}
    >
      <span className="justify-self-end">
        {loading && (
          <Loader
            className="fill-current animate-spin mr-2"
            width={20}
            height={20}
          />
        )}
      </span>
      <span>{label}</span>
    </button>
  );
}
