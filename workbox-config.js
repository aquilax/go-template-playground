module.exports = {
  globDirectory: "public/",
  globPatterns: ["**/*.{ico,html,js,svg,png,json,wasm}"],
  swDest: "public/sw.js",
  ignoreURLParametersMatching: [/^utm_/, /^fbclid$/],
  maximumFileSizeToCacheInBytes: 10 * 1024 * 1024,
};
