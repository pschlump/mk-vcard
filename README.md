# `VCard` format QR Generator

```
$ git clone https://github.com/pschlump/mk-vcard.git
$ go build
$ edit yourname.txt
$ ./mk-vcard yourname.txt
```

Should crate a `.png` file `yourname.txt.png` with the `VCard` data in it.

A `VCard` file has, for example:

```
BEGIN:VCARD
VERSION:3.0
N:Philip J Schlump
ORG:Truck Coin Swap, LLC
TITLE:Chief Comercial Officer
TEL:1232345678
URL:https://www.truckcoinswap.com
EMAIL:philip@example.com
ADR:USA
END:VCARD
```

