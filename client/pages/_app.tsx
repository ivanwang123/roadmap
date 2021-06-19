import React from "react";
import { ApolloProvider } from "@apollo/client";
import { AppProps } from "next/app";
import { useApollo } from "../lib/apollo-client";
import "../styles/tailwind.css";
import "../styles/globals.css";

function App({ Component, pageProps }: AppProps) {
  const client = useApollo(pageProps.initialApolloState);

  return (
    <ApolloProvider client={client}>
      <Component {...pageProps} />
    </ApolloProvider>
  );
}

export default App;
