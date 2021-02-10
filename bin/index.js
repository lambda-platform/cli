#!/usr/bin/env node
// const Conf = require('conf');
const env = require('yeoman-environment').createEnv();
const updateNotifier = require('update-notifier');
const chalk = require('chalk');
const minimist = require('minimist');
const pkg = require('../package.json');

class LambdaCLI {
    constructor(args) {
        this._args = args;
        this._options = minimist(args, {
            boolean: ['help', 'version'],
            string: [''],
            alias: {
                h: 'help',
                v: 'version'
            }
        });
        // this._config = new Conf({
        //     projectName: '@lambda-platform/go',
        //     defaults: {
        //         disabledAddons: {}
        //     }
        // });

        env.register(require.resolve(require.resolve('..')), '@lambda-platform/go');

    }

    async run() {

        updateNotifier({pkg}).notify();

        if (this._options.help) {
            return this._help();
        }

        if (this._options.version) {
            return console.log("@lambda-platform/cli: " + pkg.version);
        }

        switch (this._args[0]) {
            case 'c':
            case 'create':
                return this.generate(this._args.slice(1), this._options);
            default:
                return this._help();
        }
    }

    generate(projectName, options) {
        options = options || {};
        console.log(chalk.cyan(`Lambda CLI`))

        env.run('@lambda-platform/go', {"projectName": projectName}, function () {
            console.log("Starter app generated")
        });
    }

    _help() {
        const detailedHelp = `${chalk.cyan(`Lambda CLI   
`)}
Usage: lambda <command> [options]

Options:
  ${chalk.blueBright('-v, --version')}                 output the version number
  ${chalk.blueBright('-h, --help ')}                   output usage information

Commands:
  ${chalk.blueBright('lambda c, create [options] <app-name> ')}  create a new project powered by lambda-cli

Examples:
  ${chalk.blueBright('lambda create my-app, lambda c my-app')}
`;
        console.log(detailedHelp)
    }

}

module.exports = LambdaCLI;