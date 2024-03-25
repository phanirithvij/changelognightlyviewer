## changelog nightly browser

A very basic https://changelog.com/nightly browser. WIP see [docs/DEV.md](docs/DEV.md)

### To run

```shell
git clone https://github.com/phanirithvij/changelognightlyviewer
cd changelognightlyviewer
go build -o changelognightlyviewer.out .

# This won't be needed later but doing it for now
# Alternate is to proxy the urls but I did this instead
mkdir -p content && cd content
nix-shell -p httrack
httrack "https://nightly.changelog.com" -O . "-*email*"
```

