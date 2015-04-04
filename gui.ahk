; 一个简单picar.exe图形界面
; picar.exe 是一个照片归档工具，可以按“年月”文件夹形式归档并根据EXIF中
; 拍照日期信息并命名文件名。

; this script was tested on WINXP
; update: 2014-08-28 by yuekcc@qq.com
; ----------

; 主程序文件名
COMMAND_BIN := "picar.exe"

; 启动自定义前缀标记
ENABLE_PREFIX := false

; GUI CONTROL SETTINGS
; 创建GUI控件
Gui, Add, Text, x16 y17 w90 h20 , 选择目录：
Gui, Add, Button, x296 y47 w100 h30 gSelectFloder, 浏览
Gui, Add, Edit, x16 y47 w270 h30 vCtrlEdit_PicFloder
Gui, Add, GroupBox, x16 y87 w380 h80 +Left, 设置
Gui, Add, CheckBox, x26 y117 w100 gSettingPrefix, 自定义前缀
Gui, Add, ComboBox, x136 y117 w250 vCtrlEdit_Prefix, NexusS|Note3|H60
Gui, Add, Button, x296 y177 w100 h30 gGuiClose, 退出
Gui, Add, Button, x186 y177 w100 h30 gDoRename, 开始

; 设置控件初始状态
GuiControl, Disable, CtrlEdit_Prefix

; 显示主窗口
Gui, Show, x341 y145 h227 w420, picar
Return

; 退出程序
GuiClose:
	ExitApp

; 选择照片目录
SelectFloder:
	FileSelectFolder, PicFloder
	if (PicFloder = "")
		return
	GuiControl, , CtrlEdit_PicFloder, %PicFloder%
	return

; 设置文件名前缀
SettingPrefix:
	if (ENABLE_PREFIX)
	{
		ENABLE_PREFIX := false
		GuiControl, , CtrlEdit_Prefix, 
		GuiControl, Disable, CtrlEdit_Prefix
		return
	}
	GuiControl, Enable, CtrlEdit_Prefix
	ENABLE_PREFIX := true
	return
	
; start renaming
; 启动exifrename.py脚本
DoRename:
	if (ENABLE_PREFIX)
	{
		GuiControlGet, Prefix, , CtrlEdit_Prefix
		if (Prefix = "")
		{
			Msgbox, 请输入前缀
			return
		}
		Argv_Prefix = %Prefix%
	}
	
	GuiControlGet, PicFloder, , CtrlEdit_PicFloder
	if (PicFloder = "")
	{
		MsgBox, 请选择目录
		return
	}

	;Msgbox, %Argv_Prefix%
	;MsgBox, "%COMMAND_BIN% -prefix=%Argv_Prefix% -dir="%PicFloder%""
	IfExist, %COMMAND_BIN%
	{
		if (ENABLE_PREFIX)
		{
			; 重命名照片，并归档
			Run %comspec% /c "%COMMAND_BIN% -prefix=%Argv_Prefix% -dir="%PicFloder%""
		}
		else
		{
			; 只重命名照片，不归档
			Run %comspec% /c "%COMMAND_BIN% -renameonly=true -dir="%PicFloder%""
		}
		
		return
	}
	return
