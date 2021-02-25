# disblob

A blobmoji stylesheet generator for Discord.

Latest release [can be found on sr.ht](https://builds.sr.ht/~diamondburned/job/249530).

Note: importing from remote sites won't work on the website due to CORS; to
bypass this you can upload the files to discord itself, which is done in the
`discord_imports.css` file on sr.ht.

---

## Building

### Dependencies
  1. `go`
  3. `svgo` (optional, only for SVG compression)
  2. `zopfli` (optional, only for blobmoji)

### Build

```sh
git clone https://github.com/diamondburned/disblob.git
cd disblob
git submodule update --init
```

#### Building Blobmoji

Optionally, compile Blobmoji which will map the flags.

(All Blobmoji [deps](https://github.com/C1710/blobmoji/wiki/Build-instructions) are required)

```sh
cd blobmoji
make -j$(nproc)
mv build/renamed_flags/* svg/
mv build/resized_flags/  ./
cd ..
```

---

## Usage

Below are known output flags and their respective output sizes:

- `-png`: 20MB
- `-nosvgo`: 31MB
- svgo (no flags): 28MB

```sh
go run . > ./output.css
```
