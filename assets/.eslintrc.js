// http://eslint.org/docs/user-guide/configuring

module.exports = {

  root: true,

  parser: 'babel-eslint',

  'parserOptions': {
    'ecmaVersion': 6,
    'sourceType': 'module',
    'ecmaFeatures': {
      'experimentalObjectRestSpread': true,
      'impliedStrict': true
    }
  },

  'env': {
    'browser': true,
    'node': true,
    'amd': false,
    'mocha': false,
    'jasmine': false
  },

  // https://github.com/standard/standard/blob/master/docs/RULES-en.md
  extends: 'standard',

  'globals': {
    'window': true,
    '__API_URL__': true,
    '$': true,
    'jQuery': true,
    'angular': true,
    '_': true,
    'lodash': true,
    'moment': true
  },
  'rules': {
    'max-len': ['error', 120],
    'indent': ['error', 2, {'SwitchCase': 1}],
    'semi': ['error', 'always'],
    'semi-spacing': [
      'error',
      {
        'before': false,
        'after': true
      }
    ],
    'quotes': ['error', 'single'],
    'comma-dangle': ['error', 'never'],
    'accessor-pairs': 'off',
    'block-scoped-var': 'error',
    'curly': ['error', 'multi-line'],
    'dot-notation': [
      'error',
      {
        'allowKeywords': true
      }
    ],
    'eqeqeq': ['warn', 'allow-null'],
    'no-alert': 'warn',
    'no-array-constructor': 'warn',
    'no-else-return': 'warn',
    'no-empty-function': [
      'error',
      {
        'allow': [
          'arrowFunctions',
          'functions',
          'methods'
        ]
      }
    ],
    'no-dupe-args': 'error',
    'no-dupe-keys': 'error',
    'no-debugger': 'warn',
    'no-duplicate-case': 'error',
    'no-eq-null': 'warn',
    'no-empty': [
      'error',
      {
        'allowEmptyCatch': true
      }
    ],
    'no-extend-native': 'error',
    'no-extra-bind': 'error',
    'no-extra-boolean-cast': 'error',
    'no-extra-semi': 'warn',
    'no-ex-assign': 'error',
    'no-implicit-globals': 'off',
    'no-implied-eval': 'error',
    'no-irregular-whitespace': 'error',
    'no-proto': 'error',
    'no-redeclare': 'error',
    'no-regex-spaces': 'error',
    'no-self-assign': 'error',
    'no-self-compare': 'error',
    'no-sequences': 'warn',
    'no-sparse-arrays': 'error',
    'no-template-curly-in-string': 'error',
    'no-throw-literal': 'error',
    'no-unexpected-multiline': 'error',
    'no-unmodified-loop-condition': 'off',
    'no-unreachable': 'error',
    'no-useless-escape': 'warn',
    'no-void': 'error',
    'no-with': 'error',
    'radix': 'error',
    'spaced-comment': 'off',
    'no-unused-expressions': [
      'warn',
      {
        'allowShortCircuit': true,
        'allowTernary': false
      }
    ],
    'vars-on-top': 'off',
    'babel/object-curly-spacing': ['warn', 'always'],
    'babel/object-shorthand': 'warn',
    'babel/func-params-comma-dangle': 'warn'
  },
  'plugins': [
    'babel',
    'html',
    'vue'
  ]
};
