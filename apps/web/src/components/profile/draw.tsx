import { blackstoneABI } from '@/contracts/blackstone'
import { type UseWriteContractReturnType } from 'wagmi'
import { useAccount, useWriteContract } from 'wagmi'
import { Message } from '../message'

export function Draw(props: {tokenId: number, drawId: number, url: string, type:string}) {
    const { address, isConnecting, isDisconnected } = useAccount()

    const { writeContract, isError, isSuccess, error } = useWriteContract()

    console.log(props.tokenId)

    async function drawHandler () {
        writeContract({ 
            abi: blackstoneABI,
            address: '0x585a1DDaB9116F483d367bCa6eb64797252051c8',
            functionName: 'drawByVRF',
            args: [
              props.tokenId,
              props.drawId,
            ],
         })
    }

    return (
        <div>
            <div className="flex justify-center items-center">
                <div onClick={drawHandler} className="flex absolute w-40 h-40 z-50 opacity-0 text-2xl font-medium cursor-pointer hover:opacity-90 rounded-xl items-center justify-center bg-black/40">UnLock!</div>
                <div className="flex absolute w-36 h-36 z-40 opacity-30 text-2xl font-medium cursor-pointer rounded-xl items-center justify-center bg-center bg-cover bg-black/10"></div>
                <img className="rounded-lg border-2 w-40 h-40 opacity-40 p-2" src={props.url} alt="image description" />
                {isError&&<Message message={error?error.name:""}/>}
            </div>
            <p className='w-full text-lg text-center p-2 font-bolds' id='poetsen-one-regular'>{props.type}</p>
        </div>
    )
}