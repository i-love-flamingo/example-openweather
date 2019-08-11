const path = require('path')
// const mkdirp = require('mkdirp')
// const shell = require('shelljs')

class OpenWeather {
  constructor (core) {
    this.listeners = [
      {
        command: 'config',
        handler: function (core) {
          const config = core.getConfig()

          // Pug config
          config.paths.pug.src = path.join(config.paths.src, 'page')
          config.paths.pug.dist = path.join(config.paths.dist, 'template')

          config.pug.filesPattern = '/{*,.,*/page/*}/{*,.,*/*.partial}/*.pug'

          config.pugLint.filesPattern = '/**/*.pug'
        }
      },
      {
        command: 'config',
        priority: 10,
        handler: function (core) {
          // Webpack config - pre config creation

          const config = core.getConfig()

          config.paths.webpack.src = path.join(config.paths.src, 'base')
          config.webpack.rulesInclude = [
            path.posix.resolve(config.paths.src),
            /node_modules[\\/]/
          ]
        }
      },
      {
        command: 'config',
        priority: -10,
        handler: function (core) {
          // Webpack config - post config creation

          const config = core.getConfig()

          config.webpackConfig = {
            ...config.webpackConfig,
            entry: {
              base: [
                path.join(config.paths.webpack.src, 'js', 'index.js'),
                path.join(config.paths.webpack.src, 'style', 'index.sass')
              ]
            }
          }
        }
      }
      // {
      //   command: 'build',
      //   handler: function (core) {
      //     const config = core.getConfig()
      //
      //     // Copy assets
      //     mkdirp(path.join(config.paths.dist), () => {
      //       shell.cp('-r', path.join(config.paths.src, 'asset'), path.join(config.paths.dist, 'asset'))
      //     })
      //   }
      // }
    ]
  }

  getListeners () {
    return this.listeners
  }
}

module.exports = OpenWeather
