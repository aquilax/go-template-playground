module.exports = {
  globDirectory: "public/",
  globPatterns: ["**/*.{ico,html,js,svg,png}"],
  swDest: "public/sw.js",
  ignoreURLParametersMatching: [/^utm_/, /^fbclid$/],
};
