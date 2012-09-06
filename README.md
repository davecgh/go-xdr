# go-xdr

Go-xdr implements the data representation portion of the External Data
Representation (XDR) standard protocol as specified in RFC 4506 (obsoletes RFC
1832 and RFC 1014) in Pure Go (Golang).  A comprehensive suite of tests are
provided to ensure proper functionality.  It is licensed under the liberal ISC
license, so it may be used in open source or commercial projects.

## Documentation

Full `go doc` style documentation for the project can be viewed online without
installing this package by using the excellent GoPkgDoc site here:
http://go.pkgdoc.org/github.com/davecgh/go-xdr/xdr

You can also view the documentation locally once the package is installed with
the `godoc` tool by running `godoc -http=":6060"` and pointing your browser to
http://localhost:6060/pkg/github.com/davecgh/go-xdr/xdr/

## Installation

```bash
$ go get github.com/davecgh/go-xdr/xdr
```

## Sample Decode Program

```Go
package main

import (
    "fmt"
    "github.com/davecgh/go-xdr/xdr"
)

func main() {
	// Hypothetical image header format.
	type ImageHeader struct {
		Signature   [3]byte
		Version     uint32
		IsGrayscale bool
		NumSections uint32
	}

	// XDR encoded data described by the above structure.  Typically this would
	// be read from a file or across the network, but use a manual byte array
	// here as an example.
	encodedData := []byte{
		0xAB, 0xCD, 0xEF, 0x00, // Signature
		0x00, 0x00, 0x00, 0x02, // Version
		0x00, 0x00, 0x00, 0x01, // IsGrayscale
		0x00, 0x00, 0x00, 0x0A} // NumSections

	// Declare a variable to provide Unmarshal with a concrete type and instance
	// to decode into.
	var h ImageHeader
	remainingBytes, err := xdr.Unmarshal(encodedData, &h)
	if err != nil {
		fmt.Println(err)
		return
	}
  
	fmt.Println("remainingBytes:", remainingBytes)
	fmt.Printf("h: %+v", h)
}
```

The struct instance, `h`, will then contain the following values:

```Go
h.Signature = [3]byte{0xAB, 0xCD, 0xEF}
h.Version = 2
h.IsGrayscale = true
h.NumSections = 10
```

## Sample Encode Program

```Go
package main

import (
    "fmt"
    "github.com/davecgh/go-xdr/xdr"
)

func main() {
	// Hypothetical image header format.
	type ImageHeader struct {
		Signature   [3]byte
		Version     uint32
		IsGrayscale bool
		NumSections uint32
	}

	// Sample image header data.
	h := ImageHeader{[3]byte{0xAB, 0xCD, 0xEF}, 2, true, 10}

	// Use Marshal to automatically determine the appropriate underlying XDR
	// types and encode.
	encodedData, err := xdr.Marshal(&h)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("encodedData:", encodedData)
}
```

The result, `encodedData`, will then contain the following XDR encoded byte
sequence:

```
0xAB, 0xCD, 0xEF, 0x00,
0x00, 0x00, 0x00, 0x02,
0x00, 0x00, 0x00, 0x01,
0x00, 0x00, 0x00, 0x0A
```

## License

Go-xdr is licensed under the liberal ISC License.