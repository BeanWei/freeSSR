@echo off
::windows下执行定时任务>>>schtasks /create /tn "win_crontab" /tr "run.bat" /sc daily /st 12:30
go run doub.go