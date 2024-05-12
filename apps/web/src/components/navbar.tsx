"use client";
import React, { useEffect } from "react";

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

import { NavbarItemType } from "@/types/uiTypes";
import { ConnectButton } from "@/components";

const NAV_MENU: NavbarItemType[] = [
  {
    name: "Profile",
    icon: <AccountCircleIcon />,
    href: "/profile",
    target: "",
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
    <AppBar position="sticky" color="transparent" elevation={0}>
      <Box className="shadow-lg bg-white text-black">
        <Toolbar>
          <Typography
            variant="h6"
            component="div"
            sx={{ flexGrow: 1 }}
            className=""
          >
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
                  className="rounded-full hover:shadow font-bold text-black"
                >
                  {name}
                </Button>
              </Link>
            ))}
          </Box>
          <Box style={{ flexGrow: 1 }} />
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
          className="bg-white flex"
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
                    className="rounded-full hover:shadow font-bold text-black"
                  >
                    <ListItemIcon className="text-black">{Icon}</ListItemIcon>
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
