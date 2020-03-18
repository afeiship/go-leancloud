/*!
 * name: @feizheng/sso-baidu-pan
 * description: Sso for baidu pan.
 * url: https://github.com/afeiship/sso-baidu-pan
 * version: 1.0.6
 * date: 2020-03-18 23:48:29
 * license: MIT
 */

"use strict";Object.defineProperty(exports,"__esModule",{value:!0}),exports.default=void 0;var _puppeteer=_interopRequireDefault(require("puppeteer")),_ssoQqPtlogin=_interopRequireDefault(require("@feizheng/sso-qq-ptlogin2")),_nextNodeCookie=_interopRequireDefault(require("@feizheng/next-node-cookie"));function _interopRequireDefault(e){return e&&e.__esModule?e:{default:e}}function asyncGeneratorStep(e,t,n,r,o,a,u){try{var i=e[a](u),s=i.value}catch(e){return void n(e)}i.done?t(s):Promise.resolve(s).then(r,o)}function _asyncToGenerator(i){return function(){var e=this,u=arguments;return new Promise(function(t,n){var r=i.apply(e,u);function o(e){asyncGeneratorStep(r,t,n,o,a,"next",e)}function a(e){asyncGeneratorStep(r,t,n,o,a,"throw",e)}o(void 0)})}}var TOKEN_RE=/"bdstoken":"(.*?)"/,DEFAULT_OPTIONS={username:null,password:null,headless:!0,stringify:!1,args:["--no-sandbox"]},_default=function(){var t=_asyncToGenerator(regeneratorRuntime.mark(function e(i){var s,c,p;return regeneratorRuntime.wrap(function(e){for(;;)switch(e.prev=e.next){case 0:return s=Object.assign({},DEFAULT_OPTIONS,i),e.next=4,_puppeteer.default.launch(s);case 4:return c=e.sent,e.next=7,c.newPage();case 7:return p=e.sent,e.next=10,p.goto("https://pan.baidu.com/");case 10:return e.next=12,p.waitForSelector(".bd-acc-qzone");case 12:return e.next=14,p.click(".bd-acc-qzone .phoenix-btn-item");case 14:return e.abrupt("return",new Promise(function(u){p.addListener("popup",function(){var t=_asyncToGenerator(regeneratorRuntime.mark(function e(t){var n,r,o,a;return regeneratorRuntime.wrap(function(e){for(;;)switch(e.prev=e.next){case 0:return e.next=2,(0,_ssoQqPtlogin.default)(t,i);case 2:return e.next=4,p.goto("https://pan.baidu.com/disk/home",{timeout:0,waitUntil:"domcontentloaded"});case 4:return e.next=6,p.cookies();case 6:return n=e.sent,e.next=9,p.content();case 9:r=e.sent,o=r.match(TOKEN_RE),c.close(),a=s.stringify?_nextNodeCookie.default.stringify(n):n,u({cookies:a,bdstoken:o[1]});case 14:case"end":return e.stop()}},e)}));return function(e){return t.apply(this,arguments)}}())}));case 15:case"end":return e.stop()}},e)}));return function(e){return t.apply(this,arguments)}}();exports.default=_default;