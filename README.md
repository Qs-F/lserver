# cmd `lserver`

Next generation static server

## Design

- `-p` or `--port`

Configration of port number. Accept only int numbers

- `-d` or `--directory`

Configration of directory which will be published.

- `-i` or `--information`

Show information or not. Default is showing.

- `-pub`

Make accessiblility public or private in the internal network. Default is private.

- `-cors`

Configraiton on CORS. Default is allow from Cross Origin Access.

- `-v` or `--values`

Configration of values. This option must be setted at the last of line. All the options after this will be treated as values.

The configration style is that:

```shell
lserver -v hoge=fuga fuga=hoge g=google
```

And in file, you can write like this way:

```html
<!DOCTYPE html>
<html lang="en">
<head>
  <meta charset="UTF-8">
  <title></title>
</head>
<body>
  <p>$$hoge$$</p>
</body>
</html>
```

This will work on any ASCII text file. And if there is not matching value given, the part will be kept as it is.

- `-r` or `--replace`

The basic is the same as `-v` option; however, the way to use is different. The example is following:

```shell
lserver -r hoge=fuga fuga=hoge
```

```html
<!DOCTYPE html>
<html lang="en">
<head>
  <meta charset="UTF-8">
  <title></title>
</head>
<body>
  <p>hoge</p>
</body>
</html>
```

And the last product is following:

```html
<!DOCTYPE html>
<html lang="en">
<head>
  <meta charset="UTF-8">
  <title></title>
</head>
<body>
  <p>fuga</p>
</body>
</html>
```
