## Usage

Set your master key (must be 32 chars for AES-256):
```
export MASTER_KEY="32byteslongsecretkey1234567890!!"
```

## Usage 
```
docker build -t encrypt:v1.0 .
docker save -o encrypt_image.tar encrypt:v1.0
docker run --rm --env-file=./.env encrypt:v1.0 ./encrypt encrypt <plaintext>
docker run --rm --env-file=./.env encrypt:v1.0 ./encrypt decrypt <ciphertext_base64>
```

## Encrypt a password:
```
go run main.go encrypt <plaintext>
eg:
go run encryptor.go encrypt myS3cretP@ss
```

Copy the output into your config.yaml under password_enc.
```
endpoints:
  - name: service1
    target_url: https://api.service1.com/v1
    auth_type: bearer
    auth_api: https://auth.service1.com/token
    username: user1
    password_enc: "<PASTE OUTPUT HERE>"
```

## Decrypt it back (for testing):

```
go run main.go decrypt <ciphertext_base64>
eg:
go run encryptor.go decrypt k5vjYShj+3JzZPgr2X1cXyXrvPlxyaoA0h5Sdb6Tf4E=
```
