// Snowpack Configuration File
// See all supported options: https://www.snowpack.dev/reference/configuration

const httpProxy = require('http-proxy');

const proxy = httpProxy.createServer({ target: 'http://localhost:8090' });

/** @type {import("snowpack").SnowpackUserConfig } */
module.exports = {

  // semble non nécessaire
  // "root" : ".",
  
  "mount" : {
    "public" : "/"
  },
  
  //{“mount”: {“src”: “/“}}
  
  routes: [
    // dev trick to force loading tmpl directly
    // {
    //   src: '/*.html',
    //   dest: (req, res) => proxy.web(req, res),
    // },
    {
      src: '/api/.*',
      dest: (req, res) => proxy.web(req, res),
    }
  ],

  plugins: [
    /* ... */
  ],
  packageOptions: {
    /* ... */
  },
  devOptions: {
    open: "chrome"
  },
  buildOptions: {
    /* ... */
  },
};
