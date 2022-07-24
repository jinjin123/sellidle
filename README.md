# sell your idle
In China,there Cloud vps is not cheapest and bandwidth too small, and there has more idle home server, I can provide the Saas service for some guys what they want to test sth like k8s cluster etc.
```
nginx.conf see remote ip for hacker 
http block {
       set_real_ip_from  192.168.1.0/24;
        set_real_ip_from  127.0.0.1;
        real_ip_header    X-Forwarded-For;
        real_ip_recursive on;
...
}


server block  {
 outside location block 
        auth_basic "Restricted";
        auth_basic_user_file /etc/nginx/.htpasswd;

	location block {

                proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
		proxy_set_header Host $host;
                proxy_set_header X-Real-IP $remote_addr; 
		proxy_set_header REMOTE-HOST $remote_addr;
		...
	}
...
}

```
# backend
```
go run main.go
```
# fronted
```
npm run start
```

![demo2](https://github.com/jinjin123/sellidle/blob/main/server.png)
![demo1](https://github.com/jinjin123/sellidle/blob/main/bindport.png)
![demo](https://github.com/jinjin123/sellidle/blob/main/sellvps.png)

