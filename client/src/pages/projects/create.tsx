import { CreatePorject } from "@/api/projects";
import { PrimaryButton } from "@/components/Button";
import MainLayout from "@/components/MainLayout";
import Head from "next/head";
import { useRouter } from "next/router";
import { useEffect, useState } from "react";

const CreateProjectPage = () => {
  const router = useRouter();

  const [name, setName] = useState("");
  const [apiResponse, setApiResponse] = useState<any>();

  async function formSubmit(e: React.FormEvent) {
    e.preventDefault();
    let project = { Name: name };

    CreatePorject(project, setApiResponse);
  }

  useEffect(() => {
    apiResponse && apiResponse.status === "OK" && router.push("/projects");
  }, [apiResponse]);

  return (
    <MainLayout>
      <Head>
        <title>Create Projects | Watchman</title>
      </Head>
      <h1 className="text-xl font-medium">Create Project</h1>

      <form onSubmit={formSubmit} className="my-4 flex flex-col gap-4">
        <div>
          <label className="block text-sm font-medium">Name</label>
          <input
            value={name}
            onChange={(e) => setName(e.target.value)}
            className="w-80 border border-slate-200 p-1"
          />
        </div>
        <PrimaryButton type="submit" text="Create Project" />
      </form>
    </MainLayout>
  );
};
export default CreateProjectPage;
