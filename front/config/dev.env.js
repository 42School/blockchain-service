'use strict'
const merge = require('webpack-merge')
const prodEnv = require('./prod.env')

module.exports = merge(prodEnv, {
  NODE_ENV: '"development"',
  BACK_URL: '"http://localhost:3001"',
  GOOGLE_TOKEN_API: '"token"'
})
