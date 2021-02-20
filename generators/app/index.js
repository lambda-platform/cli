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

        let prompts = [{
            type: 'input',
            name: 'useAppAdmin',
            message: 'are you want example "app admin"?',
            default: 'yes'
        }];

        if(this.options["projectName"]){
            this.projectName = this.options["projectName"][0].replace(/\s+/g, '-').toLowerCase();

            this.pro = this.options["pro"];
            // this.serviceName = this.options["projectName"][0].replace(/\s+/g, '-').toLowerCase();
            this.serviceName = "app";

            return this.prompt(prompts).then(props => {

                this.useAppAdmin = props.useAppAdmin.replace(/\s+/g, '-').toLowerCase();
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
            this.pro = props.pro.replace(/\s+/g, '-').toLowerCase();
            // this.serviceName = props.appName.replace(/\s+/g, '-').toLowerCase();
            this.serviceName = "app";
            this.useAppAdmin = props.useAppAdmin.replace(/\s+/g, '-').toLowerCase();
            this.secretKey = getKey(34)

            cb()
        });

    }

    writing() {
        console.log("Creating project folders");

        let srcDir = this.destinationPath(path.join('', this.projectName));
        let serviceDir = this.destinationPath(path.join(this.projectName, this.serviceName));
        let appAdminDir = this.destinationPath(path.join(this.projectName, "appAdmin"));
        let modelsDir = this.destinationPath(path.join(this.projectName, "models"));
        let projectPublicDir = srcDir+"/public/";
        let initDir = srcDir+"/init";



        mkdir.sync(srcDir);
        mkdir.sync(projectPublicDir);
        mkdir.sync(serviceDir);
        mkdir.sync(initDir);
        if(this.useAppAdmin == "yes"){
            mkdir.sync(appAdminDir);
        }
        mkdir.sync(modelsDir);


        this.fs.copy(
            this.templatePath('_gitignore'),
            path.join(srcDir, '.gitignore')
        );
        this.fs.copy(
            this.templatePath('README.md'),
            path.join(srcDir, 'README.md')
        );



        // this.fs.copy(
        //     this.templatePath('init'),
        //     path.join(srcDir, 'init')
        // );
        this.fs.copy(
            this.templatePath('server_config'),
            path.join(srcDir, 'server_config')
        );
        this.fs.copy(
            this.templatePath('public'),
            path.join(srcDir, 'public')
        );




        let tmplContext = {
            projectName: this.projectName,
            serviceName: this.serviceName,
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
        if(this.pro == "yes"){
            if(this.useAppAdmin == "yes"){
                this.fs.copyTpl(
                    this.templatePath('main.go_pro'),
                    path.join(srcDir, 'main.go'),
                    tmplContext
                );
            } else {
                this.fs.copyTpl(
                    this.templatePath('main.go_pro_without_admin'),
                    path.join(srcDir, 'main.go'),
                    tmplContext
                );
            }
        } else {
            if(this.useAppAdmin == "yes"){
                this.fs.copyTpl(
                    this.templatePath('main.go'),
                    path.join(srcDir, 'main.go'),
                    tmplContext
                );
            } else {
                this.fs.copyTpl(
                    this.templatePath('main_witihout_admin.go'),
                    path.join(srcDir, 'main.go'),
                    tmplContext
                );
            }
        }



        this.fs.copyTpl(
            this.templatePath('webpack.appAdmin.js'),
            path.join(srcDir, 'webpack.appAdmin.js'),
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
        } else {
            this.fs.copyTpl(
                this.templatePath('go.mod'),
                path.join(srcDir, 'go.mod'),
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
                this.templatePath('init/init.go_pro'),
                path.join(initDir, 'init.go'),
                tmplContext
            );
            this.fs.copyTpl(
                this.templatePath('start.sh_pro'),
                path.join(srcDir, 'start.sh'),
                tmplContext
            );
        } else {
            this.fs.copyTpl(
                this.templatePath('init/init.go'),
                path.join(initDir, 'init.go'),
                tmplContext
            );
            this.fs.copyTpl(
                this.templatePath('start.sh'),
                path.join(srcDir, 'start.sh'),
                tmplContext
            );
        }

        this.fs.copyTpl(
            this.templatePath('exampleService/app.go'),
            path.join(serviceDir, 'app.go'),
            tmplContext
        );

        this.fs.copyTpl(
            this.templatePath('exampleService/handlers'),
            path.join(serviceDir, 'handlers'),
            tmplContext
        );

        this.fs.copyTpl(
            this.templatePath('exampleService/middlewares'),
            path.join(serviceDir, 'middlewares'),
            tmplContext
        );
        if(this.useAppAdmin == "yes") {
            this.fs.copyTpl(
                this.templatePath('appAdmin/admin.go'),
                path.join(appAdminDir, 'admin.go'),
                tmplContext
            );
            this.fs.copyTpl(
                this.templatePath('appAdmin/handlers'),
                path.join(appAdminDir, 'handlers'),
                tmplContext
            );
            this.fs.copyTpl(
                this.templatePath('appAdmin/templates'),
                path.join(appAdminDir, 'templates'),
                tmplContext
            );
        }

        this.fs.copyTpl(
            this.templatePath('models/form'),
            path.join(modelsDir, 'form'),
            tmplContext
        );
        this.fs.copyTpl(
            this.templatePath('models/grid'),
            path.join(modelsDir, 'grid'),
            tmplContext
        );
        this.fs.copyTpl(
            this.templatePath('models/form/caller'),
            path.join(modelsDir, 'form/caller'),
            tmplContext
        );

        this.fs.copyTpl(
            this.templatePath('models/form/validationCaller'),
            path.join(modelsDir, 'form/validationCaller'),
            tmplContext
        );
        this.fs.copyTpl(
            this.templatePath('models/form/validations'),
            path.join(modelsDir, 'form/validations'),
            tmplContext
        );
        this.fs.copyTpl(
            this.templatePath('models/grid/caller'),
            path.join(modelsDir, 'grid/caller'),
            tmplContext
        );
        this.fs.copyTpl(
            this.templatePath('exampleService/models'),
            path.join(serviceDir, 'models'),
            tmplContext
        );
        this.fs.copyTpl(
            this.templatePath('exampleService/routes'),
            path.join(serviceDir, 'routes'),
            tmplContext
        );
        this.fs.copyTpl(
            this.templatePath('exampleService/templates'),
            path.join(serviceDir, 'templates'),
            tmplContext
        );


    }
};

module.exports = LambdaGo;