# disblob

A blobmoji stylesheet generator for Discord.

Latest release [can be found on sr.ht](https://builds.sr.ht/~diamondburned/job/249530)

Note: importing from remote sites won't work on the website due to CORS, so to bypass this you can upload the files to discord itself, which is done in the `discord_imports.css` file on sr.ht.

---

## Building

### Dependencies
  1. `go`
  2. `zopfli`

```sh
git clone https://github.com/diamondburned/disblob.git
cd disblob
git submodule update --init
```

Optionally, compile Blobmoji which will map the flags

(All Blobmoji [deps](https://github.com/C1710/blobmoji/wiki/Build-instructions) are required)

```sh
cd blobmoji
make -j$(nproc)
cd ..
mv blobmoji/build/renamed_flags/* blobmoji/svg/
mv blobmoji/build/resized_flags/ blobmoji/
```

---

Then generate the stylesheet

```sh
go run . > ./output.css
```
