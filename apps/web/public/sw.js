if(!self.define){let e,a={};const s=(s,c)=>(s=new URL(s+".js",c).href,a[s]||new Promise((a=>{if("document"in self){const e=document.createElement("script");e.src=s,e.onload=a,document.head.appendChild(e)}else e=s,importScripts(s),a()})).then((()=>{let e=a[s];if(!e)throw new Error(`Module ${s} didn’t register its module`);return e})));self.define=(c,i)=>{const n=e||("document"in self?document.currentScript.src:"")||location.href;if(a[n])return;let t={};const f=e=>s(e,n),d={module:{uri:n},exports:t,require:f};a[n]=Promise.all(c.map((e=>d[e]||f(e)))).then((e=>(i(...e),t)))}}define(["./workbox-9564d7f6"],(function(e){"use strict";importScripts(),self.skipWaiting(),e.clientsClaim(),e.precacheAndRoute([{url:"/Albert.png",revision:"ea726bb8c9062fc5d2ccf47c91a2c208"},{url:"/Jake.png",revision:"c6e6f0e6d30c0246a2d2a8f5c0735f24"},{url:"/Solo.png",revision:"151000ceb2ea9efb49fc1f3b81845d7c"},{url:"/Sophia.png",revision:"ea9376ae496798168cc37f50a3b9841b"},{url:"/_next/static/chunks/305.3470641eb77f90f2.js",revision:"3470641eb77f90f2"},{url:"/_next/static/chunks/340-9457a50c24d300a8.js",revision:"fjk7NYZdFw9k_DcacxH3r"},{url:"/_next/static/chunks/456-2a384064e4fdccbe.js",revision:"fjk7NYZdFw9k_DcacxH3r"},{url:"/_next/static/chunks/50eabd66-bbc951c7b9f09013.js",revision:"fjk7NYZdFw9k_DcacxH3r"},{url:"/_next/static/chunks/515-b6a6273ad5d05541.js",revision:"fjk7NYZdFw9k_DcacxH3r"},{url:"/_next/static/chunks/538.2c8ff711069127f4.js",revision:"2c8ff711069127f4"},{url:"/_next/static/chunks/548.42ac187d48965f60.js",revision:"42ac187d48965f60"},{url:"/_next/static/chunks/565.e84857e45b05ac24.js",revision:"e84857e45b05ac24"},{url:"/_next/static/chunks/632.998d56dcd8c76588.js",revision:"998d56dcd8c76588"},{url:"/_next/static/chunks/658.a2c1d37e24a31c47.js",revision:"a2c1d37e24a31c47"},{url:"/_next/static/chunks/693-4ff1ae7c6b52ed1f.js",revision:"fjk7NYZdFw9k_DcacxH3r"},{url:"/_next/static/chunks/716.b7e578ed8f9d2a7f.js",revision:"b7e578ed8f9d2a7f"},{url:"/_next/static/chunks/752.fbf4b464afd712ac.js",revision:"fbf4b464afd712ac"},{url:"/_next/static/chunks/804.3948508efd883b78.js",revision:"3948508efd883b78"},{url:"/_next/static/chunks/93e2f757-e6fe9f882bc1ea88.js",revision:"fjk7NYZdFw9k_DcacxH3r"},{url:"/_next/static/chunks/app/_not-found/page-e1337a38ff54e7f4.js",revision:"fjk7NYZdFw9k_DcacxH3r"},{url:"/_next/static/chunks/app/layout-0963aa1db2a518ac.js",revision:"fjk7NYZdFw9k_DcacxH3r"},{url:"/_next/static/chunks/app/page-cde75498508fb4e5.js",revision:"fjk7NYZdFw9k_DcacxH3r"},{url:"/_next/static/chunks/app/profile/page-b00b5a6cdaa0b8ef.js",revision:"fjk7NYZdFw9k_DcacxH3r"},{url:"/_next/static/chunks/c107fa1b.445a44a080d3ab21.js",revision:"445a44a080d3ab21"},{url:"/_next/static/chunks/framework-036af537c76198b1.js",revision:"fjk7NYZdFw9k_DcacxH3r"},{url:"/_next/static/chunks/main-8e39ec0d9308a966.js",revision:"fjk7NYZdFw9k_DcacxH3r"},{url:"/_next/static/chunks/main-app-46de5de00abf01f2.js",revision:"fjk7NYZdFw9k_DcacxH3r"},{url:"/_next/static/chunks/pages/_app-65fc6011b48e8393.js",revision:"fjk7NYZdFw9k_DcacxH3r"},{url:"/_next/static/chunks/pages/_error-b242a615af41f779.js",revision:"fjk7NYZdFw9k_DcacxH3r"},{url:"/_next/static/chunks/polyfills-78c92fac7aa8fdd8.js",revision:"79330112775102f91e1010318bae2bd3"},{url:"/_next/static/chunks/webpack-e788b387cdcb23f0.js",revision:"fjk7NYZdFw9k_DcacxH3r"},{url:"/_next/static/css/ecc8bf6fe536d05f.css",revision:"ecc8bf6fe536d05f"},{url:"/_next/static/fjk7NYZdFw9k_DcacxH3r/_buildManifest.js",revision:"f6c21f93baf19d71e917a8f8e41b8424"},{url:"/_next/static/fjk7NYZdFw9k_DcacxH3r/_ssgManifest.js",revision:"b6652df95db52feb4daf4eca35380933"},{url:"/_next/static/media/0e4fe491bf84089c-s.p.woff2",revision:"5e22a46c04d947a36ea0cad07afcc9e1"},{url:"/_next/static/media/1c57ca6f5208a29b-s.woff2",revision:"491a7a9678c3cfd4f86c092c68480f23"},{url:"/_next/static/media/37b0c0a51409261e-s.woff2",revision:"5ce748f413aee42a8d4723df0d18830b"},{url:"/_next/static/media/3dbd163d3bb09d47-s.woff2",revision:"93dcb0c222437699e9dd591d8b5a6b85"},{url:"/_next/static/media/42d52f46a26971a3-s.woff2",revision:"b44d0dd122f9146504d444f290252d88"},{url:"/_next/static/media/44c3f6d12248be7f-s.woff2",revision:"705e5297b1a92dac3b13b2705b7156a7"},{url:"/_next/static/media/46c894be853ec49f-s.woff2",revision:"47891b6adb3a947dd3c594bd5196850e"},{url:"/_next/static/media/4a8324e71b197806-s.woff2",revision:"5fba57b10417c946c556545c9f348bbd"},{url:"/_next/static/media/506bd11311670951-s.woff2",revision:"7976a92314c8770252603e7813da9f67"},{url:"/_next/static/media/5647e4c23315a2d2-s.woff2",revision:"e64969a373d0acf2586d1fd4224abb90"},{url:"/_next/static/media/627622453ef56b0d-s.p.woff2",revision:"e7df3d0942815909add8f9d0c40d00d9"},{url:"/_next/static/media/71ba03c5176fbd9c-s.woff2",revision:"2effa1fe2d0dff3e7b8c35ee120e0d05"},{url:"/_next/static/media/7be645d133f3ee22-s.woff2",revision:"3ba6fb27a0ea92c2f1513add6dbddf37"},{url:"/_next/static/media/7c53f7419436e04b-s.woff2",revision:"fd4ff709e3581e3f62e40e90260a1ad7"},{url:"/_next/static/media/7d8c9b0ca4a64a5a-s.p.woff2",revision:"0772a436bbaaaf4381e9d87bab168217"},{url:"/_next/static/media/80a2a8cc25a3c264-s.woff2",revision:"2d3d8a78ef164ab6c1c62a3e57c2727b"},{url:"/_next/static/media/83e4d81063b4b659-s.woff2",revision:"bd30db6b297b76f3a3a76f8d8ec5aac9"},{url:"/_next/static/media/8db47a8bf03b7d2f-s.p.woff2",revision:"49003e0ff09f1efb8323cf35b836ba8f"},{url:"/_next/static/media/8fb72f69fba4e3d2-s.woff2",revision:"7a2e2eae214e49b4333030f789100720"},{url:"/_next/static/media/912a9cfe43c928d9-s.woff2",revision:"376ffe2ca0b038d08d5e582ec13a310f"},{url:"/_next/static/media/934c4b7cb736f2a3-s.p.woff2",revision:"1f6d3cf6d38f25d83d95f5a800b8cac3"},{url:"/_next/static/media/94300924a0693016-s.woff2",revision:"105927314bd3f089b99c0dda456171ed"},{url:"/_next/static/media/9e48537b1b020091-s.woff2",revision:"4b52fd954ca934c204d73ddbc640e5d4"},{url:"/_next/static/media/a5b77b63ef20339c-s.woff2",revision:"96e992d510ed36aa573ab75df8698b42"},{url:"/_next/static/media/a6d330d7873e7320-s.woff2",revision:"f7ec4e2d6c9f82076c56a871d1d23a2d"},{url:"/_next/static/media/baf12dd90520ae41-s.woff2",revision:"8096f9b1a15c26638179b6c9499ff260"},{url:"/_next/static/media/bbdb6f0234009aba-s.woff2",revision:"5756151c819325914806c6be65088b13"},{url:"/_next/static/media/bd976642b4f7fd99-s.woff2",revision:"cc0ffafe16e997fe75c32c5c6837e781"},{url:"/_next/static/media/cff529cd86cc0276-s.woff2",revision:"c2b2c28b98016afb2cb7e029c23f1f9f"},{url:"/_next/static/media/d117eea74e01de14-s.woff2",revision:"4d1e5298f2c7e19ba39a6ac8d88e91bd"},{url:"/_next/static/media/de9eb3a9f0fa9e10-s.woff2",revision:"7155c037c22abdc74e4e6be351c0593c"},{url:"/_next/static/media/dfa8b99978df7bbc-s.woff2",revision:"7a500aa24dccfcf0cc60f781072614f5"},{url:"/_next/static/media/e25729ca87cc7df9-s.woff2",revision:"9a74bbc5f0d651f8f5b6df4fb3c5c755"},{url:"/_next/static/media/eb52b768f62eeeb4-s.woff2",revision:"90687dc5a4b6b6271c9f1c1d4986ca10"},{url:"/_next/static/media/f06116e890b3dadb-s.woff2",revision:"2855f7c90916c37fe4e6bd36205a26a8"},{url:"/background.png",revision:"120f26099ae099dc7f4274a961ca8be1"},{url:"/background1.png",revision:"2e0d3684cfa51bcea9f7362e7dd4bb44"},{url:"/bg.png",revision:"d74693097a94531633eea8f052e057b6"},{url:"/konan-hero.png",revision:"6d1bb98517bea08bbd5720a339ada80f"},{url:"/konan-hero.webp",revision:"e952c2171f6bc972b5cd7ff11f9fbd5d"},{url:"/konan-nav-logo.webp",revision:"a7d84dc364fabbac8f39905e059db4b0"},{url:"/konan_logo.webp",revision:"b15b23cae740c8a339c64d397bd70d89"},{url:"/main1.png",revision:"68ff0d948d1742b3caec56ebff07b5c5"},{url:"/manifest.json",revision:"5e774ea1a43f1afc1a64fb70b230df02"},{url:"/sunglasses.png",revision:"7a2a56c457d8b09411a8a823b36d9a1c"},{url:"/swe-worker-5c72df51bb1f6ee0.js",revision:"5a47d90db13bb1309b25bdf7b363570e"}],{ignoreURLParametersMatching:[/^utm_/,/^fbclid$/]}),e.cleanupOutdatedCaches(),e.registerRoute("/",new e.NetworkFirst({cacheName:"start-url",plugins:[{cacheWillUpdate:async({response:e})=>e&&"opaqueredirect"===e.type?new Response(e.body,{status:200,statusText:"OK",headers:e.headers}):e}]}),"GET"),e.registerRoute(/^https:\/\/fonts\.(?:gstatic)\.com\/.*/i,new e.CacheFirst({cacheName:"google-fonts-webfonts",plugins:[new e.ExpirationPlugin({maxEntries:4,maxAgeSeconds:31536e3})]}),"GET"),e.registerRoute(/^https:\/\/fonts\.(?:googleapis)\.com\/.*/i,new e.StaleWhileRevalidate({cacheName:"google-fonts-stylesheets",plugins:[new e.ExpirationPlugin({maxEntries:4,maxAgeSeconds:604800})]}),"GET"),e.registerRoute(/\.(?:eot|otf|ttc|ttf|woff|woff2|font.css)$/i,new e.StaleWhileRevalidate({cacheName:"static-font-assets",plugins:[new e.ExpirationPlugin({maxEntries:4,maxAgeSeconds:604800})]}),"GET"),e.registerRoute(/\.(?:jpg|jpeg|gif|png|svg|ico|webp)$/i,new e.StaleWhileRevalidate({cacheName:"static-image-assets",plugins:[new e.ExpirationPlugin({maxEntries:64,maxAgeSeconds:2592e3})]}),"GET"),e.registerRoute(/\/_next\/static.+\.js$/i,new e.CacheFirst({cacheName:"next-static-js-assets",plugins:[new e.ExpirationPlugin({maxEntries:64,maxAgeSeconds:86400})]}),"GET"),e.registerRoute(/\/_next\/image\?url=.+$/i,new e.StaleWhileRevalidate({cacheName:"next-image",plugins:[new e.ExpirationPlugin({maxEntries:64,maxAgeSeconds:86400})]}),"GET"),e.registerRoute(/\.(?:mp3|wav|ogg)$/i,new e.CacheFirst({cacheName:"static-audio-assets",plugins:[new e.RangeRequestsPlugin,new e.ExpirationPlugin({maxEntries:32,maxAgeSeconds:86400})]}),"GET"),e.registerRoute(/\.(?:mp4|webm)$/i,new e.CacheFirst({cacheName:"static-video-assets",plugins:[new e.RangeRequestsPlugin,new e.ExpirationPlugin({maxEntries:32,maxAgeSeconds:86400})]}),"GET"),e.registerRoute(/\.(?:js)$/i,new e.StaleWhileRevalidate({cacheName:"static-js-assets",plugins:[new e.ExpirationPlugin({maxEntries:48,maxAgeSeconds:86400})]}),"GET"),e.registerRoute(/\.(?:css|less)$/i,new e.StaleWhileRevalidate({cacheName:"static-style-assets",plugins:[new e.ExpirationPlugin({maxEntries:32,maxAgeSeconds:86400})]}),"GET"),e.registerRoute(/\/_next\/data\/.+\/.+\.json$/i,new e.StaleWhileRevalidate({cacheName:"next-data",plugins:[new e.ExpirationPlugin({maxEntries:32,maxAgeSeconds:86400})]}),"GET"),e.registerRoute(/\.(?:json|xml|csv)$/i,new e.NetworkFirst({cacheName:"static-data-assets",plugins:[new e.ExpirationPlugin({maxEntries:32,maxAgeSeconds:86400})]}),"GET"),e.registerRoute((({sameOrigin:e,url:{pathname:a}})=>!(!e||a.startsWith("/api/auth/callback")||!a.startsWith("/api/"))),new e.NetworkFirst({cacheName:"apis",networkTimeoutSeconds:10,plugins:[new e.ExpirationPlugin({maxEntries:16,maxAgeSeconds:86400})]}),"GET"),e.registerRoute((({request:e,url:{pathname:a},sameOrigin:s})=>"1"===e.headers.get("RSC")&&"1"===e.headers.get("Next-Router-Prefetch")&&s&&!a.startsWith("/api/")),new e.NetworkFirst({cacheName:"pages-rsc-prefetch",plugins:[new e.ExpirationPlugin({maxEntries:32,maxAgeSeconds:86400})]}),"GET"),e.registerRoute((({request:e,url:{pathname:a},sameOrigin:s})=>"1"===e.headers.get("RSC")&&s&&!a.startsWith("/api/")),new e.NetworkFirst({cacheName:"pages-rsc",plugins:[new e.ExpirationPlugin({maxEntries:32,maxAgeSeconds:86400})]}),"GET"),e.registerRoute((({url:{pathname:e},sameOrigin:a})=>a&&!e.startsWith("/api/")),new e.NetworkFirst({cacheName:"pages",plugins:[new e.ExpirationPlugin({maxEntries:32,maxAgeSeconds:86400})]}),"GET"),e.registerRoute((({sameOrigin:e})=>!e),new e.NetworkFirst({cacheName:"cross-origin",networkTimeoutSeconds:10,plugins:[new e.ExpirationPlugin({maxEntries:32,maxAgeSeconds:3600})]}),"GET"),self.__WB_DISABLE_DEV_LOGS=!0}));

