# sso-baidu-pan
> Sso for baidu pan.

## installation
```shell
npm install -S @feizheng/sso-baidu-pan
```

## usage
```js
import ssoBaiduPan from '@feizheng/sso-baidu-pan';

const cookies = await ssoBaiduPan({
  username: 'YOUR_QQ_NO',
  password: 'YOUR_PASSWORD'
});
// set cookies
// page.setCookie(...cookies)
```

## ubuntu
- https://github.com/puppeteer/puppeteer/issues/2462
