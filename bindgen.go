package bindgen

type (
	Binder interface {
		BindCgo()
		BindSharedLibrary()
		C2go()
		CPP2go()
	}
)

// clang -E -P temp.h > output_file.txt
// https://github.com/KhronosGroup/OpenGL-Registry/tree/main/api/GLES3
