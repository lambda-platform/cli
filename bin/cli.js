#!/usr/bin/env node
const LambdaCLI = require('./index');
const cli = new LambdaCLI(process.argv.slice(2));
cli.run();