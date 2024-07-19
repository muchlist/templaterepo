# templaterepo template

Architecture templaterepo hexagonal golang.  
Pada mulanya template ini diperuntukkan sebagai templaterepo, ide dasarnya adalah kita dapat memakai ulang code pada beberapa aplikasi.  
Karena adanya pemakaian ulang code, yang mana dalam tanda kutip `bisa saling memakai`, maka diperlukan suatu cara agar kita dapat menghindari import cycle dependency. Maka inilah caranya.

## Overview

Tujuan :  

- memiliki cara yang seragam dalam membangun sebuah program.  
- code yang decouple antar module
- menggunakan prinsip `hexagonal` architexture.
- menghindari import cycle dependency meskipun ada banyak modul yang saling terhubung.
- code yang testable terutama pada bagian logic.  

Fokus Hexagonal arch adalah pemisahan/pemurnian antara core logic (core-nya) terhadap dependency luar. core harus bersih dan hanya terdiri dari standart library dan wrapper code yang dibangun hanya pada repository ini.

Pemakaian istilah core bisa diganti dengan service dan port bisa diganti dengan storer, pada dasarnya port ini hanyalah sebuah interface.

## Content

- [Quick start](#quick-start)
- [Project structure](#project-structure)

## Quick start

Local development:  

```sh

# 1. pastikan ketersediaan dependency seperti database dll.
# 2. menjalankan aplikasi dengan makefile (lihat file Makefile)
$ make run/api/user

# command tersebut akan mengeksekusi
$ go run ./app/api-user

```  

## Aturan penulisan code

### Layer business

1. Pada layer business, terutama bagian `service (core)` berkomunikasi dengan mengandalkan interface menggunakan prinsip `dependency inversion`. Penjelasannya adalah seperti gambar berikut.  
![interface](documents/interface.png)  
contoh bisa dilihat pada domain user `business/user` yang terhubung dengan `business/notifserv`. Penerapan dependensinya dapat dilihat pada `app/api-user/routing.go`.  
Dengan metode seperti ini maka tidak akan ada import cycle dependency dan code menjadi sangat decouple antar domain.  
2. memiliki konstruktor yang mengembalikan tipe konkrit.  
3. parameter di constructor merupakan `interfaces` dan bukan konkrit object.

### Idiom

- Penamaan Interface ([https://go.dev/doc/effective_go#interface-names](https://go.dev/doc/effective_go#interface-names))
- Agar kita bisa langsung tau sebuah tipe itu adalah interface, maka namakan tipe tersebut menggunakan akhiran -er -tor. 
- Contoh Writer, Reader, Assumer, Saver, Reader, Generator.
- Contoh pada project yang menggunakan 3 layer : UserServiceAssumer, UserStorer, UserSaver, UserLoader.

### Rules lainnya

- Uber Style Guide ([https://github.com/uber-go/guide/blob/master/style.md](https://github.com/uber-go/guide/blob/master/style.md)). Nantinya kita buat sendiri untuk menimpa rules uber. (misalnya penamaan interface).
- Konfigurasi file hanya bisa diakses di main.go saja. layer lain yang ingin mengakses konfigurasi harus menerimanya melalui parameter fungsi. 
- Konfigurasi harus memiliki nilai default yang setidaknya bisa berjalan di environment local. nilai default dapat ditimpa env, dan command line.
- Error harus dihandle hanya 1 kali dan tidak boleh di abaikan, Maksudnya adalah antara di konsumsi atau di return, tidak boleh keduanya sekaligus. contoh konsumsi : menulis error pada log, contoh return : mereturn error apabila error tidak nil.
- Jangan mengekspose variable yang hidup didalam package, ganti menjadi private dan gunakan func Public sebagai gantinya.
- Ketika suatu code menjadi banyak dipakai ditempat lain, jika satu package buatlah helper.go, jika dipakai di beda package buatlah package baru (misalnya `/user/ipkg/error_parser.go`), jika pemakaiannya semakin meluas masukkan di `pkg` (misalnya `pkg/slicer/slicer.go`, `pkg/datastructure/ds.go`, `pkg/errr/custom_error.go` dan sebagainya).


## Project structure

### `app/`

Tempat untuk code yang tidak dapat dipakai ulang, spesifik untuk input output, menghidupkan dan mematikan aplikasi. Folder app ini dapat dijadikan acuan sebagai tempat dimulainya program ketika dijalankan.

Folder app adalah sama dengan folder `cmd` pada kebanyakan projek lainnya. Dinamakan app karena posisi folder akan berada diatas (yang mana dirasa cukup bagus) dan cukup mewakili fungsi folder.  

Contoh app restfull api menggunakan fiber dapat diliat pada `app/api-user`  

Alih alih menggunakan framework seperti cobra untuk memilih menjalankan aplikasi, kita menggunakan cara paling simple seperti `go run ./app/api-user` untuk menjalankan aplikasi api user dan `go run ./app/consumer-user` untuk menjalankan aplikasi kafka consumer.


### `pkg/`

pkg layer merupakan tempat untuk meletakkan package lainnya yang dapat dipakai ulang dan dipakai dimana saja.
umumnya untuk meletakkan pondasi yang tidak terikat dengan module problem, seperti logger, web framework, data structure secara general, tempat untuk meletakkan library yang sudah di wrap agar mudah di mock.

App layer dapat mengimport pkg, business layer dapat mengimport pkg.  

penggunaan `pkg/` sebagai folder penampung code terakhir yang pada mulanya kita tidak yakin dimana ingin menaruhnya, terbukti meningkatkan percepatan kita dalam menulis code. `taruh dimana nih, bingung.... di pkg ajaa ...` 

### `business`

Berisi code yang terkait dengan logika bisnis, problem bisnis, data bisnis.

### `business/[domain]/*`

Didalam perdomain bisnis terdapat Service (atau core -- dalam istilah hexagonal) yang harus kita jaga kemurniannya dari library luar, termasuk juga layer untuk mengakses data persistance (repo) dan interface2 yang berlaku sebagai port.  

### `business/[domain]/[subfolder]`

Pada bisnis domain memungkinkan untuk di restruktur menjadi beberapa folder apabila code menjadi sangat rumit jika digabungkan. Contoh `business/xcomplex`.




## Tools

### `Makefile`

Makefile berisi command untuk membantu proses menjalankan aplikasi dengan cepat karena tidak harus mengingat semua command yang panjang. Berfungsi seperti alias. Caranya adalah dengan menuliskan cmd di file Makefile seperti contoh berikut.

Baris teratas adalah comment yang akan muncul ketika memanggil helper.  
`.PHONY` adalah penanda agar terminal tidak menganggap command makefile sebagai akses ke file.  
`run/tidy:` adalah alias untuk cmd yang ada didalam nya.

```sh
## run/tidy: run golang formater and tidying code
.PHONY: run/tidy
run/tidy:
  @echo 'Tidying and verifying module dependencies...'
  go mod tidy
  go mod verify
  @echo 'Formatting code...'
  go fmt ./...
```

### `pre-commit`

Disarankan menggunakan pre-commit (doc : [pre-commit]("https://pre-commit.com/")).  

  ```bash
  // init
  pre-commit install

  // precommit akan di trigger setiap commit

  // manual
  pre-commit run --all-files

  ```
