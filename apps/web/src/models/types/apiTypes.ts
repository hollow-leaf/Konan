interface dnftSale {
  dnft: dnft;
  price: number;
  amount: number;
}

interface dnft {
  id: string;
  address: string;
  metadata: dnftMetadata; //DA
}

interface dnftMetadata {
  uri: string; //IPFS
}

interface cat {
  glasses: string;
  background: string;
}
