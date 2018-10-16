'windwos下定时任务指令：schtasks /create /tn "freeSSR" /tr "run.vbs" /sc hourly /mo 3
'创建vbs脚本来执行go脚本，可隐藏windows控制台
CreateObject("WScript.Shell").Run "cmd /c go run doub.go",0