import React from "react";
import { useAuth0 } from "@auth0/auth0-react";
import { gql } from "@apollo/client";
import useAuthQuery from "../hooks/useAuthQuery";

const GET_USERS = gql`
  query {
    allUsers {
      id
      username
      # followingRoadmaps {
      #   id
      #   title
      #   description
      # }
      # createdRoadmaps {
      #   id
      #   title
      # }
    }
  }
`;

function Private() {
  const { logout, user, getAccessTokenSilently } = useAuth0();
  const { data } = useAuthQuery(GET_USERS);

  const getToken = async () => {
    try {
      const token = await getAccessTokenSilently({ ignoreCache: true });
      console.log("TOKEN", token);
    } catch (e) {
      console.error(e);
    }
  };

  return (
    <div>
      <h1>Private</h1>
      <pre>{JSON.stringify(user, null, 2)}</pre>
      {/* TODO: client.clearStore() when log out */}
      <button
        onClick={() => logout({ returnTo: "http://localhost:3000/login" })}
      >
        Log out
      </button>
      <button onClick={getToken}>Get token</button>
      <pre>{JSON.stringify(data, null, 2)}</pre>
    </div>
  );
}

export default Private;
