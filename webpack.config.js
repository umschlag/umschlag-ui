const ExtractTextPlugin = require('extract-text-webpack-plugin');

module.exports = {
  entry: {
    app: [
      './src/index.js'
    ]
  },

  output: {
    path: './assets',
    filename: 'harbor.js'
  },

  devtool: 'source-map',

  module: {
    loaders: [
      {
        test: /\.js?$/,
        exclude: /node_modules/,
        loader: 'babel',
        query: {
          presets: [
            'react',
            'es2015'
          ]
        }
      },
      {
        test: /\.less$/,
        loader: ExtractTextPlugin.extract(
          'css?sourceMap!' + 'less?sourceMap'
        )
      },
      {test: /\.(woff|woff2)(\?v=\d+\.\d+\.\d+)?$/, loader: 'url?limit=10000&mimetype=application/font-woff'},
      {test: /\.svg(\?v=\d+\.\d+\.\d+)?$/, loader: 'url?limit=10000&mimetype=image/svg+xml'},
      {test: /\.ttf(\?v=\d+\.\d+\.\d+)?$/, loader: 'url?limit=10000&mimetype=application/octet-stream'},
      {test: /\.eot(\?v=\d+\.\d+\.\d+)?$/, loader: 'file'}
    ]
  },

  // externals: {
  //   'immutable': 'Immutable',
  //   'moment': 'moment',
  //   'react': 'React',
  //   'react-dom': 'ReactDOM',
  //   'react-redux': 'ReactRedux',
  //   'react-router': 'ReactRouter',
  //   'redux': 'Redux',
  //   'superagent': 'superagent'
  // },

  resolve: {
    extensions: [
      '',
      '.js',
      '.jsx'
    ]
  },

  plugins: [
    new ExtractTextPlugin(
      'harbor.css'
    )
  ]
};
