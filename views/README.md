# zatweb
一款结合了yaf、dva、ant-design的开发框架


``` bash
git clone https://github.com/enorzw/zatweb.git
```
安装
``` bash
npm install
```
安装php7
参考[Yaf官网](http://www.php.net/manual/zh/book.yaf.php)安装yaf
配置php.ini设置yaf的环境
``` ini
[yaf] 
yaf.environ=develop
```
编译
``` bash
npm run build
```

前端启动
``` bash
npm start
```

lint配置，添加下列配置，去掉这两个提示

* unexpected console statement no-console
* document' is not defined no-undef

``` json
{
    "env": {
        "browser": true,
        "node": true
    }, 
    "rules": {
        "no-console":0,
    }
}

``` 

帮助：
* [安装配置php71](https://github.com/enorzw/zatweb/blob/master/doc/php71.md)  
* [dva-model的详细说明](https://github.com/enorzw/zatweb/blob/master/doc/model.md) 
* [component的详细说明](https://github.com/enorzw/zatweb/blob/master/doc/component.md)
* [路由的详细说明](https://github.com/enorzw/zatweb/blob/master/doc/router.md) 

参考链接: 
* dva 
    * [dva官方github](https://github.com/dvajs/dva)
    * [dva的由来](https://github.com/sorrycc/blog/issues/6)
    * [dva知识图谱](https://github.com/dvajs/dva-knowledgemap)
* php 
    * [yaf中文文档](http://www.php.net/manual/zh/book.yaf.php)
* react
    * [react官方文档](https://facebook.github.io/react/docs/hello-world.html)
    * [react中文文档](https://hulufei.gitbooks.io/react-tutorial/content/introduction.html) 
* router
    * [react-router官方中文文档](http://react-guide.github.io/react-router-cn/index.html)
    * [react-router中文文档](http://618cj.com/react-router4-0%E8%B7%AF%E7%94%B1%E4%B8%AD%E6%96%87%E6%96%87%E6%A1%A3api/)
* redux
    * [react-redux](http://www.redux.org.cn/docs/react-redux/api.html#connectmapstatetoprops-mapdispatchtoprops-mergeprops-options)
    * [redux中文文档](http://cn.redux.js.org/index.html)
    * [redux-saga中文文档](http://leonshi.com/redux-saga-in-chinese/docs/introduction/BeginnerTutorial.html?utm_source=tuicool&utm_medium=referral)
* 前端
    * [ant-design中文文档](https://ant.design/docs/react/introduce-cn)
    * [es6 Generator函数](http://es6.ruanyifeng.com/#docs/generator)
* roadhog
    * [roadhog](https://github.com/sorrycc/roadhog#%E9%85%8D%E7%BD%AE)