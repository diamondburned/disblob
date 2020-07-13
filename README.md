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
 - [x] Client (BD(BetterBiscord)
 - [ ] Client (Other css injectors)
 - [ ] Browsers (FIrefox, Chrome, etc. (Use [Stylish](https://userstyles.org/) to inject css)
***
|  | Contributors:|  |
|--|--|--|
|  | 		2 ([diamondburned](https://github.com/diamondburned/) & [ThatGeekyWeeb](https://github.com/ThatGeekyWeeb))

***
## Dependencies
 
 1. `go`
  2. `zopfli`
  3. 1+ GiB of storage space
 
## Building

```sh
git submodule update --init
cd ./blobmoji
make

# This will generate the stylesheet.
go run . -defpath definitions.css -datpath data.css
```

# Contributing
Any Pull Requests, Opened Issues, and forks are  Greatly welcomed!\
You help make disblob better by Contributing!

> Written with [StackEdit](https://stackedit.io/).
<!--stackedit_data:
eyJoaXN0b3J5IjpbMTE2Nzk0OTc4LDUzNDgyMDU2Ml19
-->