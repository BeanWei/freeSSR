@echo off
::windows定时任务指令>>>schtasks /create /tn "freeSSR" /tr "run.bat" /sc hourly /mo 3
::删除定时任务指令>>>schtasks /delete /tn "freeSSR"

go run doub.go