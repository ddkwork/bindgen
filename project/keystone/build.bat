@echo off
cd /d "%~dp0"
cmake -B build -G "Ninja" -DCMAKE_BUILD_TYPE=Release . && cmake --build build --config Release && dir /b build\*.dll 2>nul
