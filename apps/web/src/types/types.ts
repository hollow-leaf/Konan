export type nftOwn = {
    owner: string
    nftId: number
    nftName: string
    creatorAddr: string
    url: string
}

export type nftCreate = {
    nftId: number
    nftName: string
    price: number
    maxSupply: number
    creatorId: string
    url: string
    poolContractAddr: string
}

export type account = {
    userId: string
    address: string
    headerImage: string
}

export type nftStatus = {
    background: string
    glasses: string
}