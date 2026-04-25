@echo off
cd /d "%~dp0"
REM git clone --recursive https://github.com/zyantific/zydis.git
cmake -B build -G "Ninja" -DCMAKE_BUILD_TYPE=Release . && cmake --build build --config Release
