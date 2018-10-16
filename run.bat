@echo off
::windows定时任务指令>>>schtasks /create /tn "freeSSR" /tr "run.bat" /sc hourly /mo 3
::删除定时任务指令>>>schtasks /delete /tn "freeSSR"

::隐藏执行批处理时的控制台窗口
if "%1" == "h" goto begin   
mshta vbscript:createobject("wscript.shell").run("%~nx0 h",0)(window.close)&&exit   

:begin 
go run doub.go