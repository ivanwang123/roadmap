import { AppProps } from "next/app";
import { ApolloProvider } from "@apollo/client";
import { useApollo } from "../lib/apollo-client";
import { Auth0Provider } from "@auth0/auth0-react";
// import auth0Config from "../config/auth0.json";
import "../styles/globals.css";

function App({ Component, pageProps }: AppProps) {
  const client = useApollo(pageProps.initialApolloState);

  return (
    <ApolloProvider client={client}>
      <Auth0Provider
        domain="dev-jkn4emz6.us.auth0.com"
        clientId="2El7qcsqxaGzhb6ys9SME8Ofxdcvst34"
        redirectUri="https://roadmapper.vercel.app/private"
        audience="https://dev-jkn4emz6.us.auth0.com/api/v2/"
        // cacheLocation="localstorage"
        useRefreshTokens={true}
        // scope={auth0Config.scope}
      >
        <Component {...pageProps} />
      </Auth0Provider>
    </ApolloProvider>
  );
}

export default App;
