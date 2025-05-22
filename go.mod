module github.com/devprosvn/VNPrider

go 1.23

require golang.org/x/crypto v0.0.0

require golang.org/x/sys v0.0.0

replace golang.org/x/crypto => ./xdeps/golang.org/x/crypto

replace golang.org/x/sys => ./xdeps/golang.org/x/sys
