"use client";
import { Box, IconButton, Link, Typography } from "@mui/material";
import React from "react";
import Image from "next/image";
import { images } from "@/content/TeamData";
import { LinkedIn } from "@mui/icons-material";

export function Team() {
  return (
    <Box className="flex flex-col justify-center items-center py-10">
      <Box className="flex flex-col xl:flex-row py-10 xl:py-36 gap-5">
        <Typography className="text-4xl xl:text-8xl font-bold text-black">
          MEET OUR
        </Typography>
        <Typography className="text-4xl xl:text-8xl font-extrabold	bg-gradient-to-r from-[#3F09F9] via-[#CB6BE0] to-[#CE2CE5] inline-block text-transparent bg-clip-text">
          TEAM
        </Typography>
      </Box>
      <Box className="flex flex-col xl:flex-row mx-auto items-center justify-center align-middle">
        {images.map((image, index) => (
          <Box className="px-5" key={index}>
            <Image
              src={image.name}
              alt={image.label}
              width={300}
              height={400}
              className="rounded-2xl"
            />
            <div className="p-2 space-y-2">
              <div>
                <p className="text-2xl font-bold text-black">{image.label}</p>
                <p className="text-xl font-bold text-orange-300">{image.title}</p>
              </div>
              <div>
                <Link href={image.linkedin}>
                  <LinkedIn className="text-2xl hover:text-gray-500" />
                </Link>
              </div>
            </div>
          </Box>
        ))}
      </Box>
    </Box>
  );
}
