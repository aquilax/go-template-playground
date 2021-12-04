module.exports = {
  globDirectory: "public/",
  globPatterns: ["**/*.{ico,html,js,svg,png,json}"],
  swDest: "public/sw.js",
  ignoreURLParametersMatching: [/^utm_/, /^fbclid$/],
};
