# Generador de textos para el proyecto de word count

Este programa tiene como objetivo generar datos de prueba para el proyecto de word count. Es capaz de generar textos
aleatorios y de revisar la salida esperada de los programas de los estudiantes.

## Uso del programa

El programa soporta los siguientes argumentos:

- `-mode`: Modo de operación del programa. Puede ser `generator` o `validator`.
- `-seed`: Semilla para la generación de textos aleatorios. Por default es 0.
- `-words`: Número de palabras únicas que se desean generar. Por default es 10000.
- `-size`: Cantidad total de palabras que contendrá el texto generado. Por default es un millón.

## Ejemplos de uso

### Generar un texto aleatorio

```bash
go run wordcount.go -mode generator -seed 1234 -words 10000 -size 1000000 > input.txt
```

### Validar la salida de un programa

```bash
go run wordcount.go -mode validator < input.txt
```

### Probar su proyecto

```bash
go run wordcount.go -mode generator -seed 1234 -words 10000 -size 1000000 | ./suproyecto | go run wordcount.go -mode validator -seed 1234 -words 10000 -size 1000000
```

### Generar un archivo de salida esperada con Bash
Es posible generar un archivo de salida esperada con el siguiente comando:

```bash
cat entrada.txt | tr '[:upper:]' '[:lower:]' | tr -cs '[:alnum:]' '\n' | sort | uniq -c | sort -k2 | awk '{print $2 ": " $1}' > salida_esperada.txt
```

El comando anterior funciona de la siguiente manera:
1. `cat entrada.txt`: Lee el archivo de entrada.
2. `tr '[:upper:]' '[:lower:]'`: Convierte todas las letras a minúsculas.
3. `tr -cs '[:alnum:]' '\n'`: Convierte todos los caracteres no alfanuméricos en saltos de línea.
4. `sort`: Ordena las palabras.
5. `uniq -c`: Cuenta las palabras únicas.
6. `sort -k2`: Ordena las palabras por orden alfabético.
7. `awk '{print $2 ": " $1}'`: Imprime la palabra y su frecuencia.
8. `> salida_esperada.txt`: Guarda la salida en un archivo.
