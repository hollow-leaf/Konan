export function Loading(props: any) {
    return (
        
        <div className="flex absolute w-full h-full z-50 rounded-xl items-center justify-center py-12 px-4 sm:px-6 lg:px-8 bg-black/70">
            <div>
            <svg className="ip mt-20" viewBox="0 0 256 128" width="256px" height="128px" xmlns="http://www.w3.org/2000/svg">
                <defs>
                    <linearGradient id="grad1" x1="0" y1="0" x2="1" y2="0">
                        <stop offset="0%" stopColor="#5ebd3e" />
                        <stop offset="33%" stopColor="#ffb900" />
                        <stop offset="67%" stopColor="#f78200" />
                        <stop offset="100%" stopColor="#e23838" />
                    </linearGradient>
                    <linearGradient id="grad2" x1="1" y1="0" x2="0" y2="0">
                        <stop offset="0%" stopColor="#e23838" />
                        <stop offset="33%" stopColor="#973999" />
                        <stop offset="67%" stopColor="#009cdf" />
                        <stop offset="100%" stopColor="#5ebd3e" />
                    </linearGradient>
                </defs>
                <g fill="none" strokeLinecap="round" strokeWidth="16">
                    <g className="ip__track" stroke="#ddd">
                        <path d="M8,64s0-56,60-56,60,112,120,112,60-56,60-56"/>
                        <path d="M248,64s0-56-60-56-60,112-120,112S8,64,8,64"/>
                    </g>
                    <g strokeDasharray="180 656">
                        <path className="ip__worm1" stroke="url(#grad1)" strokeDashoffset="0" d="M8,64s0-56,60-56,60,112,120,112,60-56,60-56"/>
                        <path className="ip__worm2" stroke="url(#grad2)" strokeDashoffset="358" d="M248,64s0-56-60-56-60,112-120,112S8,64,8,64"/>
                    </g>
                </g>
            </svg>
            <h2 className="mt-10 text-center text-3xl font-bold tracking-tight text-white">
                        Loading...
            </h2>
            <div className="flex w-full items-center justify-center mt-10">
                {props.close&&<button type="button" onClick={props.close} className="py-2.5 px-5 me-2 mb-2 text-sm font-medium text-gray-900 focus:outline-none bg-white/70 rounded-lg border border-gray-200/70 hover:bg-gray-100 hover:text-blue-700 focus:z-10 focus:ring-4 focus:ring-gray-100">Cancel</button>}
            </div>
            </div>
        </div>
    )
}