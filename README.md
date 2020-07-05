# disblob

A stylesheet generator for use with Discord.

## Building

```sh
# This will generate the stylesheet.
go run . >> style.css
```

This generator will skip emojis that are already generated. To reset this,
remove `state.db` and use `>` instead of `>>`.
