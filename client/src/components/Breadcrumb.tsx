import Link from "next/link";
import { useRouter } from "next/router";

const Breadcrumb = () => {
  const router = useRouter();
  return (
    <div className="bg-slate-100 py-1 px-2 border-b border-slate-200">
      {router.pathname.startsWith("/projects") && (
        <div className="flex gap-0.5 items-center">
          <p>/</p>
          <Link className="hover:underline" href="/projects">
            Projects
          </Link>
          {router.pathname.replace("/projects", "") === "/create" && (
            <>
              <p>/</p>
              <Link className="hover:underline" href={"/projects/create"}>
                Create Project
              </Link>
            </>
          )}
          {/* if path contains only one slash, it means page is project details */}
          {(router.pathname.match(new RegExp("/", "g")) || []).length == 2 &&
            !router.pathname.includes("/create") &&
            !router.pathname.includes("/update") &&
            !router.pathname.includes("/delete") && (
              <>
                <p>/</p>
                <Link
                  className="hover:underline"
                  href={`/projects/${router.query.id}`}
                >
                  Project Details
                </Link>
              </>
            )}
        </div>
      )}
    </div>
  );
};

export default Breadcrumb;
