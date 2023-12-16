1. **Función `getMovies`**:
   - `w http.ResponseWriter`: Este es un parámetro que proporciona métodos para enviar respuestas HTTP. Cualquier salida que se escriba en `w` se enviará como parte de la respuesta HTTP.
   - `r *http.Request`: Este es un puntero a la estructura que representa la solicitud HTTP entrante. Contiene detalles como la URL solicitada, el método HTTP (GET, POST, etc.), los encabezados, el cuerpo de la solicitud, y más.

2. **`w.Header().Set("Content-Type", "application/json")`**:
   - `w.Header()` devuelve los encabezados HTTP que se enviarán en la respuesta.
   - `.Set("Content-Type", "application/json")` establece el encabezado `Content-Type` en `application/json`. Esto indica al cliente (como un navegador o una aplicación API) que el contenido de la respuesta está en formato JSON.

3. **`json.NewEncoder(w).Encode(movies)`**:
   - `json.NewEncoder(w)`: Crea un nuevo codificador que escribirá la salida JSON en `w`, el `http.ResponseWriter`.
   - `.Encode(movies)`: Este método toma los datos de `movies` (que probablemente sea una variable o estructura que contiene una lista de películas) y los convierte en formato JSON. Luego, escribe este JSON en el `http.ResponseWriter` (`w`), que a su vez lo envía como la respuesta HTTP.

En resumen, la función `getMovies` es un controlador de solicitud HTTP que responde con una lista de películas en formato JSON. Cuando se llama a esta función en respuesta a una solicitud HTTP, establece el tipo de contenido de la respuesta como JSON y luego envía los datos de las películas en ese formato.


2. - Función deleteMovie:
params := mux.Vars(r):

Aquí se utiliza un paquete para analizar las variables de la ruta de la URL. mux.Vars(r) extrae los parámetros de la ruta de la solicitud HTTP. Por ejemplo, si la ruta es /movies/{id}, mux.Vars(r) obtendrá el valor de {id}. Bucle for index, item := range movies:

Este es un bucle que itera sobre la lista movies.
index es el índice actual en la lista.
item es el elemento actual de la lista (en este caso, una película).
Condición if item.ID == params["id"]:

Comprueba si el ID de la película actual (item.ID) coincide con el ID obtenido de los parámetros de la URL (params["id"]).
Eliminación de la película:

movies = append(movies[:index], movies[index+1:]...): Si se encuentra una película con el ID coincidente, esta línea elimina esa película de la lista movies.
movies[:index] es un segmento de la lista hasta, pero sin incluir, la película que se va a eliminar.
movies[index+1:] es un segmento de la lista después de la película que se va a eliminar.
append(...) combina estos dos segmentos, excluyendo efectivamente la película que se va a eliminar.
break:

Sale del bucle for una vez que se ha encontrado y eliminado la película correspondiente.
En resumen, la función deleteMovie busca en la lista movies una película con un ID específico proporcionado en la URL. Cuando encuentra la película con ese ID, la elimina de la lista y luego finaliza el procesamiento de la solicitud.
