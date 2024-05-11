export interface dnftSale {
  dnft: dnft;
  price: number;
  amount: number;
}

export interface dnft {
  id: string;
  address: string;
  metadata: dnftMetadata; //DA
}

export interface dnftMetadata {
  uri: string; //IPFS
}

export interface cat {
  glasses: string;
  background: string;
}
