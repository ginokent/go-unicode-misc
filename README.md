# go-unicode-misc

Convert a single byte to a Unicode Escape Sequence.  


```go
b := unicodemisc.UnicodeEscapeSequence(0x7f)
os.Stdout.Write(b)
// -> \u007f
```

```go
s := unicodemisc.UnicodeEscapeSequenceString(0x7f)
fmt.Print(s)
// -> \u007f
```
