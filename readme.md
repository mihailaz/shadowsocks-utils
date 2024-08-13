# Shadowsocks utils

### ssinfo

Shows shadowsocks connection info parsed from url (ss or ssconf schema).

Usage:

```shell
ssinfo <ss or ssconf url>
```

Examples:

```shell
ssinfo ss://c29tZS1lbmNyaXB0aW9uLW1ldGhvZDpzb21lLXBhc3N3b3JkCg@example.com
ssinfo ssconf://example.com/some-path
```

```
ss://c29tZS1lbmNyaXB0aW9uLW1ldGhvZDpzb21lLXBhc3N3b3Jk@example.com
encryption method:  some-encription-method
host:  example.com
password:  some-password
```