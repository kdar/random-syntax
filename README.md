# random-syntax theme

![random-syntax in action](https://raw.githubusercontent.com/kdar/random-syntax/master/examples/jt64H9QFt3.gif)

This is a syntax theme for [Atom](http://atom.io) that provides a Go tool to randomly generate it.

The code and idea was taken from [HUSL](http://www.husl-colors.org/syntax/).

Base theme files were taken from [Monokai](https://github.com/kevinsawicki/monokai).

## Usage

Just run `go run generate.go` in the random-syntax directory. This will replace the
`styles/colorsgen.less` file. If you generate a theme you like, you should backup
that file so you don't overwrite it.

## TODO

Use a theme that supports base16 colors (such as [github.com/Alchiadus/base16-syntax](github.com/Alchiadus/base16-syntax)).

Instead of using a Go program, make an atom package to generate a theme.

## Contributing

Pull requests to use more of the 16 colors or any other fix are greatly welcome.
