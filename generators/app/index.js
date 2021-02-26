'use strict';
const path = require('path');
const Generator = require('yeoman-generator');
const mkdir = require('mkdirp');
function getKey(length) {
    var randomChars = 'ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789';
    var result = '';
    for ( var i = 0; i < length; i++ ) {
        result += randomChars.charAt(Math.floor(Math.random() * randomChars.length));
    }
    return result;
}
class LambdaGo extends Generator {

    paths() {
        this.destinationRoot('./');
    }

    prompting() {

        let cb = this.async();

        let prompts = [
            // {
            //     type: 'input',
            //     name: 'useAppAdmin',
            //     message: 'are you want example "app admin"?',
            //     default: 'yes'
            // }
        ];

        if(this.options["projectName"]){
            this.projectName = this.options["projectName"][0].replace(/\s+/g, '-').toLowerCase();
            this.pro = this.options["pro"];

            return this.prompt(prompts).then(props => {

                this.secretKey = getKey(34)

                cb()
            });
        } else {
            prompts = [{
                type: 'input',
                name: 'projectName',
                message: 'Please insert Project?',
                default: 'example'
            }, prompts[0]]
        }
        return this.prompt(prompts).then(props => {
            this.projectName = props.projectName.replace(/\s+/g, '-').toLowerCase();
            this.secretKey = getKey(34)

            cb()
        });

    }

    writing() {
        console.log("Creating project folders");

        let srcDir = this.destinationPath(path.join('', this.projectName));

        let projectPublicDir = srcDir+"/public/";

        mkdir.sync(srcDir);
        mkdir.sync(projectPublicDir);



        this.fs.copy(
            this.templatePath('_gitignore'),
            path.join(srcDir, '.gitignore')
        );
        this.fs.copy(
            this.templatePath('README.md'),
            path.join(srcDir, 'README.md')
        );

        this.fs.copy(
            this.templatePath('server_config'),
            path.join(srcDir, 'server_config')
        );
        this.fs.copy(
            this.templatePath('public'),
            path.join(srcDir, 'public')
        );
        this.fs.copy(
            this.templatePath('assets'),
            path.join(srcDir, 'assets')
        );
        this.fs.copy(
            this.templatePath('database'),
            path.join(srcDir, 'database')
        );

        let tmplContext = {
            projectName: this.projectName,
            secretKey: this.secretKey,
        };
        this.fs.copyTpl(
            this.templatePath('lambda.json'),
            path.join(srcDir, 'lambda.json'),
        );

        this.fs.copyTpl(
            this.templatePath('__env'),
            path.join(srcDir, '.env'),
            tmplContext
        );

        this.fs.copyTpl(
            this.templatePath('_editorconfig'),
            path.join(srcDir, '.editorconfig'),
        );
        this.fs.copyTpl(
            this.templatePath('main.go'),
            path.join(srcDir, 'main.go'),
            tmplContext
        );



        this.fs.copyTpl(
            this.templatePath('webpack.app.js'),
            path.join(srcDir, 'webpack.app.js'),
        );

        this.fs.copyTpl(
            this.templatePath('webpack.lambda.js'),
            path.join(srcDir, 'webpack.lambda.js'),
        );

        this.fs.copyTpl(
            this.templatePath('webpack.mix.js'),
            path.join(srcDir, 'webpack.mix.js'),
        );

        this.fs.copyTpl(
            this.templatePath('package.json'),
            path.join(srcDir, 'package.json'),
        );

        if(this.pro == "yes"){
            this.fs.copyTpl(
                this.templatePath('go.mod_pro'),
                path.join(srcDir, 'go.mod'),
                tmplContext
            );
            this.fs.copyTpl(
                this.templatePath('start.sh_pro'),
                path.join(srcDir, 'start.sh'),
                tmplContext
            );
        } else {
            this.fs.copyTpl(
                this.templatePath('go.mod'),
                path.join(srcDir, 'go.mod'),
                tmplContext
            );
            this.fs.copyTpl(
                this.templatePath('start.sh'),
                path.join(srcDir, 'start.sh'),
                tmplContext
            );
        }

        this.fs.copyTpl(
            this.templatePath('runner.conf'),
            path.join(srcDir, 'runner.conf'),
            tmplContext
        );

        if(this.pro == "yes") {

            this.fs.copyTpl(
                this.templatePath('bootstrap_pro'),
                path.join(srcDir, 'bootstrap'),
                tmplContext
            );
        } else {
            this.fs.copyTpl(
                this.templatePath('bootstrap'),
                path.join(srcDir, 'bootstrap'),
                tmplContext
            );
        }
        this.fs.copyTpl(
            this.templatePath('routes'),
            path.join(srcDir, 'routes'),
            tmplContext
        );
        this.fs.copyTpl(
            this.templatePath('views'),
            path.join(srcDir, 'views'),
            tmplContext
        );
        this.fs.copyTpl(
            this.templatePath('app'),
            path.join(srcDir, 'app'),
            tmplContext
        );




    }
};

module.exports = LambdaGo;