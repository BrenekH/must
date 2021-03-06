# Maintainer: Brenek Harrison <brenekharrison @ gmail d0t com>
pkgname=must-local
pkgver="$(git rev-parse --short HEAD)"
pkgrel=1
epoch=
pkgdesc="AUR helper with apt-like syntax"
arch=("x86_64" "aarch64")
url="https://github.com/BrenekH/must#readme"
license=("GPL")
groups=()
depends=("glibc" "git")
makedepends=("go")
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

	go build -o must -ldflags="-X 'main.Version=local-$pkgver'" ./cmd/main.go
}

check() {
	cd "$pkgname"

	go test ./...

	./must --version
}

package() {
	cd "$pkgname"

	install -Dm755 must "$pkgdir/usr/bin/must"
}
