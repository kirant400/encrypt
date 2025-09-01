## Usage

Set your master key (must be 32 chars for AES-256):
```
export MASTER_KEY="32byteslongsecretkey1234567890!!"
```

## Encrypt a password:
```
go run encryptor.go myS3cretP@ss
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
go run encryptor.go decrypt k5vjYShj+3JzZPgr2X1cXyXrvPlxyaoA0h5Sdb6Tf4E=
```

→ Output:
```
myS3cretP@ss
```