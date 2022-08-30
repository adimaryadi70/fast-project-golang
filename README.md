# fast-project-golang
# Version 1.1.2
sample services sederhana
- login jwt
- db query gorm
- database mysql
- database postgres
- info sesi token di file config go terdapat token_duration itu adalah per menit token_secret boleh di ganti bebas 


Template ini semoga bermamfaat bagi pemula

data base mengunakan mysql 
auto migrasi mengunakan gorm 
Standar MVC
KODE RC response 

KODE 00 suksess
KODE 08 Gagal
KODE 01 Unauthorize

generate Proto GRPC
protoc --proto_path=D:\DATA\Backend\fast-project-golang\common\proto D:\DATA\Backend\fast-project-golang\common\proto\transaction\transaction.proto --go-grpc_ou
t=./proto/transaction

