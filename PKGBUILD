# Maintainer: Brenek Harrison <brenekharrison @ gmail d0t com>
pkgname=must-local
pkgver=0.0.1
pkgrel=1
epoch=
pkgdesc="Hobby AUR helper with apt-like syntax"
arch=("x86_64" "aarch64")
url="https://github.com/BrenekH/must#readme"
license=("GPL")
groups=()
depends=("glibc")
makedepends=("go" "git")
checkdepends=()
optdepends=()
provides=("must")
conflicts=()
replaces=()
backup=()
options=()
install=
changelog=
source=("$pkgname"::"git+file://$(pwd)")
noextract=()
sha256sums=('SKIP')
validpgpkeys=()

build() {
	cd "$pkgname"

	export CGO_CPPFLAGS="${CPPFLAGS}"
	export CGO_CFLAGS="${CFLAGS}"
	export CGO_CXXFLAGS="${CXXFLAGS}"
	export CGO_LDFLAGS="${LDFLAGS}"
	export GOFLAGS="-buildmode=pie -trimpath -ldflags=-linkmode=external -mod=readonly -modcacherw"

	go build -o must ./cmd/main.go
}

check() {
	cd "$pkgname"

	go test ./...
}

package() {
	cd "$pkgname"

	install -Dm755 must "$pkgdir/usr/bin/must"
}
