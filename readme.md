# API Description

This api provide create cells to given any shape of polygon and can find out given coordinate inside that create polygon. 
Used technology for polygon Create ApI
Go language 
Google/s2 GEO library 


### How to setup GO Environment 

GO language canonical Git repository is located at https://go.googlesource.com/go.
There is a mirror of the repository at https://github.com/golang/go.

Unless otherwise noted, the Go source files are distributed under the BSD-style license found in the LICENSE file.

### Download and Install

#### Binary Distributions

Official binary distributions are available at https://golang.org/dl/.
After downloading a binary release, visit https://golang.org/doc/install or load doc/install.html in your web browser for installation instructions.

#### Install From Source
If a binary distribution is not available for your combination of operating system and architecture, visithttps://golang.org/doc/install/source or load doc/install-source.html in your web browser for source installation instructions.
Referenced - https://github.com/golang/go


# About Google S2 Library

Google’s S2 library is a real treasure, not only due to its capabilities for spatial indexing but also because it is a library that was released more than 4 years ago and it didn’t get the attention it deserved. The S2 library is used by Google itself on Google Maps, MongoDB engine and also by Foursquare, but you’re not going to find any documentation or articles about the library anywhere except for a paper by Foursquare, a Google presentation and the source code comments. You’ll also struggle to find bindings for the library, the official repository has missing Swig files for the Python library and thanks to some folks we can have a partial binding for the Python language. We heard that Google is actively working on the library right now and we are probably soon going to get more details about it when they release this work, but We decided to share some examples about the library and the reasons why We think that this library is so cool.

Get to more about - http://blog.christianperone.com/2015/08/googles-s2-geometry-on-the-sphere-cells-and-hilbert-curve/

###  API Description 
This API provide 2 Different APIs to do Geographical representations 

`/createpolygon`

* getting coordinate given custom polygon 
* filtering the polygon coordinates
* create rect according to given polygon
* create covering area using created rect 

`/getselldetails`

* getting coordinate from selected point 
* creating s2 cellid given point coordinate
* compare with created cellid 

#####  USE AT YOUR OWN RISK
