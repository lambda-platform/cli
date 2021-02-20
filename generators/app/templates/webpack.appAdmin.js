let mix = require("laravel-mix");
const path = require("path");
require("laravel-mix-merge-manifest");
const lambdaRoot = "../../vue";

mix.webpackConfig({
    output: {
        chunkFilename: mix.inProduction()
            ? "js/chunks/[name].[chunkhash].js"
            : "js/chunks/[name].js"
    },

    resolve: {
        modules: [
            path.resolve(`${lambdaRoot}/appAdmin/`, 'node_modules'),
        ],
        alias: {
            modules: path.resolve(__dirname, "./node_modules"),

            vue$: "vue/dist/vue.common.js"
        },
        extensions: ["*", ".js", ".ts", ".vue", ".json"],
        symlinks: false
    }
});

mix.options({
    processCssUrls: false
}).setPublicPath("public");

console.log("Compiling - client only app");
mix.js(lambdaRoot+"/appAdmin/index.js", "public/assets/appAdmin/js/appAdmin.js").vue();
mix.sass(lambdaRoot+"/appAdmin/scss/style.scss", "public/assets/appAdmin/css/appAdmin.css");

//mix.styles(
//    [
//        `${lambdaRoot}/template/node_modules/` + "iview/dist/styles/iview.css",
//    ],
//    "public/assets/appAdmin/css/vendor.css"
//);
//mix.extract(
//    [
//        "axios",
//        "vue",
//        "iview",
//        "lodash",
//        "moment",
//        "numeral",
//        "portal-vue",
//        "vue-axios",
//        "vuex",
//        "vue-router"
//    ],
//    "public/assets/appAdmin/js/vendor.js"
//);

if (mix.inProduction()) {
    mix.version();
}

mix.mergeManifest();
