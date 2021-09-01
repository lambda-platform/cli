const { WebpackManifestPlugin } = require('webpack-manifest-plugin');
module.exports = {
    configureWebpack: config => {
        config.plugins = config.plugins.concat(
            new WebpackManifestPlugin({
            })
        );
        config.devServer = {
            headers: { 'Access-Control-Allow-Origin': '*' }
        };
    },
    publicPath: process.env.NODE_ENV === 'production' ? "/assets/app" : '',
    outputDir:"dist/assets/app/",

};
