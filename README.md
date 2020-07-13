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
|  | 		All tested		 |  |
|  |-  (Tests have been done with the client and Browsers)  |  
***
 - [x] Client (BD(BetterBiscord)~(100% Working!))
 - [x] Firefox Browser 
 - [x] Google Chrome
 > Use [stylus](https://github.com/openstyles/stylus) for injecting css Within browsers
 > > DO NOT USE STYLISH IT IS [SPYWARE](https://robertheaton.com/2018/07/02/stylish-browser-extension-steals-your-internet-history/)
 
 > NOTE: use within browsers requires one use the `discord_imports.css` from [here](https://builds.sr.ht/~diamondburned/job/249530)(Latest build)
 > Or use this
 > ```css
 > @import url("https://cdn.discordapp.com/attachments/729185721771884623/730214842333528134/style_part00.css");
 > @import url("https://cdn.discordapp.com/attachments/729185721771884623/730214847161041026/style_part01.css");
 > @import url("https://cdn.discordapp.com/attachments/729185721771884623/730214851900735570/style_part02.css");
 > @import url("https://cdn.discordapp.com/attachments/729185721771884623/730214854991806464/style_part03.css");
 > ```
 > Importing with stylus requires one to upload the css to discord's servers, to bypass CORS
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
**Optionally, compile Blobmoji (Flags will not be mapped if you do not)**
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
