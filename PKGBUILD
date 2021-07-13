# Maintainer: Brenek Harrison <brenekharrison @ gmail dot com>
pkgname=my-aur-helper-git
pkgver=0.0.1
pkgrel=1
epoch=
pkgdesc="Prototype AUR helper."
arch=("x86_64" "aarch64")
url="https://github.com/BrenekH/my-aur-helper#readme"
license=("GPL")
groups=()
depends=("glibc")
makedepends=("go" "git")
checkdepends=()
optdepends=()
provides=()
conflicts=()
replaces=()
backup=()
options=()
install=
changelog=
source=("$pkgname"::"git+file:///home/brenekh/repos/my-aur-helper")
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

	go build -o mah ./cmd/main.go
}

check() {
	cd "$pkgname"

	go test ./...
}

package() {
	cd "$pkgname"

	install -Dm755 mah "$pkgdir/usr/bin/mah"
}
