import React from "react";
import Grid from "@mui/material/Grid";
import Box from "@mui/material/Box";
import Typography from "@mui/material/Typography";
import Skeleton from "@mui/material/Skeleton";
import Image from "next/image";
import { NFTdata } from "@/content/NftData";
import { MediaProps } from "@/interface/MediaInterface";

function Media(props: MediaProps) {
  const { loading = false } = props;

  return (
    <Grid container>
      {(loading ? Array.from(new Array(4)) : NFTdata).map((item, index) => (
        <Box key={index} sx={{ width: 420, marginRight: 0.5, my: 5 }}>
          {item ? (
            <Image
              width={420}
              height={236}
              alt={item.title}
              src={item.src}
              className="rounded-3xl"
            />
          ) : (
            <Skeleton variant="rectangular" width={420} height={236} />
          )}
          {item ? (
            <Box sx={{ pr: 2 }}>
              <Typography gutterBottom variant="body2">
                {item.title}
              </Typography>
              <Typography
                display="block"
                variant="caption"
                color="text.secondary"
              >
                {item.channel}
              </Typography>
              <Typography variant="caption" color="text.secondary">
                {`${item.views} • ${item.createdAt}`}
              </Typography>
            </Box>
          ) : (
            <Box sx={{ pt: 0.5 }}>
              <Skeleton />
              <Skeleton width="60%" />
            </Box>
          )}
        </Box>
      ))}
    </Grid>
  );
}

export function NftTable() {
  return (
    <Box sx={{ overflow: "hidden" }}>
      <Media />
    </Box>
  );
}
