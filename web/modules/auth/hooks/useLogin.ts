import { useMutation } from "@apollo/client";
import { useNotificationStore } from "modules/notification";
import { getApolloClient } from "lib/apollo-client";
import { useRouter } from "next/router";
import { LOGIN_MUTATION, ME_QUERY } from "../api";

export function useLogin() {
  const router = useRouter();

  const { setNotification } = useNotificationStore();

  const [login, { loading }] = useMutation(LOGIN_MUTATION);

  const handleLogin = (data: any) => {
    console.log("SUBMIT", data);
    login({
      variables: data,
      update: (cache, { data }) => {
        console.log("LOGIN DATA", data);
        if (!data || !data.login) {
          return;
        }

        cache.writeQuery({
          query: ME_QUERY,
          data: {
            me: data.login,
          },
        });
      },
    })
      .then(() => {
        const client = getApolloClient();
        client.resetStore();

        setNotification({
          type: "success",
          message: "Successfully logged in",
        });

        if (router.query.redirect !== undefined) {
          router.push(router.query.redirect as string);
        } else {
          router.push("/");
        }
      })
      .catch((err) => {
        setNotification({
          type: "error",
          message: err.message,
        });
      });
  };

  return {
    login: handleLogin,
    loading,
  };
}
