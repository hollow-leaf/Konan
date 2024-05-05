import React from "react";

import Image from "next/image";
import Link from "next/link";
import {
  Instagram as InstagramIcon,
  Telegram as TelegramIcon,
  X as XIcon,
  Facebook as FacebookIcon,
  LinkedIn as LinkedInIcon,
} from "@mui/icons-material";
import { IconButton, Typography, Box } from "@mui/material";

export function Footer() {
  return (
    <Box className="flex w-full shrink-0 flex-col items-center justify-center gap-2 border-t px-4 py-6 sm:flex-row md:px-6 text-black">
      <Typography variant="body1" className="text-black">
        © 2024 Konan foundation. All rights reserved.
      </Typography>
      <Box sx={{ display: "flex", gap: 2 }}>
        <IconButton
          component="a"
          href="/"
          target="_blank"
          rel="noopener noreferrer"
        >
          <InstagramIcon />
        </IconButton>
        <IconButton
          component="a"
          href="/"
          target="_blank"
          rel="noopener noreferrer"
        >
          <XIcon />
        </IconButton>
        <IconButton
          component="a"
          href="/"
          target="_blank"
          rel="noopener noreferrer"
        >
          <TelegramIcon />
        </IconButton>
        <IconButton
          component="a"
          href="/"
          target="_blank"
          rel="noopener noreferrer"
        >
          <FacebookIcon />
        </IconButton>
        <IconButton
          component="a"
          href="/"
          target="_blank"
          rel="noopener noreferrer"
        >
          <LinkedInIcon />
        </IconButton>
      </Box>
      <Box sx={{ display: "flex", gap: 2 }}>
        <Typography
          variant="body2"
        >
          <Link href="/">Privacy Policy</Link>
        </Typography>
      </Box>
    </Box>
  );
}

export default Footer;
