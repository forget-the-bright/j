# j
Java Version  Manager

![image](https://user-images.githubusercontent.com/56473277/228444767-6cc15881-69a6-4ed1-8b30-f7c19413fde0.png)


## 安装
  [下载](https://github.com/forget-the-bright/j/releases) 下载自己需要的版本, 到自己自定义的目录 修改可执行文件名称为j
  
  默认文件下载安装在用户目录下 ```.j```目录，目录下  ```versions```, ```downloads```, ```java```  分别是本地安装目录，安装包下载目录，当前使用的java版本目录 

  将 JAVA_HOME 配置为 ```USER_HOME\.j\java```  

  指定安装目录需要 添加环境变量 ```J_HOME```
## 命令

### 列出

列出所有可安装版本
```
j ls-remote
```
![image](https://user-images.githubusercontent.com/56473277/228444893-1ae5779e-74a2-4884-9c7d-09aa533d644e.png)

列出本地安装版本
```
j ls
```
![image](https://user-images.githubusercontent.com/56473277/228177030-defae4d2-77ba-4ded-9598-953330ac6cd8.png)

### 下载
```
j install 8
```
![image](https://user-images.githubusercontent.com/56473277/228178315-9491c998-c839-441f-8d5e-78578273f57c.png)
![image](https://user-images.githubusercontent.com/56473277/228178398-59382b04-ae1d-443a-9e20-ba41ab65ae7c.png)

### 切换版本
```
j use 17
```
![image](https://user-images.githubusercontent.com/56473277/228178562-509f752f-134e-4b44-b220-26ed7fc9b33c.png)




