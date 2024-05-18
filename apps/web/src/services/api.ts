import { nftOwn } from "@/types/types";


const  headers = {
    "Content-Type": "application/json",
    "Accept": "application/json",
}

//serverless call
const HOST = "https://psyduck-app.wayneies1206.workers.dev"

export async function getNftcsByUserId(userId:string) {
    const address = await userId2Address(userId)
    try {
        let body = {
            "creatorId": userId
        }
        const res = await fetch(HOST + '/nftCreateByAddress', {
            method: 'POST',
            headers: headers,
            body: JSON.stringify(body)
          })
        if(res){
            const rres = await res.json()
            return rres.nfts
        }else{
            return []
        }
    }
    catch (err) {
        console.log("error", err);
        return []
    }
}

export async function userId2Address(userId:string){
    try {

        let body = {
            "userId": userId
        }

        const res = await fetch(HOST + '/userIdToAddress', {
            method: 'POST',
            headers: headers,
            body: JSON.stringify(body)
          })
        if(res){
            const rres = await res.json()
            return {"address":rres.address, "poolContractAddr": rres.poolContractAddr}
        }else{
            return ""
        }
    }
    catch (err) {
        console.log("error", err);
        return ""
    }
}

export async function getNftosByAddress(Address:string): Promise<nftOwn[]> {
    try {
        let body = {
            "address": Address
        }
        const res = await fetch(HOST + '/nftOwnByAddress', {
            method: 'POST',
            headers: headers,
            body: JSON.stringify(body)
          })
        if(res){
            const rres = await res.json()
            return rres.nfts
        }else{
            return []
        }
    }
    catch (err) {
        console.log("error", err);
        return []
    }
}

export async function getNftImageByIndex(index: number): Promise<string> {
    try {
        let body = {
            "id": index
        }
        const res = await fetch(HOST + '/getNftURIById', {
            method: 'POST',
            headers: headers,
            body: JSON.stringify(body)
          })
        if(res){
            const rres = await res.json()
            return rres.uri.split("//")[1]
        }else{
            return ""
        }
    }
    catch (err) {
        console.log("error", err);
        return ""
    }
}