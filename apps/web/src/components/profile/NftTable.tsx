import React from "react";
import Grid from "@mui/material/Grid";
import Box from "@mui/material/Box";
import Skeleton from "@mui/material/Skeleton";
import { MediaProps } from "@/interface/MediaInterface";
import { NFTCard } from "./NFTCard";
import { nftIPFS } from "@/content/ipfs/ipfs-nft";
import { IPFSTyeps } from "@/types/IPFSTypes";

function Media(props: MediaProps) {
  const { loading = false } = props;

  return (
    <Grid container>
      {(loading ? Array.from(new Array(nftIPFS.length)) : nftIPFS).map(
        (item, index) => (
          <Box key={index} sx={{ width: 420, marginRight: 0.5, my: 5 }}>
            {item ? (
              <NFTCard index={index} />
            ) : (
              <Skeleton variant="rectangular" width={345} height={360} />
            )}
          </Box>
        ),
      )}
    </Grid>
  );
}

export function NftTable() {
  return (
    <Box sx={{ overflow: "hidden" }}>
      {/* <Media loading /> */}
      <Media />
    </Box>
  );
}
