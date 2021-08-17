import { Layout } from "components/layout";
import { withAuth, WithAuthProps } from "modules/auth";
import Router from "next/router";
import React from "react";
import { Loading, Redirecting } from "../components/placeholder";

type InputProps = {
  roadmaps: string[];
};

function Private({
  data: { me, loading, error },
  roadmaps,
}: WithAuthProps<InputProps>) {
  console.log("ROADMAPS", roadmaps);
  if (loading) return <Loading />;
  if (error) {
    Router.push({ pathname: "/login", query: { redirect: Router.pathname } });
    return <Redirecting />;
  }
  return (
    <Layout title="Private | Roadmap">
      <h1>Private</h1>
      <pre>{JSON.stringify(me, null, 2)}</pre>
    </Layout>
  );
}

export default withAuth<InputProps>()(Private);

export const getServerSideProps = async () => {
  const props: InputProps = {
    roadmaps: ["hi again"],
  };

  return {
    props,
  };
};

// const GET_USERS = gql`
//   query {
//     allUsers {
//       id
//       username
//       # followingRoadmaps {
//       #   id
//       #   title
//       #   description
//       # }
//       # createdRoadmaps {
//       #   id
//       #   title
//       # }
//     }
//   }
// `;
// { me }: InferGetStaticPropsType<typeof getServerSideProps>

// export const getServerSideProps: GetServerSideProps = async () => {
//   const client = getApolloClient();
//   const { data, error } = await client.query({
//     query: gql`
//       query Me {
//         me {
//           id
//           username
//           email
//         }
//       }
//     `,
//   });

//   if (!data.me || error) {
//     return {
//       redirect: {
//         destination: "/login",
//         permanent: false,
//       },
//     };
//   }

//   return {
//     props: {
//       me: data.me,
//     },
//   };
// };
