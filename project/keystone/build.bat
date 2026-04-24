cmake -B build -G "Ninja" -DCMAKE_BUILD_TYPE=Release -DCMAKE_MAKE_PROGRAM=ninja.exe . && cmake --build build --config Release && dir /b build\*.dll 2>nul
