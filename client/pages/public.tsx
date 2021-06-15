import React from "react";
import { useUser } from "@auth0/nextjs-auth0";

function Public() {
  const { user } = useUser();

  return (
    <div>
      <h1>Public</h1>
      <pre>{JSON.stringify(user, null, 2)}</pre>
    </div>
  );
}

export default Public;
