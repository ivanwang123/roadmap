import React, { useEffect } from "react";
import { useAuth0 } from "@auth0/auth0-react";
import { useRouter } from "next/router";
import { Auth0Client } from "@auth0/auth0-spa-js";

const auth0 = new Auth0Client({
  domain: process.env.NEXT_PUBLIC_AUTH0_DOMAIN || "",
  client_id: process.env.NEXT_PUBLIC_AUTH0_CLIENT_ID || "",
  audience: process.env.NEXT_PUBLIC_AUTH0_AUDIENCE || "",
  redirect_uri: "http://localhost:3000/private",
});

function Login() {
  const { loginWithRedirect } = useAuth0();
  const router = useRouter();

  const login = async () => {
    await auth0
      .loginWithRedirect({
        redirect_uri: "http://localhost:3000/login",
      })
      .then(() => {
        //logged in. you can get the user profile like this:
        auth0.getUser().then((user) => {
          console.log(user);
        });
      });
    //logged in. you can get the user profile like this:
  };

  useEffect(() => {
    const getSession = async () => {
      const session = await auth0.checkSession();
      const user = await auth0.getUser();
      console.log("USER", user);
      console.log("SESSION", session);
    };
    getSession();
  }, []);

  return (
    <div>
      <h1>Login</h1>
      <div>
        <button onClick={() => loginWithRedirect()}>Log in</button>
        <a
          href={`https://dev-jkn4emz6.us.auth0.com/authorize?response_type=id_token token&client_id=2El7qcsqxaGzhb6ys9SME8Ofxdcvst34&redirect_uri=${
            "http://localhost:3000" + router.pathname
          }&scope=openid profile&state=xyzABC123&prompt=none&nonce=NONCE&audience=https://dev-jkn4emz6.us.auth0.com/api/v2/`}
        >
          Login link
        </a>
        <button onClick={login}>Login SPA</button>
      </div>
    </div>
  );
}

export default Login;
