import { RedirectLink } from "components/element";
import { Form, InputField, SubmitButton } from "components/form";
import { Layout } from "components/layout";
import { useLogin } from "modules/auth";
import { Notification } from "modules/notification";
import React from "react";
import { z } from "zod";

type LoginValues = {
  email: string;
  password: string;
};

const schema = z.object({
  email: z.string().email("Must be a valid email"),
  password: z.string().min(1, "Password is required"),
});

function Login() {
  const { login, loading } = useLogin();

  return (
    <Layout title="Login | Roadmap">
      <main className="grid grid-cols-12">
        {/* TODO: Add side image */}
        <section className="col-start-2 col-end-7"></section>
        <section className="max-w-sm col-start-7 col-end-12">
          <div className="bg-secondary p-8 mt-8 rounded-sm">
            <h1 className="text-3xl text-gray-800 font-medium tracking-wide mb-3">
              Log in
            </h1>

            <Form<LoginValues, typeof schema> onSubmit={login} schema={schema}>
              <>
                <InputField label="EMAIL" id="email" />
                <InputField label="PASSWORD" id="password" type="password" />
                <SubmitButton label="Log in" loading={loading} />
              </>
            </Form>
            <Notification type="text" style="mt-1" showOnly="error" />

            <div className="text-gray-500 text-sm text-center mt-6">
              Don't have an account?{" "}
              <RedirectLink label="Sign up" pathname="/register" />
            </div>
          </div>
        </section>
      </main>
    </Layout>
  );
}

export default Login;
