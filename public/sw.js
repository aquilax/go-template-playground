if(!self.define){let e,i={};const s=(s,c)=>(s=new URL(s+".js",c).href,i[s]||new Promise((i=>{if("document"in self){const e=document.createElement("script");e.src=s,e.onload=i,document.head.appendChild(e)}else e=s,importScripts(s),i()})).then((()=>{let e=i[s];if(!e)throw new Error(`Module ${s} didn’t register its module`);return e})));self.define=(c,n)=>{const o=e||("document"in self?document.currentScript.src:"")||location.href;if(i[o])return;let r={};const t=e=>s(e,o),f={module:{uri:o},exports:r,require:t};i[o]=Promise.all(c.map((e=>f[e]||t(e)))).then((e=>(n(...e),r)))}}define(["./workbox-f683aea5"],(function(e){"use strict";self.addEventListener("message",(e=>{e.data&&"SKIP_WAITING"===e.data.type&&self.skipWaiting()})),e.precacheAndRoute([{url:"favicon.ico",revision:"b4ad1213c636ba43c7e52628d80fae97"},{url:"images/icons-192.png",revision:"78a2d98c4593b1f501e45744235f2d39"},{url:"images/icons-512.png",revision:"58ba208658ab7cc24fdb17ae5436c5dc"},{url:"images/icons-vector.svg",revision:"d79470f21f2e6cfb8caf8904724b7f51"},{url:"index.html",revision:"507cc2b94f8c81733b3dcd8373f6b642"},{url:"manifest.json",revision:"80d536776add9d0ecb2fca9e4c6e6289"},{url:"tgo.js",revision:"4923119a9cc4e01b53624654dfb8d6fa"}],{ignoreURLParametersMatching:[/^utm_/,/^fbclid$/]})}));
//# sourceMappingURL=sw.js.map
