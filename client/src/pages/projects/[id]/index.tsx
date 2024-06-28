import { GetProjectById, ListProjects } from "@/api/projects";
import { PrimaryButton } from "@/components/Button";
import MainLayout from "@/components/MainLayout";
import Head from "next/head";
import { useRouter } from "next/router";
import { useEffect, useState } from "react";

const ProjectDetailsPage = () => {
  const router = useRouter();
  const { id } = router.query;

  const [project, setProject] = useState<any>();

  useEffect(() => {
    GetProjectById(id ? id.toString() : "", setProject);
  }, [id]);

  return (
    <MainLayout>
      <Head>
        <title>Project Details | Watchman</title>
      </Head>
      <div className="mb-4 flex justify-between items-center uppercase">
        {project?.data?.name && (
          <h1 className="text-xl font-medium">{project.data.name}</h1>
        )}
        <div className="flex items-center gap-2">
          <PrimaryButton
            text="Update"
            onClick={() => router.push("/projects/" + id + "/update")}
          />
          <PrimaryButton
            text="Delete"
            intent="danger"
            onClick={() => router.push("/projects/" + id + "/delete")}
          />
        </div>
      </div>
      {project && project.data && project.data.name && (
        <>
          <div className="flex justify-start">
            <div className="w-48">Name</div>
            <p>{project.data.name}</p>
          </div>
          <div className="flex justify-start">
            <div className="w-48">ID</div>
            <p>{project.data.id}</p>
          </div>
        </>
      )}
    </MainLayout>
  );
};
export default ProjectDetailsPage;
