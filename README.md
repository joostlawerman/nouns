[![Build Status](https://travis-ci.org/joostlawerman/nouns.svg?branch=master)](https://travis-ci.org/joostlawerman/nouns)
[![Coverage Status](https://coveralls.io/repos/github/joostlawerman/nouns/badge.svg?branch=master)](https://coveralls.io/github/joostlawerman/nouns?branch=master)
# Nouns
A simple plugin to Pluralize and Singularize nouns.

# Install
```
	go get github.com/joostlawerman/nouns
```
# Examples
```
	plural, err := nouns.Pluralize("noun") // nouns
	singular, err := nouns.Singularize("nouns") // noun
```
# Contributions
English is not my first language so if you find anything missing or a conversion not working please consider to open an issue or send me a pull request.

# License
MIT License

Copyright (c) 2016 Joost Lawerman

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
