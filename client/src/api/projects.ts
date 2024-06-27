// type APIResponseType = {
//   Status: string;
//   Message: string;
//   RequestID: string;
//   Data: any;
// };

type ProjectType = {
  ID?: string;
  Name?: string;
};

async function CreatePorject(
  projectObject: ProjectType,
  setter: (response: any) => void,
) {
  fetch("http://127.0.0.1:4000/projects", {
    method: "POST",
    body: JSON.stringify(projectObject),
  })
    .then((res) => res.json())
    .then((data) => {
      setter(data);
    });
}

function ListProjects(setter: any) {
  fetch("http://127.0.0.1:4000/projects")
    .then((res) => res.json())
    .then((data) => {
      setter(data);
    });
}

function GetProjectById(projectID: string, setter: any) {
  projectID &&
    fetch("http://127.0.0.1:4000/project?id=" + projectID)
      .then((res) => res.json())
      .then((data) => {
        setter(data);
      });
}

function UpdateProject(
  projectObject: ProjectType,
  setter: (response: any) => void,
) {
  fetch("http://127.0.0.1:4000/project", {
    method: "PUT",
    body: JSON.stringify(projectObject),
  })
    .then((res) => res.json())
    .then((data) => {
      setter(data);
    });
}

function DeleteProject(
  projectObject: ProjectType,
  setter: (response: any) => void,
) {
  fetch("http://127.0.0.1:4000/project", {
    method: "DELETE",
    body: JSON.stringify(projectObject),
  })
    .then((res) => res.json())
    .then((data) => {
      setter(data);
    });
}

export {
  CreatePorject,
  ListProjects,
  GetProjectById,
  UpdateProject,
  DeleteProject,
};
