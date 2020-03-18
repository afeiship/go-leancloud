import puppeteer from 'puppeteer';
import ssoQqPtlogin2 from '@feizheng/sso-qq-ptlogin2';
import NxNodeCookie from '@feizheng/next-node-cookie';

const TOKEN_RE = /"bdstoken":"(.*?)"/;
const DEFAULT_OPTIONS = {
  username: null,
  password: null,
  headless: true,
  stringify: false,
  args: ['--no-sandbox']
};

export default async (inOptions) => {
  const options = Object.assign({}, DEFAULT_OPTIONS, inOptions);
  const browser = await puppeteer.launch(options);
  const page = await browser.newPage();
  await page.goto('https://pan.baidu.com/');
  await page.waitForSelector('.bd-acc-qzone');
  await page.click('.bd-acc-qzone .phoenix-btn-item');
  return new Promise((resolve) => {
    page.addListener('popup', async (popup) => {
      await ssoQqPtlogin2(popup, inOptions);
      await page.goto('https://pan.baidu.com/disk/home', { timeout: 0, waitUntil: 'domcontentloaded' });
      const cookies = await page.cookies();
      const html = await page.content();
      const matches = html.match(TOKEN_RE);
      browser.close();
      const result = options.stringify ? NxNodeCookie.stringify(cookies) : cookies;
      resolve({ cookies: result, bdstoken: matches[1] });
    });
  });
};
