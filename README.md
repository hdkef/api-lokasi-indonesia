# README

# API LOKASI INDONESIA
API ini dibuat untuk memperoleh lokasi di Indonesia, seperti provinsi, kota, kabupaten, dan kelurahan. API ini ditulis dalam bahasa Golang yang menyebabkan respon time menjadi cukup cepat.

# CARA INSTALL

## Dengan cloning
clone github ini dengan cara
    git clone github.com/hdkef/api-lokasi-indonesia

ubah file .env sesuai yang diinginkan : PORT dan GIN_MODE.
Port adalah 8080 secara default dan GIN_MODE adalah debug secara default,
ganti GIN_MODE mejadi release

compile kode golang ke binary dengan cara

    go build server.go

jalankan binary golang

     ./server

## Dengan docker (Soon)

jalankan perintah berikut di terminal

    docker run -d -p 1010:8080 --name api-lokasi-indonesia 081218068401/api-lokasi-indonesia

aplikasi akan berjalan pada localhost:1010 dengan nama container api-lokasi-indonesia



# CARA MENGGUNAKAN

## Mendapatkan semua provinsi

    HOST:PORT/provinces

response seperti :



## Penjelasan singkat API

    HOST:PORT/get/[objekpertama]/[bywhat]/[objectkedua]/[nilai]

objectpertama : merupakan elemen yang ingin dicari (provinsi/kota/kecamatan/kelurahan)

bywhat : merupakan jenis filter, dapat filter menggunakan ID dengan byid atau menggunakan nama dengan byname

objectkedua : merupakan elemen untuk difilter dari (provinsi/kota/kecamatan/kelurahan). misalnya
ingin mendapatkan kota yang didapat dari nama provinsi ACEH, maka menjadi /byname/ACEH

nilai : merupakan nilai dari ID atau nama dari objectkedua.

PERHATIAN!

untuk bywhat == byname maka nilai harus HURUF KAPITAL dan DIPISAH DENGAN SPASI atau %20

## Objek

provinsi ==> province

kota / kabupaten ==> city

kecamatan ==> district

kelurahan ==> village

## Mendapatkan Semua Provinsi

hit end point dengan

    HOST:PORT/provinces

Respon yang akan didapat dalam bentuk JSON
[{
    id:nomor id dari provinsi,
    name:nama dari provinsi
},...]

seperti berikut

[{"id":"11","name":"ACEH"},{"id":"12","name":"SUMATERA UTARA"},..]

## Contoh Mendapatkan Kota / Kabupaten dari ID Provinsi

hit end point dengan

    HOST:PORT/get/city/byid/province/[ID]

[ID] adalah nomor ID provinsi

contoh :

    HOST:PORT/get/city/byid/province/12

Respon yang akan didapat dalam bentuk JSON
[{
    id:nomor id dari kota / kabupaten,
    provinceid:nomor id dari provinsi,
    name:nama dari provinsi
},...]

seperti berikut

[{"id":"1201","provinceid":"12","name":"KABUPATEN NIAS"},{"id":"1202","provinceid":"12","name":"KABUPATEN MANDAILING NATAL"},...]

## Contoh Mendapatkan Kota / Kabupaten dari Nama Provinsi

hit end point dengan

    HOST:PORT/get/city/byname/province/[name]

[name] adalah nama provinsi

contoh :

    HOST:PORT/get/city/byname/province/NUSA TENGGARA BARAT

ATAU

    HOST:PORT/get/city/byname/province/NUSA%20TENGGARA%20BARAT

Respon yang akan didapat dalam bentuk JSON
[{
    id:nomor id dari kota / kabupaten,
    provinceid:nomor id dari provinsi,
    name:nama dari provinsi
},...]

seperti berikut

[{"id":"1201","provinceid":"12","name":"KABUPATEN NIAS"},{"id":"1202","provinceid":"12","name":"KABUPATEN MANDAILING NATAL"},...]