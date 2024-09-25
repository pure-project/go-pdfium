echo prefix=$(pwd)/pdfium > pdfium.pc
cat >> pdfium.pc << EOF
includedir=\${prefix}/include
libdir=\${prefix}/lib
Name: PDFium
Description: PDFium
Version: 6392
Requires:
Libs: -L\${libdir} -lpdfium
Cflags: -I\${includedir}
EOF

export CGO_ENABLED=1
export CGO_LDFLAGS='-Wl,-rpath=$ORIGIN'
export PKG_CONFIG_PATH=$PKG_CONFIG_PATH:$(pwd)
export CC=aarch64-linux-gnu-gcc
export LD=aarch64-linux-gnu-ld
export AS=aarch64-linux-gnu-as
export GOOS=linux
export GOARCH=arm64
go build -tags pdfium_experimental -o pdfium-worker worker.go