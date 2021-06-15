import Layout from "../components/Layout";
// import { gql } from "@apollo/client";
// import { getApolloClient } from "../apollo-client";
// import { InferGetStaticPropsType } from "next";

// { users }: InferGetStaticPropsType<typeof getStaticProps>

function Home() {
  return (
    <Layout title="Home | Roadmap">
      <h1 className="font-bold text-emerald-500">Roadmap</h1>
      <div>
        {/* {users.map((user: any) => (
          <div key={user.id}>
            <h3>{user.id}</h3>
            <p>
              {user.username} - {user.email}
            </p>
          </div>
        ))} */}
      </div>
    </Layout>
  );
}

// export const getStaticProps = async () => {
//   const client = getApolloClient();

//   const { data } = await client.query({
//     query: gql`
//       query Users {
//         users {
//           id
//           username
//           email
//         }
//       }
//     `,
//   });

//   return {
//     props: {
//       users: data.users.slice(0, 4),
//     },
//   };
// };

export default Home;
