# README

# API LOKASI INDONESIA
API ini dibuat untuk memperoleh lokasi di Indonesia, seperti provinsi, kota, kabupaten, kelurahan dan sebagainya. API ini ditulis dalam bahasa Golang yang menyebabkan respon time jauh lebih cepat daripada
PHP dan sejenisnya.

# CARA MENGGUNAKAN

## Mengatur PORT
PORT dapat diubah dengan mengubah nilai PORT pada file .env,
default 8080
## Mendapatkan Semua Provinsi

    hit end point dengan

    <q>HOST:PORT/provinces</q>

    Respon yang akan didapat dalam bentuk JSON
    [{
        id:nomor id dari provinsi,
        name:nama dari provinsi
    },...]

    seperti berikut

    [{"id":"11","name":"ACEH"},{"id":"12","name":"SUMATERA UTARA"},..]

## Mendapatkan Kota / Kabupaten dari Province-ID

    hit end point dengan

    <quote>HOST:PORT/get/city/byid/province/[ID]</quote>

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

## Mendapatkan Kota / Kabupaten dari Nama Provinsi

    hit end point dengan

    <quote>HOST:PORT/get/city/byname/province/[name]</quote>

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