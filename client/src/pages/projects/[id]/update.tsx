import { GetProjectById, UpdateProject } from "@/api/projects";
import { PrimaryButton } from "@/components/Button";
import MainLayout from "@/components/MainLayout";
import Head from "next/head";
import { useRouter } from "next/router";
import { useEffect, useState } from "react";

const UpdateProjectPage = () => {
  const router = useRouter();
  const { id } = router.query;

  const [project, setProject] = useState<any>();

  useEffect(() => {
    GetProjectById(id ? id.toString() : "", setProject);
  }, [id]);

  const [name, setName] = useState(project?.data?.name || "");
  const [apiResponse, setApiResponse] = useState<any>();

  async function formSubmit(e: React.FormEvent) {
    e.preventDefault();
    let project = { Name: name, Id: id };

    UpdateProject(project, setApiResponse);
  }

  useEffect(() => {
    apiResponse && apiResponse.status === "OK" && router.push("/projects");
  }, [apiResponse]);

  useEffect(() => {
    setName(project?.data?.name || "");
  }, [project]);

  return (
    <MainLayout>
      <Head>
        <title>Update Project | Watchman</title>
      </Head>
      <h1 className="text-xl font-medium">Update Project</h1>

      <form onSubmit={formSubmit} className="my-4 flex flex-col gap-4">
        <div>
          <label className="block text-sm font-medium">Name</label>
          <input
            value={name}
            onChange={(e) => setName(e.target.value)}
            className="w-80 border border-slate-200 p-1"
          />
        </div>
        <PrimaryButton type="submit" text="Update Project" />
      </form>
    </MainLayout>
  );
};
export default UpdateProjectPage;
