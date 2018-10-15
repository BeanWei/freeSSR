@echo off
::windows定时任务指令>>>schtasks /create /tn "freeSSR" /tr "run.bat" /sc hourly /mo 3
go run doub.go