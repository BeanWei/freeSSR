@echo off
::windows��ִ�ж�ʱ����>>>schtasks /create /tn "win_crontab" /tr "run.bat" /sc daily /st 12:30
go run doub.go