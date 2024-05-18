import { blackstoneABI } from '@/contracts/blackstone'
import { useAccount, useWriteContract } from 'wagmi'
import { Message } from './message'

export function MintNftButton(props: {price: number}) {
    const { address, isConnected, isDisconnected } = useAccount()

    const { writeContract, isError, isSuccess, error, isPending } = useWriteContract()

    async function mintHandler () {
        console.log(123)
        writeContract({ 
            abi: blackstoneABI,
            address: '0x585a1DDaB9116F483d367bCa6eb64797252051c8',
            functionName: 'publicSaleMint',
            args: [
              1
            ],
            value: BigInt(1000000000000)
         })
    }

    return (
        <div>
            <div className="flex flex-col justify-center items-center">
                {isConnected?
                <button onClick={mintHandler} type="button" className="text-white bg-gradient-to-r from-purple-500 to-pink-500 hover:bg-gradient-to-l focus:ring-4 focus:outline-none focus:ring-purple-200 dark:focus:ring-purple-800 font-medium rounded-full text-xl px-10 py-6 text-center me-2 mb-2">Mint Now!</button>
                :
                <button type="button" className="text-white bg-gray-200 hover:bg-gradient-to-l focus:ring-4 focus:outline-none focus:ring-purple-200 dark:focus:ring-purple-800 font-medium rounded-full text-xl px-10 py-6 text-center me-2 mb-2" disabled>Mint Now!</button>
                }
                {isConnected?<></>
                :
                <p className='text-gray-400'>Connect wallet and mint!</p>
                }
            </div>
        </div>
    )
}