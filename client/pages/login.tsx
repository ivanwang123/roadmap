import React from "react";
import { useAuth0 } from "@auth0/auth0-react";

function Login() {
  const { loginWithRedirect } = useAuth0();

  return (
    <div>
      <h1>Login</h1>
      <div>
        <button onClick={() => loginWithRedirect()}>Log in</button>
      </div>
    </div>
  );
}

export default Login;
