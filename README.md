# disblob
A stylesheet(css) generator for use with Discord.
***
***All original code is done by the insane but genius [diamondburned](https://github.com/diamondburned/)***
> [Master repo](https://github.com/diamondburned/disblob)

***This fork is an attempt to iron out some bugs, and refine the `README`***
***
# Status
|   |Status |	|
|--|--|--|
|   |Edits: |	|
|  | `README` Only |  |

|   |Fork|  |
|--|--|--|
|  | Fork Only |  |
|  |  (No PR)  |  |

|  | Testing |  |
|--|--|--|
|  | 		-Client Only-		 |  |
|  |-  (No Tests have been done outside of the client)  |  
***
 - [x] Client (BD(BetterBiscord))
 - [ ] Client (Other css injectors)
 - [ ] Browsers (FIrefox, Chrome, etc. (Use [stylus](https://github.com/openstyles/stylus) to inject css)
***
|  | Contributors:|  |
|--|--|--|
|  | 		2 ([diamondburned](https://github.com/diamondburned/) & [ThatGeekyWeeb](https://github.com/ThatGeekyWeeb))

***
## Dependencies
  1. (All Blobmoji deps(https://github.com/C1710/blobmoji/blob/master/README.md#building-notocoloremoji))
  2. `go`
  3. `zopfli`
  4. 1+ GiB of storage space
 
## Building

```sh
git submodule update --init
cd blobmoji
make -j$(nproc)
cd ..
mv blobmoji/build/renamed_flags/* blobmoji/svg/
mv blobmoji/build/resized_flags/ blobmoji/

# This will generate the stylesheet.
go run . -defpath definitions.css -datpath data.css > <css file>
```

# Contributing
Any Pull Requests, Opened Issues, and forks are  Greatly welcomed!\
You help make disblob better by Contributing!

> Written with [StackEdit](https://stackedit.io/).
