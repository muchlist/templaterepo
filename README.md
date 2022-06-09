# Monorepo template

Architecture monorepo hexagonal golang.


## Overview

Tujuan dari dibuatnya template ini adalah agar:  

- memiliki cara yang seragam dalam membangun sebuah service/program.  
- code yang decouple antar module dengan menggunakan prinsip pada hexagonal architexture.
- menghindari import cycle dependency meskipun ada banyak modul yang saling terhubung.
- code yang selalu bisa di test terutama pada layer core (dikenal juga sebagai usecase atau service)  

Fokus Hexa arch adalah pemisahan/pemurnian antara core logic (core-nya) terhadap dependency luar. core harus bersih dan hanya terdiri dari standart library dan wrapper code yang dibangun pada repository ini.  

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

1. Pada layer business, terutama core dan repo/storer/client (layer yang menerima data dari luar) dihubungkan dengan interface dengan menggunakan prinsip `dependency inversion principle` yang mana menggunakan golang style seperti gambar berikut.  
![interface](docs/interface.png)  

contoh bisa dilihat pada domain user `internal/business/user` yang terhubung dengan `internal/business/notifserv` dan penerapan dependensi pada `app/api-user/routing.go`.  
Dengan metode seperti ini maka tidak akan ada import cycle dependency dan code menjadi sangan decouple antar domain.
2. memiliki konstruktor yang mengembalikan tipe konkrit.
3. parameter di constructor di wajibkan berupa `interfaces` bukan konkrit object.


## Project structure

### `app/`

Tempat untuk code yang tidak dapat dipakai ulang, spesifik untuk input output, menghidupkan dan mematikan aplikasi. Folder app ini dapat dijadikan acuan sebagai tempat dimulainya program ketika dijalankan.

Folder app adalah sama dengan folder cmd pada kebanyakan projek lainnya. Dinamakan app karena posisi folder akan berada diatas (yang mana dirasa cukup bagus) dan cukup mewakili fungsi folder.  

Contoh app restfull api menggunakan fiber dapat diliat pada `app/api-user` 


### `pkg/`

pkg layer merupakan tempat untuk meletakkan package lainnya yang dapat dipakai ulang tidak hanya untuk repository ini, namun juga dapat dipakai repository yang lain.
umumnya untuk meletakkan pondasi yang tidak terikat dengan module problem, seperti logger, web framework, data structure secara general, wrapper library.

App layer dapat mengimport pkg dan ipkg, business layer dapat mengimport pkg dan ipkg.

### `internal/business`

internal/module berisi code yang terkait dengan logika bisnis, problem bisnis, data bisnis.

### `internal/business/[domain]/*`

Didalam perdomain bisnis terdapat Core (yang mana harus kita jaga kemurniannya dari library luar), termasuk juga layer untuk mengakses data persistance (repo) dan interface2 yang berlaku sebagai port.  

### `internal/business/[domain]/[subfolder]`

Pada bisnis domain memungkinkan untuk di restruktur menjadi beberapa folder apabila code menjadi sangat rumit jika digabungkan. Contoh `internal/core/complex`

### `internal/ipkg/`

internal/ipkg/ berisi library internal seperti auth, database, metric, validator, middleware yang mana package ini biasanya hanya diperlukan terbatas pada bussines layer saja.  
Misalnya pada `internal/ipkg/cjwt` kemungkinan claimsnya akan berbeda dengan aplikasi lain. `internal/ipkg/mid` adalah middleware yang di costumisasi hanya untuk aplikasi ini saja.  
Biasanya akan ada keraguan untuk meletakkan suatu code di `internal/ipkg/` atau di `pkg` yang mana perlu di diskusikan terlebih dahulu.

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
