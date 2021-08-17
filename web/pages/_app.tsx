import { ApolloProvider } from "@apollo/client";
import { useApollo } from "lib/apollo-client";
import { AppProps } from "next/app";
import React from "react";
import "styles/globals.css";
import "styles/tailwind.css";

function App({ Component, pageProps }: AppProps) {
  const client = useApollo(pageProps);

  return (
    <ApolloProvider client={client}>
      <Component {...pageProps} />
    </ApolloProvider>
  );
}

export default App;
