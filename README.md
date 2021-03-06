# SuperListsAPI
  <div>
    <img align="right" width="350px" src="https://i.ibb.co/5nZgKFf/golang-gopher-chamberos.png" alt="gophers-chamberos" style="max-width:100%;">
  </div>
SuperLists es un proyecto que intenta ayudar a dejar de utilizar papel para poder crear listas. Cualquier tipo de lista es bienvenida: 
<br>
<br>

* Lista de compras
* Lista de tareas
* Cualquier tipo de listas!

<br>

### Listas de una sola persona?
No! Las listas también pueden compartirse para poder ir completándolas en conjunto. Las tareas finalizadas pueden marcarse como tal y como próximo feature se enviarán notificaciones cuando una tarea haya sido marcada como finalizada.

<br>
<br>
Confiamos en las tecnologías open-source y solo por mencionar algunas de las empleadas para este proyecto podemos mencionar:

<br>
<br>

* Go
* Gin Gonic
* Gorm
* Postgresql
* JWT (Golang package)



## Sobre el API de SuperLists

### Endpoints
<br>

#### Auth endpoints

Estos endpoints son utilizados para lo que refiere a authentication & authorization. Utilizan algunos servicios/repositorios externos, como los de User.

*`POST` - `v1/auth/login` Login

*`POST` - `v1/auth/signup` Create account

<br>

#### List endpoints

Endpoints destinados para el handleo del CRUD de Lists.

*`POST` - `v1/lists/` Create list

*`PUT`  - `v1/lists/{id}` Update list

*`GET`  - `v1/lists/{id}` Get list

*`GET` - `v1/lists/` Get lists

*`DELETE` - `v1/lists/{id}` Delete list

*`POST` - `v1/lists/joinList/{listID}` Join to other user list with an invitation code.

<br>

#### User Lists endpoints

Endpoints empleados para el CRUD de UserLists, entidad que relaciona los Users con sus Lists.

*`POST` - `v1/userLists/` Create user_list

*`GET`  - `v1/userLists/{id}` Get user_list

*`GET` - `v1/userLists/` Get user_lists by user ID

*`DELETE` - `v1/userLists/{id}` Delete user_list

<br>

#### List items endpoints
//TODO actualizar este API doc
Los endpoints de List Items, implementan el CRUD del objeto que modela una tarea, la cual tiene asociada una lista a la que corresponde, entre otras cosas.

*`POST` - `v1/listItems/` Create list_item

*`PUT`  - `v1/listItems/{id}` Update list_item

*`GET`  - `v1/listItems/{id}` Get list_item
   
*`DELETE` - `v1/listItems/{id}` Delete list_item

<br>

#### List Management endpoints
* Coming soon
<br>


#### Profile endpoints
* Coming soon
<br>

 