import React from "react"

import Image from "next/image"
import Link from "next/link"
import {
  Instagram as InstagramIcon,
  Telegram as TelegramIcon,
  X as XIcon,
  Facebook as FacebookIcon,
  LinkedIn as LinkedInIcon,
} from "@mui/icons-material"
import { IconButton, Typography, Box } from "@mui/material"

export function Footer() {
  return (
    <Box className="flex w-full shrink-0 flex-col items-center justify-center gap-2 border-t bg-black px-4 py-6 sm:flex-row md:px-6">
      <Typography variant="body1" className="text-gray-400">
        © 2024 XueDAO foundation. All rights reserved.
      </Typography>
      <Box sx={{ display: "flex", gap: 2 }}>
        <IconButton
          className="bg-gray-500 hover:text-slate-300"
          component="a"
          href="/"
          target="_blank"
          rel="noopener noreferrer"
        >
          <InstagramIcon />
        </IconButton>
        <IconButton
          className="bg-gray-500 hover:text-slate-300"
          component="a"
          href="/"
          target="_blank"
          rel="noopener noreferrer"
        >
          <XIcon />
        </IconButton>
        <IconButton
          className="bg-gray-500 hover:text-slate-300"
          component="a"
          href="/"
          target="_blank"
          rel="noopener noreferrer"
        >
          <TelegramIcon />
        </IconButton>
        <IconButton
          className="bg-gray-500 hover:text-slate-300"
          component="a"
          href="/"
          target="_blank"
          rel="noopener noreferrer"
        >
          <FacebookIcon />
        </IconButton>
        <IconButton
          className="bg-gray-500 hover:text-slate-300"
          component="a"
          href="/"
          target="_blank"
          rel="noopener noreferrer"
        >
          <LinkedInIcon />
        </IconButton>
      </Box>
      <Box sx={{ display: "flex", gap: 2 }}>
        <Typography variant="body2" className="text-gray-400 opacity-50 hover:text-slate-300">
          <Link href="/">Privacy Policy</Link>
        </Typography>
      </Box>
    </Box>
  )
}

export default Footer
