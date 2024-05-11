import React from "react";
import { X as XIcon, GitHub } from "@mui/icons-material";
import { IconButton, Typography, Box } from "@mui/material";

export function Footer() {
  return (
    <Box className="flex w-full shrink-0 flex-col items-center justify-center gap-2 border-t px-4 py-6 sm:flex-row md:px-6 text-black">
      <Typography variant="body1" className="text-black">
        © 2024 Konan Team. All rights reserved.
      </Typography>
      <Box sx={{ display: "flex", gap: 2 }}>
        <IconButton
          component="a"
          href="https://twitter.com/psyduck_eth"
          target="_blank"
          rel="noopener noreferrer"
        >
          <XIcon />
        </IconButton>
        <IconButton
          component="a"
          href="https://github.com/hollow-leaf/Konan"
          target="_blank"
          rel="noopener noreferrer"
        >
          <GitHub />
        </IconButton>
      </Box>
    </Box>
  );
}

export default Footer;
