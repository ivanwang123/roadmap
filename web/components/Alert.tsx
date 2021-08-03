import React from "react";

type Props = {
  message: string;
  error: boolean;
};

function Alert({ message, error }: Props) {
  if (message.length) {
    return (
      <div
        className={`w-full text-white px-6 py-2 rounded ${
          error ? "bg-red-400" : "bg-green-400"
        }`}
      >
        {message}
      </div>
    );
  } else {
    return <></>;
  }
}

export default Alert;
