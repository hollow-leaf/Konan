import React from "react";

export function CollectionCard() {
  return (
    <div className="bg-white rounded-3xl shadow p-4 m-4 bg-opacity-30">
      <div className="flex mb-2">
        <div className="m-2">
          <img
            src={"/konan-hero.webp"}
            alt="xuedao logo"
            width={80}
            height={80}
            className="rounded-full"
          />
        </div>
        <div className="m-2 content-center">
          <p className="text-xl font-bold	">#Polly Doll</p>
          <p>@By TheSalvare</p>
        </div>
      </div>
      <div className="flex">
        <div className="mt-2 min-w-40 min-h-56 rounded-lg bg-center bg-cover bg-[url('/konan-hero.webp')]"></div>
        <div className="mt-2 ml-4">
          <div className="min-w-40 min-h-28 rounded-lg bg-center bg-cover bg-[url('/konan-hero.webp')]"></div>
          <div className="mt-4 min-w-40 min-h-28 rounded-lg bg-center bg-cover bg-[url('/konan-hero.webp')]"></div>
        </div>
      </div>
    </div>
  );
}
