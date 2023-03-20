# TIFFer

A library for parsing TIFF files.  
The main purpose of this library is reading meta-data from TIFF files.  

## Description

This library supports reading and parsing TIFF tags and their values.  
Decoding of image formats is not supported by this library.  
Writing and encoding data is also not part of this library.  

This library supports following tags:
* **Baseline tags**  
  (of the in the TIFF 6.0 Specification)

* **Extension tags**  
  (including but not limited to) 
  * XMP tag

* **Private tags**  
  (including but not limited to)
  * DNG tags
  * EXIF tags
  * GPS tags
  * Interoperability tags
  * IPTC tag

The library is able to read tags stored in so called Sub-IFDs which is really 
a bit of a "hack" of the TIFF encoding scheme. TIFF format was designed to 
support only the flat layout of IFDs. SubIFDs in this library can be accessed 
via the `SubIFD` field of the Directory Entry (Tag). A couple of usage examples 
of the library can be viewed in the `example` folder.

## Links
* TIFF Tag Reference at AWARE SYSTEMS  
https://www.awaresystems.be/imaging/tiff/tifftags.html

* TIFF 6.0 Specification (Revision 6.0 Final — June 3, 1992)  
https://developer.adobe.com/content/dam/udp/en/open/standards/tiff/TIFF6.pdf

## History

TIFF format was developed by Aldus Corporation. At this moment, in the year 
2023, the owner of this technology is Adobe Inc., who bought Aldus Corporation 
in September 1994. More information about TIFF format can be found in 
Wikipedia: https://en.wikipedia.org/wiki/TIFF

## Notes

There are some mismatches between Golang's data types and TIFF format data 
types. First of all, it concerns the Rational data types. Go programming 
language has very poor support for built-in rational number types. More 
information can be found in the comments of the `basic-types.go` file.
