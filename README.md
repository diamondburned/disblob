# disblob
A stylesheet(css) generator for use with Discord.
***
***All original code is done by the insane but genius [diamondburned](https://github.com/diamondburned/)***
> [Master repo](https://github.com/diamondburned/disblob)

***This fork is an attempt to refine the `README` and make it easier for others to setup!***
***
# Status

|  | Testing |  |
|--|--|--|
|  | 		-Client + Firefox-		 |  |
|  |-  (Tests have been done with the client and Firefox)  |  
***
 - [x] Client (BD(BetterBiscord)~(100% Working!))
 - [ ] Client (Other css injectors)
 - [x] Firefox Browser (Testing done with stylish(DO NOT USE(Considered spyware)) 
 - [x] Google Chrome
***
***
|  | Contributors:|  |
|--|--|--|
|  | 		2 ([diamondburned](https://github.com/diamondburned/) & [ThatGeekyWeeb](https://github.com/ThatGeekyWeeb))

***
## Dependencies
  1. `go`
  2. `zopfli`
***
## Building

```sh
git clone https://github.com/diamondburned/disblob.git
cd disblob
git submodule update --init
```
**Optinally, compile Blobmoji (Flags will not be mapped if you do not)**
>(All Blobmoji [deps](https://github.com/C1710/blobmoji/wiki/Build-instructions) are required)
> >```sh
> >cd blobmoji
> >make -j$(nproc)
> >cd ..
> >mv blobmoji/build/renamed_flags/* blobmoji/svg/
> >mv blobmoji/build/resized_flags/ blobmoji/
> >```
***
```sh
# This will generate the stylesheet.
go run . > ./output.css
```
