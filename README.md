# Módulo de Clientes (Go)

Proyecto de ejemplo para el Hito 2: módulo CRUD de clientes en Go.

## Estructura

- `cmd/api/main.go` — entrada principal del servidor
- `internal/models/` — definición del struct `Client`
- `internal/handlers/` — handlers HTTP
- `internal/storage/` — almacenamiento en memoria

## Endpoints

Base: `http://localhost:8080/api/v1/clientes`

- `GET /` — lista todos los clientes (200)
- `POST /` — crea cliente (201)
- `GET /{id}` — obtiene cliente por id (200 / 404)
- `PUT /{id}` — actualiza cliente (200 / 404 / 400)
- `DELETE /{id}` — elimina cliente (204 / 404)

## Validación

- `nombre` es obligatorio en `POST` y `PUT`
- JSON inválido responde `400`
- Id no encontrado responde `404`

## Ejecutar

```powershell
cd "C:\Users\AlexGC\OneDrive\Escritorio\Modulo de clientes"
go mod tidy
go run ./cmd/api
```

## Ejemplos de `curl`

Crear cliente:

```bash
curl -X POST http://localhost:8080/api/v1/clientes/ \
  -H "Content-Type: application/json" \
  -d '{"nombre":"Juan Pérez","correo":"juan@example.com","telefono":"1234567890","direccion":"Calle Falsa 123","rfc":"JUAP800101"}'
```

Listar clientes:

```bash
curl http://localhost:8080/api/v1/clientes/
```

Obtener cliente por ID:

```bash
curl http://localhost:8080/api/v1/clientes/1
```

Actualizar cliente:

```bash
curl -X PUT http://localhost:8080/api/v1/clientes/1 \
  -H "Content-Type: application/json" \
  -d '{"nombre":"Juan Pérez","correo":"juan.nuevo@example.com","telefono":"1234567890","direccion":"Calle Falsa 123","rfc":"JUAP800101"}'
```

Eliminar cliente:

```bash
curl -X DELETE http://localhost:8080/api/v1/clientes/1
```

## Checklist para el Hito 2

- [x] `cmd/api/main.go` con router Chi y subrouter `/api/v1`
- [x] `internal/models/` con struct `Client` y tags JSON
- [x] `internal/handlers/` con los 5 handlers CRUD
- [x] `internal/storage/` con almacenamiento en memoria
- [x] Validación básica de `nombre` obligatorio
- [x] Status codes correctos: `200`, `201`, `404`, `400`, `204`
- [x] `DOCUMENTO_TECNICO.md` añadido
- [x] `curl_commands.sh` con ejemplos de demo
- [x] `.gitignore` para archivos binarios y temporales
