import { useMemo } from "react";
import {
  ApolloClient,
  HttpLink,
  InMemoryCache,
  NormalizedCacheObject,
} from "@apollo/client";
// import { relayStylePagination } from "@apollo/client/utilities";
import merge from "deepmerge";
import { isEqual } from "lodash";

export const APOLLO_STATE_PROP_NAME = "__APOLLO_STATE__";

let apolloClient: ApolloClient<NormalizedCacheObject> | undefined;

function createApolloClient() {
  return new ApolloClient({
    ssrMode: typeof window === "undefined",
    link: new HttpLink({
      uri: "http://localhost:8080/query", // Server URL (must be absolute)
      credentials: "include", // Additional fetch() options like `credentials` or `headers`
    }),
    cache: new InMemoryCache({
      typePolicies: {
        Query: {
          fields: {
            roadmaps: {
              keyArgs: ["type", "input", ["sort"]],
              merge(existing = [], incoming, { args: { input } }: any) {
                const merged = existing.slice(0);

                let offset = offsetFromCursor(merged, input.cursorId);

                for (let i = 0; i < incoming.length; i++) {
                  merged[offset + i] = incoming[i];
                }

                return merged;
              },
            },
          },
        },
      },
    }),
  });
}

// TODO: Change depending on cursor value
function offsetFromCursor(items: any[], cursor: number) {
  for (let i = items.length - 1; i >= 0; i--) {
    if (items[i].id === cursor) {
      return i + 1;
    }
  }

  return items.length;
}

function initializeApollo(initialState: NormalizedCacheObject | null = null) {
  const _apolloClient = apolloClient ?? createApolloClient();

  // If your page has Next.js data fetching methods that use Apollo Client, the initial state
  // gets hydrated here
  if (initialState) {
    // Get existing cache, loaded during client side data fetching
    const existingCache = _apolloClient.extract();

    // Merge the existing cache into data passed from getStaticProps/getServerSideProps
    const data = merge(initialState, existingCache, {
      // combine arrays using object equality (like in sets)
      arrayMerge: (destinationArray, sourceArray) => [
        ...sourceArray,
        ...destinationArray.filter((d) =>
          sourceArray.every((s) => !isEqual(d, s))
        ),
      ],
    });

    // Restore the cache with the merged data
    _apolloClient.cache.restore(data);
  }

  // For SSG and SSR always create a new Apollo Client
  if (typeof window === "undefined") return _apolloClient;

  // Create the Apollo Client once in the client
  if (!apolloClient) apolloClient = _apolloClient;

  return _apolloClient;
}

export const getApolloClient = initializeApollo;

export function addApolloState(
  client: ApolloClient<NormalizedCacheObject>,
  pageProps: any
) {
  if (pageProps?.props) {
    pageProps.props[APOLLO_STATE_PROP_NAME] = client.cache.extract();
  }

  return pageProps;
}

export function useApollo(pageProps: any) {
  const state = pageProps[APOLLO_STATE_PROP_NAME];
  const store = useMemo(() => initializeApollo(state), [state]);
  return store;
}

// TODO: Deprecated

// import {
//   ApolloClient,
//   HttpLink,
//   InMemoryCache,
//   NormalizedCacheObject,
// } from "@apollo/client";
// import { useMemo } from "react";

// let apolloClient: ApolloClient<NormalizedCacheObject> | undefined;

// function createApolloClient(): ApolloClient<any> {
//   console.log("CREATE APOLLO CLIENT");
//   return new ApolloClient({
//     /**
//      * Enable SSR mode when not running on the client-side
//      */
//     ssrMode: typeof window === "undefined",
//     link: new HttpLink({
//       uri: "http://localhost:8080/query",
//       credentials: "include",
//     }),
//     cache: new InMemoryCache(),
//     connectToDevTools: true,
//   });
// }

// export function initializeApollo(initialState: any = null): ApolloClient<any> {
//   const _apolloClient = apolloClient ?? createApolloClient();

//   // If your page has Next.js data fetching methods that use Apollo Client, the initial state
//   // get hydrated here
//   if (initialState) {
//     _apolloClient.cache.restore(initialState);
//   }

//   /**
//    * SSG and SSR
//    * Always create a new Apollo Client
//    */
//   if (typeof window === "undefined") {
//     return _apolloClient;
//   }

//   // Create the Apollo Client once in the client
//   apolloClient = apolloClient ?? _apolloClient;

//   return apolloClient;
// }

// export const getApolloClient = initializeApollo;

// export function useApollo(initialState: any) {
//   const apolloStore = useMemo(
//     () => initializeApollo(initialState),
//     [initialState]
//   );
//   return apolloStore;
// }
