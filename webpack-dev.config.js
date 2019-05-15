const path = require('path')
const join = path.join
const resolve = path.resolve
const {
    camelCase
} = require('lodash')
const webpack = require('webpack')
const {
    // TsConfigPathsPlugin,
    CheckerPlugin
} = require('awesome-typescript-loader')
const HtmlWebpackPlugin = require('html-webpack-plugin')
// const WebpackShellPlugin = require('webpack-shell-plugin');
const env = process && process.env && process.env.NODE_ENV
const serverPort = process.env.npm_package_config_devPort || 8082
const dev = !(env && env === 'production')
const pkgInfo = require('./package.json')
/**
 * Update this variable if you change your library name
 */
const libraryName = 'myGo'
const plugins = [
    // new WebpackShellPlugin({
    //     onBuildStart: ['sh bin/prepare_hooks.sh']
    // }),
    new CheckerPlugin(),
    // new TsConfigPathsPlugin(),
    new HtmlWebpackPlugin({
        inject: 'head', //true 默认值，script标签位于html文件的 body 底部
        title: libraryName,
        filename: 'index.html',
        template: join(__dirname, './go-webassembly/fib/index.html'),
        hash: true,
        chunks: ['common', 'index']
    }),
    new webpack.DefinePlugin({
        'PACKAGE_NAME': JSON.stringify(pkgInfo.name),
        'VERSION': JSON.stringify(pkgInfo.version),
        'BUILD_TIME': JSON.stringify(new Date())
    })
]
let entry = [
    // 'react-hot-loader/patch',
    `webpack-dev-server/client?http://localhost:${serverPort}`,
    // bundle the client for webpack-dev-servers and connect to the provided endpoint
    'webpack/hot/only-dev-server',
    // bundle the client for hot reloading
    // `./src/${libraryName}.ts`
    `./go-webassembly/fib/index.ts`
]
if (dev === false) {

} else {
    plugins.push(new webpack.HotModuleReplacementPlugin())
}
// plugins.push(new TsConfigPathsPlugin({
//     configFile: "examples-map/tsconfig.json"
// }))
module.exports = {
    mode: 'development',
    entry: {
        index: entry
    },
    // Currently cheap-module-source-map is broken https://github.com/webpack/webpack/issues/4176
    devtool: 'source-map',
    output: {
        path: join(__dirname, 'dist'),
        libraryTarget: 'umd',
        library: camelCase(libraryName),
        filename: `${libraryName}.js`
    },
    resolve: {
        extensions: ['.ts', '.js']
    },
    module: {
        unknownContextCritical : false,
        rules: [{
                test: /\.ts$/,
                use: [{
                    loader: 'ts-loader'
                }]
            },
            {
                test: /\.scss$/,
                use: ['style-loader', 'css-loader', 'sass-loader']
            },
            {
                test: /\._exec\.js$/,
                use: ['script-loader']
            }
        ]
    },
    plugins: plugins,
    devServer: {
        hot: true,
        contentBase: resolve(__dirname, 'go-webassembly/fib'),
        watchContentBase: true,
        port: serverPort,
        publicPath: '/'
    },
    node: {
        fs: 'empty'
    }
}