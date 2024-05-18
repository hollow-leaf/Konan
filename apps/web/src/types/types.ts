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

export type nftMeatadata = {
    description: string
    external_url: string
    image: string
    name: string
    attributes: nftAttributes[]
}

export type nftAttributes = {
    trait_type: string
    value: string
}