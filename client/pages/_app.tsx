import { AppProps } from "next/app";
import { ApolloProvider } from "@apollo/client";
import { useApollo } from "../lib/apollo-client";
import { Auth0Provider } from "@auth0/auth0-react";
import "../styles/globals.css";

function App({ Component, pageProps }: AppProps) {
  const client = useApollo(pageProps.initialApolloState);

  return (
    <ApolloProvider client={client}>
      <Auth0Provider
        domain={process.env.NEXT_PUBLIC_AUTH0_DOMAIN || ""}
        clientId={process.env.NEXT_PUBLIC_AUTH0_CLIENT_ID || ""}
        redirectUri="http://localhost:3000/private"
        audience={process.env.NEXT_PUBLIC_AUTH0_AUDIENCE || ""}
        scope="openid profile"
        // cacheLocation="localstorage"
        // useRefreshTokens={true}
        // scope={auth0Config.scope}
      >
        <Component {...pageProps} />
      </Auth0Provider>
    </ApolloProvider>
  );
}

export default App;
