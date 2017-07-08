# 路由详细说明

dva的路由模块为react-router，官方有[中文文档](http://react-guide.github.io/react-router-cn/index.html)

## 1.入口
在dva的入口文件里，我们可以设置history
``` javascript
import dva from 'dva';
import { browserHistory } from 'dva/router';

const app = dva({
  history: browserHistory,
});
```
其中history有三种：
* <font color='#dd7e6b'>browserHistory</font>：是使用react-router推荐的history。它使用浏览器中的 History API 用于处理 URL，创建一个像example.com/some/path这样真实的 URL 。
* <font color='#dd7e6b'>hashHistory</font>：Hash history 不需要服务器任何配置就可以运行，如果你刚刚入门，那就使用它吧。但是不推荐在实际线上环境中用到它
* <font color='#dd7e6b'>memoryHistory</font>：Memory history 不会在地址栏被操作或读取。

最后在入口文件里,通过下面代码就可以加载路由：
``` javascript
app.router(require('./router'));
```
## 2.路由设置
上面介绍了入口，最后说的"./router"就是路由定义的地方了。其中这个require返回的是一个函数，这个函数返回的就是定义路由的Router对象。

路由设置有两种方式：
* 使用JSX语法
``` html
function ConfigRouter(){
  return (
    <Router>
      <Route path="/" component={App}>
        {/* 当 url 为/时渲染 Dashboard */}
        <IndexRoute component={Dashboard} />
        <Route path="about" component={About} />
        <Route path="inbox" component={Inbox}>
          <Route path="messages/:id" component={Message} />
        </Route>
      </Route>
    </Router> 
  )
}
```
* 使用json对象
``` javascript
function ConfigRouter(){
  const routeConfig = [
    { path: '/',
      component: App,
      indexRoute: { component: Dashboard },
      childRoutes: [
        { path: 'about', component: About },
        { path: 'inbox',
          component: Inbox,
          childRoutes: [
            { path: '/messages/:id', component: Message },
            { path: 'messages/:id',
              onEnter: function (nextState, replaceState) {
                replaceState(null, '/messages/' + nextState.params.id)
              }
            }
          ]
        }
      ]
    }
  ] 
  return <Router routes={routeConfig} /> 
}
```
使用json对象还有两个方法：
* 同步：即是纯粹的json对象
* 异步：也可以叫做动态路由，即使用getComponents，getIndexRoute，getChildRoutes异步函数，并且在此基础上可以配合require.ensure函数来实现按需加载，代码如下：
``` javascript
const CourseRoute = {
  path: 'course/:courseId',

  getChildRoutes(location, callback) {
    require.ensure([], function (require) {
      callback(null, [
        require('./routes/Announcements'),
        require('./routes/Assignments'),
        require('./routes/Grades'),
      ])
    })
  },

  getIndexRoute(location, callback) {
    require.ensure([], function (require) {
      callback(null, require('./components/Index'))
    })
  },

  getComponents(location, callback) {
    require.ensure([], function (require) {
      callback(null, require('./components/Course'))
    })
  }
}
```