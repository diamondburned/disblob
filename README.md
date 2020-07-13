# disblob
A stylesheet(css) generator for use with Discord.
***
***All original code is done by the insane but genius [diamondburned](https://github.com/diamondburned/)***
> [Master repo](https://github.com/diamondburned/disblob)

***This fork is an attempt to refine the `README` and make it easier for others to setup!***
***
# Status
|   |Status |	|
|--|--|--|
|   |Edits: |	|
|  | `README` Only |  |

|   |Fork|  |
|--|--|--|
|  | Fork Only |  |
|  |  (PR [Open](https://github.com/diamondburned/disblob/pull/1))  |  |

|  | Testing |  |
|--|--|--|
|  | 		-Client Only-		 |  |
|  |-  (No Tests have been done outside of the client)  |  
***
 - [x] Client (BD(BetterBiscord)~(100% Working!)
 - [ ] Client (Other css injectors)
 - [ ] Browsers (Chrome, etc. (Use [stylus](https://github.com/openstyles/stylus) to inject css)
 - [x] Firefox Browser (Testing done with stylish(DO NOT USE(Considered spyware)) 
***
|  | Contributors:|  |
|--|--|--|
|  | 		2 ([diamondburned](https://github.com/diamondburned/) & [ThatGeekyWeeb](https://github.com/ThatGeekyWeeb))

***
## Dependencies
  1. (All Blobmoji deps(https://github.com/C1710/blobmoji/wiki/Build-instructions))
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
## Releases
My github pages link can be used to `@import` the css
```css
@import url("https://thatgeekyweeb.github.io/src/import.data.css");
```

Additionally there are releases [here!](https://github.com/ThatGeekyWeeb/disblob/releases)


# Contributing
Any Pull Requests, Opened Issues, and forks are  Greatly welcomed!\
You help make disblob better by Contributing!

> Written with [StackEdit](https://stackedit.io/).

<!--stackedit_data:
eyJoaXN0b3J5IjpbMTYyMDA3NjI2Nyw3Mjk5OTIyMjIsLTE0OT
E2OTMyMzQsMTMzNjM1MTI5Nl19
-->
<!--stackedit_data:
eyJoaXN0b3J5IjpbLTcwODM1NzY2MywtMTE5NDc1MDI4MV19
-->
