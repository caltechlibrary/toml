
# Fork notes

Looks like there are many pull requests lingering in BurntSushi's
toml package. I am forking this so I can add three functions 
that make toml a drop in replacement for how I am using
the go json package.

+ MarshalIndent()
+ Marshal()
+ Unmarshal()

Will make the additiona as a single file "marshaler.go" and submit
issue with it as an attachment.

