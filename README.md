# Go Beta template

Architecture template for beta squad.


## Overview
Tujuan dari dibuatnya template ini adalah agar:
- memiliki cara yang seragam dalam membangun sebuah service/program.
- mengatur code agar code yang sudah ditulis sebelumnya tidak berubah menjadi spaghetti code.
- mengatur semua potensi kebingungan yang bisa timbul seiring berkembangnya repository dengan cara menulis segala macam keputusan yang diambil atas kebingungan tersebut.

Menggunakan prinsip prinsip pemrograman golang yang dikombinasikan dengan clean architecture dan hexagonal architexture.

Fokus clean arch / Hexa arch adalah pemisahan antara bussinesss logic (core-nya) dengan layer-layer aplikasi lainnya dan dependency pihak ketiga, mudahnya testing, mudahnya penggantian library pihak ketiga.

## Content
- [Quick start](#quick-start)
- [Project structure](#project-structure)

## Quick start
Local development:  
```sh
# 1. pastikan ketersediaan dependency seperti database dll.
# 2. menjalankan aplikasi dengan makefile (lihat file Makefile)
$ make run/api/user
```


## Project structure
### `app/`
Tempat untuk code yang tidak dapat dipakai ulang, spesifik untuk input output, menghidupkan dan mematikan aplikasi. Folder app ini dapat dijadikan acuan sebagai tempat dimulainya program ketika dijalankan.

Folder app adalah sama dengan folder cmd pada kebanyakan projek lainnya. Dinamakan app karena posisi folder akan berada diatas (yang mana dirasa cukup bagus) dan cukup mewakili fungsi folder.


### `pkg/`
pkg layer merupakan tempat untuk meletakkan package lainnya yang dapat dipakai ulang tidak hanya untuk repository ini, namun juga dapat dipakai repository yang lain.
umumnya untuk meletakkan pondasi yang tidak terikat dengan bussiness problem, seperti logger, web framework, data structure secara general, wrapper library.

App layer dapat mengimport pkg, bussiness layer dapat mengimport pkg.


### `bussiness/core/`
bussiness/core berisi busines logic, busines problem, Kita mengenal ini dengan sebutan Usecase atau Service, termasuk juga layer untuk mengakses data persistance (repo) dan api-luar. Pada repository dengan jumlah aplikasi sedikit disarankan untuk membuat struktur yang di group berdasarkan layer untuk menghindari `import cycle dependency`.  
misalnya :
```
    bussiness 
      - usecase
        - user.go
        - product.go
      - store
        - user.go
        - product.go
      - model
        - user.go
        - product.go
```
Dengan begini memungkinkan untuk saling import usecase (bussiness logic). Namun struktur seperti ini tidak berdasarkan pembagian berdasarkan domain dan akan menjadi rumit ketika usecase sudah banyak.

Opsi lain adalah vertical layer seperti :
```
    bussiness 
      - user
        - model.go
        - store.go
        - storer.go
        - model.go
      - product
        - model.go
        - store.go
        - storer.go
        - model.go
```
yang mana menjadi tidak memungkinkan untuk saling mengimport. (misalnya user membutuhkan product dan product membutuhkan user).   

### `bussiness/sys/`
bussiness/sys berisi logic yang berhubungan dengan system seperti auth, database, metric, validator, middleware yang mana package ini biasanya hanya diperlukan terbatas pada bussines layer saja.  
Misalnya pada `bussiness/sys/cjwt` kemungkinan claimsnya akan berbeda dengan aplikasi lain. `bussiness/sys/mid` adalah middleware yang di costumisasi hanya untuk aplikasi ini.  
Biasanya akan ada keraguan untuk meletakkan suatu code di `bussiness/sys/` atau di `pkg` yang mana perlu di diskusikan terlebih dahulu.

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
