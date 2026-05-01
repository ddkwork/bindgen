@echo off
cd /d "%~dp0"
cmake -B build -G "Ninja" -DCMAKE_BUILD_TYPE=Release . && cmake --build build --config Release
copy build\glfw3.dll .\glfw3.dll
