import React from 'react';

export function Business() {
    return (
        <div className="flex h-screen w-full px-32 py-10 text-white">
            <div className=" w-1/2 flex flex-col items-center justify-center">
                <p className="text-8xl font-extrabold bg-gradient-to-r from-blue-600 via-green-500 to-indigo-400 inline-block text-transparent bg-clip-text">Features</p>
            </div>
            <div className="w-1/2 flex flex-col p-32 gap-10">
                <div className="basis-1/4">
                    <p className="text-4xl font-bold">1</p>
                </div>
                <div className="basis-1/4 flex flex-row gap-5">
                    <p className="text-4xl font-bold">2</p>
                    <p className="py-2">123</p>
                </div>
                <div className="basis-1/4">
                    <p className="text-4xl font-bold">3</p>
                </div>
                <div className="basis-1/4">
                    <p className="text-4xl font-bold">4</p>
                </div>
            </div>
        </div>
    );
}
