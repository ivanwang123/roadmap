import { RedirectLink } from "components/element";
import { Form, InputField, SubmitButton } from "components/form";
import { Layout } from "components/layout";
import { useRegister } from "modules/auth";
import { Notification } from "modules/notification";
import React from "react";
import { z } from "zod";

type RegisterValues = {
  email: string;
  username: string;
  password: string;
};

const schema = z.object({
  email: z.string().email("Must be a valid email"),
  username: z.string().min(1, "Username is required"),
  password: z.string().min(8, "Password must be at least 8 characters long"),
});

function Register() {
  const { register: signup, loading } = useRegister();

  return (
    <Layout title="Sign up | Roadmap">
      <main className="grid grid-cols-12">
        <section className="col-start-2 col-end-7"></section>
        <section className="max-w-sm col-start-7 col-end-12">
          <div className="bg-secondary p-8 mt-8 rounded-sm">
            <h1 className="text-3xl text-black font-semibold mb-3">Sign up</h1>

            <Form<RegisterValues, typeof schema>
              onSubmit={signup}
              schema={schema}
            >
              <>
                <InputField label="EMAIL" id="email" />
                <InputField label="USERNAME" id="username" />
                <InputField label="PASSWORD" id="password" type="password" />
                <SubmitButton label="Sign up" loading={loading} />
              </>
            </Form>
            <Notification type="text" style="mt-1" showOnly="error" />

            <div className="text-gray-500 text-sm text-center mt-6">
              Already have an account?{" "}
              <RedirectLink label="Log in" pathname="/login" />
            </div>
          </div>
        </section>
      </main>
    </Layout>
  );
}

export default Register;
