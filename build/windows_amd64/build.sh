# use MinGW-w64
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
go build -tags pdfium_experimental -o pdfium-worker.exe worker.go
