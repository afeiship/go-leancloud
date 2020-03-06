import nx from '@feizheng/next-js-core2';
import ssoBaiduPan from '../../../dist';
class App {
  static start() {
    ssoBaiduPan({ username: '88603982', password: '---', headless: false }).then((res) => {
      console.log(res);
    });
  }
}

// main start:
App.start();
