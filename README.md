# Sistema de gestion de ecommerce

El objetivo del proyecto es implementar un sistema básico de gestión para un ecommerce

## Pre-requisitos

Antes de comenzar, asegúrate de tener instalado lo siguiente:

* **Go**: Versión 1.20 o superior. [Descargar aquí](https://go.dev/dl/)
* **XAMPP**: Para la gestión de la base de datos MySQL. [Descargar aquí](https://www.apachefriends.org/es/index.html)
* **Git**: Para clonar el repositorio.

---

## Configuración Inicial

### 1. Clonar el proyecto
Abre una terminal y ejecuta el siguiente comando:
```bash
git clone [https://github.com/teban99-rp/ecommerce-golang](https://github.com/teban99-rp/ecommerce-golang)
cd ecommerce-golang
```
### 2. Prepearar la Base de Datos (XAMPP)
1. Abre el XAMPP Control Panel
2. Inicia los módulos Apache y MySQL
3. Entra a [admin mysql](http://localhost/phpmyadmin)
4. Crea la base de datos **ecommerce**

### 3. Ejecución
- Ejecutar el comando
```bash
go mod tidy
```
Para instalar todos los paquetes necesarios que ocupa el proyecto
- Ejecutar la aplicación
```bash
go run main.go
```
Si todo esta correcto, se podrá ver un mensaje indicando que el servidor esta corriendo. (__Server started on :8080__)

### 4. Video Explicativo
[Enlace del video](https://mailinternacionaledu-my.sharepoint.com/:f:/g/personal/esriospe_uide_edu_ec/IgAerqxojjFRQpjX5I4fj_z0ATRCSGZcQKLwA0k9Adq2F00?e=r9yJdc)