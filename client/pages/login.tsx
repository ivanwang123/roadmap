import { useMutation } from "@apollo/client";
import Link from "next/link";
import Router, { useRouter } from "next/router";
import React from "react";
import { useForm } from "react-hook-form";
import { Input } from "../components/Input";
import Layout from "../components/Layout";
import { LOGIN_MUTATION } from "../graphql/mutations/login";
import { ME_QUERY } from "../graphql/queries/me";
import { getApolloClient } from "../lib/apollo-client";
import Loader from "../svgs/loader.svg";

function Login() {
  // const [error, setError] = useState<boolean>(false);
  // const [message, setMessage] = useState<string>("");

  const {
    register,
    handleSubmit,
    formState: { errors },
  } = useForm();

  const [login, { loading }] = useMutation(LOGIN_MUTATION);

  const onSubmit = (data: any) => {
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
        console.log("LOGIN SUCCESS");
        const client = getApolloClient();
        client.resetStore();
        if (Router.query.redirect !== undefined) {
          Router.push(Router.query.redirect as string);
        } else {
          Router.push("/");
        }
      })
      .catch((err) => {
        console.error(err);
        // setError(true);
        // setMessage(err.message);
      });
  };

  return (
    <Layout title="Login | Roadmap">
      <main className="grid grid-cols-12">
        {/* TODO: Add side image */}
        <section className="col-start-2 col-end-7"></section>
        <section className="max-w-sm col-start-7 col-end-12">
          <h1 className="text-3xl text-gray-800 font-medium tracking-wide mt-8 mb-3">
            Log in
          </h1>
          {/* TODO: Add error alerts */}
          {/* <Alert message={"message"} error={error} /> */}

          <form className="flex flex-col" onSubmit={handleSubmit(onSubmit)}>
            <Input
              id="email"
              name="Email"
              error={errors.email}
              register={register("email", {
                required: { value: true, message: "Email is required" },
                pattern: {
                  value: new RegExp(
                    "^[a-zA-Z0-9_.+-]+@[a-zA-Z0-9-]+.[a-zA-Z0-9-.]+$"
                  ),
                  message: "Must be a valid email",
                },
              })}
            />

            <Input
              id="password"
              name="Password"
              type="password"
              error={errors.password}
              register={register("password", {
                required: { value: true, message: "Password is required" },
              })}
            />

            <button
              type="submit"
              className="icon-btn-grid items-center bg-green-500 text-white font-medium tracking-wide py-2 rounded disabled:opacity-70"
              disabled={loading}
            >
              <span className="justify-self-end">
                {loading && (
                  <Loader
                    className="fill-current animate-spin mr-2"
                    width={20}
                    height={20}
                  />
                )}
              </span>
              <span>Log in</span>
            </button>
          </form>

          <div className="text-sm text-gray-500 text-center mt-6">
            Don't have an account? <RedirectToSignup />
          </div>
        </section>
      </main>
    </Layout>
  );
}

function RedirectToSignup() {
  const router = useRouter();
  let href: any = {
    pathname: "/register",
  };
  if (router.query.redirect !== undefined) {
    href["query"] = { redirect: router.query.redirect };
  }
  return (
    <Link href={href}>
      <a className="text-blue-500 hover:underline">Sign up</a>
    </Link>
  );
}

export default Login;
