import React from 'react';
import Card from '@mui/material/Card';
import CardContent from '@mui/material/CardContent';
import CardMedia from '@mui/material/CardMedia';
import Typography from '@mui/material/Typography';
import { CardActionArea, CardActions } from '@mui/material';
import { nftIPFS } from '@/content/ipfs/ipfs-nft';
import Link from 'next/link';

export function NFTCard(props: { index: number }) {

  return (
    <Card sx={{ maxWidth: 345 }}>
      <CardActionArea>
        <CardMedia
          component="img"
          height="140"
          image={nftIPFS[props.index].src}
          alt={nftIPFS[props.index].name}
          className="transform transition-transform hover:scale-110"
        />
        <CardContent>
          <Typography gutterBottom variant="h5" component="div">
            {nftIPFS[props.index].name}
          </Typography>
          <Typography variant="body2" color="text.secondary">
            Solo is gold System Architect.
          </Typography>
        </CardContent>
      </CardActionArea>
      <CardActions>
        <Link color="primary" href={nftIPFS[props.index].src} className='hover:text-shadow-lg'>
          Show details
        </Link>
      </CardActions>
    </Card>
  );
}
