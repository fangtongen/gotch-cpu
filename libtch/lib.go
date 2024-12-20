package libtch

// #cgo CFLAGS: -I -O3 -Wall -Wno-unused-variable -Wno-deprecated-declarations -Wno-c++11-narrowing -g -Wno-sign-compare -Wno-unused-function
// #cgo CFLAGS: -I/usr/local/include
// #cgo CFLAGS: -D_GLIBCXX_USE_CXX11_ABI=1
// #cgo LDFLAGS: -lstdc++ -ltorch -lc10 -ltorch_cpu -L/lib64
// #cgo CXXFLAGS: -std=c++17 -I -g -O3
// #cgo CFLAGS: -I/libtorch/lib -I/libtorch/include -I/libtorch/include/torch/csrc/api/include -I/libtorch/include/torch/csrc
// #cgo LDFLAGS: -L/libtorch/lib
// #cgo CXXFLAGS: -I/libtorch/lib -I/libtorch/include -I/libtorch/include/torch/csrc/api/include -I/libtorch/include/torch/csrc
import "C"
