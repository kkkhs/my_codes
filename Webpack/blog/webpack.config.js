const path = require('path')
const HtmlWebpackPlugin = require('html-webpack-plugin')
const TerserPlugin = require("terser-webpack-plugin")
const BundleAnalyzerPlugin = require("webpack-bundle-analyzer").BundleAnalyzerPlugin

module.exports = {
  mode: 'development',
  devtool:"inline-source-map",
  entry: {
    index: './src/index.js',
    other: './src/other.js',
  },
  output: {
    filename: '[name].[contentHash:8].js',    //只有文件改变，打包的hash才会改变，更好利用浏览器缓存
    path: path.resolve(__dirname, 'dist')   // 当前目录下的dist路径
  },
  // 配置路径别名(在引入路径时直接使用 utils/...)
  resolve: {  
    alias: {
      utils: path.resolve(__dirname, 'src/utils')
    }
  },
  optimization:{
    minimize: true,   // 是否需要压缩
    minimizer: [new TerserPlugin()]   //压缩使用的插件
  },
  // 本地服务插件
  devServer: {
    static: "./dist", //  指定运行目录
  },
  plugins:[
    new HtmlWebpackPlugin({  // 自动引入js文件
      template: './src/index.html',
      filename: 'index.html',
      chunks: ['index'] //只引用 index.js
    }),
    new HtmlWebpackPlugin({  // 自动引入js文件
      template: './src/other.html',
      filename: 'other.html',
      chunks: ['other']  //只引用 other.js
    }),
    new BundleAnalyzerPlugin()  // 可视化打包工具
  ],
  module: {
    rules: [
      // 配置css-loader(需要yarn add)
      {
        test: /\.css$/i,    // 匹配文件名
        use: ["style-loader", "css-loader"] // 使用的外部下载的loader
      },
      // 配置图片loader(内置)
      {
        test: /\.(png|svg|jpg|jpeg|gif)$/i,
        type: "asset/resource"    //使用的内置的loader
      },
      {
        test: /\.js$/,
        exclude: /node_modules/,   // 排除的文件
        use: {
          loader: "babel-loader",
          options: {
            presets: ["@babel/preset-env"]
          }
        }
      }
    ]
  }
}