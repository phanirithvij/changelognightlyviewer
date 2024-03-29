## changelog nightly browser

A very basic https://changelog.com/nightly browser. WIP see [docs/DEV.md](docs/DEV.md)

### To run

```shell
git clone https://github.com/phanirithvij/changelognightlyviewer
cd changelognightlyviewer
go build -o changelognightlyviewer.out .

# In a new shell
mkdir -p content && cd content
nix-shell -p httrack # or apt-get/pacman/yum/brew
httrack "https://nightly.changelog.com" -O . "-*email*" --preserve
# This won't be needed later but doing it for now
# Alternate is to proxy the urls but I did this instead
```

Then run `./changelognightlyviewer.out -port 8080`
visit http://localhost:8080/browse

Disable javascript and try it too.

- [Screenshot_20240325_174452](https://github.com/phanirithvij/changelognightlyviewer/assets/29627898/bf4cf987-c712-4ca2-bc85-3ea6ee3c3d63)
- [demo video](https://youtu.be/VE5J3XzIMIk)
