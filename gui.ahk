; a simple gui for picar.exe
; test on Windows 7 64bit
Enable_Prefix := false
Enable_Rename_Only := false

; GUI CONTROL SETTINGS
Gui, Add, Text, x16 y17 w90 h20 , 选择目录：
Gui, Add, Button, x296 y47 w100 h30 gSelectFloder, 浏览
Gui, Add, Edit, x16 y47 w270 h30 vCtrlEdit_PicFloder
Gui, Add, GroupBox, x16 y87 w380 h110 +Left, 设置
Gui, Add, CheckBox, x26 y117 w100 h30 gSetPrefix, 自定义前缀
Gui, Add, CheckBox, x26 y157 w100 h30 gSetRenameOnly, 不归档
Gui, Add, Edit, x136 y117 w250 h30 vCtrlEdit_Prefix
Gui, Add, Button, x296 y210 w100 h30 gGuiClose, 退出
Gui, Add, Button, x186 y210 w100 h30 gDoRename, 开始

GuiControl, Disable, CtrlEdit_Prefix

Gui, Show, x341 y145 h250 w410, Picar Shell
Return

; program exit
GuiClose:
	ExitApp

; select folder for photos
SelectFloder:
	FileSelectFolder, PicFloder
	if (PicFloder = "")
		return
	GuiControl, , CtrlEdit_PicFloder, %PicFloder%
	return

SetPrefix:
	if (Enable_Prefix)
	{
		Enable_Prefix := false
		GuiControl, , CtrlEdit_Prefix, 
		GuiControl, Disable, CtrlEdit_Prefix
		return
	}
	GuiControl, Enable, CtrlEdit_Prefix
	Enable_Prefix := true
	return

SetRenameOnly:
    if(Enable_Rename_Only)
    {
        Enable_Rename_Only := false
        return
    }
    Enable_Rename_Only := true
    return
	
; start renaming
DoRename:
	if (Enable_Prefix)
	{
		GuiControlGet, Prefix, , CtrlEdit_Prefix
		if (Prefix = "")
		{
			Msgbox, 请输入前缀
			return
		}
		Argv_Prefix = -prefix=%Prefix%
	}
	else
	{
        Argv_Prefix = 
	}
	
	if (Enable_Rename_Only)
	{
        Argv_Rename_only = -renameonly
    }
    else
    {
        Argv_Rename_only = 
    }
	
	GuiControlGet, PicFloder, , CtrlEdit_PicFloder
	if (PicFloder = "")
	{
		MsgBox, 请选择目录
		return
	}

	IfExist, picar.exe
	{
		Run %comspec% /c "picar.exe %Argv_Prefix% %Argv_Rename_only% "%PicFloder%""
		return
	}
	MsgBox, 没有找到 picar.exe
	return