# img2gray
This CLI generates gray image.  
.png, .jpg(.jpeg), are only supported.

## Install

```
go get github.com/Le0tk0k/img2gray
```

## Option

|  Flag  |  Description  | Default |
| ---- | ---- | --- |
|  -r  |  Remove sorce file | false |

## Usage

```
$ # default, Do not delete sorce file
$ ./main sample.jpeg

$ Delete sorce file
$ ./main -r sample.jpeg
```
