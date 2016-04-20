const ExtractTextPlugin = require('extract-text-webpack-plugin');
const path = require('path');

module.exports = {
  entry: [
    path.join(__dirname, '/src/index.js')
  ],

  output: {
    filename: path.join(__dirname, '/dist/harbor.js')
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
      }
    ]
  },

  externals: {
    'immutable': 'Immutable',
    'moment': 'moment',
    'react': 'React',
    'react-dom': 'ReactDOM',
    'react-redux': 'ReactRedux',
    'react-router': 'ReactRouter',
    'redux': 'Redux',
    'superagent': 'superagent'
  },

  resolve: {
    extensions: [
      '',
      '.js',
      '.jsx'
    ]
  },

  plugins: [
    new ExtractTextPlugin(
      path.join(__dirname, '/dist/harbor.css')
    )
  ]
};
