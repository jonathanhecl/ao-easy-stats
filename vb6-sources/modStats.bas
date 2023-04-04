Attribute VB_Name = "modStats"
Option Explicit

Public Const EVENT_LOGIN = "LOGIN"
Public Const EVENT_LOGOUT = "LOGOUT"
Public Const EVENT_CONTINUE = "CONTINUE"
Public Const EVENT_INITIALIZED = "INITIALIZED"

Public Sub RecordStat(SEvent As String, CharName As String)

    Dim file As Integer
    Dim SDate As String
    Dim SHour As String
    
    SDate = Format(Now, "yyyy-MM-dd")
    SHour = Format(Now, "hh:mm:ss")
    
    If Dir(App.Path & "\stats\", vbDirectory) = vbNullString Then
        MkDir App.Path & "\stats\"
    End If

    file = FreeFile
    Open App.Path & "\stats\" & SDate & ".txt" For Append Shared As #file
    Print #file, SHour & vbTab & CharName & vbTab & SEvent
    Close #file

End Sub
