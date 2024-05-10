"use client";
import React from "react";
import { Web3Provider } from "@/config/wagmi.config";
import { ThemeProvider, createTheme } from "@mui/material/styles";

const theme = createTheme({
  palette: {
    primary: {
      main: "#000000",
    },
  },
});

export function Layout({ children }: { children: React.ReactNode }) {
  return (
    <Web3Provider>
      <ThemeProvider theme={theme}>
        <div>{children}</div>
      </ThemeProvider>
    </Web3Provider>
  );
}

export default Layout;
