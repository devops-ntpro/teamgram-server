# Teamgram - Open source [mtproto](https://core.telegram.org/mtproto) server written in golang
> open source mtproto server implemented in golang with compatible telegram client.

English | [简体中文](readme-cn.md)

### Introduce
Open source [mtproto](https://core.telegram.org/mtproto) server written in golang

### Architecture
![Architecture](docs/image/architecture-001.png)

### Documents
[Diffie–Hellman key exchange](docs/dh-key-exchange.md)

[Creating an Authorization Key](docs/Creating_an_Authorization_Key.md)

[Mobile Protocol: Detailed Description (v.1.0, DEPRECATED)](docs/Mobile_Protocol-Detailed_Description_v.1.0_DEPRECATED.md)

[Encrypted CDNs for Speed and Security](docs/cdn.md) Translate By [@steedfly](https://github.com/steedfly)

### Quick start with Docker
> TODO...

### [Centos 9 Stream Build and Install](docs/install-centos-9.md) [@A Feel]

### Manual Build and Install
#### Depends
- **mysql5.7**
- [redis](https://redis.io/)
- [etcd](https://etcd.io/)
- [kafka](https://kafka.apache.org/quickstart)
- [minio](https://docs.min.io/docs/minio-quickstart-guide.html#GNU/Linux)
- [ffmpeg](https://www.johnvansickle.com/ffmpeg/)

#### Install Teamgram
- Get source code　
```
git clone https://github.com/devops-ntpro/teamgram-server.git
cd teamgram-server
```
- grant permissions to access without password
```
1. get mysql config file
    sudo docker cp teamgram-server_mysql_1:/etc/mysql/mysql.cnf .
2. add
    [mysqld]
    skip-grant-tables
3. save changes
    sudo docker cp mysql.cnf teamgram-server_mysql_1:/etc/mysql
4. restart mysql
    sudo docker exec -it teamgram-server_mysql_1 bash
    /etc/init.d/mysql restart
```  
- init database
```
1. create database teamgram
2. init teamgram database
   mysql -h127.0.0.1 -P3306 -uroot teamgram < teamgramd/sql/teamgram2.sql
   mysql -h127.0.0.1 -P3306 -uroot teamgram < teamgramd/sql/migrate-20220321.sql
  ...  
```

- init minio buckets, bucket names:
  - `documents`
  - `encryptedfiles`
  - `photos`
  - `videos`

- Build
```
cd scripts
./build.sh
```
- Create log dirs
```
cd ./teamgramd/logs
mkdir authsession bff biz dfs gateway idgen media msg session status sync
```
    
- Run
```
cd ./teamgramd/bin
./runall2.sh
```


### Compatible clients
**Important**: default signIn verify code is **12345**

[Android client for Teamgram](clients/teamgram-android.md)

[iOS client for Teamgram](clients/teamgram-ios.md)

[tdesktop for Teamgram](clients/teamgram-tdesktop.md)

## Feedback
Please report bugs, concerns, suggestions by issues, or join telegram group [Teamgram中文社区](https://t.me/cnteamgram) Or [Teamgram](https://t.me/enteamgram) to discuss problems around source code.

## Notes
If need enterprise edition, please PM the **[author](https://t.me/benqi)**

## Give a Star! ⭐

If you like or are using this project to learn or start your solution, please give it a star. Thanks!
