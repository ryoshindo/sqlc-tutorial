import Head from "next/head";
import "../styles/globals.css";
import type { AppProps } from "next/app";

const App = ({ Component, pageProps }: AppProps) => {
  return (
    <>
      <Head>
        <title>sqlc-tutorial</title>
      </Head>
      <Component {...pageProps} />
    </>
  );
};

export default App;
