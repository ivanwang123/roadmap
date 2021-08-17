import { useMutation } from "@apollo/client";
import { getApolloClient } from "lib/apollo-client";
import { useNotificationStore } from "modules/notification";
import { LOGOUT_MUTATION } from "../api";

export function useLogout() {
  const { setNotification } = useNotificationStore();

  const [logout, { loading }] = useMutation(LOGOUT_MUTATION);

  const handleLogout = () => {
    logout()
      .then(() => {
        const client = getApolloClient();
        client.clearStore().then(() => {
          client.resetStore();
        });

        setNotification({
          type: "success",
          message: "Successfully logged out",
        });
      })
      .catch(() => {
        setNotification({
          type: "error",
          message: "Something went wrong, unable to logout",
        });
      });
  };

  return {
    logout: handleLogout,
    loading: loading,
  };
}
