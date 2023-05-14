import Head from "next/head";
import "../styles/globals.css";
import type { AppProps } from "next/app";
import {
  ApolloClient,
  ApolloLink,
  ApolloProvider,
  createHttpLink,
  InMemoryCache,
} from "@apollo/client";
import { StyledEngineProvider } from "@mui/material/styles";

const App = ({ Component, pageProps }: AppProps) => {
  const client = new ApolloClient({
    cache: new InMemoryCache(),
    link: ApolloLink.from([
      createHttpLink({
        uri: process.env.NEXT_PUBLIC_API_ENDPOINT,
        credentials: "include",
      }),
    ]),
  });

  return (
    <>
      <Head>
        <title>sqlc-tutorial</title>
      </Head>

      <ApolloProvider client={client}>
        <StyledEngineProvider injectFirst>
          <Component {...pageProps} />
        </StyledEngineProvider>
      </ApolloProvider>
    </>
  );
};

export default App;
