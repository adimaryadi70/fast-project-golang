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
protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative common\proto\transaction\transaction.proto


