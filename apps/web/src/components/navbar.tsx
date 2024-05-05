"use client";

import React, { useEffect } from "react";

import Image from "next/image";
import Link from "next/link";

import {
  Web as WebIcon,
  AccountCircle as AccountCircleIcon,
  Code as CodeIcon,
  Close as CloseIcon,
  Menu as MenuIcon,
  Book as BookIcon,
} from "@mui/icons-material";
import {
  AppBar,
  Button,
  IconButton,
  Toolbar,
  Typography,
  Collapse,
  List,
  ListItem,
  ListItemIcon,
  ListItemText,
  Box,
} from "@mui/material";

import { NavbarItemType } from "@/models/types/uiTypes";
import { ConnectButton } from "@/components";

const NAV_MENU: NavbarItemType[] = [
  {
    name: "Profile",
    icon: <AccountCircleIcon />,
    href: "/nft-mint",
    target: "_blank",
  },
  {
    name: "GitHub",
    icon: <CodeIcon />,
    href: "https://github.com/hollow-leaf/Konan",
    target: "_blank",
  },
];

export function Navbar() {
  const [open, setOpen] = React.useState(false);

  const handleOpen = () => setOpen((cur) => !cur);

  useEffect(() => {
    window.addEventListener(
      "resize",
      () => window.innerWidth >= 960 && setOpen(false),
    );
  }, []);

  useEffect(() => {
    if (window.location.hash) {
      const element = document.getElementById(
        window.location.hash.substring(1),
      );
      if (element) {
        element.scrollIntoView({ behavior: "smooth" });
      }
    }
  }, []);

  return (
    <AppBar
      position="sticky"
      color="transparent"
      elevation={0}
    >
      <Box className="mx-5 rounded-3xl mt-5 shadow-lg bg-slate-200 bg-opacity-30 text-black">
      <Toolbar>
        <Typography variant="h6" component="div" sx={{ flexGrow: 1 }} className="">
          <Link href="/">
            <Typography className="font-bold text-xl">Konan</Typography>
          </Link>
        </Typography>
        <Box
          sx={{
            display: { xs: "none", lg: "flex" },
            gap: 2,
            alignItems: "center",
          }}
        >
          {NAV_MENU.map(({ name, icon: Icon, href, target }) => (
            <Link key={name} href={href} target={target} passHref>
              <Button
                startIcon={Icon}
                className="rounded-full hover:shadow font-bold"
              >
                {name}
              </Button>
            </Link>
          ))}
        </Box>
        <div style={{ flexGrow: 1 }} />
        <Box sx={{ display: { xs: "none", lg: "flex" } }}>
          <ConnectButton />
        </Box>
        <IconButton
          edge="start"
          color="inherit"
          aria-label="menu"
          onClick={handleOpen}
          sx={{ display: { xs: "block", lg: "none" } }}
        >
          {open ? <CloseIcon /> : <MenuIcon />}
        </IconButton>
      </Toolbar>
      <Collapse
        className="bg-sky-500/10 flex"
        in={open}
        timeout="auto"
        unmountOnExit
      >
        <List component="nav" className="flex flex-col items-start">
          {NAV_MENU.map(({ name, icon: Icon, href, target }) => (
            <Link key={name} href={href} target={target} passHref>
              <Box sx={{ display: "block", textAlign: "left", pt: 1 }}>
                <Button
                  onClick={handleOpen}
                  className="text-white rounded-full hover:shadow  font-bold"
                >
                  <ListItemIcon className="text-white">
                    {Icon}
                  </ListItemIcon>
                  <ListItemText primary={name} />
                </Button>
              </Box>
            </Link>
          ))}
          <ListItem>
            <ConnectButton />
          </ListItem>
        </List>
      </Collapse>
      </Box>
    </AppBar>
  );
}

export default Navbar;
