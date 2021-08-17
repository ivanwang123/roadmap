import { useMutation } from "@apollo/client";
import { getApolloClient } from "lib/apollo-client";
import { useNotificationStore } from "modules/notification";
import { useRouter } from "next/router";
import { ME_QUERY, REGISTER_MUTATION } from "../api";

export function useRegister() {
  const router = useRouter();

  const { setNotification } = useNotificationStore();

  const [register, { loading }] = useMutation(REGISTER_MUTATION);

  const handleRegister = (data: any) => {
    console.log("SUBMIT", data);
    register({
      variables: data,
      update: (cache, { data }) => {
        if (!data || !data.createUser) {
          return;
        }

        cache.writeQuery({
          query: ME_QUERY,
          data: {
            me: data.createUser,
          },
        });
      },
    })
      .then(() => {
        const client = getApolloClient();
        client.resetStore();

        setNotification({
          type: "success",
          message: "Successfully created a new account",
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
    register: handleRegister,
    loading,
  };
}
