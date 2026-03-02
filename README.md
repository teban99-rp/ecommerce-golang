# Sistema de gestion de ecommerce

Este proyecto es una plataforma de comercio electrónico desarrollada en **Go (Golang)**, diseñada para gestionar usuarios, productos y procesos de compra mediante una arquitectura de microservicios web.

---

## I. Información del Proyecto
* **Nombre:** Esteban Rios
* **Objetivo:** Desarrollar un backend escalable y eficiente para un E-commerce, aplicando concurrencia y servicios API Rest.
* **Fecha:** 01 de marzo de 2026
* **Tecnologías Principales:** 
    * Lenguaje: Go
    * Framework: Gin Gonic
    * Base de Datos: MySQL
    * Formato de Datos: JSON

---

## II. Pre-requisitos

Antes de comenzar, asegúrate de tener instalado lo siguiente:

* **Go**: Versión 1.20 o superior. [Descargar aquí](https://go.dev/dl/)
* **XAMPP**: Para la gestión de la base de datos MySQL. [Descargar aquí](https://www.apachefriends.org/es/index.html)
* **Git**: Para clonar el repositorio.

---

## III. Configuración Inicial

### a. Clonar el proyecto
Abre una terminal y ejecuta el siguiente comando:
```bash
git clone https://github.com/teban99-rp/ecommerce-golang
cd ecommerce-golang
```
### b. Prepearar la Base de Datos (XAMPP)
1. Abre el XAMPP Control Panel
2. Inicia los módulos Apache y MySQL
3. Entra a [admin mysql](http://localhost/phpmyadmin)
4. Crea la base de datos **ecommerce**

### c. Ejecución
- Ejecutar el comando
```bash
go mod tidy
```
Para instalar todos los paquetes necesarios que ocupa el proyecto
- Configurar archivo .env
```
DB_USER= root
DB_PASSWORD= 
DB_HOST= localhost
DB_PORT= 3306
DB_NAME= ecommerce
```
- Ejecutar la aplicación
```bash
go run main.go
```
Si todo esta correcto, se podrá ver un mensaje indicando que el servidor esta corriendo. (__Server started on :8080__)

## IV. Introducción y Funcionalidades

### a. Generación de Servicios Web (API)
El sistema cuenta con una serie de endpoints que permiten la interacción completa con el ecosistema de la tienda. Las APIS se dividen en tres capas de acceso (públicom, protegido y administrativo). 

Se han implementado los siguientes servicios web:

#### 🔓 Rutas Públicas (Autenticación y Catálogo)
1.  **Login (`POST /api/login`):** Validación de credenciales y entrega de Token JWT.
2.  **Registro (`POST /api/register`):** Creación de nuevos usuarios en el sistema.
3.  **Catálogo (`GET /api/products`):** Visualización de productos para usuarios no registrados.

#### 🔐 Rutas Protegidas (Clientes con JWT)
4.  **Añadir al Carrito (`POST /api/add_cart`):** Gestión de persistencia de productos por usuario.
5.  **Ver Carrito (`GET /api/cart/:user_id`):** Consulta de artículos seleccionados.
6.  **Crear Orden (`POST /api/create_order`):** Conversión de carrito a orden de compra.
7.  **Mis Órdenes (`GET /api/orders/:user_id`):** Historial de compras del cliente.
8.  **Procesar Pago (`POST /api/orders/payment`):** Simulación y validación de pagos.

#### 🛡️ Rutas Administrativas (Role-Based Access Control)
9.  **Gestión de Usuarios (`GET /api/admin/users`):** Listado y control de cuentas.
10. **CRUD de Productos:** Creación, edición, actualización y eliminación de productos.
    
    - **Obtener Producto: (`GET /api/admin/product/:product_id`).**
    - **Crear Producto: (`POST /api/admin/products`).**
    - **Editar Producto: (`PUT /api/admin/product/:product_id`).**
    - **Eliminar Producto: (`DELETE /api/admin/delete/product/:product_id`).**
11. **Logística de Órdenes:** Servicios para cambiar el estado de las ordenes.
    - **Enviar Orden: (`POST /api/admin/orders/ship/:id`).**
    - **Cancelar Orden: (`POST /api/admin/orders/cancel/:id`).**

### b. Serialización de Datos
Toda la transferencia de información entre el cliente y el servidor se realiza mediante **JSON**, asegurando que la API sea compatible con aplicaciones móviles, web o de escritorio.

### c. Visualización del Futuro

En base al gran potencial que tiene Go, se puede considerar a futuro la administración de imágenes para darle un mayor realce a la plataforma. También se puede considerar la opción de integrar el backend con algoritmos de **Machine Learning** para la predicción de demanda. En el futuro, el sistema podrá gestionar inventarios de forma autónoma, comunicándose con proveedores mediante APIs automatizadas para reabastecer productos antes de que se agoten, optimizando la cadena de suministro global.

---

## V. Tabla de Integración de Conocimientos

### 🗓️ Cronograma de Desarrollo e Integración (8 Semanas)

A continuación se detalla la evolución del proyecto a través de las unidades de la asignatura:

| Semanas 1 - 2 (Unidad 1 y 2) | Semanas 3 - 6 (Unidades 2 y 3) | Semanas 7 - 8 (Unidad 4) |
| :--- | :--- | :--- |
| **Selección del Sistema:** Desarrollo de un sistema de gestión de E-commerce. | **Modelado de Datos:** Identificación de clases/estructuras: `Cart`, `CartItem`, `Inventory`, `Order`, `OrderItem`, `Product` y `User`. | **Rutas y Servicios:** Implementación de rutas para el manejo de la aplicación y servicios Web. |
| **Definición de Objetivos:** Establecimiento de módulos y funcionalidades principales. | **Lógica de Negocio:** Implementación de encapsulación, manejo de errores e interfaces para un flujo dinámico. | **Frontend/Vistas:** Desarrollo de templates para una interfaz de usuario amigable. |
| **Módulos Definidos:** Usuarios, Productos, Carrito de Compras, Órdenes y Almacenamiento. | **Planificación y Código:** Avance significativo en la programación y documentación interna mediante comentarios. | **QA y Testing:** Aplicación de pruebas para validar el funcionamiento óptimo del sistema. |
| **Arquitectura:** Definición de paquetes necesarios para el funcionamiento del proyecto. | **Serialización:** Implementación de estructuras adaptadas al formato **JSON** para el manejo de información. | **Documentación Final:** Entrega de la documentación técnica completa del aplicativo. |

---

## VI. Conclusiones y Reflexión
El desarrollo de este proyecto permitió comprender la potencia de Go para manejar múltiples peticiones simultáneas. La mayor dificultad fue la estructuración de los controladores de base de datos, pero se logró una arquitectura limpia y también se logro aplicar todos los conocimientos vistos a lo largo de la materia.

---

## VII. Anexos

### Pruebas de la API (Postman)
Para facilitar la revisión de los servicios web, se ha incluido una colección de Postman en la carpeta:
`files/Entregable_final/ecommerce-go.postman_collection.json`

Solo debes importarla en Postman para probar todos los endpoints (Login, Carrito, Admin, etc.).

### Video Explicativo
[Enlace del video](https://mailinternacionaledu-my.sharepoint.com/:v:/g/personal/esriospe_uide_edu_ec/IQAWHcRecD_dQZG9za4kRaB0AXZRjZJmmdPmuYGD5-5wOho?nav=eyJyZWZlcnJhbEluZm8iOnsicmVmZXJyYWxBcHAiOiJPbmVEcml2ZUZvckJ1c2luZXNzIiwicmVmZXJyYWxBcHBQbGF0Zm9ybSI6IldlYiIsInJlZmVycmFsTW9kZSI6InZpZXciLCJyZWZlcnJhbFZpZXciOiJNeUZpbGVzTGlua0NvcHkifX0&e=Qv8pak)

### Archivo Informe
`files/Entregable_final/Informe_proyecto_final.pdf`

Solo debes importarla en Postman para probar todos los endpoints (Login, Carrito, Admin, etc.).