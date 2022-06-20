@echo off

set PROJECT_NAME=Azgenol
set BUILD_DIR=bin

set BUILD_MOD=%1
set BUILD_ID_FLAG=-ldflags=""

IF DEFINED BUILD_MOD (
  IF "%BUILD_MOD%"=="init" (
    echo Init...
    go mod init %PROJECT_NAME%
    go install -v "github.com/tc-hib/go-winres@latest"
    go mod tidy
    goto PREPARE_BUILD_DIR
  )ELSE IF "%BUILD_MOD%"=="help" (
    echo Supported commands:
    echo - make = Just build
    echo - make init = First make
    echo - make reliase = Make reliase build
    echo - make format = Auto call formating all libs

    echo - make help = support info
    goto END
  )ELSE IF "%BUILD_MOD%"=="reliase" (
    echo Reliase build
    set BUILD_ID_FLAG=-ldflags="-H=windowsgui"
    goto PREPARE_BUILD_DIR
  )ELSE IF "%BUILD_MOD%"=="format" (
    goto FORMAT
  )
  echo Undefine command
  goto END
)
goto PREPARE_BUILD_DIR



:PREPARE_BUILD_DIR
IF exist %BUILD_DIR% (
  echo Go to bin
) ELSE (
  md %BUILD_DIR%
)
goto PREPARE_ASSETS



:PREPARE_ASSETS
IF exist %BUILD_DIR%/Assets (
  echo Assets exist
) ELSE (
  cd %BUILD_DIR%
  md Assets
  cd ..
  xcopy Assets "%BUILD_DIR%/Assets" /E /H /I /Q
  echo Assets copied
)
goto PREPARE_WIN_DATA_RES



:PREPARE_WIN_DATA_RES
go-winres make --in="./BuildData/WinDataRes/WindowsDataRes.json" --out="rsrc_%PROJECT_NAME%"
goto BUILD



:BUILD
echo Build
cd %BUILD_DIR%
IF exist %PROJECT_NAME%.exe (
  del %PROJECT_NAME%.exe
  echo File was deleted
)
cd ..
go build %BUILD_ID_FLAG% -o="%BUILD_DIR%/%PROJECT_NAME%.exe" .
goto CLEAR_WIN_DATA_RES



:CLEAR_WIN_DATA_RES
del "rsrc_%PROJECT_NAME%_windows_amd64.syso"
del "rsrc_%PROJECT_NAME%_windows_386.syso"
goto RUN




:FORMAT
echo Formating
cd AzgenolKernelLib
@echo on
go fmt ./...
@echo off
cd ..
cd AzgenolExecLib
@echo on
go fmt ./...
@echo off
cd ..
goto END



:RUN
cd %BUILD_DIR%
%PROJECT_NAME%
cd ..

:END
::pause