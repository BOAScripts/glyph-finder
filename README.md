# What

2 tools:
- `parser.py`: Python scripts that generate a `glyphs.json` with all the glyphs with their name and group from local.
- `nerd-glyphs`: a CLI tool that fuzzzy find against the `glyphs.json` embedded in the binary (`main.go`).

Some helper scripts: // TODO
- `update-glyph-list.sh`: Update the glyphs.json from NerdFont git and run `parser.py`
- `install.sh` - idk how to distribute a program. update then install, or just install.

# Why

The NerdFont website does not have all the glyphs from all their fonts for whatever valable reason (license I guess). This project aims to be able to search for specific glyphs easily within your terminal.

You probably already have a nerd-font installed with a lot of glyphs if you are here. And `yazi` have a dependency on `ttf-nerd-font-symbols` that install a font with all the nerd-font glyphs available.

# How

## Option 1 - install.sh // TODO

- Review the code then:
- `curl https://raw.github....install.sh | bash`

## Option 2 - Copy binary into $PATH

- Download the binary `nerd-glyphs`
- Move it into your $PATH ($HOME/.local/bin, /usr/local/bin, ...)
- Execute `nerd-glyphs`

## Option 3 - build from source

> This will try to update the `glyphs.json` from nerdfont github that will later be embedded into `nerd-glyphs`

- `git clone` this repo
- `./update-glyph-list.sh` //TODO - for now it's a manual git clone --depth 1 of  NerdFont git then python script with hardcoded path.
  - it will download the nerdfont necesarry files (`wget` or `curl`) in a temp folder here (`./nerd-font-src/`)
  - and then execute the `parser.py` against the files to generate the `./glyphs.json`
- `go build -o nerd-glyphs main.go`
- `./nerd-glyphs`
- copy `nerd-glyphs` in your $PATH

# Sources

## NerdFont

> NerdFont git have the relatition between glyph and glyph-name

From Nerd Font git
[GitHub](https://github.com/ryanoasis/nerd-fonts): `https://github.com/ryanoasis/nerd-fonts`
file path url: `https://raw.githubusercontent.com/ryanoasis/nerd-fonts/refs/heads/master/bin/scripts/lib/{file-name}`

| file-name | group-name |
| -- | -- |
| i_cod.sh | Codicons |
| i_dev.sh | Devicons |
| i_extra.sh | Progress indicators |
| i_fa.sh | Font Awesome |
| i_fae.sh | Font Awesome Extension |
| i_iec.sh | IEC Power Symbols |
| i_logos.sh | Font Logos |
| i_md.sh | Material |
| i_oct.sh | Octicons |
| i_ple.sh | Powerline Extra |
| i_pom.sh | Pomicons |
| i_seti.sh | Seti-UI |
| i_weather.sh | Weather |
