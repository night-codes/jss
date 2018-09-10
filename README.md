# JSS
CommonJS and AMD modules compiler written in Golang. 

## How To Install   
```bash
go install github.com/night-codes/jss
```

## How to use

```bash
Usage: jss [--watch WATCH] [--import IMPORT] INPUTDIR OUTPUT

Positional arguments:
  INPUTDIR               Input source directory ("jss/").
  OUTPUT                 Output JS file ("script.js").

Options:
  --watch WATCH, -w WATCH
                         Watch directory for changes.
  --import IMPORT, -i IMPORT
                         Specify a Scss import path.
  --help, -h             display this help and exit
```


## License

MIT License

Copyright (c) 2018 Oleksiy Chechel (alex.mirrr@gmail.com)   

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
