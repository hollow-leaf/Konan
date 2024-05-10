"use client";
import React, { useEffect, useState } from "react";
import Image from "next/image";
import { LinkedIn } from "@mui/icons-material";
import Link from "next/link";
import SwipeableViews from "react-swipeable-views";
import { autoPlay } from "react-swipeable-views-utils";
import { ImageType } from "@/models/types/uiTypes";
import { Box, Typography } from "@mui/material";

const AutoPlaySwipeableViews = autoPlay(SwipeableViews);
const images: ImageType[] = [
  {
    label: "Solo",
    name: "/Solo.png",
    title: "System Architect",
    linkedin: "https://www.linkedin.com/in/cheng-wei-lin-350698226/",
  },
  {
    label: "Jake",
    name: "/Jake.png",
    title: "Frontend Engineer",
    linkedin: "https://www.linkedin.com/in/jake-kuo-83b186245/",
  },
  {
    label: "Albert",
    name: "/Albert.png",
    title: "Smart Contract Developer",
    linkedin: "https://www.linkedin.com/in/chun-yuan-cheng-321870263/",
  },
  {
    label: "Sophia",
    name: "/Sophia.png",
    title: "UI/UX Designer",
    linkedin: "https://www.linkedin.com/in/liang1997/",
  },
];

export function AutoPlaySwipe() {
  const [index, setIndex] = useState(0);

  const handleChangeIndex = (index: number) => {
    setIndex(index);
  };

  const [isMobile, setIsMobile] = useState(false);
  useEffect(() => {
    function handleResize() {
      setIsMobile(window.innerWidth <= 768); //if window <= 768px, set the widow is mobile.
    }
    //initialize handle
    handleResize();
    //monitor window size
    window.addEventListener("resize", handleResize);
    // Clean up
    return () => {
      window.removeEventListener("resize", handleResize);
    };
  }, []);

  return (
    <>
      {isMobile ? (
        <Box className="flex flex-col justify-center items-center">
          <AutoPlaySwipeableViews
            index={index}
            onChangeIndex={handleChangeIndex}
            axis="x"
            animateHeight={true}
            enableMouseEvents
          >
            {images.map((image, index) => (
              <Image
                key={index}
                src={image.name}
                alt={`${image.label}`}
                className="rounded-3xl"
                width={500}
                height={300}
                priority
              />
            ))}
          </AutoPlaySwipeableViews>

          <Box className="flex flex-col items-center justify-center">
            <Typography className="text-2xl font-bold text-black">
              {images[index].label}
            </Typography>
            <Typography className="text-2xl font-bold text-black">
              {images[index].title}
            </Typography>
            <Typography className="text-2xl font-bold text-black">
              {index + 1}/4
            </Typography>
          </Box>
        </Box>
      ) : (
        <>
          <AutoPlaySwipeableViews
            index={index}
            onChangeIndex={handleChangeIndex}
            axis="x"
            animateHeight={true}
            enableMouseEvents
          >
            {images.map((image, index) => (
              <Box
                key={index}
                className="rounded-3xl justify-center items-center flex flex-col"
              >
                <Box className="flex flex-col xl:flex-row gap-10">
                  <Image
                    src={image.name}
                    alt={`${image.label}`}
                    className="rounded-3xl"
                    width={500}
                    height={300}
                    priority
                  />

                  <Box className="flex flex-col justify-center">
                    <Typography className="text-4xl xl:text-8xl font-bold">
                      {image.label}
                    </Typography>
                    <Typography className="text-2xl xl:text-4xl font-bold">
                      {image.title}
                    </Typography>
                    <Link
                      href={`${image.linkedin}`}
                      className="xl:mt-72"
                      target="_blank"
                    >
                      <LinkedIn className="text-black text-4xl" />
                    </Link>
                  </Box>
                </Box>

                <Box className="text-black text-3xl font-bold flex justify-center items-center my-5">
                  {index + 1}/4
                </Box>
              </Box>
            ))}
          </AutoPlaySwipeableViews>
        </>
      )}
    </>
  );
}
