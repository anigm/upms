import 'babel-polyfill';
import "intl";
import dva from 'dva';
import { browserHistory } from 'dva/router';
import { message } from 'antd';
import './index.less';


// 1. Initialize
const app = dva({
  history: browserHistory,
  onError(e) {
    message.error(e.message, /* duration */3);
  },
});

global.DvaApp = app;

// 2. Plugins
// app.use({});

// 3. Model
// app.model(require('./models/example'));

// 4. Router
app.router(require('./router'));

// 5. Start
app.start('#root');
