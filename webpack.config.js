// Based on https://github.com/preboot/angular-webpack/blob/master/webpack.config.js

// Helper: root() is defined at the bottom
var path = require('path');
var webpack = require('webpack');
var CopyWebpackPlugin = require('copy-webpack-plugin');
var WebpackCleanupPlugin = require('webpack-cleanup-plugin');

var ENV = process.env.ENV;
var isProd = (ENV === 'prod');
var envText = isProd ? 'prod' : 'dev'

console.debug('Resolved build environment: '  + envText);

// Webpack Plugins
var CommonsChunkPlugin = webpack.optimize.CommonsChunkPlugin;
var autoprefixer = require('autoprefixer');
var HtmlWebpackPlugin = require('html-webpack-plugin');

module.exports = () => {
  var config = {};

  if (!isProd) {
    config.devtool = 'eval-source-map';
  } else {
    config.devtool = 'source-map';
  }

  config.entry = {
    app: root('pubg-hack-ui', 'main.ts'),
    polyfill: root('pubg-hack-ui', 'polyfill.ts'),
    vendor: root('pubg-hack-ui', 'vendor.ts'),
  };

  config.output = {
    path: root('dist'),
    filename: '[name].[hash].'+envText+'.bundle.js',
  };

  config.resolve = {
    extensions: ['.ts', '.js', '.json', '.css', '.scss', '.html'],
    alias: {
      'fonts': root('pubg-hack-ui', 'fonts'),
      'jquery': 'jquery/dist/jquery.min',
    },
    modules: [
      root('.'),
      'node_modules',
    ]
  };

  var atlConfigFile = root('pubg-hack-ui', 'tsconfig.json');
  config.module = {
    rules: [
      {test: /\.ts$/, loader: 'awesome-typescript-loader?configFileName=' + atlConfigFile},
      {test: /\.(gif|png|jpg|svg|woff|woff2|ttf|eot)$/, loader: 'file-loader'},
      {test: /\.json$/, loader: 'json-loader'},
      {test: /\.(css|scss|sass)$/, loaders: ['to-string-loader', 'css-loader', 'sass-loader']},
      {test: /\.html$/, loader: 'html-loader'},
    ]
  };

  config.plugins = [
    new CommonsChunkPlugin({
      name: ['vendor', 'polyfill'],
    }),
    new webpack.ProvidePlugin({
      jQuery: 'jquery',
      $: 'jquery',
      jquery: 'jquery',
    }),
    new webpack.LoaderOptionsPlugin({
      options: {
        postcss: [
          require('autoprefixer')(),
        ]
      }
    }),
    new webpack.DefinePlugin({
      'process.env': {
        ENV: JSON.stringify(envText),
      }
    }),
    new webpack.NoEmitOnErrorsPlugin(),
    new HtmlWebpackPlugin({
      template: root('pubg-hack-ui', 'index.html'),
      chunksSortMode: 'dependency',
    }),
    new CopyWebpackPlugin([{
      from: root('pubg-hack-ui', 'coherent.js'),
    }]),
    new WebpackCleanupPlugin(),
  ];

  if (isProd) {
    config.plugins.push(
      new webpack.optimize.UglifyJsPlugin({sourceMap: true, mangle: { keep_fnames: true }})
    );
  }

  config.devServer = {
    contentBase: root('pubg-hack-ui'),
    historyApiFallback: true,
    quiet: false,
    stats: 'normal', // none (or false), errors-only, minimal, normal (or true) and verbose
    port: process.env.WEBPACK_DEV_SERVER_PORT || 8888,
  };

  return config
}


// Helper functions
function root(args) {
  args = Array.prototype.slice.call(arguments, 0);
  return path.join.apply(path, [__dirname].concat(args));
}
