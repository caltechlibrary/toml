
# Fork notes

Looks like there are many pull requests lingering in BurntSushi's
toml package. I am forking this so I can add two functions 
that make toml a drop in replacement for where I am using
the go json package.

+ MarshalIndent()
+ Marshal()

Will make the addition as a single file "marshaler.go" and submited
the suggestion as issue #248 with included code for the two functions.

RSD, 2019-08-01

