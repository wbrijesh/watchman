import { ListProjects } from "@/api/projects";
import { PrimaryButton } from "@/components/Button";
import MainLayout from "@/components/MainLayout";
import Table from "@/components/table";
import Link from "next/link";
import { useRouter } from "next/router";
import { useEffect, useState } from "react";

const ProjectsHomePage = () => {
  const router = useRouter();

  const [projects, setProjects] = useState<any>([]);

  useEffect(() => {
    ListProjects(setProjects);
  }, []);

  return (
    <MainLayout>
      <div className="mb-4 flex justify-between items-center">
        <h1 className="text-xl font-medium">Projects</h1>
        <PrimaryButton
          text="Create Project"
          onClick={() => router.push("/projects/create")}
        />
      </div>
      {projects.data !== undefined && (
        <Table
          columns={["Name", "ID"]}
          data={projects.data.map((project: any) => {
            return {
              name: (
                <Link href={"/projects/" + project.id}>{project.name}</Link>
              ),
              id: project.id,
            };
          })}
        />
      )}
    </MainLayout>
  );
};
export default ProjectsHomePage;
