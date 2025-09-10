# Instalación de GO

- Descargar instalador y seguir instrucciones
🔗[Pagina de Go](https://go.dev/dl/)
- Varificar instalación ´go version´

# Go Path (Estructura de directorios)

- GOPATH (Donde estan los proyectos)
  - bin (binarios)
  - pkg (dependencias internas y externas. Las externas se almacenan en la carpeta mod)
  -src (Se almacenan los proyectos)
- GOROOT (ejecuta el programa)

# Modules & Packages

Los paquetes son carpatas donde se van a almacenar los archivos de codigo Go incluyendo el main. Se debe evitar almacenar un paquete dentro de otro.

Los modulos es la estructura del proyecto el cual contiene todos lo paquetes que hacen que el codigo funcione.