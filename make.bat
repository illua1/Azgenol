@echo off

set PROJECT_NAME=Azgenol
set BUILD_DIR=bin

set BUILD_MOD=%1



IF DEFINED BUILD_MOD (
  echo Debug mod
) ELSE (
  echo 
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
goto BUILD


:BUILD
echo Build
cd %BUILD_DIR%
IF exist %PROJECT_NAME%.exe (
  del %PROJECT_NAME%.exe
  echo File was deleted
)
cd ..
go build -o="%BUILD_DIR%/%PROJECT_NAME%.exe" .
goto RUN


:RUN
cd %BUILD_DIR%
%PROJECT_NAME%
cd ..

:END
::pause