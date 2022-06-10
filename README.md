# VCard format QR Generator

```
$ git clone https://github.com/pschlump/mk-vcard.git
$ go build
$ edit yourname.txt
$ ./mk-vcard yourname.txt
```

should crate a .png file `yourname.txt.png` with the vcard data in it.

A vcard file has:

```
BEGIN:VCARD
VERSION:3.0
N:Philip J Schlump
ORG:Truck Coin Swap, LLC
TITLE:Chief Comercial Officer
TEL:+123-234-5678
URL:https://www.truckcoinswap.com
EMAIL:philip@example.com
ADR:USA
END:VCARD
```

