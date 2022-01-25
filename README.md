# Pkmnchooser

## Introduction

pkmnchooser (or Pokémon Chooser) aims to help to choose a Pokémon to your team based on moves you desire. It is constructed on the top of PokéAPI (see https://pokeapi.co/).

## Objectives

This project has as it main goal to improve my knowledge on the Go programming language. More improvements will be add, so this is still WIP.

## Getting Started

To use pkmncs, first clone the repository, cd into project and compile everything with Go Build tool
```
git clone https://github.com/pvskp/pkmnchooser.git
cd pkmnchooser
go build
```
Then, you can just execute the binary file using each desired move as an argument. For example:

`./pkmncs thunderbolt flamethrower water-gun`

Should return

```
castform
castform-rainy
castform-snowy
castform-sunny
chansey
clefable
clefairy
dragonair
dragonite
dratini
gyarados
jigglypuff
kangaskhan
lickitung
mew
mewtwo
nidoking
nidoqueen
rhydon
snorlax
wigglytuff
```
## LICENSE

Feel free to use this project as you wish.
