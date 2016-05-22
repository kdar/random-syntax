# random-syntax theme

![random-syntax in action](https://raw.githubusercontent.com/kdar/random-syntax/master/examples/jt64H9QFt3.gif)

This is a syntax theme for [Atom](http://atom.io) that provides a Go tool to randomly generate it.

The code and idea was taken from [HUSL](http://www.husl-colors.org/syntax/).

Base theme files were taken from [Monokai](https://github.com/kevinsawicki/monokai).

## Usage

Just run `go run generate.go` in the random-syntax directory. This will replace the
`styles/colorsgen.less` file. If you generate a theme you like, you should backup
that file so you don't overwrite it.

## Notes

I pretty much haphazardly matched the base16 colors generated to other colors already in the theme. It may be worth while to actually making a theme that is designed to take these 16 colors.

## Contributing

Pull requests to use more of the 16 colors or any other fix are greatly welcome.
