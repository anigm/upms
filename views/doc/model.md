# dva-model详细说明
使用dva时，我们可以使用dva的工具自动生成model，例如我们要生成一个计数的model
``` bash
> dva g model count
```
这个命令会做两件事：

1.自动在models文件夹里生成count.js文件，如下代码所示，这是dva中model的标准结构：
``` javascript
export default {
  namespace: 'count',
  state: {},
  reducers: {},
  effects: {},
  subscriptions: {},
};
```
2.在index.js入口文件中加载model：
``` javascript
app.model(require("./models/count")); 
```
* namespace：model的命令空间
* state：model的state
* reducers：在这里定义修改state的方法
* effects：异步处理的方法，可以通过dispatch调用reducer方法
* subscriptions：事件的订阅
