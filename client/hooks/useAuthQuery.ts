import { useEffect } from "react";
import { DocumentNode, useLazyQuery } from "@apollo/client";
import { getApolloClient } from "../lib/apollo-client";
import { useAuth0 } from "@auth0/auth0-react";
import { setContext } from "@apollo/client/link/context";

function useAuthQuery<TData = any, TVariables = any>(query: DocumentNode) {
  const [getData, { data, loading, error }] =
    useLazyQuery<TData, TVariables>(query);
  const { isLoading, getAccessTokenSilently } = useAuth0();

  useEffect(() => {
    if (!isLoading) {
      const client = getApolloClient();
      getAccessTokenSilently()
        .then((token) => {
          console.log("ACCESS TOKEN", token);
          console.log("CLIENT", client);

          const authLink = setContext((_, { headers }) => {
            return {
              headers: {
                ...headers,
                authorization: `Bearer ${token}`,
              },
            };
          });

          client.setLink(authLink.concat(client.link));

          getData();
        })
        .catch((e) => console.error(e));
    }
  }, [isLoading]);

  return { data, loading, error };
}

export default useAuthQuery;
