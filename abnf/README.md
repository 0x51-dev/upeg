# Augmented BNF for Syntax Specifications: ABNF

## References

- [RFC5234](https://www.rfc-editor.org/rfc/rfc5234.html)

## Errata

```abnf
elements       =  alternation *c-wsp
--- becomes ---
elements       =  alternation *WSP
```

```abnf
rulelist       =  1*( rule / (*c-wsp c-nl) )
--- becomes ---
rulelist       =  1*( rule / (*WSP c-nl) )
```

```abnf
repeat         =  1*DIGIT / (*DIGIT "*" *DIGIT)
--- becomes ---
repeat         =  (*DIGIT "*" *DIGIT) / 1*DIGIT
```

```abnf
CRLF           =  CR LF
--- becomes ---
CRLF           =  [CR] LF
```

```abnf
HEXDIG         =  DIGIT / "A" / "B" / "C" / "D" / "E" / "F"
--- becomes ---
HEXDIG         =  DIGIT / %x41-46 / %x61-66   ; A-F / a-f
```
