@echo off
REM 모든 .proto 파일에 대해 gRPC 및 Protobuf 코드를 컴파일하는 배치 파일

REM Display current directory
echo Current directory: %CD%

REM proto 디렉토리로 이동
cd proto

REM 환경변수에서 C# 코드 생성 경로와 gRPC 플러그인 경로를 가져옴
if "%UNITY_PROTO_PATH%"=="" (
    echo Error: UNITY_PROTO_PATH 환경변수가 설정되지 않았습니다.
    exit /b 1
)

if "%GRPC_PLUGIN_PATH%"=="" (
    echo Error: GRPC_PLUGIN_PATH 환경변수가 설정되지 않았습니다.
    exit /b 1
)

REM 환경 변수 값 출력
echo UNITY_PROTO_PATH: "%UNITY_PROTO_PATH%"
echo GRPC_PLUGIN_PATH: "%GRPC_PLUGIN_PATH%""

REM 모든 .proto 파일에 대해 Go 및 C# 코드를 컴파일
for %%f in (*.proto) do (
    echo Generating gRPC code for %%f...

    @REM REM Go 코드 생성
    @REM protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative %%f

    REM Go 코드 생성
    protoc --go_out=../src/%%~nf/pb --go_opt=paths=source_relative --go-grpc_out=../src/%%~nf/pb --go-grpc_opt=paths=source_relative %%f

    REM C# 코드 생성
    protoc --proto_path=. --csharp_out="%UNITY_PROTO_PATH%/%%~nf" --grpc_out="%UNITY_PROTO_PATH%/%%~nf" --plugin=protoc-gen-grpc="%GRPC_PLUGIN_PATH%" %%f
)

echo gRPC code generation completed for all .proto files.
pause
