### AO Easy Stats

Este proyecto genera un gráfico estadístico a partir de datos de texto generados por Argentum Online.

![Image](https://raw.githubusercontent.com/jonathanhecl/ao-easy-stats/main/portada.png)

### Implementación en el servidor de AO (VB6)

1. Incluir el módulo `modStats.bas` en el servidor de Argentum Online.
2. Cuando se inicia el servidor, en `General.bas Sub Main()` añadir al inicio de la función:
````
   Call modStats.RecordStat(modStats.EVENT_INITIALIZED, "")
````
3. Cada vez que inicia sesión un personaje, en `TCP.bas Sub ConnectUser()` añadir debajo de `.flags.UserLogged = True`:
````
    Call modStats.RecordStat(modStats.EVENT_LOGIN, .Name)
````
4. Cada vez que cierre sesión un personaje, en `TCP.bas Sub CloseUser()` añadir debajo de `.flags.UserLogged = False`:
````
    Call modStats.RecordStat(modStats.EVENT_LOGOUT, .Name)
````
5. En algún Timer del servidor, que se ejecute cada 10~60 segundos, para verificar si ha cambiado el día, añadir:
````
   If lastStatDate <> Day(Date) Then
   Dim LoopC As Integer

        For LoopC = 1 To MaxUsers
            With UserList(LoopC)
                If .ConnIDValida And _
                    .flags.UserLogged Then
                        Call modStats.RecordStat(modStats.EVENT_CONTINUE, .Name)
                End If
            End With
        Next
   Else
        lastStatDate = Day(Date)
   End If
````
6. No olvidarse de añadir en el mismo lugar donde se encuentre el Timer, pero fuera de él, declarar la variable `lastStatDate`:
````
    Private lastStatDate As Byte
````

### ¿Cómo funciona?

* Cada vez que se inicie el servidor, se conecte un personaje, se desconecte, se creara una archivo .txt por día, en donde guardará el registro de eventos.
* Estos archivos se guardan en la carpeta `stats` del servidor.
* Estos eventos se pueden visualizar en el archivo `stats.html` cuando se ejecute el generador de gráficos.
* El archivo `stats.html` se puede abrir en cualquier navegador web o incluir dentro de otra web mediante iframe.

### Generador de gráficos

1. Descargar el archivo compilado `ao-easy-stats.exe` desde la sección de [releases](https://github.com/jonathanhecl/ao-easy-stats/releases).
2. Colocar el archivo `ao-easy-stats.exe` junto con la carpeta `stats` y ejecutar para generar `stats.html`.
3. Abrir el archivo `stats.html` en un navegador web.

### GS-Zone (c) 2023