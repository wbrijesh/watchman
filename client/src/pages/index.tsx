import MainLayout from "@/components/MainLayout";
import Head from "next/head";
import Link from "next/link";

const Homepage = () => {
  return (
    <MainLayout>
      <Head>
        <title>Watchman</title>
        <Link
          rel="icon"
          href="data:image/svg+xml,<svg xmlns=%22http://www.w3.org/2000/svg%22 viewBox=%220 0 100 100%22><text y=%22.9em%22 font-size=%2290%22>🎯</text></svg>"
        />
      </Head>
      <h1 className="text-xl font-medium">Homepage</h1>
      <p className="my-2">
        Please check this page later, as it is still under construction.
      </p>
    </MainLayout>
  );
};

export default Homepage;
