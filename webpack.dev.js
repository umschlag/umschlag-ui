const argv = require('yargs').argv;
const webpack = require('webpack');
const webpackConfig = require('./webpack.config.js');
const WebpackDevServer = require('webpack-dev-server');

const port = argv.port || 9000;
const scheme = argv.scheme || 'http';
const host = argv.host || 'localhost:8080';
const token = argv.token || null;
const url = `${scheme}://${host}`

webpackConfig.entry.app.unshift(
  `webpack-dev-server/client?http://localhost:${port}/`,
  'webpack/hot/dev-server'
);

webpackConfig.plugins.unshift(
  new webpack.HotModuleReplacementPlugin()
);

const compiler = webpack(
  webpackConfig
);

if (token == null) {
  var params = {
    target: url,
    xfwd: true,
    changeOrigin: true
  };
} else {
  var params = {
    target: url,
    headers: {
      'Authorization': `Bearer ${token}`
    },
    xfwd: true,
    changeOrigin: true
  };
}

const server = new WebpackDevServer(compiler, {
  hot: true,
  inline: true,
  historyApiFallback: true,
  proxy: {
    '/api/*': params
  },
  stats: {
    assets: false,
    colors: true,
    version: false,
    modules: false,
    hash: false,
    timings: false,
    chunks: false,
    chunkModules: false,
    reasons: false,
    cached: true,
    chunkOrigins: true,
    children: false
  }
});

server.listen(port, (err) => {
  if (err) {
    return console.err(err);
  }

  console.log(`Now listening on http://localhost:${port}`);
});
