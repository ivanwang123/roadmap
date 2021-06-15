import React, { useEffect } from "react";
import { useAuth0 } from "@auth0/auth0-react";

function Public() {
  const { user, getAccessTokenSilently } = useAuth0();

  useEffect(() => {
    (async () => {
      try {
        const token = await getAccessTokenSilently();
        // const res = await axios.post(
        //   "https://dev-jkn4emz6.us.auth0.com/oauth/token",
        //   {
        //     client_id: "2El7qcsqxaGzhb6ys9SME8Ofxdcvst34",
        //     client_secret:
        //       "bdT8rWiG4OMRilWdAH9POZBnR-Kn7URJ1HG5T1VYYxdhoWmUOurxMZFGNDtEfi9m",
        //     audience: "https://dev-jkn4emz6.us.auth0.com/api/v2/",
        //     grant_type: "client_credentials",
        //   }
        // );
        console.log("ACCESS TOKEN", token);
      } catch (e) {
        console.error("ACCESS TOKEN ERROR", e);
      }
    })();
  }, []);
  return (
    <div>
      <h1>Public</h1>
      <pre>{JSON.stringify(user, null, 2)}</pre>
    </div>
  );
}

export default Public;
