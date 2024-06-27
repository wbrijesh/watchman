import { Kode_Mono } from "next/font/google";
import Link from "next/link";
import { useRouter } from "next/router";
import Breadcrumb from "./Breadcrumb";

const LogoFont = Kode_Mono({ subsets: ["latin"] });

type SidebarOptionType = {
  name: string;
  link: string;
};

const SidebarOptions: SidebarOptionType[] = [
  { name: "Home", link: "/" },
  { name: "Projects", link: "/projects" },
  { name: "Logs", link: "/logs" },
  { name: "Metrics", link: "/metrics" },
  { name: "Analytics", link: "/analytics" },
];

const BookmarkOptions: SidebarOptionType[] = [
  { name: "Example Project", link: "/projects/1" },
];

const MainLayout = ({ children }: { children: React.ReactNode }) => {
  const router = useRouter();
  return (
    <main className="text-slate-600 text-sm">
      <div className="flex">
        <div className="w-72 h-[calc(100vh-0px)] bg-white border-r border-slate-200 flex flex-col justify-between">
          <div>
            <div className="h-14 w-full px-4 border-b border-slate-200 flex items-center justify-between">
              <div>
                <p
                  className={`text-xl text-black font-semibold ${LogoFont.className}`}
                >
                  WATCHMAN
                </p>
              </div>
            </div>
            <div className="flex flex-col">
              <div className="bg-slate-400 text-white text-xs font-bold py-0.5 px-3">
                MENU
              </div>
              {SidebarOptions.map((option) => (
                <Link
                  key={option.name}
                  href={`${option.link}`}
                  className={`border-b border-slate-200 py-2 px-3 hover:bg-slate-100 ${
                    router.pathname.startsWith(option.link) &&
                    option.link !== "/"
                      ? "bg-slate-200"
                      : ""
                  } ${
                    router.pathname === option.link && option.link == "/"
                      ? "bg-slate-200"
                      : ""
                  }`}
                >
                  {option.name}
                </Link>
              ))}
              <div className="bg-slate-400 text-white text-xs font-bold py-0.5 px-3">
                BOOKMARKS
              </div>
              {BookmarkOptions.map((option) => (
                <Link
                  key={option.name}
                  href={option.link}
                  className={`border-b border-slate-200 py-2 px-3 hover:bg-slate-100 ${
                    router.pathname == option.link ? "bg-slate-200" : ""
                  }`}
                >
                  {option.name}
                </Link>
              ))}
            </div>
          </div>

          <div className="border-b border-slate-200 py-2 px-3">
            <p>Settings</p>
          </div>
        </div>
        <div className="w-[calc(100vw-286px)]">
          <Breadcrumb />

          <div className="p-4">{children}</div>
        </div>
      </div>
    </main>
  );
};

export default MainLayout;
