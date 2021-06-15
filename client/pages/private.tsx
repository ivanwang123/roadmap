import React from "react";
import { useUser } from "@auth0/nextjs-auth0";

function Private() {
  const { user } = useUser();

  return (
    <div>
      <h1>Private</h1>
      <pre>{JSON.stringify(user, null, 2)}</pre>
      <a href="/api/auth/logout">Logout</a>
    </div>
  );
}

export default Private;
