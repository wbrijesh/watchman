import { DeleteProject, GetProjectById, ListProjects } from "@/api/projects";
import { PrimaryButton } from "@/components/Button";
import MainLayout from "@/components/MainLayout";
import Head from "next/head";
import { useRouter } from "next/router";
import { useEffect, useState } from "react";

const ProjectDeletePage = () => {
  const router = useRouter();
  const { id } = router.query;

  const [project, setProject] = useState<any>();
  const [apiResponse, setApiResponse] = useState<any>();

  useEffect(() => {
    GetProjectById(id ? id.toString() : "", setProject);
  }, [id]);

  function deleteProject() {
    let project = { ID: id ? id.toString() : "" };

    DeleteProject(project, setApiResponse);
  }

  useEffect(() => {
    apiResponse && apiResponse.status === "OK" && router.push("/projects");
  }, [apiResponse]);

  return (
    <MainLayout>
      <Head>
        <title>Delete Project | Watchman</title>
      </Head>
      <div className="mb-4 flex justify-between items-center uppercase">
        {project?.data?.name && (
          <h1 className="text-xl font-medium">{project.data.name}</h1>
        )}
      </div>
      {project && project.data && project.data.name && (
        <p>Are you sure you want to delete this {project.data.name} project?</p>
      )}
      <div className="my-4 flex gap-2">
        <PrimaryButton
          text="Delete"
          intent="danger"
          onClick={() => deleteProject()}
        />
        <PrimaryButton text="Cancel" onClick={() => router.push("/projects")} />
      </div>
    </MainLayout>
  );
};
export default ProjectDeletePage;
